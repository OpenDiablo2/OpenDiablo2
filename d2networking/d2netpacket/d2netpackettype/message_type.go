package d2netpackettype

type NetPacketType uint32

// Warning: Do NOT re-arrange the order of these packet values unless you want to
//          break compatibility between clients of slightly different versions.
//          Also note that the packet id is a byte, so if we use more than 256
//          of these then we are doing something very wrong.
const (
	UpdateServerInfo                NetPacketType = iota
	GenerateMap                                   // Sent by the server to generate a map
	AddPlayer                                     // Server sends to the client to add a player
	MovePlayer                                    // Sent to the client or server to indicate player movement
	PlayerConnectionRequest                       // Client sends to server to request a connection
	PlayerDisconnectionNotification               // Client notifies the server that it is disconnecting
	Ping                                          // Ping message type
	Pong                                          // Pong message type
	ServerClosed                                  // Local host has closed the server
)
