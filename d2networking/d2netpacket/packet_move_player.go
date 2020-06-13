package d2netpacket

import "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"

// TODO: Need to handle being on different maps

type MovePlayerPacket struct {
	PlayerId string  `json:"playerId"`
	StartX   float64 `json:"startX"`
	StartY   float64 `json:"startY"`
	DestX    float64 `json:"destX"`
	DestY    float64 `json:"destY"`
}

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
