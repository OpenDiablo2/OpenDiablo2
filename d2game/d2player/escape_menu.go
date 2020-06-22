package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type screenID int

const (
	exitScreenID    screenID = -1
	mainScreenID             = 0
	optionsScreenID          = 1
)

type EscapeMenu struct {
	screens      []Screen
	activeScreen screenID
	isOpen       bool
	pentLeft     *d2ui.Sprite
	pentRight    *d2ui.Sprite
	selectSound  d2audio.SoundEffect
}

func NewEscapeMenu() *EscapeMenu {
	m := &EscapeMenu{}
	m.screens = []Screen{
		mainScreenID:    newMainScreen(m.switchScreen),
		optionsScreenID: newOptionsScreen(m.switchScreen),
	}
	m.activeScreen = mainScreenID
	return m
}

func (m *EscapeMenu) OnLoad() {
	pentLeftAnim, _ := d2asset.LoadAnimation(d2resource.PentSpin, d2resource.PaletteUnits)
	pentLeft, _ := d2ui.LoadSprite(pentLeftAnim)
	pentLeft.SetBlend(false)
	pentLeft.PlayBackward()

	pentRightAnim, _ := d2asset.LoadAnimation(d2resource.PentSpin, d2resource.PaletteUnits)
	pentRight, _ := d2ui.LoadSprite(pentRightAnim)
	pentRight.SetBlend(false)
	pentRight.PlayForward()

	selectSound, _ := d2audio.LoadSoundEffect(d2resource.SFXCursorSelect)

	for _, screen := range m.screens {
		screen.Load(pentLeft, pentRight, selectSound)
	}
}

func (m *EscapeMenu) OnEscKey() {
	if !m.isOpen {
		m.reset()
		m.isOpen = true
		return
	}

	switch m.activeScreen {
	case optionsScreenID:
		m.switchScreen(mainScreenID)
		return
	}

	m.isOpen = false
}

func (m *EscapeMenu) reset() {
	m.activeScreen = mainScreenID
	for _, screen := range m.screens {
		screen.Reset()
	}
}

func (m *EscapeMenu) Render(target d2render.Surface) {
	if !m.isOpen {
		return
	}
	m.screens[m.activeScreen].Render(target)
}

func (m *EscapeMenu) Advance(elapsed float64) error {
	if !m.isOpen {
		return nil
	}
	return m.screens[m.activeScreen].Advance(elapsed)
}

func (m *EscapeMenu) Toggle() {
	m.isOpen = !m.isOpen
}

func (m *EscapeMenu) IsOpen() bool {
	return m.isOpen
}

func (m *EscapeMenu) OnUpKey() {
	if !m.isOpen {
		return
	}
	m.screens[m.activeScreen].OnUpKey()
}

func (m *EscapeMenu) OnDownKey() {
	if !m.isOpen {
		return
	}
	m.screens[m.activeScreen].OnDownKey()
}

func (m *EscapeMenu) OnEnterKey() {
	if !m.isOpen {
		return
	}
	m.screens[m.activeScreen].OnEnterKey()
}

func (m *EscapeMenu) OnMouseMove(event d2input.MouseMoveEvent) bool {
	if !m.isOpen {
		return false
	}
	return m.screens[m.activeScreen].OnMouseMove(event)
}

func (m *EscapeMenu) OnMouseButtonDown(event d2input.MouseEvent) bool {
	if !m.isOpen {
		return false
	}
	if event.Button == d2input.MouseButtonLeft {
		return m.screens[m.activeScreen].OnLeftClick(event)
	}
	return false
}

func (m *EscapeMenu) switchScreen(screenID screenID) {
	if screenID == exitScreenID {
		m.Toggle()
		return
	}
	// Prevent visual glitches
	prev := m.activeScreen
	m.activeScreen = screenID
	m.screens[prev].Reset()
}
