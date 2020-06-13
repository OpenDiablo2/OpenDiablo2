package d2netpackettype

type NetPacketType uint32

const (
	UpdateServerInfo NetPacketType = iota
	GenerateMap
	AddPlayer
)
