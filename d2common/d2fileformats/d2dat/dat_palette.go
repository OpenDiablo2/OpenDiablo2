package d2dat

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

const (
	numColors = 256
)

// DATPalette represents a 256 color palette.
type DATPalette struct {
	colors [numColors]d2interface.Color
}

// NumColors returns the number of colors in the palette
func (p *DATPalette) NumColors() int {
	return len(p.colors)
}

// GetColors returns the slice of colors in the palette
func (p *DATPalette) GetColors() [numColors]d2interface.Color {
	return p.colors
}

// GetColor returns a color by index
func (p *DATPalette) GetColor(idx int) (d2interface.Color, error) {
	if color := p.colors[idx]; color != nil {
		return color, nil
	}

	return nil, fmt.Errorf("cannot find color index '%d in palette'", idx)
}
