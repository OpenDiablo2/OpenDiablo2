package d2gui

import (
	"image/color"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type guiWidget interface {
	getLayer() int
	render(target d2render.Surface) error
	advance(elapsed float64) error
}

type guiManager struct {
	cursorSprite *Sprite
	loadSprite   *Sprite
	widgets      []guiWidget
	loading      bool
}

func createGuiManager() (*guiManager, error) {
	cursorSprite, err := CreateSprite(d2resource.CursorDefault, d2resource.PaletteUnits)
	if err != nil {
		return nil, err
	}

	loadSprite, err := CreateSprite(d2resource.LoadingScreen, d2resource.PaletteLoading)
	if err != nil {
		return nil, err
	}

	width, height := loadSprite.getSize()
	loadSprite.SetPosition(400-width/2, 300+height/2)

	manager := &guiManager{
		cursorSprite: cursorSprite,
		loadSprite:   loadSprite,
	}

	if err := d2input.BindHandler(manager); err != nil {
		return nil, err
	}

	return manager, nil
}

func (gui *guiManager) OnMouseMove(event d2input.MouseMoveEvent) bool {
	gui.cursorSprite.SetPosition(event.X, event.Y)
	return false
}

func (gui *guiManager) render(target d2render.Surface) error {
	if gui.loading {
		target.Clear(color.Black)
		if err := gui.loadSprite.render(target); err != nil {
			return err
		}
	} else {
		for _, widget := range gui.widgets {
			if err := widget.render(target); err != nil {
				return err
			}
		}
	}

	if err := gui.cursorSprite.render(target); err != nil {
		return err
	}

	return nil
}

func (gui *guiManager) advance(elapsed float64) error {
	if gui.loading {
		gui.loadSprite.Show()
		if err := gui.loadSprite.advance(elapsed); err != nil {
			return err
		}
	} else {
		gui.loadSprite.Hide()
		for _, widget := range gui.widgets {
			if err := widget.advance(elapsed); err != nil {
				return err
			}
		}
	}

	if err := gui.loadSprite.advance(elapsed); err != nil {
		return err
	}

	return nil
}

func (gui *guiManager) showLoadScreen(progress float64) {
	progress = math.Min(progress, 1.0)
	progress = math.Max(progress, 0.0)

	animation := gui.loadSprite.animation
	frameCount := animation.GetFrameCount()
	animation.SetCurrentFrame(int(float64(frameCount-1.0) * progress))

	gui.loading = true
}

func (gui *guiManager) hideLoadScreen() {
	gui.loading = false
}

func (gui *guiManager) showCursor() {
	gui.cursorSprite.Show()
}

func (gui *guiManager) hideCursor() {
	gui.cursorSprite.Hide()
}

func (gui *guiManager) clear() {
	gui.widgets = nil
	gui.hideLoadScreen()
}
