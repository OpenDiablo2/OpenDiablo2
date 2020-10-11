package d2systems

import (
	"github.com/gravestench/ecs"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

func NewFileHandleResolver() *FileHandleResolutionSystem {
	// this filter is for entities that have a file path and file type but no file handle.
	filesToSource := ecs.NewFilter().
		Require(d2components.FilePath, d2components.FileType).
		Forbid(d2components.FileHandle, d2components.FileSource).
		Build()

	return &FileHandleResolutionSystem{
		SubscriberSystem: ecs.NewSubscriberSystem(filesToSource),
	}
}

type FileHandleResolutionSystem struct {
	*ecs.SubscriberSystem
	fileSub     *ecs.Subscription
	filePaths   *d2components.FilePathMap
	fileTypes   *d2components.FileTypeMap
	fileSources *d2components.FileSourceMap
	fileHandles *d2components.FileHandleMap
}

// Init initializes the system with the given world
func (m *FileHandleResolutionSystem) Init(world *ecs.World) {
	m.World = world

	for subIdx := range m.Subscriptions {
		m.AddSubscription(m.Subscriptions[subIdx])
	}

	if world == nil {
		m.SetActive(false)
		return
	}

	m.fileSub = m.Subscriptions[0]

	// try to inject the components we require, then cast the returned
	// abstract ComponentMap back to the concrete implementation
	m.filePaths = m.InjectMap(d2components.FilePath).(*d2components.FilePathMap)
	m.fileTypes = m.InjectMap(d2components.FileType).(*d2components.FileTypeMap)
	m.fileHandles = m.InjectMap(d2components.FileHandle).(*d2components.FileHandleMap)
	m.fileSources = m.InjectMap(d2components.FileSource).(*d2components.FileSourceMap)
}

// Process processes all of the Entities
func (m *FileHandleResolutionSystem) Process() {
	for _, EID := range m.fileSub.GetEntities() {
		m.ProcessEntity(EID)
	}
}

// ProcessEntity updates an individual entity in the system
func (m *FileHandleResolutionSystem) ProcessEntity(id ecs.EID) {
	fp, found := m.filePaths.GetFilePath(id)
	if !found {
		return
	}

	for _, source := range m.fileSources.GetFileSources() {
		data, err := source.Open(fp)
		if err != nil {
			continue
		}

		dataComponent := m.fileHandles.AddFileHandle(id)
		dataComponent.Data = data

		break
	}
}
