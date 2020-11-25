package d2netpacket

import (
	"encoding/json"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// PlayerDisconnectRequestPacket contains a player ID and game state.
// It is sent by a remote client to close the connection (leave a game).
type PlayerDisconnectRequestPacket struct {
	ID          string            `json:"id"`
	PlayerState *d2hero.HeroState `json:"gameState"`
}

// CreatePlayerDisconnectRequestPacket returns a NetPacket which defines a
// PlayerDisconnectRequestPacket with the given ID.
func CreatePlayerDisconnectRequestPacket(id string) (NetPacket, error) {
	playerDisconnectRequest := PlayerDisconnectRequestPacket{
		ID: id,
	}

	b, err := json.Marshal(playerDisconnectRequest)
	if err != nil {
		return NetPacket{PacketType: d2netpackettype.PlayerDisconnectionNotification}, err
	}

	return NetPacket{
		PacketType: d2netpackettype.PlayerDisconnectionNotification,
		PacketData: b,
	}, nil
}

// UnmarshalPlayerDisconnectionRequest unmarshals the given data to a
// PlayerDisconnectRequestPacket struct
func UnmarshalPlayerDisconnectionRequest(packet []byte) (PlayerDisconnectRequestPacket, error) {
	var resp PlayerDisconnectRequestPacket

	if err := json.Unmarshal(packet, &resp); err != nil {
		return resp, err
	}

	return resp, nil
}
