package d2networking

import "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"

// ClientListener is an interface used to pass packet data from
// ClientConnections to GameServer and GameClient.
type ClientListener interface {
	OnPacketReceived(packet d2netpacket.NetPacket) error
}
