package ui

import (
	"image/color"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2helper"

	"github.com/OpenDiablo2/OpenDiablo2/d2data/datadict"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2render"

	"github.com/hajimehoshi/ebiten"
)

var fontCache = map[string]*Font{}

// FontSize represents the size of a character in a font
type FontSize struct {
	Width  uint8
	Height uint8
}

// Font represents a font
type Font struct {
	fontSprite *d2render.Sprite
	metrics    map[uint8]FontSize
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
		metrics: make(map[uint8]FontSize),
	}
	result.fontSprite = d2render.CreateSprite(fileProvider.LoadFile(font+".dc6"), datadict.Palettes[palette])
	woo := "Woo!\x01"
	fontData := fileProvider.LoadFile(font + ".tbl")
	if string(fontData[0:5]) != woo {
		panic("No woo :(")
	}
	for i := 12; i < len(fontData); i += 14 {
		fontSize := FontSize{
			Width:  fontData[i+3],
			Height: fontData[i+4],
		}
		result.metrics[fontData[i+8]] = fontSize
	}
	return result
}

// GetTextMetrics returns the size of the specified text
func (v *Font) GetTextMetrics(text string) (width, height uint32) {
	width = uint32(0)
	curWidth := uint32(0)
	height = uint32(0)
	maxCharHeight := uint32(0)
	for _, m := range v.fontSprite.Frames {
		maxCharHeight = d2helper.Max(maxCharHeight, uint32(m.Height))
	}
	for i := 0; i < len(text); i++ {
		ch := text[i]
		if ch == '\n' {
			width = d2helper.Max(width, curWidth)
			curWidth = 0
			height += maxCharHeight + 6
			continue
		}
		metric := v.metrics[uint8(ch)]
		curWidth += uint32(metric.Width)
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
			char := uint8(ch)
			metric := v.metrics[char]
			v.fontSprite.Frame = char
			v.fontSprite.MoveTo(xPos, y+int(v.fontSprite.Frames[char].Height))
			v.fontSprite.Draw(target)
			xPos += int(metric.Width)
		}

		if lineIdx >= len(lines)-1 {
			break
		}

		xPos = x
		y += int(maxCharHeight + 6)
	}
}
