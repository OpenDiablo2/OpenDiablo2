// Package ebiten provides graphics and input API to develop a 2D game.
package ebiten

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var (
	//nolint:gochecknoglobals This is a constant in all but by name, no constant map in go
	keyToEbiten = map[d2interface.Key]ebiten.Key{
		d2interface.Key0:            ebiten.Key0,
		d2interface.Key1:            ebiten.Key1,
		d2interface.Key2:            ebiten.Key2,
		d2interface.Key3:            ebiten.Key3,
		d2interface.Key4:            ebiten.Key4,
		d2interface.Key5:            ebiten.Key5,
		d2interface.Key6:            ebiten.Key6,
		d2interface.Key7:            ebiten.Key7,
		d2interface.Key8:            ebiten.Key8,
		d2interface.Key9:            ebiten.Key9,
		d2interface.KeyA:            ebiten.KeyA,
		d2interface.KeyB:            ebiten.KeyB,
		d2interface.KeyC:            ebiten.KeyC,
		d2interface.KeyD:            ebiten.KeyD,
		d2interface.KeyE:            ebiten.KeyE,
		d2interface.KeyF:            ebiten.KeyF,
		d2interface.KeyG:            ebiten.KeyG,
		d2interface.KeyH:            ebiten.KeyH,
		d2interface.KeyI:            ebiten.KeyI,
		d2interface.KeyJ:            ebiten.KeyJ,
		d2interface.KeyK:            ebiten.KeyK,
		d2interface.KeyL:            ebiten.KeyL,
		d2interface.KeyM:            ebiten.KeyM,
		d2interface.KeyN:            ebiten.KeyN,
		d2interface.KeyO:            ebiten.KeyO,
		d2interface.KeyP:            ebiten.KeyP,
		d2interface.KeyQ:            ebiten.KeyQ,
		d2interface.KeyR:            ebiten.KeyR,
		d2interface.KeyS:            ebiten.KeyS,
		d2interface.KeyT:            ebiten.KeyT,
		d2interface.KeyU:            ebiten.KeyU,
		d2interface.KeyV:            ebiten.KeyV,
		d2interface.KeyW:            ebiten.KeyW,
		d2interface.KeyX:            ebiten.KeyX,
		d2interface.KeyY:            ebiten.KeyY,
		d2interface.KeyZ:            ebiten.KeyZ,
		d2interface.KeyApostrophe:   ebiten.KeyApostrophe,
		d2interface.KeyBackslash:    ebiten.KeyBackslash,
		d2interface.KeyBackspace:    ebiten.KeyBackspace,
		d2interface.KeyCapsLock:     ebiten.KeyCapsLock,
		d2interface.KeyComma:        ebiten.KeyComma,
		d2interface.KeyDelete:       ebiten.KeyDelete,
		d2interface.KeyDown:         ebiten.KeyDown,
		d2interface.KeyEnd:          ebiten.KeyEnd,
		d2interface.KeyEnter:        ebiten.KeyEnter,
		d2interface.KeyEqual:        ebiten.KeyEqual,
		d2interface.KeyEscape:       ebiten.KeyEscape,
		d2interface.KeyF1:           ebiten.KeyF1,
		d2interface.KeyF2:           ebiten.KeyF2,
		d2interface.KeyF3:           ebiten.KeyF3,
		d2interface.KeyF4:           ebiten.KeyF4,
		d2interface.KeyF5:           ebiten.KeyF5,
		d2interface.KeyF6:           ebiten.KeyF6,
		d2interface.KeyF7:           ebiten.KeyF7,
		d2interface.KeyF8:           ebiten.KeyF8,
		d2interface.KeyF9:           ebiten.KeyF9,
		d2interface.KeyF10:          ebiten.KeyF10,
		d2interface.KeyF11:          ebiten.KeyF11,
		d2interface.KeyF12:          ebiten.KeyF12,
		d2interface.KeyGraveAccent:  ebiten.KeyGraveAccent,
		d2interface.KeyHome:         ebiten.KeyHome,
		d2interface.KeyInsert:       ebiten.KeyInsert,
		d2interface.KeyKP0:          ebiten.KeyKP0,
		d2interface.KeyKP1:          ebiten.KeyKP1,
		d2interface.KeyKP2:          ebiten.KeyKP2,
		d2interface.KeyKP3:          ebiten.KeyKP3,
		d2interface.KeyKP4:          ebiten.KeyKP4,
		d2interface.KeyKP5:          ebiten.KeyKP5,
		d2interface.KeyKP6:          ebiten.KeyKP6,
		d2interface.KeyKP7:          ebiten.KeyKP7,
		d2interface.KeyKP8:          ebiten.KeyKP8,
		d2interface.KeyKP9:          ebiten.KeyKP9,
		d2interface.KeyKPAdd:        ebiten.KeyKPAdd,
		d2interface.KeyKPDecimal:    ebiten.KeyKPDecimal,
		d2interface.KeyKPDivide:     ebiten.KeyKPDivide,
		d2interface.KeyKPEnter:      ebiten.KeyKPEnter,
		d2interface.KeyKPEqual:      ebiten.KeyKPEqual,
		d2interface.KeyKPMultiply:   ebiten.KeyKPMultiply,
		d2interface.KeyKPSubtract:   ebiten.KeyKPSubtract,
		d2interface.KeyLeft:         ebiten.KeyLeft,
		d2interface.KeyLeftBracket:  ebiten.KeyLeftBracket,
		d2interface.KeyMenu:         ebiten.KeyMenu,
		d2interface.KeyMinus:        ebiten.KeyMinus,
		d2interface.KeyNumLock:      ebiten.KeyNumLock,
		d2interface.KeyPageDown:     ebiten.KeyPageDown,
		d2interface.KeyPageUp:       ebiten.KeyPageUp,
		d2interface.KeyPause:        ebiten.KeyPause,
		d2interface.KeyPeriod:       ebiten.KeyPeriod,
		d2interface.KeyPrintScreen:  ebiten.KeyPrintScreen,
		d2interface.KeyRight:        ebiten.KeyRight,
		d2interface.KeyRightBracket: ebiten.KeyRightBracket,
		d2interface.KeyScrollLock:   ebiten.KeyScrollLock,
		d2interface.KeySemicolon:    ebiten.KeySemicolon,
		d2interface.KeySlash:        ebiten.KeySlash,
		d2interface.KeySpace:        ebiten.KeySpace,
		d2interface.KeyTab:          ebiten.KeyTab,
		d2interface.KeyUp:           ebiten.KeyUp,
		d2interface.KeyAlt:          ebiten.KeyAlt,
		d2interface.KeyControl:      ebiten.KeyControl,
		d2interface.KeyShift:        ebiten.KeyShift,
	}
	//nolint:gochecknoglobals This is a constant in all but by name, no constant map in go
	mouseButtonToEbiten = map[d2interface.MouseButton]ebiten.MouseButton{
		d2interface.MouseButtonLeft:   ebiten.MouseButtonLeft,
		d2interface.MouseButtonMiddle: ebiten.MouseButtonMiddle,
		d2interface.MouseButtonRight:  ebiten.MouseButtonRight,
	}
)

// InputService provides an abstraction on ebiten to support handling input events
type InputService struct{}

// CursorPosition returns a position of a mouse cursor relative to the game screen (window).
func (is InputService) CursorPosition() (x, y int) {
	return ebiten.CursorPosition()
}

// InputChars return "printable" runes read from the keyboard at the time update is called.
func (is InputService) InputChars() []rune {
	return ebiten.InputChars()
}

// IsKeyPressed checks if the provided key is down.
func (is InputService) IsKeyPressed(key d2interface.Key) bool {
	return ebiten.IsKeyPressed(keyToEbiten[key])
}

// IsKeyJustPressed checks if the provided key is just transitioned from up to down.
func (is InputService) IsKeyJustPressed(key d2interface.Key) bool {
	return inpututil.IsKeyJustPressed(keyToEbiten[key])
}

// IsKeyJustReleased checks if the provided key is just transitioned from down to up.
func (is InputService) IsKeyJustReleased(key d2interface.Key) bool {
	return inpututil.IsKeyJustReleased(keyToEbiten[key])
}

// IsMouseButtonPressed checks if the provided mouse button is down.
func (is InputService) IsMouseButtonPressed(button d2interface.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(mouseButtonToEbiten[button])
}

// IsMouseButtonJustPressed checks if the provided mouse button is just transitioned from up to down.
func (is InputService) IsMouseButtonJustPressed(button d2interface.MouseButton) bool {
	return inpututil.IsMouseButtonJustPressed(mouseButtonToEbiten[button])
}

// IsMouseButtonJustReleased checks if the provided mouse button is just transitioned from down to up.
func (is InputService) IsMouseButtonJustReleased(button d2interface.MouseButton) bool {
	return inpututil.IsMouseButtonJustReleased(mouseButtonToEbiten[button])
}

// KeyPressDuration returns how long the key is pressed in frames.
func (is InputService) KeyPressDuration(key d2interface.Key) int {
	return inpututil.KeyPressDuration(keyToEbiten[key])
}
