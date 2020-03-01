package d2input

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/d2input_ebiten"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/keyboard"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/mouse"
)

var (
	ErrHasReg         = errors.New("input system already has provided handler")
	ErrNotReg         = errors.New("input system does not have provided handler")
	ErrInvalidBackend = errors.New("invalid input system backend specified")
)

type BackendType int

const (
	Ebiten = iota
)

type Priority int

const (
	PriorityLow Priority = iota
	PriorityDefault
	PriorityHigh
)

type KeyMod int

const (
	KeyModAlt = KeyMod(1 << iota)
	KeyModControl
	KeyModShift
)

type MouseButtonMod int

const (
	MouseButtonModLeft = MouseButtonMod(1 << iota)
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
	Key keyboard.Key
}

type KeyCharsEvent struct {
	HandlerEvent
	Chars []rune
}

type MouseEvent struct {
	HandlerEvent
	Button mouse.MouseButton
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

func Initialize(t BackendType) error {
	switch t {
	case Ebiten:
		singleton.backend = &d2input_ebiten.Backend{}
	default:
		return ErrInvalidBackend
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

const (
	MouseButtonLeft   = mouse.ButtonLeft
	MouseButtonMiddle = mouse.ButtonMiddle
	MouseButtonRight  = mouse.ButtonRight
	Key0              = keyboard.Key0
	Key1              = keyboard.Key1
	Key2              = keyboard.Key2
	Key3              = keyboard.Key3
	Key4              = keyboard.Key4
	Key5              = keyboard.Key5
	Key6              = keyboard.Key6
	Key7              = keyboard.Key7
	Key8              = keyboard.Key8
	Key9              = keyboard.Key9
	KeyA              = keyboard.KeyA
	KeyB              = keyboard.KeyB
	KeyC              = keyboard.KeyC
	KeyD              = keyboard.KeyD
	KeyE              = keyboard.KeyE
	KeyF              = keyboard.KeyF
	KeyG              = keyboard.KeyG
	KeyH              = keyboard.KeyH
	KeyI              = keyboard.KeyI
	KeyJ              = keyboard.KeyJ
	KeyK              = keyboard.KeyK
	KeyL              = keyboard.KeyL
	KeyM              = keyboard.KeyM
	KeyN              = keyboard.KeyN
	KeyO              = keyboard.KeyO
	KeyP              = keyboard.KeyP
	KeyQ              = keyboard.KeyQ
	KeyR              = keyboard.KeyR
	KeyS              = keyboard.KeyS
	KeyT              = keyboard.KeyT
	KeyU              = keyboard.KeyU
	KeyV              = keyboard.KeyV
	KeyW              = keyboard.KeyW
	KeyX              = keyboard.KeyX
	KeyY              = keyboard.KeyY
	KeyZ              = keyboard.KeyZ
	KeyApostrophe     = keyboard.KeyApostrophe
	KeyBackslash      = keyboard.KeyBackslash
	KeyBackspace      = keyboard.KeyBackspace
	KeyCapsLock       = keyboard.KeyCapsLock
	KeyComma          = keyboard.KeyComma
	KeyDelete         = keyboard.KeyDelete
	KeyDown           = keyboard.KeyDown
	KeyEnd            = keyboard.KeyEnd
	KeyEnter          = keyboard.KeyEnter
	KeyEqual          = keyboard.KeyEqual
	KeyEscape         = keyboard.KeyEscape
	KeyF1             = keyboard.KeyF1
	KeyF2             = keyboard.KeyF2
	KeyF3             = keyboard.KeyF3
	KeyF4             = keyboard.KeyF4
	KeyF5             = keyboard.KeyF5
	KeyF6             = keyboard.KeyF6
	KeyF7             = keyboard.KeyF7
	KeyF8             = keyboard.KeyF8
	KeyF9             = keyboard.KeyF9
	KeyF10            = keyboard.KeyF10
	KeyF11            = keyboard.KeyF11
	KeyF12            = keyboard.KeyF12
	KeyGraveAccent    = keyboard.KeyGraveAccent
	KeyHome           = keyboard.KeyHome
	KeyInsert         = keyboard.KeyInsert
	KeyKP0            = keyboard.KeyKP0
	KeyKP1            = keyboard.KeyKP1
	KeyKP2            = keyboard.KeyKP2
	KeyKP3            = keyboard.KeyKP3
	KeyKP4            = keyboard.KeyKP4
	KeyKP5            = keyboard.KeyKP5
	KeyKP6            = keyboard.KeyKP6
	KeyKP7            = keyboard.KeyKP7
	KeyKP8            = keyboard.KeyKP8
	KeyKP9            = keyboard.KeyKP9
	KeyKPAdd          = keyboard.KeyKPAdd
	KeyKPDecimal      = keyboard.KeyKPDecimal
	KeyKPDivide       = keyboard.KeyKPDivide
	KeyKPEnter        = keyboard.KeyKPEnter
	KeyKPEqual        = keyboard.KeyKPEqual
	KeyKPMultiply     = keyboard.KeyKPMultiply
	KeyKPSubtract     = keyboard.KeyKPSubtract
	KeyLeft           = keyboard.KeyLeft
	KeyLeftBracket    = keyboard.KeyLeftBracket
	KeyMenu           = keyboard.KeyMenu
	KeyMinus          = keyboard.KeyMinus
	KeyNumLock        = keyboard.KeyNumLock
	KeyPageDown       = keyboard.KeyPageDown
	KeyPageUp         = keyboard.KeyPageUp
	KeyPause          = keyboard.KeyPause
	KeyPeriod         = keyboard.KeyPeriod
	KeyPrintScreen    = keyboard.KeyPrintScreen
	KeyRight          = keyboard.KeyRight
	KeyRightBracket   = keyboard.KeyRightBracket
	KeyScrollLock     = keyboard.KeyScrollLock
	KeySemicolon      = keyboard.KeySemicolon
	KeySlash          = keyboard.KeySlash
	KeySpace          = keyboard.KeySpace
	KeyTab            = keyboard.KeyTab
	KeyUp             = keyboard.KeyUp
	KeyAlt            = keyboard.KeyAlt
	KeyControl        = keyboard.KeyControl
	KeyShift          = keyboard.KeyShift
)
