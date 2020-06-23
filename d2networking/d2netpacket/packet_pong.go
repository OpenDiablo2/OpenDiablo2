package d2netpacket

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"time"
)

type PongPacket struct {
	ID string    `json:"id"`
	TS time.Time `json:"ts"`
}

func CreatePongPacket(id string) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.Pong,
		PacketData: PongPacket{
			ID: id,
			TS: time.Now(),
		},
	}
}
