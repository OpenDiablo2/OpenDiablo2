package d2asset

import (
	"encoding/binary"
	"errors"
	"fmt"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

const (
	fontBudget = 64
)

// Static checks to confirm struct conforms to interface
var _ d2interface.FontManager = &fontManager{}
var _ d2interface.Cacher = &fontManager{}

type fontManager struct {
	*AssetManager
	cache d2interface.Cache
}

// LoadFont loads a font from the archives managed by the ArchiveManager
func (fm *fontManager) LoadFont(tablePath, spritePath, palettePath string) (d2interface.Font,
	error) {
	cachePath := fmt.Sprintf("%s;%s;%s", tablePath, spritePath, palettePath)
	if font, found := fm.cache.Retrieve(cachePath); found {
		return font.(d2interface.Font), nil
	}

	sheet, err := fm.LoadAnimation(spritePath, palettePath)
	if err != nil {
		return nil, err
	}

	data, err := fm.LoadFile(tablePath)
	if err != nil {
		return nil, err
	}

	if string(data[:5]) != "Woo!\x01" {
		return nil, errors.New("invalid font table format")
	}

	_, maxCharHeight := sheet.GetFrameBounds()

	glyphs := make(map[rune]fontGlyph)

	for i := 12; i < len(data); i += 14 {
		code := rune(binary.LittleEndian.Uint16(data[i : i+2]))

		var glyph fontGlyph
		glyph.frame = int(binary.LittleEndian.Uint16(data[i+8 : i+10]))
		glyph.width = int(data[i+3])
		glyph.height = maxCharHeight

		glyphs[code] = glyph
	}

	font := &Font{
		sheet:  sheet,
		glyphs: glyphs,
		color:  color.White,
	}

	if err != nil {
		return nil, err
	}

	if err := fm.cache.Insert(cachePath, font, 1); err != nil {
		return nil, err
	}

	return font, nil
}

// ClearCache clears the font cache
func (fm *fontManager) ClearCache() {
	fm.cache.Clear()
}

// GetCache returns the font managers cache
func (fm *fontManager) GetCache() d2interface.Cache {
	return fm.cache
}
