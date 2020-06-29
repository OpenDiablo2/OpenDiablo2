// Package d2netpacket defines types which are encoded to JSON and sent in  network
// packet payloads.
/*
Package d2netpacket/d2netpackettype defines a uint32 enumerable representing each
packet type.

A struct is defined for each packet type. Each struct comes with a function which
returns a NetPacket declaring the type enum (header) followed by the associated
struct (body). The NetPacket is marshalled to JSON for transport. On receipt of
the packet, the enum is read as a single byte then the remaining data (the struct)
is unmarshalled to the type associated with the type enum.*/
package d2netpacket

import "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"

// NetPacket is used to wrap and send all packet types under d2netpacket.
// When decoding a packet: First the PacketType byte is read, then the
// PacketData is unmarshalled to a struct of the type associated with
// PacketType.
type NetPacket struct {
	PacketType d2netpackettype.NetPacketType `json:"packetType"`
	PacketData interface{}                   `json:"packetData"`
}
