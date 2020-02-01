package ebiten

import (
	"image"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2config"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type EbitenRenderer struct {
}

func CreateRenderer() (*EbitenRenderer, error) {
	result := &EbitenRenderer{}

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

func (*EbitenRenderer) GetRendererName() string {
	return "Ebiten"
}

func (*EbitenRenderer) SetWindowIcon(fileName string) {
	_, iconImage, err := ebitenutil.NewImageFromFile(fileName, ebiten.FilterLinear)
	if err == nil {
		ebiten.SetWindowIcon([]image.Image{iconImage})
	}

}

func (r *EbitenRenderer) IsDrawingSkipped() bool {
	return ebiten.IsDrawingSkipped()
}

func (r *EbitenRenderer) Run(f func(surface d2common.Surface) error, width, height int, title string) error {
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

func (r *EbitenRenderer) CreateSurface(surface d2common.Surface) (error, d2common.Surface) {
	result := &ebitenSurface{
		image: surface.(*ebitenSurface).image,
		stateCurrent: surfaceState{
			filter: ebiten.FilterNearest,
			mode:   ebiten.CompositeModeSourceOver,
		},
	}
	return nil, result
}

func (r *EbitenRenderer) NewSurface(width, height int, filter d2common.Filter) (error, d2common.Surface) {
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

func (r *EbitenRenderer) IsFullScreen() (bool, error) {
	return ebiten.IsFullscreen(), nil
}

func (r *EbitenRenderer) SetFullScreen(fullScreen bool) error {
	ebiten.SetFullscreen(fullScreen)
	return nil
}

func (r *EbitenRenderer) SetVSyncEnabled(vsync bool) error {
	ebiten.SetVsyncEnabled(vsync)
	return nil
}

func (r *EbitenRenderer) GetVSyncEnabled() (bool, error) {
	return ebiten.IsVsyncEnabled(), nil
}

func (r *EbitenRenderer) GetCursorPos() (int, int, error) {
	cx, cy := ebiten.CursorPosition()
	return cx, cy, nil
}

func (r *EbitenRenderer) CurrentFPS() float64 {
	return ebiten.CurrentFPS()
}
