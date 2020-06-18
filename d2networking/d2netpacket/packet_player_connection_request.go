package d2netpacket

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gamestate"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

type PlayerConnectionRequestPacket struct {
	Id        string                 `json:"id"`
	GameState *d2gamestate.GameState `json:"gameState"`
}

func CreatePlayerConnectionRequestPacket(id string, gameState *d2gamestate.GameState) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.PlayerConnectionRequest,
		PacketData: PlayerConnectionRequestPacket{
			Id:        id,
			GameState: gameState,
		},
	}
}
