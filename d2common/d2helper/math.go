package d2helper

import "math"

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

func NextPow2(x int32) int32 {
	result := int32(1)
	for result < x {
		result *= 2
	}
	return result
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
func BytesToInt32(b []byte) int32 {
	// equivalnt of return int32(binary.LittleEndian.Uint32(b))
	return int32(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24)
}

// IsoToScreen converts isometric coordinates to screenspace coordinates
func IsoToScreen(isoX, isoY, modX, modY float64) (float64, float64) {
	screenX := (isoX - isoY) * 80
	screenY := (isoX + isoY) * 40
	return screenX + modX, screenY + modY
}

// ScreenToIso converts screenspace coordinates to isometric coordinates
func ScreenToIso(sx, sy float64) (float64, float64) {
	x := (sx/80 + sy/40) / 2
	y := (sy/40 - (sx / 80)) / 2
	return x, y
}

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

// AlmostEqual returns true if two values are within threshold from each other
func AlmostEqual(a, b, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}
