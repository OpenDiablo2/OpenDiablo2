package d2util

import "image/color"

// Color converts an rgba uint32 to a colorEnabled.RGBA
func Color(rgba uint32) color.RGBA {
	result := color.RGBA{}
	a, b, g, r := 0, 1, 2, 3
	byteWidth := 8
	byteMask := 0xff

	for idx := 0; idx < 4; idx++ {
		shift := idx * byteWidth
		component := uint8(rgba>>shift) & uint8(byteMask)

		switch idx {
		case a:
			result.A = component
		case b:
			result.B = component
		case g:
			result.G = component
		case r:
			result.R = component
		}
	}

	return result
}
