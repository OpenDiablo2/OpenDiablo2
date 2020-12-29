package ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

// PanicScreen represents the system to draw a panic message on the screen.
type PanicScreen struct {
	errorMessage string
}

// CreatePanicScreen creates a new panic screen.
func CreatePanicScreen(errorMessage string) *PanicScreen {
	result := &PanicScreen{
		errorMessage: errorMessage,
	}

	ebiten.SetWindowTitle("OpenDiablo 2 - PANIC SCREEN")
	ebiten.SetWindowResizable(true)

	return result
}

// Update updates a game by one tick. The given argument represents a screen image.
//
// Update updates only the game logic and Draw draws the screen.
//
// In the first frame, it is ensured that Update is called at least once before Draw. You can use Update
// to initialize the game state.
//
// After the first frame, Update might not be called or might be called once
// or more for one frame. The frequency is determined by the current TPS (tick-per-second).
func (s *PanicScreen) Update() error {
	return nil
}

// Draw draws the game screen by one frame.
//
// The give argument represents a screen image. The updated content is adopted as the game screen.
func (s *PanicScreen) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(colornames.Darkred)

	ebitenutil.DebugPrint(screen, s.errorMessage)
}

// Layout accepts a native outside size in device-independent pixels and returns the game's logical screen
// size.
//
// On desktops, the outside is a window or a monitor (fullscreen mode). On browsers, the outside is a body
// element. On mobiles, the outside is the view's size.
//
// Even though the outside size and the screen size differ, the rendering scale is automatically adjusted to
// fit with the outside.
//
// Layout is called almost every frame.
//
// If Layout returns non-positive numbers, the caller can panic.
//
// You can return a fixed screen size if you don't care, or you can also return a calculated screen size
// adjusted with the given outside size.
func (s *PanicScreen) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
