package d2pl2

const (
	bitShift0 = 8 * iota
	bitShift8
	bitShift16
	bitShift24

	mask = 0xff
)

// PL2Color represents an RGBA color
type PL2Color struct {
	R uint8
	G uint8
	B uint8
	_ uint8
}

// RGBA returns RGBA values of PL2Color
func (p *PL2Color) RGBA() uint32 {
	return toComposite(p.R, p.G, p.B, mask)
}

// SetRGBA sets PL2Color's value to rgba given
func (p *PL2Color) SetRGBA(rgba uint32) {
	p.R, p.G, p.B = toComponent(rgba)
}

func toComposite(w, x, y, z uint8) uint32 {
	composite := uint32(w) << bitShift24
	composite += uint32(x) << bitShift16
	composite += uint32(y) << bitShift8
	composite += uint32(z) << bitShift0

	return composite
}

func toComponent(wxyz uint32) (w, x, y uint8) {
	w = uint8(wxyz >> bitShift24 & mask)
	x = uint8(wxyz >> bitShift16 & mask)
	y = uint8(wxyz >> bitShift8 & mask)

	return w, x, y
}
