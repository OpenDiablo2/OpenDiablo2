package d2math

import "math"

const (
	// Epsilon is used as the threshold for 'almost equal' operations.
	Epsilon float64 = 0.0001

	RadToDeg float64 = 57.29578
	RadFull  float64 = 6.283185253783088
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

// ClampFloat64 returns a clamped to min and max.
func ClampFloat64(a, min, max float64) float64 {
	if a > max {
		return max
	} else if a < min {
		return min
	}
	return a
}

// Sign returns the sign of a.
func Sign(a float64) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}
