//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that FileSource implements Component
var _ akara.Component = &FileSource{}

// AbstractSource is the abstract representation of what a file source is
type AbstractSource interface {
	Path() string // the path of the source itself
	Open(path *File) (d2interface.DataStream, error)
}

// FileSource contains an embedded file source interface, something that can open files
type FileSource struct {
	AbstractSource
}

// New returns a FileSource component. By default, it contains a nil instance.
func (*FileSource) New() akara.Component {
	return &FileSource{}
}

// FileSourceFactory is a wrapper for the generic component factory that returns FileSource component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a FileSource.
type FileSourceFactory struct {
	*akara.ComponentFactory
}

// Add adds a FileSource component to the given entity and returns it
func (m *FileSourceFactory) Add(id akara.EID) *FileSource {
	return m.ComponentFactory.Add(id).(*FileSource)
}

// Get returns the FileSource component for the given entity, and a bool for whether or not it exists
func (m *FileSourceFactory) Get(id akara.EID) (*FileSource, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FileSource), found
}
