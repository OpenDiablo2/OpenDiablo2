//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that FileLoaded implements Component
var _ akara.Component = &FileLoaded{}

// FileLoaded is used to flag file entities as having been loaded. it is an empty struct.
type FileLoaded struct {}

// New returns a FileLoaded component. By default, it contains an empty string.
func (*FileLoaded) New() akara.Component {
	return &FileLoaded{}
}

// FileLoadedFactory is a wrapper for the generic component factory that returns FileLoaded component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a FileLoaded.
type FileLoadedFactory struct {
	FileLoaded *akara.ComponentFactory
}

// AddFileLoaded adds a FileLoaded component to the given entity and returns it
func (m *FileLoadedFactory) AddFileLoaded(id akara.EID) *FileLoaded {
	return m.FileLoaded.Add(id).(*FileLoaded)
}

// GetFileLoaded returns the FileLoaded component for the given entity, and a bool for whether or not it exists
func (m *FileLoadedFactory) GetFileLoaded(id akara.EID) (*FileLoaded, bool) {
	component, found := m.FileLoaded.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FileLoaded), found
}