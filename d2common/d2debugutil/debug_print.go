package d2debugutil

import (
	"image"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2debugutil/internal/assets"
	"github.com/hajimehoshi/ebiten"
)

var (
	debugPrintTextImage     *ebiten.Image
	debugPrintTextSubImages = map[rune]*ebiten.Image{}
)

func init() {
	img := assets.CreateTextImage()
	debugPrintTextImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
}

// DebugPrint draws the string str on the image on left top corner.
//
// The available runes are in U+0000 to U+00FF, which is C0 Controls and Basic Latin and C1 Controls and Latin-1 Supplement.
//
// DebugPrint always returns nil as of 1.5.0-alpha.
func D2DebugPrint(image *ebiten.Image, str string) error {
	D2DebugPrintAt(image, str, 0, 0)
	return nil
}

// DebugPrintAt draws the string str on the image at (x, y) position.
//
// The available runes are in U+0000 to U+00FF, which is C0 Controls and Basic Latin and C1 Controls and Latin-1 Supplement.
func D2DebugPrintAt(image *ebiten.Image, str string, x, y int) {
	drawDebugText(image, str, x, y, false)
}

func drawDebugText(rt *ebiten.Image, str string, ox, oy int, shadow bool) {
	op := &ebiten.DrawImageOptions{}
	if shadow {
		op.ColorM.Scale(0, 0, 0, 0.5)
	}
	x := 0
	y := 0
	w, _ := debugPrintTextImage.Size()
	for _, c := range str {
		const (
			cw = assets.CharWidth
			ch = assets.CharHeight
		)
		if c == '\n' {
			x = 0
			y += ch
			continue
		}
		s, ok := debugPrintTextSubImages[c]
		if !ok {
			n := w / cw
			sx := (int(c) % n) * cw
			sy := (int(c) / n) * ch
			s = debugPrintTextImage.SubImage(image.Rect(sx, sy, sx+cw, sy+ch)).(*ebiten.Image)
			debugPrintTextSubImages[c] = s
		}
		op.GeoM.Reset()
		op.GeoM.Translate(float64(x), float64(y))
		op.GeoM.Translate(float64(ox+1), float64(oy))
		op.CompositeMode = ebiten.CompositeModeLighter
		_ = rt.DrawImage(s, op)
		x += cw
	}
}
