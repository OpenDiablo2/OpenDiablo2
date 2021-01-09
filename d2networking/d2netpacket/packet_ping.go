package d2netpacket //nolint:dupl // ServerClosed and Ping just happen to be very similar packets

import (
	"encoding/json"
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
func CreatePingPacket() (NetPacket, error) {
	ping := PingPacket{
		TS: time.Now(),
	}

	b, err := json.Marshal(ping)
	if err != nil {
		return NetPacket{PacketType: d2netpackettype.Ping}, err
	}

	return NetPacket{
		PacketType: d2netpackettype.Ping,
		PacketData: b,
	}, nil
}

// UnmarshalPing unmarshals the given data to a PingPacket struct
func UnmarshalPing(packet []byte) (PingPacket, error) {
	var p PingPacket
	if err := json.Unmarshal(packet, &p); err != nil {
		return p, err
	}

	return p, nil
}
