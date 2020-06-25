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
	singletonServer.udpConnection.SetReadBuffer(4096)
}

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
		io.Copy(sb, reader)
		stringData := sb.String()
		switch packetType {
		case d2netpackettype.PlayerConnectionRequest:
			packetData := d2netpacket.PlayerConnectionRequestPacket{}
			json.Unmarshal([]byte(stringData), &packetData)
			clientConnection := d2udpclientconnection.CreateUDPClientConnection(singletonServer.udpConnection, packetData.Id, addr)
			clientConnection.SetPlayerState(packetData.PlayerState)
			OnClientConnected(clientConnection)
		case d2netpackettype.MovePlayer:
			packetData := d2netpacket.MovePlayerPacket{}
			json.Unmarshal([]byte(stringData), &packetData)
			netPacket := d2netpacket.NetPacket{
				PacketType: packetType,
				PacketData: packetData,
			}

			for _, player := range singletonServer.clientConnections {
				player.SendPacketToClient(netPacket)
			}
		case d2netpackettype.Pong:
			packetData := d2netpacket.PlayerConnectionRequestPacket{}
			json.Unmarshal([]byte(stringData), &packetData)
			singletonServer.manager.Recv(packetData.Id)
		case d2netpackettype.ServerClosed:
			singletonServer.manager.Shutdown()
		case d2netpackettype.PlayerDisconnectionNotification:
			var packet d2netpacket.PlayerDisconnectRequestPacket
			json.Unmarshal([]byte(stringData), &packet)
			log.Printf("Received disconnect: %s", packet.Id)
		}
	}
}

func Run() {
	log.Print("Starting GameServer")
	singletonServer.running = true
	singletonServer.scriptEngine.RunScript("scripts/server/server.js")
	if singletonServer.udpConnection != nil {
		go runNetworkServer()
	}
	log.Print("Network server has been started")
}

func Stop() {
	log.Print("Stopping GameServer")
	singletonServer.running = false
	if singletonServer.udpConnection != nil {
		singletonServer.udpConnection.Close()
	}
}

func Destroy() {
	if singletonServer == nil {
		return
	}
	log.Print("Destroying GameServer")
	Stop()
}

func OnClientConnected(client ClientConnection) {
	// Temporary position hack --------------------------------------------
	sx, sy := singletonServer.mapEngines[0].GetStartPosition() // TODO: Another temporary hack
	clientPlayerState := client.GetPlayerState()
	clientPlayerState.X = sx
	clientPlayerState.Y = sy
	// --------------------------------------------------------------------

	log.Printf("Client connected with an id of %s", client.GetUniqueId())
	singletonServer.clientConnections[client.GetUniqueId()] = client
	client.SendPacketToClient(d2netpacket.CreateUpdateServerInfoPacket(singletonServer.seed, client.GetUniqueId()))
	client.SendPacketToClient(d2netpacket.CreateGenerateMapPacket(d2enum.RegionAct1Town))

	playerState := client.GetPlayerState()
	createPlayerPacket := d2netpacket.CreateAddPlayerPacket(client.GetUniqueId(), playerState.HeroName, int(sx*5)+3, int(sy*5)+3,
		playerState.HeroType, *playerState.Stats, playerState.Equipment)
	for _, connection := range singletonServer.clientConnections {
		connection.SendPacketToClient(createPlayerPacket)
		if connection.GetUniqueId() == client.GetUniqueId() {
			continue
		}

		conPlayerState := connection.GetPlayerState()
		client.SendPacketToClient(d2netpacket.CreateAddPlayerPacket(connection.GetUniqueId(), conPlayerState.HeroName,
			int(conPlayerState.X*5)+3, int(conPlayerState.Y*5)+3, conPlayerState.HeroType, *conPlayerState.Stats, conPlayerState.Equipment))
	}

}

func OnClientDisconnected(client ClientConnection) {
	log.Printf("Client disconnected with an id of %s", client.GetUniqueId())
	delete(singletonServer.clientConnections, client.GetUniqueId())
}

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
			player.SendPacketToClient(packet)
		}
	}
	return nil
}
