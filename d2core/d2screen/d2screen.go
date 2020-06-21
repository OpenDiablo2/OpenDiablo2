package d2screen

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type Screen interface{}

type ScreenLoadHandler interface {
	OnLoad() error
}

type ScreenUnloadHandler interface {
	OnUnload() error
}

type ScreenRenderHandler interface {
	Render(target d2render.Surface) error
}

type ScreenAdvanceHandler interface {
	Advance(elapsed float64) error
}

var singleton struct {
	nextScreen    Screen
	loadingScreen Screen
	currentScreen Screen
}

func SetNextScreen(screen Screen) {
	singleton.nextScreen = screen
}

func Advance(elapsed float64) error {
	if singleton.nextScreen != nil {
		if handler, ok := singleton.currentScreen.(ScreenUnloadHandler); ok {
			if err := handler.OnUnload(); err != nil {
				return err
			}
		}

		d2ui.Reset()
		d2gui.SetLayout(nil)

		if _, ok := singleton.nextScreen.(ScreenLoadHandler); ok {
			d2gui.ShowLoadScreen(0)
			d2gui.HideCursor()
			singleton.currentScreen = nil
			singleton.loadingScreen = singleton.nextScreen
		} else {
			singleton.currentScreen = singleton.nextScreen
			singleton.loadingScreen = nil
		}

		singleton.nextScreen = nil
	} else if singleton.loadingScreen != nil {
		handler := singleton.loadingScreen.(ScreenLoadHandler)
		if err := handler.OnLoad(); err != nil {
			return err
		}

		singleton.currentScreen = singleton.loadingScreen
		singleton.loadingScreen = nil
		d2gui.ShowCursor()
		d2gui.HideLoadScreen()
	} else if singleton.currentScreen != nil {
		if handler, ok := singleton.currentScreen.(ScreenAdvanceHandler); ok {
			if err := handler.Advance(elapsed); err != nil {
				return err
			}
		}
	}

	return nil
}

func Render(surface d2render.Surface) error {
	if handler, ok := singleton.currentScreen.(ScreenRenderHandler); ok {
		if err := handler.Render(surface); err != nil {
			return err
		}
	}

	return nil
}
