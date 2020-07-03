package main

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render/ebiten"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"

	"github.com/OpenDiablo2/OpenDiablo2/d2app"
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

	////////////////////////////////////////////////////////////////////////////
	// TODO these things need to be converted to AppComponents
	d2config.Load()
	d2asset.Initialize()

	renderer, err := ebiten.CreateRenderer()
	if err != nil {
		panic(err)
	}

	if err := d2render.Initialize(renderer); err != nil {
		panic(err)
	}
	////////////////////////////////////////////////////////////////////////////

	app := d2app.Create(GitBranch, GitCommit)
	app.Run()
}
