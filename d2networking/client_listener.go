package d2networking

import "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"

// ClientListener is an interface used to call GameClient.OnPacketReceived
// from LocalClientConnection and RemoteClientConnection.
type ClientListener interface {
	OnPacketReceived(packet d2netpacket.NetPacket) error
}
