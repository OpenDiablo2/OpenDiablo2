package d2remoteclient

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	uuid "github.com/satori/go.uuid"
)

type RemoteClientConnection struct {
	clientListener d2networking.ClientListener
	uniqueId       string
}

func (l RemoteClientConnection) GetUniqueId() string {
	return l.uniqueId
}

func (l RemoteClientConnection) GetConnectionType() string {
	return "Local Client"
}

func (l *RemoteClientConnection) SendPacketToClient(packet d2netpacket.NetPacket) error {
	return l.clientListener.OnPacketReceived(packet)
}

func Create() *RemoteClientConnection {
	result := &RemoteClientConnection{
		uniqueId: uuid.NewV4().String(),
	}

	return result
}

func (l *RemoteClientConnection) Open(connectionString string) error {
	// TODO: Connect to the server
	return nil
}

func (l *RemoteClientConnection) Close() error {
	// TODO: Disconnect from the server
	return nil
}

func (l *RemoteClientConnection) SendPacketToServer(packet d2netpacket.NetPacket) error {
	// TODO: Send to the server
	return nil
}

func (l *RemoteClientConnection) SetClientListener(listener d2networking.ClientListener) {
	l.clientListener = listener
}
