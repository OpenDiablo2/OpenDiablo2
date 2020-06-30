package ebiten

import (
	"fmt"
	"image"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type ebitenSurface struct {
	stateStack   []surfaceState
	stateCurrent surfaceState
	image        *ebiten.Image
}

func (s *ebitenSurface) PushTranslation(x, y int) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.x += x
	s.stateCurrent.y += y
}

func (s *ebitenSurface) PushCompositeMode(mode d2enum.CompositeMode) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.mode = d2ToEbitenCompositeMode(mode)
}

func (s *ebitenSurface) PushFilter(filter d2interface.Filter) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.filter = d2ToEbitenFilter(filter)
}

func (s *ebitenSurface) PushColor(color color.Color) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.color = color
}

func (s *ebitenSurface) PushBrightness(brightness float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.brightness = brightness
}

func (s *ebitenSurface) Pop() {
	count := len(s.stateStack)
	if count == 0 {
		panic("empty stack")
	}

	s.stateCurrent = s.stateStack[count-1]
	s.stateStack = s.stateStack[:count-1]
}

func (s *ebitenSurface) PopN(n int) {
	for i := 0; i < n; i++ {
		s.Pop()
	}
}

func (s *ebitenSurface) Render(sfc d2interface.Surface) error {
	opts := &ebiten.DrawImageOptions{CompositeMode: s.stateCurrent.mode}
	opts.GeoM.Translate(float64(s.stateCurrent.x), float64(s.stateCurrent.y))
	opts.Filter = s.stateCurrent.filter
	if s.stateCurrent.color != nil {
		opts.ColorM = ColorToColorM(s.stateCurrent.color)
	}
	if s.stateCurrent.brightness != 0 {
		opts.ColorM.ChangeHSV(0, 1, s.stateCurrent.brightness)
	}

	var img = sfc.(*ebitenSurface).image
	return s.image.DrawImage(img, opts)
}

// Renders the section of the animation frame enclosed by bounds
func (s *ebitenSurface) RenderSection(sfc d2interface.Surface, bound image.Rectangle) error {
	opts := &ebiten.DrawImageOptions{CompositeMode: s.stateCurrent.mode}
	opts.GeoM.Translate(float64(s.stateCurrent.x), float64(s.stateCurrent.y))
	opts.Filter = s.stateCurrent.filter
	if s.stateCurrent.color != nil {
		opts.ColorM = ColorToColorM(s.stateCurrent.color)
	}
	if s.stateCurrent.brightness != 0 {
		opts.ColorM.ChangeHSV(0, 1, s.stateCurrent.brightness)
	}

	var img = sfc.(*ebitenSurface).image
	return s.image.DrawImage(img.SubImage(bound).(*ebiten.Image), opts)
}

func (s *ebitenSurface) DrawText(format string, params ...interface{}) {
	ebitenutil.DebugPrintAt(s.image, fmt.Sprintf(format, params...), s.stateCurrent.x, s.stateCurrent.y)
}

func (s *ebitenSurface) DrawLine(x, y int, color color.Color) {
	ebitenutil.DrawLine(
		s.image,
		float64(s.stateCurrent.x),
		float64(s.stateCurrent.y),
		float64(s.stateCurrent.x+x),
		float64(s.stateCurrent.y+y),
		color,
	)
}

func (s *ebitenSurface) DrawRect(width, height int, color color.Color) {
	ebitenutil.DrawRect(
		s.image,
		float64(s.stateCurrent.x),
		float64(s.stateCurrent.y),
		float64(width),
		float64(height),
		color,
	)
}

func (s *ebitenSurface) Clear(color color.Color) error {
	return s.image.Fill(color)
}

func (s *ebitenSurface) GetSize() (int, int) {
	return s.image.Size()
}

func (s *ebitenSurface) GetDepth() int {
	return len(s.stateStack)
}

func (s *ebitenSurface) ReplacePixels(pixels []byte) error {
	return s.image.ReplacePixels(pixels)
}

func (s *ebitenSurface) Screenshot() *image.RGBA {
	width, height := s.GetSize()
	bounds := image.Rectangle{image.Point{0, 0}, image.Point{width, height}}
	image := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			image.Set(x, y, s.image.At(x, y))
		}
	}

	return image
}
