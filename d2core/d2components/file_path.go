//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that FilePathComponent implements Component
var _ akara.Component = &FilePathComponent{}

// static check that FilePathMap implements ComponentMap
var _ akara.ComponentMap = &FilePathMap{}

// FilePathComponent represents a file path for a file
type FilePathComponent struct {
	*akara.BaseComponent
	Path string
}

// FilePathMap is a map of entity ID's to FilePath
type FilePathMap struct {
	*akara.BaseComponentMap
}

// AddFilePath adds a new FilePathComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *FilePathComponent instead of an akara.Component
func (cm *FilePathMap) AddFilePath(id akara.EID) *FilePathComponent {
	return cm.Add(id).(*FilePathComponent)
}

// GetFilePath returns the FilePathComponent associated with the given entity id
func (cm *FilePathMap) GetFilePath(id akara.EID) (*FilePathComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*FilePathComponent), found
}

// FilePath is a convenient reference to be used as a component identifier
var FilePath = newFilePath() // nolint:gochecknoglobals // global by design

func newFilePath() akara.Component {
	return &FilePathComponent{
		BaseComponent: akara.NewBaseComponent(FilePathCID, newFilePath, newFilePathMap),
	}
}

func newFilePathMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(FilePathCID, newFilePath, newFilePathMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &FilePathMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
