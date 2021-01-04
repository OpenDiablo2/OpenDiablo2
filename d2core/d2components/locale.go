//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that Locale implements Component
var _ akara.Component = &Locale{}

// Locale represents a file as a path
type Locale struct {
	Code byte
	String string
}

// New returns a Locale component. By default, it contains an empty string.
func (*Locale) New() akara.Component {
	return &Locale{}
}

// LocaleFactory is a wrapper for the generic component factory that returns Locale component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Locale.
type LocaleFactory struct {
	Locale *akara.ComponentFactory
}

// AddLocale adds a Locale component to the given entity and returns it
func (m *LocaleFactory) AddLocale(id akara.EID) *Locale {
	return m.Locale.Add(id).(*Locale)
}

// GetLocale returns the Locale component for the given entity, and a bool for whether or not it exists
func (m *LocaleFactory) GetLocale(id akara.EID) (*Locale, bool) {
	component, found := m.Locale.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Locale), found
}
