package d2input

import (
	"sort"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/keyboard"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/mouse"
)

type InputBackend interface {
	Initialize() error
	Advance(float64) error

	CursorPosition() (int, int)

	IsKeyPressed(keyboard.Key) bool
	IsKeyJustPressed(keyboard.Key) bool
	IsKeyJustReleased(keyboard.Key) bool

	IsMouseButtonPressed(mouse.MouseButton) bool
	IsMouseButtonJustPressed(mouse.MouseButton) bool
	IsMouseButtonJustReleased(mouse.MouseButton) bool

	InputChars() []rune
}

type handlerEntry struct {
	handler  Handler
	priority Priority
}

type handlerEntryList []handlerEntry

func (lel handlerEntryList) Len() int {
	return len(lel)
}

func (lel handlerEntryList) Swap(i, j int) {
	lel[i], lel[j] = lel[j], lel[i]
}

func (lel handlerEntryList) Less(i, j int) bool {
	return lel[i].priority > lel[j].priority
}

type inputManager struct {
	cursorX int
	cursorY int

	buttonMod MouseButtonMod
	keyMod    KeyMod

	entries handlerEntryList

	backend InputBackend
}

func (im *inputManager) advance(elapsed float64) error {
	if err := im.backend.Advance(elapsed); err != nil {
		return err
	}

	cursorX, cursorY := im.backend.CursorPosition()

	im.keyMod = 0
	if im.backend.IsKeyPressed(keyboard.KeyAlt) {
		im.keyMod |= KeyModAlt
	}
	if im.backend.IsKeyPressed(keyboard.KeyControl) {
		im.keyMod |= KeyModControl
	}
	if im.backend.IsKeyPressed(keyboard.KeyShift) {
		im.keyMod |= KeyModShift
	}

	im.buttonMod = 0
	if im.backend.IsMouseButtonPressed(mouse.ButtonLeft) {
		im.buttonMod |= MouseButtonModLeft
	}
	if im.backend.IsMouseButtonPressed(mouse.ButtonMiddle) {
		im.buttonMod |= MouseButtonModMiddle
	}
	if im.backend.IsMouseButtonPressed(mouse.ButtonRight) {
		im.buttonMod |= MouseButtonModRight
	}

	eventBase := HandlerEvent{
		im.keyMod,
		im.buttonMod,
		cursorX,
		cursorY,
	}

	keys := keyboard.GetKeys()
	for _, key := range keys {
		if im.backend.IsKeyJustPressed(key) {
			event := KeyEvent{eventBase, key}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(KeyDownHandler); ok {
					return l.OnKeyDown(event)
				}

				return false
			})
		}

		if im.backend.IsKeyPressed(key) {
			event := KeyEvent{eventBase, key}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(KeyRepeatHandler); ok {
					return l.OnKeyRepeat(event)
				}

				return false
			})
		}

		if im.backend.IsKeyJustReleased(key) {
			event := KeyEvent{eventBase, key}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(KeyUpHandler); ok {
					return l.OnKeyUp(event)
				}

				return false
			})
		}
	}

	if chars := im.backend.InputChars(); len(chars) > 0 {
		event := KeyCharsEvent{eventBase, chars}
		im.propagate(func(handler Handler) bool {
			if l, ok := handler.(KeyCharsHandler); ok {
				l.OnKeyChars(event)
			}

			return false
		})
	}

	for _, button := range mouse.MouseButtons {
		if im.backend.IsMouseButtonJustPressed(button) {
			event := MouseEvent{eventBase, button}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(MouseButtonDownHandler); ok {
					return l.OnMouseButtonDown(event)
				}

				return false
			})
		}

		if im.backend.IsMouseButtonJustReleased(button) {
			event := MouseEvent{eventBase, button}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(MouseButtonUpHandler); ok {
					return l.OnMouseButtonUp(event)
				}

				return false
			})
		}
	}

	if im.cursorX != cursorX || im.cursorY != cursorY {
		event := MouseMoveEvent{eventBase}
		im.propagate(func(handler Handler) bool {
			if l, ok := handler.(MouseMoveHandler); ok {
				return l.OnMouseMove(event)
			}

			return false
		})

		im.cursorX, im.cursorY = cursorX, cursorY
	}

	return nil
}

func (im *inputManager) bindHandler(handler Handler, priority Priority) error {
	for _, entry := range im.entries {
		if entry.handler == handler {
			return ErrHasReg
		}
	}

	im.entries = append(im.entries, handlerEntry{handler, priority})
	sort.Sort(im.entries)

	return nil
}

func (im *inputManager) unbindHandler(handler Handler) error {
	for i, entry := range im.entries {
		if entry.handler == handler {
			copy(im.entries[i:], im.entries[i+1:])
			im.entries = im.entries[:len(im.entries)-1]
			return nil
		}
	}

	return ErrNotReg
}

func (im *inputManager) propagate(callback func(Handler) bool) {
	var priority Priority
	var handled bool

	for _, entry := range im.entries {
		if priority > entry.priority && handled {
			break
		}

		if callback(entry.handler) {
			handled = true
		}

		priority = entry.priority
	}
}
