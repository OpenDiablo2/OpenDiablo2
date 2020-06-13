package d2server

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"log"
)

type GameServer struct {
	clientConnections map[string]ClientConnection
}

var singletonServer * GameServer

func Create() {
	log.Print("Creating GameServer")
	if singletonServer != nil {
		return
	}

	singletonServer = &GameServer{
		clientConnections: make(map[string]ClientConnection),
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
}

func OnClientDisconnected(client ClientConnection) {
	log.Printf("Client disconnected with an id of %s", client.GetUniqueId())
	delete(singletonServer.clientConnections, client.GetUniqueId())
}

func OnPacketReceived(client ClientConnection, packet d2netpacket.NetPacket) error {
	return nil
}
