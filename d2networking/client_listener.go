package d2networking

import "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"

type ClientListener interface {
	OnPacketReceived(packet d2netpacket.NetPacket) error
}
