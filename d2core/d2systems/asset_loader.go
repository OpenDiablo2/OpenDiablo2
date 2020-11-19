package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"io"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2cache"

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
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	assetCacheBudget      = 1024
	assetCacheEntryWeight = 1 // may want to make different weights for different asset types
)

const (
	LogPrefixAssetLoader = "Asset Loader System"
)

// NewAssetLoader creates a new asset loader instance
func NewAssetLoader() *AssetLoaderSystem {
	// we are going to check entities that dont yet have loaded asset types
	filesToLoad := akara.NewFilter().
		Require(d2components.FilePath). // we want to process entities with these file components
		Require(d2components.FileType).
		Require(d2components.FileHandle).
		Forbid(d2components.FileSource). // but we forbid files that are already loaded
		Forbid(d2components.GameConfig).
		Forbid(d2components.StringTable).
		Forbid(d2components.DataDictionary).
		Forbid(d2components.Palette).
		Forbid(d2components.PaletteTransform).
		Forbid(d2components.Cof).
		Forbid(d2components.Dc6).
		Forbid(d2components.Dcc).
		Forbid(d2components.Ds1).
		Forbid(d2components.Dt1).
		Forbid(d2components.Wav).
		Forbid(d2components.AnimData).
		Build()

	fileSources := akara.NewFilter().
		Require(d2components.FileSource).
		Build()

	assetLoader := &AssetLoaderSystem{
		SubscriberSystem: akara.NewSubscriberSystem(filesToLoad, fileSources),
		cache:            d2cache.CreateCache(assetCacheBudget).(*d2cache.Cache),
		Logger:           d2util.NewLogger(),
	}

	assetLoader.SetPrefix(LogPrefixAssetLoader)

	return assetLoader
}

var _ akara.System = &AssetLoaderSystem{}

// AssetLoaderSystem is responsible for parsing file handle data into various structs, like COF or DC6
type AssetLoaderSystem struct {
	*akara.SubscriberSystem
	*d2util.Logger
	fileSub          *akara.Subscription
	sourceSub        *akara.Subscription
	cache            *d2cache.Cache
	*d2components.FilePathMap
	*d2components.FileTypeMap
	*d2components.FileHandleMap
	*d2components.FileSourceMap
	*d2components.StringTableMap
	*d2components.FontTableMap
	*d2components.DataDictionaryMap
	*d2components.PaletteMap
	*d2components.PaletteTransformMap
	*d2components.CofMap
	*d2components.Dc6Map
	*d2components.DccMap
	*d2components.Ds1Map
	*d2components.Dt1Map
	*d2components.WavMap
	*d2components.AnimDataMap
}

// Init injects component maps related to various asset types
func (m *AssetLoaderSystem) Init(world *akara.World) {
	m.World = world

	if world == nil {
		m.SetActive(false)
		return
	}

	m.Info("initializing ...")

	for subIdx := range m.Subscriptions {
		m.Subscriptions[subIdx] = m.AddSubscription(m.Subscriptions[subIdx].Filter)
	}

	m.fileSub = m.Subscriptions[0]
	m.sourceSub = m.Subscriptions[1]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.FilePathMap = m.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	m.FileTypeMap = m.InjectMap(d2components.FileType).(*d2components.FileTypeMap)
	m.FileHandleMap = m.InjectMap(d2components.FileHandle).(*d2components.FileHandleMap)
	m.FileSourceMap = m.InjectMap(d2components.FileSource).(*d2components.FileSourceMap)
	m.StringTableMap = m.InjectMap(d2components.StringTable).(*d2components.StringTableMap)
	m.DataDictionaryMap = m.InjectMap(d2components.DataDictionary).(*d2components.DataDictionaryMap)
	m.PaletteMap = m.InjectMap(d2components.Palette).(*d2components.PaletteMap)
	m.PaletteTransformMap = m.InjectMap(d2components.PaletteTransform).(*d2components.PaletteTransformMap)
	m.FontTableMap = m.InjectMap(d2components.FontTable).(*d2components.FontTableMap)
	m.CofMap = m.InjectMap(d2components.Cof).(*d2components.CofMap)
	m.Dc6Map = m.InjectMap(d2components.Dc6).(*d2components.Dc6Map)
	m.DccMap = m.InjectMap(d2components.Dcc).(*d2components.DccMap)
	m.Ds1Map = m.InjectMap(d2components.Ds1).(*d2components.Ds1Map)
	m.Dt1Map = m.InjectMap(d2components.Dt1).(*d2components.Dt1Map)
	m.WavMap = m.InjectMap(d2components.Wav).(*d2components.WavMap)
	m.AnimDataMap = m.InjectMap(d2components.AnimData).(*d2components.AnimDataMap)
}

// Update processes all of the Entities in the subscription of file entities that need to be processed
func (m *AssetLoaderSystem) Update() {
	for _, eid := range m.fileSub.GetEntities() {
		m.loadAsset(eid)
	}
}

func (m *AssetLoaderSystem) loadAsset(id akara.EID) {
	// make sure everything is kosher
	fp, found := m.GetFilePath(id)
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
		m.AddPaletteTransform(id).Transform = entry.(*d2pl2.PL2)
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
		m.AddAnimData(id).AnimationData = entry.(*d2animdata.AnimationData)
	}

	return found
}

func (m *AssetLoaderSystem) parseAndCache(id akara.EID, path string, t d2enum.FileType, data []byte) {
	go func() {
		switch t {
		case d2enum.FileTypeStringTable:
			m.Infof("Loading string table: %s", path)
			m.loadStringTable(id, path, data) // TODO: add error handling for string table load
		case d2enum.FileTypeFontTable:
			m.Infof("Loading font table: %s", path)
			m.loadFontTable(id, path, data) // TODO: add error handling for string table load
		case d2enum.FileTypeDataDictionary:
			m.Infof("Loading data dictionary: %s", path)
			m.loadDataDictionary(id, path, data) // TODO: add error handling for data dict load
		case d2enum.FileTypePalette:
			m.Infof("Loading palette: %s", path)
			m.loadPalette(id, path, data)
		case d2enum.FileTypePaletteTransform:
			m.Infof("Loading palette transform: %s", path)
			m.loadPaletteTransform(id, path, data)
		case d2enum.FileTypeCOF:
			m.Infof("Loading COF: %s", path)
			m.loadCOF(id, path, data)
		case d2enum.FileTypeDC6:
			m.Infof("Loading DC6: %s", path)
			m.loadDC6(id, path, data)
		case d2enum.FileTypeDCC:
			m.Infof("Loading DCC: %s", path)
			m.loadDCC(id, path, data)
		case d2enum.FileTypeDS1:
			m.Infof("Loading DS1: %s", path)
			m.loadDS1(id, path, data)
		case d2enum.FileTypeDT1:
			m.Infof("Loading DT1: %s", path)
			m.loadDT1(id, path, data)
		case d2enum.FileTypeWAV:
			m.Infof("Loading WAV: %s", path)
			fh, found := m.GetFileHandle(id)
			if !found {
				return
			}

			m.loadWAV(id, path, fh.Data)
		case d2enum.FileTypeD2:
			m.Infof("Loading animation data: %s", path)
			m.loadAnimData(id, path, data)
		}
	}()
}

func (m *AssetLoaderSystem) loadStringTable(id akara.EID, path string, data []byte) {
	txt := d2tbl.LoadTextDictionary(data)
	loaded := &txt
	m.AddStringTable(id).TextDictionary = loaded
	m.cache.Insert(path, loaded, assetCacheEntryWeight)
}

func (m *AssetLoaderSystem) loadFontTable(id akara.EID, path string, data []byte) {
	m.AddFontTable(id).Data = data
	m.cache.Insert(path, data, assetCacheEntryWeight)
}

func (m *AssetLoaderSystem) loadDataDictionary(id akara.EID, path string, data []byte) {
	loaded := d2txt.LoadDataDictionary(data)
	m.AddDataDictionary(id).DataDictionary = loaded
	m.cache.Insert(path, loaded, assetCacheEntryWeight)
}

func (m *AssetLoaderSystem) loadPalette(id akara.EID, path string, data []byte) error {
	loaded, err := d2dat.Load(data)
	if err == nil {
		m.AddPalette(id).Palette = loaded
		m.cache.Insert(path, loaded, assetCacheEntryWeight)
	}

	return err
}

func (m *AssetLoaderSystem) loadPaletteTransform(id akara.EID, path string, data []byte) error {
	loaded, err := d2pl2.Load(data)
	if err == nil {
		m.AddPaletteTransform(id).Transform = loaded
		m.cache.Insert(path, loaded, assetCacheEntryWeight)
	}

	return err
}

func (m *AssetLoaderSystem) loadCOF(id akara.EID, path string, data []byte) error {
	loaded, err := d2cof.Load(data)
	if err == nil {
		m.AddCof(id).COF = loaded
		m.cache.Insert(path, loaded, assetCacheEntryWeight)
	}

	return err
}

func (m *AssetLoaderSystem) loadDC6(id akara.EID, path string, data []byte) error {
	loaded, err := d2dc6.Load(data)
	if err == nil {
		m.AddDc6(id).DC6 = loaded
		m.cache.Insert(path, loaded, assetCacheEntryWeight)
	}

	return err
}

func (m *AssetLoaderSystem) loadDCC(id akara.EID, path string, data []byte) error {
	loaded, err := d2dcc.Load(data)
	if err == nil {
		m.AddDcc(id).DCC = loaded
		m.cache.Insert(path, loaded, assetCacheEntryWeight)
	}

	return err
}

func (m *AssetLoaderSystem) loadDS1(id akara.EID, path string, data []byte) error {
	loaded, err := d2ds1.LoadDS1(data)
	if err == nil {
		m.AddDs1(id).DS1 = loaded
		m.cache.Insert(path, loaded, assetCacheEntryWeight)
	}

	return err
}

func (m *AssetLoaderSystem) loadDT1(id akara.EID, path string, data []byte) error {
	loaded, err := d2dt1.LoadDT1(data)
	if err == nil {
		m.AddDt1(id).DT1 = loaded
		m.cache.Insert(path, loaded, assetCacheEntryWeight)
	}

	return err
}

func (m *AssetLoaderSystem) loadWAV(id akara.EID, path string, seeker io.ReadSeeker) {
	component := m.AddWav(id)
	component.Data = seeker
	m.cache.Insert(path, seeker, assetCacheEntryWeight)
}

func (m *AssetLoaderSystem) loadAnimData(id akara.EID, path string, data []byte) error {
	loaded, err := d2animdata.Load(data)
	if err == nil {
		m.AddAnimData(id).AnimationData = loaded
		m.cache.Insert(path, loaded, assetCacheEntryWeight)
	}

	return err
}
