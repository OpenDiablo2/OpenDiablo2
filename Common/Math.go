package Common

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

// BytesToInt32 converts 4 bytes to int32
func BytesToInt32(b []byte) int32 {
	// equivalnt of return int32(binary.LittleEndian.Uint32(b))
	return int32(uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24)
}
