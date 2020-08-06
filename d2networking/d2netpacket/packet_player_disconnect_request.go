package d2netpacket

import (
	"encoding/json"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// PlayerDisconnectRequestPacket contains a player ID and game state.
// It is sent by a remote client to close the connection (leave a game).
type PlayerDisconnectRequestPacket struct {
	ID          string                `json:"id"`
	PlayerState *d2player.PlayerState `json:"gameState"` // TODO: remove this? It isn't used.
}

// CreatePlayerDisconnectRequestPacket returns a NetPacket which defines a
// PlayerDisconnectRequestPacket with the given ID.
func CreatePlayerDisconnectRequestPacket(id string) NetPacket {
	playerDisconnectRequest := PlayerDisconnectRequestPacket{
		ID: id,
	}
	b, _ := json.Marshal(playerDisconnectRequest)

	return NetPacket{
		PacketType: d2netpackettype.PlayerDisconnectionNotification,
		PacketData: b,
	}
}

func UnmarshalPlayerDisconnectionRequest(packet []byte) (PlayerDisconnectRequestPacket, error) {
	var resp PlayerDisconnectRequestPacket

	if err := json.Unmarshal(packet, &resp); err != nil {
		return resp, err
	}
	return resp, nil
}
