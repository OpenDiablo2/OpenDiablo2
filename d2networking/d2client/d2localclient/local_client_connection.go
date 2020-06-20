// Package d2localclient facilitates communication between a local client and server.
package d2localclient

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2server"
	uuid "github.com/satori/go.uuid"
)

// LocalClientConnection is the implementation of ClientConnection
// for a local client.
type LocalClientConnection struct {
	clientListener    d2networking.ClientListener // The game client
	uniqueId          string                      // Unique ID generated on construction
	openNetworkServer bool                        // True if this is a server
	playerState       *d2player.PlayerState       // Local player state
}

// GetUniqueId returns LocalClientConnection.uniqueId.
func (l LocalClientConnection) GetUniqueId() string {
	return l.uniqueId
}

// GetConnectionType returns an enum representing the connection type.
// See: d2clientconnectiontype
func (l LocalClientConnection) GetConnectionType() d2clientconnectiontype.ClientConnectionType {
	return d2clientconnectiontype.Local
}

// SendPacketToClient passes a packet to the game client for processing.
func (l *LocalClientConnection) SendPacketToClient(packet d2netpacket.NetPacket) error {
	return l.clientListener.OnPacketReceived(packet)
}

// Create constructs a new LocalClientConnection and returns
// a pointer to it.
func Create(openNetworkServer bool) *LocalClientConnection {
	result := &LocalClientConnection{
		uniqueId:          uuid.NewV4().String(),
		openNetworkServer: openNetworkServer,
	}

	return result
}

// Open creates a new GameServer, runs the server and connects this client to it.
func (l *LocalClientConnection) Open(_ string, saveFilePath string) error {
	l.SetPlayerState(d2player.LoadPlayerState(saveFilePath))
	d2server.Create(l.openNetworkServer)

	go d2server.Run()
	d2server.OnClientConnected(l)

	return nil
}

// Close disconnects from the server and destroys it.
func (l *LocalClientConnection) Close() error {
	err := l.SendPacketToServer(d2netpacket.CreateServerClosedPacket())
	if err != nil {
		return err
	}

	d2server.OnClientDisconnected(l)
	d2server.Destroy()

	return nil
}

// SendPacketToServer calls d2server.OnPacketReceived with the given packet.
func (l *LocalClientConnection) SendPacketToServer(packet d2netpacket.NetPacket) error {
	// TODO: This is going to blow up if the server has ceased to be.
	return d2server.OnPacketReceived(l, packet)
}

// SetClientListener sets LocalClientConnection.clientListener to the given value.
func (l *LocalClientConnection) SetClientListener(listener d2networking.ClientListener) {
	l.clientListener = listener
}

// GetPlayerState returns LocalClientConnection.playerState.
func (l *LocalClientConnection) GetPlayerState() *d2player.PlayerState {
	return l.playerState
}

// SetPlayerState sets LocalClientConnection.playerState to the given value.
func (l *LocalClientConnection) SetPlayerState(playerState *d2player.PlayerState) {
	l.playerState = playerState
}
