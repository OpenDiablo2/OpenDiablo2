package d2netpacket

import (
	"encoding/json"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"log"
)

// ServerFullPacket contains the current time. It is sent by the server
// to inform a client that the server has reached the max number of allowed connections.
type ServerFullPacket struct{}

// CreateServerFullPacket returns a NetPacket which declares a
// ServerFullPacket with the current time.
func CreateServerFullPacket() NetPacket {
	serverClosed := ServerFullPacket{}

	b, err := json.Marshal(serverClosed)
	if err != nil {
		log.Print(err)
	}

	return NetPacket{
		PacketType: d2netpackettype.ServerFull,
		PacketData: b,
	}
}

// UnmarshalServerFull unmarshalls the given data to a ServerFullPacket struct
func UnmarshalServerFull(packet []byte) (ServerFullPacket, error) {
	var resp ServerFullPacket

	if err := json.Unmarshal(packet, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
