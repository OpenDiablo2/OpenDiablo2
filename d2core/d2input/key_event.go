package d2input

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// KeyEvent represents key events
type KeyEvent struct {
	HandlerEvent
	key d2enum.Key
	// Duration represents the number of frames this key has been pressed for
	duration int
}

// Key returns the key
func (e *KeyEvent) Key() d2enum.Key {
	return e.key
}

// Duration returns the duration
func (e *KeyEvent) Duration() int {
	return e.duration
}
