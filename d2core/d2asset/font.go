package d2asset

import (
	"encoding/binary"
	"errors"
	"image/color"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type fontGlyph struct {
	frame  int
	width  int
	height int
}

type Font struct {
	sheet  *Animation
	glyphs map[rune]fontGlyph
	color  color.Color
}

func loadFont(tablePath, spritePath, palettePath string) (*Font, error) {
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
		glyph.height = maxCharHeight // int(data[i+4])

		glyphs[code] = glyph
	}

	font := &Font{
		sheet:  sheet,
		glyphs: glyphs,
		color:  color.White,
	}

	return font, nil
}

func (f *Font) GetTextMetrics(text string) (int, int) {
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

func (f *Font) Clone() *Font {
	return &Font{
		sheet:  f.sheet,
		glyphs: f.glyphs,
		color:  f.color,
	}
}

func (f *Font) RenderText(text string, target d2render.Surface) error {
	f.sheet.SetColorMod(f.color)
	f.sheet.SetBlend(false)

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		var (
			lineHeight int
			lineLength int
		)

		for _, c := range line {
			if glyph, ok := f.glyphs[c]; ok {
				f.sheet.SetCurrentFrame(glyph.frame)
				f.sheet.Render(target)
				lineHeight = d2common.MaxInt(lineHeight, glyph.height)
				target.PushTranslation(glyph.width, 0)
				lineLength++
			}
		}

		target.PopN(lineLength)
		target.PushTranslation(0, lineHeight)
	}

	target.PopN(len(lines))

	return nil
}
