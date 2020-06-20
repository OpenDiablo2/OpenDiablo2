// Package d2netpackettype defines the enumerable NetPacketType.
package d2netpackettype

// NetPacketType is an enum referring to all packet types in package
// d2netpacket.
type NetPacketType uint32

// (Except NetPacket which declares a NetPacketType to specify the packet body
// type. See d2netpackettype.NetPacket.)
//
// Warning
//
// Do NOT re-arrange the order of these packet values unless you want to
// break compatibility between clients of slightly different versions.
// Also note that the packet id is a byte, so if we use more than 256 of
// these then we are doing something very wrong.
const (
	UpdateServerInfo                NetPacketType = iota // Sent by the server, client sets the given player ID and map seed
	GenerateMap                                          // Sent by the server, client generates a map
	AddPlayer                                            // Sent by the server, client adds a player
	MovePlayer                                           // Sent by client or server, moves a player entity
	PlayerConnectionRequest                              // Sent by the remote client when connecting
	PlayerDisconnectionNotification                      // Sent by the remote client when disconnecting
	Ping                                                 // Requests a Pong packet
	Pong                                                 // Responds to a Ping packet
	ServerClosed                                         // Sent by the local host when it has closed the server
	CastSkill                                            // Sent by client or server, indicates entity casting skill
)

func (n NetPacketType) String() string {
	strings := map[NetPacketType]string{
		UpdateServerInfo:                "UpdateServerInfo",
		GenerateMap:                     "GenerateMap",
		AddPlayer:                       "AddPlayer",
		MovePlayer:                      "MovePlayer",
		PlayerConnectionRequest:         "PlayerConnectionRequest",
		PlayerDisconnectionNotification: "PlayerDisconnectionNotification",
		Ping:                            "Ping",
		Pong:                            "Pong",
		ServerClosed:                    "ServerClosed",
		CastSkill:                       "CastSkill",
	}

	return strings[n]
}
