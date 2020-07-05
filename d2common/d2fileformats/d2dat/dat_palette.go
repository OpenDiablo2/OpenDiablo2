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

func (p *DATPalette) NumColors() int {
	return len(p.colors)
}

func (p *DATPalette) GetColors() [numColors]d2interface.Color {
	return p.colors
}

func (p *DATPalette) GetColor(idx int) (d2interface.Color, error) {
	if color := p.colors[idx]; color != nil {
		return color, nil
	}
	return nil, fmt.Errorf("Cannot find color index '%d in palette'", idx)
}

