package d2client

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapengine"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	d2cct "github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2localclient"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2remoteclient"
)

// Creates a connections to the server and returns a game client instance
func Create(connectionType d2cct.ClientConnectionType) (*GameClient, error) {
	result := &GameClient{
		// TODO: Mapgen - Needs levels.txt stuff
		MapEngine:      d2mapengine.CreateMapEngine(),
		Players:        make(map[string]*d2mapentity.Player),
		connectionType: connectionType,
		realm:          &d2mapengine.MapRealm{},
	}

	switch connectionType {
	case d2cct.LANClient:
		result.clientConnection = d2remoteclient.Create()
	case d2cct.LANServer:
		openSocket := true
		result.clientConnection = d2localclient.Create(openSocket)
	case d2cct.Local:
		dontOpenSocket := false
		result.clientConnection = d2localclient.Create(dontOpenSocket)
	default:
		str := "unknown client connection type specified: %d"
		return nil, fmt.Errorf(str, connectionType)
	}
	result.clientConnection.SetClientListener(result)
	return result, nil
}
