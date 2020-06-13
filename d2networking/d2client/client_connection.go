package d2client

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

type ClientConnection interface {
	Open(connectionString string) error
	Close() error
	SendPacketToServer(packet d2netpacket.NetPacket) error
	SetClientListener(listener d2networking.ClientListener)
}
