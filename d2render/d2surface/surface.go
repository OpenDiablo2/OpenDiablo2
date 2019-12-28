package d2surface

import (
	"fmt"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2corehelper"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type surfaceState struct {
	x      int
	y      int
	mode   ebiten.CompositeMode
	filter ebiten.Filter
	color  color.Color
}

type Surface struct {
	stateStack   []surfaceState
	stateCurrent surfaceState
	image        *ebiten.Image
}

func CreateSurface(image *ebiten.Image) *Surface {
	return &Surface{
		image: image,
		stateCurrent: surfaceState{
			filter: ebiten.FilterNearest,
			mode:   ebiten.CompositeModeSourceOver,
		},
	}
}

func (s *Surface) PushTranslation(x, y int) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.x += x
	s.stateCurrent.y += y
}

func (s *Surface) PushCompositeMode(mode ebiten.CompositeMode) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.mode = mode
}

func (s *Surface) PushFilter(filter ebiten.Filter) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.filter = filter
}

func (s *Surface) PushColor(color color.Color) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.color = color
}

func (s *Surface) Pop() {
	count := len(s.stateStack)
	if count == 0 {
		panic("empty stack")
	}

	s.stateCurrent = s.stateStack[count-1]
	s.stateStack = s.stateStack[:count-1]
}

func (s *Surface) PopN(n int) {
	for i := 0; i < n; i++ {
		s.Pop()
	}
}

func (s *Surface) Render(image *ebiten.Image) error {
	opts := &ebiten.DrawImageOptions{CompositeMode: s.stateCurrent.mode}
	opts.GeoM.Translate(float64(s.stateCurrent.x), float64(s.stateCurrent.y))
	opts.Filter = s.stateCurrent.filter
	if s.stateCurrent.color != nil {
		opts.ColorM = d2corehelper.ColorToColorM(s.stateCurrent.color)
	}

	return s.image.DrawImage(image, opts)
}

func (s *Surface) DrawText(format string, params ...interface{}) {
	ebitenutil.DebugPrintAt(s.image, fmt.Sprintf(format, params...), s.stateCurrent.x, s.stateCurrent.y)
}

func (s *Surface) DrawLine(x, y int, color color.Color) {
	ebitenutil.DrawLine(
		s.image,
		float64(s.stateCurrent.x),
		float64(s.stateCurrent.y),
		float64(s.stateCurrent.x+x),
		float64(s.stateCurrent.y+y),
		color,
	)
}

func (s *Surface) DrawRect(width, height int, color color.Color) {
	ebitenutil.DrawRect(
		s.image,
		float64(s.stateCurrent.x),
		float64(s.stateCurrent.y),
		float64(width),
		float64(height),
		color,
	)
}

func (s *Surface) Clear(color color.Color) error {
	return s.image.Fill(color)
}

func (s *Surface) GetSize() (int, int) {
	return s.image.Size()
}

func (s *Surface) GetDepth() int {
	return len(s.stateStack)
}
