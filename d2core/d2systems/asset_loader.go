package d2systems

import (
	"errors"
	"io"

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

// NewAssetLoader creates a new asset loader instance
func NewAssetLoader() *AssetLoaderSystem {
	// we are going to check entities that dont yet have loaded asset types
	filesToLoad := akara.NewFilter().
		Require(d2components.FilePath).
		Require(d2components.FileType).
		Require(d2components.FileHandle).
		Forbid(d2components.FileSource).
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

	return &AssetLoaderSystem{
		SubscriberSystem: akara.NewSubscriberSystem(filesToLoad, fileSources),
	}
}

var _ akara.System = &AssetLoaderSystem{}

type AssetLoaderSystem struct {
	*akara.SubscriberSystem
	fileSub          *akara.Subscription
	sourceSub        *akara.Subscription
	filePaths        *d2components.FilePathMap
	fileTypes        *d2components.FileTypeMap
	fileHandles      *d2components.FileHandleMap
	fileSources      *d2components.FileSourceMap
	stringTables     *d2components.StringTableMap
	dataDictionaries *d2components.DataDictionaryMap
	palettes         *d2components.PaletteMap
	transforms       *d2components.PaletteTransformMap
	cof              *d2components.CofMap
	dc6              *d2components.Dc6Map
	dcc              *d2components.DccMap
	ds1              *d2components.Ds1Map
	dt1              *d2components.Dt1Map
	wav              *d2components.WavMap
	animDatas        *d2components.AnimDataMap
}

// Init initializes the system with the given world
func (m *AssetLoaderSystem) Init(world *akara.World) {
	m.World = world

	if world == nil {
		m.SetActive(false)
		return
	}

	for subIdx := range m.Subscriptions {
		m.AddSubscription(m.Subscriptions[subIdx])
	}

	m.fileSub = m.Subscriptions[0]
	m.sourceSub = m.Subscriptions[1]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.filePaths = m.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	m.fileTypes = m.InjectMap(d2components.FileType).(*d2components.FileTypeMap)
	m.fileHandles = m.InjectMap(d2components.FileHandle).(*d2components.FileHandleMap)
	m.fileSources = m.InjectMap(d2components.FileSource).(*d2components.FileSourceMap)
	m.stringTables = m.InjectMap(d2components.StringTable).(*d2components.StringTableMap)
	m.dataDictionaries = m.InjectMap(d2components.DataDictionary).(*d2components.DataDictionaryMap)
	m.palettes = m.InjectMap(d2components.Palette).(*d2components.PaletteMap)
	m.transforms = m.InjectMap(d2components.PaletteTransform).(*d2components.PaletteTransformMap)
	m.cof = m.InjectMap(d2components.Cof).(*d2components.CofMap)
	m.dc6 = m.InjectMap(d2components.Dc6).(*d2components.Dc6Map)
	m.dcc = m.InjectMap(d2components.Dcc).(*d2components.DccMap)
	m.ds1 = m.InjectMap(d2components.Ds1).(*d2components.Ds1Map)
	m.dt1 = m.InjectMap(d2components.Dt1).(*d2components.Dt1Map)
	m.wav = m.InjectMap(d2components.Wav).(*d2components.WavMap)
	m.animDatas = m.InjectMap(d2components.AnimData).(*d2components.AnimDataMap)
}

// Process processes all of the Entities
func (m *AssetLoaderSystem) Process() {
	for _, eid := range m.fileSub.GetEntities() {
		m.loadAsset(eid)
	}
}

func (m *AssetLoaderSystem) loadAsset(id akara.EID) {
	ft, found := m.fileTypes.GetFileType(id)
	if !found {
		return
	}

	fh, found := m.fileHandles.GetFileHandle(id)
	if !found {
		return
	}

	data, buf := make([]byte, 0), make([]byte, 16)

	for {
		numRead, err := fh.Data.Read(buf)
		data = append(data, buf[:numRead]...)

		if numRead < 1 || err != nil {
			break
		}
	}

	err := m.tryToParse(id, data, ft.Type)
	if err != nil {
		ft.Type = d2enum.FileTypeUnknown
	}
}

func (m *AssetLoaderSystem) tryToParse(id akara.EID, data []byte, t d2enum.FileType) error {
	var err error

	switch t {
	case d2enum.FileTypeStringTable:
		m.loadFileTypeStringTable(id, data) // TODO: add error handling for string table load
	case d2enum.FileTypeDataDictionary:
		m.loadFileTypeDataDictionary(id, data) // TODO: add error handling for data dict load
	case d2enum.FileTypePalette:
		err = m.loadFileTypePalette(id, data)
	case d2enum.FileTypePaletteTransform:
		err = m.loadFileTypePaletteTransform(id, data)
	case d2enum.FileTypeCOF:
		err = m.loadFileTypeCOF(id, data)
	case d2enum.FileTypeDC6:
		err = m.loadFileTypeDC6(id, data)
	case d2enum.FileTypeDCC:
		err = m.loadFileTypeDCC(id, data)
	case d2enum.FileTypeDS1:
		err = m.loadFileTypeDS1(id, data)
	case d2enum.FileTypeDT1:
		err = m.loadFileTypeDT1(id, data)
	case d2enum.FileTypeWAV:
		fh, found := m.fileHandles.GetFileHandle(id)
		if !found {
			return errors.New("no file handle for wav file")
		}

		m.loadFileTypeWAV(id, fh.Data)
	case d2enum.FileTypeD2:
		err = m.loadFileTypeD2(id, data)
	}

	return err
}

func (m *AssetLoaderSystem) loadFileTypeStringTable(id akara.EID, data []byte) {
	txt := d2tbl.LoadTextDictionary(data)
	m.stringTables.AddStringTable(id).TextDictionary = &txt
}

func (m *AssetLoaderSystem) loadFileTypeDataDictionary(id akara.EID, data []byte) {
	m.dataDictionaries.AddDataDictionary(id).DataDictionary = d2txt.LoadDataDictionary(data)
}

func (m *AssetLoaderSystem) loadFileTypePalette(id akara.EID, data []byte) error {
	loaded, err := d2dat.Load(data)
	if err == nil {
		m.palettes.AddPalette(id).Palette = loaded
	}

	return err
}

func (m *AssetLoaderSystem) loadFileTypePaletteTransform(id akara.EID, data []byte) error {
	loaded, err := d2pl2.Load(data)
	if err == nil {
		m.transforms.AddPaletteTransform(id).Transform = loaded
	}

	return err
}

func (m *AssetLoaderSystem) loadFileTypeCOF(id akara.EID, data []byte) error {
	loaded, err := d2cof.Load(data)
	if err == nil {
		m.cof.AddCof(id).COF = loaded
	}

	return err
}

func (m *AssetLoaderSystem) loadFileTypeDC6(id akara.EID, data []byte) error {
	loaded, err := d2dc6.Load(data)
	if err == nil {
		m.dc6.AddDc6(id).DC6 = loaded
	}

	return err
}

func (m *AssetLoaderSystem) loadFileTypeDCC(id akara.EID, data []byte) error {
	loaded, err := d2dcc.Load(data)
	if err == nil {
		m.dcc.AddDcc(id).DCC = loaded
	}

	return err
}

func (m *AssetLoaderSystem) loadFileTypeDS1(id akara.EID, data []byte) error {
	loaded, err := d2ds1.LoadDS1(data)
	if err == nil {
		m.ds1.AddDs1(id).DS1 = loaded
	}

	return err
}

func (m *AssetLoaderSystem) loadFileTypeDT1(id akara.EID, data []byte) error {
	loaded, err := d2dt1.LoadDT1(data)
	if err == nil {
		m.dt1.AddDt1(id).DT1 = loaded
	}

	return err
}

func (m *AssetLoaderSystem) loadFileTypeWAV(id akara.EID, seeker io.ReadSeeker) {
	component := m.wav.AddWav(id)
	component.Data = seeker
}

func (m *AssetLoaderSystem) loadFileTypeD2(id akara.EID, data []byte) error {
	loaded, err := d2animdata.Load(data)
	if err == nil {
		m.animDatas.AddAnimData(id).AnimationData = loaded
	}

	return err
}
