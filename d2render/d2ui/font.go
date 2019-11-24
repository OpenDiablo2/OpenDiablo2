package d2ui

import (
	"github.com/OpenDiablo2/D2Shared/d2data/d2dc6"
	"image/color"
	"strings"

	"github.com/OpenDiablo2/D2Shared/d2helper"

	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"

	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"

	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2render"

	"github.com/hajimehoshi/ebiten"

	"encoding/binary"

	"unicode"
)

var fontCache = map[string]*Font{}

// FontSize represents the size of a character in a font
type FontSize struct {
	Width  uint8
	Height uint8
}

// Font represents a font
type Font struct {
	fontSprite d2render.Sprite
	fontTable  map[uint16]uint16
	metrics    map[uint16]FontSize
}

// GetFont creates or loads an existing font
func GetFont(font string, palette d2enum.PaletteType, fileProvider d2interface.FileProvider) *Font {
	cacheItem, exists := fontCache[font+"_"+string(palette)]
	if exists {
		return cacheItem
	}
	newFont := CreateFont(font, palette, fileProvider)
	fontCache[font+"_"+string(palette)] = newFont
	return newFont
}

// CreateFont creates an instance of a MPQ Font
func CreateFont(font string, palette d2enum.PaletteType, fileProvider d2interface.FileProvider) *Font {
	result := &Font{
		fontTable: make(map[uint16]uint16),
		metrics:   make(map[uint16]FontSize),
	}
	// bug: performance issue when using CJK fonts, because ten thousand frames will be rendered PER font
	dc6, _ := d2dc6.LoadDC6(fileProvider.LoadFile(font+".dc6"), d2datadict.Palettes[palette])
	result.fontSprite = d2render.CreateSpriteFromDC6(dc6)
	woo := "Woo!\x01"
	fontData := fileProvider.LoadFile(font + ".tbl")
	if string(fontData[0:5]) != woo {
		panic("No woo :(")
	}

	containsCjk := false
	for i := 12; i < len(fontData); i += 14 {
		// font mappings, map unicode code points to array indics
		unicodeCode := binary.LittleEndian.Uint16(fontData[i : i+2])
		fontIndex := binary.LittleEndian.Uint16(fontData[i+8 : i+10])
		result.fontTable[unicodeCode] = fontIndex

		if unicodeCode < unicode.MaxLatin1 {
			result.metrics[unicodeCode] = FontSize{
				Width:  fontData[i+3],
				Height: fontData[i+4],
			}
		} else if !containsCjk {
			// CJK characters are all in the same size
			result.metrics[unicode.MaxLatin1] = FontSize{
				Width:  fontData[i+3],
				Height: fontData[i+4],
			}
			containsCjk = true
		}
	}

	return result
}

// GetTextMetrics returns the size of the specified text
func (v *Font) GetTextMetrics(text string) (width, height uint32) {
	width = uint32(0)
	curWidth := uint32(0)
	height = uint32(0)
	maxCharHeight := uint32(0)
	// todo: it can be saved as a struct member, since it only depends on `.Frames`
	for _, m := range v.fontSprite.Frames {
		maxCharHeight = d2helper.Max(maxCharHeight, uint32(m.Height))
	}
	for _, ch := range text {
		if ch == '\n' {
			width = d2helper.Max(width, curWidth)
			curWidth = 0
			height += maxCharHeight + 6
			continue
		}

		curWidth += v.getCharWidth(ch)
	}
	width = d2helper.Max(width, curWidth)
	height += maxCharHeight
	return
}

// Draw draws the font on the target surface
func (v *Font) Draw(x, y int, text string, color color.Color, target *ebiten.Image) {
	v.fontSprite.ColorMod = color
	v.fontSprite.Blend = false

	maxCharHeight := uint32(0)
	for _, m := range v.metrics {
		maxCharHeight = d2helper.Max(maxCharHeight, uint32(m.Height))
	}

	targetWidth, _ := target.Size()
	lines := strings.Split(text, "\n")
	for lineIdx, line := range lines {
		lineWidth, _ := v.GetTextMetrics(line)
		xPos := x + ((targetWidth / 2) - int(lineWidth/2))

		for _, ch := range line {
			width := v.getCharWidth(ch)
			index := v.fontTable[uint16(ch)]
			v.fontSprite.Frame = int16(index)
			v.fontSprite.MoveTo(xPos, y+int(v.fontSprite.Frames[index].Height))
			v.fontSprite.Draw(target)
			xPos += int(width)
		}

		if lineIdx >= len(lines)-1 {
			break
		}

		xPos = x
		y += int(maxCharHeight + 6)
	}
}

func (v *Font) getCharWidth(char rune) (width uint32) {
	if char < unicode.MaxLatin1 {
		return uint32(v.metrics[uint16(char)].Width)
	}
	return uint32(v.metrics[unicode.MaxLatin1].Width)
}
