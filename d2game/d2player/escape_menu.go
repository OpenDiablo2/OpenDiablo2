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
	mainScreenID screenID = iota
)

type EscapeMenu struct {
	screens      []Screen
	activeScreen Screen

	isOpen bool

	pentLeft    *d2ui.Sprite
	pentRight   *d2ui.Sprite
	selectSound d2audio.SoundEffect
}

func NewEscapeMenu() *EscapeMenu {
	mainScreen := newMainScreen()

	screens := []Screen{
		mainScreenID: mainScreen,
	}

	return &EscapeMenu{
		screens:      screens,
		activeScreen: screens[mainScreenID],
	}
}

func (m *EscapeMenu) OnLoad() {
	for _, screen := range m.screens {
		screen.OnLoad()
	}

	animation, _ := d2asset.LoadAnimation(d2resource.PentSpin, d2resource.PaletteUnits)
	m.pentLeft, _ = d2ui.LoadSprite(animation)
	m.pentLeft.SetBlend(false)
	m.pentLeft.PlayBackward()

	m.pentRight, _ = d2ui.LoadSprite(animation)
	m.pentRight.SetBlend(false)
	m.pentRight.PlayForward()

	m.selectSound, _ = d2audio.LoadSoundEffect(d2resource.SFXCursorSelect)
}

func (m *EscapeMenu) OnEscKey() {
	if !m.isOpen {
		m.reset()
		m.isOpen = true
		return
	}

	if m.activeScreen.PrevScreen() != nil {
		m.activeScreen = m.activeScreen.PrevScreen()
		return
	}

	m.isOpen = false
}

func (m *EscapeMenu) reset() {
	m.activeScreen = m.screens[0]
	for _, screen := range m.screens {
		screen.Reset()
	}
}

func (m *EscapeMenu) Render(target d2render.Surface) {
	if !m.isOpen {
		return
	}
	m.activeScreen.Render(target)
}

func (m *EscapeMenu) Advance(elapsed float64) error {
	if !m.isOpen {
		return nil
	}
	return m.activeScreen.Advance(elapsed)
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
	m.activeScreen.OnUpKey()
}

func (m *EscapeMenu) OnDownKey() {
	if !m.isOpen {
		return
	}
	m.activeScreen.OnDownKey()
}

func (m *EscapeMenu) OnEnterKey() {
	if !m.isOpen {
		return
	}
	m.activeScreen.OnEnterKey()
}

func (m *EscapeMenu) OnMouseMove(event d2input.MouseMoveEvent) bool {
	if !m.isOpen {
		return false
	}
	return m.activeScreen.OnMouseMove(event)
}

func (m *EscapeMenu) OnMouseButtonDown(event d2input.MouseEvent) bool {
	if !m.isOpen {
		return false
	}
	return m.activeScreen.OnMouseButtonDown(event)
}
