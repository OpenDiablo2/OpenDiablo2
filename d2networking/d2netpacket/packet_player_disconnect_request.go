package d2netpacket

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

// PlayerDisconnectRequestPacket contains a player ID and game state.
// It is sent by a remote client to close the connection (leave a game).
type PlayerDisconnectRequestPacket struct {
	Id          string                `json:"id"`
	PlayerState *d2player.PlayerState `json:"gameState"` // TODO: remove this? It isn't used.
}

// CreatePlayerDisconnectRequestPacket returns a NetPacket which defines a
// PlayerDisconnectRequestPacket with the given ID.
func CreatePlayerDisconnectRequestPacket(id string) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.PlayerDisconnectionNotification,
		PacketData: PlayerDisconnectRequestPacket{
			Id: id,
		},
	}
}
