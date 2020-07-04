package d2vector

import (
	"math/big"
)

// Float64 is the implementation of Vector using float64
// to store x and y.
type Float64 struct {
	x, y float64
}

// NewFloat64 creates a new Float64 and returns a pointer to it.
func NewFloat64(x, y float64) *Float64 {
	return &Float64{x, y}
}

// XBig returns the big.Float value of x.
func (f *Float64) XBig() *big.Float {
	return big.NewFloat(f.x)
}

// YBig returns the big.Float value of y.
func (f *Float64) YBig() *big.Float {
	return big.NewFloat(f.y)
}

// X64 returns the float64 value of x.
func (f *Float64) X64() float64 {
	return f.x
}

// Y64 returns the float64 value of y.
func (f *Float64) Y64() float64 {
	return f.y
}

// AsBigFloat returns a pointer to a new BigFloat
// based on the values of f.
func (f *Float64) AsBigFloat() *BigFloat {
	return NewBigFloat(f.x, f.y)
}

// CopyBigFloat copies the values from a BigFloat to f.
func (f *Float64) CopyBigFloat(bf BigFloat) {
	f.x, _ = bf.X64()
	f.y, _ = bf.Y64()
}

// comparison, in a separate interface?
