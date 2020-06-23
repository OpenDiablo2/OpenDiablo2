package d2netpacket

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket/d2netpackettype"
)

type PlayerDisconnectRequestPacket struct {
	Id          string                `json:"id"`
	PlayerState *d2player.PlayerState `json:"gameState"`
}

func CreatePlayerDisconnectRequestPacket(id string) NetPacket {
	return NetPacket{
		PacketType: d2netpackettype.PlayerDisconnectionNotification,
		PacketData: PlayerDisconnectRequestPacket{
			Id: id,
		},
	}
}
