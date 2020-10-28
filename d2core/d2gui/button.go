package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

type buttonState int

const (
	buttonStateDefault buttonState = iota
	buttonStatePressed
	buttonStatePressedToggled
)

const (
	grey = 0x404040ff
)

// Button is a user actionable drawable toggle switch
type Button struct {
	widgetBase

	width    int
	height   int
	state    buttonState
	surfaces []d2interface.Surface
}

func (b *Button) onMouseButtonDown(_ d2interface.MouseEvent) bool {
	b.state = buttonStatePressed

	return false
}

func (b *Button) onMouseButtonUp(_ d2interface.MouseEvent) bool {
	b.state = buttonStateDefault

	return false
}

func (b *Button) onMouseLeave(_ d2interface.MouseMoveEvent) bool {
	b.state = buttonStateDefault

	return false
}

func (b *Button) render(target d2interface.Surface) {
	target.Render(b.surfaces[b.state])
}

func (b *Button) getSize() (width, height int) {
	return b.width, b.height
}
