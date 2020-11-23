package d2localclient

import (
	"github.com/google/uuid"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2server"
)

// LocalClientConnection is the implementation of ClientConnection
// for a local client.
type LocalClientConnection struct {
	asset             *d2asset.AssetManager
	heroState         *d2hero.HeroStateFactory
	clientListener    d2networking.ClientListener // The game client
	uniqueID          string                      // Unique ID generated on construction
	openNetworkServer bool                        // True if this is a server
	playerState       *d2hero.HeroState           // Local player state
	gameServer        *d2server.GameServer        // Game Server

	logLevel d2util.LogLevel
}

// GetUniqueID returns LocalClientConnection.uniqueID.
func (l *LocalClientConnection) GetUniqueID() string {
	return l.uniqueID
}

// GetConnectionType returns an enum representing the connection type.
// See: d2clientconnectiontype
func (l *LocalClientConnection) GetConnectionType() d2clientconnectiontype.ClientConnectionType {
	return d2clientconnectiontype.Local
}

// SendPacketToClient passes a packet to the game client for processing.
func (l *LocalClientConnection) SendPacketToClient(packet d2netpacket.NetPacket) error {
	return l.clientListener.OnPacketReceived(packet)
}

// Create constructs a new LocalClientConnection and returns
// a pointer to it.
func Create(
	asset *d2asset.AssetManager,
	l d2util.LogLevel,
	openNetworkServer bool) (*LocalClientConnection, error) {
	heroStateFactory, err := d2hero.NewHeroStateFactory(asset)
	if err != nil {
		return nil, err
	}

	result := &LocalClientConnection{
		heroState:         heroStateFactory,
		asset:             asset,
		uniqueID:          uuid.New().String(),
		openNetworkServer: openNetworkServer,
		logLevel:          l,
	}

	return result, nil
}

// Open creates a new GameServer, runs the server and connects this client to it.
func (l *LocalClientConnection) Open(_, saveFilePath string) error {
	var err error

	l.SetPlayerState(l.heroState.LoadHeroState(saveFilePath))

	l.gameServer, err = d2server.NewGameServer(l.asset, l.openNetworkServer, l.logLevel, 30)
	if err != nil {
		return err
	}

	if err := l.gameServer.Start(); err != nil {
		return err
	}

	l.gameServer.OnClientConnected(l)

	return nil
}

// Close disconnects from the server and destroys it.
func (l *LocalClientConnection) Close() error {
	sc, err := d2netpacket.CreateServerClosedPacket()
	if err != nil {
		return err
	}

	err = l.SendPacketToServer(sc)
	if err != nil {
		return err
	}

	l.gameServer.OnClientDisconnected(l)
	l.gameServer.Stop()

	return nil
}

// SendPacketToServer calls d2server.OnPacketReceived with the given packet.
func (l *LocalClientConnection) SendPacketToServer(packet d2netpacket.NetPacket) error {
	return l.gameServer.OnPacketReceived(l, packet)
}

// SetClientListener sets LocalClientConnection.clientListener to the given value.
func (l *LocalClientConnection) SetClientListener(listener d2networking.ClientListener) {
	l.clientListener = listener
}

// GetPlayerState returns LocalClientConnection.playerState.
func (l *LocalClientConnection) GetPlayerState() *d2hero.HeroState {
	return l.playerState
}

// SetPlayerState sets LocalClientConnection.playerState to the given value.
func (l *LocalClientConnection) SetPlayerState(playerState *d2hero.HeroState) {
	l.playerState = playerState
}
