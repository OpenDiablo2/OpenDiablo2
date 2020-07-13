package d2asset

import (
	"encoding/binary"
	"errors"
	"image/color"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var _ d2interface.Font = &Font{} // Static check to confirm struct conforms to interface

type fontGlyph struct {
	frame  int
	width  int
	height int
}

// Font represents a displayable font
type Font struct {
	sheet  d2interface.Animation
	glyphs map[rune]fontGlyph
	color  color.Color
}

func loadFont(tablePath, spritePath, palettePath string) (d2interface.Font, error) {
	sheet, err := LoadAnimation(spritePath, palettePath)
	if err != nil {
		return nil, err
	}

	data, err := LoadFile(tablePath)
	if err != nil {
		return nil, err
	}

	if string(data[:5]) != "Woo!\x01" {
		return nil, errors.New("invalid font table format")
	}

	_, maxCharHeight := sheet.GetFrameBounds()

	glyphs := make(map[rune]fontGlyph)

	for i := 12; i < len(data); i += 14 {
		code := rune(binary.LittleEndian.Uint16(data[i : i+2]))

		var glyph fontGlyph
		glyph.frame = int(binary.LittleEndian.Uint16(data[i+8 : i+10]))
		glyph.width = int(data[i+3])
		glyph.height = maxCharHeight

		glyphs[code] = glyph
	}

	font := &Font{
		sheet:  sheet,
		glyphs: glyphs,
		color:  color.White,
	}

	return font, nil
}

// SetColor sets the fonts color
func (f *Font) SetColor(c color.Color) {
	f.color = c
}

// GetTextMetrics returns the dimensions of the Font element in pixels
func (f *Font) GetTextMetrics(text string) (width, height int) {
	var (
		lineWidth   int
		lineHeight  int
		totalWidth  int
		totalHeight int
	)

	for _, c := range text {
		if c == '\n' {
			totalWidth = d2common.MaxInt(totalWidth, lineWidth)
			totalHeight += lineHeight
			lineWidth = 0
			lineHeight = 0
		} else if glyph, ok := f.glyphs[c]; ok {
			lineWidth += glyph.width
			lineHeight = d2common.MaxInt(lineHeight, glyph.height)
		}
	}

	totalWidth = d2common.MaxInt(totalWidth, lineWidth)
	totalHeight += lineHeight

	return totalWidth, totalHeight
}

// RenderText prints a text using its configured style on a Surface (multi-lines are left-aligned, use label otherwise)
func (f *Font) RenderText(text string, target d2interface.Surface) error {
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

			if err := f.sheet.Render(target); err != nil {
				return err
			}

			lineHeight = d2common.MaxInt(lineHeight, glyph.height)
			lineLength++

			target.PushTranslation(glyph.width, 0)
		}

		target.PopN(lineLength)
		target.PushTranslation(0, lineHeight)
	}

	target.PopN(len(lines))

	return nil
}
