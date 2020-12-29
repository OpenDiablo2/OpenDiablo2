package d2netpacket

import (
	"encoding/json"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// PlayerConnectionRequestPacket contains a player ID and game state.
// It is sent by a remote client to initiate a connection (join a game).
type PlayerConnectionRequestPacket struct {
	ID          string            `json:"id"`
	PlayerState *d2hero.HeroState `json:"gameState"`
}

// CreatePlayerConnectionRequestPacket returns a NetPacket which defines a
// PlayerConnectionRequestPacket with the given ID and game state.
func CreatePlayerConnectionRequestPacket(id string, playerState *d2hero.HeroState) (NetPacket, error) {
	playerConnectionRequest := PlayerConnectionRequestPacket{
		ID:          id,
		PlayerState: playerState,
	}

	b, err := json.Marshal(playerConnectionRequest)
	if err != nil {
		return NetPacket{PacketType: d2netpackettype.PlayerConnectionRequest}, err
	}

	return NetPacket{
		PacketType: d2netpackettype.PlayerConnectionRequest,
		PacketData: b,
	}, nil
}

// UnmarshalPlayerConnectionRequest unmarshals the given data to a
// PlayerConnectionRequestPacket struct
func UnmarshalPlayerConnectionRequest(packet []byte) (PlayerConnectionRequestPacket, error) {
	var resp PlayerConnectionRequestPacket

	if err := json.Unmarshal(packet, &resp); err != nil {
		return PlayerConnectionRequestPacket{}, err
	}

	return resp, nil
}
