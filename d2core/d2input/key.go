// Package d2input provides interaction with input services providing key and mouse interactions.
package d2input

// Key represents button on a traditional keyboard.
type Key int

const (
	// Key0 is the number 0
	Key0 Key = iota
	// Key1 is the number 1
	Key1
	// Key2 is the number 2
	Key2
	// Key3 is the number 3
	Key3
	// Key4 is the number 4
	Key4
	// Key5 is the number 5
	Key5
	// Key6 is the number 6
	Key6
	// Key7 is the number 7
	Key7
	// Key8 is the number 8
	Key8
	// Key9 is the number 9
	Key9
	// KeyA is the letter A
	KeyA
	// KeyB is the letter B
	KeyB
	// KeyC is the letter C
	KeyC
	// KeyD is the letter D
	KeyD
	// KeyE is the letter E
	KeyE
	// KeyF is the letter F
	KeyF
	// KeyG is the letter G
	KeyG
	// KeyH is the letter H
	KeyH
	// KeyI is the letter I
	KeyI
	// KeyJ is the letter J
	KeyJ
	// KeyK is the letter K
	KeyK
	// KeyL is the letter L
	KeyL
	// KeyM is the letter M
	KeyM
	// KeyN is the letter N
	KeyN
	// KeyO is the letter O
	KeyO
	// KeyP is the letter P
	KeyP
	// KeyQ is the letter Q
	KeyQ
	// KeyR is the letter R
	KeyR
	// KeyS is the letter S
	KeyS
	// KeyT is the letter T
	KeyT
	// KeyU is the letter U
	KeyU
	// KeyV is the letter V
	KeyV
	// KeyW is the letter W
	KeyW
	// KeyX is the letter X
	KeyX
	// KeyY is the letter Y
	KeyY
	// KeyZ is the letter Z
	KeyZ
	// KeyApostrophe is the Apostrophe
	KeyApostrophe
	// KeyBackslash is the Backslash
	KeyBackslash
	// KeyBackspace is the Backspace
	KeyBackspace
	// KeyCapsLock is the CapsLock
	KeyCapsLock
	// KeyComma is the Comma
	KeyComma
	// KeyDelete is the Delete
	KeyDelete
	// KeyDown is the down arrow key
	KeyDown
	// KeyEnd is the End
	KeyEnd
	// KeyEnter is the Enter
	KeyEnter
	// KeyEqual is the Equal
	KeyEqual
	// KeyEscape is the Escape
	KeyEscape
	// KeyF1 is the function F1
	KeyF1
	// KeyF2 is the function F2
	KeyF2
	// KeyF3 is the function F3
	KeyF3
	// KeyF4 is the function F4
	KeyF4
	// KeyF5 is the function F5
	KeyF5
	// KeyF6 is the function F6
	KeyF6
	// KeyF7 is the function F7
	KeyF7
	// KeyF8 is the function F8
	KeyF8
	// KeyF9 is the function F9
	KeyF9
	// KeyF10 is the function F10
	KeyF10
	// KeyF11 is the function F11
	KeyF11
	// KeyF12 is the function F12
	KeyF12
	// KeyGraveAccent is the Grave Accent
	KeyGraveAccent
	// KeyHome is the home key
	KeyHome
	// KeyInsert is the insert key
	KeyInsert
	// KeyKP0 is keypad 0
	KeyKP0
	// KeyKP1 is keypad 1
	KeyKP1
	// KeyKP2 is keypad 2
	KeyKP2
	// KeyKP3 is keypad 3
	KeyKP3
	// KeyKP4 is keypad 4
	KeyKP4
	// KeyKP5 is keypad 5
	KeyKP5
	// KeyKP6 is keypad 6
	KeyKP6
	// KeyKP7 is keypad 7
	KeyKP7
	// KeyKP8 is keypad 8
	KeyKP8
	// KeyKP9 is keypad 9
	KeyKP9
	// KeyKPAdd is keypad Add
	KeyKPAdd
	// KeyKPDecimal is keypad Decimal
	KeyKPDecimal
	// KeyKPDivide is keypad Divide
	KeyKPDivide
	// KeyKPEnter is keypad Enter
	KeyKPEnter
	// KeyKPEqual is keypad Equal
	KeyKPEqual
	// KeyKPMultiply is keypad Multiply
	KeyKPMultiply
	// KeyKPSubtract is keypad Subtract
	KeyKPSubtract
	// KeyLeft is the left arrow key
	KeyLeft
	// KeyLeftBracket is the left bracket
	KeyLeftBracket
	// KeyMenu is the Menu key
	KeyMenu
	// KeyMinus is the Minus key
	KeyMinus
	// KeyNumLock is the NumLock key
	KeyNumLock
	// KeyPageDown is the PageDown key
	KeyPageDown
	// KeyPageUp is the PageUp key
	KeyPageUp
	// KeyPause is the Pause key
	KeyPause
	// KeyPeriod is the Period key
	KeyPeriod
	// KeyPrintScreen is the PrintScreen key
	KeyPrintScreen
	// KeyRight is the right arrow key
	KeyRight
	// KeyRightBracket is the right bracket key
	KeyRightBracket
	// KeyScrollLock is the scroll lock key
	KeyScrollLock
	// KeySemicolon is the semicolon key
	KeySemicolon
	// KeySlash is the front slash key
	KeySlash
	// KeySpace is the space key
	KeySpace
	// KeyTab is the tab key
	KeyTab
	// KeyUp is the up arrow key
	KeyUp
	// KeyAlt is the alt key
	KeyAlt
	// KeyControl is the control key
	KeyControl
	// KeyShift is the shift key
	KeyShift

	// Lowest key in key constants
	keyMin = Key0
	// Highest key is key constants
	keyMax = KeyShift
)

// KeyMod represents a "modified" key action. This could mean, for example, ctrl-S
type KeyMod int

const (
	// KeyModAlt is the Alt key modifier
	KeyModAlt KeyMod = 1 << iota
	// KeyModControl is the Control key modifier
	KeyModControl
	// KeyModShift is the Shift key modifier
	KeyModShift
)
