package d2input

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// MouseEvent represents a mouse event
type MouseEvent struct {
	HandlerEvent
	mouseButton d2enum.MouseButton
}

// KeyMod returns the key mod
func (e *MouseEvent) KeyMod() d2enum.KeyMod {
	return e.HandlerEvent.keyMod
}

// ButtonMod represents a button mod
func (e *MouseEvent) ButtonMod() d2enum.MouseButtonMod {
	return e.HandlerEvent.buttonMod
}

// X returns the event's X position
func (e *MouseEvent) X() int {
	return e.HandlerEvent.x
}

// Y returns the event's Y position
func (e *MouseEvent) Y() int {
	return e.HandlerEvent.y
}

// Button returns the mouse button
func (e *MouseEvent) Button() d2enum.MouseButton {
	return e.mouseButton
}
