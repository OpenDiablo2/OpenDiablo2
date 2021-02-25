package d2cof

// SpeedToFPS returns FPS value basing on cof's speed
func (c *COF) SpeedToFPS() float64 {
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
