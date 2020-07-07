package d2math

import "math"

const (
	// Epsilon is used as the threshold for 'almost equal' operations.
	Epsilon float64 = 0.0001
)

// CompareFloat64Fuzzy returns an integer between -1 and 1 describing
// the comparison of floats a and b. 0 will be returned if the
// absolute difference between a and b is less than Epsilon.
func CompareFloat64Fuzzy(a, b float64) int {
	delta := a - b
	if math.Abs(delta) < Epsilon {
		return 0
	}

	if delta > 0 {
		return 1
	}

	return -1
}
