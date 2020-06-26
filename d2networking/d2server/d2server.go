package d2server

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"io"
	"log"
	"net"
	"strings"
	"time"

	// "github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapgen"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"

	// "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	packet "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	packettype "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"github.com/OpenDiablo2/OpenDiablo2/d2script"
	"github.com/robertkrimen/otto"
)

var singletonServer *GameServer

func advance() {
	now := d2common.Now()
	elapsed := now - singletonServer.lastAdvance
	singletonServer.realm.Advance(elapsed)
	singletonServer.lastAdvance = now
}

func Create(openNetworkServer bool) {
	log.Print("Creating GameServer")
	if singletonServer != nil {
		return
	}

	config := d2config.Get()
	maxConnections := config.MaxConnections
	seed := time.Now().UnixNano()

	singletonServer = &GameServer{
		clientConnections: make(map[string]ClientConnection),
		realm:             &d2mapengine.MapRealm{},
		scriptEngine:      d2script.CreateScriptEngine(),
		seed:              seed,
		maxClients:        maxConnections,
		lastAdvance:       d2common.Now(),
	}

	singletonServer.realm.Init(seed)
	singletonServer.manager = CreateConnectionManager(singletonServer)

	// mapEngine := d2mapengine.CreateMapEngine()
	// mapEngine.SetSeed(singletonServer.seed)
	// TODO: Mapgen - Needs levels.txt stuff
	// mapEngine.ResetMap(d2enum.RegionAct1Town, 100, 100)
	// d2mapgen.GenerateAct1Overworld(mapEngine)
	// singletonServer.mapEngines = append(singletonServer.mapEngines, mapEngine)

	addScriptEngineFunctions()

	if openNetworkServer {
		createNetworkServer()
	}
}

func addScriptEngineFunctions() {
	singletonServer.scriptEngine.AddFunction("getMapEngines", ottoTestFunc)
}

func ottoTestFunc(call otto.FunctionCall) otto.Value {
	val, err := singletonServer.scriptEngine.ToValue(singletonServer.realm)
	if err != nil {
		fmt.Print(err.Error())
	}
	return val
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
	srv := singletonServer
	for srv.running {
		advance()

		_, addr, err := srv.udpConnection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Socket error: %s\n", err)
			continue
		}
		buff := bytes.NewBuffer(buffer)
		packetTypeId, err := buff.ReadByte()
		packetType := packettype.NetPacketType(packetTypeId)
		reader, err := gzip.NewReader(buff)
		sb := new(strings.Builder)
		io.Copy(sb, reader)
		stringData := sb.String()

		switch packetType {
		case packettype.PlayerConnectionRequest:
			srv.handlePlayerConnRequest(addr, stringData)
		case packettype.MovePlayer:
			srv.handleMovePlayer(addr, stringData)
		case packettype.Pong:
			srv.handlePong(addr, stringData)
		case packettype.ServerClosed:
			srv.manager.Shutdown()
		case packettype.PlayerDisconnectionNotification:
			srv.handlePlayerDisconnectNotification(stringData)
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
	srv := singletonServer
	realm := srv.realm
	seed := srv.seed
	state := client.GetPlayerState()

	actId := state.Act
	levelId := d2datadict.GetFirstLevelIdByActId(actId)
	engine := realm.GetMapEngine(actId, levelId)

	// params for AddPlayer packet, of new player
	id := client.GetUniqueId()

	name := state.HeroName
	hero := state.HeroType
	stats := *state.Stats
	equip := state.Equipment
	x, y := engine.GetStartPosition()
	state.X = x
	state.Y = y

	infoPacket := packet.CreateUpdateServerInfoPacket(seed, id)
	mapgenPacket := packet.CreateGenerateMapPacket(actId, levelId)

	// UNSURE: maybe we should pass a struct instead of all of these args
	addNew := packet.CreateAddPlayerPacket(
		id, name, actId, int(x*5), int(y*5), hero, stats, equip,
	)

	srv.clientConnections[id] = client

	client.SendPacketToClient(infoPacket)
	client.SendPacketToClient(mapgenPacket)

	log.Printf("Client connected with an id of %s", id)
	realm.AddPlayer(id, state.Act)

	// for each connection, send the AddPlayer packet for the new player
	for _, connection := range srv.clientConnections {
		conId := connection.GetUniqueId()
		connection.SendPacketToClient(addNew)

		if conId == id {
			continue
		}

		// send an AddPlayer for existing connections to the new player
		conState := connection.GetPlayerState()
		conActId := conState.Act
		conName := conState.HeroName
		conHero := conState.HeroType
		conEquip := conState.Equipment
		conStats := *conState.Stats
		conX, conY := 0, 0

		addExisting := packet.CreateAddPlayerPacket(
			conId, conName, conActId, conX, conY, conHero, conStats, conEquip,
		)

		client.SendPacketToClient(addExisting)
	}

}

func OnClientDisconnected(client ClientConnection) {
	log.Printf("Client disconnected with an id of %s", client.GetUniqueId())
	clientId := client.GetUniqueId()
	delete(singletonServer.clientConnections, clientId)
	singletonServer.realm.RemovePlayer(clientId)
}

func OnPacketReceived(client ClientConnection, p packet.NetPacket) error {
	switch p.PacketType {
	case packettype.MovePlayer:
		// TODO: This needs to be verified on the server (here) before sending to other clients....
		// TODO: Hacky, this should be updated in realtime ----------------
		// TODO: Verify player id
		playerState := singletonServer.clientConnections[client.GetUniqueId()].GetPlayerState()
		playerState.X = p.PacketData.(packet.MovePlayerPacket).DestX
		playerState.Y = p.PacketData.(packet.MovePlayerPacket).DestY
		// ----------------------------------------------------------------
		for _, player := range singletonServer.clientConnections {
			player.SendPacketToClient(p)
		}
	}
	return nil
}
