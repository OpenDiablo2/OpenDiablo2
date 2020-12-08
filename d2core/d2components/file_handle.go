//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that FileHandle implements Component
var _ akara.Component = &FileHandle{}

// FileHandle is a component that contains a data stream for file data
type FileHandle struct {
	Data d2interface.DataStream
}

// New returns a FileHandle component. By default, it contains a nil instance.
func (*FileHandle) New() akara.Component {
	return &FileHandle{}
}

// FileHandleFactory is a wrapper for the generic component factory that returns FileHandle component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a FileHandle.
type FileHandleFactory struct {
	*akara.ComponentFactory
}

// Add adds a FileHandle component to the given entity and returns it
func (m *FileHandleFactory) Add(id akara.EID) *FileHandle {
	return m.ComponentFactory.Add(id).(*FileHandle)
}

// Get returns the FileHandle component for the given entity, and a bool for whether or not it exists
func (m *FileHandleFactory) Get(id akara.EID) (*FileHandle, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FileHandle), found
}
