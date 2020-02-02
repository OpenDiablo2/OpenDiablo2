package d2input

import (
	"errors"

	"github.com/hajimehoshi/ebiten"
)

var (
	ErrHasInit = errors.New("input system is already initialized")
	ErrNotInit = errors.New("input system is not initialized")
	ErrHasReg  = errors.New("input system already has provided handler")
	ErrNotReg  = errors.New("input system does not have provided handler")
)

type Priority int

const (
	PriorityLow Priority = iota
	PriorityDefault
	PriorityHigh
)

type Key int

//noinspection GoUnusedConst
const (
	Key0 = Key(ebiten.Key0)
	Key1 = Key(ebiten.Key1)
	Key2 = Key(ebiten.Key2)
	Key3 = Key(ebiten.Key3)
	Key4 = Key(ebiten.Key4)
	Key5 = Key(ebiten.Key5)
	Key6 = Key(ebiten.Key6)
	Key7 = Key(ebiten.Key7)
	Key8 = Key(ebiten.Key8)
	Key9 = Key(ebiten.Key9)
	KeyA = Key(ebiten.KeyA)
	KeyB     = Key(ebiten.KeyB)
	KeyC     = Key(ebiten.KeyC)
	KeyD     = Key(ebiten.KeyD)
	KeyE     = Key(ebiten.KeyE)
	KeyF     = Key(ebiten.KeyF)
	KeyG     = Key(ebiten.KeyG)
	KeyH     = Key(ebiten.KeyH)
	KeyI     = Key(ebiten.KeyI)
	KeyJ     = Key(ebiten.KeyJ)
	KeyK     = Key(ebiten.KeyK)
	KeyL     = Key(ebiten.KeyL)
	KeyM     = Key(ebiten.KeyM)
	KeyN     = Key(ebiten.KeyN)
	KeyO     = Key(ebiten.KeyO)
	KeyP     = Key(ebiten.KeyP)
	KeyQ     = Key(ebiten.KeyQ)
	KeyR     = Key(ebiten.KeyR)
	KeyS     = Key(ebiten.KeyS)
	KeyT     = Key(ebiten.KeyT)
	KeyU     = Key(ebiten.KeyU)
	KeyV     = Key(ebiten.KeyV)
	KeyW     = Key(ebiten.KeyW)
	KeyX     = Key(ebiten.KeyX)
	KeyY     = Key(ebiten.KeyY)
	KeyZ     = Key(ebiten.KeyZ)
	KeyApostrophe     = Key(ebiten.KeyApostrophe)
	KeyBackslash      = Key(ebiten.KeyBackslash)
	KeyBackspace      = Key(ebiten.KeyBackspace)
	KeyCapsLock       = Key(ebiten.KeyCapsLock)
	KeyComma          = Key(ebiten.KeyComma)
	KeyDelete         = Key(ebiten.KeyDelete)
	KeyDown           = Key(ebiten.KeyDown)
	KeyEnd            = Key(ebiten.KeyEnd)
	KeyEnter          = Key(ebiten.KeyEnter)
	KeyEqual          = Key(ebiten.KeyEqual)
	KeyEscape         = Key(ebiten.KeyEscape)
	KeyF1             = Key(ebiten.KeyF1)
	KeyF2             = Key(ebiten.KeyF2)
	KeyF3             = Key(ebiten.KeyF3)
	KeyF4             = Key(ebiten.KeyF4)
	KeyF5             = Key(ebiten.KeyF5)
	KeyF6             = Key(ebiten.KeyF6)
	KeyF7             = Key(ebiten.KeyF7)
	KeyF8             = Key(ebiten.KeyF8)
	KeyF9             = Key(ebiten.KeyF9)
	KeyF10            = Key(ebiten.KeyF10)
	KeyF11           = Key(ebiten.KeyF11)
	KeyF12           = Key(ebiten.KeyF12)
	KeyGraveAccent     = Key(ebiten.KeyGraveAccent)
	KeyHome            = Key(ebiten.KeyHome)
	KeyInsert          = Key(ebiten.KeyInsert)
	KeyKP0             = Key(ebiten.KeyKP0)
	KeyKP1             = Key(ebiten.KeyKP1)
	KeyKP2             = Key(ebiten.KeyKP2)
	KeyKP3             = Key(ebiten.KeyKP3)
	KeyKP4             = Key(ebiten.KeyKP4)
	KeyKP5             = Key(ebiten.KeyKP5)
	KeyKP6             = Key(ebiten.KeyKP6)
	KeyKP7             = Key(ebiten.KeyKP7)
	KeyKP8             = Key(ebiten.KeyKP8)
	KeyKP9             = Key(ebiten.KeyKP9)
	KeyKPAdd           = Key(ebiten.KeyKPAdd)
	KeyKPDecimal       = Key(ebiten.KeyKPDecimal)
	KeyKPDivide        = Key(ebiten.KeyKPDivide)
	KeyKPEnter         = Key(ebiten.KeyKPEnter)
	KeyKPEqual         = Key(ebiten.KeyKPEqual)
	KeyKPMultiply      = Key(ebiten.KeyKPMultiply)
	KeyKPSubtract      = Key(ebiten.KeyKPSubtract)
	KeyLeft            = Key(ebiten.KeyLeft)
	KeyLeftBracket     = Key(ebiten.KeyLeftBracket)
	KeyMenu            = Key(ebiten.KeyMenu)
	KeyMinus           = Key(ebiten.KeyMinus)
	KeyNumLock         = Key(ebiten.KeyNumLock)
	KeyPageDown        = Key(ebiten.KeyPageDown)
	KeyPageUp          = Key(ebiten.KeyPageUp)
	KeyPause           = Key(ebiten.KeyPause)
	KeyPeriod          = Key(ebiten.KeyPeriod)
	KeyPrintScreen     = Key(ebiten.KeyPrintScreen)
	KeyRight           = Key(ebiten.KeyRight)
	KeyRightBracket     = Key(ebiten.KeyRightBracket)
	KeyScrollLock       = Key(ebiten.KeyScrollLock)
	KeySemicolon        = Key(ebiten.KeySemicolon)
	KeySlash            = Key(ebiten.KeySlash)
	KeySpace            = Key(ebiten.KeySpace)
	KeyTab              = Key(ebiten.KeyTab)
	KeyUp               = Key(ebiten.KeyUp)
	KeyAlt              = Key(ebiten.KeyAlt)
	KeyControl          = Key(ebiten.KeyControl)
	KeyShift            = Key(ebiten.KeyShift)
)

type KeyMod int

const (
	KeyModAlt = 1 << iota
	KeyModControl
	KeyModShift
)

type MouseButton int

const (
	MouseButtonLeft   = MouseButton(ebiten.MouseButtonLeft)
	MouseButtonMiddle = MouseButton(ebiten.MouseButtonMiddle)
	MouseButtonRight  = MouseButton(ebiten.MouseButtonRight)
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

var singleton *inputManager

func Initialize() error {
	if singleton != nil {
		return ErrHasInit
	}

	singleton = &inputManager{}
	return nil
}

func Shutdown() {
	singleton = nil
}

func Advance(elapsed float64) error {
	if singleton == nil {
		return ErrNotInit
	}

	return singleton.advance(elapsed)
}

func BindHandlerWithPriority(handler Handler, priority Priority) error {
	if singleton == nil {
		return ErrNotInit
	}

	return singleton.bindHandler(handler, priority)
}

func BindHandler(handler Handler) error {
	return BindHandlerWithPriority(handler, PriorityDefault)
}

func UnbindHandler(handler Handler) error {
	if singleton == nil {
		return ErrNotInit
	}

	return singleton.unbindHandler(handler)
}
