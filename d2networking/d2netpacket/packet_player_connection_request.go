package d2netpacket

import (
	"encoding/json"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
	"log"
)

// PlayerConnectionRequestPacket contains a player ID and game state.
// It is sent by a remote client to initiate a connection (join a game).
type PlayerConnectionRequestPacket struct {
	ID          string                `json:"id"`
	PlayerState *d2player.PlayerState `json:"gameState"`
}

// CreatePlayerConnectionRequestPacket returns a NetPacket which defines a
// PlayerConnectionRequestPacket with the given ID and game state.
func CreatePlayerConnectionRequestPacket(id string, playerState *d2player.PlayerState) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.PlayerConnectionRequest,
		PacketData: PlayerConnectionRequestPacket{
			ID:          id,
			PlayerState: playerState,
		},
	}
}

func UnmarshalPlayerConnectionRequest(packet []byte) (PlayerConnectionRequestPacket, error) {
	var p NetPacket
	if err := json.Unmarshal(packet, &p); err != nil {
		return PlayerConnectionRequestPacket{}, err
	}

	log.Println(p)
	return p.PacketData.(PlayerConnectionRequestPacket), nil
}
