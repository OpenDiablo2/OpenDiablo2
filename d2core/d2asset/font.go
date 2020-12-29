package d2asset

import (
	"encoding/binary"
	"image/color"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

type fontGlyph struct {
	frame  int
	width  int
	height int
}

// Font represents a displayable font
type Font struct {
	sheet  d2interface.Animation
	table  []byte
	glyphs map[rune]fontGlyph
	color  color.Color
}

// SetColor sets the fonts color
func (f *Font) SetColor(c color.Color) {
	f.color = c
}

// GetTextMetrics returns the dimensions of the Font element in pixels
func (f *Font) GetTextMetrics(text string) (width, height int) {
	if f.glyphs == nil {
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
		} else if glyph, ok := f.glyphs[c]; ok {
			lineWidth += glyph.width
			lineHeight = d2math.MaxInt(lineHeight, glyph.height)
		}
	}

	width = d2math.MaxInt(width, lineWidth)
	height += lineHeight

	return width, height
}

// RenderText prints a text using its configured style on a Surface (multi-lines are left-aligned, use label otherwise)
func (f *Font) RenderText(text string, target d2interface.Surface) error {
	if f.glyphs == nil {
		f.sheet.BindRenderer(target.Renderer())
		f.initGlyphs()
	}

	f.sheet.SetColorMod(f.color)

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		var (
			lineHeight int
			lineLength int
		)

		for _, c := range line {
			glyph, ok := f.glyphs[c]
			if !ok {
				continue
			}

			if err := f.sheet.SetCurrentFrame(glyph.frame); err != nil {
				return err
			}

			f.sheet.Render(target)

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

func (f *Font) initGlyphs() {
	_, maxCharHeight := f.sheet.GetFrameBounds()

	glyphs := make(map[rune]fontGlyph)

	for i := 12; i < len(f.table); i += 14 {
		code := rune(binary.LittleEndian.Uint16(f.table[i : i+2]))

		var glyph fontGlyph
		glyph.frame = int(binary.LittleEndian.Uint16(f.table[i+8 : i+10]))
		glyph.width = int(f.table[i+3])
		glyph.height = maxCharHeight

		glyphs[code] = glyph
	}

	f.glyphs = glyphs
}
