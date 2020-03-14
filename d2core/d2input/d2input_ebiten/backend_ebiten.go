package d2input_ebiten

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/keyboard"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/mouse"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Backend struct{}

var (
	keyMap    map[keyboard.Key]ebiten.Key
	buttonMap map[mouse.MouseButton]ebiten.MouseButton
)

func (b *Backend) Initialize() error {
	setupKeyMap()
	setupButtonMap()
	return nil
}

func (b *Backend) Advance(elapsed float64) error {
	// ebiten doesnt need to do anything here
	return nil
}

func (b *Backend) CursorPosition() (int, int) {
	return ebiten.CursorPosition()
}

func (b *Backend) IsKeyPressed(k keyboard.Key) bool {
	return ebiten.IsKeyPressed(keyMap[k])
}

func (b *Backend) IsKeyJustPressed(k keyboard.Key) bool {
	return inpututil.IsKeyJustPressed(keyMap[k])
}

func (b *Backend) IsKeyJustReleased(k keyboard.Key) bool {
	return inpututil.IsKeyJustReleased(keyMap[k])
}

func (b *Backend) IsMouseButtonPressed(mb mouse.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(buttonMap[mb])
}

func (b *Backend) IsMouseButtonJustPressed(mb mouse.MouseButton) bool {
	return inpututil.IsMouseButtonJustPressed(buttonMap[mb])
}

func (b *Backend) IsMouseButtonJustReleased(mb mouse.MouseButton) bool {
	return inpututil.IsMouseButtonJustReleased(buttonMap[mb])
}

func (b *Backend) InputChars() []rune {
	return ebiten.InputChars()
}

func setupKeyMap() {
	keyMap = make(map[keyboard.Key]ebiten.Key)
	keyMap[keyboard.Key0] = ebiten.Key0
	keyMap[keyboard.Key1] = ebiten.Key1
	keyMap[keyboard.Key2] = ebiten.Key2
	keyMap[keyboard.Key3] = ebiten.Key3
	keyMap[keyboard.Key4] = ebiten.Key4
	keyMap[keyboard.Key5] = ebiten.Key5
	keyMap[keyboard.Key6] = ebiten.Key6
	keyMap[keyboard.Key7] = ebiten.Key7
	keyMap[keyboard.Key8] = ebiten.Key8
	keyMap[keyboard.Key9] = ebiten.Key9
	keyMap[keyboard.KeyA] = ebiten.KeyA
	keyMap[keyboard.KeyB] = ebiten.KeyB
	keyMap[keyboard.KeyC] = ebiten.KeyC
	keyMap[keyboard.KeyD] = ebiten.KeyD
	keyMap[keyboard.KeyE] = ebiten.KeyE
	keyMap[keyboard.KeyF] = ebiten.KeyF
	keyMap[keyboard.KeyG] = ebiten.KeyG
	keyMap[keyboard.KeyH] = ebiten.KeyH
	keyMap[keyboard.KeyI] = ebiten.KeyI
	keyMap[keyboard.KeyJ] = ebiten.KeyJ
	keyMap[keyboard.KeyK] = ebiten.KeyK
	keyMap[keyboard.KeyL] = ebiten.KeyL
	keyMap[keyboard.KeyM] = ebiten.KeyM
	keyMap[keyboard.KeyN] = ebiten.KeyN
	keyMap[keyboard.KeyO] = ebiten.KeyO
	keyMap[keyboard.KeyP] = ebiten.KeyP
	keyMap[keyboard.KeyQ] = ebiten.KeyQ
	keyMap[keyboard.KeyR] = ebiten.KeyR
	keyMap[keyboard.KeyS] = ebiten.KeyS
	keyMap[keyboard.KeyT] = ebiten.KeyT
	keyMap[keyboard.KeyU] = ebiten.KeyU
	keyMap[keyboard.KeyV] = ebiten.KeyV
	keyMap[keyboard.KeyW] = ebiten.KeyW
	keyMap[keyboard.KeyX] = ebiten.KeyX
	keyMap[keyboard.KeyY] = ebiten.KeyY
	keyMap[keyboard.KeyZ] = ebiten.KeyZ
	keyMap[keyboard.KeyApostrophe] = ebiten.KeyApostrophe
	keyMap[keyboard.KeyBackslash] = ebiten.KeyBackslash
	keyMap[keyboard.KeyBackspace] = ebiten.KeyBackspace
	keyMap[keyboard.KeyCapsLock] = ebiten.KeyCapsLock
	keyMap[keyboard.KeyComma] = ebiten.KeyComma
	keyMap[keyboard.KeyDelete] = ebiten.KeyDelete
	keyMap[keyboard.KeyDown] = ebiten.KeyDown
	keyMap[keyboard.KeyEnd] = ebiten.KeyEnd
	keyMap[keyboard.KeyEnter] = ebiten.KeyEnter
	keyMap[keyboard.KeyEqual] = ebiten.KeyEqual
	keyMap[keyboard.KeyEscape] = ebiten.KeyEscape
	keyMap[keyboard.KeyF1] = ebiten.KeyF1
	keyMap[keyboard.KeyF2] = ebiten.KeyF2
	keyMap[keyboard.KeyF3] = ebiten.KeyF3
	keyMap[keyboard.KeyF4] = ebiten.KeyF4
	keyMap[keyboard.KeyF5] = ebiten.KeyF5
	keyMap[keyboard.KeyF6] = ebiten.KeyF6
	keyMap[keyboard.KeyF7] = ebiten.KeyF7
	keyMap[keyboard.KeyF8] = ebiten.KeyF8
	keyMap[keyboard.KeyF9] = ebiten.KeyF9
	keyMap[keyboard.KeyF10] = ebiten.KeyF10
	keyMap[keyboard.KeyF11] = ebiten.KeyF11
	keyMap[keyboard.KeyF12] = ebiten.KeyF12
	keyMap[keyboard.KeyGraveAccent] = ebiten.KeyGraveAccent
	keyMap[keyboard.KeyHome] = ebiten.KeyHome
	keyMap[keyboard.KeyInsert] = ebiten.KeyInsert
	keyMap[keyboard.KeyKP0] = ebiten.KeyKP0
	keyMap[keyboard.KeyKP1] = ebiten.KeyKP1
	keyMap[keyboard.KeyKP2] = ebiten.KeyKP2
	keyMap[keyboard.KeyKP3] = ebiten.KeyKP3
	keyMap[keyboard.KeyKP4] = ebiten.KeyKP4
	keyMap[keyboard.KeyKP5] = ebiten.KeyKP5
	keyMap[keyboard.KeyKP6] = ebiten.KeyKP6
	keyMap[keyboard.KeyKP7] = ebiten.KeyKP7
	keyMap[keyboard.KeyKP8] = ebiten.KeyKP8
	keyMap[keyboard.KeyKP9] = ebiten.KeyKP9
	keyMap[keyboard.KeyKPAdd] = ebiten.KeyKPAdd
	keyMap[keyboard.KeyKPDecimal] = ebiten.KeyKPDecimal
	keyMap[keyboard.KeyKPDivide] = ebiten.KeyKPDivide
	keyMap[keyboard.KeyKPEnter] = ebiten.KeyKPEnter
	keyMap[keyboard.KeyKPEqual] = ebiten.KeyKPEqual
	keyMap[keyboard.KeyKPMultiply] = ebiten.KeyKPMultiply
	keyMap[keyboard.KeyKPSubtract] = ebiten.KeyKPSubtract
	keyMap[keyboard.KeyLeft] = ebiten.KeyLeft
	keyMap[keyboard.KeyLeftBracket] = ebiten.KeyLeftBracket
	keyMap[keyboard.KeyMenu] = ebiten.KeyMenu
	keyMap[keyboard.KeyMinus] = ebiten.KeyMinus
	keyMap[keyboard.KeyNumLock] = ebiten.KeyNumLock
	keyMap[keyboard.KeyPageDown] = ebiten.KeyPageDown
	keyMap[keyboard.KeyPageUp] = ebiten.KeyPageUp
	keyMap[keyboard.KeyPause] = ebiten.KeyPause
	keyMap[keyboard.KeyPeriod] = ebiten.KeyPeriod
	keyMap[keyboard.KeyPrintScreen] = ebiten.KeyPrintScreen
	keyMap[keyboard.KeyRight] = ebiten.KeyRight
	keyMap[keyboard.KeyRightBracket] = ebiten.KeyRightBracket
	keyMap[keyboard.KeyScrollLock] = ebiten.KeyScrollLock
	keyMap[keyboard.KeySemicolon] = ebiten.KeySemicolon
	keyMap[keyboard.KeySlash] = ebiten.KeySlash
	keyMap[keyboard.KeySpace] = ebiten.KeySpace
	keyMap[keyboard.KeyTab] = ebiten.KeyTab
	keyMap[keyboard.KeyUp] = ebiten.KeyUp
	keyMap[keyboard.KeyAlt] = ebiten.KeyAlt
	keyMap[keyboard.KeyControl] = ebiten.KeyControl
	keyMap[keyboard.KeyShift] = ebiten.KeyShift
}

func setupButtonMap() {
	buttonMap = make(map[mouse.MouseButton]ebiten.MouseButton)
	buttonMap[mouse.ButtonRight] = ebiten.MouseButtonRight
	buttonMap[mouse.ButtonMiddle] = ebiten.MouseButtonMiddle
	buttonMap[mouse.ButtonLeft] = ebiten.MouseButtonLeft
}
