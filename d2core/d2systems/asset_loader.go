package d2systems

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gravestench/ecs"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

func NewAssetLoader() *AssetLoaderSystem {
	filesToLoad := ecs.NewFilter()

	// subscribe to entities with a file path+type+handle, ready to be loaded
	filesToLoad.Require(d2components.FilePath, d2components.FileType, d2components.FileHandle)

	// exclude entities that have already been loaded
	filesToLoad.Forbid(d2components.GameConfig).
		//Forbid(d2components.AssetStringTableCID).
		//Forbid(d2components.AssetDataDictionaryCID).
		//Forbid(d2components.AssetPaletteCID).
		//Forbid(d2components.AssetPaletteTransformCID).
		//Forbid(d2components.AssetCofCID).
		//Forbid(d2components.AssetDc6CID).
		//Forbid(d2components.AssetDccCID).
		//Forbid(d2components.AssetDs1CID).
		//Forbid(d2components.AssetDt1CID).
		//Forbid(d2components.AssetWavCID).
		//Forbid(d2components.AssetD2CID).
		Build()

	// subscribe to entities that have a source type and a source component
	fileSources := ecs.NewFilter().
		Require(d2components.FileSource).
		Build()

	return &AssetLoaderSystem{
		SubscriberSystem: ecs.NewSubscriberSystem(filesToLoad.Build(), fileSources),
	}
}

var _ ecs.System = &AssetLoaderSystem{}

type AssetLoaderSystem struct {
	*ecs.SubscriberSystem
	fileSub     *ecs.Subscription
	sourceSub   *ecs.Subscription
	filePaths   *d2components.FilePathMap
	fileTypes   *d2components.FileTypeMap
	fileHandles *d2components.FileHandleMap
	fileSources *d2components.FileSourceMap
}

// Init initializes the system with the given world
func (m *AssetLoaderSystem) Init(world *ecs.World) {
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
}

// Process processes all of the Entities
func (m *AssetLoaderSystem) Process() {
	for _, eid := range m.fileSub.GetEntities() {
		m.ProcessEntity(eid)
	}
}

// ProcessEntity updates an individual entity in the system
func (m *AssetLoaderSystem) ProcessEntity(id ecs.EID) {
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

	var err error

	switch ft.Type {
	case d2enum.FileTypeJSON:
		err = m.loadFileTypeJSON(id, data)
		//case d2enum.FileTypeStringTable:
		//	err = m.loadFileTypeStringTable(id, data)
		//case d2enum.FileTypeDataDictionary:
		//	err = m.loadFileTypeDataDictionary(id, data)
		//case d2enum.FileTypePalette:
		//	err = m.loadFileTypePalette(id, data)
		//case d2enum.FileTypePaletteTransform:
		//	err = m.loadFileTypePaletteTransform(id, data)
		//case d2enum.FileTypeCOF:
		//	err = m.loadFileTypeCOF(id, data)
		//case d2enum.FileTypeDC6:
		//	err = m.loadFileTypeDC6(id, data)
		//case d2enum.FileTypeDCC:
		//	err = m.loadFileTypeDCC(id, data)
		//case d2enum.FileTypeDS1:
		//	err = m.loadFileTypeDS1(id, data)
		//case d2enum.FileTypeDT1:
		//	err = m.loadFileTypeDT1(id, data)
		//case d2enum.FileTypeWAV:
		//	err = m.loadFileTypeWAV(id, data)
		//case d2enum.FileTypeD2:
		//	err = m.loadFileTypeD2(id, data)
	}

	if err != nil {
		ft.Type = d2enum.FileTypeUnknown
	}
}

func (m *AssetLoaderSystem) loadFileTypeJSON(id ecs.EID, data []byte) error {
	_, found := m.filePaths.GetFilePath(id)
	if !found {
		return errors.New("file path component for entity not found")
	}

	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	mpq := result["MpqLoadOrder"].([]interface{})
	fmt.Println("Address :", mpq)

	return nil
}

//func (m *AssetLoaderSystem) loadFileTypeStringTable(id ecs.EID, data []byte) error {
//
//}
//
//func (m *AssetLoaderSystem) loadFileTypeDataDictionary(id ecs.EID, data []byte) error {
//
//}
//
//func (m *AssetLoaderSystem) loadFileTypePalette(id ecs.EID, data []byte) error {
//
//}
//
//func (m *AssetLoaderSystem) loadFileTypePaletteTransform(id ecs.EID, data []byte) error {
//
//}
//
//func (m *AssetLoaderSystem) loadFileTypeCOF(id ecs.EID, data []byte) error {
//
//}
//
//func (m *AssetLoaderSystem) loadFileTypeDC6(id ecs.EID, data []byte) error {
//
//}
//
//func (m *AssetLoaderSystem) loadFileTypeDCC(id ecs.EID, data []byte) error {
//
//}
//
//func (m *AssetLoaderSystem) loadFileTypeDS1(id ecs.EID, data []byte) error {
//
//}
//
//func (m *AssetLoaderSystem) loadFileTypeDT1(id ecs.EID, data []byte) error {
//
//}
//
//func (m *AssetLoaderSystem) loadFileTypeWAV(id ecs.EID, data []byte) error {
//
//}
//
//func (m *AssetLoaderSystem) loadFileTypeD2(id ecs.EID, data []byte) error {
//
//}
