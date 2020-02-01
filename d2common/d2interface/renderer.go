package d2interface

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type Renderer interface {
	GetRendererName() string
	SetWindowIcon(fileName string)
	Run(f func(d2common.Surface) error, width, height int, title string) error
	IsDrawingSkipped() bool
	CreateSurface(surface d2common.Surface) (error, d2common.Surface)
	NewSurface(width, height int, filter d2common.Filter) (error, d2common.Surface)
	IsFullScreen() (bool, error)
	SetFullScreen(fullScreen bool) error
	SetVSyncEnabled(vsync bool) error
	GetVSyncEnabled() (bool, error)
	GetCursorPos() (int, int, error)
	CurrentFPS() float64
}
