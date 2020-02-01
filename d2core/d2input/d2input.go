package d2input

import (
	"errors"

	"github.com/hajimehoshi/ebiten"
)

var (
	ErrHasInit error = errors.New("input system is already initialized")
	ErrNotInit error = errors.New("input system is not initialized")
	ErrHasReg  error = errors.New("input system already has provided handler")
	ErrNotReg  error = errors.New("input system does not have provided handler")
)

type Priority int

const (
	PriorityLow Priority = iota
	PriorityDefault
	PriorityHigh
)

type Key int

const (
	Key0            Key = Key(ebiten.Key0)
	Key1            Key = Key(ebiten.Key1)
	Key2            Key = Key(ebiten.Key2)
	Key3            Key = Key(ebiten.Key3)
	Key4            Key = Key(ebiten.Key4)
	Key5            Key = Key(ebiten.Key5)
	Key6            Key = Key(ebiten.Key6)
	Key7            Key = Key(ebiten.Key7)
	Key8            Key = Key(ebiten.Key8)
	Key9            Key = Key(ebiten.Key9)
	KeyA            Key = Key(ebiten.KeyA)
	KeyB            Key = Key(ebiten.KeyB)
	KeyC            Key = Key(ebiten.KeyC)
	KeyD            Key = Key(ebiten.KeyD)
	KeyE            Key = Key(ebiten.KeyE)
	KeyF            Key = Key(ebiten.KeyF)
	KeyG            Key = Key(ebiten.KeyG)
	KeyH            Key = Key(ebiten.KeyH)
	KeyI            Key = Key(ebiten.KeyI)
	KeyJ            Key = Key(ebiten.KeyJ)
	KeyK            Key = Key(ebiten.KeyK)
	KeyL            Key = Key(ebiten.KeyL)
	KeyM            Key = Key(ebiten.KeyM)
	KeyN            Key = Key(ebiten.KeyN)
	KeyO            Key = Key(ebiten.KeyO)
	KeyP            Key = Key(ebiten.KeyP)
	KeyQ            Key = Key(ebiten.KeyQ)
	KeyR            Key = Key(ebiten.KeyR)
	KeyS            Key = Key(ebiten.KeyS)
	KeyT            Key = Key(ebiten.KeyT)
	KeyU            Key = Key(ebiten.KeyU)
	KeyV            Key = Key(ebiten.KeyV)
	KeyW            Key = Key(ebiten.KeyW)
	KeyX            Key = Key(ebiten.KeyX)
	KeyY            Key = Key(ebiten.KeyY)
	KeyZ            Key = Key(ebiten.KeyZ)
	KeyApostrophe   Key = Key(ebiten.KeyApostrophe)
	KeyBackslash    Key = Key(ebiten.KeyBackslash)
	KeyBackspace    Key = Key(ebiten.KeyBackspace)
	KeyCapsLock     Key = Key(ebiten.KeyCapsLock)
	KeyComma        Key = Key(ebiten.KeyComma)
	KeyDelete       Key = Key(ebiten.KeyDelete)
	KeyDown         Key = Key(ebiten.KeyDown)
	KeyEnd          Key = Key(ebiten.KeyEnd)
	KeyEnter        Key = Key(ebiten.KeyEnter)
	KeyEqual        Key = Key(ebiten.KeyEqual)
	KeyEscape       Key = Key(ebiten.KeyEscape)
	KeyF1           Key = Key(ebiten.KeyF1)
	KeyF2           Key = Key(ebiten.KeyF2)
	KeyF3           Key = Key(ebiten.KeyF3)
	KeyF4           Key = Key(ebiten.KeyF4)
	KeyF5           Key = Key(ebiten.KeyF5)
	KeyF6           Key = Key(ebiten.KeyF6)
	KeyF7           Key = Key(ebiten.KeyF7)
	KeyF8           Key = Key(ebiten.KeyF8)
	KeyF9           Key = Key(ebiten.KeyF9)
	KeyF10          Key = Key(ebiten.KeyF10)
	KeyF11          Key = Key(ebiten.KeyF11)
	KeyF12          Key = Key(ebiten.KeyF12)
	KeyGraveAccent  Key = Key(ebiten.KeyGraveAccent)
	KeyHome         Key = Key(ebiten.KeyHome)
	KeyInsert       Key = Key(ebiten.KeyInsert)
	KeyKP0          Key = Key(ebiten.KeyKP0)
	KeyKP1          Key = Key(ebiten.KeyKP1)
	KeyKP2          Key = Key(ebiten.KeyKP2)
	KeyKP3          Key = Key(ebiten.KeyKP3)
	KeyKP4          Key = Key(ebiten.KeyKP4)
	KeyKP5          Key = Key(ebiten.KeyKP5)
	KeyKP6          Key = Key(ebiten.KeyKP6)
	KeyKP7          Key = Key(ebiten.KeyKP7)
	KeyKP8          Key = Key(ebiten.KeyKP8)
	KeyKP9          Key = Key(ebiten.KeyKP9)
	KeyKPAdd        Key = Key(ebiten.KeyKPAdd)
	KeyKPDecimal    Key = Key(ebiten.KeyKPDecimal)
	KeyKPDivide     Key = Key(ebiten.KeyKPDivide)
	KeyKPEnter      Key = Key(ebiten.KeyKPEnter)
	KeyKPEqual      Key = Key(ebiten.KeyKPEqual)
	KeyKPMultiply   Key = Key(ebiten.KeyKPMultiply)
	KeyKPSubtract   Key = Key(ebiten.KeyKPSubtract)
	KeyLeft         Key = Key(ebiten.KeyLeft)
	KeyLeftBracket  Key = Key(ebiten.KeyLeftBracket)
	KeyMenu         Key = Key(ebiten.KeyMenu)
	KeyMinus        Key = Key(ebiten.KeyMinus)
	KeyNumLock      Key = Key(ebiten.KeyNumLock)
	KeyPageDown     Key = Key(ebiten.KeyPageDown)
	KeyPageUp       Key = Key(ebiten.KeyPageUp)
	KeyPause        Key = Key(ebiten.KeyPause)
	KeyPeriod       Key = Key(ebiten.KeyPeriod)
	KeyPrintScreen  Key = Key(ebiten.KeyPrintScreen)
	KeyRight        Key = Key(ebiten.KeyRight)
	KeyRightBracket Key = Key(ebiten.KeyRightBracket)
	KeyScrollLock   Key = Key(ebiten.KeyScrollLock)
	KeySemicolon    Key = Key(ebiten.KeySemicolon)
	KeySlash        Key = Key(ebiten.KeySlash)
	KeySpace        Key = Key(ebiten.KeySpace)
	KeyTab          Key = Key(ebiten.KeyTab)
	KeyUp           Key = Key(ebiten.KeyUp)
	KeyAlt          Key = Key(ebiten.KeyAlt)
	KeyControl      Key = Key(ebiten.KeyControl)
	KeyShift        Key = Key(ebiten.KeyShift)
)

type KeyMod int

const (
	KeyModAlt = 1 << iota
	KeyModControl
	KeyModShift
)

type MouseButton int

const (
	MouseButtonLeft   MouseButton = MouseButton(ebiten.MouseButtonLeft)
	MouseButtonMiddle MouseButton = MouseButton(ebiten.MouseButtonMiddle)
	MouseButtonRight  MouseButton = MouseButton(ebiten.MouseButtonRight)
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
