package d2input

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	ebiten_input "github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/ebiten"
)

var (
	// ErrHasReg shows the input system already has a registered handler
	ErrHasReg = errors.New("input system already has provided handler")
	// ErrNotReg shows the input system has no registered handler
	ErrNotReg = errors.New("input system does not have provided handler")
)

var singleton *inputManager // TODO remove this singleton

// Initialize creates a single global input manager based on a specific input service
func Create() (d2interface.InputManager, error) {
	singleton = &inputManager{
		inputService: ebiten_input.InputService{},
	}

	return singleton, nil
}

// Advance moves the input manager with the elapsed number of seconds.
func Advance(elapsed, current float64) error {
	return singleton.Advance(elapsed, current)
}

// BindHandlerWithPriority adds an event handler with a specific call priority
func BindHandlerWithPriority(handler d2interface.InputEventHandler, priority d2interface.Priority) error {
	return singleton.BindHandlerWithPriority(handler, priority)
}

// BindHandler adds an event handler
func BindHandler(handler d2interface.InputEventHandler) error {
	return BindHandlerWithPriority(handler, d2interface.PriorityDefault)
}

// UnbindHandler removes a previously bound event handler
func UnbindHandler(handler d2interface.InputEventHandler) error {
	return singleton.UnbindHandler(handler)
}
