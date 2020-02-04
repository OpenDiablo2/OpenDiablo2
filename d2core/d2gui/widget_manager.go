package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type widget interface {
	Render(target d2render.Surface) error
	Advance(elapsed float64) error
}

type widgetManager struct {
}
