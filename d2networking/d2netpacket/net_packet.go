package d2netpacket

import (
	"encoding/json"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"log"
)

// NetPacket is used to wrap and send all packet types under d2netpacket.
// When decoding a packet: First the PacketType byte is read, then the
// PacketData is unmarshalled to a struct of the type associated with
// PacketType.
type NetPacket struct {
	PacketType d2netpackettype.NetPacketType `json:"packetType"`
	PacketData interface{}                   `json:"packetData"`
}

func InspectPacketType(b []byte) d2netpackettype.NetPacketType {
	var packet NetPacket

	if err := json.Unmarshal(b, &packet); err != nil {
		log.Println(err)
	}

	return packet.PacketType
}

func UnmarshalNetPacket(packet []byte) (NetPacket, error) {
	var p NetPacket
	if err := json.Unmarshal(packet, &p); err != nil {
		return p, err
	}

	return p, nil
}
