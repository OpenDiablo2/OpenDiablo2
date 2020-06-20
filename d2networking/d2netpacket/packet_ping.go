package d2netpacket

import (
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// PingPacket contains the time at which it was sent. It is sent by the
// server and instructs the client to respond with a Pong packet.
type PingPacket struct {
	TS time.Time `json:"ts"`
}

// CreatePingPacket returns a NetPacket which declares a GenerateMapPacket
// with the the current time.
func CreatePingPacket() NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.Ping,
		PacketData: PingPacket{
			TS: time.Now(),
		},
	}
}
