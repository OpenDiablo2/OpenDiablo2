package d2gamescreen

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type EscapeOption int

const (
	EscapeOptions  = EscapeOption(iota)
	EscapeSaveExit = EscapeOption(iota)
	EscapeReturn   = EscapeOption(iota)
)

type mouseRegion int

const (
	regAbove = mouseRegion(iota)
	regIn    = mouseRegion(iota)
	regBelow = mouseRegion(iota)
)

// EscapeMenu is the overlay menu shown in-game when pressing Escape
type EscapeMenu struct {
	current     EscapeOption
	isOpen      bool
	labels      []d2ui.Label
	pentLeft    *d2ui.Sprite
	pentRight   *d2ui.Sprite
	selectSound d2audio.SoundEffect

	// pre-computations
	pentWidth  int
	pentHeight int
	textHeight int
}

// Creates an default instance of the EscapeMenu
func NewEscapeMenu() *EscapeMenu {
	return &EscapeMenu{
		labels: make([]d2ui.Label, 0),
	}
}

func (m *EscapeMenu) OnKeyDown(event d2input.KeyEvent) bool {
	switch event.Key {
	case d2input.KeyEscape:
		m.Toggle()
	case d2input.KeyUp:
		m.OnUpKey()
	case d2input.KeyDown:
		m.OnDownKey()
	case d2input.KeyEnter:
		m.OnEnterKey()
	default:
		return false
	}
	return false
}

func (m *EscapeMenu) OnLoad() error {
	d2input.BindHandler(m)
	m.labels = []d2ui.Label{
		d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
		d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
		d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
	}

	m.labels[EscapeOptions].SetText("OPTIONS")
	m.labels[EscapeSaveExit].SetText("SAVE AND EXIT GAME")
	m.labels[EscapeReturn].SetText("RETURN TO GAME")

	for i := range m.labels {
		m.labels[i].Alignment = d2ui.LabelAlignCenter
	}

	animation, _ := d2asset.LoadAnimation(d2resource.PentSpin, d2resource.PaletteUnits)
	m.pentLeft, _ = d2ui.LoadSprite(animation)
	m.pentLeft.SetBlend(false)
	m.pentLeft.PlayBackward()

	m.pentRight, _ = d2ui.LoadSprite(animation)
	m.pentRight.SetBlend(false)
	m.pentRight.PlayForward()

	m.pentWidth, m.pentHeight = m.pentLeft.GetFrameBounds()
	_, m.textHeight = m.labels[EscapeOptions].GetSize()

	m.selectSound, _ = d2audio.LoadSoundEffect(d2resource.SFXCursorSelect)

	return nil
}

func (m *EscapeMenu) Render(target d2render.Surface) error {
	if !m.isOpen {
		return nil
	}

	tw, _ := target.GetSize()
	// X Position of the mid-render target.
	midX := tw / 2

	// Y Coordinates for the center of the first option
	choiceStart := 210
	// Y Delta, in pixels, between center of choices
	choiceDx := 50
	// X Delta, in pixels, between center of pentagrams
	betwPentDist := 275

	for i := range m.labels {
		m.labels[i].SetPosition(midX, choiceStart+i*choiceDx-m.textHeight/2)
		m.labels[i].Render(target)
	}

	m.pentLeft.SetPosition(midX-(betwPentDist+m.pentWidth/2), choiceStart+int(m.current)*choiceDx+m.pentHeight/2)
	m.pentRight.SetPosition(midX+(betwPentDist-m.pentWidth/2), choiceStart+int(m.current)*choiceDx+m.pentHeight/2)

	m.pentLeft.Render(target)
	m.pentRight.Render(target)

	return nil
}

func (m *EscapeMenu) Advance(elapsed float64) error {
	if !m.isOpen {
		return nil
	}

	m.pentLeft.Advance(elapsed)
	m.pentRight.Advance(elapsed)
	return nil
}

func (m *EscapeMenu) IsOpen() bool {
	return m.isOpen
}

func (m *EscapeMenu) Toggle() {
	if !m.isOpen {
		m.reset()
	}
	m.isOpen = !m.isOpen
}

func (m *EscapeMenu) reset() {
	m.current = EscapeOptions
}

func (m *EscapeMenu) OnUpKey() {
	switch m.current {
	case EscapeSaveExit:
		m.current = EscapeOptions
	case EscapeReturn:
		m.current = EscapeSaveExit
	}
}

func (m *EscapeMenu) OnDownKey() {
	switch m.current {
	case EscapeOptions:
		m.current = EscapeSaveExit
	case EscapeSaveExit:
		m.current = EscapeReturn
	}
}

func (m *EscapeMenu) OnEnterKey() {
	m.selectCurrent()
}

// Moves current selection marker to closes option to mouse.
func (m *EscapeMenu) OnMouseMove(event d2input.MouseMoveEvent) bool {
	if !m.isOpen {
		return false
	}
	lbl := &m.labels[EscapeSaveExit]
	reg := m.toMouseRegion(event.HandlerEvent, lbl)

	switch reg {
	case regAbove:
		m.current = EscapeOptions
	case regIn:
		m.current = EscapeSaveExit
	case regBelow:
		m.current = EscapeReturn
	}

	return false
}

// Allows user to click on menu options in Y coord. of mouse is over label.
func (m *EscapeMenu) OnMouseButtonDown(event d2input.MouseEvent) bool {
	if !m.isOpen {
		return false
	}

	lbl := &m.labels[EscapeOptions]
	if m.toMouseRegion(event.HandlerEvent, lbl) == regIn {
		m.current = EscapeOptions
		m.selectCurrent()
		return false
	}

	lbl = &m.labels[EscapeSaveExit]
	if m.toMouseRegion(event.HandlerEvent, lbl) == regIn {
		m.current = EscapeSaveExit
		m.selectCurrent()
		return false
	}

	lbl = &m.labels[EscapeReturn]
	if m.toMouseRegion(event.HandlerEvent, lbl) == regIn {
		m.current = EscapeReturn
		m.selectCurrent()
		return false
	}

	return false
}

func (m *EscapeMenu) selectCurrent() {
	switch m.current {
	case EscapeOptions:
		m.onOptions()
		m.selectSound.Play()
	case EscapeSaveExit:
		m.onSaveAndExit()
		m.selectSound.Play()
	case EscapeReturn:
		m.onReturnToGame()
		m.selectSound.Play()
	}
}

// User clicked on "OPTIONS"
func (m *EscapeMenu) onOptions() error {
	log.Println("OPTIONS Clicked from Escape Menu")
	return nil
}

// User clicked on "SAVE AND EXIT"
func (m *EscapeMenu) onSaveAndExit() error {
	log.Println("SAVE AND EXIT GAME Clicked from Escape Menu")
	mainMenu := CreateMainMenu()
	mainMenu.SetScreenMode(ScreenModeMainMenu)
	d2screen.SetNextScreen(mainMenu)
	return nil
}

// User clicked on "RETURN TO GAME"
func (m *EscapeMenu) onReturnToGame() error {
	m.Toggle()
	return nil
}

// Where is the Y coordinate of the mouse compared to this label.
func (m *EscapeMenu) toMouseRegion(event d2input.HandlerEvent, lbl *d2ui.Label) mouseRegion {
	_, h := lbl.GetSize()
	y := lbl.Y
	my := event.Y

	if my < y {
		return regAbove
	}
	if my > (y + h) {
		return regBelow
	}
	return regIn
}
