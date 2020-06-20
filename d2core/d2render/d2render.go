package d2render

import (
	"errors"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var (
	// ErrWasInit holding an error instance for initialized rendering system
	ErrWasInit = errors.New("rendering system is already initialized")
	// ErrNotInit holding an error instance for non-initialized rendering system
	ErrNotInit = errors.New("rendering system has not been initialized")
	// ErrInvalidRenderer holding an error instance for invalid rendering system specification
	ErrInvalidRenderer = errors.New("invalid rendering system specified")
)

var singleton d2interface.Renderer

// Initialize the renderer
func Initialize(rend d2interface.Renderer) error {
	verifyNotInit()
	singleton = rend
	log.Printf("Initialized the %s renderer...", singleton.GetRendererName())
	return nil
}

// SetWindowIcon sets the window icon by a given file name as string
func SetWindowIcon(fileName string) {
	verifyWasInit()
	singleton.SetWindowIcon(fileName)
}

// Run will run the renderer
func Run(f func(d2interface.Surface) error, width, height int, title string) error {
	verifyWasInit()
	singleton.Run(f, width, height, title)
	return nil
}

// IsDrawingSkipped checks whether the drawing is skipped
func IsDrawingSkipped() bool {
	verifyWasInit()
	return singleton.IsDrawingSkipped()
}

// CreateSurface creates a new surface, which returns the newly created surface or error
func CreateSurface(surface d2interface.Surface) (d2interface.Surface, error) {
	verifyWasInit()
	return singleton.CreateSurface(surface)
}

// NewSurface adds a new surface, and returns the new surface or error
func NewSurface(width, height int, filter d2interface.Filter) (d2interface.Surface, error) {
	verifyWasInit()
	return singleton.NewSurface(width, height, filter)
}

// IsFullScreen checks whether the window is on full screen
func IsFullScreen() bool {
	verifyWasInit()
	return singleton.IsFullScreen()
}

// SetFullScreen sets the window in fullscreen or windowed mode depending on the fullScreen flag
func SetFullScreen(fullScreen bool) {
	verifyWasInit()
	singleton.SetFullScreen(fullScreen)
}

// SetVSyncEnabled sets or unsets the VSync depending on the given vsync parameter flag
func SetVSyncEnabled(vsync bool) {
	verifyWasInit()
	singleton.SetVSyncEnabled(vsync)
}

// GetVSyncEnabled checks whether the VSync is enabled or not
func GetVSyncEnabled() bool {
	verifyWasInit()
	return singleton.GetVSyncEnabled()
}

// GetCursorPos returns the exact current position of the cursor
func GetCursorPos() (int, int) {
	verifyWasInit()
	return singleton.GetCursorPos()
}

// CurrentFPS returns the current frames per second
func CurrentFPS() float64 {
	verifyWasInit()
	return singleton.CurrentFPS()
}

func verifyWasInit() {
	if singleton == nil {
		panic(ErrNotInit)
	}
}

func verifyNotInit() {
	if singleton != nil {
		panic(ErrWasInit)
	}
}
