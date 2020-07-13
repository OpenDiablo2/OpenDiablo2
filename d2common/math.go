package d2common

import (
	"math"
)

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

// AlmostEqual returns true if two values are within threshold from each other
func AlmostEqual(a, b, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}
