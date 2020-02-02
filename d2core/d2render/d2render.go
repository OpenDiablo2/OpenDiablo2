package d2render

import (
	"errors"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
)

var (
	ErrHasInit         = errors.New("rendering system is already initialized")
	ErrNotInit         = errors.New("rendering system has not been initialized")
	ErrInvalidRenderer = errors.New("invalid rendering system specified")
)

var singleton Renderer

func Initialize(rend Renderer) error {
	if singleton != nil {
		return d2input.ErrHasInit
	}
	singleton = rend
	log.Printf("Initialized the %s renderer...", singleton.GetRendererName())
	return nil
}

func SetWindowIcon(fileName string) error {
	if singleton == nil {
		return ErrNotInit
	}
	singleton.SetWindowIcon(fileName)
	return nil
}

func Run(f func(Surface) error, width, height int, title string) error {
	if singleton == nil {
		return ErrNotInit
	}
	singleton.Run(f, width, height, title)
	return nil
}

func IsDrawingSkipped() (error, bool) {
	if singleton == nil {
		return ErrNotInit, true
	}
	return nil, singleton.IsDrawingSkipped()
}

func CreateSurface(surface Surface) (error, Surface) {
	if singleton == nil {
		return ErrNotInit, nil
	}
	return singleton.CreateSurface(surface)
}

func NewSurface(width, height int, filter Filter) (error, Surface) {
	if singleton == nil {
		return ErrNotInit, nil
	}
	return singleton.NewSurface(width, height, filter)
}

func IsFullScreen() (bool, error) {
	if singleton == nil {
		return false, ErrNotInit
	}
	return singleton.IsFullScreen()
}

func SetFullScreen(fullScreen bool) error {
	if singleton == nil {
		return ErrNotInit
	}
	return singleton.SetFullScreen(fullScreen)
}

func SetVSyncEnabled(vsync bool) error {
	if singleton == nil {
		return ErrNotInit
	}
	return singleton.SetVSyncEnabled(vsync)
}

func GetVSyncEnabled() (bool, error) {
	if singleton == nil {
		return false, ErrNotInit
	}
	return singleton.GetVSyncEnabled()
}

func GetCursorPos() (int, int, error) {
	if singleton == nil {
		return 0, 0, ErrNotInit
	}
	return singleton.GetCursorPos()
}

func CurrentFPS() (float64, error) {
	if singleton == nil {
		return 0, ErrNotInit
	}
	return singleton.CurrentFPS(), nil
}
