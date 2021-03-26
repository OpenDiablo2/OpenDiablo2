package d2pl2

// PL2Color24Bits represents an RGB color
type PL2Color24Bits struct {
	R uint8
	G uint8
	B uint8
}

// RGBA returns RGBA values of PL2Color
func (p *PL2Color24Bits) RGBA() uint32 {
	return toComposite(p.R, p.G, p.B, mask)
}

// SetRGBA sets PL2Color's value to rgba given
func (p *PL2Color24Bits) SetRGBA(rgba uint32) {
	p.R, p.G, p.B = toComponent(rgba)
}
