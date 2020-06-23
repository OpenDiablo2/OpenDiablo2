package d2player

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
)

type layoutID int

const (
	noLayoutID           layoutID = -1
	mainLayoutID                  = 0
	optionsLayoutID               = 1
	soundOptionsLayoutID          = 2

	labelGutter = 2
)

type EscapeMenu struct {
	isOpen        bool
	selectSound   d2audio.SoundEffect
	currentLayout layoutID
	layouts       []*d2gui.Layout
}

func NewEscapeMenu() *EscapeMenu {
	m := &EscapeMenu{}
	m.layouts = []*d2gui.Layout{
		mainLayoutID:         newMainLayout(m.showLayout),
		optionsLayoutID:      newOptionsLayout(m.showLayout),
		soundOptionsLayoutID: newSoundOptionsLayout(m.showLayout),
	}
	return m
}

func (m *EscapeMenu) Close() {
	m.isOpen = false
	d2gui.SetLayout(nil)
}

func (m *EscapeMenu) Open() {
	m.isOpen = true
	m.setLayout(mainLayoutID)
}

func (m *EscapeMenu) showLayout(id layoutID) {
	m.selectSound.Play()

	if id == noLayoutID {
		m.Close()
		return
	}

	m.setLayout(id)
}

func (m *EscapeMenu) setLayout(id layoutID) {
	d2gui.SetLayout(m.layouts[id])
	m.currentLayout = id
}

func newMainLayout(showLayoutFn func(layoutID)) *d2gui.Layout {
	mainLayout := d2gui.CreateLayout(d2gui.PositionTypeHorizontal)
	mainLayout.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	mainLayout.AddSpacerDynamic()
	mainLayout.SetMouseClickHandler(func(_ d2input.MouseEvent) {
		fmt.Println("click main layout")
	})

	left := mainLayout.AddLayout(d2gui.PositionTypeVertical)
	left.AddSprite(d2resource.PentSpin, d2resource.PaletteUnits)

	center := mainLayout.AddLayout(d2gui.PositionTypeVertical)
	center.SetHorizontalAlign(d2gui.HorizontalAlignCenter)
	center.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	center.AddSpacerDynamic()
	addBigSelectionLabel(center, showLayoutFn, "options", optionsLayoutID)
	addBigSelectionLabel(center, showLayoutFn, "save and exit game", noLayoutID)
	addBigSelectionLabel(center, showLayoutFn, "return to game", noLayoutID)
	center.AddSpacerDynamic()

	right := mainLayout.AddLayout(d2gui.PositionTypeVertical)
	right.AddSprite(d2resource.PentSpin, d2resource.PaletteUnits)

	mainLayout.AddSpacerDynamic()
	return mainLayout
}

func newOptionsLayout(showLayoutFn func(layoutID)) *d2gui.Layout {
	mainLayout := d2gui.CreateLayout(d2gui.PositionTypeHorizontal)
	mainLayout.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	mainLayout.AddSpacerDynamic()

	left := mainLayout.AddLayout(d2gui.PositionTypeVertical)
	left.AddSprite(d2resource.PentSpin, d2resource.PaletteUnits)

	center := mainLayout.AddLayout(d2gui.PositionTypeVertical)
	center.SetHorizontalAlign(d2gui.HorizontalAlignCenter)
	center.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	center.AddSpacerDynamic()
	addBigSelectionLabel(center, showLayoutFn, "sound options", soundOptionsLayoutID)
	addBigSelectionLabel(center, showLayoutFn, "video options", soundOptionsLayoutID)
	addBigSelectionLabel(center, showLayoutFn, "automap options", soundOptionsLayoutID)
	addBigSelectionLabel(center, showLayoutFn, "configure controls", soundOptionsLayoutID)
	addBigSelectionLabel(center, showLayoutFn, "previous menu", mainLayoutID)
	center.AddSpacerDynamic()

	right := mainLayout.AddLayout(d2gui.PositionTypeVertical)
	right.AddSprite(d2resource.PentSpin, d2resource.PaletteUnits)

	mainLayout.AddSpacerDynamic()
	return mainLayout
}

func newSoundOptionsLayout(showLayoutFn func(layoutID)) *d2gui.Layout {
	mainLayout := d2gui.CreateLayout(d2gui.PositionTypeHorizontal)
	mainLayout.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	mainLayout.AddSpacerDynamic()

	left := mainLayout.AddLayout(d2gui.PositionTypeVertical)
	left.AddSprite(d2resource.PentSpin, d2resource.PaletteUnits)

	center := mainLayout.AddLayout(d2gui.PositionTypeVertical)
	center.SetHorizontalAlign(d2gui.HorizontalAlignCenter)
	center.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	center.AddSpacerDynamic()
	addTitle(center, "sound options")
	addOnOffLabel(center, "3d sound")
	addOnOffLabel(center, "environmental effects")
	addSmallSelectionLabel(center, showLayoutFn, "previous menu", optionsLayoutID)
	center.AddSpacerDynamic()

	right := mainLayout.AddLayout(d2gui.PositionTypeVertical)
	right.AddSprite(d2resource.PentSpin, d2resource.PaletteUnits)

	mainLayout.AddSpacerDynamic()
	return mainLayout
}

func addTitle(layout *d2gui.Layout, text string) {
	layout.AddLabel("sound options", d2gui.FontStyle42Units)
	layout.AddSpacerStatic(10, labelGutter)
}

func addSmallSelectionLabel(layout *d2gui.Layout, showLayoutFn func(layoutID), text string, targetLayout layoutID) {
	label, _ := layout.AddLabel(text, d2gui.FontStyle30Units)
	label.SetMouseClickHandler(func(_ d2input.MouseEvent) {
		showLayoutFn(targetLayout)
	})
	layout.AddSpacerStatic(10, labelGutter)
}

func addBigSelectionLabel(layout *d2gui.Layout, showLayoutFn func(layoutID), text string, targetLayout layoutID) {
	label, _ := layout.AddLabel(text, d2gui.FontStyle42Units)
	label.SetMouseClickHandler(func(_ d2input.MouseEvent) {
		showLayoutFn(targetLayout)
	})
	layout.AddSpacerStatic(10, labelGutter)
}

func addOnOffLabel(layout *d2gui.Layout, text string) {
	l := layout.AddLayout(d2gui.PositionTypeHorizontal)
	l.AddLabel(text, d2gui.FontStyle30Units)
	l.AddSpacerDynamic()
	lbl, _ := l.AddLabel("on", d2gui.FontStyle30Units)
	l.SetMouseClickHandler(func(_ d2input.MouseEvent) {
		if lbl.GetText() == "on" {
			lbl.SetText("off")
			return
		}
		lbl.SetText("on")
	})
	layout.AddSpacerStatic(10, labelGutter)
}

func (m *EscapeMenu) OnLoad() {
	m.selectSound, _ = d2audio.LoadSoundEffect(d2resource.SFXCursorSelect)
}

func (m *EscapeMenu) OnEscKey() {
	if !m.isOpen {
		m.Open()
		return
	}

	switch m.currentLayout {
	case optionsLayoutID:
		m.setLayout(mainLayoutID)
		return
	case soundOptionsLayoutID:
		m.setLayout(optionsLayoutID)
		return
	}

	m.Close()
}

func (m *EscapeMenu) IsOpen() bool {
	return m.isOpen
}
