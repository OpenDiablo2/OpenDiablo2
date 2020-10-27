package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2app"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render/ebiten"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2server"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	ebiten2 "github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio/ebiten"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"
	"github.com/OpenDiablo2/OpenDiablo2/d2script"
)

// GitBranch is set by the CI build process to the name of the branch
//nolint:gochecknoglobals // This is filled in by the build system
var GitBranch string

// GitCommit is set by the CI build process to the commit hash
//nolint:gochecknoglobals // This is filled in by the build system
var GitCommit string

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("OpenDiablo2 - Open source Diablo 2 engine")

	if err := d2config.Load(); err != nil {
		panic(err)
	}

	// NewAssetManager our providers
	renderer, err := ebiten.CreateRenderer()
	if err != nil {
		panic(err)
	}

	asset, err := d2asset.NewAssetManager(d2config.Config)
	if err != nil {
		panic(err)
	}

	audio, err := ebiten2.CreateAudio(asset)
	if err != nil {
		panic(err)
	}

	inputManager := d2input.NewInputManager()

	term, err := d2term.New(inputManager)
	if err != nil {
		log.Fatal(err)
	}

	err = asset.BindTerminalCommands(term)
	if err != nil {
		log.Fatal(err)
	}

	scriptEngine := d2script.CreateScriptEngine()

	var listen bool

	var maxplayers int = 3

	for argCount, arg := range os.Args {
		if arg == "-listen" {
			listen = true
		}

		if arg == "-maxplayers" {
			max, _ := strconv.ParseInt(os.Args[argCount+1], 10, 32)
			maxplayers = int(max)
		}
	}

	if listen {
		server, _ := d2server.NewGameServer(asset, true, maxplayers)
		serverErr := server.Start()

		if serverErr != nil {
			log.Fatal(serverErr)
			os.Exit(1)
		}

		for {
			fmt.Println("Ah yes, another second of being a server, lovely")
			// This is jank to get the goroutine of server.Start to continue
			time.Sleep(time.Second)
		}
	}

	app := d2app.Create(GitBranch, GitCommit, inputManager, term, scriptEngine, audio, renderer, asset)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
