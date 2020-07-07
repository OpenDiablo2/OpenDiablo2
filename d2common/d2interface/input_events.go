package d2interface

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// HandlerEvent holds the qualifiers for a key or mouse event
type HandlerEvent interface {
	KeyMod() d2enum.KeyMod
	ButtonMod() d2enum.MouseButtonMod
	X() int
	Y() int
}

// KeyEvent represents an event associated with a keyboard key
type KeyEvent interface {
	HandlerEvent
	Key() d2enum.Key
	// Duration represents the number of frames this key has been pressed for
	Duration() int
}

// KeyCharsEvent represents an event associated with a keyboard character being pressed
type KeyCharsEvent interface {
	HandlerEvent
	Chars() []rune
}

// MouseEvent represents a mouse event
type MouseEvent interface {
	HandlerEvent
	Button() d2enum.MouseButton
}

// MouseMoveEvent represents a mouse movement event
type MouseMoveEvent interface {
	HandlerEvent
}
