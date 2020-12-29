package d2netpacket

import (
	"encoding/json"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// NetPacket is used to wrap and send all packet types under d2netpacket.
// When decoding a packet: First the PacketType byte is read, then the
// PacketData is unmarshalled to a struct of the type associated with
// PacketType.
type NetPacket struct {
	PacketType d2netpackettype.NetPacketType `json:"packetType"`
	PacketData json.RawMessage               `json:"packetData"`
}

// InspectPacketType determines the packet type from the given data
func InspectPacketType(b []byte) (d2netpackettype.NetPacketType, error) {
	var packet NetPacket

	if err := json.Unmarshal(b, &packet); err != nil {
		return d2netpackettype.UnknownPacketType, err
	}

	return packet.PacketType, nil
}

// UnmarshalNetPacket unmarshals the byte slice into a NetPacket struct
func UnmarshalNetPacket(packet []byte) (NetPacket, error) {
	var p NetPacket
	if err := json.Unmarshal(packet, &p); err != nil {
		return p, err
	}

	return p, nil
}

// MarshalPacket is a quick helper function to Marshal very anything UNSAFELY, meaning the error is not checked before sending.
func MarshalPacket(packet interface{}) ([]byte, error) {
	b, err := json.Marshal(packet)
	if err != nil {
		return nil, err
	}

	return b, nil
}
