package d2util

import (
	"image"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cw = assets.CharWidth
	ch = assets.CharHeight
)

// NewDebugPrinter creates a new debug printer
func NewDebugPrinter() *GlyphPrinter {
	img := ebiten.NewImageFromImage(assets.CreateTextImage())

	printer := &GlyphPrinter{
		glyphImageTable: img,
		glyphs:          make(map[rune]*ebiten.Image),
	}

	return printer
}

// GlyphPrinter uses an image containing glyphs to draw text onto ebiten images
type GlyphPrinter struct {
	glyphImageTable *ebiten.Image
	glyphs          map[rune]*ebiten.Image
}

// Print draws the string str on the image on left top corner.
//
// The available runes are in U+0000 to U+00FF, which is C0 Controls and
// Basic Latin and C1 Controls and Latin-1 Supplement.
//
// DebugPrint always returns nil as of 1.5.0-alpha.
func (p *GlyphPrinter) Print(target *ebiten.Image, str string) error {
	p.PrintAt(target, str, 0, 0)
	return nil
}

// PrintAt draws the string str on the image at (x, y) position.
// The available runes are in U+0000 to U+00FF, which is C0 Controls and
// Basic Latin and C1 Controls and Latin-1 Supplement.
func (p *GlyphPrinter) PrintAt(target *ebiten.Image, str string, x, y int) {
	p.drawDebugText(target, str, x, y, false)
}

func (p *GlyphPrinter) drawDebugText(target *ebiten.Image, str string, ox, oy int, shadow bool) {
	op := &ebiten.DrawImageOptions{}

	if shadow {
		op.ColorM.Scale(0, 0, 0, 0.5)
	}

	x := 0
	y := 0

	w, _ := p.glyphImageTable.Size()

	for _, c := range str {
		if c == '\n' {
			x = 0
			y += ch

			continue
		}

		s, ok := p.glyphs[c]
		if !ok {
			n := w / cw
			sx := (int(c) % n) * cw
			sy := (int(c) / n) * ch
			rect := image.Rect(sx, sy, sx+cw, sy+ch)
			s = p.glyphImageTable.SubImage(rect).(*ebiten.Image)
			p.glyphs[c] = s
		}

		op.GeoM.Reset()
		op.GeoM.Translate(float64(x), float64(y))
		op.GeoM.Translate(float64(ox+1), float64(oy))

		op.CompositeMode = ebiten.CompositeModeLighter
		target.DrawImage(s, op)

		x += cw
	}
}
