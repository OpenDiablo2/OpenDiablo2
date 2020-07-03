package d2input

import (
	"sort"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

type inputManager struct {
	inputService d2interface.InputService
	cursorX      int
	cursorY      int

	buttonMod d2interface.MouseButtonMod
	keyMod    d2interface.KeyMod

	entries handlerEntryList
}

// Advance advances the inputManager
func (im *inputManager) Advance(_, _ float64) error {
	im.updateKeyMod()
	im.updateButtonMod()

	cursorX, cursorY := im.inputService.CursorPosition()
	eventBase := HandlerEvent{
		im.keyMod,
		im.buttonMod,
		cursorX,
		cursorY,
	}

	for key := d2interface.KeyMin; key <= d2interface.KeyMax; key++ {
		im.updateJustPressedKey(key, eventBase)
		im.updateJustReleasedKey(key, eventBase)
		im.updatePressedKey(key, eventBase)
	}

	im.updateInputChars(eventBase)

	for button := d2interface.MouseButtonMin; button <= d2interface.MouseButtonMax; button++ {
		im.updateJustPressedButton(button, eventBase)
		im.updateJustReleasedButton(button, eventBase)
		im.updatePressedButton(button, eventBase)
	}

	im.updateCursor(cursorX, cursorY, eventBase)

	return nil
}

func (im *inputManager) updateKeyMod() {
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
}

func (im *inputManager) updateButtonMod() {
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
}

func (im *inputManager) updateJustPressedKey(k d2interface.Key, e HandlerEvent) {
	if im.inputService.IsKeyJustPressed(k) {
		event := KeyEvent{HandlerEvent: e, key: k}

		fn := func(handler d2interface.InputEventHandler) bool {
			if l, ok := handler.(d2interface.KeyDownHandler); ok {
				return l.OnKeyDown(&event)
			}

			return false
		}

		im.propagate(fn)
	}
}

func (im *inputManager) updateJustReleasedKey(k d2interface.Key, e HandlerEvent) {
	if im.inputService.IsKeyJustReleased(k) {
		event := KeyEvent{HandlerEvent: e, key: k}

		fn := func(handler d2interface.InputEventHandler) bool {
			if l, ok := handler.(d2interface.KeyUpHandler); ok {
				return l.OnKeyUp(&event)
			}

			return false
		}
		im.propagate(fn)
	}
}

func (im *inputManager) updatePressedKey(k d2interface.Key, e HandlerEvent) {
	if im.inputService.IsKeyPressed(k) {
		event := KeyEvent{
			HandlerEvent: e,
			key:          k,
			duration:     im.inputService.KeyPressDuration(k),
		}

		fn := func(handler d2interface.InputEventHandler) bool {
			if l, ok := handler.(d2interface.KeyRepeatHandler); ok {
				return l.OnKeyRepeat(&event)
			}

			return false
		}
		im.propagate(fn)
	}
}

func (im *inputManager) updateInputChars(eventBase HandlerEvent) {
	if chars := im.inputService.InputChars(); len(chars) > 0 {
		event := KeyCharsEvent{eventBase, chars}

		fn := func(handler d2interface.InputEventHandler) bool {
			if l, ok := handler.(d2interface.KeyCharsHandler); ok {
				l.OnKeyChars(&event)
			}

			return false
		}
		im.propagate(fn)
	}
}

func (im *inputManager) updateJustPressedButton(b d2interface.MouseButton, e HandlerEvent) {
	if im.inputService.IsMouseButtonJustPressed(b) {
		event := MouseEvent{e, b}

		fn := func(handler d2interface.InputEventHandler) bool {
			if l, ok := handler.(d2interface.MouseButtonDownHandler); ok {
				return l.OnMouseButtonDown(&event)
			}

			return false
		}
		im.propagate(fn)
	}
}

func (im *inputManager) updateJustReleasedButton(b d2interface.MouseButton, e HandlerEvent) {
	if im.inputService.IsMouseButtonJustReleased(b) {
		event := MouseEvent{e, b}

		fn := func(handler d2interface.InputEventHandler) bool {
			if l, ok := handler.(d2interface.MouseButtonUpHandler); ok {
				return l.OnMouseButtonUp(&event)
			}

			return false
		}
		im.propagate(fn)
	}
}

func (im *inputManager) updatePressedButton(b d2interface.MouseButton, e HandlerEvent) {
	if im.inputService.IsMouseButtonPressed(b) {
		event := MouseEvent{e, b}

		fn := func(handler d2interface.InputEventHandler) bool {
			if l, ok := handler.(d2interface.MouseButtonRepeatHandler); ok {
				return l.OnMouseButtonRepeat(&event)
			}

			return false
		}
		im.propagate(fn)
	}
}

func (im *inputManager) updateCursor(cursorX, cursorY int, e HandlerEvent) {
	if im.cursorX != cursorX || im.cursorY != cursorY {
		event := MouseMoveEvent{e}

		fn := func(handler d2interface.InputEventHandler) bool {
			if l, ok := handler.(d2interface.MouseMoveHandler); ok {
				return l.OnMouseMove(&event)
			}

			return false
		}
		im.propagate(fn)

		im.cursorX, im.cursorY = cursorX, cursorY
	}
}

// BindHandlerWithPriority adds an event handler with a specific call priority
func (im *inputManager) BindHandlerWithPriority(
	h d2interface.InputEventHandler,
	p d2interface.Priority) error {
	return singleton.bindHandler(h, p)
}

// BindHandler adds an event handler
func (im *inputManager) BindHandler(h d2interface.InputEventHandler) error {
	return im.bindHandler(h, d2interface.PriorityDefault)
}

// BindHandler adds an event handler
func (im *inputManager) bindHandler(h d2interface.InputEventHandler, p d2interface.Priority) error {
	for _, entry := range im.entries {
		if entry.handler == h {
			return ErrHasReg
		}
	}

	entry := handlerEntry{h, p}
	im.entries = append(im.entries, entry)
	sort.Sort(im.entries)

	return nil
}

// UnbindHandler removes a previously bound event handler
func (im *inputManager) UnbindHandler(handler d2interface.InputEventHandler) error {
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
