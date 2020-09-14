package ebiten

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
)

const (
	screenWidth       = 800
	screenHeight      = 600
	defaultSaturation = 1.0
	defaultBrightness = 1.0
	defaultSkewX      = 0.0
	defaultSkewY      = 0.0
	defaultScaleX     = 1.0
	defaultScaleY     = 1.0
)

// Renderer is an implementation of a renderer
type Renderer struct {
	renderCallback func(surface d2interface.Surface) error
}

// Update updates the screen with the given *ebiten.Image
func (r *Renderer) Update(screen *ebiten.Image) error {
	err := r.renderCallback(createEbitenSurface(r, screen))
	if err != nil {
		return err
	}

	return nil
}

// Layout returns the renderer screen width and height
func (r *Renderer) Layout(_, _ int) (width, height int) {
	return screenWidth, screenHeight
}

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

// GetRendererName returns the name of the renderer
func (*Renderer) GetRendererName() string {
	return "Ebiten"
}

// SetWindowIcon sets the icon for the window, visible in the chrome of the window
func (*Renderer) SetWindowIcon(fileName string) {
	_, iconImage, err := ebitenutil.NewImageFromFile(fileName, ebiten.FilterLinear)
	if err == nil {
		ebiten.SetWindowIcon([]image.Image{iconImage})
	}
}

// IsDrawingSkipped returns a bool for whether or not the drawing has been skipped
func (r *Renderer) IsDrawingSkipped() bool {
	return ebiten.IsDrawingSkipped()
}

// Run initializes the renderer
func (r *Renderer) Run(f func(surface d2interface.Surface) error, width, height int, title string) error {
	r.renderCallback = f

	ebiten.SetWindowTitle(title)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(width, height)

	return ebiten.RunGame(r)
}

// CreateSurface creates a renderer surface from an existing surface
func (r *Renderer) CreateSurface(surface d2interface.Surface) (d2interface.Surface, error) {
	img := surface.(*ebitenSurface).image
	sfcState := surfaceState{
		filter:     ebiten.FilterNearest,
		effect:     d2enum.DrawEffectNone,
		saturation: defaultSaturation,
		brightness: defaultBrightness,
		skewX:      defaultSkewX,
		skewY:      defaultSkewY,
		scaleX:     defaultScaleX,
		scaleY:     defaultScaleY,
	}
	result := createEbitenSurface(r, img, sfcState)

	return result, nil
}

// NewSurface creates a new surface
func (r *Renderer) NewSurface(width, height int, filter d2enum.Filter) (d2interface.Surface, error) {
	ebitenFilter := d2ToEbitenFilter(filter)
	img, err := ebiten.NewImage(width, height, ebitenFilter)

	if err != nil {
		return nil, err
	}

	return createEbitenSurface(r, img), nil
}

// IsFullScreen returns a boolean for whether or not the renderer is currently set to fullscreen
func (r *Renderer) IsFullScreen() bool {
	return ebiten.IsFullscreen()
}

// SetFullScreen sets the renderer to fullscreen, given a boolean
func (r *Renderer) SetFullScreen(fullScreen bool) {
	ebiten.SetFullscreen(fullScreen)
}

// SetVSyncEnabled enables vsync, given a boolean
func (r *Renderer) SetVSyncEnabled(vsync bool) {
	ebiten.SetVsyncEnabled(vsync)
}

// GetVSyncEnabled returns a boolean for whether or not vsync is enabled
func (r *Renderer) GetVSyncEnabled() bool {
	return ebiten.IsVsyncEnabled()
}

// GetCursorPos returns the current cursor position x,y coordinates
func (r *Renderer) GetCursorPos() (x, y int) {
	return ebiten.CursorPosition()
}

// CurrentFPS returns the current frames per second of the renderer
func (r *Renderer) CurrentFPS() float64 {
	return ebiten.CurrentFPS()
}
