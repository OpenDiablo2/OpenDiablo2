package d2render

type Renderer interface {
	GetRendererName() string
	SetWindowIcon(fileName string)
	Run(f func(Surface) error, width, height int, title string) error
	IsDrawingSkipped() bool
	CreateSurface(surface Surface) (error, Surface)
	NewSurface(width, height int, filter Filter) (error, Surface)
	IsFullScreen() (bool, error)
	SetFullScreen(fullScreen bool) error
	SetVSyncEnabled(vsync bool) error
	GetVSyncEnabled() (bool, error)
	GetCursorPos() (int, int, error)
	CurrentFPS() float64
}
