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
