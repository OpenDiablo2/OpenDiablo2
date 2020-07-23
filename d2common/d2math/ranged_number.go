package d2math

import "fmt"

// RangedNumber is a number with a min and max range
type RangedNumber struct {
	min int
	max int
}

// Min returns the min value, swapping min/max if not ordered
func (rn *RangedNumber) Min() int {
	if rn.min > rn.max {
		rn.Set(rn.max, rn.min)
	}

	return rn.min
}

// Max returns the max value, swapping min/max if not ordered
func (rn *RangedNumber) Max() int {
	if rn.min > rn.max {
		rn.Set(rn.max, rn.min)
	}

	return rn.max
}

// Set sets the min and max values, ordering the arguments if necessary
func (rn *RangedNumber) Set(min, max int) *RangedNumber {
	rn.SetMin(min)
	rn.SetMax(max)

	return rn
}

// SetMin sets the minimum value
func (rn *RangedNumber) SetMin(min int) *RangedNumber {
	rn.min = min
	if rn.min > rn.max {
		rn.min = rn.max
		rn.max = min
	}

	return rn
}

// SetMax sets the maximum value
func (rn *RangedNumber) SetMax(max int) *RangedNumber {
	rn.max = max
	if rn.min > rn.max {
		rn.max = rn.min
		rn.min = max
	}

	return rn
}

// Clone creates a new copy of a ranged number, with the same min/max
func (rn RangedNumber) Clone() *RangedNumber {
	return &rn
}

// Copy copies the min/max values of the given ranged number
func (rn *RangedNumber) Copy(other *RangedNumber) *RangedNumber {
	return rn.Set(other.min, other.max)
}

// Equals checks equality with the given ranged number
func (rn *RangedNumber) Equals(other *RangedNumber) bool {
	return rn.min == other.min && rn.max == other.max
}

// Add adds the given ranged number to this one, returning this one
func (rn *RangedNumber) Add(other *RangedNumber) *RangedNumber {
	return rn.Set(rn.min+other.min, rn.max+other.max)
}

// Sub subtracts the given ranged number from this one, returning this one
func (rn *RangedNumber) Sub(other *RangedNumber) *RangedNumber {
	return rn.Set(rn.min-other.min, rn.max-other.max)
}

// Mul multiplies this ranged number by the given ranged number, returning this one
func (rn *RangedNumber) Mul(other *RangedNumber) *RangedNumber {
	return rn.Set(rn.min*other.min, rn.max*other.max)
}

// Div divides this ranged number by the given ranged number, returning this one
func (rn *RangedNumber) Div(other *RangedNumber) *RangedNumber {
	return rn.Set(rn.min/other.min, rn.max/other.max)
}

func (rn *RangedNumber) String() string {
	if rn.Min() == rn.Max() { // ensures ordering
		return fmt.Sprintf("%d", rn.min)
	}

	return fmt.Sprintf("%d to %d", rn.min, rn.max)
}
