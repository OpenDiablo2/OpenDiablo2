package d2asset

import (
	"fmt"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"

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
	paletteBudget          = 64
	paletteTransformBudget = 64
)

const (
	defaultLanguage    = "ENG"
	logPrefix          = "Asset Manager"
	fmtLoadAsset       = "could not load file stream %s (%v)"
	fmtLoadAnimation   = "loading animation %s with palette %s, draw effect %d"
	fmtLoadComposite   = "loading composite: type %d, token %s, palette %s"
	fmtLoadFont        = "loading font: table %s, sprite %s, palette %s"
	fmtLoadPalette     = "loading palette %s"
	fmtLoadStringTable = "loading string table: %s"
	fmtLoadTransform   = "loading palette transform: %s"
	fmtLoadDict        = "loading data dictionary: %s"
)

// AssetManager loads files and game objects
type AssetManager struct {
	*d2util.Logger
	*d2loader.Loader
	tables     []d2tbl.TextDictionary
	animations d2interface.Cache
	fonts      d2interface.Cache
	palettes   d2interface.Cache
	transforms d2interface.Cache
	Records    *d2records.RecordManager
	language   string
}

// SetLogLevel sets the log level for the asset manager,  record manager, and file loader
func (am *AssetManager) SetLogLevel(level d2util.LogLevel) {
	am.Logger.SetLevel(level)
	am.Records.Logger.SetLevel(level)
	am.Loader.Logger.SetLevel(level)
}

// LoadAsset loads an asset
func (am *AssetManager) LoadAsset(filePath string) (asset.Asset, error) {
	data, err := am.Loader.Load(filePath)
	if err != nil {
		errStr := fmt.Sprintf(fmtLoadAsset, filePath, err.Error())

		am.Error(errStr)
	}

	return data, err
}

// LoadFileStream streams an MPQ file from a source file path
func (am *AssetManager) LoadFileStream(filePath string) (d2interface.DataStream, error) {
	am.Logger.Debugf("Loading FileStream: %s", filePath)
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
	am.Logger.Debugf("Checking if file exists %s", filePath)

	if loadedAsset, err := am.Loader.Load(filePath); err != nil || loadedAsset == nil {
		return false, err
	}

	return true, nil
}

// LoadLanguage loads language from resource path
func (am *AssetManager) LoadLanguage(languagePath string) string {
	languageByte, err := am.LoadFile(languagePath)
	if err != nil {
		am.Debugf("Unable to load language file: %s", err)
		return defaultLanguage
	}

	languageCode := languageByte[0]
	am.Debugf("Language code: %#02x", languageCode)

	language := d2resource.GetLanguageLiteral(languageCode)
	am.Infof("Language: %s", language)

	am.language = language

	return language
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

	am.Debugf(fmtLoadAnimation, animationPath, palettePath, effect)

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
	am.Debugf(fmtLoadComposite, baseType, token, palettePath)

	c := &Composite{
		AssetManager: am,
		baseType:     baseType,
		basePath:     baseString(baseType),
		token:        token,
		palettePath:  palettePath,
	}

	c.SetDirection(0)

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

	am.Debugf(fmtLoadFont, tablePath, spritePath, palettePath)

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

	am.Debugf(fmtLoadPalette, palettePath)

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
	data, err := am.LoadFile(tablePath)
	if err != nil {
		return nil, err
	}

	table := d2tbl.LoadTextDictionary(data)
	if table == nil {
		return nil, fmt.Errorf("table not found: %s", tablePath)
	}

	am.Debugf(fmtLoadStringTable, tablePath)

	am.tables = append(am.tables, table)

	return table, err
}

// TranslateString returns the translation of the given string. The string is retrieved from
// the loaded string tables.
func (am *AssetManager) TranslateString(input interface{}) string {
	var key string

	switch s := input.(type) {
	case string:
		key = s
	case fmt.Stringer:
		key = s.String()
	}

	for idx := range am.tables {
		if value, found := am.tables[idx][key]; found {
			return value
		}
	}

	// Fix to allow v.setDescLabels("#123") to be bypassed for a patch in issue #360. Reenable later.
	// log.Panicf("Could not find a string for the key '%s'", key)
	return key
}

// TranslateLabel translates the label taking into account its shift in the table
func (am *AssetManager) TranslateLabel(label int) string {
	return am.TranslateString(fmt.Sprintf("#%d", d2enum.BaseLabelNumbers(label+d2resource.GetLabelModifier(am.language))))
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

	am.Debugf(fmtLoadTransform, path)

	if err := am.transforms.Insert(path, pl2, 1); err != nil {
		return nil, err
	}

	return pl2, nil
}

// LoadDataDictionary loads a txt data file
func (am *AssetManager) LoadDataDictionary(path string) (*d2txt.DataDictionary, error) {
	// we purposefully do not cache data dictionaries because we are already
	// caching the file data. The underlying csv.Reader does not implement io.Seeker,
	// so after it has been iterated through, we cannot iterate through it again.
	//
	// The easy way around this is to not cache d2txt.DataDictionary objects, and just create
	// a new instance from cached file data if/when we ever need to reload the data dict
	data, err := am.LoadFile(path)
	if err != nil {
		return nil, err
	}

	am.Debugf(fmtLoadDict, path)

	return d2txt.LoadDataDictionary(data), nil
}

// LoadRecords will load the records for the given path into the record manager.
// This is dependant on the record manager having bound a loader for the given path.
func (am *AssetManager) LoadRecords(path string) error {
	dict, err := am.LoadDataDictionary(path)
	if err != nil {
		return err
	}

	err = am.Records.Load(path, dict)
	if err != nil {
		return err
	}

	return nil
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
	if err != nil {
		return nil, err
	}

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
