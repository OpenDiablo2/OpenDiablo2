package rectangle

import (
	"math"

	"github.com/gravestench/pho/phomath"
)

// GetAspectRatio returns the aspect ratio (width/height) of the given rectangle
func GetAspectRatio(r *Rectangle) float64 {
	if r.Height < phomath.Epsilon {
		return math.NaN()
	}

	return r.Width / r.Height
}
