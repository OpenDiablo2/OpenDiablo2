package d2common

import (
	"math"
)

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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

// GetAngleBetween returns the angle between two points. 0deg is facing to the right.
func GetAngleBetween(p1X, p1Y, p2X, p2Y float64) int {
	deltaY := p1Y - p2Y
	deltaX := p2X - p1X

	result := math.Atan2(deltaY, deltaX) * (180 / math.Pi)
	iResult := int(result)
	for iResult < 0 {
		iResult += 360
	}
	for iResult >= 360 {
		iResult -= 360
	}
	return iResult
}

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

// Return the new adjusted value, as well as any remaining amount after the max
func AdjustWithRemainder(sourceValue, adjustment, targetvalue float64) (newValue, remainder float64) {
	if adjustment == 0 || math.Abs(adjustment) < 0.000001 {
		return sourceValue, 0
	}
	adjustNegative := adjustment < 0.0
	maxNegative := targetvalue-sourceValue < 0.0
	if adjustNegative != maxNegative {
		// FIXME: This shouldn't happen but it happens all the time..
		return sourceValue, 0
		//panic("Cannot move towards the opposite direction...")
	}

	finalValue := sourceValue + adjustment
	if !adjustNegative {
		if finalValue > targetvalue {
			diff := finalValue - targetvalue //  RoundToDecial(finalValue-targetvalue, 6)
			return targetvalue, diff
		}
		return finalValue, 0
	}

	if finalValue < targetvalue {
		return targetvalue, RoundToDecial(finalValue-targetvalue, 6)
	}
	return finalValue, 0
}

func RoundToDecial(f float64, d int) float64 {
	digits := float64(math.Pow10(d))
	return math.Trunc(f*digits) / digits
}
