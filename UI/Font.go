package UI

import (
	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
)

var fontCache = map[string]*Font{}

// FontSize represents the size of a character in a font
type FontSize struct {
	Width  uint8
	Height uint8
}

// Font represents a font
type Font struct {
	FontSprite *Common.Sprite
	Metrics    map[uint8]FontSize
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
		Metrics: make(map[uint8]FontSize),
	}
	result.FontSprite = fileProvider.LoadSprite(font+".dc6", palette)
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
		result.Metrics[fontData[i+8]] = fontSize
	}
	return result
}
