//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// static check that FileType implements Component
var _ akara.Component = &FileType{}

// FileType is used to flag file entities with a file type
type FileType struct {
	Type d2enum.FileType
}

// New returns a FileType component. By default, it contains a nil instance.
func (*FileType) New() akara.Component {
	return &FileType{}
}

// FileTypeFactory is a wrapper for the generic component factory that returns FileType component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a FileType.
type FileTypeFactory struct {
	FileType *akara.ComponentFactory
}

// AddFileType adds a FileType component to the given entity and returns it
func (m *FileTypeFactory) AddFileType(id akara.EID) *FileType {
	return m.FileType.Add(id).(*FileType)
}

// GetFileType returns the FileType component for the given entity, and a bool for whether or not it exists
func (m *FileTypeFactory) GetFileType(id akara.EID) (*FileType, bool) {
	component, found := m.FileType.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FileType), found
}
