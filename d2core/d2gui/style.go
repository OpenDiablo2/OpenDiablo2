package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type FontStyle int

const (
	FontStyle16Units FontStyle = iota
	FontStyle30Units
	FontStyle42Units
	FontStyleFormal10Static
	FontStyleFormal11Units
	FontStyleFormal12Static
)

type fontStyleConfig struct {
	fontBasePath string
	palettePath  string
}

var fontStyleConfigs = map[FontStyle]fontStyleConfig{
	FontStyle16Units:        {d2resource.Font16, d2resource.PaletteUnits},
	FontStyle30Units:        {d2resource.Font30, d2resource.PaletteUnits},
	FontStyle42Units:        {d2resource.Font42, d2resource.PaletteUnits},
	FontStyleFormal10Static: {d2resource.FontFormal10, d2resource.PaletteStatic},
	FontStyleFormal11Units:  {d2resource.FontFormal11, d2resource.PaletteUnits},
	FontStyleFormal12Static: {d2resource.FontFormal12, d2resource.PaletteStatic},
}

func loadFont(fontStyle FontStyle) (*d2asset.Font, error) {
	config := fontStyleConfigs[fontStyle]
	return d2asset.LoadFont(config.fontBasePath+".tbl", config.fontBasePath+".dc6", config.palettePath)
}

// type ButtonFlag int
//
// const (
// 	ButtonFlagToggle = 1 << iota
// )
//
// type ButtonStyle int
//
// const (
// 	ButtonStyleWide ButtonStyle = iota
// 	ButtonStyleMedium
// 	ButtonStyleNarrow
// 	ButtonStyleCancel
// 	ButtonStyleTall
// 	ButtonStyleShort
// 	ButtonStyleOkCancel
// )
//
// type buttonStyleConfig struct {
// 	segmentsX     int
// 	segmentsY     int
// 	spritePath    string
// 	palettePath   string
// 	fontStyle     FontStyle
// 	textOffset    int
// 	disabledFrame int
// 	flags         ButtonFlag
// }
//
// type ButtonLayout struct {
// 	XSegments        int              //1
// 	YSegments        int              // 1
// 	ResourceName     string           // Font Name
// 	PaletteName      string           // PaletteType
// 	Toggleable       bool             // false
// 	BaseFrame        int              // 0
// 	DisabledFrame    int              // -1
// 	FontPath         string           // ResourcePaths.FontExocet10
// 	ClickableRect    *image.Rectangle // nil
// 	AllowFrameChange bool             // true
// 	TextOffset       int              // 0
// }
//
// var buttonStyleConfigs = map[ButtonStyle]buttonStyleConfig{
// 	ButtonStyleWide:     {2, 1, d2resource.WideButtonBlank, d2resource.PaletteUnits, false, 0, -1, d2resource.FontExocet10, nil, true, 1},
// 	ButtonStyleShort:    {1, 1, d2resource.ShortButtonBlank, d2resource.PaletteUnits, false, 0, -1, d2resource.FontRediculous, nil, true, -1},
// 	ButtonStyleMedium:   {1, 1, d2resource.MediumButtonBlank, d2resource.PaletteUnits, false, 0, 0, d2resource.FontExocet10, nil, true, 0},
// 	ButtonStyleTall:     {1, 1, d2resource.TallButtonBlank, d2resource.PaletteUnits, false, 0, 0, d2resource.FontExocet10, nil, true, 5},
// 	ButtonStyleOkCancel: {1, 1, d2resource.CancelButton, d2resource.PaletteUnits, false, 0, -1, d2resource.FontRediculous, nil, true, 0},
// }
