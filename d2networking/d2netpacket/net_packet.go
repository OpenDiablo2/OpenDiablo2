package d2netpacket

import "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"

type NetPacket struct {
	PacketType d2netpackettype.NetPacketType `json:"packetType"`
	PacketData interface{}                   `json:"packetData"`
}
