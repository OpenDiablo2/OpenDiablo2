package d2surface

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type surfaceState struct {
	x    int
	y    int
	mode ebiten.CompositeMode
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
			mode: ebiten.CompositeModeSourceOver,
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

func (s *Surface) Pop() {
	count := len(s.stateStack)
	if count == 0 {
		panic("empty stack")
	}

	s.stateCurrent = s.stateStack[count-1]
	s.stateStack = s.stateStack[:count-1]
}

func (s *Surface) Render(image *ebiten.Image) error {
	opts := &ebiten.DrawImageOptions{CompositeMode: s.stateCurrent.mode}
	opts.GeoM.Translate(float64(s.stateCurrent.x), float64(s.stateCurrent.y))
	return s.image.DrawImage(image, opts)
}

func (s *Surface) DrawText(text string) {
	ebitenutil.DebugPrintAt(s.image, text, s.stateCurrent.x, s.stateCurrent.y)
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

func (s *Surface) GetSize() (int, int) {
	return s.image.Size()
}
