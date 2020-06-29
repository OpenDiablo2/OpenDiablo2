package d2netpacket

import "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"

// MovePlayerPacket contains a movement command for a specific player entity.
// It is sent by the server to move a player entity on a client.
// TODO: Need to handle being on different maps
type MovePlayerPacket struct {
	PlayerId string  `json:"playerId"`
	StartX   float64 `json:"startX"`
	StartY   float64 `json:"startY"`
	DestX    float64 `json:"destX"`
	DestY    float64 `json:"destY"`
}

// CreateMovePlayerPacket returns a NetPacket which declares a MovePlayerPacket
// with the given ID and movement command.
func CreateMovePlayerPacket(playerId string, startX, startY, destX, destY float64) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.MovePlayer,
		PacketData: MovePlayerPacket{
			PlayerId: playerId,
			StartX:   startX,
			StartY:   startY,
			DestX:    destX,
			DestY:    destY,
		},
	}
}
