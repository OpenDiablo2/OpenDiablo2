package d2asset

import (
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2font"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dat"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2pl2"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2tbl"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader"
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
	dt1Budget              = 4096 * 2048 * 128
	ds1Budget              = 4096 * 2048 * 128
	cofBudget              = 4096 * 2048 * 128
	dccBudget              = 4096 * 2048 * 128
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

	tables           []d2tbl.TextDictionary
	dt1s             d2interface.Cache
	ds1s             d2interface.Cache
	cofs             d2interface.Cache
	dccs             d2interface.Cache
	animations       d2interface.Cache
	fonts            d2interface.Cache
	palettes         d2interface.Cache
	transforms       d2interface.Cache
	Records          *d2records.RecordManager
	language         string
	languageModifier int
}

// SetLogLevel sets the log level for the asset manager,  record manager, and file loader
func (am *AssetManager) SetLogLevel(level d2util.LogLevel) {
	am.Logger.SetLevel(level)
	am.Records.Logger.SetLevel(level)
	am.Loader.Logger.SetLevel(level)
}

// LoadAsset loads an asset
func (am *AssetManager) LoadAsset(filePath string) (io.ReadSeeker, error) {
	data, err := am.Loader.Load(filePath)
	if err != nil {
		errStr := fmt.Sprintf(fmtLoadAsset, filePath, err.Error())

		am.Error(errStr)
	}

	return data, err
}

// LoadFileStream streams an MPQ file from a source file path
func (am *AssetManager) LoadFileStream(filePath string) (io.ReadSeeker, error) {
	am.Logger.Debugf("Loading FileStream: %s", filePath)
	return am.LoadAsset(filePath)
}

// LoadFile loads an entire file from a source file path as a []byte
func (am *AssetManager) LoadFile(filePath string) ([]byte, error) { // I DO NOT LIKE THIS! - Essial
	fileAsset, err := am.LoadAsset(filePath)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(fileAsset)
	if err != nil {
		return nil, err
	}

	return data, err
}

// FileExists checks if a file exists on the underlying file system at the given file path.
func (am *AssetManager) FileExists(filePath string) (bool, error) {
	filePath = filepath.Clean(filePath)

	am.Logger.Debugf("Checking if file exists %s", filePath)

	return am.Loader.Exists(filePath), nil
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
	am.languageModifier = d2resource.GetLabelModifier(language)

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

	palette, err := am.LoadPalette(palettePath)
	if err != nil {
		return nil, err
	}

	var animation d2interface.Animation

	switch types.Ext2AssetType(filepath.Ext(animationPath)) {
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
		return nil, fmt.Errorf("unknown Animation format for file: %s", animationPath)
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
func (am *AssetManager) LoadFont(tablePath, spritePath, palettePath string) (*d2font.Font, error) {
	cachePath := fmt.Sprintf("%s;%s;%s", tablePath, spritePath, palettePath)

	if cached, found := am.fonts.Retrieve(cachePath); found {
		return cached.(*d2font.Font), nil
	}

	sheet, err := am.LoadAnimation(spritePath, palettePath)
	if err != nil {
		return nil, err
	}

	tableData, err := am.LoadFile(tablePath)
	if err != nil {
		return nil, err
	}

	am.Debugf(fmtLoadFont, tablePath, spritePath, palettePath)

	font, err := d2font.Load(tableData)
	if err != nil {
		return nil, fmt.Errorf("error while loading font table %s: %v", tablePath, err)
	}

	font.SetBackground(sheet)

	err = am.fonts.Insert(cachePath, font, defaultCacheEntryWeight)

	return font, err
}

// LoadPalette loads a palette from a given palette path
func (am *AssetManager) LoadPalette(palettePath string) (d2interface.Palette, error) {
	if cached, found := am.palettes.Retrieve(palettePath); found {
		return cached.(d2interface.Palette), nil
	}

	if types.Ext2AssetType(filepath.Ext(palettePath)) != types.AssetTypePalette {
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

	table, err := d2tbl.LoadTextDictionary(data)
	if err != nil {
		return table, err
	}

	am.Debugf(fmtLoadStringTable, tablePath)

	am.tables = append(am.tables, table)

	return table, err
}

// TranslateString returns the translation of the given string. The string is retrieved from
// the loaded string tables. If input value is int (e.g. from d2enum/numeric_labels.go)
// output string is translation for # + input
func (am *AssetManager) TranslateString(input interface{}) string {
	var key string

	switch s := input.(type) {
	case string:
		key = s
	case fmt.Stringer:
		key = s.String()
	case int:
		key = fmt.Sprintf("#%d", d2enum.BaseLabelNumbers(s+am.languageModifier))
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
	dcc, err := am.LoadDCC(path)
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
	if err := term.Bind("assetspam", "display verbose asset manager logs", nil, am.commandAssetSpam(term)); err != nil {
		return err
	}

	if err := term.Bind("assetstat", "display asset manager cache statistics", nil, am.commandAssetStat(term)); err != nil {
		return err
	}

	if err := term.Bind("assetclear", "clear asset manager cache", nil, am.commandAssetClear); err != nil {
		return err
	}

	return nil
}

// UnbindTerminalCommands unbinds commands from the terminal
func (am *AssetManager) UnbindTerminalCommands(term d2interface.Terminal) error {
	return term.Unbind("assetspam", "assetstat", "assetclear")
}

func (am *AssetManager) commandAssetSpam(term d2interface.Terminal) func([]string) error {
	return func(args []string) error {
		verbose, err := strconv.ParseBool(args[0])
		if err != nil {
			term.Errorf("asset manager verbose invalid argument")
			return nil
		}

		if verbose {
			term.Infof("asset manager verbose logging enabled")
		} else {
			term.Infof("asset manager verbose logging disabled")
		}

		am.palettes.SetVerbose(verbose)
		am.fonts.SetVerbose(verbose)
		am.transforms.SetVerbose(verbose)
		am.animations.SetVerbose(verbose)
		am.dt1s.SetVerbose(verbose)
		am.ds1s.SetVerbose(verbose)
		am.dccs.SetVerbose(verbose)
		am.cofs.SetVerbose(verbose)

		return nil
	}
}

func (am *AssetManager) commandAssetStat(term d2interface.Terminal) func([]string) error {
	return func([]string) error {
		var cacheStatistics = func(c d2interface.Cache) float64 {
			const percent = 100.0
			return float64(c.GetWeight()) / float64(c.GetBudget()) * percent
		}

		term.Infof("palette cache: %f", cacheStatistics(am.palettes))
		term.Infof("palette transform cache: %f", cacheStatistics(am.transforms))
		term.Infof("Animation cache: %f", cacheStatistics(am.animations))
		term.Infof("font cache: %f", cacheStatistics(am.fonts))

		return nil
	}
}

func (am *AssetManager) commandAssetClear([]string) error {
	am.palettes.Clear()
	am.transforms.Clear()
	am.animations.Clear()
	am.fonts.Clear()
	am.dt1s.Clear()
	am.ds1s.Clear()
	am.dccs.Clear()
	am.cofs.Clear()

	return nil
}

// LoadDT1 loads and returns the given path as a DT1
func (am *AssetManager) LoadDT1(dt1Path string) (*d2dt1.DT1, error) {
	if dt1Value, found := am.dt1s.Retrieve(dt1Path); found {
		return dt1Value.(*d2dt1.DT1), nil
	}

	fileData, err := am.LoadFile("/data/global/tiles/" + dt1Path)
	if err != nil {
		return nil, fmt.Errorf("could not load /data/global/tiles/%s", dt1Path)
	}

	dt1, err := d2dt1.LoadDT1(fileData)
	if err != nil {
		return nil, err
	}

	if err := am.dt1s.Insert(dt1Path, dt1, defaultCacheEntryWeight); err != nil {
		return nil, err
	}

	return dt1, nil
}

// LoadDS1 loads and returns the given path as a DS1
func (am *AssetManager) LoadDS1(ds1Path string) (*d2ds1.DS1, error) {
	if ds1Value, found := am.dt1s.Retrieve(ds1Path); found {
		return ds1Value.(*d2ds1.DS1), nil
	}

	fileData, err := am.LoadFile("/data/global/tiles/" + ds1Path)
	if err != nil {
		return nil, err
	}

	ds1, err := d2ds1.LoadDS1(fileData)
	if err != nil {
		return nil, err
	}

	if err := am.dt1s.Insert(ds1Path, ds1, defaultCacheEntryWeight); err != nil {
		return nil, err
	}

	return ds1, nil
}

// LoadCOF loads and returns the given path as a COF
func (am *AssetManager) LoadCOF(cofPath string) (*d2cof.COF, error) {
	if cofValue, found := am.cofs.Retrieve(cofPath); found {
		return cofValue.(*d2cof.COF), nil
	}

	fileData, err := am.LoadFile(cofPath)
	if err != nil {
		return nil, err
	}

	cof, err := d2cof.Unmarshal(fileData)
	if err != nil {
		return nil, err
	}

	if err := am.cofs.Insert(cofPath, cof, defaultCacheEntryWeight); err != nil {
		return nil, err
	}

	return cof, nil
}

// LoadDCC loads and returns the given path as a DCC
func (am *AssetManager) LoadDCC(dccPath string) (*d2dcc.DCC, error) {
	if dccValue, found := am.dccs.Retrieve(dccPath); found {
		return dccValue.(*d2dcc.DCC), nil
	}

	fileData, err := am.LoadFile(dccPath)
	if err != nil {
		return nil, err
	}

	dcc, err := d2dcc.Load(fileData)
	if err != nil {
		return nil, err
	}

	if err := am.dccs.Insert(dccPath, dcc, defaultCacheEntryWeight); err != nil {
		return nil, err
	}

	return dcc, nil
}
