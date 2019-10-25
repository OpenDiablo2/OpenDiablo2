package OpenDiablo2

import (
	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
)

// MPQFontSize represents the size of a character in a font
type MPQFontSize struct {
	Width  uint8
	Height uint8
}

// MPQFont represents a font
type MPQFont struct {
	Engine     *Engine
	FontSprite *Common.Sprite
	Metrics    map[uint8]MPQFontSize
}

// CreateMPQFont creates an instance of a MPQ Font
func CreateMPQFont(engine *Engine, font string, palette Palettes.Palette) *MPQFont {
	result := &MPQFont{
		Engine:  engine,
		Metrics: make(map[uint8]MPQFontSize),
	}
	result.FontSprite = result.Engine.LoadSprite(font+".dc6", palette)
	woo := "Woo!\x01"
	fontData := result.Engine.GetFile(font + ".tbl")
	if string(fontData[0:5]) != woo {
		panic("No woo :(")
	}
	for i := 12; i < len(fontData); i += 14 {
		fontSize := MPQFontSize{
			Width:  fontData[i+3],
			Height: fontData[i+4],
		}
		result.Metrics[fontData[i+8]] = fontSize
	}
	return result
}
