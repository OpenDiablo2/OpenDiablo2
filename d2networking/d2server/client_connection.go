package d2server

import "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"

type ClientConnection interface {
	GetUniqueId() string
	GetConnectionType() string
	SendPacketToClient(packet d2netpacket.NetPacket) error
}
