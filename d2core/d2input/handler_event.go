package d2input

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// HandlerEvent is an event that EventHandlers will process and respond to
type HandlerEvent struct {
	keyMod    d2enum.KeyMod
	buttonMod d2enum.MouseButtonMod
	x         int
	y         int
}

// KeyMod yields the modifier for a key action
func (e *HandlerEvent) KeyMod() d2enum.KeyMod {
	return e.keyMod
}

// ButtonMod yields the modifier for a button action
func (e *HandlerEvent) ButtonMod() d2enum.MouseButtonMod {
	return e.buttonMod
}

// X returns the x screen coordinate for the event
func (e *HandlerEvent) X() int {
	return e.x
}

// Y returns the y screen coordinate for the event
func (e *HandlerEvent) Y() int {
	return e.y
}
