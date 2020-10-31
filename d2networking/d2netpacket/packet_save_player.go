package d2netpacket

import (
	"encoding/json"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// SavePlayerPacket has the actual selected left and right skill
// the Server has to check if these skills are actually allowed for the Player
type SavePlayerPacket struct {
	LeftSkill  int `json:"leftSkill"`
	RightSkill int `json:"rightSkill"`
}

// CreateSavePlayerPacket sends a packet which instructs the server to save the Player
func CreateSavePlayerPacket(playerState *d2mapentity.Player) NetPacket {
	savePlayerData := SavePlayerPacket{
		LeftSkill:  playerState.LeftSkill.ID,
		RightSkill: playerState.RightSkill.ID,
	}

	b, err := json.Marshal(savePlayerData)
	if err != nil {
		log.Print(err)
	}

	return NetPacket{
		PacketType: d2netpackettype.SavePlayer,
		PacketData: b,
	}
}

// UnmarshalSavePlayer unmarshalls the given data to a SavePlayerPacket struct
func UnmarshalSavePlayer(packet []byte) (SavePlayerPacket, error) {
	var p SavePlayerPacket
	if err := json.Unmarshal(packet, &p); err != nil {
		return p, err
	}

	return p, nil
}
