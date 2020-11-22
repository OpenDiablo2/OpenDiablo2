//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that FileHandleComponent implements Component
var _ akara.Component = &FileHandleComponent{}

// static check that FileHandleMap implements ComponentMap
var _ akara.ComponentMap = &FileHandleMap{}

// FileHandleComponent is a component that contains a data stream for file data
type FileHandleComponent struct {
	*akara.BaseComponent
	Data d2interface.DataStream
}

// FileHandleMap is a map of entity ID's to FileHandle
type FileHandleMap struct {
	*akara.BaseComponentMap
}

// AddFileHandle adds a new FileHandleComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *FileHandleComponent instead of an akara.Component
func (cm *FileHandleMap) AddFileHandle(id akara.EID) *FileHandleComponent {
	return cm.Add(id).(*FileHandleComponent)
}

// GetFileHandle returns the FileHandleComponent associated with the given entity id
func (cm *FileHandleMap) GetFileHandle(id akara.EID) (*FileHandleComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*FileHandleComponent), found
}

// FileHandle is a convenient reference to be used as a component identifier
var FileHandle = newFileHandle() // nolint:gochecknoglobals // global by design

func newFileHandle() akara.Component {
	return &FileHandleComponent{
		BaseComponent: akara.NewBaseComponent(FileHandleCID, newFileHandle, newFileHandleMap),
	}
}

func newFileHandleMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(FileHandleCID, newFileHandle, newFileHandleMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &FileHandleMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
