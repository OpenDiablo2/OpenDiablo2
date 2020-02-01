package d2map

type Camera struct {
	x float64
	y float64
}

func (c *Camera) MoveTo(x, y float64) {
	c.x = x
	c.y = y
}

func (c *Camera) MoveBy(x, y float64) {
	c.x += x
	c.y += y
}

func (c *Camera) GetPosition() (float64, float64) {
	return c.x, c.y
}
