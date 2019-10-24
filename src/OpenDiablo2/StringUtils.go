package OpenDiablo2

import "strconv"

// StringToInt converts a string to an integer
func StringToInt(text string) int {
	result, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return result
}

// StringToUint8 converts a string to an uint8
func StringToUint8(text string) uint8 {
	result, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	if result < 0 || result > 255 {
		panic("value out of range of byte")
	}
	return uint8(result)
}

// StringToInt8 converts a string to an int8
func StringToInt8(text string) int8 {
	result, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	if result < -128 || result > 122 {
		panic("value out of range of a signed byte")
	}
	return int8(result)
}
