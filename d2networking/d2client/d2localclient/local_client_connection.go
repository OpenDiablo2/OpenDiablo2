package d2localclient

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2server"
	uuid "github.com/satori/go.uuid"
)

type LocalClientConnection struct {
	clientListener    d2networking.ClientListener
	uniqueId          string
	openNetworkServer bool
}

func (l LocalClientConnection) GetUniqueId() string {
	return l.uniqueId
}

func (l LocalClientConnection) GetConnectionType() string {
	return "Local Client"
}

func (l *LocalClientConnection) SendPacketToClient(packet d2netpacket.NetPacket) error {
	return l.clientListener.OnPacketReceived(packet)
}

func Create(openNetworkServer bool) *LocalClientConnection {
	result := &LocalClientConnection{
		uniqueId:          uuid.NewV4().String(),
		openNetworkServer: openNetworkServer,
	}

	return result
}

func (l *LocalClientConnection) Open(connectionString string, saveFilePath string) error {
	d2server.Create(saveFilePath, l.openNetworkServer)
	go d2server.Run()
	d2server.OnClientConnected(l)
	return nil
}

func (l *LocalClientConnection) Close() error {
	d2server.OnClientDisconnected(l)
	d2server.Destroy()
	return nil
}

func (l *LocalClientConnection) SendPacketToServer(packet d2netpacket.NetPacket) error {
	// TODO: This is going to blow up if the server has ceased to be.
	return d2server.OnPacketReceived(l, packet)
}

func (l *LocalClientConnection) SetClientListener(listener d2networking.ClientListener) {
	l.clientListener = listener
}
