package main

import (
	"image"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/OpenDiablo2/OpenDiablo2/common"

	"github.com/OpenDiablo2/OpenDiablo2/core"
	"github.com/OpenDiablo2/OpenDiablo2/mpq"
	"github.com/hajimehoshi/ebiten"
)

// GitBranch is set by the CI build process to the name of the branch
var GitBranch string

// GitCommit is set by the CI build process to the commit hash
var GitCommit string
var d2Engine *core.Engine

func main() {
	//runtime.LockOSThread()
	//defer runtime.UnlockOSThread()
	if len(GitBranch) == 0 {
		GitBranch = "Local Build"
		GitCommit = ""
	}
	common.SetBuildInfo(GitBranch, GitCommit)
	log.SetFlags(log.Ldate | log.LUTC | log.Lmicroseconds | log.Llongfile)
	log.Println("OpenDiablo2 - Open source Diablo 2 engine")
	_, iconImage, err := ebitenutil.NewImageFromFile("d2logo.png", ebiten.FilterLinear)
	if err == nil {
		ebiten.SetWindowIcon([]image.Image{iconImage})
	}
	mpq.InitializeCryptoBuffer()
	d2Engine = core.CreateEngine()
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
	ebitenutil.DebugPrintAt(screen, "vsync:"+strconv.FormatBool(ebiten.IsVsyncEnabled())+"\nFPS:"+strconv.Itoa(int(ebiten.CurrentFPS())), 5, 565)
	return nil
}
