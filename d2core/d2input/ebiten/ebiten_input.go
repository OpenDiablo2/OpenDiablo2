package ebiten

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var (
	keyToEbiten = map[d2input.Key]ebiten.Key{
		d2input.Key0:            ebiten.Key0,
		d2input.Key1:            ebiten.Key1,
		d2input.Key2:            ebiten.Key2,
		d2input.Key3:            ebiten.Key3,
		d2input.Key4:            ebiten.Key4,
		d2input.Key5:            ebiten.Key5,
		d2input.Key6:            ebiten.Key6,
		d2input.Key7:            ebiten.Key7,
		d2input.Key8:            ebiten.Key8,
		d2input.Key9:            ebiten.Key9,
		d2input.KeyA:            ebiten.KeyA,
		d2input.KeyB:            ebiten.KeyB,
		d2input.KeyC:            ebiten.KeyC,
		d2input.KeyD:            ebiten.KeyD,
		d2input.KeyE:            ebiten.KeyE,
		d2input.KeyF:            ebiten.KeyF,
		d2input.KeyG:            ebiten.KeyG,
		d2input.KeyH:            ebiten.KeyH,
		d2input.KeyI:            ebiten.KeyI,
		d2input.KeyJ:            ebiten.KeyJ,
		d2input.KeyK:            ebiten.KeyK,
		d2input.KeyL:            ebiten.KeyL,
		d2input.KeyM:            ebiten.KeyM,
		d2input.KeyN:            ebiten.KeyN,
		d2input.KeyO:            ebiten.KeyO,
		d2input.KeyP:            ebiten.KeyP,
		d2input.KeyQ:            ebiten.KeyQ,
		d2input.KeyR:            ebiten.KeyR,
		d2input.KeyS:            ebiten.KeyS,
		d2input.KeyT:            ebiten.KeyT,
		d2input.KeyU:            ebiten.KeyU,
		d2input.KeyV:            ebiten.KeyV,
		d2input.KeyW:            ebiten.KeyW,
		d2input.KeyX:            ebiten.KeyX,
		d2input.KeyY:            ebiten.KeyY,
		d2input.KeyZ:            ebiten.KeyZ,
		d2input.KeyApostrophe:   ebiten.KeyApostrophe,
		d2input.KeyBackslash:    ebiten.KeyBackslash,
		d2input.KeyBackspace:    ebiten.KeyBackspace,
		d2input.KeyCapsLock:     ebiten.KeyCapsLock,
		d2input.KeyComma:        ebiten.KeyComma,
		d2input.KeyDelete:       ebiten.KeyDelete,
		d2input.KeyDown:         ebiten.KeyDown,
		d2input.KeyEnd:          ebiten.KeyEnd,
		d2input.KeyEnter:        ebiten.KeyEnter,
		d2input.KeyEqual:        ebiten.KeyEqual,
		d2input.KeyEscape:       ebiten.KeyEscape,
		d2input.KeyF1:           ebiten.KeyF1,
		d2input.KeyF2:           ebiten.KeyF2,
		d2input.KeyF3:           ebiten.KeyF3,
		d2input.KeyF4:           ebiten.KeyF4,
		d2input.KeyF5:           ebiten.KeyF5,
		d2input.KeyF6:           ebiten.KeyF6,
		d2input.KeyF7:           ebiten.KeyF7,
		d2input.KeyF8:           ebiten.KeyF8,
		d2input.KeyF9:           ebiten.KeyF9,
		d2input.KeyF10:          ebiten.KeyF10,
		d2input.KeyF11:          ebiten.KeyF11,
		d2input.KeyF12:          ebiten.KeyF12,
		d2input.KeyGraveAccent:  ebiten.KeyGraveAccent,
		d2input.KeyHome:         ebiten.KeyHome,
		d2input.KeyInsert:       ebiten.KeyInsert,
		d2input.KeyKP0:          ebiten.KeyKP0,
		d2input.KeyKP1:          ebiten.KeyKP1,
		d2input.KeyKP2:          ebiten.KeyKP2,
		d2input.KeyKP3:          ebiten.KeyKP3,
		d2input.KeyKP4:          ebiten.KeyKP4,
		d2input.KeyKP5:          ebiten.KeyKP5,
		d2input.KeyKP6:          ebiten.KeyKP6,
		d2input.KeyKP7:          ebiten.KeyKP7,
		d2input.KeyKP8:          ebiten.KeyKP8,
		d2input.KeyKP9:          ebiten.KeyKP9,
		d2input.KeyKPAdd:        ebiten.KeyKPAdd,
		d2input.KeyKPDecimal:    ebiten.KeyKPDecimal,
		d2input.KeyKPDivide:     ebiten.KeyKPDivide,
		d2input.KeyKPEnter:      ebiten.KeyKPEnter,
		d2input.KeyKPEqual:      ebiten.KeyKPEqual,
		d2input.KeyKPMultiply:   ebiten.KeyKPMultiply,
		d2input.KeyKPSubtract:   ebiten.KeyKPSubtract,
		d2input.KeyLeft:         ebiten.KeyLeft,
		d2input.KeyLeftBracket:  ebiten.KeyLeftBracket,
		d2input.KeyMenu:         ebiten.KeyMenu,
		d2input.KeyMinus:        ebiten.KeyMinus,
		d2input.KeyNumLock:      ebiten.KeyNumLock,
		d2input.KeyPageDown:     ebiten.KeyPageDown,
		d2input.KeyPageUp:       ebiten.KeyPageUp,
		d2input.KeyPause:        ebiten.KeyPause,
		d2input.KeyPeriod:       ebiten.KeyPeriod,
		d2input.KeyPrintScreen:  ebiten.KeyPrintScreen,
		d2input.KeyRight:        ebiten.KeyRight,
		d2input.KeyRightBracket: ebiten.KeyRightBracket,
		d2input.KeyScrollLock:   ebiten.KeyScrollLock,
		d2input.KeySemicolon:    ebiten.KeySemicolon,
		d2input.KeySlash:        ebiten.KeySlash,
		d2input.KeySpace:        ebiten.KeySpace,
		d2input.KeyTab:          ebiten.KeyTab,
		d2input.KeyUp:           ebiten.KeyUp,
		d2input.KeyAlt:          ebiten.KeyAlt,
		d2input.KeyControl:      ebiten.KeyControl,
		d2input.KeyShift:        ebiten.KeyShift,
	}
	mouseButtonToEbiten = map[d2input.MouseButton]ebiten.MouseButton{
		d2input.MouseButtonLeft:   ebiten.MouseButtonLeft,
		d2input.MouseButtonMiddle: ebiten.MouseButtonMiddle,
		d2input.MouseButtonRight:  ebiten.MouseButtonRight,
	}
)

// InputService provides an abstraction on ebiten to support handling input events
type InputService struct{}

func (is InputService) CursorPosition() (x int, y int) {
	return ebiten.CursorPosition()
}

func (is InputService) InputChars() []rune {
	return ebiten.InputChars()
}

func (is InputService) IsKeyPressed(key d2input.Key) bool {
	return ebiten.IsKeyPressed(keyToEbiten[key])
}

func (is InputService) IsKeyJustPressed(key d2input.Key) bool {
	return inpututil.IsKeyJustPressed(keyToEbiten[key])
}

func (is InputService) IsKeyJustReleased(key d2input.Key) bool {
	return inpututil.IsKeyJustReleased(keyToEbiten[key])
}

func (is InputService) IsMouseButtonPressed(button d2input.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(mouseButtonToEbiten[button])
}

func (is InputService) IsMouseButtonJustPressed(button d2input.MouseButton) bool {
	return inpututil.IsMouseButtonJustPressed(mouseButtonToEbiten[button])
}

func (is InputService) IsMouseButtonJustReleased(button d2input.MouseButton) bool {
	return inpututil.IsMouseButtonJustReleased(mouseButtonToEbiten[button])
}

func (is InputService) KeyPressDuration(key d2input.Key) int {
	return inpututil.KeyPressDuration(keyToEbiten[key])
}
