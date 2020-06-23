package d2netpacket

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"time"
)

type ServerClosedPacket struct {
	TS time.Time `json:"ts"`
}

func CreateServerClosedPacket() NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.ServerClosed,
		PacketData: ServerClosedPacket{
			TS: time.Now(),
		},
	}
}
