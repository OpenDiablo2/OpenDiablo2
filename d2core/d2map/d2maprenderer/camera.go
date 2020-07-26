package d2maprenderer

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"

// Camera is the position of the Camera perspective in orthogonal world space. See viewport.go.
type Camera struct {
	position *d2vector.Position
	target   *d2vector.Position
}

// MoveTo sets the position of the Camera to the given position
func (c *Camera) MoveTo(position *d2vector.Position) {
	c.position = position
	c.target = position
}

// MoveBy adds the given vector to the current position of the Camera.
func (c *Camera) MoveBy(vector *d2vector.Vector) {
	c.position.Add(vector)
}

// SetTarget sets the target position
func (c *Camera) SetTarget(target *d2vector.Position) {
	c.target = target
}

// MoveTargetBy adds the given vector to the current position of the Camera.
func (c *Camera) MoveTargetBy(vector *d2vector.Vector) {
	if c.target == nil {
		v := c.position.Clone()
		c.target = &d2vector.Position{Vector: *v}
	}

	c.target.Add(vector)
}

// ClearTarget sets the target position
func (c *Camera) ClearTarget() {
	c.target = nil
}

// GetPosition returns the Camera position
func (c *Camera) GetPosition() *d2vector.Position {
	return c.position
}

// Advance returns the Camera position
func (c *Camera) Advance(elapsed float64) {
	c.advanceToTarget(elapsed)
}

func (c *Camera) advanceToTarget(_ float64) {
	if c.target != nil {
		diff := c.position.World().Subtract(c.target.World())
		diff.Scale(-0.85)
		c.MoveBy(diff)
	}
}
