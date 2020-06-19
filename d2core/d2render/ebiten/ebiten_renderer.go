package ebiten

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Renderer struct {
	renderCallback func(surface d2render.Surface) error
}

func (r *Renderer) Update(screen *ebiten.Image) error {
	err := r.renderCallback(&ebitenSurface{image: screen})
	if err != nil {
		return err
	}
	return nil
}

func (r *Renderer) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}

func CreateRenderer() (*Renderer, error) {
	result := &Renderer{}

	config := d2config.Get()

	ebiten.SetCursorVisible(false)
	ebiten.SetFullscreen(config.FullScreen)
	ebiten.SetRunnableInBackground(config.RunInBackground)
	ebiten.SetVsyncEnabled(config.VsyncEnabled)
	ebiten.SetMaxTPS(config.TicksPerSecond)

	return result, nil
}

func (*Renderer) GetRendererName() string {
	return "Ebiten"
}

func (*Renderer) SetWindowIcon(fileName string) {
	_, iconImage, err := ebitenutil.NewImageFromFile(fileName, ebiten.FilterLinear)
	if err == nil {
		ebiten.SetWindowIcon([]image.Image{iconImage})
	}
}

func (r *Renderer) IsDrawingSkipped() bool {
	return ebiten.IsDrawingSkipped()
}

func (r *Renderer) Run(f func(surface d2render.Surface) error, width, height int, title string) error {
	r.renderCallback = f
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(width, height)
	return ebiten.RunGame(r)
}

func (r *Renderer) CreateSurface(surface d2render.Surface) (d2render.Surface, error) {
	result := &ebitenSurface{
		image: surface.(*ebitenSurface).image,
		stateCurrent: surfaceState{
			filter: ebiten.FilterNearest,
			mode:   ebiten.CompositeModeSourceOver,
		},
	}
	return result, nil
}

func (r *Renderer) NewSurface(width, height int, filter d2render.Filter) (d2render.Surface, error) {
	ebitenFilter := d2ToEbitenFilter(filter)
	img, err := ebiten.NewImage(width, height, ebitenFilter)
	if err != nil {
		return nil, err
	}
	result := &ebitenSurface{image: img}
	return result, nil
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
