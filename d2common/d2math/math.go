package d2math

import "math"

const (
	// Epsilon is used as the threshold for 'almost equal' operations.
	epsilon float64 = 0.0001
)

// CompareFloat64Fuzzy returns an integer between -1 and 1 describing
// the comparison of floats a and b. 0 will be returned if the
// absolute difference between a and b is less than epsilon.
func CompareFloat64Fuzzy(a, b *float64) int {
	difference := *a - *b
	if math.Abs(difference) < epsilon {
		return 0
	}

	if difference > 0 {
		return 1
	}

	return -1
}
