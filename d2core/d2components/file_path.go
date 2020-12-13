//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that File implements Component
var _ akara.Component = &File{}

// File represents a file as a path
type File struct {
	Path string
}

// New returns a File component. By default, it contains an empty string.
func (*File) New() akara.Component {
	return &File{}
}

// FileFactory is a wrapper for the generic component factory that returns File component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a File.
type FileFactory struct {
	*akara.ComponentFactory
}

// Add adds a File component to the given entity and returns it
func (m *FileFactory) Add(id akara.EID) *File {
	return m.ComponentFactory.Add(id).(*File)
}

// Get returns the File component for the given entity, and a bool for whether or not it exists
func (m *FileFactory) Get(id akara.EID) (*File, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*File), found
}
