package d2input

// KeyCharsEvent represents a key character event
type KeyCharsEvent struct {
	HandlerEvent
	chars []rune
}

// Chars returns the characters
func (e *KeyCharsEvent) Chars() []rune {
	return e.chars
}
