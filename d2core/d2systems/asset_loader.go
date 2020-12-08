package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"io"

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

var _ akara.System = &AssetLoaderSystem{}

// AssetLoaderSystem is responsible for parsing data from file handles into various structs, like COF or DC6
type AssetLoaderSystem struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	fileSub   *akara.Subscription
	sourceSub *akara.Subscription
	cache     *d2cache.Cache
	localeString string // related to file "/data/local/use"
	d2components.FileFactory
	d2components.FileTypeFactory
	d2components.FileHandleFactory
	d2components.FileSourceFactory
	d2components.StringTableFactory
	d2components.FontTableFactory
	d2components.DataDictionaryFactory
	d2components.PaletteFactory
	d2components.PaletteTransformFactory
	d2components.CofFactory
	d2components.Dc6Factory
	d2components.DccFactory
	d2components.Ds1Factory
	d2components.Dt1Factory
	d2components.WavFactory
	d2components.AnimationDataFactory
	d2components.LocaleFactory
	d2components.BitmapFontFactory
	d2components.FileLoadedFactory
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

	m.fileSub = m.AddSubscription(filesToLoad)
	m.sourceSub = m.AddSubscription(fileSources)
}

func (m *AssetLoaderSystem) setupFactories() {
	m.Debug("setting up component factories")

	m.InjectComponent(&d2components.File{}, &m.File)
	m.InjectComponent(&d2components.FileType{}, &m.FileType)
	m.InjectComponent(&d2components.FileHandle{}, &m.FileHandle)
	m.InjectComponent(&d2components.FileSource{}, &m.FileSource)
	m.InjectComponent(&d2components.StringTable{}, &m.StringTable)
	m.InjectComponent(&d2components.FontTable{}, &m.FontTable)
	m.InjectComponent(&d2components.DataDictionary{}, &m.DataDictionary)
	m.InjectComponent(&d2components.Palette{}, &m.Palette)
	m.InjectComponent(&d2components.PaletteTransform{}, &m.PaletteTransform)
	m.InjectComponent(&d2components.Cof{}, &m.Cof)
	m.InjectComponent(&d2components.Dc6{}, &m.Dc6)
	m.InjectComponent(&d2components.Dcc{}, &m.Dcc)
	m.InjectComponent(&d2components.Ds1{}, &m.Ds1)
	m.InjectComponent(&d2components.Dt1{}, &m.Dt1)
	m.InjectComponent(&d2components.Wav{}, &m.Wav)
	m.InjectComponent(&d2components.AnimationData{}, &m.AnimationData)
	m.InjectComponent(&d2components.Locale{}, &m.Locale)
	m.InjectComponent(&d2components.BitmapFont{}, &m.BitmapFont)
	m.InjectComponent(&d2components.FileLoaded{}, &m.FileLoaded)
}

// Update processes all of the Entities in the subscription of file entities that need to be processed
func (m *AssetLoaderSystem) Update() {
	for _, eid := range m.fileSub.GetEntities() {
		m.loadAsset(eid)
	}
}

func (m *AssetLoaderSystem) loadAsset(id akara.EID) {
	// make sure everything is kosher
	fp, found := m.GetFile(id)
	if !found {
		m.Errorf("filepath component not found for entity %d", id)
		return
	}

	ft, found := m.GetFileType(id)
	if !found {
		m.Errorf("filetype component not found for entity %d", id)
		return
	}

	fh, found := m.GetFileHandle(id)
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
		m.AddStringTable(id).TextDictionary = entry.(*d2tbl.TextDictionary)
	case d2enum.FileTypeFontTable:
		m.AddFontTable(id).Data = entry.([]byte)
	case d2enum.FileTypeDataDictionary:
		m.AddDataDictionary(id).DataDictionary = entry.(*d2txt.DataDictionary)
	case d2enum.FileTypePalette:
		m.AddPalette(id).Palette = entry.(d2interface.Palette)
	case d2enum.FileTypePaletteTransform:
		m.AddPaletteTransform(id).PL2 = entry.(*d2pl2.PL2)
	case d2enum.FileTypeCOF:
		m.AddCof(id).COF = entry.(*d2cof.COF)
	case d2enum.FileTypeDC6:
		m.AddDc6(id).DC6 = entry.(*d2dc6.DC6)
	case d2enum.FileTypeDCC:
		m.AddDcc(id).DCC = entry.(*d2dcc.DCC)
	case d2enum.FileTypeDS1:
		m.AddDs1(id).DS1 = entry.(*d2ds1.DS1)
	case d2enum.FileTypeDT1:
		m.AddDt1(id).DT1 = entry.(*d2dt1.DT1)
	case d2enum.FileTypeWAV:
		m.AddWav(id).Data = entry.(d2interface.DataStream)
	case d2enum.FileTypeD2:
		m.AddAnimationData(id).AnimationData = entry.(*d2animdata.AnimationData)
	}

	m.AddFileLoaded(id)

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

		fh, found := m.GetFileHandle(id)
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

	m.AddFileLoaded(id)
}

func (m *AssetLoaderSystem) loadLocale(id akara.EID, data []byte) {
	locale := m.AddLocale(id)

	locale.Code = data[0]
	locale.String = d2resource.GetLanguageLiteral(locale.Code)

	m.localeString = locale.String
}

func (m *AssetLoaderSystem) loadStringTable(id akara.EID, path string, data []byte) {
	txt := d2tbl.LoadTextDictionary(data)
	loaded := &txt
	m.AddStringTable(id).TextDictionary = loaded

	if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
		m.Error(cacheErr.Error())
	}
}

func (m *AssetLoaderSystem) loadFontTable(id akara.EID, path string, data []byte) {
	m.AddFontTable(id).Data = data

	if cacheErr := m.cache.Insert(path, data, assetCacheEntryWeight); cacheErr != nil {
		m.Error(cacheErr.Error())
	}
}

func (m *AssetLoaderSystem) loadDataDictionary(id akara.EID, path string, data []byte) {
	loaded := d2txt.LoadDataDictionary(data)
	m.AddDataDictionary(id).DataDictionary = loaded

	if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
		m.Error(cacheErr.Error())
	}
}

func (m *AssetLoaderSystem) loadPalette(id akara.EID, path string, data []byte) error {
	loaded, err := d2dat.Load(data)
	if err == nil {
		m.AddPalette(id).Palette = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadPaletteTransform(id akara.EID, path string, data []byte) error {
	loaded, err := d2pl2.Load(data)
	if err == nil {
		m.AddPaletteTransform(id).PL2 = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadCOF(id akara.EID, path string, data []byte) error {
	loaded, err := d2cof.Load(data)
	if err == nil {
		m.AddCof(id).COF = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadDC6(id akara.EID, path string, data []byte) error {
	loaded, err := d2dc6.Load(data)
	if err == nil {
		m.AddDc6(id).DC6 = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadDCC(id akara.EID, path string, data []byte) error {
	loaded, err := d2dcc.Load(data)
	if err == nil {
		m.AddDcc(id).DCC = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadDS1(id akara.EID, path string, data []byte) error {
	loaded, err := d2ds1.LoadDS1(data)
	if err == nil {
		m.AddDs1(id).DS1 = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadDT1(id akara.EID, path string, data []byte) error {
	loaded, err := d2dt1.LoadDT1(data)
	if err == nil {
		m.AddDt1(id).DT1 = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}

func (m *AssetLoaderSystem) loadWAV(id akara.EID, path string, seeker io.ReadSeeker) {
	component := m.AddWav(id)
	component.Data = seeker

	if cacheErr := m.cache.Insert(path, seeker, assetCacheEntryWeight); cacheErr != nil {
		m.Error(cacheErr.Error())
	}
}

func (m *AssetLoaderSystem) loadAnimationData(id akara.EID, path string, data []byte) error {
	loaded, err := d2animdata.Load(data)
	if err == nil {
		m.AddAnimationData(id).AnimationData = loaded

		if cacheErr := m.cache.Insert(path, loaded, assetCacheEntryWeight); cacheErr != nil {
			m.Error(cacheErr.Error())
		}
	}

	return err
}
