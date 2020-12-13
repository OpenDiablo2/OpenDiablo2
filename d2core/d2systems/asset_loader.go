package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"io"
	"time"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2animdata"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2cof"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dat"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dcc"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2pl2"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2tbl"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	assetCacheBudget      = 1024
	assetCacheEntryWeight = 1 // may want to make different weights for different asset types
)

const (
	logPrefixAssetLoader = "Asset Loader System"
)

const (
	maxTimePerUpdate = time.Millisecond * 16
)

var _ akara.System = &AssetLoaderSystem{}

// AssetLoaderSystem is responsible for parsing data from file handles into various structs, like COF or DC6
type AssetLoaderSystem struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	fileSub   *akara.Subscription
	sourceSub *akara.Subscription
	gameConfigs *akara.Subscription
	cache     *d2cache.Cache
	localeString string // related to file "/data/local/use"
	Components struct {
		File d2components.FileFactory
		FileType d2components.FileTypeFactory
		FileHandle d2components.FileHandleFactory
		FileSource d2components.FileSourceFactory
		GameConfig d2components.GameConfigFactory
		StringTable d2components.StringTableFactory
		FontTable d2components.FontTableFactory
		DataDictionary d2components.DataDictionaryFactory
		Palette d2components.PaletteFactory
		PaletteTransform d2components.PaletteTransformFactory
		Cof d2components.CofFactory
		Dc6 d2components.Dc6Factory
		Dcc d2components.DccFactory
		Ds1 d2components.Ds1Factory
		Dt1 d2components.Dt1Factory
		Wav d2components.WavFactory
		AnimationData d2components.AnimationDataFactory
		Locale d2components.LocaleFactory
		BitmapFont d2components.BitmapFontFactory
		FileLoaded d2components.FileLoadedFactory
	}
}

// Init injects component maps related to various asset types
func (m *AssetLoaderSystem) Init(world *akara.World) {
	m.World = world

	m.setupLogger()

	m.Debug("initializing ...")

	m.setupSubscriptions()
	m.setupFactories()

	m.cache = d2cache.CreateCache(assetCacheBudget).(*d2cache.Cache)
}

func (m *AssetLoaderSystem) setupLogger() {
	m.Logger = d2util.NewLogger()
	m.SetPrefix(logPrefixAssetLoader)
}

func (m *AssetLoaderSystem) setupSubscriptions() {
	m.Debug("setting up component subscriptions")

	// we are going to check entities that dont yet have loaded asset types
	filesToLoad := m.NewComponentFilter().
		Require(
			&d2components.File{}, // we want to process entities with these file components
			&d2components.FileType{},
			&d2components.FileHandle{},
		).
		Forbid(
			&d2components.FileSource{},
			&d2components.FileLoaded{},
		).
		Build()

	fileSources := m.NewComponentFilter().
		Require(&d2components.FileSource{}).
		Build()

	m.fileSub = m.World.AddSubscription(filesToLoad)
	m.sourceSub = m.World.AddSubscription(fileSources)

	gameConfigs := m.NewComponentFilter().Require(&d2components.GameConfig{}).Build()
	m.gameConfigs = m.World.AddSubscription(gameConfigs)
}

func (m *AssetLoaderSystem) setupFactories() {
	m.Debug("setting up component factories")

	m.InjectComponent(&d2components.File{}, &m.Components.File.ComponentFactory)
	m.InjectComponent(&d2components.FileType{}, &m.Components.FileType.ComponentFactory)
	m.InjectComponent(&d2components.FileHandle{}, &m.Components.FileHandle.ComponentFactory)
	m.InjectComponent(&d2components.FileSource{}, &m.Components.FileSource.ComponentFactory)
	m.InjectComponent(&d2components.GameConfig{}, &m.Components.GameConfig.ComponentFactory)
	m.InjectComponent(&d2components.StringTable{}, &m.Components.StringTable.ComponentFactory)
	m.InjectComponent(&d2components.FontTable{}, &m.Components.FontTable.ComponentFactory)
	m.InjectComponent(&d2components.DataDictionary{}, &m.Components.DataDictionary.ComponentFactory)
	m.InjectComponent(&d2components.Palette{}, &m.Components.Palette.ComponentFactory)
	m.InjectComponent(&d2components.PaletteTransform{}, &m.Components.PaletteTransform.ComponentFactory)
	m.InjectComponent(&d2components.Cof{}, &m.Components.Cof.ComponentFactory)
	m.InjectComponent(&d2components.Dc6{}, &m.Components.Dc6.ComponentFactory)
	m.InjectComponent(&d2components.Dcc{}, &m.Components.Dcc.ComponentFactory)
	m.InjectComponent(&d2components.Ds1{}, &m.Components.Ds1.ComponentFactory)
	m.InjectComponent(&d2components.Dt1{}, &m.Components.Dt1.ComponentFactory)
	m.InjectComponent(&d2components.Wav{}, &m.Components.Wav.ComponentFactory)
	m.InjectComponent(&d2components.AnimationData{}, &m.Components.AnimationData.ComponentFactory)
	m.InjectComponent(&d2components.Locale{}, &m.Components.Locale.ComponentFactory)
	m.InjectComponent(&d2components.BitmapFont{}, &m.Components.BitmapFont.ComponentFactory)
	m.InjectComponent(&d2components.FileLoaded{}, &m.Components.FileLoaded.ComponentFactory)
}

// Update processes all of the Entities in the subscription of file entities that need to be processed
func (m *AssetLoaderSystem) Update() {
	start := time.Now()

	for _, eid := range m.fileSub.GetEntities() {
		m.loadAsset(eid)
		if time.Since(start) > maxTimePerUpdate {
			break
		}
	}

	for _, eid := range m.gameConfigs.GetEntities() {
		cfg, found := m.Components.GameConfig.Get(eid)
		if !found {
			continue
		}

		m.SetLevel(cfg.LogLevel)
	}
}

func (m *AssetLoaderSystem) loadAsset(id akara.EID) {
	// make sure everything is kosher
	fp, found := m.Components.File.Get(id)
	if !found {
		m.Errorf("filepath component not found for entity %d", id)
		return
	}

	ft, found := m.Components.FileType.Get(id)
	if !found {
		m.Errorf("filetype component not found for entity %d", id)
		return
	}

	fh, found := m.Components.FileHandle.Get(id)
	if !found {
		m.Errorf("filehandle component not found for entity %d", id)
		return
	}

	// try to pull from the cache and assign to the given entity id
	if found := m.assignFromCache(id, fp.Path, ft.Type); found {
		m.Debugf("Retrieving %s from cache", fp.Path)
		return
	}

	m.Debugf("Loading file: %s", fp.Path)

	// make sure to seek back to 0 if the filehandle was cached
	_, _ = fh.Data.Seek(0, 0)

	data, buf := make([]byte, 0), make([]byte, 16)

	// read, parse, and cache the data
	for {
		numRead, err := fh.Data.Read(buf)
		data = append(data, buf[:numRead]...)

		if numRead < 1 || err != nil {
			break
		}
	}

	m.parseAndCache(id, fp.Path, ft.Type, data)
}

func (m *AssetLoaderSystem) assignFromCache(id akara.EID, path string, t d2enum.FileType) bool {
	entry, found := m.cache.Retrieve(path)
	if !found {
		return found
	}

	// if we found what we're looking for, create the appropriate component and assign what we retrieved
	switch t {
	case d2enum.FileTypeStringTable:
		m.Components.StringTable.Add(id).TextDictionary = entry.(*d2tbl.TextDictionary)
	case d2enum.FileTypeFontTable:
		m.Components.FontTable.Add(id).Data = entry.([]byte)
	case d2enum.FileTypeDataDictionary:
		m.Components.DataDictionary.Add(id).DataDictionary = entry.(*d2txt.DataDictionary)
	case d2enum.FileTypePalette:
		m.Components.Palette.Add(id).Palette = entry.(d2interface.Palette)
	case d2enum.FileTypePaletteTransform:
		m.Components.PaletteTransform.Add(id).PL2 = entry.(*d2pl2.PL2)
	case d2enum.FileTypeCOF:
		m.Components.Cof.Add(id).COF = entry.(*d2cof.COF)
	case d2enum.FileTypeDC6:
		m.Components.Dc6.Add(id).DC6 = entry.(*d2dc6.DC6)
	case d2enum.FileTypeDCC:
		m.Components.Dcc.Add(id).DCC = entry.(*d2dcc.DCC)
	case d2enum.FileTypeDS1:
		m.Components.Ds1.Add(id).DS1 = entry.(*d2ds1.DS1)
	case d2enum.FileTypeDT1:
		m.Components.Dt1.Add(id).DT1 = entry.(*d2dt1.DT1)
	case d2enum.FileTypeWAV:
		m.Components.Wav.Add(id).Data = entry.(d2interface.DataStream)
	case d2enum.FileTypeD2:
		m.Components.AnimationData.Add(id).AnimationData = entry.(*d2animdata.AnimationData)
	}

	m.Components.FileLoaded.Add(id)

	return found
}

//nolint:gocyclo // this big switch statement is unfortunate, but necessary
func (m *AssetLoaderSystem) parseAndCache(id akara.EID, path string, t d2enum.FileType, data []byte) {
	switch t {
	case d2enum.FileTypeStringTable:
		m.Debugf("Loading string table: %s", path)
		m.loadStringTable(id, path, data)
	case d2enum.FileTypeFontTable:
		m.Debugf("Loading font table: %s", path)
		m.loadFontTable(id, path, data)
	case d2enum.FileTypeDataDictionary:
		m.Debugf("Loading data dictionary: %s", path)
		m.loadDataDictionary(id, path, data)
	case d2enum.FileTypePalette:
		m.Debugf("Loading palette: %s", path)

		if err := m.loadPalette(id, path, data); err != nil {
			m.Error(err.Error())
		}
	case d2enum.FileTypePaletteTransform:
		m.Debugf("Loading palette transform: %s", path)

		if err := m.loadPaletteTransform(id, path, data); err != nil {
			m.Error(err.Error())
		}
	case d2enum.FileTypeCOF:
		m.Debugf("Loading COF: %s", path)

		if err := m.loadCOF(id, path, data); err != nil {
			m.Error(err.Error())
		}
	case d2enum.FileTypeDC6:
		m.Debugf("Loading DC6: %s", path)

		if err := m.loadDC6(id, path, data); err != nil {
			m.Error(err.Error())
		}
	case d2enum.FileTypeDCC:
		m.Debugf("Loading DCC: %s", path)

		if err := m.loadDCC(id, path, data); err != nil {
			m.Error(err.Error())
		}
	case d2enum.FileTypeDS1:
		m.Debugf("Loading DS1: %s", path)

		if err := m.loadDS1(id, path, data); err != nil {
			m.Error(err.Error())
		}
	case d2enum.FileTypeDT1:
		m.Debugf("Loading DT1: %s", path)

		if err := m.loadDT1(id, path, data); err != nil {
			m.Error(err.Error())
		}
	case d2enum.FileTypeWAV:
		m.Debugf("Loading WAV: %s", path)

		fh, found := m.Components.FileHandle.Get(id)
		if !found {
			return
		}

		m.loadWAV(id, path, fh.Data)
	case d2enum.FileTypeD2:
		m.Debugf("Loading animation data: %s", path)

		if err := m.loadAnimationData(id, path, data); err != nil {
			m.Error(err.Error())
		}
	case d2enum.FileTypeLocale:
		m.Debugf("Loading locale: %s", path)

		m.loadLocale(id, data)
	}

	m.Components.FileLoaded.Add(id)
}

func (m *AssetLoaderSystem) loadLocale(id akara.EID, data []byte) {
	locale := m.Components.Locale.Add(id)

	locale.Code = data[0]
	locale.String = d2resource.GetLanguageLiteral(locale.Code)

	m.localeString = locale.String
}

func (m *AssetLoaderSystem) loadStringTable(id akara.EID, path string, data []byte) {
	txt := d2tbl.LoadTextDictionary(data)
	loaded := &txt
	m.Components.StringTable.Add(id).TextDictionary = loaded

	if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
		m.Error(cacheErr.Error())
	}
}

func (m *AssetLoaderSystem) loadFontTable(id akara.EID, path string, data []byte) {
	m.Components.FontTable.Add(id).Data = data

	if cacheErr := m.cache.Insert(path, data, assetCacheEntryWeight); cacheErr != nil {
		m.Error(cacheErr.Error())
	}
}

func (m *AssetLoaderSystem) loadDataDictionary(id akara.EID, path string, data []byte) {
	loaded := d2txt.LoadDataDictionary(data)
	m.Components.DataDictionary.Add(id).DataDictionary = loaded

	if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
		m.Error(cacheErr.Error())
	}
}

func (m *AssetLoaderSystem) loadPalette(id akara.EID, path string, data []byte) error {
	loaded, err := d2dat.Load(data)
	if err == nil {
		m.Components.Palette.Add(id).Palette = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadPaletteTransform(id akara.EID, path string, data []byte) error {
	loaded, err := d2pl2.Load(data)
	if err == nil {
		m.Components.PaletteTransform.Add(id).PL2 = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadCOF(id akara.EID, path string, data []byte) error {
	loaded, err := d2cof.Load(data)
	if err == nil {
		m.Components.Cof.Add(id).COF = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadDC6(id akara.EID, path string, data []byte) error {
	loaded, err := d2dc6.Load(data)
	if err == nil {
		m.Components.Dc6.Add(id).DC6 = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadDCC(id akara.EID, path string, data []byte) error {
	loaded, err := d2dcc.Load(data)
	if err == nil {
		m.Components.Dcc.Add(id).DCC = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadDS1(id akara.EID, path string, data []byte) error {
	loaded, err := d2ds1.LoadDS1(data)
	if err == nil {
		m.Components.Ds1.Add(id).DS1 = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadDT1(id akara.EID, path string, data []byte) error {
	loaded, err := d2dt1.LoadDT1(data)
	if err == nil {
		m.Components.Dt1.Add(id).DT1 = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadWAV(id akara.EID, path string, seeker io.ReadSeeker) {
	component := m.Components.Wav.Add(id)
	component.Data = seeker

	if cacheErr := m.cache.Insert(path, seeker, assetCacheEntryWeight); cacheErr != nil {
		m.Error(cacheErr.Error())
	}
}

func (m *AssetLoaderSystem) loadAnimationData(id akara.EID, path string, data []byte) error {
	loaded, err := d2animdata.Load(data)
	if err == nil {
		m.Components.AnimationData.Add(id).AnimationData = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}
