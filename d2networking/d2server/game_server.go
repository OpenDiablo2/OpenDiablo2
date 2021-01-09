package d2server

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"sync"
	"time"

	"github.com/robertkrimen/otto"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2server/d2tcpclientconnection"
	"github.com/OpenDiablo2/OpenDiablo2/d2script"
)

const logPrefix = "Game Server"

const (
	port                   = "6669"
	chunkSize          int = 4096 // nolint:deadcode,unused,varcheck // WIP
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
	packetManagerChan chan ReceivedPacket
	heroStateFactory  *d2hero.HeroStateFactory

	*d2util.Logger
}

// ReceivedPacket encapsulates the data necessary for the packet manager goroutine to process data from clients.
// The packet manager needs to know who sent the data, in addition to the data itself.
type ReceivedPacket struct {
	Client ClientConnection
	Packet d2netpacket.NetPacket
}

// NewGameServer builds a new GameServer that can be started
//
// ctx: required context item
// networkServer: true = 0.0.0.0 | false = 127.0.0.1
// maxConnections (default: 8): maximum number of TCP connections allowed open
func NewGameServer(asset *d2asset.AssetManager,
	networkServer bool,
	l d2util.LogLevel,
	maxConnections ...int) (*GameServer,
	error) {
	if len(maxConnections) == 0 {
		maxConnections = []int{8}
	}

	heroStateFactory, err := d2hero.NewHeroStateFactory(asset)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	gameServer := &GameServer{
		ctx:               ctx,
		cancel:            cancel,
		asset:             asset,
		connections:       make(map[string]ClientConnection),
		networkServer:     networkServer,
		maxConnections:    maxConnections[0],
		packetManagerChan: make(chan ReceivedPacket),
		mapEngines:        make([]*d2mapengine.MapEngine, 0),
		scriptEngine:      d2script.CreateScriptEngine(),
		seed:              time.Now().UnixNano(),
		heroStateFactory:  heroStateFactory,
	}

	gameServer.Logger = d2util.NewLogger()
	gameServer.Logger.SetPrefix(logPrefix)
	gameServer.Logger.SetLevel(l)

	mapEngine := d2mapengine.CreateMapEngine(l, asset)
	mapEngine.SetSeed(gameServer.seed)
	mapEngine.ResetMap(d2enum.RegionAct1Town, 100, 100)

	mapGen, err := d2mapgen.NewMapGenerator(asset, l, mapEngine)
	if err != nil {
		return nil, err
	}

	mapGen.GenerateAct1Overworld()

	gameServer.mapEngines = append(gameServer.mapEngines, mapEngine)

	gameServer.scriptEngine.AddFunction("getMapEngines", func(call otto.FunctionCall) otto.Value {
		val, err := gameServer.scriptEngine.ToValue(gameServer.mapEngines)
		if err != nil {
			gameServer.Error(err.Error())
		}
		return val
	})

	return gameServer, nil
}

// Start essentially starts all of the game server go routines as well as begins listening for connection. This will
// return an error if it is unable to bind to a socket.
func (g *GameServer) Start() error {
	listenerAddress := "127.0.0.1:" + port
	if g.networkServer {
		listenerAddress = "0.0.0.0:" + port
	}

	g.Infof("Starting Game Server @ %s\n", listenerAddress)

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
				select {
				case <-g.ctx.Done():
					// this error was just a result of the server closing, don't worry about it
				default:
					g.Errorf("Unable to accept connection: %s", err)
				}

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
	g.connections = make(map[string]ClientConnection)

	if err := g.listener.Close(); err != nil {
		g.Errorf("failed to close the listener %s, err: %v\n", g.listener.Addr(), err)
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
			err := g.OnPacketReceived(p.Client, p.Packet)
			if err != nil {
				g.Errorf("failed to handle packet received from client %s: %v", p.Client.GetUniqueID(), err)
			}
		}
	}
}

func (g *GameServer) sendPacketToClients(packet d2netpacket.NetPacket) {
	for _, c := range g.connections {
		if err := c.SendPacketToClient(packet); err != nil {
			g.Errorf("GameServer: error sending packet: %s to client %s: %s", packet.PacketType, c.GetUniqueID(), err)
		}
	}
}

// handleConnection accepts an individual connection and starts pooling for new packets. It is recommended this is called
// via Go Routine. Context should be a property of the GameServer Struct.
func (g *GameServer) handleConnection(conn net.Conn) {
	var (
		connected int
		client    ClientConnection
	)

	g.Infof("Accepting connection: %s\n", conn.RemoteAddr().String())

	defer func() {
		if err := conn.Close(); err != nil {
			g.Errorf("failed to close the connection: %s\n", conn.RemoteAddr())
		}
	}()

	decoder := json.NewDecoder(conn)

	for {
		var packet d2netpacket.NetPacket

		err := decoder.Decode(&packet)
		if err != nil {
			switch err {
			case io.EOF:
				break // the other side closed the connection
			default:
				g.Error(err.Error())
			}

			return // allow the connection to close
		}

		// If this is the first packet we are seeing from this specific connection we first need to see if the client
		// is sending a valid request. If this is a valid request, we will register it and flip the connected switch
		// to.
		if connected == 0 {
			if packet.PacketType != d2netpackettype.PlayerConnectionRequest {
				g.Infof("Closing connection with %s: did not receive new player connection request...", conn.RemoteAddr().String())
			}

			if client, err = g.registerConnection(packet.PacketData, conn); err != nil {
				return
			}

			connected = 1
		}

		select {
		case <-g.ctx.Done():
			return
		default:
			g.packetManagerChan <- ReceivedPacket{
				Client: client,
				Packet: packet,
			}
		}
	}
}

// registerConnection accepts a PlayerConnectionRequestPacket and thread safely updates the connection pool
//
// Errors:
// - errServerFull
// - errPlayerAlreadyExists
func (g *GameServer) registerConnection(b []byte, conn net.Conn) (ClientConnection, error) {
	var client ClientConnection

	g.Lock()
	defer g.Unlock()

	// check to see if the server is full
	if len(g.connections) >= g.maxConnections {
		sf, serverFullErr := d2netpacket.CreateServerFullPacket()
		if serverFullErr != nil {
			g.Errorf("ServerFullPacket: %v", serverFullErr)
		}

		msf, marshalServerFullErr := d2netpacket.MarshalPacket(sf)
		if marshalServerFullErr != nil {
			g.Errorf("MarshalPacket: %v", marshalServerFullErr)
		}

		_, errServerFullPacket := conn.Write(msf)
		g.Warningf("%v", errServerFullPacket)

		return client, errServerFull
	}

	// if it is not full, unmarshal the playerConnectionRequest
	packet, err := d2netpacket.UnmarshalPlayerConnectionRequest(b)
	if err != nil {
		g.Errorf("Failed to unmarshal PlayerConnectionRequest: %s\n", err)
	}

	// check to see if the player is already registered
	if _, ok := g.connections[packet.ID]; ok {
		g.Errorf("%v", errPlayerAlreadyExists)
		return client, errPlayerAlreadyExists
	}

	// Client a new TCP Client Connection and add it to the connections map
	client = d2tcpclientconnection.CreateTCPClientConnection(conn, packet.ID)
	client.SetPlayerState(packet.PlayerState)

	g.OnClientConnected(client)

	return client, nil
}

// OnClientConnected initializes the given ClientConnection. It sends the
// following packets to the newly connected client: UpdateServerInfoPacket,
// GenerateMapPacket, AddPlayerPacket.
//
// It also sends AddPlayerPackets for each other player entity to the new
// player and vice versa, so all player entities exist on all clients.
//
// For more information, see d2networking.d2netpacket.
func (g *GameServer) OnClientConnected(client ClientConnection) {
	// Temporary position hack --------------------------------------------
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/829
	sx, sy := g.mapEngines[0].GetStartPosition()
	clientPlayerState := client.GetPlayerState()
	clientPlayerState.X = sx
	clientPlayerState.Y = sy
	// --------------------------------------------------------------------

	g.Infof("Client connected with an id of %s", client.GetUniqueID())
	g.connections[client.GetUniqueID()] = client

	g.handleClientConnection(client, sx, sy)
}

func (g *GameServer) handleClientConnection(client ClientConnection, x, y float64) {
	usi, err := d2netpacket.CreateUpdateServerInfoPacket(g.seed, client.GetUniqueID())
	if err != nil {
		g.Errorf("UpdateServerInfoPacket: %v", err)
	}

	err = client.SendPacketToClient(usi)
	if err != nil {
		g.Errorf("GameServer: error sending UpdateServerInfoPacket to client %s: %s", client.GetUniqueID(), err)
	}

	gmp, err := d2netpacket.CreateGenerateMapPacket(d2enum.RegionAct1Town)
	if err != nil {
		g.Errorf("GenerateMapPacket: %v", err)
	}

	err = client.SendPacketToClient(gmp)
	if err != nil {
		g.Errorf("GameServer: error sending GenerateMapPacket to client %s: %s", client.GetUniqueID(), err)
	}

	playerState := client.GetPlayerState()

	// these are in subtiles
	playerX := int(x*subtilesPerTile) + middleOfTileOffset
	playerY := int(y*subtilesPerTile) + middleOfTileOffset

	d2hero.HydrateSkills(playerState.Skills, g.asset)

	createPlayerPacket, err := d2netpacket.CreateAddPlayerPacket(
		client.GetUniqueID(),
		playerState.HeroName,
		playerX,
		playerY,
		playerState.HeroType,
		playerState.Stats,
		playerState.Skills,
		playerState.Equipment,
		playerState.LeftSkill,
		playerState.RightSkill,
		playerState.Gold,
	)
	if err != nil {
		g.Errorf("AddPlayerPacket: %v", err)
	}

	for _, connection := range g.connections {
		err := connection.SendPacketToClient(createPlayerPacket)
		if err != nil {
			g.Errorf("GameServer: error sending %T to client %s: %s", createPlayerPacket, connection.GetUniqueID(), err)
		}

		if connection.GetUniqueID() == client.GetUniqueID() {
			continue
		}

		conPlayerState := connection.GetPlayerState()
		playerX := int(conPlayerState.X*subtilesPerTile) + middleOfTileOffset
		playerY := int(conPlayerState.Y*subtilesPerTile) + middleOfTileOffset
		app, err := d2netpacket.CreateAddPlayerPacket(
			connection.GetUniqueID(),
			conPlayerState.HeroName,
			playerX,
			playerY,
			conPlayerState.HeroType,
			conPlayerState.Stats,
			conPlayerState.Skills,
			conPlayerState.Equipment,
			conPlayerState.LeftSkill,
			conPlayerState.RightSkill,
			conPlayerState.Gold,
		)

		if err != nil {
			g.Errorf("AddPlayerPacket: %v", err)
		}

		err = client.SendPacketToClient(app)

		if err != nil {
			g.Errorf("GameServer: error sending CreateAddPlayerPacket to client %s: %s", connection.GetUniqueID(), err)
		}
	}
}

// OnClientDisconnected removes the given client from the list
// of client connections.
// If this client was the host, disconnects all clients and kills GameServer.
func (g *GameServer) OnClientDisconnected(client ClientConnection) {
	g.Infof("Client disconnected with an id of %s", client.GetUniqueID())
	delete(g.connections, client.GetUniqueID())

	if client.GetConnectionType() == d2clientconnectiontype.Local {
		g.Info("Host disconnected, game server shuting down")

		serverClosed, err := d2netpacket.CreateServerClosedPacket()
		if err != nil {
			g.Errorf("failed to generate ServerClosed packet after host disconnected: %s", err)
		} else {
			g.sendPacketToClients(serverClosed)
		}

		g.Stop()
	}
}

// OnPacketReceived is called when a packet has been received from a remote client,
// and by the local client to 'send' a packet to the server,
// nolint:gocyclo // switch statement on packet type makes sense, no need to change
func (g *GameServer) OnPacketReceived(client ClientConnection, packet d2netpacket.NetPacket) error {
	if g == nil {
		return errors.New("game server is nil")
	}

	switch packet.PacketType {
	case d2netpackettype.MovePlayer:
		movePacket, err := d2netpacket.UnmarshalMovePlayer(packet.PacketData)
		if err != nil {
			return err
		}

		playerState := g.connections[client.GetUniqueID()].GetPlayerState()
		playerState.X = movePacket.DestX
		playerState.Y = movePacket.DestY

		g.sendPacketToClients(packet)
	case d2netpackettype.CastSkill, d2netpackettype.SpawnItem:
		g.sendPacketToClients(packet)
	case d2netpackettype.SavePlayer:
		savePacket, err := d2netpacket.UnmarshalSavePlayer(packet.PacketData)
		if err != nil {
			return err
		}

		playerState := g.connections[client.GetUniqueID()].GetPlayerState()
		playerState.LeftSkill = savePacket.Player.LeftSkill.Shallow.SkillID
		playerState.RightSkill = savePacket.Player.RightSkill.Shallow.SkillID
		playerState.Stats = savePacket.Player.Stats
		playerState.Act = savePacket.Player.Act
		playerState.Difficulty = savePacket.Difficulty

		err = g.heroStateFactory.Save(playerState)
		if err != nil {
			g.Errorf("GameServer: error saving saving Player: %s", err)
		}
	case d2netpackettype.PlayerConnectionRequest:
		break // prevent log message. these are handled by handleConnection
	case d2netpackettype.PlayerDisconnectionNotification:
		g.sendPacketToClients(packet)
		g.OnClientDisconnected(client)
	default:
		g.Warningf("GameServer: received unknown packet %s", packet.PacketType)
	}

	return nil
}
