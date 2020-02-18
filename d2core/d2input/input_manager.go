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

type checkInput		func(Key)	bool
type checkCursor	func()		(int int)

type inputManager struct {
	cursorX					int
	cursorY					int

	buttonMod				MouseButtonMod
	keyMod					KeyMod

	eventBase				HandlerEvent
	entries					handlerEntryList

	key						map[Key]int
	mouseButton				map[MouseButton]int

	keyMax					int

	cursorPosition			func() (int, int)
	isKeyPressed			func(k Key) bool
	isMouseButtonPressed	func(b MouseButton) bool
	inputChars				func() []rune

	_backendAdvance			func()
}

var cursorX, cursorY int
func (im *inputManager) advance(elapsed float64) error {
	cursorX, cursorY = im.cursorPosition()

	im.keyMod = 0
	if im.isKeyPressed(KeyAlt) {
		im.keyMod |= KeyModAlt
	}
	if im.isKeyPressed(KeyControl) {
		im.keyMod |= KeyModControl
	}
	if im.isKeyPressed(KeyShift) {
		im.keyMod |= KeyModShift
	}

	im.buttonMod = 0
	if im.isMouseButtonPressed(MouseButtonLeft) {
		im.buttonMod |= MouseButtonModLeft
	}
	if im.isMouseButtonPressed(MouseButtonMiddle) {
		im.buttonMod |= MouseButtonModMiddle
	}
	if im.isMouseButtonPressed(MouseButtonRight) {
		im.buttonMod |= MouseButtonModRight
	}

	im.eventBase = HandlerEvent{
		im.keyMod,
		im.buttonMod,
		cursorX,
		cursorY,
	}

	im._backendAdvance()

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
