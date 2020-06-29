package d2netpacket

import (
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// ServerClosedPacket contains the current time. It is sent by the server
// to inform a client that the server has shut down.
type ServerClosedPacket struct {
	TS time.Time `json:"ts"`
}

// CreateServerClosedPacket returns a NetPacket which declares a
// ServerClosedPacket with the current time.
func CreateServerClosedPacket() NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.ServerClosed,
		PacketData: ServerClosedPacket{
			TS: time.Now(),
		},
	}
}
