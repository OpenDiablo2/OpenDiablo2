package d2scene

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type Scene interface{}

type SceneLoadHandler interface {
	OnLoad() error
}

type SceneUnloadHandler interface {
	OnUnload() error
}

type SceneRenderHandler interface {
	Render(target d2render.Surface) error
}

type SceneAdvanceHandler interface {
	Advance(elapsed float64) error
}

var singleton struct {
	nextScene    Scene
	loadingScene Scene
	currentScene Scene
}

func SetNextScene(scene Scene) {
	singleton.nextScene = scene
}

func Advance(elapsed float64) error {
	if singleton.nextScene != nil {
		if handler, ok := singleton.currentScene.(SceneUnloadHandler); ok {
			if err := handler.OnUnload(); err != nil {
				return err
			}
		}

		d2ui.Reset()
		d2gui.Clear()

		if _, ok := singleton.nextScene.(SceneLoadHandler); ok {
			d2gui.ShowLoadScreen(0)
			d2gui.HideCursor()
			singleton.currentScene = nil
			singleton.loadingScene = singleton.nextScene
		} else {
			singleton.currentScene = singleton.nextScene
			singleton.loadingScene = nil
		}

		singleton.nextScene = nil
	} else if singleton.loadingScene != nil {
		handler := singleton.loadingScene.(SceneLoadHandler)
		if err := handler.OnLoad(); err != nil {
			return err
		}

		singleton.currentScene = singleton.loadingScene
		singleton.loadingScene = nil
		d2gui.ShowCursor()
		d2gui.HideLoadScreen()
	} else if singleton.currentScene != nil {
		if handler, ok := singleton.currentScene.(SceneAdvanceHandler); ok {
			if err := handler.Advance(elapsed); err != nil {
				return err
			}
		}
	}

	return nil
}

func Render(surface d2render.Surface) error {
	if handler, ok := singleton.currentScene.(SceneRenderHandler); ok {
		if err := handler.Render(surface); err != nil {
			return err
		}
	}

	return nil
}
