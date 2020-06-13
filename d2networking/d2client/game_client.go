package d2client

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2localclient"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
)

type GameClient struct {
	clientConnection ClientConnection
}

func Create(connectionType d2clientconnectiontype.ClientConnectionType) (*GameClient, error) {
	result := &GameClient{
	}


	switch connectionType {
	case d2clientconnectiontype.Local:
		result.clientConnection = d2localclient.Create()
		result.clientConnection.SetClientListener(result)
	default:
		return nil, fmt.Errorf("unknown client connection type specified: %d", connectionType)
	}

	return result, nil
}

func (g *GameClient) Open(connectionString string) error {
	return g.clientConnection.Open(connectionString)
}

func (g *GameClient) Close() error {
	return g.clientConnection.Close()
}

func (g * GameClient) Destroy() error {
	return g.clientConnection.Close()
}

func (g * GameClient) OnPacketReceived(packet d2netpacket.NetPacket) error {
	return nil
}

func (g * GameClient) SendPacketToServer(packet d2netpacket.NetPacket) error {
	return g.clientConnection.SendPacketToServer(packet)
}
