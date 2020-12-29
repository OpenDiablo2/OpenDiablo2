package d2math

import "math"

const (
	// Epsilon is used as the threshold for 'almost equal' operations.
	Epsilon float64 = 0.0001

	// RadToDeg is used to convert anges in radians to degrees by multiplying the radians by RadToDeg. Similarly,degrees
	// are converted to radians when dividing by RadToDeg.
	RadToDeg float64 = 57.29578

	// RadFull is the radian equivalent of 360 degrees.
	RadFull float64 = 6.283185253783088
)

// EqualsApprox returns true if the difference between a and b is less than Epsilon.
func EqualsApprox(a, b float64) bool {
	return Abs(a-b) < Epsilon
}

// CompareApprox returns an integer between -1 and 1 describing the comparison of floats a and b. 0 will be returned if
// the absolute difference between a and b is less than Epsilon.
func CompareApprox(a, b float64) int {
	delta := a - b

	if Abs(delta) < Epsilon {
		return 0
	}

	if delta > 0 {
		return 1
	}

	return -1
}

// Abs returns the absolute value of a. It is a less CPU intensive version of the standard library math.Abs().
func Abs(a float64) float64 {
	if a < 0 {
		return -a
	}

	return a
}

// Clamp returns a clamped to min and max.
func Clamp(a, min, max float64) float64 {
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

// Lerp returns the linear interpolation from a to b using interpolator x.
func Lerp(a, b, x float64) float64 {
	return a + x*(b-a)
}

// Unlerp returns the intepolator Lerp would require to return x when given
// a and b. The x argument of this function can be thought of as the return
// value of lerp. The return value of this function can be used as x in
// Lerp.
func Unlerp(a, b, x float64) float64 {
	return (x - a) / (b - a)
}

// WrapInt wraps x to between 0 and max. For example WrapInt(450, 360) would return 90.
func WrapInt(x, max int) int {
	wrapped := x % max

	if wrapped < 0 {
		return max + wrapped
	}

	return wrapped
}

// MinInt returns the minimum of the given values
func MinInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// MaxInt returns the maximum of the given values
func MaxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Min returns the lower of two values
func Min(a, b uint32) uint32 {
	if a < b {
		return a
	}

	return b
}

// Max returns the higher of two values
func Max(a, b uint32) uint32 {
	if a > b {
		return a
	}

	return b
}

// MaxInt32 returns the higher of two values
func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

// AbsInt32 returns the absolute of the given int32
func AbsInt32(a int32) int32 {
	if a < 0 {
		return -a
	}

	return a
}

// MinInt32 returns the higher of two values
func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}

	return b
}

// BytesToInt32 converts 4 bytes to int32

// IsoToScreen converts isometric coordinates to screenspace coordinates

// ScreenToIso converts screenspace coordinates to isometric coordinates

// GetRadiansBetween returns the radians between two points. 0rad is facing to the right.
func GetRadiansBetween(p1X, p1Y, p2X, p2Y float64) float64 {
	deltaY := p2Y - p1Y
	deltaX := p2X - p1X

	return math.Atan2(deltaY, deltaX)
}

// ClampInt ensures that the given value is between or equal to the given min or max
func ClampInt(value, min, max int) int {
	return int(math.Min(math.Max(float64(value), float64(min)), float64(max)))
}
