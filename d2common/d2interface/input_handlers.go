package d2interface

// InputEventHandler is an event handler
type InputEventHandler interface{}

// KeyDownHandler represents a handler for a keyboard key pressed event
type KeyDownHandler interface {
	OnKeyDown(event KeyEvent) bool
}

// KeyRepeatHandler represents a handler for a keyboard key held-down event; between a pressed and released.
type KeyRepeatHandler interface {
	OnKeyRepeat(event KeyEvent) bool
}

// KeyUpHandler represents a handler for a keyboard key release event
type KeyUpHandler interface {
	OnKeyUp(event KeyEvent) bool
}

// KeyCharsHandler represents a handler associated with a keyboard character pressed event
type KeyCharsHandler interface {
	OnKeyChars(event KeyCharsEvent) bool
}

// MouseButtonDownHandler represents a handler for a mouse button pressed event
type MouseButtonDownHandler interface {
	OnMouseButtonDown(event MouseEvent) bool
}

// MouseButtonRepeatHandler represents a handler for a mouse button held-down event; between a pressed and released.
type MouseButtonRepeatHandler interface {
	OnMouseButtonRepeat(event MouseEvent) bool
}

// MouseButtonUpHandler represents a handler for a mouse button release event
type MouseButtonUpHandler interface {
	OnMouseButtonUp(event MouseEvent) bool
}

// MouseMoveHandler represents a handler for a mouse button release event
type MouseMoveHandler interface {
	OnMouseMove(event MouseMoveEvent) bool
}
