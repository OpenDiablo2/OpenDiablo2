package d2server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/robertkrimen/otto"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2server/d2tcpclientconnection"
	"github.com/OpenDiablo2/OpenDiablo2/d2script"
)

const (
	port                   = "6669"
	chunkSize          int = 4096
	subtilesPerTile        = 5
	middleOfTileOffset     = 3
)

var (
	errPlayerAlreadyExists = errors.New("player already exists")
	errServerFull          = errors.New("server full") // Server currently at maximum TCP connections
)

// GameServer manages a copy of the map and entities as well as manages packet routing and connections.
// It can accept connections from localhost as well remote clients. It can also be started in a standalone mode.
type GameServer struct {
	sync.RWMutex
	connections       map[string]ClientConnection
	listener          net.Listener
	networkServer     bool
	ctx               context.Context
	cancel            context.CancelFunc
	asset             *d2asset.AssetManager
	mapEngines        []*d2mapengine.MapEngine
	scriptEngine      *d2script.ScriptEngine
	seed              int64
	maxConnections    int
	packetManagerChan chan []byte
}

//nolint:gochecknoglobals // currently singleton by design
var singletonServer *GameServer

// NewGameServer builds a new GameServer that can be started
//
// ctx: required context item
// networkServer: true = 0.0.0.0 | false = 127.0.0.1
// maxConnections (default: 8): maximum number of TCP connections allowed open
func NewGameServer(asset *d2asset.AssetManager, networkServer bool,
	maxConnections ...int) (*GameServer,
	error) {
	if len(maxConnections) == 0 {
		maxConnections = []int{8}
	}

	ctx, cancel := context.WithCancel(context.Background())

	gameServer := &GameServer{
		ctx:               ctx,
		cancel:            cancel,
		asset:             asset,
		connections:       make(map[string]ClientConnection),
		networkServer:     networkServer,
		maxConnections:    maxConnections[0],
		packetManagerChan: make(chan []byte),
		mapEngines:        make([]*d2mapengine.MapEngine, 0),
		scriptEngine:      d2script.CreateScriptEngine(),
		seed:              time.Now().UnixNano(),
	}

	// TODO: In order to support dedicated mode we need to load the levels txt and files. Revisit this once this we can
	//   load files independent of the app.
	mapEngine := d2mapengine.CreateMapEngine(asset)
	mapEngine.SetSeed(gameServer.seed)
	mapEngine.ResetMap(d2enum.RegionAct1Town, 100, 100) // TODO: Mapgen - Needs levels.txt stuff

	mapGen, err := d2mapgen.NewMapGenerator(asset, mapEngine)
	if err != nil {
		return nil, err
	}

	mapGen.GenerateAct1Overworld()

	gameServer.mapEngines = append(gameServer.mapEngines, mapEngine)

	gameServer.scriptEngine.AddFunction("getMapEngines", func(call otto.FunctionCall) otto.Value {
		val, err := gameServer.scriptEngine.ToValue(singletonServer.mapEngines)
		if err != nil {
			fmt.Print(err.Error())
		}
		return val
	})

	// TODO: Temporary hack to work around local connections. Possible that we can move away from the singleton pattern here
	// 		but for now this will work.
	singletonServer = gameServer

	return gameServer, nil
}

// Start essentially starts all of the game server go routines as well as begins listening for connection. This will
// return an error if it is unable to bind to a socket.
func (g *GameServer) Start() error {
	listenerAddress := "127.0.0.1:" + port
	if g.networkServer {
		listenerAddress = "0.0.0.0:" + port
	}

	log.Printf("Starting Game Server @ %s\n", listenerAddress)

	l, err := net.Listen("tcp4", listenerAddress)
	if err != nil {
		return err
	}

	g.listener = l

	go g.packetManager()

	go func() {
		for {
			c, err := g.listener.Accept()
			if err != nil {
				log.Printf("Unable to accept connection: %s\n", err)
				return
			}

			go g.handleConnection(c)
		}
	}()

	return nil
}

// Stop stops the game server
func (g *GameServer) Stop() {
	g.Lock()
	g.cancel()

	if err := g.listener.Close(); err != nil {
		log.Printf("failed to close the listener %s, err: %v\n", g.listener.Addr(), err)
	}
}

// packetManager is meant to be started as a Goroutine and is used to manage routing of packets to clients.
func (g *GameServer) packetManager() {
	defer close(g.packetManagerChan)

	for {
		select {
		// If the server is stopped we need to clean up the packet manager goroutine
		case <-g.ctx.Done():
			return
		case p := <-g.packetManagerChan:
			switch d2netpacket.InspectPacketType(p) {
			case d2netpackettype.PlayerConnectionRequest:
				player, err := d2netpacket.UnmarshalNetPacket(p)
				if err != nil {
					log.Printf("Unable to unmarshal PlayerConnectionRequestPacket: %s\n", err)
				}

				g.sendPacketToClients(player)
			case d2netpackettype.MovePlayer:
				move, err := d2netpacket.UnmarshalNetPacket(p)
				if err != nil {
					log.Println(err)
					continue
				}

				g.sendPacketToClients(move)
			case d2netpackettype.SpawnItem:
				item, err := d2netpacket.UnmarshalNetPacket(p)
				if err != nil {
					log.Println(err)
					continue
				}

				g.sendPacketToClients(item)
			case d2netpackettype.ServerClosed:
				g.Stop()
			}
		}
	}
}

func (g *GameServer) sendPacketToClients(packet d2netpacket.NetPacket) {
	for _, c := range g.connections {
		if err := c.SendPacketToClient(packet); err != nil {
			log.Printf("GameServer: error sending packet: %s to client %s: %s", packet.PacketType, c.GetUniqueID(), err)
		}
	}
}

// handleConnection accepts an individual connection and starts pooling for new packets. It is recommended this is called
// via Go Routine. Context should be a property of the GameServer Struct.
func (g *GameServer) handleConnection(conn net.Conn) {
	var connected int

	var packet d2netpacket.NetPacket

	log.Printf("Accepting connection: %s\n", conn.RemoteAddr().String())

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("failed to close the connection: %s\n", conn.RemoteAddr())
		}
	}()

	decoder := json.NewDecoder(conn)

	for {
		err := decoder.Decode(&packet)
		if err != nil {
			log.Println(err)
			return // exit this connection as we could not read the first packet
		}

		// If this is the first packet we are seeing from this specific connection we first need to see if the client
		// is sending a valid request. If this is a valid request, we will register it and flip the connected switch
		// to.
		if connected == 0 {
			if packet.PacketType != d2netpackettype.PlayerConnectionRequest {
				log.Printf("Closing connection with %s: did not receive new player connection request...\n", conn.RemoteAddr().String())
			}

			// TODO: I do not think this error check actually works. Need to retrofit with Errors.Is().
			if err := g.registerConnection(packet.PacketData, conn); err != nil {
				switch err {
				case errServerFull: // Server is currently full and not accepting new connections.
					// TODO: Need to create a new Server Full packet to return to clients.
					log.Println(err)
					return
				case errPlayerAlreadyExists: // Player is already registered and did not disconnection correctly.
					log.Println(err)
					return
				}
			}

			connected = 1
		}

		select {
		case <-g.ctx.Done():
			return
		default:
			g.packetManagerChan <- packet.PacketData
		}
	}
}

// registerConnection accepts a PlayerConnectionRequestPacket and thread safely updates the connection pool
//
// Errors:
// - errServerFull
// - errPlayerAlreadyExists
func (g *GameServer) registerConnection(b []byte, conn net.Conn) error {
	g.Lock()

	// check to see if the server is full
	if len(g.connections) >= g.maxConnections {
		return errServerFull
	}

	// if it is not full, unmarshal the playerConnectionRequest
	packet, err := d2netpacket.UnmarshalPlayerConnectionRequest(b)
	if err != nil {
		log.Printf("Failed to unmarshal PlayerConnectionRequest: %s\n", err)
	}

	// check to see if the player is already registered
	if _, ok := g.connections[packet.ID]; ok {
		return errPlayerAlreadyExists
	}

	// Client a new TCP Client Connection and add it to the connections map
	client := d2tcpclientconnection.CreateTCPClientConnection(conn, packet.ID)
	client.SetPlayerState(packet.PlayerState)
	log.Printf("Client connected with an id of %s", client.GetUniqueID())
	g.connections[client.GetUniqueID()] = client

	// Temporary position hack --------------------------------------------
	sx, sy := g.mapEngines[0].GetStartPosition() // TODO: Another temporary hack
	clientPlayerState := client.GetPlayerState()
	clientPlayerState.X = sx
	clientPlayerState.Y = sy
	// ---------

	// This really should be deferred however to much time will be spend holding a lock when we attempt to send a packet
	g.Unlock()

	handleClientConnection(g, client, sx, sy)

	return nil
}

// OnClientConnected initializes the given ClientConnection. It sends the
// following packets to the newly connected client: UpdateServerInfoPacket,
// GenerateMapPacket, AddPlayerPacket.
//
// It also sends AddPlayerPackets for each other player entity to the new
// player and vice versa, so all player entities exist on all clients.
//
// For more information, see d2networking.d2netpacket.
func OnClientConnected(client ClientConnection) {
	// Temporary position hack --------------------------------------------
	sx, sy := singletonServer.mapEngines[0].GetStartPosition() // TODO: Another temporary hack
	clientPlayerState := client.GetPlayerState()
	clientPlayerState.X = sx
	clientPlayerState.Y = sy
	// --------------------------------------------------------------------

	log.Printf("Client connected with an id of %s", client.GetUniqueID())
	singletonServer.connections[client.GetUniqueID()] = client

	handleClientConnection(singletonServer, client, sx, sy)
}

func handleClientConnection(gameServer *GameServer, client ClientConnection, x, y float64) {
	err := client.SendPacketToClient(d2netpacket.CreateUpdateServerInfoPacket(gameServer.seed, client.GetUniqueID()))
	if err != nil {
		log.Printf("GameServer: error sending UpdateServerInfoPacket to client %s: %s", client.GetUniqueID(), err)
	}

	err = client.SendPacketToClient(d2netpacket.CreateGenerateMapPacket(d2enum.RegionAct1Town))
	if err != nil {
		log.Printf("GameServer: error sending GenerateMapPacket to client %s: %s", client.GetUniqueID(), err)
	}

	playerState := client.GetPlayerState()

	// these are in subtiles
	playerX := int(x*subtilesPerTile) + middleOfTileOffset
	playerY := int(y*subtilesPerTile) + middleOfTileOffset

	createPlayerPacket := d2netpacket.CreateAddPlayerPacket(
		client.GetUniqueID(),
		playerState.HeroName,
		playerX,
		playerY,
		playerState.HeroType,
		playerState.Stats,
		playerState.Skills,
		playerState.Equipment,
	)

	for _, connection := range gameServer.connections {
		err := connection.SendPacketToClient(createPlayerPacket)
		if err != nil {
			log.Printf("GameServer: error sending %T to client %s: %s", createPlayerPacket, connection.GetUniqueID(), err)
		}

		if connection.GetUniqueID() == client.GetUniqueID() {
			continue
		}

		conPlayerState := connection.GetPlayerState()
		playerX := int(conPlayerState.X*subtilesPerTile) + middleOfTileOffset
		playerY := int(conPlayerState.Y*subtilesPerTile) + middleOfTileOffset
		err = client.SendPacketToClient(
			d2netpacket.CreateAddPlayerPacket(
				connection.GetUniqueID(),
				conPlayerState.HeroName,
				playerX,
				playerY,
				conPlayerState.HeroType,
				conPlayerState.Stats,
				conPlayerState.Skills,
				conPlayerState.Equipment,
			),
		)

		if err != nil {
			log.Printf("GameServer: error sending CreateAddPlayerPacket to client %s: %s", connection.GetUniqueID(), err)
		}
	}
}

// OnClientDisconnected removes the given client from the list
// of client connections.
func OnClientDisconnected(client ClientConnection) {
	log.Printf("Client disconnected with an id of %s", client.GetUniqueID())
	delete(singletonServer.connections, client.GetUniqueID())
}

// OnPacketReceived is called by the local client to 'send' a packet to the server.
func OnPacketReceived(client ClientConnection, packet d2netpacket.NetPacket) error {
	switch packet.PacketType {
	case d2netpackettype.MovePlayer:
		movePacket, err := d2netpacket.UnmarshalMovePlayer(packet.PacketData)
		if err != nil {
			return err
		}
		// TODO: This needs to be verified on the server (here) before sending to other clients....
		// TODO: Hacky, this should be updated in realtime ----------------
		// TODO: Verify player id
		playerState := singletonServer.connections[client.GetUniqueID()].GetPlayerState()
		playerState.X = movePacket.DestX
		playerState.Y = movePacket.DestY
		// ----------------------------------------------------------------
		for _, player := range singletonServer.connections {
			err := player.SendPacketToClient(packet)
			if err != nil {
				log.Printf("GameServer: error sending %T to client %s: %s", packet, player.GetUniqueID(), err)
			}
		}
	case d2netpackettype.CastSkill:
		for _, player := range singletonServer.connections {
			err := player.SendPacketToClient(packet)
			if err != nil {
				log.Printf("GameServer: error sending %T to client %s: %s", packet, player.GetUniqueID(), err)
			}
		}
	case d2netpackettype.SpawnItem:
		for _, player := range singletonServer.connections {
			err := player.SendPacketToClient(packet)
			if err != nil {
				log.Printf("GameServer: error sending %T to client %s: %s", packet, player.GetUniqueID(), err)
			}
		}
	}

	return nil
}
