package d2font

import (
	//"os"

	//"encoding/binary"
	"fmt"
	"image/color"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
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

// Load loads a new font from byte slice
func Load(data []byte, sheet d2interface.Animation) (*Font, error) {
	if string(data[:5]) != "Woo!\x01" {
		return nil, fmt.Errorf("invalid font table format")
	}

	font := &Font{
		table: data,
		sheet: sheet,
		color: color.White,
	}

	font.initGlyphs()

	ok := true
	dw := font.Marshal()
	for i := range dw {
		if dw[i] != data[i] {
			ok = false
		}
	}

	_ = ok
	//fmt.Println(ok)
	//fmt.Println(len(data) == len(dw))
	//os.Exit(0)

	return font, nil
}

// SetColor sets the fonts color
func (f *Font) SetColor(c color.Color) {
	f.color = c
}

// GetTextMetrics returns the dimensions of the Font element in pixels
func (f *Font) GetTextMetrics(text string) (width, height int) {

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

func (f *Font) initGlyphs() error {
	sr := d2datautils.CreateStreamReader(f.table)
	sr.SkipBytes(12)

	_, maxCharHeight := f.sheet.GetFrameBounds()

	glyphs := make(map[rune]fontGlyph)

	for i := 12; i < len(f.table); i += 14 {
		sr.SetPosition(uint64(i))

		code, err := sr.ReadUInt16()
		if err != nil {
			return err
		}

		var glyph fontGlyph

		sr.SkipBytes(1)

		width, err := sr.ReadByte()
		if err != nil {
			return err
		}

		glyph.width = int(width)

		glyph.height = maxCharHeight

		sr.SkipBytes(4)

		frame, err := sr.ReadUInt16()
		if err != nil {
			return err
		}

		glyph.frame = int(frame)

		glyphs[rune(code)] = glyph
	}

	f.glyphs = glyphs

	return nil
}

func (f *Font) Marshal() []byte {
	sw := d2datautils.CreateStreamWriter()

	sw.PushBytes([]byte("Woo!\x01")...)

	return sw.GetBytes()
}
