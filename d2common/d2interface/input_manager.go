package d2interface

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// InputManager manages an InputService
type InputManager interface {
	Advance(elapsedTime, currentTime float64) error
	BindHandlerWithPriority(InputEventHandler, d2enum.Priority) error
	BindHandler(h InputEventHandler) error
	UnbindHandler(handler InputEventHandler) error
}
