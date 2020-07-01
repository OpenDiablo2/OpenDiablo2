package d2server

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2server/d2udpclientconnection"
	"github.com/OpenDiablo2/OpenDiablo2/d2script"
	"github.com/robertkrimen/otto"
)

// GameServer owns the authoritative copy of the map and entities
// It accepts incoming connections from local (host) and remote
// clients.
type GameServer struct {
	sync.RWMutex
	clientConnections map[string]ClientConnection
	manager           *ConnectionManager
	mapEngines        []*d2mapengine.MapEngine
	scriptEngine      *d2script.ScriptEngine
	udpConnection     *net.UDPConn
	seed              int64
	running           bool
}

var singletonServer *GameServer

// Create constructs a new GameServer and assigns it as a singleton. It
// also generates the initial map and entities for the server.
//
// If openNetworkServer is true, the GameServer starts listening for UDP
// packets.
func Create(openNetworkServer bool) {
	log.Print("Creating GameServer")
	if singletonServer != nil {
		return
	}

	singletonServer = &GameServer{
		clientConnections: make(map[string]ClientConnection),
		mapEngines:        make([]*d2mapengine.MapEngine, 0),
		scriptEngine:      d2script.CreateScriptEngine(),
		seed:              time.Now().UnixNano(),
	}

	singletonServer.manager = CreateConnectionManager(singletonServer)

	mapEngine := d2mapengine.CreateMapEngine()
	mapEngine.SetSeed(singletonServer.seed)
	mapEngine.ResetMap(d2enum.RegionAct1Town, 100, 100) // TODO: Mapgen - Needs levels.txt stuff
	d2mapgen.GenerateAct1Overworld(mapEngine)
	singletonServer.mapEngines = append(singletonServer.mapEngines, mapEngine)

	singletonServer.scriptEngine.AddFunction("getMapEngines", func(call otto.FunctionCall) otto.Value {
		val, err := singletonServer.scriptEngine.ToValue(singletonServer.mapEngines)
		if err != nil {
			fmt.Print(err.Error())
		}
		return val
	})

	if openNetworkServer {
		createNetworkServer()
	}
}

func createNetworkServer() {
	s, err := net.ResolveUDPAddr("udp4", "0.0.0.0:6669")
	if err != nil {
		panic(err)
	}

	singletonServer.udpConnection, err = net.ListenUDP("udp4", s)
	if err != nil {
		panic(err)
	}
	err = singletonServer.udpConnection.SetReadBuffer(4096)
	if err != nil {
		log.Print("GameServer: error setting UDP read buffer:", err)
	}
}

// runNetworkServer runs a while loop, reading from the GameServer's UDP
// connection.
func runNetworkServer() {
	buffer := make([]byte, 4096)
	for singletonServer.running {
		_, addr, err := singletonServer.udpConnection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Socket error: %s\n", err)
			continue
		}
		buff := bytes.NewBuffer(buffer)
		packetTypeId, err := buff.ReadByte()
		packetType := d2netpackettype.NetPacketType(packetTypeId)
		reader, err := gzip.NewReader(buff)
		sb := new(strings.Builder)

		// This will throw errors where packets are not compressed. This doesn't
		// break anything, so the gzip.ErrHeader error, is currently ignored to
		// avoid noisy logging.
		written, err := io.Copy(sb, reader)
		if err != nil && err != gzip.ErrHeader {
			log.Printf("GameServer: error copying bytes from %v packet: %s", packetType, err)

		}
		if written == 0 {
			log.Printf("GameServer: empty packet %v packet received", packetType)
			continue
		}

		stringData := sb.String()
		switch packetType {
		case d2netpackettype.PlayerConnectionRequest:
			packetData := d2netpacket.PlayerConnectionRequestPacket{}
			err := json.Unmarshal([]byte(stringData), &packetData)
			if err != nil {
				log.Printf("GameServer: error unmarshalling packet of type %T: %s", packetData, err)
				continue
			}
			clientConnection := d2udpclientconnection.CreateUDPClientConnection(singletonServer.udpConnection, packetData.Id, addr)
			clientConnection.SetPlayerState(packetData.PlayerState)
			OnClientConnected(clientConnection)
		case d2netpackettype.MovePlayer:
			packetData := d2netpacket.MovePlayerPacket{}
			err := json.Unmarshal([]byte(stringData), &packetData)
			if err != nil {
				log.Printf("GameServer: error unmarshalling %T: %s", packetData, err)
				continue
			}
			netPacket := d2netpacket.NetPacket{
				PacketType: packetType,
				PacketData: packetData,
			}

			for _, player := range singletonServer.clientConnections {
				err = player.SendPacketToClient(netPacket)
				if err != nil {
					log.Printf("GameServer: error sending %T to client %s: %s", packetData, player.GetUniqueId(), err)
				}
			}
		case d2netpackettype.Pong:
			packetData := d2netpacket.PlayerConnectionRequestPacket{}
			err := json.Unmarshal([]byte(stringData), &packetData)
			if err != nil {
				log.Printf("GameServer: error unmarshalling packet of type %T: %s", packetData, err)
				continue
			}
			singletonServer.manager.Recv(packetData.Id)
		case d2netpackettype.ServerClosed:
			singletonServer.manager.Shutdown()
		case d2netpackettype.PlayerDisconnectionNotification:
			var packet d2netpacket.PlayerDisconnectRequestPacket
			err := json.Unmarshal([]byte(stringData), &packet)
			if err != nil {
				log.Printf("GameServer: error unmarshalling packet of type %T: %s", packet, err)
				continue
			}
			log.Printf("Received disconnect: %s", packet.Id)
		}
	}
}

// Run sets GameServer.running to true and call runNetworkServer
// in a goroutine.
func Run() {
	log.Print("Starting GameServer")
	singletonServer.running = true
	_, err := singletonServer.scriptEngine.RunScript("scripts/server/server.js")
	if err != nil {
		log.Printf("GameServer: error initializing debug script: %s", err)
	}
	if singletonServer.udpConnection != nil {
		go runNetworkServer()
	}
	log.Print("Network server has been started")
}

// Stop sets GameServer.running to false and closes the
// GameServer's UDP connection.
func Stop() {
	log.Print("Stopping GameServer")
	singletonServer.running = false
	if singletonServer.udpConnection != nil {
		err := singletonServer.udpConnection.Close()
		if err != nil {
			log.Printf("GameServer: error when trying to close UDP connection: %s", err)
		}
	}
}

// Destroy calls Stop() if the server exists.
func Destroy() {
	if singletonServer == nil {
		return
	}
	log.Print("Destroying GameServer")
	Stop()
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

	log.Printf("Client connected with an id of %s", client.GetUniqueId())
	singletonServer.clientConnections[client.GetUniqueId()] = client
	err := client.SendPacketToClient(d2netpacket.CreateUpdateServerInfoPacket(singletonServer.seed, client.GetUniqueId()))
	if err != nil {
		log.Printf("GameServer: error sending UpdateServerInfoPacket to client %s: %s", client.GetUniqueId(), err)
	}
	err = client.SendPacketToClient(d2netpacket.CreateGenerateMapPacket(d2enum.RegionAct1Town))
	if err != nil {
		log.Printf("GameServer: error sending GenerateMapPacket to client %s: %s", client.GetUniqueId(), err)
	}

	playerState := client.GetPlayerState()
	createPlayerPacket := d2netpacket.CreateAddPlayerPacket(client.GetUniqueId(), playerState.HeroName, int(sx*5)+3, int(sy*5)+3,
		playerState.HeroType, *playerState.Stats, playerState.Equipment)
	for _, connection := range singletonServer.clientConnections {
		err := connection.SendPacketToClient(createPlayerPacket)
		if err != nil {
			log.Printf("GameServer: error sending %T to client %s: %s", createPlayerPacket, connection.GetUniqueId(), err)
		}
		if connection.GetUniqueId() == client.GetUniqueId() {
			continue
		}

		conPlayerState := connection.GetPlayerState()
		err = client.SendPacketToClient(d2netpacket.CreateAddPlayerPacket(connection.GetUniqueId(), conPlayerState.HeroName,
			int(conPlayerState.X*5)+3, int(conPlayerState.Y*5)+3, conPlayerState.HeroType, *conPlayerState.Stats, conPlayerState.Equipment))
		if err != nil {
			log.Printf("GameServer: error sending CreateAddPlayerPacket to client %s: %s", connection.GetUniqueId(), err)
		}
	}

}

// OnClientDisconnected removes the given client from the list
// of client connections.
func OnClientDisconnected(client ClientConnection) {
	log.Printf("Client disconnected with an id of %s", client.GetUniqueId())
	delete(singletonServer.clientConnections, client.GetUniqueId())
}

// OnPacketReceived is called by the local client to 'send' a packet to the server.
func OnPacketReceived(client ClientConnection, packet d2netpacket.NetPacket) error {
	switch packet.PacketType {
	case d2netpackettype.MovePlayer:
		// TODO: This needs to be verified on the server (here) before sending to other clients....
		// TODO: Hacky, this should be updated in realtime ----------------
		// TODO: Verify player id
		playerState := singletonServer.clientConnections[client.GetUniqueId()].GetPlayerState()
		playerState.X = packet.PacketData.(d2netpacket.MovePlayerPacket).DestX
		playerState.Y = packet.PacketData.(d2netpacket.MovePlayerPacket).DestY
		// ----------------------------------------------------------------
		for _, player := range singletonServer.clientConnections {
			err := player.SendPacketToClient(packet)
			if err != nil {
				log.Printf("GameServer: error sending %T to client %s: %s", packet, player.GetUniqueId(), err)
			}
		}
	case d2netpackettype.CastSkill:
		for _, player := range singletonServer.clientConnections {
			err := player.SendPacketToClient(packet)
			if err != nil {
				log.Printf("GameServer: error sending %T to client %s: %s", packet, player.GetUniqueId(), err)
			}
		}
	}
	return nil
}
