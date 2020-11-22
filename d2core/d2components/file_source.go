//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// AbstractSource is the abstract representation of what a file source is
type AbstractSource interface {
	Path() string // the path of the source itself
	Open(path *FilePathComponent) (d2interface.DataStream, error)
}

// static check that FileSourceComponent implements Component
var _ akara.Component = &FileSourceComponent{}

// static check that FileSourceMap implements ComponentMap
var _ akara.ComponentMap = &FileSourceMap{}

// FileSourceComponent contains an embedded file source interface, something that can open files
type FileSourceComponent struct {
	*akara.BaseComponent
	AbstractSource
}

// FileSourceMap is a map of entity ID's to FileSource
type FileSourceMap struct {
	*akara.BaseComponentMap
}

// AddFileSource adds a new FileSourceComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *FileSourceComponent instead of an akara.Component
func (cm *FileSourceMap) AddFileSource(id akara.EID) *FileSourceComponent {
	return cm.Add(id).(*FileSourceComponent)
}

// GetFileSource returns the FileSourceComponent associated with the given entity id
func (cm *FileSourceMap) GetFileSource(id akara.EID) (*FileSourceComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*FileSourceComponent), found
}

// FileSource is a convenient reference to be used as a component identifier
var FileSource = newFileSource() // nolint:gochecknoglobals // global by design

func newFileSource() akara.Component {
	return &FileSourceComponent{
		BaseComponent: akara.NewBaseComponent(FileSourceCID, newFileSource, newFileSourceMap),
	}
}

func newFileSourceMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(FileSourceCID, newFileSource, newFileSourceMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &FileSourceMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
