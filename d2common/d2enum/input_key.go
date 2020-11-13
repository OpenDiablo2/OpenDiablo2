package d2enum

// Key represents button on a traditional keyboard.
type Key int

// Input keys
const (
	Key0 Key = iota
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
	KeyTilde
	KeyMouse3
	KeyMouse4
	KeyMouse5
	KeyMouseWheelUp
	KeyMouseWheelDown

	KeyMin = Key0
	KeyMax = KeyMouseWheelDown
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
