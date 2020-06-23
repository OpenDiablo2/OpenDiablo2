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

type layoutCfg struct {
	showLayoutFn func(id layoutID)
}

func NewEscapeMenu() *EscapeMenu {
	m := &EscapeMenu{}
	cfg := &layoutCfg{
		showLayoutFn: m.showLayout,
	}
	m.layouts = []*d2gui.Layout{
		mainLayoutID:         newMainLayout(cfg),
		optionsLayoutID:      newOptionsLayout(cfg),
		soundOptionsLayoutID: newSoundOptionsLayout(cfg),
	}
	return m
}

func wrapLayout(fn func(base *d2gui.Layout)) *d2gui.Layout {
	base := d2gui.CreateLayout(d2gui.PositionTypeHorizontal)
	base.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	base.AddSpacerDynamic()

	left := base.AddLayout(d2gui.PositionTypeVertical)
	left.SetSize(52, 52)
	left.AddAnimatedSprite(d2resource.PentSpin, d2resource.PaletteUnits, d2gui.DirectionForward)

	center := base.AddLayout(d2gui.PositionTypeVertical)
	center.SetSize(menuSize, 0)
	center.SetHorizontalAlign(d2gui.HorizontalAlignCenter)
	center.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	center.AddSpacerDynamic()

	fn(center)

	center.AddSpacerDynamic()

	right := base.AddLayout(d2gui.PositionTypeVertical)
	right.SetSize(52, 52)
	right.AddAnimatedSprite(d2resource.PentSpin, d2resource.PaletteUnits, d2gui.DirectionBackward)

	base.AddSpacerDynamic()

	return base
}

func newMainLayout(cfg *layoutCfg) *d2gui.Layout {
	return wrapLayout(func(base *d2gui.Layout) {
		addBigSelectionLabel(base, cfg, "options", optionsLayoutID)
		addBigSelectionLabel(base, cfg, "save and exit game", noLayoutID)
		addBigSelectionLabel(base, cfg, "return to game", noLayoutID)
	})
}

func newOptionsLayout(cfg *layoutCfg) *d2gui.Layout {
	return wrapLayout(func(base *d2gui.Layout) {
		addBigSelectionLabel(base, cfg, "sound options", soundOptionsLayoutID)
		addBigSelectionLabel(base, cfg, "video options", soundOptionsLayoutID)
		addBigSelectionLabel(base, cfg, "automap options", soundOptionsLayoutID)
		addBigSelectionLabel(base, cfg, "configure controls", soundOptionsLayoutID)
		addBigSelectionLabel(base, cfg, "previous menu", mainLayoutID)
	})
}

func newSoundOptionsLayout(cfg *layoutCfg) *d2gui.Layout {
	return wrapLayout(func(base *d2gui.Layout) {
		addTitle(base, "sound options")
		addOnOffLabel(base, "3d sound")
		addOnOffLabel(base, "environmental effects")
		addSmallSelectionLabel(base, cfg, "previous menu", optionsLayoutID)
	})
}

func addTitle(layout *d2gui.Layout, text string) {
	layout.AddLabel(text, d2gui.FontStyle42Units)
	layout.AddSpacerStatic(10, labelGutter)
}

func addSmallSelectionLabel(layout *d2gui.Layout, cfg *layoutCfg, text string, targetLayout layoutID) {
	label, _ := layout.AddLabel(text, d2gui.FontStyle30Units)
	label.SetMouseClickHandler(func(_ d2input.MouseEvent) {
		cfg.showLayoutFn(targetLayout)
	})
	layout.AddSpacerStatic(10, labelGutter)
}

func addBigSelectionLabel(layout *d2gui.Layout, cfg *layoutCfg, text string, targetLayout layoutID) {
	label, _ := layout.AddLabel(text, d2gui.FontStyle42Units)
	label.SetMouseClickHandler(func(_ d2input.MouseEvent) {
		cfg.showLayoutFn(targetLayout)
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
