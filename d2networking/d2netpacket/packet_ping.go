package d2netpacket

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"time"
)

type PingPacket struct {
	TS time.Time `json:"ts"`
}

func CreatePingPacket() NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.Ping,
		PacketData: PingPacket{
			TS: time.Now(),
		},
	}
}
