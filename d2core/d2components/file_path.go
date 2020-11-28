//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that FilePath implements Component
var _ akara.Component = &FilePath{}

// FilePath represents a file path for a file
type FilePath struct {
	Path string
}

// New returns a FilePath component. By default, it contains an empty string.
func (*FilePath) New() akara.Component {
	return &FilePath{}
}

// FilePathFactory is a wrapper for the generic component factory that returns FilePath component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a FilePath.
type FilePathFactory struct {
	FilePath *akara.ComponentFactory
}

// AddFilePath adds a FilePath component to the given entity and returns it
func (m *FilePathFactory) AddFilePath(id akara.EID) *FilePath {
	return m.FilePath.Add(id).(*FilePath)
}

// GetFilePath returns the FilePath component for the given entity, and a bool for whether or not it exists
func (m *FilePathFactory) GetFilePath(id akara.EID) (*FilePath, bool) {
	component, found := m.FilePath.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FilePath), found
}
