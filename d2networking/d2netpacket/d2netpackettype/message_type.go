package d2netpackettype

type NetPacketType uint32

// Warning: Do NOT re-arrange the order of these packet values unless you want to
//          break compatibility between clients of slightly different versions.
const (
	UpdateServerInfo        NetPacketType = iota
	GenerateMap                           // Sent by the server to generate a map
	AddPlayer                             // Server sends to the client to add a player
	MovePlayer                            // Sent to the client or server to indicate player movement
	PlayerConnectionRequest               // Client sends to server to request a connection
)
