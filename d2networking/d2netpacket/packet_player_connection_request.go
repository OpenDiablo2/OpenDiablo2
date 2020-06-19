package d2netpacket

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

type PlayerConnectionRequestPacket struct {
	Id          string                `json:"id"`
	PlayerState *d2player.PlayerState `json:"gameState"`
}

func CreatePlayerConnectionRequestPacket(id string, playerState *d2player.PlayerState) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.PlayerConnectionRequest,
		PacketData: PlayerConnectionRequestPacket{
			Id:          id,
			PlayerState: playerState,
		},
	}
}
