//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"io"

	"github.com/gravestench/akara"
)

// static check that Wav implements Component
var _ akara.Component = &Wav{}

// Wav is a component that contains an embedded io.ReadSeeker for streaming wav audio files
type Wav struct {
	Data io.ReadSeeker
}

// New returns a new Wav component. By default, it contains a nil instance.
func (*Wav) New() akara.Component {
	return &Wav{}
}

// WavFactory is a wrapper for the generic component factory that returns Wav component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Wav.
type WavFactory struct {
	Wav *akara.ComponentFactory
}

// AddWav adds a Wav component to the given entity and returns it
func (m *WavFactory) AddWav(id akara.EID) *Wav {
	return m.Wav.Add(id).(*Wav)
}

// GetWav returns the Wav component for the given entity, and a bool for whether or not it exists
func (m *WavFactory) GetWav(id akara.EID) (*Wav, bool) {
	component, found := m.Wav.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Wav), found
}
