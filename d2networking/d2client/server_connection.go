package d2client

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

// ServerConnection is an interface for abstracting local and
// remote server connections.
type ServerConnection interface {
	Open(connectionString string, saveFilePath string) error
	Close() error
	SendPacketToServer(packet d2netpacket.NetPacket) error
	SetClientListener(listener d2networking.ClientListener)
}
