package d2input

import (
	"errors"
	"sort"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

type inputManager struct {
	app          d2interface.App
	inputService d2interface.InputService
	cursorX      int
	cursorY      int

	buttonMod d2interface.MouseButtonMod
	keyMod    d2interface.KeyMod

	entries handlerEntryList
}

func (im *inputManager) Advance(_, _ float64) error {
	cursorX, cursorY := im.inputService.CursorPosition()

	im.keyMod = 0
	if im.inputService.IsKeyPressed(d2interface.KeyAlt) {
		im.keyMod |= d2interface.KeyModAlt
	}
	if im.inputService.IsKeyPressed(d2interface.KeyControl) {
		im.keyMod |= d2interface.KeyModControl
	}
	if im.inputService.IsKeyPressed(d2interface.KeyShift) {
		im.keyMod |= d2interface.KeyModShift
	}

	im.buttonMod = 0
	if im.inputService.IsMouseButtonPressed(d2interface.MouseButtonLeft) {
		im.buttonMod |= d2interface.MouseButtonModLeft
	}
	if im.inputService.IsMouseButtonPressed(d2interface.MouseButtonMiddle) {
		im.buttonMod |= d2interface.MouseButtonModMiddle
	}
	if im.inputService.IsMouseButtonPressed(d2interface.MouseButtonRight) {
		im.buttonMod |= d2interface.MouseButtonModRight
	}

	eventBase := HandlerEvent{
		im.keyMod,
		im.buttonMod,
		cursorX,
		cursorY,
	}

	for key := d2interface.KeyMin; key <= d2interface.KeyMax; key++ {
		if im.inputService.IsKeyJustPressed(key) {
			event := KeyEvent{HandlerEvent: eventBase, key: key}
			im.propagate(func(handler d2interface.InputEventHandler) bool {
				if l, ok := handler.(d2interface.KeyDownHandler); ok {
					return l.OnKeyDown(&event)
				}

				return false
			})
		}

		if im.inputService.IsKeyPressed(key) {
			event := KeyEvent{
				HandlerEvent: eventBase,
				key:          key,
				duration:     im.inputService.KeyPressDuration(key),
			}
			im.propagate(func(handler d2interface.InputEventHandler) bool {
				if l, ok := handler.(d2interface.KeyRepeatHandler); ok {
					return l.OnKeyRepeat(&event)
				}

				return false
			})
		}

		if im.inputService.IsKeyJustReleased(key) {
			event := KeyEvent{HandlerEvent: eventBase, key: key}
			im.propagate(func(handler d2interface.InputEventHandler) bool {
				if l, ok := handler.(d2interface.KeyUpHandler); ok {
					return l.OnKeyUp(&event)
				}

				return false
			})
		}
	}

	if chars := im.inputService.InputChars(); len(chars) > 0 {
		event := KeyCharsEvent{eventBase, chars}
		im.propagate(func(handler d2interface.InputEventHandler) bool {
			if l, ok := handler.(d2interface.KeyCharsHandler); ok {
				l.OnKeyChars(&event)
			}

			return false
		})
	}

	for button := d2interface.MouseButtonMin; button <= d2interface.
		MouseButtonMax; button++ {
		if im.inputService.IsMouseButtonJustPressed(button) {
			event := MouseEvent{eventBase, button}
			im.propagate(func(handler d2interface.InputEventHandler) bool {
				if l, ok := handler.(d2interface.MouseButtonDownHandler); ok {
					return l.OnMouseButtonDown(&event)
				}

				return false
			})
		}

		if im.inputService.IsMouseButtonJustReleased(button) {
			event := MouseEvent{eventBase, button}
			im.propagate(func(handler d2interface.InputEventHandler) bool {
				if l, ok := handler.(d2interface.MouseButtonUpHandler); ok {
					return l.OnMouseButtonUp(&event)
				}

				return false
			})
		}
		if im.inputService.IsMouseButtonPressed(button) {
			event := MouseEvent{eventBase, button}
			im.propagate(func(handler d2interface.InputEventHandler) bool {
				if l, ok := handler.(d2interface.MouseButtonRepeatHandler); ok {
					return l.OnMouseButtonRepeat(&event)
				}

				return false
			})
		}
	}

	if im.cursorX != cursorX || im.cursorY != cursorY {
		event := MouseMoveEvent{eventBase}
		im.propagate(func(handler d2interface.InputEventHandler) bool {
			if l, ok := handler.(d2interface.MouseMoveHandler); ok {
				return l.OnMouseMove(&event)
			}

			return false
		})

		im.cursorX, im.cursorY = cursorX, cursorY
	}

	return nil
}

func (im *inputManager) Render(_ d2interface.Surface) error {
	return nil // nothing to render
}

func (im *inputManager) Initialize() error {
	return nil // nothing to initialize, yet
}

func (im *inputManager) BindApp(app d2interface.App) error {
	if im.app != nil {
		return errors.New("input manager already bound to an app")
	}

	im.app = app

	return nil
}

func (im *inputManager) UnbindApp(app d2interface.App) error {
	if im.app == nil {
		return errors.New("input manager not bound to an app")
	}

	im.app = nil

	return nil
}

func (im *inputManager) bindHandler(handler d2interface.InputEventHandler, priority d2interface.Priority) error {
	for _, entry := range im.entries {
		if entry.handler == handler {
			return ErrHasReg
		}
	}

	im.entries = append(im.entries, handlerEntry{handler, priority})
	sort.Sort(im.entries)

	return nil
}

func (im *inputManager) unbindHandler(handler d2interface.InputEventHandler) error {
	for i, entry := range im.entries {
		if entry.handler == handler {
			copy(im.entries[i:], im.entries[i+1:])
			im.entries = im.entries[:len(im.entries)-1]
			return nil
		}
	}

	return ErrNotReg
}

func (im *inputManager) propagate(callback func(d2interface.InputEventHandler) bool) {
	var priority d2interface.Priority
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

type handlerEntry struct {
	handler  d2interface.InputEventHandler
	priority d2interface.Priority
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
