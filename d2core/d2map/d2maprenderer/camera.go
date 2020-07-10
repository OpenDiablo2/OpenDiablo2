package d2maprenderer

// Camera is the position of the camera perspective in orthogonal world space. See viewport.go.
// TODO: Has a coordinate (issue #456)
type Camera struct {
	x float64
	y float64
}

// MoveTo sets the position of the camera to the given x and y coordinates.
func (c *Camera) MoveTo(x, y float64) {
	c.x = x
	c.y = y
}

// MoveBy adds the given vector to the current position of the camera.
func (c *Camera) MoveBy(x, y float64) {
	c.x += x
	c.y += y
}

// GetPosition returns the camera x and y position.
func (c *Camera) GetPosition() (float64, float64) {
	return c.x, c.y
}
