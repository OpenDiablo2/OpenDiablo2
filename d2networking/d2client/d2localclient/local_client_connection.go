package d2localclient

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	uuid "github.com/satori/go.uuid"

	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2server"
)

// LocalClientConnection is the implementation of ClientConnection
// for a local client.
type LocalClientConnection struct {
	asset             *d2asset.AssetManager
	clientListener    d2networking.ClientListener // The game client
	uniqueID          string                      // Unique ID generated on construction
	openNetworkServer bool                        // True if this is a server
	playerState       *d2player.PlayerState       // Local player state
	gameServer        *d2server.GameServer        // Game Server
}

// GetUniqueID returns LocalClientConnection.uniqueID.
func (l LocalClientConnection) GetUniqueID() string {
	return l.uniqueID
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
func Create(asset *d2asset.AssetManager, openNetworkServer bool) *LocalClientConnection {
	result := &LocalClientConnection{
		asset:             asset,
		uniqueID:          uuid.NewV4().String(),
		openNetworkServer: openNetworkServer,
	}

	return result
}

// Open creates a new GameServer, runs the server and connects this client to it.
func (l *LocalClientConnection) Open(_, saveFilePath string) error {
	var err error

	l.SetPlayerState(d2player.LoadPlayerState(saveFilePath))

	l.gameServer, err = d2server.NewGameServer(l.asset, l.openNetworkServer, 30)
	if err != nil {
		return err
	}

	if err := l.gameServer.Start(); err != nil {
		return err
	}

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
	l.gameServer.Stop()

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
