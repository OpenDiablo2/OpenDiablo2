package d2font

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

const (
	knownSignature = "Woo!\x01"
)

type fontGlyph struct {
	unknown1 []byte
	unknown2 []byte
	unknown3 []byte
	frame    int
	width    int
	height   int
}

// Font represents a displayable font
type Font struct {
	unknownHeaderBytes []byte
	sheet              d2interface.Animation
	table              []byte
	glyphs             map[rune]fontGlyph
	color              color.Color
}

// Load loads a new font from byte slice
func Load(data []byte, sheet d2interface.Animation) (*Font, error) {
	sr := d2datautils.CreateStreamReader(data)

	signature, err := sr.ReadBytes(5)
	if err != nil {
		return nil, err
	}

	if string(signature) != knownSignature {
		return nil, fmt.Errorf("invalid font table format")
	}

	font := &Font{
		table: data,
		sheet: sheet,
		color: color.White,
	}

	font.unknownHeaderBytes, err = sr.ReadBytes(7)
	if err != nil {
		return nil, err
	}

	font.initGlyphs(sr)

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

func (f *Font) initGlyphs(sr *d2datautils.StreamReader) error {
	_, maxCharHeight := f.sheet.GetFrameBounds()

	glyphs := make(map[rune]fontGlyph)

	for i := 12; i < len(f.table); i += 14 {
		code, err := sr.ReadUInt16()
		if err != nil {
			return err
		}

		var glyph fontGlyph

		glyph.unknown1, err = sr.ReadBytes(1)
		if err != nil {
			return err
		}

		width, err := sr.ReadByte()
		if err != nil {
			return err
		}

		glyph.width = int(width)

		glyph.height = maxCharHeight

		glyph.unknown2, err = sr.ReadBytes(4)
		if err != nil {
			return err
		}

		frame, err := sr.ReadUInt16()
		if err != nil {
			return err
		}

		glyph.frame = int(frame)

		glyph.unknown3, err = sr.ReadBytes(4)
		if err != nil {
			return err
		}

		glyphs[rune(code)] = glyph
	}

	f.glyphs = glyphs

	return nil
}

func (f *Font) Marshal() []byte {
	sw := d2datautils.CreateStreamWriter()

	sw.PushBytes([]byte("Woo!\x01")...)
	sw.PushBytes(f.unknownHeaderBytes...)

	for c, i := range f.glyphs {
		sw.PushUint16(uint16(c))
		sw.PushBytes(i.unknown1...)
		sw.PushBytes(byte(i.width))
		sw.PushBytes(i.unknown2...)
		sw.PushUint16(uint16(i.frame))
		sw.PushBytes(i.unknown3...)
	}

	return sw.GetBytes()
}
