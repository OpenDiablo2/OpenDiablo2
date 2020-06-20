package d2ui

import (
	"encoding/binary"
	"image/color"
	"strings"
	"unicode"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

var fontCache = map[string]*Font{}

// FontSize represents the size of a character in a font
type FontSize struct {
	Width  uint8
	Height uint8
}

// Font represents a font
type Font struct {
	fontSprite *Sprite
	fontTable  map[uint16]uint16
	metrics    map[uint16]FontSize
}

// GetFont creates or loads an existing font
func GetFont(fontPath string, palettePath string) *Font {
	cacheItem, exists := fontCache[fontPath+"_"+palettePath]
	if exists {
		return cacheItem
	}
	newFont := CreateFont(fontPath, palettePath)
	fontCache[fontPath+"_"+palettePath] = newFont
	return newFont
}

// CreateFont creates an instance of a MPQ Font
func CreateFont(font string, palettePath string) *Font {
	result := &Font{
		fontTable: make(map[uint16]uint16),
		metrics:   make(map[uint16]FontSize),
	}
	// bug: performance issue when using CJK fonts, because ten thousand frames will be rendered PER font
	animation, _ := d2asset.LoadAnimation(font+".dc6", palettePath)
	result.fontSprite, _ = LoadSprite(animation)
	woo := "Woo!\x01"
	fontData, err := d2asset.LoadFile(font + ".tbl")
	if err != nil {
		panic(err)
	}
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
func (v *Font) GetTextMetrics(text string) (width, height int) {
	width = 0
	curWidth := 0
	height = 0
	_, maxCharHeight := v.fontSprite.GetFrameBounds()
	for _, ch := range text {
		if ch == '\n' {
			width = d2common.MaxInt(width, curWidth)
			curWidth = 0
			height += maxCharHeight + 6
			continue
		}

		curWidth += v.getCharWidth(ch)
	}
	width = d2common.MaxInt(width, curWidth)
	height += maxCharHeight
	return
}

// Render draws the font on the target surface
func (v *Font) Render(x, y int, text string, color color.Color, target d2interface.Surface) {
	v.fontSprite.SetColorMod(color)
	v.fontSprite.SetBlend(false)

	maxCharHeight := uint32(0)
	for _, m := range v.metrics {
		maxCharHeight = d2common.Max(maxCharHeight, uint32(m.Height))
	}

	targetWidth, _ := target.GetSize()
	lines := strings.Split(text, "\n")
	for lineIdx, line := range lines {
		lineWidth, _ := v.GetTextMetrics(line)
		xPos := x + ((targetWidth / 2) - lineWidth/2)

		for _, ch := range line {
			width := v.getCharWidth(ch)
			index := v.fontTable[uint16(ch)]
			v.fontSprite.SetCurrentFrame(int(index))
			_, height := v.fontSprite.GetCurrentFrameSize()
			v.fontSprite.SetPosition(xPos, y+height)
			v.fontSprite.Render(target)
			xPos += width
		}

		if lineIdx >= len(lines)-1 {
			break
		}

		xPos = x
		y += int(maxCharHeight + 6)
	}
}

func (v *Font) getCharWidth(char rune) (width int) {
	if char < unicode.MaxLatin1 {
		return int(v.metrics[uint16(char)].Width)
	}
	return int(v.metrics[unicode.MaxLatin1].Width)
}
