package d2server

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gamestate"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

type GameServer struct {
	gameState         *d2gamestate.GameState
	clientConnections map[string]ClientConnection
}

var singletonServer *GameServer

func Create(gameStatePath string) {
	log.Print("Creating GameServer")
	if singletonServer != nil {
		return
	}

	singletonServer = &GameServer{
		clientConnections: make(map[string]ClientConnection),
		gameState:         d2gamestate.LoadGameState(gameStatePath),
	}

}

func Run() {
	log.Print("Starting GameServer")
}

func Stop() {
	log.Print("Stopping GameServer")
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
	createPlayerPacket := d2netpacket.CreateAddPlayerPacket(client.GetUniqueId(), 8, 8,
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
	return nil
}
