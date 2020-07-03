package d2interface

// HandlerEvent holds the qualifiers for a key or mouse event
type HandlerEvent interface {
	KeyMod() KeyMod
	ButtonMod() MouseButtonMod
	X() int
	Y() int
}

// KeyEvent represents an event associated with a keyboard key
type KeyEvent interface {
	HandlerEvent
	Key() Key
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
	Button() MouseButton
}

// MouseMoveEvent represents a mouse movement event
type MouseMoveEvent interface {
	HandlerEvent
}
