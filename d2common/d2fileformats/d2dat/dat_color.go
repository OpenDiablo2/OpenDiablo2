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
	mask      = 0xff
)

const (
	bitShift0 = iota * colorBits
	bitShift8
	bitShift16
	bitShift24
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
	return toComposite(c.r, c.g, c.b, c.a)
}

// SetRGBA sets the color components using the given RGBA form
func (c *DATColor) SetRGBA(rgba uint32) {
	c.r, c.g, c.b, c.a = toComponent(rgba)
}

// BGRA gets the combination of the color components (0xBBGGRRAA)
func (c *DATColor) BGRA() uint32 {
	return toComposite(c.b, c.g, c.r, c.a)
}

// SetBGRA sets the color components using the given BGRA form
func (c *DATColor) SetBGRA(bgra uint32) {
	c.b, c.g, c.r, c.a = toComponent(bgra)
}

func toComposite(w, x, y, z uint8) uint32 {
	composite := uint32(w) << bitShift24
	composite += uint32(x) << bitShift16
	composite += uint32(y) << bitShift8
	composite += uint32(z) << bitShift0

	return composite
}

func toComponent(wxyz uint32) (w, x, y, z uint8) {
	w = uint8(wxyz >> bitShift24 & mask)
	x = uint8(wxyz >> bitShift16 & mask)
	y = uint8(wxyz >> bitShift8 & mask)
	z = uint8(wxyz >> bitShift0 & mask)

	return w, x, y, z
}
