package d2systems

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

func NewFileHandleResolver() *FileHandleResolutionSystem {
	// this filter is for entities that have a file path and file type but no file handle.
	filesToSource := akara.NewFilter().
		Require(d2components.FilePath).
		Require(d2components.FileType).
		Forbid(d2components.FileHandle).
		Forbid(d2components.FileSource).
		Build()

	sourcesToUse := akara.NewFilter().
		RequireOne(d2components.FileSource).
		Build()

	return &FileHandleResolutionSystem{
		SubscriberSystem: akara.NewSubscriberSystem(filesToSource, sourcesToUse),
	}
}

type FileHandleResolutionSystem struct {
	*akara.SubscriberSystem
	filesToLoad  *akara.Subscription
	sourcesToUse *akara.Subscription
	filePaths    *d2components.FilePathMap
	fileTypes    *d2components.FileTypeMap
	fileSources  *d2components.FileSourceMap
	fileHandles  *d2components.FileHandleMap
}

// Init initializes the system with the given world
func (m *FileHandleResolutionSystem) Init(world *akara.World) {
	m.World = world

	for subIdx := range m.Subscriptions {
		m.AddSubscription(m.Subscriptions[subIdx])
	}

	if world == nil {
		m.SetActive(false)
		return
	}

	m.filesToLoad = m.Subscriptions[0]
	m.sourcesToUse = m.Subscriptions[1]

	testBS := akara.NewBitSet(int(d2components.FileSourceCID), 1)
	truth := m.sourcesToUse.Filter.Allow(testBS)
	_ = truth

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.filePaths = m.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	m.fileTypes = m.InjectMap(d2components.FileType).(*d2components.FileTypeMap)
	m.fileHandles = m.InjectMap(d2components.FileHandle).(*d2components.FileHandleMap)
	m.fileSources = m.InjectMap(d2components.FileSource).(*d2components.FileSourceMap)
}

// Process processes all of the Entities
func (m *FileHandleResolutionSystem) Process() {
	filesToLoad := m.filesToLoad.GetEntities()
	sourcesToUse := m.sourcesToUse.GetEntities()

	for _, fileID := range filesToLoad {
		for _, sourceID := range sourcesToUse {
			if m.loadFileWithSource(fileID, sourceID) {
				break
			}
		}
	}
}

// try to load a file with a source, returns true if loaded
func (m *FileHandleResolutionSystem) loadFileWithSource(fileID, sourceID akara.EID) bool {
	fp, found := m.filePaths.GetFilePath(fileID)
	if !found {
		return false
	}

	source, found := m.fileSources.GetFileSource(sourceID)
	if !found {
		return false
	}

	data, err := source.Open(fp)
	if err != nil {
		return false
	}

	dataComponent := m.fileHandles.AddFileHandle(fileID)
	dataComponent.Data = data

	return true
}
