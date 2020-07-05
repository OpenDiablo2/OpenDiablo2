package d2vector

import (
	"fmt"
	"math"
	"math/big"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// Float64 is the implementation of Vector using float64.
type Float64 struct {
	x, y float64
}

// NewFloat64 creates a new Float64 and returns a pointer to it.
func NewFloat64(x, y float64) d2interface.Vector {
	return &Float64{x, y}
}

// XYBigFloat returns the values as big.Float.
func (f *Float64) XYBigFloat() (*big.Float, *big.Float) {
	bf := NewBigFloat(f.x, f.y)
	return bf.XYBigFloat()
}

// XYFloat64 returns the values as float64.
func (f *Float64) XYFloat64() (*float64, *float64) {
	return &f.x, &f.y
}

// Equals check whether this Vector is equal to a given Vector.
func (f *Float64) Equals(v d2interface.Vector) bool {
	vx, vy := v.XYFloat64()
	return f.x == *vx && f.y == *vy
}

// EqualsF checks if the Vector is approximately equal
// to the given Vector.
func (f *Float64) EqualsF(v d2interface.Vector) bool {
	x, y := f.CompareF(v)
	return x == 0 && y == 0
}

// CompareF performs a fuzzy comparison and returns 2
// ints represending the -1 to 1 comparison of x and y.
func (f *Float64) CompareF(v d2interface.Vector) (int, int) {
	vx, vy := v.XYFloat64()
	return d2math.CompareFloat64Fuzzy(&f.x, vx),
		d2math.CompareFloat64Fuzzy(&f.y, vy)
}

// Set sets the vector values to the given float64 values.
func (f *Float64) Set(x, y float64) d2interface.Vector {
	f.x = x
	f.y = y

	return f
}

// Clone creates a copy of this Vector.
func (f *Float64) Clone() d2interface.Vector {
	return NewFloat64(f.x, f.y)
}

// Floor rounds the vector down to the nearest whole numbers.
func (f *Float64) Floor() d2interface.Vector {
	f.x = math.Floor(f.x)
	f.y = math.Floor(f.y)

	return f
}

// Add to this Vector the components of the given Vector.
func (f *Float64) Add(v d2interface.Vector) d2interface.Vector {
	vx, vy := v.XYFloat64()
	f.x += *vx
	f.y += *vy

	return f
}

// Subtract from this Vector from the components of the given Vector.
func (f *Float64) Subtract(v d2interface.Vector) d2interface.Vector {
	vx, vy := v.XYFloat64()
	f.x -= *vx
	f.y -= *vy

	return v
}

// Multiply this Vector by the components of the given Vector.
func (f *Float64) Multiply(v d2interface.Vector) d2interface.Vector {
	vx, vy := v.XYFloat64()
	f.x *= *vx
	f.y *= *vy

	return v
}

func (f *Float64) String() string {
	return fmt.Sprintf("Float64{%g, %g}", f.x, f.y)
}
