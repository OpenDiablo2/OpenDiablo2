package d2netpacket

import (
	"encoding/json"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// SavePlayerPacket has the actual selected left and right skill
// the Server has to check if these skills are actually allowed for the Player
type SavePlayerPacket struct {
	Player     *d2mapentity.Player   `json:"Player"`
	Difficulty d2enum.DifficultyType `json:"Difficulty"`
}

// CreateSavePlayerPacket sends a packet which instructs the server to save the Player
func CreateSavePlayerPacket(playerState *d2mapentity.Player, difficulty d2enum.DifficultyType) (NetPacket, error) {
	savePlayerData := SavePlayerPacket{
		Player:     playerState,
		Difficulty: difficulty,
	}

	b, err := json.Marshal(savePlayerData)
	if err != nil {
		return NetPacket{PacketType: d2netpackettype.SavePlayer}, err
	}

	return NetPacket{
		PacketType: d2netpackettype.SavePlayer,
		PacketData: b,
	}, nil
}

// UnmarshalSavePlayer unmarshalls the given data to a SavePlayerPacket struct
func UnmarshalSavePlayer(packet []byte) (SavePlayerPacket, error) {
	var p SavePlayerPacket
	if err := json.Unmarshal(packet, &p); err != nil {
		return p, err
	}

	return p, nil
}
