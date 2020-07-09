package ebiten

import (
	"errors"
	"image"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// CreateRenderer creates an ebiten renderer instance
func CreateRenderer() (*Renderer, error) {
	result := &Renderer{}

	config := d2config.Config
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ebiten.SetFullscreen(config.FullScreen)
	ebiten.SetRunnableOnUnfocused(config.RunInBackground)
	ebiten.SetVsyncEnabled(config.VsyncEnabled)
	ebiten.SetMaxTPS(config.TicksPerSecond)

	return result, nil
}

// Renderer renders the game
type Renderer struct {
	app d2interface.App
	renderCallback func(surface d2interface.Surface) error
}

// BindApp binds the renderer to the App instance
func (r *Renderer) BindApp(app d2interface.App) error {
	if r.app != nil {
		return errors.New("renderer already bound to app instance")
	}
	r.app = app
	return nil
}

// Initialize is called by the App after all AppComponents are bound
func (r *Renderer) Initialize() error {
	// TODO any renderer init need to be done ?
	return nil
}

// Update is the render function
func (r *Renderer) Update(screen *ebiten.Image) error {
	err := r.renderCallback(createEbitenSurface(screen))
	if err != nil {
		return err
	}

	return nil
}

// Layout defines the screen dimensions for the renderer
func (r *Renderer) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}


// GetRendererName returns the name of the renderer
func (*Renderer) GetRendererName() string {
	return "Ebiten"
}

// SetWindowIcon sets the window icon
func (*Renderer) SetWindowIcon(fileName string) {
	_, iconImage, err := ebitenutil.NewImageFromFile(fileName, ebiten.FilterLinear)
	if err == nil {
		ebiten.SetWindowIcon([]image.Image{iconImage})
	}
}

//
func (r *Renderer) IsDrawingSkipped() bool {
	return ebiten.IsDrawingSkipped()
}

func (r *Renderer) Run(f func(surface d2interface.Surface) error, width, height int, title string) error {
	r.renderCallback = f
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(width, height)
	return ebiten.RunGame(r)
}

func (r *Renderer) CreateSurface(surface d2interface.Surface) (d2interface.Surface, error) {
	result := createEbitenSurface(
		surface.(*ebitenSurface).image,
		surfaceState{
			filter: ebiten.FilterNearest,
			effect: d2enum.DrawEffectNone,
		},
	)

	return result, nil
}

func (r *Renderer) NewSurface(width, height int, filter d2enum.Filter) (d2interface.Surface, error) {
	ebitenFilter := d2ToEbitenFilter(filter)
	img, err := ebiten.NewImage(width, height, ebitenFilter)
	if err != nil {
		return nil, err
	}
	return createEbitenSurface(img), nil
}

func (r *Renderer) IsFullScreen() bool {
	return ebiten.IsFullscreen()
}

func (r *Renderer) SetFullScreen(fullScreen bool) {
	ebiten.SetFullscreen(fullScreen)
}

func (r *Renderer) SetVSyncEnabled(vsync bool) {
	ebiten.SetVsyncEnabled(vsync)
}

func (r *Renderer) GetVSyncEnabled() bool {
	return ebiten.IsVsyncEnabled()
}

func (r *Renderer) GetCursorPos() (int, int) {
	return ebiten.CursorPosition()
}

func (r *Renderer) CurrentFPS() float64 {
	return ebiten.CurrentFPS()
}
