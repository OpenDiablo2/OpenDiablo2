package main

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2systems"
	"github.com/gravestench/akara"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/OpenDiablo2/OpenDiablo2/d2app"
)

// GitBranch is set by the CI build process to the name of the branch
//nolint:gochecknoglobals // This is filled in by the build system
var GitBranch string = "local"

// GitCommit is set by the CI build process to the commit hash
//nolint:gochecknoglobals // This is filled in by the build system
var GitCommit string = "build"

func main() {
	ecs := kingpin.Flag("ecs", "start the ecs implementation").Bool()
	kingpin.Parse()

	if *ecs {
		cfg := akara.NewWorldConfig()

		cfg.
			With(d2systems.NewAppBootstrapSystem()).
			With(d2systems.NewGameClientBootstrapSystem())

		akara.NewWorld(cfg)

		return
	}

	log.SetFlags(log.Lshortfile)

	instance := d2app.Create(GitBranch, GitCommit)

	if err := instance.Run(); err != nil {
		return
	}
}
