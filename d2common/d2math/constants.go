package d2math

import "math"

const (
	PI               = math.Pi
	PI2              = PI * 2
	TAU              = PI / 2
	Epsilon          = 1e-6
	DegreesToRadians = PI / 180

	// RadToDeg is used to convert anges in radians to degrees by multiplying the radians by RadToDeg. Similarly,degrees
	// are converted to radians when dividing by RadToDeg.
	RadToDeg float64 = 57.29578

	// RadFull is the radian equivalent of 360 degrees.
	RadFull float64 = 6.283185253783088
)
