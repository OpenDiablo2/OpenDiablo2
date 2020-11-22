//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// static check that FileTypeComponent implements Component
var _ akara.Component = &FileTypeComponent{}

// static check that FileTypeMap implements ComponentMap
var _ akara.ComponentMap = &FileTypeMap{}

// FileTypeComponent is used to flag file entities with a file type
type FileTypeComponent struct {
	*akara.BaseComponent
	Type d2enum.FileType
}

// FileTypeMap is a map of entity ID's to FileType
type FileTypeMap struct {
	*akara.BaseComponentMap
}

// AddFileType adds a new FileTypeComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *FileTypeComponent instead of an akara.Component
func (cm *FileTypeMap) AddFileType(id akara.EID) *FileTypeComponent {
	return cm.Add(id).(*FileTypeComponent)
}

// GetFileType returns the FileTypeComponent associated with the given entity id
func (cm *FileTypeMap) GetFileType(id akara.EID) (*FileTypeComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*FileTypeComponent), found
}

// FileType is a convenient reference to be used as a component identifier
var FileType = newFileType() // nolint:gochecknoglobals // global by design

func newFileType() akara.Component {
	return &FileTypeComponent{
		BaseComponent: akara.NewBaseComponent(FileTypeCID, newFileType, newFileTypeMap),
	}
}

func newFileTypeMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(FileTypeCID, newFileType, newFileTypeMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &FileTypeMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
