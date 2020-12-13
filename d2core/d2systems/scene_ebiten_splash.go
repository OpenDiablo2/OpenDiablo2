package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2input"
	"github.com/gravestench/akara"
	"image/color"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

const (
	sceneKeyEbitenSplash = "Ebiten Splash"
)

const (
	splashDelaySeconds = 0.5
	splashTimeout = 3
)

// static check that EbitenSplashScene implements the scene interface
var _ d2interface.Scene = &EbitenSplashScene{}

// NewEbitenSplashScene creates a new main menu scene. This is the first screen that the user
// will see when launching the game.
func NewEbitenSplashScene() *EbitenSplashScene {
	scene := &EbitenSplashScene{
		BaseScene: NewBaseScene(sceneKeyEbitenSplash),
		delay: splashDelaySeconds,
	}

	scene.backgroundColor = color.Black

	return scene
}

// EbitenSplashScene represents the in-game terminal for typing commands
type EbitenSplashScene struct {
	*BaseScene
	booted  bool
	squares []akara.EID
	timeElapsed float64
	delay float64
}

// Init the terminal
func (s *EbitenSplashScene) Init(world *akara.World) {
	s.World = world

	s.Debug("initializing ...")
}

func (s *EbitenSplashScene) boot() {
	if !s.BaseScene.booted {
		s.BaseScene.boot()
		return
	}

	s.createSplash()

	s.booted = true
}

// Update and render the terminal to the terminal viewport
func (s *EbitenSplashScene) Update() {
	for _, id := range s.Viewports {
		s.Components.Priority.Add(id).Priority = scenePriorityEbitenSplash
	}

	if s.Paused() {
		return
	}

	if !s.booted {
		s.boot()
	}

	if !s.booted {
		return
	}

	s.updateSplash()

	s.BaseScene.Update()
}

func (s *EbitenSplashScene) createSplash() {
	//                                                         @
	//  @ @ @   @       @     @      @ @ @    @                @ @
	//  @       @           @ @ @    @   @    @ @ @        @ @
	//  @ @     @ @ @   @     @      @ @ @    @   @      @ @ @
	//  @       @   @   @     @      @        @   @    @ @ @
	//  @ @ @   @ @ @   @     @ @    @ @ @    @   @    @ @
	flags := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0, 1, 1, 0, 0},
		{1, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0},
		{1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 0},
		{1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 0, 0, 0, 0},
	}

	orange := color.RGBA{0xdb, 0x56, 0x20, 255}
	squares := make([]akara.EID, 0)
	size := 10

	totalW, totalH := len(flags[0])*size, len(flags)*size
	ox, oy := (800-totalW)/2, (600-totalH)/2

	for y, row := range flags {
		for x, col := range row {
			if col == 0 {
				continue
			}

			square := s.Add.Rectangle(ox+x*size, oy+y*size, size, size, orange)

			s.Components.Alpha.Add(square).Alpha = 0

			squares = append(squares, square)
		}
	}

	interactive := s.Components.Interactive.Add(s.NewEntity())

	interactive.InputVector.SetMouseButton(d2input.MouseButtonLeft)

	interactive.Callback = func() bool {
		s.Debug("hiding splash scene")

		s.timeElapsed = splashTimeout

		interactive.Enabled = false

		return true // prevent propagation
	}

	s.squares = squares
}

func (s *EbitenSplashScene) updateSplash() {
	if s.delay > 0 {
		s.delay -= s.TimeDelta.Seconds()
		return
	}

	// fade out after timeout
	if s.timeElapsed >= splashTimeout {
		vpAlpha, _ := s.Components.Alpha.Get(s.Viewports[0])
		vpAlpha.Alpha -= 0.0425
		if vpAlpha.Alpha <= 0 {
			vpAlpha.Alpha = 0

			s.Debug("finished, deactivating")
			s.SetActive(false)
		}
	}

	numSquares := float64(len(s.squares))
	s.timeElapsed += s.TimeDelta.Seconds()

	// fade all of the squares
	for idx, id := range s.squares {
		a := math.Sin(s.timeElapsed*2 + -90 + (float64(idx)/numSquares))
		a = (a+1)/2 // clamp between 0..1

		alpha, found := s.Components.Alpha.Get(id)
		if !found {
			s.Components.Alpha.Add(id)
		}

		alpha.Alpha = a
	}
}
