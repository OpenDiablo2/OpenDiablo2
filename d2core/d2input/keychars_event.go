package d2input

// KeyCharsEvent represents a key character event
type KeyCharsEvent struct {
	chars []rune
	HandlerEvent
}

// Chars returns the characters
func (e *KeyCharsEvent) Chars() []rune {
	return e.chars
}
