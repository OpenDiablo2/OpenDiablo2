package d2input

import "github.com/hajimehoshi/ebiten/v2"

// Modifier represents a keyboard modifier key
type Modifier = int

// Modifiers
const (
	ModAlt     = Modifier(ebiten.KeyAlt)
	ModControl = Modifier(ebiten.KeyControl)
	ModShift   = Modifier(ebiten.KeyShift)
)
