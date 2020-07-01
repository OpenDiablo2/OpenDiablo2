package d2input

import (
	"errors"
)

var (
	// ErrHasReg shows the input system already has a registered handler
	ErrHasReg = errors.New("input system already has provided handler")
	// ErrNotReg shows the input system has no registered handler
	ErrNotReg = errors.New("input system does not have provided handler")
)

var singleton inputManager

// HandlerEvent holds the qualifiers for a key or mouse event
type HandlerEvent struct {
	KeyMod    KeyMod
	ButtonMod MouseButtonMod
	X         int
	Y         int
}

// KeyEvent represents an event associated with a keyboard key
type KeyEvent struct {
	HandlerEvent
	Key Key
	// Duration represents the number of frames this key has been pressed for
	Duration int
}

// KeyCharsEvent represents an event associated with a keyboard character being pressed
type KeyCharsEvent struct {
	HandlerEvent
	Chars []rune
}

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

// MouseEvent represents a mouse event
type MouseEvent struct {
	HandlerEvent
	Button MouseButton
}

// MouseEvent represents a mouse movement event
type MouseMoveEvent struct {
	HandlerEvent
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

// Initialize creates a single global input manager based on a specific input service
func Initialize(inputService InputService) {
	singleton = inputManager{
		inputService: inputService,
	}
}

// Advance moves the input manager with the elapsed number of seconds.
func Advance(elapsed float64) error {
	return singleton.advance(elapsed)
}

// BindHandlerWithPriority adds an event handler with a specific call priority
func BindHandlerWithPriority(handler Handler, priority Priority) error {
	return singleton.bindHandler(handler, priority)
}

// BindHandler adds an event handler
func BindHandler(handler Handler) error {
	return BindHandlerWithPriority(handler, PriorityDefault)
}

// UnbindHandler removes a previously bound event handler
func UnbindHandler(handler Handler) error {
	return singleton.unbindHandler(handler)
}
