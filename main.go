package main

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2app"

	ebiten_input "github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/ebiten"

	ebiten2 "github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio/ebiten"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"
)

// GitBranch is set by the CI build process to the name of the branch
//nolint:gochecknoglobals This is filled in by the build system
var GitBranch string

// GitCommit is set by the CI build process to the commit hash
//nolint:gochecknoglobals This is filled in by the build system
var GitCommit string

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("OpenDiablo2 - Open source Diablo 2 engine")

	// Initialize our providers
	audio, err := ebiten2.CreateAudio()
	if err != nil {
		panic(err)
	}

	d2input.Initialize(ebiten_input.InputService{}) // TODO d2input singleton must be init before d2term
	term, err := d2term.Initialize()

	if err != nil {
		log.Fatal(err)
	}

	app := d2app.Create(GitBranch, GitCommit, term, audio)
	app.Run()
}
