package ebiten

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Renderer struct {
}

func CreateRenderer() (*Renderer, error) {
	result := &Renderer{}

	config, err := d2config.Get()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

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
	config, err := d2config.Get()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return ebiten.Run(func(img *ebiten.Image) error {
		err := f(&ebitenSurface{image: img})
		if err != nil {
			return err
		}
		return nil
	}, width, height, config.Scale, title)
}

func (r *Renderer) CreateSurface(surface d2render.Surface) (error, d2render.Surface) {
	result := &ebitenSurface{
		image: surface.(*ebitenSurface).image,
		stateCurrent: surfaceState{
			filter: ebiten.FilterNearest,
			mode:   ebiten.CompositeModeSourceOver,
		},
	}
	return nil, result
}

func (r *Renderer) NewSurface(width, height int, filter d2render.Filter) (error, d2render.Surface) {
	ebitenFilter := d2ToEbitenFilter(filter)
	img, err := ebiten.NewImage(width, height, ebitenFilter)
	if err != nil {
		return err, nil
	}
	result := &ebitenSurface{
		image: img,
	}
	return nil, result
}

func (r *Renderer) IsFullScreen() (bool, error) {
	return ebiten.IsFullscreen(), nil
}

func (r *Renderer) SetFullScreen(fullScreen bool) error {
	ebiten.SetFullscreen(fullScreen)
	return nil
}

func (r *Renderer) SetVSyncEnabled(vsync bool) error {
	ebiten.SetVsyncEnabled(vsync)
	return nil
}

func (r *Renderer) GetVSyncEnabled() (bool, error) {
	return ebiten.IsVsyncEnabled(), nil
}

func (r *Renderer) GetCursorPos() (int, int, error) {
	cx, cy := ebiten.CursorPosition()
	return cx, cy, nil
}

func (r *Renderer) CurrentFPS() float64 {
	return ebiten.CurrentFPS()
}
