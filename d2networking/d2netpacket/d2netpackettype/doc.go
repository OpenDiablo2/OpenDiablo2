// Package d2netpackettype defines types which are encoded to JSON and sent in  network
//  packet payloads.
// Package d2netpacket/d2netpackettype defines a uint32 enumerable representing each
// packet type.
//
// A struct is defined for each packet type. Each struct comes with a function which
// returns a NetPacket declaring the type enum (header) followed by the associated
// struct (body). The NetPacket is marshaled to JSON for transport. On receipt of
// the packet, the enum is read as a single byte then the remaining data (the struct)
// is unmarshalled to the type associated with the type enum.
package d2netpackettype
