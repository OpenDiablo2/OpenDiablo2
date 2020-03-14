package d2input_midi

//this input backend is just an example.

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/keyboard"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/mouse"

	"errors"
	"sync"
	//"github.com/xlab/midievent"
	"github.com/xlab/portmidi"
)

var (
	ErrNoDefaultInput = errors.New("No default midi input device was found.")
)

var (
	inName        string
	inDevID       int
	in            *portmidi.Stream
	keyState      map[int]bool // current state of the key
	keyPressed    map[int]bool // just pressed, zeroed every Advance
	keyReleased   map[int]bool // just released, cleared every advance
	mouseState    map[int]bool
	mousePressed  map[int]bool
	mouseReleased map[int]bool
	cursorX       int = 0
	cursorY       int = 0
	mux           *sync.Mutex
)

type Backend struct{}

func (b *Backend) Initialize() error {
	portmidi.Initialize()
	mux = &sync.Mutex{}

	keyState = make(map[int]bool)
	keyPressed = make(map[int]bool)
	keyReleased = make(map[int]bool)
	mouseState = make(map[int]bool)
	mousePressed = make(map[int]bool)
	mouseReleased = make(map[int]bool)

	defaultDev, defaultExists := portmidi.DefaultInputDeviceID()

	if !defaultExists {
		return ErrNoDefaultInput
	}

	in, _ = portmidi.NewInputStream(defaultDev, 1024, 0)

	return nil
}

func processMidi() {
	clearPressedReleasedState() // clear just pressed/released
	// this was blocking, so we call as go routine in Advance
	for ev := range in.Source() {
		message := portmidi.Message(ev.Message)

		status := int(message.Status())
		d1 := int(message.Data1())
		d2 := int(message.Data2())

		handleCursor(status, d1, d2)
		handleKey(status, d1, d2)
		handleMouse(status, d1, d2)
	}
}

func handleCursor(status, d1, d2 int) {
	if status == 0xB0 { // MIDI CC, channel 0
		switch d1 {
		case 7:
			cursorX = int(800 * (float64(d2) / 128.0))
		case 8:
			cursorY = int(600 * (float64(d2) / 128.0))
		}
	}
}

func handleKey(status, d1, d2 int) {
	mux.Lock()
	hasVelocity := d2 > 0
	if status == 0x80 { // Note off, channel 0
		keyState[d1] = false
		keyPressed[d1] = false
		keyReleased[d1] = true
	}
	if status == 0x90 { // Note on, channel 0
		keyState[d1] = hasVelocity
		keyPressed[d1] = hasVelocity
		keyReleased[d1] = !hasVelocity
	}
	mux.Unlock()
}

func handleMouse(status, d1, d2 int) {
	mux.Lock()
	if status == 0x90 { // Note On, channel 0
		hasVelocity := d2 > 0
		if d1 == 125 {
			mouseState[0] = hasVelocity
			mousePressed[0] = hasVelocity
			mouseReleased[0] = !hasVelocity
		}
		if d1 == 126 {
			mouseState[1] = hasVelocity
			mousePressed[1] = hasVelocity
			mouseReleased[1] = !hasVelocity
		}
		if d1 == 127 {
			mouseState[2] = hasVelocity
			mousePressed[2] = hasVelocity
			mouseReleased[2] = !hasVelocity
		}
	}
	mux.Unlock()
}

func clearPressedReleasedState() {
	mux.Lock()
	for i := range keyPressed {
		keyPressed[i] = false
	}
	for i := range keyReleased {
		keyReleased[i] = false
	}
	for i := range mousePressed {
		mousePressed[i] = false
	}
	for i := range mouseReleased {
		mouseReleased[i] = false
	}
	mux.Unlock()
}

func (b *Backend) Advance(elapsed float64) error {
	go processMidi()
	return nil
}

func (b *Backend) CursorPosition() (int, int) {
	return cursorX, cursorY
}

func (b *Backend) IsKeyPressed(k keyboard.Key) bool {
	return keyState[int(k)]
}

func (b *Backend) IsKeyJustPressed(k keyboard.Key) bool {
	return keyPressed[int(k)]
}

func (b *Backend) IsKeyJustReleased(k keyboard.Key) bool {
	return keyReleased[int(k)]
}

func (b *Backend) IsMouseButtonPressed(mb mouse.MouseButton) bool {
	return mouseState[int(mb)]
}

func (b *Backend) IsMouseButtonJustPressed(mb mouse.MouseButton) bool {
	return mousePressed[int(mb)]
}

func (b *Backend) IsMouseButtonJustReleased(mb mouse.MouseButton) bool {
	return mouseReleased[int(mb)]
}

var tmp []rune

func (b *Backend) InputChars() []rune {
	return tmp // How are we supposed to input chars via midi? lol
}
