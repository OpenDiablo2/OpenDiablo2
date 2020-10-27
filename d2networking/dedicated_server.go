package d2networking

import (
	"os"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2server"
)

const (
	stopServer = 0b1
)

/*
StartDedicatedServer Checks whether or not we should start a server i.e the -listen parameter has been passed in, and if so launches a
server hosted to the network, in theory. (this is still WIP)
*/
func StartDedicatedServer(manager *d2asset.AssetManager, in, out chan byte, log chan string) (e error) {
	var listen bool

	var maxPlayers int = 3

	for argCount, arg := range os.Args {
		if arg == "-listen" {
			listen = true
		}

		if arg == "-maxplayers" {
			max, _ := strconv.ParseInt(os.Args[argCount+1], 10, 32)
			maxPlayers = int(max)
		}
	}

	if !listen {
		return nil
	}

	srvAsset, err := d2asset.NewAssetManager(d2config.Config)
	if err != nil {
		panic(err)
	}

	server, _ := d2server.NewGameServer(srvAsset, true, maxPlayers)
	serverErr := server.Start()

	if serverErr != nil {
		panic(err)
	}

	// I have done the messaging with a bitmask for memory efficiency, this can easilly be translated to pretty error
	// messages later, sue me.
	for {
		msgIn := <-in
		/* For those who do not know an AND operation denoted by & discards bits which do not line up so for instance:
		01011001 & 00000001 = 00000001 or 1
		00100101 & 00000010 = 00000000 or 0
		01100110 & 01100000 = 01100000 or 96
		these can be used to have multiple messages in just 8 bits, that's a quarter of a rune in go!
		*/
		if (msgIn & stopServer) == stopServer {
			break
		}
	}

	return nil
}
