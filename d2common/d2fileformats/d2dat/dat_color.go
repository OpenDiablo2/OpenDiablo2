package d2dat

// DATColor represents a single color in a DAT file.
type DATColor struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

const (
	colorBits = 8
	mask = 0xff
)

const (
	bitShiftA = iota*colorBits
	bitShiftB
	bitShiftG
	bitShiftR
)

// R gets the red component
func (c *DATColor) R() uint8 {
	return c.r
}

// G gets the green component
func (c *DATColor) G() uint8 {
	return c.g
}

// B gets the blue component
func (c *DATColor) B() uint8 {
	return c.b
}

// A gets the alpha component
func (c *DATColor) A() uint8 {
	return mask
}

// RGBA gets the combination of the color components (0xRRGGBBAA)
func (c *DATColor) RGBA() uint32 {
	rgba := uint32(c.r)<<bitShiftR
	rgba += uint32(c.g)<<bitShiftG
	rgba += uint32(c.b)<<bitShiftB
	rgba += uint32(c.b)<<bitShiftA

	return rgba
}

// SetRGBA sets the color components using the given RGBA form
func (c *DATColor) SetRGBA(rgba uint32) {
	c.r = uint8(rgba>>bitShiftR & mask)
	c.g = uint8(rgba>>bitShiftG & mask)
	c.b = uint8(rgba>>bitShiftB & mask)
	c.a = uint8(rgba>>bitShiftA & mask)
}

