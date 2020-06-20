package d2netpacket

import "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"

// CastPacket contains a cast command for an entity. It is sent by the server
// and instructs the client to trigger the use of the given skill on the given
// entity.
// TODO: Need to handle being on different maps
type CastPacket struct {
	SourceEntityID string  `json:"sourceEntityId"`
	SkillID        int     `json:"skillId"`
	TargetX        float64 `json:"targetX"`
	TargetY        float64 `json:"targetY"`
	TargetEntityID string  `json:"targetEntityId"`
}

// CreateCastPacket returns a NetPacket which declares a CastPacket with the
// given skill command.
func CreateCastPacket(entityID string, skillID int, targetX, targetY float64) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.CastSkill,
		PacketData: CastPacket{
			SourceEntityID: entityID,
			SkillID:        skillID,
			TargetX:        targetX,
			TargetY:        targetY,
			TargetEntityID: "", // TODO implement targeting entities
		},
	}
}
