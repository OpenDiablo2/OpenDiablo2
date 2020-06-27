package d2input

import (
	"sort"
)

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

type InputService interface {
	CursorPosition() (x int, y int)
	InputChars() []rune
	IsKeyPressed(key Key) bool
	IsKeyJustPressed(key Key) bool
	IsKeyJustReleased(key Key) bool
	IsMouseButtonPressed(button MouseButton) bool
	IsMouseButtonJustPressed(button MouseButton) bool
	IsMouseButtonJustReleased(button MouseButton) bool
	KeyPressDuration(key Key) int
}

type inputManager struct {
	inputService InputService
	cursorX      int
	cursorY      int

	buttonMod MouseButtonMod
	keyMod    KeyMod

	entries handlerEntryList
}

func (im *inputManager) advance(_ float64) error {
	cursorX, cursorY := im.inputService.CursorPosition()

	im.keyMod = 0
	if im.inputService.IsKeyPressed(KeyAlt) {
		im.keyMod |= KeyModAlt
	}
	if im.inputService.IsKeyPressed(KeyControl) {
		im.keyMod |= KeyModControl
	}
	if im.inputService.IsKeyPressed(KeyShift) {
		im.keyMod |= KeyModShift
	}

	im.buttonMod = 0
	if im.inputService.IsMouseButtonPressed(MouseButtonLeft) {
		im.buttonMod |= MouseButtonModLeft
	}
	if im.inputService.IsMouseButtonPressed(MouseButtonMiddle) {
		im.buttonMod |= MouseButtonModMiddle
	}
	if im.inputService.IsMouseButtonPressed(MouseButtonRight) {
		im.buttonMod |= MouseButtonModRight
	}

	eventBase := HandlerEvent{
		im.keyMod,
		im.buttonMod,
		cursorX,
		cursorY,
	}

	for key := keyMin; key <= keyMax; key++ {
		if im.inputService.IsKeyJustPressed(key) {
			event := KeyEvent{HandlerEvent: eventBase, Key: key}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(KeyDownHandler); ok {
					return l.OnKeyDown(event)
				}

				return false
			})
		}

		if im.inputService.IsKeyPressed(key) {
			event := KeyEvent{HandlerEvent: eventBase, Key: key, Duration: im.inputService.KeyPressDuration(key)}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(KeyRepeatHandler); ok {
					return l.OnKeyRepeat(event)
				}

				return false
			})
		}

		if im.inputService.IsKeyJustReleased(key) {
			event := KeyEvent{HandlerEvent: eventBase, Key: key}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(KeyUpHandler); ok {
					return l.OnKeyUp(event)
				}

				return false
			})
		}
	}

	if chars := im.inputService.InputChars(); len(chars) > 0 {
		event := KeyCharsEvent{eventBase, chars}
		im.propagate(func(handler Handler) bool {
			if l, ok := handler.(KeyCharsHandler); ok {
				l.OnKeyChars(event)
			}

			return false
		})
	}

	for button := mouseButtonMin; button <= mouseButtonMax; button++ {
		if im.inputService.IsMouseButtonJustPressed(button) {
			event := MouseEvent{eventBase, MouseButton(button)}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(MouseButtonDownHandler); ok {
					return l.OnMouseButtonDown(event)
				}

				return false
			})
		}

		if im.inputService.IsMouseButtonJustReleased(button) {
			event := MouseEvent{eventBase, MouseButton(button)}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(MouseButtonUpHandler); ok {
					return l.OnMouseButtonUp(event)
				}

				return false
			})
		}
		if im.inputService.IsMouseButtonPressed(button) {
			event := MouseEvent{eventBase, MouseButton(button)}
			im.propagate(func(handler Handler) bool {
				if l, ok := handler.(MouseButtonRepeatHandler); ok {
					return l.OnMouseButtonRepeat(event)
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
