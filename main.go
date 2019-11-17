package main

import (
	"image"
	"log"
	"runtime"

	"github.com/pkg/profile"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2scene"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"

	"github.com/OpenDiablo2/OpenDiablo2/d2core"
	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2mpq"
	"github.com/hajimehoshi/ebiten"
)

// GitBranch is set by the CI build process to the name of the branch
var GitBranch string

// GitCommit is set by the CI build process to the commit hash
var GitCommit string
var d2Engine d2core.Engine

func main() {
	procs := runtime.GOMAXPROCS(16)
	log.Printf("Setting gomaxprocs to 16, it was previously set to %d", procs)
	//runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	defer profile.Start(profile.ProfilePath(".")).Stop()
	if len(GitBranch) == 0 {
		GitBranch = "Local Build"
		GitCommit = ""
	}
	d2common.SetBuildInfo(GitBranch, GitCommit)
	log.SetFlags(log.Ldate | log.LUTC | log.Lmicroseconds | log.Llongfile)
	log.Println("OpenDiablo2 - Open source Diablo 2 engine")
	_, iconImage, err := ebitenutil.NewImageFromFile("d2logo.png", ebiten.FilterLinear)
	if err == nil {
		ebiten.SetWindowIcon([]image.Image{iconImage})
	}
	d2mpq.InitializeCryptoBuffer()
	d2Engine = d2core.CreateEngine()
	d2Engine.SetNextScene(d2scene.CreateMainMenu(&d2Engine, &d2Engine, d2Engine.UIManager, d2Engine.SoundManager))
	ebiten.SetCursorVisible(false)
	ebiten.SetFullscreen(d2Engine.Settings.FullScreen)
	ebiten.SetRunnableInBackground(d2Engine.Settings.RunInBackground)
	ebiten.SetVsyncEnabled(d2Engine.Settings.VsyncEnabled)
	ebiten.SetMaxTPS(d2Engine.Settings.TicksPerSecond)
	if err := ebiten.Run(update, 800, 600, d2Engine.Settings.Scale, "OpenDiablo 2 ("+GitBranch+")"); err != nil {
		log.Fatal(err)
	}
}

func update(screen *ebiten.Image) error {
	d2Engine.Update()
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	d2Engine.Draw(screen)
	return nil
}
