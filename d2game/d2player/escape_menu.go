package d2player

import (
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
	menuSize    = 500
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

func wrapLayout(fn func(base *d2gui.Layout)) *d2gui.Layout {
	base := d2gui.CreateLayout(d2gui.PositionTypeHorizontal)
	base.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	base.AddSpacerDynamic()

	left := base.AddLayout(d2gui.PositionTypeVertical)
	left.SetSize(52, 0)
	left.AddAnimatedSprite(d2resource.PentSpin, d2resource.PaletteUnits, d2gui.DirectionForward)

	center := base.AddLayout(d2gui.PositionTypeVertical)
	center.SetSize(menuSize, 0)
	center.SetHorizontalAlign(d2gui.HorizontalAlignCenter)
	center.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	center.AddSpacerDynamic()

	fn(center)

	center.AddSpacerDynamic()

	right := base.AddLayout(d2gui.PositionTypeVertical)
	right.SetSize(52, 0)
	right.AddAnimatedSprite(d2resource.PentSpin, d2resource.PaletteUnits, d2gui.DirectionBackward)

	base.AddSpacerDynamic()

	return base
}

func newMainLayout(showLayoutFn func(layoutID)) *d2gui.Layout {
	return wrapLayout(func(base *d2gui.Layout) {
		addBigSelectionLabel(base, showLayoutFn, "options", optionsLayoutID)
		addBigSelectionLabel(base, showLayoutFn, "save and exit game", noLayoutID)
		addBigSelectionLabel(base, showLayoutFn, "return to game", noLayoutID)
	})
}

func newOptionsLayout(showLayoutFn func(layoutID)) *d2gui.Layout {
	return wrapLayout(func(base *d2gui.Layout) {
		addBigSelectionLabel(base, showLayoutFn, "sound options", soundOptionsLayoutID)
		addBigSelectionLabel(base, showLayoutFn, "video options", soundOptionsLayoutID)
		addBigSelectionLabel(base, showLayoutFn, "automap options", soundOptionsLayoutID)
		addBigSelectionLabel(base, showLayoutFn, "configure controls", soundOptionsLayoutID)
		addBigSelectionLabel(base, showLayoutFn, "previous menu", mainLayoutID)
	})
}

func newSoundOptionsLayout(showLayoutFn func(layoutID)) *d2gui.Layout {
	return wrapLayout(func(base *d2gui.Layout) {
		addTitle(base, "sound options")
		addOnOffLabel(base, "3d sound")
		addOnOffLabel(base, "environmental effects")
		addSmallSelectionLabel(base, showLayoutFn, "previous menu", optionsLayoutID)
	})
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
	l.SetSize(menuSize, 0)
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
