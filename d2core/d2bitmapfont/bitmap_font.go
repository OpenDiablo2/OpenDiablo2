package d2bitmapfont

import (
	"encoding/binary"
	"image/color"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

func New(s d2interface.Sprite, table []byte, col color.Color) *BitmapFont {
	return &BitmapFont{
		Sprite: s,
		Table:  table,
		Color:  col,
	}
}

type Glyph struct {
	frame  int
	width  int
	height int
}

// BitmapFont represents a rasterized font, made from a font Table, sprite, and palette
type BitmapFont struct {
	Sprite d2interface.Sprite
	Table  []byte
	Glyphs map[rune]Glyph
	Color  color.Color
}

// SetColor sets the fonts Color
func (f *BitmapFont) SetColor(c color.Color) {
	f.Color = c
}

// GetTextMetrics returns the dimensions of the BitmapFont element in pixels
func (f *BitmapFont) GetTextMetrics(text string) (width, height int) {
	if f.Glyphs == nil {
		f.initGlyphs()
	}

	var (
		lineWidth  int
		lineHeight int
	)

	for _, c := range text {
		if c == '\n' {
			width = d2math.MaxInt(width, lineWidth)
			height += lineHeight
			lineWidth = 0
			lineHeight = 0
		} else if glyph, ok := f.Glyphs[c]; ok {
			lineWidth += glyph.width
			lineHeight = d2math.MaxInt(lineHeight, glyph.height)
		}
	}

	width = d2math.MaxInt(width, lineWidth)
	height += lineHeight

	return width, height
}

// RenderText prints a text using its configured style on a Surface (multi-lines are left-aligned, use label otherwise)
func (f *BitmapFont) RenderText(text string, target d2interface.Surface) error {
	if f.Glyphs == nil {
		f.initGlyphs()
	}

	f.Sprite.SetColorMod(f.Color)

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		var (
			lineHeight int
			lineLength int
		)

		for _, c := range line {
			glyph, ok := f.Glyphs[c]
			if !ok {
				continue
			}

			if err := f.Sprite.SetCurrentFrame(glyph.frame); err != nil {
				return err
			}

			f.Sprite.Render(target)

			lineHeight = d2math.MaxInt(lineHeight, glyph.height)
			lineLength++

			target.PushTranslation(glyph.width, 0)
		}

		target.PopN(lineLength)
		target.PushTranslation(0, lineHeight)
	}

	target.PopN(len(lines))

	return nil
}

func (f *BitmapFont) initGlyphs() {
	_, maxCharHeight := f.Sprite.GetFrameBounds()

	glyphs := make(map[rune]Glyph)

	for i := 12; i < len(f.Table); i += 14 {
		code := rune(binary.LittleEndian.Uint16(f.Table[i : i+2]))

		var glyph Glyph
		glyph.frame = int(binary.LittleEndian.Uint16(f.Table[i+8 : i+10]))
		glyph.width = int(f.Table[i+3])
		glyph.height = maxCharHeight

		glyphs[code] = glyph
	}

	f.Glyphs = glyphs
}
