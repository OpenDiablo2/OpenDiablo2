package d2pl2

// PL2Color represents an RGBA color
type PL2Color struct {
	r uint8
	g uint8
	b uint8
	_ uint8
}

func (c PL2Color) R() uint8 {
	return c.r
}

func (c PL2Color) G() uint8 {
	return c.g
}

func (c PL2Color) B() uint8 {
	return c.g
}

func (c PL2Color) A() uint8 {
	return 0xff
}

func (c PL2Color) RGBA() uint32 {
	return uint32(c.r)<<24 | uint32(c.g)<<16 | uint32(c.b)<<8 | uint32(0xff)
}

func (c PL2Color) SetRGBA(u uint32) {
	c.r = byte((u>>24)&0xff)
	c.g = byte((u>>16)&0xff)
	c.b = byte((u>>8)&0xff)
}

func (c PL2Color) BGRA() uint32 {
	return uint32(c.b)<<8 | uint32(c.g)<<16 | uint32(c.r)<<24 | uint32(0xff)
}

func (c PL2Color) SetBGRA(u uint32) {
	c.b = byte((u>>24)&0xff)
	c.g = byte((u>>16)&0xff)
	c.r = byte((u>>8)&0xff)
}



