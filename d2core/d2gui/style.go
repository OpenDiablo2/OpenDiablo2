package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

// FontStyle is a representation of a font with a palette
type FontStyle int

// Font styles
const (
	FontStyle16Units FontStyle = iota
	FontStyle30Units
	FontStyle42Units
	FontStyleExocet10
	FontStyleFormal10Static
	FontStyleFormal11Units
	FontStyleFormal12Static
	FontStyleRediculous
)

type fontStyleConfig struct {
	fontBasePath string
	palettePath  string
}

func getFontStyleConfig(f FontStyle) *fontStyleConfig {
	fontStyles := map[FontStyle]*fontStyleConfig{
		FontStyle16Units:        {d2resource.Font16, d2resource.PaletteUnits},
		FontStyle30Units:        {d2resource.Font30, d2resource.PaletteUnits},
		FontStyle42Units:        {d2resource.Font42, d2resource.PaletteUnits},
		FontStyleExocet10:       {d2resource.FontExocet10, d2resource.PaletteUnits},
		FontStyleFormal10Static: {d2resource.FontFormal10, d2resource.PaletteStatic},
		FontStyleFormal11Units:  {d2resource.FontFormal11, d2resource.PaletteUnits},
		FontStyleFormal12Static: {d2resource.FontFormal12, d2resource.PaletteStatic},
		FontStyleRediculous:     {d2resource.FontRediculous, d2resource.PaletteUnits},
	}

	return fontStyles[f]
}

// ButtonStyle is a representation of a button style. Button styles have
// x and y sebment counts, an image, a palette, a font, and a text offset
type ButtonStyle int

// Button styles
const (
	ButtonStyleMedium ButtonStyle = iota
	ButtonStyleNarrow
	ButtonStyleOkCancel
	ButtonStyleShort
	ButtonStyleTall
	ButtonStyleWide
)

type buttonStyleConfig struct {
	segmentsX     int
	segmentsY     int
	animationPath string
	palettePath   string
	fontStyle     FontStyle
	textOffset    int
}

func getButtonStyleConfig(b ButtonStyle) *buttonStyleConfig {
	buttonStyleConfigs := map[ButtonStyle]*buttonStyleConfig{
		ButtonStyleMedium:   {1, 1, d2resource.MediumButtonBlank, d2resource.PaletteUnits, FontStyleExocet10, 0},
		ButtonStyleOkCancel: {1, 1, d2resource.CancelButton, d2resource.PaletteUnits, FontStyleRediculous, 0},
		ButtonStyleShort:    {1, 1, d2resource.ShortButtonBlank, d2resource.PaletteUnits, FontStyleRediculous, -1},
		ButtonStyleTall:     {1, 1, d2resource.TallButtonBlank, d2resource.PaletteUnits, FontStyleExocet10, 5},
		ButtonStyleWide:     {2, 1, d2resource.WideButtonBlank, d2resource.PaletteUnits, FontStyleExocet10, 1},
	}

	return buttonStyleConfigs[b]
}
