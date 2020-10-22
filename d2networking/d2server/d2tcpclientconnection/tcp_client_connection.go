// Package d2tcpclientconnection provides a TCP protocol implementation of a client connection
package d2tcpclientconnection

import (
	"encoding/json"
	"net"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

// TCPClientConnection represents a client connection over TCP
type TCPClientConnection struct {
	id            string
	tcpConnection net.Conn
	playerState   *d2hero.HeroState
}

// CreateTCPClientConnection creates a new tcp client connection instance
func CreateTCPClientConnection(tcpConnection net.Conn, id string) *TCPClientConnection {
	return &TCPClientConnection{
		tcpConnection: tcpConnection,
		id:            id,
	}
}

// GetUniqueID returns the unique ID for the tcp client connection
func (t TCPClientConnection) GetUniqueID() string {
	return t.id
}

// SendPacketToClient marshals and sends (writes) NetPackets
func (t *TCPClientConnection) SendPacketToClient(p d2netpacket.NetPacket) error {
	packet, err := json.Marshal(p)
	if err != nil {
		return err
	}

	_, err = t.tcpConnection.Write(packet)
	if err != nil {
		return err
	}

	return nil
}

// SetPlayerState sets the game client player state
func (t *TCPClientConnection) SetPlayerState(playerState *d2hero.HeroState) {
	t.playerState = playerState
}

// GetPlayerState gets the game client player state
func (t *TCPClientConnection) GetPlayerState() *d2hero.HeroState {
	return t.playerState
}

// GetConnectionType returns an enum representing the connection type.
// See: d2clientconnectiontype.
func (t TCPClientConnection) GetConnectionType() d2clientconnectiontype.ClientConnectionType {
	return d2clientconnectiontype.LANClient
}
