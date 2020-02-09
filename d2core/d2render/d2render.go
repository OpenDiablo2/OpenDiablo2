package d2render

import (
	"errors"
	"log"
)

var (
	ErrWasInit         = errors.New("rendering system is already initialized")
	ErrNotInit         = errors.New("rendering system has not been initialized")
	ErrInvalidRenderer = errors.New("invalid rendering system specified")
)

var singleton Renderer

func Initialize(rend Renderer) error {
	verifyNotInit()
	singleton = rend
	log.Printf("Initialized the %s renderer...", singleton.GetRendererName())
	return nil
}

func SetWindowIcon(fileName string) {
	verifyWasInit()
	singleton.SetWindowIcon(fileName)
}

func Run(f func(Surface) error, width, height int, title string) error {
	verifyWasInit()
	singleton.Run(f, width, height, title)
	return nil
}

func IsDrawingSkipped() bool {
	verifyWasInit()
	return singleton.IsDrawingSkipped()
}

func CreateSurface(surface Surface) (Surface, error) {
	verifyWasInit()
	return singleton.CreateSurface(surface)
}

func NewSurface(width, height int, filter Filter) (Surface, error) {
	verifyWasInit()
	return singleton.NewSurface(width, height, filter)
}

func IsFullScreen() bool {
	verifyWasInit()
	return singleton.IsFullScreen()
}

func SetFullScreen(fullScreen bool) {
	verifyWasInit()
	singleton.SetFullScreen(fullScreen)
}

func SetVSyncEnabled(vsync bool) {
	verifyWasInit()
	singleton.SetVSyncEnabled(vsync)
}

func GetVSyncEnabled() bool {
	verifyWasInit()
	return singleton.GetVSyncEnabled()
}

func GetCursorPos() (int, int) {
	verifyWasInit()
	return singleton.GetCursorPos()
}

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
