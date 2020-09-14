package d2asset

import (
	"fmt"
	"image/color"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dat"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2pl2"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2tbl"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"
)

const (
	defaultCacheEntryWeight = 1
)

const (
	animationBudget        = 1024 * 1024 * 128
	fontBudget             = 128
	tableBudget            = 64
	paletteBudget          = 64
	paletteTransformBudget = 64
)

// AssetManager loads files and game objects
type AssetManager struct {
	loader     *d2loader.Loader
	tables     d2interface.Cache
	animations d2interface.Cache
	fonts      d2interface.Cache
	palettes   d2interface.Cache
	transforms d2interface.Cache
}

// LoadAsset loads an asset
func (am *AssetManager) LoadAsset(filePath string) (asset.Asset, error) {
	data, err := am.loader.Load(filePath)
	if err != nil {
		log.Printf("error loading file stream %s (%v)", filePath, err.Error())
	}

	return data, err
}

// LoadFileStream streams an MPQ file from a source file path
func (am *AssetManager) LoadFileStream(filePath string) (d2interface.DataStream, error) {
	return am.LoadAsset(filePath)
}

// LoadFile loads an entire file from a source file path as a []byte
func (am *AssetManager) LoadFile(filePath string) ([]byte, error) {
	fileAsset, err := am.LoadAsset(filePath)
	if err != nil {
		return nil, err
	}

	data, err := fileAsset.Data()
	if err != nil {
		return nil, err
	}

	return data, err
}

// FileExists checks if a file exists on the underlying file system at the given file path.
func (am *AssetManager) FileExists(filePath string) (bool, error) {
	if loadedAsset, err := am.loader.Load(filePath); err != nil || loadedAsset == nil {
		return false, err
	}

	return true, nil
}

// LoadAnimation loads an Animation by its resource path and its palette path
func (am *AssetManager) LoadAnimation(animationPath, palettePath string) (d2interface.Animation, error) {
	return am.LoadAnimationWithEffect(animationPath, palettePath, d2enum.DrawEffectNone)
}

// LoadAnimationWithEffect loads an Animation by its resource path and its palette path with a given transparency value
func (am *AssetManager) LoadAnimationWithEffect(animationPath, palettePath string,
	effect d2enum.DrawEffect) (d2interface.Animation, error) {
	cachePath := fmt.Sprintf("%s;%s;%d", animationPath, palettePath, effect)

	if animation, found := am.animations.Retrieve(cachePath); found {
		return animation.(d2interface.Animation).Clone(), nil
	}

	animAsset, err := am.LoadAsset(animationPath)
	if err != nil {
		return nil, err
	}

	palette, err := am.LoadPalette(palettePath)
	if err != nil {
		return nil, err
	}

	var animation d2interface.Animation

	switch animAsset.Type() {
	case types.AssetTypeDC6:
		animation, err = am.loadDC6(animationPath, palette, effect)
		if err != nil {
			return nil, err
		}
	case types.AssetTypeDCC:
		animation, err = am.loadDCC(animationPath, palette, effect)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown Animation format for file: %s", animAsset.Path())
	}

	err = am.animations.Insert(cachePath, animation, defaultCacheEntryWeight)

	return animation, err
}

// LoadComposite creates a composite object from a ObjectLookupRecord and palettePath describing it
func (am *AssetManager) LoadComposite(baseType d2enum.ObjectType, token, palettePath string) (*Composite, error) {
	c := &Composite{
		AssetManager: am,
		baseType:     baseType,
		basePath:     baseString(baseType),
		token:        token,
		palettePath:  palettePath,
	}

	return c, nil
}

// LoadFont loads a font the resource files
func (am *AssetManager) LoadFont(tablePath, spritePath, palettePath string) (*Font, error) {
	cachePath := fmt.Sprintf("%s;%s;%s", tablePath, spritePath, palettePath)

	if cached, found := am.fonts.Retrieve(cachePath); found {
		return cached.(*Font), nil
	}

	sheet, err := am.LoadAnimation(spritePath, palettePath)
	if err != nil {
		return nil, err
	}

	tableData, err := am.LoadFile(tablePath)
	if err != nil {
		return nil, err
	}

	if string(tableData[:5]) != "Woo!\x01" {
		return nil, fmt.Errorf("invalid font table format: %s", tablePath)
	}

	font := &Font{
		table: tableData,
		sheet: sheet,
		color: color.White,
	}

	err = am.fonts.Insert(cachePath, font, defaultCacheEntryWeight)

	return font, err
}

// LoadPalette loads a palette from a given palette path
func (am *AssetManager) LoadPalette(palettePath string) (d2interface.Palette, error) {
	if cached, found := am.palettes.Retrieve(palettePath); found {
		return cached.(d2interface.Palette), nil
	}

	paletteAsset, err := am.LoadAsset(palettePath)
	if err != nil {
		return nil, err
	}

	if paletteAsset.Type() != types.AssetTypePalette {
		return nil, fmt.Errorf("not an instance of a palette: %s", palettePath)
	}

	data, err := am.LoadFile(palettePath)
	if err != nil {
		return nil, err
	}

	palette, err := d2dat.Load(data)
	if err != nil {
		return nil, err
	}

	err = am.palettes.Insert(palettePath, palette, defaultCacheEntryWeight)

	return palette, err
}

// LoadStringTable loads a string table from the given path
func (am *AssetManager) LoadStringTable(tablePath string) (d2tbl.TextDictionary, error) {
	if cached, found := am.tables.Retrieve(tablePath); found {
		return cached.(d2tbl.TextDictionary), nil
	}

	data, err := am.LoadFile(tablePath)
	if err != nil {
		return nil, err
	}

	table := d2tbl.LoadTextDictionary(data)
	if table == nil {
		return nil, fmt.Errorf("table not found: %s", tablePath)
	}

	err = am.tables.Insert(tablePath, table, defaultCacheEntryWeight)

	return table, err
}

// LoadPaletteTransform loads a palette transform file
func (am *AssetManager) LoadPaletteTransform(path string) (*d2pl2.PL2, error) {
	if pl2, found := am.transforms.Retrieve(path); found {
		return pl2.(*d2pl2.PL2), nil
	}

	data, err := am.LoadFile(path)
	if err != nil {
		return nil, err
	}

	pl2, err := d2pl2.Load(data)
	if err != nil {
		return nil, err
	}

	if err := am.transforms.Insert(path, pl2, 1); err != nil {
		return nil, err
	}

	return pl2, nil
}

// loadDC6 creates an Animation from d2dc6.DC6 and d2dat.DATPalette
func (am *AssetManager) loadDC6(path string,
	palette d2interface.Palette, effect d2enum.DrawEffect) (d2interface.Animation, error) {
	dc6Data, err := am.LoadFile(path)
	if err != nil {
		return nil, err
	}

	dc6, err := d2dc6.Load(dc6Data)
	if err != nil {
		return nil, err
	}

	animation, err := newDC6Animation(dc6, palette, effect)

	return animation, err
}

// loadDCC creates an Animation from d2dcc.DCC and d2dat.DATPalette
func (am *AssetManager) loadDCC(path string,
	palette d2interface.Palette, effect d2enum.DrawEffect) (d2interface.Animation, error) {
	dccData, err := am.LoadFile(path)
	if err != nil {
		return nil, err
	}

	dcc, err := d2dcc.Load(dccData)
	if err != nil {
		return nil, err
	}

	animation, err := newDCCAnimation(dcc, palette, effect)

	return animation, nil
}

// BindTerminalCommands binds the in-game terminal comands for the asset manager.
func (am *AssetManager) BindTerminalCommands(term d2interface.Terminal) error {
	if err := term.BindAction("assetspam", "display verbose asset manager logs", func(verbose bool) {
		if verbose {
			term.OutputInfof("asset manager verbose logging enabled")
		} else {
			term.OutputInfof("asset manager verbose logging disabled")
		}

		am.palettes.SetVerbose(verbose)
		am.fonts.SetVerbose(verbose)
		am.transforms.SetVerbose(verbose)
		am.animations.SetVerbose(verbose)
	}); err != nil {
		return err
	}

	if err := term.BindAction("assetstat", "display asset manager cache statistics", func() {
		var cacheStatistics = func(c d2interface.Cache) float64 {
			const percent = 100.0
			return float64(c.GetWeight()) / float64(c.GetBudget()) * percent
		}

		term.OutputInfof("palette cache: %f", cacheStatistics(am.palettes))
		term.OutputInfof("palette transform cache: %f", cacheStatistics(am.transforms))
		term.OutputInfof("Animation cache: %f", cacheStatistics(am.animations))
		term.OutputInfof("font cache: %f", cacheStatistics(am.fonts))
	}); err != nil {
		return err
	}

	if err := term.BindAction("assetclear", "clear asset manager cache", func() {
		am.palettes.Clear()
		am.transforms.Clear()
		am.animations.Clear()
		am.fonts.Clear()
	}); err != nil {
		return err
	}

	return nil
}
