package d2netpacket

import (
	"encoding/json"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// CastPacket contains a cast command for an entity. It is sent by the server
// and instructs the client to trigger the use of the given skill on the given
// entity.
type CastPacket struct {
	SourceEntityID string  `json:"sourceEntityId"`
	SkillID        int     `json:"skillId"`
	TargetX        float64 `json:"targetX"`
	TargetY        float64 `json:"targetY"`
	TargetEntityID string  `json:"targetEntityId"`
}

// CreateCastPacket returns a NetPacket which declares a CastPacket with the
// given skill command.
func CreateCastPacket(entityID string, skillID int, targetX, targetY float64) (NetPacket, error) {
	castPacket := CastPacket{
		SourceEntityID: entityID,
		SkillID:        skillID,
		TargetX:        targetX,
		TargetY:        targetY,
		TargetEntityID: "", // https://github.com/OpenDiablo2/OpenDiablo2/issues/826
	}

	b, err := json.Marshal(castPacket)
	if err != nil {
		return NetPacket{PacketType: d2netpackettype.CastSkill}, err
	}

	return NetPacket{
		PacketType: d2netpackettype.CastSkill,
		PacketData: b,
	}, nil
}

// UnmarshalCast unmarshals the given data to a CastPacket struct
func UnmarshalCast(packet []byte) (CastPacket, error) {
	var p CastPacket
	if err := json.Unmarshal(packet, &p); err != nil {
		return p, err
	}

	return p, nil
}
