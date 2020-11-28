package d2input

import "github.com/hajimehoshi/ebiten/v2"

type Modifier = int

const (
	ModAlt     = Modifier(ebiten.KeyAlt)
	ModControl = Modifier(ebiten.KeyControl)
	ModShift   = Modifier(ebiten.KeyShift)
)
