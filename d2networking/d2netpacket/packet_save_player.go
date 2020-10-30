package d2netpacket

import (
	"encoding/json"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"log"
)

// SavePlayerPacket has the actual selected left and right skill
// the Server has to check if these skills are actually allowed for the Player
type SavePlayerPacket struct {
	LeftSkill  *d2hero.HeroSkill `json:"leftSkill"`
	RightSkill *d2hero.HeroSkill `json:"rightSkill"`
}

// CreateSavePlayerPacket sends a packet which instructs the server to save the Player
func CreateSavePlayerPacket(playerState *d2mapentity.Player) NetPacket {
	ping := SavePlayerPacket{
		LeftSkill:  playerState.LeftSkill,
		RightSkill: playerState.RightSkill,
	}

	b, err := json.Marshal(ping)
	if err != nil {
		log.Print(err)
	}

	return NetPacket{
		PacketType: d2netpackettype.SavePlayer,
		PacketData: b,
	}
}

// UnmarshalSavePlayer unmarshals the given data to a SavePlayerPacket struct
func UnmarshalSavePlayer(packet []byte) (SavePlayerPacket, error) {
	var p SavePlayerPacket
	if err := json.Unmarshal(packet, &p); err != nil {
		return p, err
	}

	return p, nil
}
