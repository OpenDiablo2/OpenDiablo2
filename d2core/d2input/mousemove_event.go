package d2input

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// MouseMoveEvent represents a mouse movement event
type MouseMoveEvent struct {
	HandlerEvent
}

// KeyMod represents the key mod
func (e *MouseMoveEvent) KeyMod() d2enum.KeyMod {
	return e.HandlerEvent.keyMod
}

// ButtonMod represents the button mod
func (e *MouseMoveEvent) ButtonMod() d2enum.MouseButtonMod {
	return e.HandlerEvent.buttonMod
}

// X represents the X position
func (e *MouseMoveEvent) X() int {
	return e.HandlerEvent.x
}

// Y represents the Y position
func (e *MouseMoveEvent) Y() int {
	return e.HandlerEvent.y
}
