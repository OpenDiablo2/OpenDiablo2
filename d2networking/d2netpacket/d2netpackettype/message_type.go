package d2netpackettype

type NetPacketType uint32

// Warning: Do NOT re-arrange the order of these packet values unless you want to
//          break compatibility between clients of slightly different versions.
const (
	UpdateServerInfo NetPacketType = iota
	GenerateMap
	AddPlayer
	MovePlayer
)
