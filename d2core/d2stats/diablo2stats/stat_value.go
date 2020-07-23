package diablo2stats

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

// static check that Diablo2StatValue implements StatValue
var _ d2stats.StatValue = &Diablo2StatValue{}

// Diablo2StatValue is a diablo 2 implementation of a StatValue
type Diablo2StatValue struct {
	int       int
	float     float64
	_stringer func(d2stats.StatValue) string
	_type     d2stats.StatValueType
}

// Type returns the stat value type
func (sv *Diablo2StatValue) Type() d2stats.StatValueType {
	return sv._type
}

// Clone returns a deep copy of the stat value
func (sv Diablo2StatValue) Clone() d2stats.StatValue {
	clone := &Diablo2StatValue{}

	switch sv._type {
	case d2stats.StatValueInt:
		clone.SetInt(sv.Int())
	case d2stats.StatValueFloat:
		clone.SetFloat(sv.Float())
	}

	clone._stringer = sv._stringer

	return clone
}

// Int returns the integer version of the stat value
func (sv *Diablo2StatValue) Int() int {
	return sv.int
}

// String returns a string version of the value
func (sv *Diablo2StatValue) String() string {
	return sv._stringer(sv)
}

// Float returns a float64 version of the value
func (sv *Diablo2StatValue) Float() float64 {
	return sv.float
}

// SetInt sets the stat value using an int
func (sv *Diablo2StatValue) SetInt(i int) d2stats.StatValue {
	sv.int = i
	sv.float = float64(i)

	return sv
}

// SetFloat sets the stat value using a float64
func (sv *Diablo2StatValue) SetFloat(f float64) d2stats.StatValue {
	sv.int = int(f)
	sv.float = f

	return sv
}

// Stringer returns the string evaluation function
func (sv *Diablo2StatValue) Stringer() func(d2stats.StatValue) string {
	return sv._stringer
}

// SetStringer sets the string evaluation function
func (sv *Diablo2StatValue) SetStringer(f func(d2stats.StatValue) string) d2stats.StatValue {
	sv._stringer = f
	return sv
}
