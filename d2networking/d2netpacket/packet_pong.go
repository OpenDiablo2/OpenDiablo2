package d2netpacket

import (
	"encoding/json"
	"log"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// PongPacket contains the time at which it was sent and the ID of the
// client. It is sent by the client in response to a Pong packet.
type PongPacket struct {
	ID string    `json:"id"`
	TS time.Time `json:"ts"`
}

// CreatePongPacket returns a NetPacket which declares a PongPacket with
// the current time and given ID.
func CreatePongPacket(id string) NetPacket {
	pong := PongPacket{
		ID: id,
		TS: time.Now(),
	}
	b, err := json.Marshal(pong)
	if err != nil {
		log.Print(err)
	}

	return NetPacket{
		PacketType: d2netpackettype.Pong,
		PacketData: b,
	}
}

func UnmarshalPong(packet []byte) (PongPacket, error) {
	var resp PongPacket

	if err := json.Unmarshal(packet, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
