package d2input

import (
	"errors"
)

var (
	ErrHasReg = errors.New("input system already has provided handler")
	ErrNotReg = errors.New("input system does not have provided handler")
)

type Priority int

const (
	PriorityLow Priority = iota
	PriorityDefault
	PriorityHigh
)

type HandlerEvent struct {
	KeyMod    KeyMod
	ButtonMod MouseButtonMod
	X         int
	Y         int
}

type KeyEvent struct {
	HandlerEvent
	Key Key
	// Duration represents the number of frames this key has been pressed for
	Duration int
}

type KeyCharsEvent struct {
	HandlerEvent
	Chars []rune
}

type MouseEvent struct {
	HandlerEvent
	Button MouseButton
}

type MouseMoveEvent struct {
	HandlerEvent
}

type Handler interface{}

type KeyDownHandler interface {
	OnKeyDown(event KeyEvent) bool
}

type KeyRepeatHandler interface {
	OnKeyRepeat(event KeyEvent) bool
}

type KeyUpHandler interface {
	OnKeyUp(event KeyEvent) bool
}

type KeyCharsHandler interface {
	OnKeyChars(event KeyCharsEvent) bool
}

type MouseButtonDownHandler interface {
	OnMouseButtonDown(event MouseEvent) bool
}

type MouseButtonRepeatHandler interface {
	OnMouseButtonRepeat(event MouseEvent) bool
}

type MouseButtonUpHandler interface {
	OnMouseButtonUp(event MouseEvent) bool
}

type MouseMoveHandler interface {
	OnMouseMove(event MouseMoveEvent) bool
}

var singleton inputManager

func Initialize(inputService InputService) {
	singleton = inputManager{
		inputService: inputService,
	}
}

func Advance(elapsed float64) error {
	return singleton.advance(elapsed)
}

func BindHandlerWithPriority(handler Handler, priority Priority) error {
	return singleton.bindHandler(handler, priority)
}

func BindHandler(handler Handler) error {
	return BindHandlerWithPriority(handler, PriorityDefault)
}

func UnbindHandler(handler Handler) error {
	return singleton.unbindHandler(handler)
}
