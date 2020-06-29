package d2interface

type Renderer interface {
	GetRendererName() string
	SetWindowIcon(fileName string)
	Run(f func(Surface) error, width, height int, title string) error
	IsDrawingSkipped() bool
	CreateSurface(surface Surface) (Surface, error)
	NewSurface(width, height int, filter Filter) (Surface, error)
	IsFullScreen() bool
	SetFullScreen(fullScreen bool)
	SetVSyncEnabled(vsync bool)
	GetVSyncEnabled() bool
	GetCursorPos() (int, int)
	CurrentFPS() float64
}
