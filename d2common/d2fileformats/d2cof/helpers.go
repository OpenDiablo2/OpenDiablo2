package d2cof

// FPS returns FPS value basing on cof's speed
func (c *COF) FPS() float64 {
	const (
		baseFPS      = 25
		speedDivisor = 256
	)

	fps := baseFPS * (float64(c.Speed) / speedDivisor)
	if fps == 0 {
		fps = baseFPS
	}

	return fps
}

// Duration returns animation's duration
func (c *COF) Duration() float64 {
	const (
		milliseconds = 1000
	)

	frameDelay := milliseconds / c.FPS()

	return float64(c.FramesPerDirection) * frameDelay
}
