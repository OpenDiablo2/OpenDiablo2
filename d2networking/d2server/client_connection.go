package d2server

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

// ClientConnection is an interface for abstracting local and remote
// clients.
type ClientConnection interface {
	GetUniqueId() string
	GetConnectionType() d2clientconnectiontype.ClientConnectionType
	SendPacketToClient(packet d2netpacket.NetPacket) error
	GetPlayerState() *d2player.PlayerState
	SetPlayerState(playerState *d2player.PlayerState)
}
