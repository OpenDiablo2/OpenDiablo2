package d2interface

// InputManager manages an InputService
type InputManager interface {
	Advance(elapsedTime, currentTime float64) error
	BindHandlerWithPriority(InputEventHandler, Priority) error
	BindHandler(h InputEventHandler) error
	UnbindHandler(handler InputEventHandler) error
}
