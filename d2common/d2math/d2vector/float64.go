package d2vector

// Float64 is the implementation of Vector using float64
// to store x and y.
type Float64 struct {
	x, y float64
}

// NewFloat64 creates a new Float64 and returns a pointer to it.
func NewFloat64(x, y float64) *Float64 {
	return &Float64{x, y}
}
