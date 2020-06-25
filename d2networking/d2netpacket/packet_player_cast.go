package d2netpacket

import "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"

// TODO: Need to handle being on different maps

type PlayerCastPacket struct {
	PlayerId string  `json:"playerId"`
	SkillId  int     `json:"skillId"`
	StartX   float64 `json:"startX"`
	StartY   float64 `json:"startY"`
	TargetX  float64 `json:"targetX"`
	TargetY  float64 `json:"targetY"`
}

func CreatePlayerCastPacket(playerId string, skillId int, startX, startY, targetX, targetY float64) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.PlayerCast,
		PacketData: PlayerCastPacket{
			PlayerId: playerId,
			SkillId:  skillId,
			StartX:   startX,
			StartY:   startY,
			TargetX:  targetX,
			TargetY:  targetY,
		},
	}
}
