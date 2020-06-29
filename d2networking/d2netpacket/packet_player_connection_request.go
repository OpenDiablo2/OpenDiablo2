package d2netpacket

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// PlayerConnectionRequestPacket contains a player ID and game state.
// It is sent by a remote client to initiate a connection (join a game).
type PlayerConnectionRequestPacket struct {
	Id          string                `json:"id"`
	PlayerState *d2player.PlayerState `json:"gameState"`
}

// CreatePlayerConnectionRequestPacket returns a NetPacket which defines a
// PlayerConnectionRequestPacket with the given ID and game state.
func CreatePlayerConnectionRequestPacket(id string, playerState *d2player.PlayerState) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.PlayerConnectionRequest,
		PacketData: PlayerConnectionRequestPacket{
			Id:          id,
			PlayerState: playerState,
		},
	}
}
