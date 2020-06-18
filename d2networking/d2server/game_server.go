package d2server

import (
	"fmt"
	"log"
	"net"

	"github.com/robertkrimen/otto"

	"github.com/OpenDiablo2/OpenDiablo2/d2script"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gamestate"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

type GameServer struct {
	gameState         *d2gamestate.GameState
	clientConnections map[string]ClientConnection
	mapEngines        []*d2map.MapEngine
	scriptEngine      *d2script.ScriptEngine
	udpConnection     *net.UDPConn
	running           bool
}

var singletonServer *GameServer

func Create(gameStatePath string, openNetworkServer bool) {
	log.Print("Creating GameServer")
	if singletonServer != nil {
		return
	}

	singletonServer = &GameServer{
		clientConnections: make(map[string]ClientConnection),
		mapEngines:        make([]*d2map.MapEngine, 0),
		gameState:         d2gamestate.LoadGameState(gameStatePath),
		scriptEngine:      d2script.CreateScriptEngine(),
	}

	mapEngine := d2map.CreateMapEngine()
	mapEngine.GenerateMap(d2enum.RegionAct1Town, 1, 0, false)
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

	go runNetworkServer()
	log.Print("Network server has been started")
}

func runNetworkServer() {
	buffer := make([]byte, 4096)
	for singletonServer.running {
		n, addr, err := singletonServer.udpConnection.ReadFromUDP(buffer)
		if err != nil {
			continue
		}
		fmt.Printf("%s: %s", addr.IP.String(), string(buffer[0:n-1]))
	}
}

func Run() {
	log.Print("Starting GameServer")
	singletonServer.running = true
	singletonServer.scriptEngine.RunScript("scripts/server/server.js")
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
	log.Printf("Client connected with an id of %s", client.GetUniqueId())
	singletonServer.clientConnections[client.GetUniqueId()] = client
	client.SendPacketToClient(d2netpacket.CreateUpdateServerInfoPacket(singletonServer.gameState.Seed, client.GetUniqueId()))
	client.SendPacketToClient(d2netpacket.CreateGenerateMapPacket(d2enum.RegionAct1Town, 1, 0))

	// TODO: This needs to use a real method of loading characters instead of cloning the 'save file character'
	sx, sy := singletonServer.mapEngines[0].GetStartPosition() // TODO: Another temporary hack
	createPlayerPacket := d2netpacket.CreateAddPlayerPacket(client.GetUniqueId(), int(sx*5)+3, int(sy*5)+3,
		singletonServer.gameState.HeroType, singletonServer.gameState.Equipment)
	for _, connection := range singletonServer.clientConnections {
		connection.SendPacketToClient(createPlayerPacket)
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
		for _, player := range singletonServer.clientConnections {
			player.SendPacketToClient(packet)
		}
		break
	}
	return nil
}
