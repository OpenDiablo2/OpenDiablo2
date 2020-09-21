package d2server

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

// ClientConnection is an interface for abstracting local and remote
// clients.
type ClientConnection interface {
	GetUniqueID() string
	GetConnectionType() d2clientconnectiontype.ClientConnectionType
	SendPacketToClient(packet d2netpacket.NetPacket) error
	GetPlayerState() *d2hero.HeroState
	SetPlayerState(playerState *d2hero.HeroState)
}
