package d2networking

import (
	"os"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2server"
)

// ServerEventFlag represents a server event
type ServerEventFlag = int

// Server events
const (
	ServerEventStop ServerEventFlag = iota
)

// ServerMinPlayers is the minimum number of players a server can have
const (
	ServerMinPlayers        = 1
	ServerMaxPlayersDefault = 8
)

func hasFlag(value, flag int) bool {
	return (value & flag) == flag
}

/*
StartDedicatedServer Checks whether or not we should start a server i.e the -listen parameter has been passed in, and if so launches a
server hosted to the network, in theory. (this is still WIP)
*/
func StartDedicatedServer(
	manager *d2asset.AssetManager,
	in chan int,
	log chan string,
	l d2util.LogLevel,
	maxPlayers int,
) error {
	server, err := d2server.NewGameServer(manager, true, l, maxPlayers)
	if err != nil {
		return err
	}

	err = server.Start()
	if err != nil {
		return err
	}

	for {
		msgIn := <-in
		if hasFlag(msgIn, ServerEventStop) {
			log <- "Stopping server"

			server.Stop()
			log <- "Exiting..."

			os.Exit(0)
		}
	}
}

// ServerOptions represents game server options
type ServerOptions struct {
	Dedicated  *bool
	MaxPlayers *int
}
