package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/OpenDiablo2/OpenDiablo2/d2networking"

	"github.com/OpenDiablo2/OpenDiablo2/d2app"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	ebiten2 "github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio/ebiten"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render/ebiten"
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

	// Attempt to load the configuration file
	configError := d2config.Load()

	// Create our renderer
	renderer, err := ebiten.CreateRenderer()
	if err != nil {
		panic(err)
	}

	// If we failed to load our config, lets show the boot panic screen
	if configError != nil {
		renderer.ShowPanicScreen(configError.Error())
		return
	}

	// Create the asset manager
	asset, err := d2asset.NewAssetManager(d2config.Config)
	if err != nil {
		renderer.ShowPanicScreen(err.Error())
		return
	}

	srvChanIn := make(chan byte)
	srvChanLog := make(chan string)
	started, srvErr := d2networking.StartDedicatedServer(asset, srvChanIn, srvChanLog)

	if srvErr != nil {
		panic(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // This traps Control-c to safely shut down the server

	go func() {
		<-c
		srvChanIn <- 0b1
	}()

	if started {
		for data := range srvChanLog {
			log.Println(data)
		}
	}

	audio := ebiten2.CreateAudio(asset)

	inputManager := d2input.NewInputManager()

	term, err := d2term.New(inputManager)
	if err != nil {
		renderer.ShowPanicScreen(err.Error())
		return
	}

	err = asset.BindTerminalCommands(term)
	if err != nil {
		renderer.ShowPanicScreen(err.Error())
		return
	}

	scriptEngine := d2script.CreateScriptEngine()

	app := d2app.Create(GitBranch, GitCommit, inputManager, term, scriptEngine, audio, renderer, asset)
	if err := app.Run(); err != nil {
		renderer.ShowPanicScreen(err.Error())
		return
	}
}
