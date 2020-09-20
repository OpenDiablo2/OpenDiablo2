package d2netpacket

import (
	"encoding/json"
	"log"
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
	serverClosed := ServerClosedPacket{
		TS: time.Now(),
	}
	b, err := json.Marshal(serverClosed)
	if err != nil {
		log.Print(err)
	}

	return NetPacket{
		PacketType: d2netpackettype.ServerClosed,
		PacketData: b,
	}
}

func UnmarshalServerClosed(packet []byte) (ServerClosedPacket, error) {
	var resp ServerClosedPacket

	if err := json.Unmarshal(packet, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
