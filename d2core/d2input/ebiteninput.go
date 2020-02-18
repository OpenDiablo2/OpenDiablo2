package d2input

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func ebitenInput(im *inputManager) error {
	im.key[Key0]				= int(ebiten.Key0)
	im.key[Key1]				= int(ebiten.Key1)
	im.key[Key2]				= int(ebiten.Key2)
	im.key[Key3]				= int(ebiten.Key3)
	im.key[Key4]				= int(ebiten.Key4)
	im.key[Key5]				= int(ebiten.Key5)
	im.key[Key6]				= int(ebiten.Key6)
	im.key[Key7]				= int(ebiten.Key7)
	im.key[Key8]				= int(ebiten.Key8)
	im.key[Key9]				= int(ebiten.Key9)
	im.key[KeyA]				= int(ebiten.KeyA)
	im.key[KeyB]				= int(ebiten.KeyB)
	im.key[KeyC]				= int(ebiten.KeyC)
	im.key[KeyD]				= int(ebiten.KeyD)
	im.key[KeyE]				= int(ebiten.KeyE)
	im.key[KeyF]				= int(ebiten.KeyF)
	im.key[KeyG]				= int(ebiten.KeyG)
	im.key[KeyH]				= int(ebiten.KeyH)
	im.key[KeyI]				= int(ebiten.KeyI)
	im.key[KeyJ]				= int(ebiten.KeyJ)
	im.key[KeyK]				= int(ebiten.KeyK)
	im.key[KeyL]				= int(ebiten.KeyL)
	im.key[KeyM]				= int(ebiten.KeyM)
	im.key[KeyN]				= int(ebiten.KeyN)
	im.key[KeyO]				= int(ebiten.KeyO)
	im.key[KeyP]				= int(ebiten.KeyP)
	im.key[KeyQ]				= int(ebiten.KeyQ)
	im.key[KeyR]				= int(ebiten.KeyR)
	im.key[KeyS]				= int(ebiten.KeyS)
	im.key[KeyT]				= int(ebiten.KeyT)
	im.key[KeyU]				= int(ebiten.KeyU)
	im.key[KeyV]				= int(ebiten.KeyV)
	im.key[KeyW]				= int(ebiten.KeyW)
	im.key[KeyX]				= int(ebiten.KeyX)
	im.key[KeyY]				= int(ebiten.KeyY)
	im.key[KeyZ]				= int(ebiten.KeyZ)
	im.key[KeyApostrophe]		= int(ebiten.KeyApostrophe)
	im.key[KeyBackslash]		= int(ebiten.KeyBackslash)
	im.key[KeyBackspace]		= int(ebiten.KeyBackspace)
	im.key[KeyCapsLock]			= int(ebiten.KeyCapsLock)
	im.key[KeyComma]			= int(ebiten.KeyComma)
	im.key[KeyDelete]			= int(ebiten.KeyDelete)
	im.key[KeyDown]				= int(ebiten.KeyDown)
	im.key[KeyEnd]				= int(ebiten.KeyEnd)
	im.key[KeyEnter]			= int(ebiten.KeyEnter)
	im.key[KeyEqual]			= int(ebiten.KeyEqual)
	im.key[KeyEscape]			= int(ebiten.KeyEscape)
	im.key[KeyF1]				= int(ebiten.KeyF1)
	im.key[KeyF2]				= int(ebiten.KeyF2)
	im.key[KeyF3]				= int(ebiten.KeyF3)
	im.key[KeyF4]				= int(ebiten.KeyF4)
	im.key[KeyF5]				= int(ebiten.KeyF5)
	im.key[KeyF6]				= int(ebiten.KeyF6)
	im.key[KeyF7]				= int(ebiten.KeyF7)
	im.key[KeyF8]				= int(ebiten.KeyF8)
	im.key[KeyF9]				= int(ebiten.KeyF9)
	im.key[KeyF10]				= int(ebiten.KeyF10)
	im.key[KeyF11]				= int(ebiten.KeyF11)
	im.key[KeyF12]				= int(ebiten.KeyF12)
	im.key[KeyGraveAccent]		= int(ebiten.KeyGraveAccent)
	im.key[KeyHome]				= int(ebiten.KeyHome)
	im.key[KeyInsert]			= int(ebiten.KeyInsert)
	im.key[KeyKP0]				= int(ebiten.KeyKP0)
	im.key[KeyKP1]				= int(ebiten.KeyKP1)
	im.key[KeyKP2]				= int(ebiten.KeyKP2)
	im.key[KeyKP3]				= int(ebiten.KeyKP3)
	im.key[KeyKP4]				= int(ebiten.KeyKP4)
	im.key[KeyKP5]				= int(ebiten.KeyKP5)
	im.key[KeyKP6]				= int(ebiten.KeyKP6)
	im.key[KeyKP7]				= int(ebiten.KeyKP7)
	im.key[KeyKP8]				= int(ebiten.KeyKP8)
	im.key[KeyKP9]				= int(ebiten.KeyKP9)
	im.key[KeyKPAdd]			= int(ebiten.KeyKPAdd)
	im.key[KeyKPDecimal]		= int(ebiten.KeyKPDecimal)
	im.key[KeyKPDivide]			= int(ebiten.KeyKPDivide)
	im.key[KeyKPEnter]			= int(ebiten.KeyKPEnter)
	im.key[KeyKPEqual]			= int(ebiten.KeyKPEqual)
	im.key[KeyKPMultiply]		= int(ebiten.KeyKPMultiply)
	im.key[KeyKPSubtract]		= int(ebiten.KeyKPSubtract)
	im.key[KeyLeft]				= int(ebiten.KeyLeft)
	im.key[KeyLeftBracket]		= int(ebiten.KeyLeftBracket)
	im.key[KeyMenu]				= int(ebiten.KeyMenu)
	im.key[KeyMinus]			= int(ebiten.KeyMinus)
	im.key[KeyNumLock]			= int(ebiten.KeyNumLock)
	im.key[KeyPageDown]			= int(ebiten.KeyPageDown)
	im.key[KeyPageUp]			= int(ebiten.KeyPageUp)
	im.key[KeyPause]			= int(ebiten.KeyPause)
	im.key[KeyPeriod]			= int(ebiten.KeyPeriod)
	im.key[KeyPrintScreen]		= int(ebiten.KeyPrintScreen)
	im.key[KeyRight]			= int(ebiten.KeyRight)
	im.key[KeyRightBracket]		= int(ebiten.KeyRightBracket)
	im.key[KeyScrollLock]		= int(ebiten.KeyScrollLock)
	im.key[KeySemicolon]		= int(ebiten.KeySemicolon)
	im.key[KeySlash]			= int(ebiten.KeySlash)
	im.key[KeySpace]			= int(ebiten.KeySpace)
	im.key[KeyTab]				= int(ebiten.KeyTab)
	im.key[KeyUp]				= int(ebiten.KeyUp)
	im.key[KeyAlt]				= int(ebiten.KeyAlt)
	im.key[KeyControl]			= int(ebiten.KeyControl)
	im.key[KeyShift]			= int(ebiten.KeyShift)
	im.mouseButton[MouseButtonLeft]		= int(ebiten.MouseButtonLeft)
	im.mouseButton[MouseButtonMiddle]	= int(ebiten.MouseButtonMiddle)
	im.mouseButton[MouseButtonRight]	= int(ebiten.MouseButtonRight)

	int2key := func(n int) ebiten.Key {
		return ebiten.Key(n)
	}

	int2button := func(n int) ebiten.MouseButton {
		return ebiten.MouseButton(n)
	}

	im.isKeyPressed = func(k Key) bool {
		return ebiten.IsKeyPressed(int2key(im.key[k]))
	}

	im.isMouseButtonPressed	= func(b MouseButton) bool {
		return ebiten.IsMouseButtonPressed(int2button(im.mouseButton[b]))
	}

	im.cursorPosition = ebiten.CursorPosition
	im.keyMax = int(ebiten.KeyMax)
	im.inputChars = ebiten.InputChars

	im._backendAdvance = func() {

		for key, i := range im.key {
			eKey := int2key(i) // ebiten key

			if inpututil.IsKeyJustPressed(eKey) {
				event := KeyEvent{im.eventBase, Key(key)}
				im.propagate(func(handler Handler) bool {
					if l, ok := handler.(KeyDownHandler); ok {
						return l.OnKeyDown(event)
					}

					return false
				})
			}

			if im.isKeyPressed(key) {
				event := KeyEvent{im.eventBase, key}
				im.propagate(func(handler Handler) bool {
					if l, ok := handler.(KeyRepeatHandler); ok {
						return l.OnKeyRepeat(event)
					}

					return false
				})
			}

			if inpututil.IsKeyJustReleased(eKey) {
				event := KeyEvent{im.eventBase, key}
				im.propagate(func(handler Handler) bool {
					if l, ok := handler.(KeyUpHandler); ok {
						return l.OnKeyUp(event)
					}

					return false
				})
			}
		}

		if chars := im.inputChars(); len(chars) > 0 {
			event := KeyCharsEvent{im.eventBase, chars}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(KeyCharsHandler); ok {
					l.OnKeyChars(event)
				}

				return false
			})
		}

		for button, i := range im.mouseButton {
			eButton := int2button(i) // ebiten type button
			if inpututil.IsMouseButtonJustPressed(eButton) {
				event := MouseEvent{im.eventBase, button}
				im.propagate(func(handler Handler) bool {
					if l, ok := handler.(MouseButtonDownHandler); ok {
						return l.OnMouseButtonDown(event)
					}

					return false
				})
			}

			if inpututil.IsMouseButtonJustReleased(eButton) {
				event := MouseEvent{im.eventBase, button}
				im.propagate(func(handler Handler) bool {
					if l, ok := handler.(MouseButtonUpHandler); ok {
						return l.OnMouseButtonUp(event)
					}

					return false
				})
			}
		}

		if im.cursorX != cursorX || im.cursorY != cursorY {
			event := MouseMoveEvent{im.eventBase}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(MouseMoveHandler); ok {
					return l.OnMouseMove(event)
				}

				return false
			})

			im.cursorX, im.cursorY = cursorX, cursorY
		}

	}



	return nil
}
