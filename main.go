package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/OpenDiablo2/OpenDiablo2/Common"

	"github.com/OpenDiablo2/OpenDiablo2/Core"
	"github.com/OpenDiablo2/OpenDiablo2/MPQ"
	"github.com/hajimehoshi/ebiten"
)

var GitBranch string
var GitCommit string
var d2Engine *Core.Engine

func main() {
	if len(GitBranch) == 0 {
		GitBranch = "Local Build"
		GitCommit = ""
	}
	Common.SetBuildInfo(GitBranch, GitCommit)
	log.SetFlags(log.Ldate | log.LUTC | log.Lmicroseconds | log.Llongfile)
	log.Println("OpenDiablo2 - Open source Diablo 2 engine")
	_, iconImage, err := ebitenutil.NewImageFromFile("d2logo.png", ebiten.FilterLinear)
	if err == nil {
		ebiten.SetWindowIcon([]image.Image{iconImage})
	}
	MPQ.InitializeCryptoBuffer()
	d2Engine = Core.CreateEngine()
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
	//ebitenutil.DebugPrint(screen, "FPS:"+strconv.Itoa(int(ebiten.CurrentFPS())))
	return nil
}
