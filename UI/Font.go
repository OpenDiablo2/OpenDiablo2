package UI

import (
	"image/color"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
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
	fontSprite *Common.Sprite
	metrics    map[uint8]FontSize
}

// GetFont creates or loads an existing font
func GetFont(font string, palette Palettes.Palette, fileProvider Common.FileProvider) *Font {
	cacheItem, exists := fontCache[font+"_"+string(palette)]
	if exists {
		return cacheItem
	}
	newFont := CreateFont(font, palette, fileProvider)
	fontCache[font+"_"+string(palette)] = newFont
	return newFont
}

// CreateFont creates an instance of a MPQ Font
func CreateFont(font string, palette Palettes.Palette, fileProvider Common.FileProvider) *Font {
	result := &Font{
		metrics: make(map[uint8]FontSize),
	}
	result.fontSprite = fileProvider.LoadSprite(font+".dc6", palette)
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
	height = uint32(0)
	for _, ch := range text {
		metric := v.metrics[uint8(ch)]
		width += uint32(metric.Width)
		_, h := v.fontSprite.GetFrameSize(int(ch))
		height = Common.Max(height, h)
	}
	return
}

// Draw draws the font on the target surface
func (v *Font) Draw(x, y int, text string, color color.Color, target *ebiten.Image) {
	v.fontSprite.ColorMod = color
	v.fontSprite.Blend = false
	_, height := v.GetTextMetrics(text)
	for _, ch := range text {
		char := uint8(ch)
		metric := v.metrics[char]
		v.fontSprite.Frame = char
		v.fontSprite.MoveTo(x, y+int(height))
		v.fontSprite.Draw(target)
		x += int(metric.Width)
	}
}
