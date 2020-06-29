package d2client

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

// ClientConnection is an interface for abstracting local and
// remote server connections.
type ClientConnection interface {
	Open(connectionString string, saveFilePath string) error
	Close() error
	SendPacketToServer(packet d2netpacket.NetPacket) error
	SetClientListener(listener d2networking.ClientListener)
}
