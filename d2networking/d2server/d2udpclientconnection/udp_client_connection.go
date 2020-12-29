// Package d2udpclientconnection provides an implementation of a UDP client connection with a game state.
package d2udpclientconnection

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

const logPrefix = "UDP Connection"

// UDPClientConnection is the implementation of the
// d2server.ClientConnection interface to represent remote client from the
// server perspective.
type UDPClientConnection struct {
	id            string            // ID of the associated RemoteClientConnection
	address       *net.UDPAddr      // IP address of the associated RemoteClientConnection
	udpConnection *net.UDPConn      // Server's UDP Connection
	playerState   *d2hero.HeroState // Client's game state

	*d2util.Logger
}

// CreateUDPClientConnection constructs a new UDPClientConnection and
// returns a pointer to it.
func CreateUDPClientConnection(udpConnection *net.UDPConn, id string, l d2util.LogLevel, address *net.UDPAddr) *UDPClientConnection {
	result := &UDPClientConnection{
		id:            id,
		address:       address,
		udpConnection: udpConnection,
	}

	result.Logger = d2util.NewLogger()
	result.Logger.SetPrefix(logPrefix)
	result.Logger.SetLevel(l)

	return result
}

// GetUniqueID returns UDPClientConnection.id
func (u UDPClientConnection) GetUniqueID() string {
	return u.id
}

// GetConnectionType returns an enum representing the connection type.
// See: d2clientconnectiontype.
func (u UDPClientConnection) GetConnectionType() d2clientconnectiontype.ClientConnectionType {
	return d2clientconnectiontype.LANClient
}

// SendPacketToClient compresses the JSON encoding of a NetPacket and
// sends it to the client.
func (u *UDPClientConnection) SendPacketToClient(packet d2netpacket.NetPacket) error {
	data, err := json.Marshal(packet.PacketData)
	if err != nil {
		return err
	}

	var buff bytes.Buffer

	buff.WriteByte(byte(packet.PacketType))

	writer, err := gzip.NewWriterLevel(&buff, gzip.BestCompression)
	if err != nil {
		u.Error(err.Error())
	}

	if written, writeErr := writer.Write(data); writeErr != nil {
		return writeErr
	} else if written == 0 {
		return fmt.Errorf("RemoteClientConnection: attempted to send empty %v packet body",
			packet.PacketType)
	}

	if writeErr := writer.Close(); writeErr != nil {
		return writeErr
	}

	if _, udpErr := u.udpConnection.WriteToUDP(buff.Bytes(), u.address); udpErr != nil {
		return udpErr
	}

	return nil
}

// SetHeroState sets UDP.playerState to the given value.
func (u *UDPClientConnection) SetHeroState(playerState *d2hero.HeroState) {
	u.playerState = playerState
}

// GetHeroState returns UDPClientConnection.playerState.
func (u *UDPClientConnection) GetHeroState() *d2hero.HeroState {
	return u.playerState
}
