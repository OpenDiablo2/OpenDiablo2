package d2input

import (
	"errors"
)

var (
	ErrHasReg = errors.New("input system already has provided handler")
	ErrNotReg = errors.New("input system does not have provided handler")
)

type InputBackend int

const (
	Ebiten InputBackend = iota
)

type Priority int

const (
	PriorityLow Priority = iota
	PriorityDefault
	PriorityHigh
)

type Key int

const (
	Key0			Key = iota
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeyA
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	KeyApostrophe
	KeyBackslash
	KeyBackspace
	KeyCapsLock
	KeyComma
	KeyDelete
	KeyDown
	KeyEnd
	KeyEnter
	KeyEqual
	KeyEscape
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyGraveAccent
	KeyHome
	KeyInsert
	KeyKP0
	KeyKP1
	KeyKP2
	KeyKP3
	KeyKP4
	KeyKP5
	KeyKP6
	KeyKP7
	KeyKP8
	KeyKP9
	KeyKPAdd
	KeyKPDecimal
	KeyKPDivide
	KeyKPEnter
	KeyKPEqual
	KeyKPMultiply
	KeyKPSubtract
	KeyLeft
	KeyLeftBracket
	KeyMenu
	KeyMinus
	KeyNumLock
	KeyPageDown
	KeyPageUp
	KeyPause
	KeyPeriod
	KeyPrintScreen
	KeyRight
	KeyRightBracket
	KeyScrollLock
	KeySemicolon
	KeySlash
	KeySpace
	KeyTab
	KeyUp
	KeyAlt
	KeyControl
	KeyShift
)

type KeyMod int

const (
	KeyModAlt KeyMod = 1 << iota
	KeyModControl
	KeyModShift
)

type MouseButton int

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonMiddle
	MouseButtonRight
)

type MouseButtonMod int

const (
	MouseButtonModLeft MouseButtonMod = 1 << iota
	MouseButtonModMiddle
	MouseButtonModRight
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

type MouseButtonUpHandler interface {
	OnMouseButtonUp(event MouseEvent) bool
}

type MouseMoveHandler interface {
	OnMouseMove(event MouseMoveEvent) bool
}

var singleton inputManager

func Initialize(backend InputBackend) error {
	singleton.key = make(map[Key]int)
	singleton.mouseButton = make(map[MouseButton]int)

	if backend == Ebiten {
		if err := ebitenInput(&singleton); err != nil {
			return err
		}
	}
	return nil
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
