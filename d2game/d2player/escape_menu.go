package d2player

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
)

type (
	layoutID int
	optionID int
)

const (
	noLayoutID           layoutID = -1
	mainLayoutID                  = 0
	optionsLayoutID               = 1
	soundOptionsLayoutID          = 2

	opt3dSound    optionID = 0
	optEnvEffects          = 1

	labelGutter    = 2
	sidePanelsSize = 100
	pentSize       = 52
	menuSize       = 400
)

type EscapeMenu struct {
	isOpen        bool
	selectSound   d2audio.SoundEffect
	currentLayout layoutID

	leftPent  *d2gui.AnimatedSprite
	rightPent *d2gui.AnimatedSprite
	layouts   []*layout
}

type enumLabel struct {
	*d2gui.Label
	optionID optionID
	values   []string
	current  int
}

type hoverableElement interface {
	GetOffset() (int, int)
}

type layout struct {
	*d2gui.Layout
	leftPent  *d2gui.AnimatedSprite
	rightPent *d2gui.AnimatedSprite

	hoverableElements []hoverableElement
}

type layoutCfg struct {
	playSound    func()
	showLayout   func(id layoutID)
	wrapLayout   func(func(*d2gui.Layout)) *layout
	hoverElement func(el hoverableElement)
	updateValue  func(optID optionID, value string)
}

func NewEscapeMenu() *EscapeMenu {
	m := &EscapeMenu{}
	cfg := &layoutCfg{
		playSound:    m.playSound,
		showLayout:   m.showLayout,
		wrapLayout:   m.wrapLayout,
		hoverElement: m.onHoverElement,
		updateValue:  m.onUpdateValue,
	}
	m.layouts = []*layout{
		mainLayoutID:         newMainLayout(cfg),
		optionsLayoutID:      newOptionsLayout(cfg),
		soundOptionsLayoutID: newSoundOptionsLayout(cfg),
	}
	return m
}

func (m *EscapeMenu) wrapLayout(fn func(*d2gui.Layout)) *layout {
	base := d2gui.CreateLayout(d2gui.PositionTypeHorizontal)
	base.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	base.AddSpacerDynamic()

	center := base.AddLayout(d2gui.PositionTypeHorizontal)
	center.SetSize(menuSize, 0)

	left := center.AddLayout(d2gui.PositionTypeHorizontal)
	left.SetSize(sidePanelsSize, 0)
	leftPent, _ := left.AddAnimatedSprite(d2resource.PentSpin, d2resource.PaletteUnits, d2gui.DirectionForward)
	m.leftPent = leftPent

	f := center.AddLayout(d2gui.PositionTypeVertical)
	f.SetHorizontalAlign(d2gui.HorizontalAlignCenter)
	f.AddSpacerDynamic()
	fn(f)
	f.AddSpacerDynamic()

	right := center.AddLayout(d2gui.PositionTypeHorizontal)
	// For some reason, aligning the panel to the right won't align the pentagram, so we need to add a static spacer.
	right.AddSpacerStatic(sidePanelsSize-pentSize, 0)
	right.SetSize(sidePanelsSize, 0)
	rightPent, _ := right.AddAnimatedSprite(d2resource.PentSpin, d2resource.PaletteUnits, d2gui.DirectionBackward)
	m.rightPent = rightPent

	base.AddSpacerDynamic()
	return &layout{
		Layout:    base,
		leftPent:  leftPent,
		rightPent: rightPent,
	}
}

func newMainLayout(cfg *layoutCfg) *layout {
	return cfg.wrapLayout(func(base *d2gui.Layout) {
		addBigSelectionLabel(base, cfg, "options", optionsLayoutID)
		addBigSelectionLabel(base, cfg, "save and exit game", noLayoutID)
		addBigSelectionLabel(base, cfg, "return to game", noLayoutID)
	})
}

func newOptionsLayout(cfg *layoutCfg) *layout {
	return cfg.wrapLayout(func(base *d2gui.Layout) {
		addBigSelectionLabel(base, cfg, "sound options", soundOptionsLayoutID)
		addBigSelectionLabel(base, cfg, "video options", soundOptionsLayoutID)
		addBigSelectionLabel(base, cfg, "automap options", soundOptionsLayoutID)
		addBigSelectionLabel(base, cfg, "configure controls", soundOptionsLayoutID)
		addBigSelectionLabel(base, cfg, "previous menu", mainLayoutID)
	})
}

func newSoundOptionsLayout(cfg *layoutCfg) *layout {
	return cfg.wrapLayout(func(base *d2gui.Layout) {
		addTitle(base, "sound options")
		addEnumLabel(base, cfg, opt3dSound, "3d sound", []string{"on", "off"})
		addEnumLabel(base, cfg, optEnvEffects, "environmental effects", []string{"on", "off"})
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
		cfg.showLayout(targetLayout)
	})
	label.SetMouseEnterHandler(func(_ d2input.MouseMoveEvent) {
		cfg.hoverElement(label)
	})
	layout.AddSpacerStatic(10, labelGutter)
}

func addBigSelectionLabel(layout *d2gui.Layout, cfg *layoutCfg, text string, targetLayout layoutID) {
	label, _ := layout.AddLabel(text, d2gui.FontStyle42Units)
	label.SetMouseClickHandler(func(_ d2input.MouseEvent) {
		cfg.showLayout(targetLayout)
	})
	label.SetMouseEnterHandler(func(_ d2input.MouseMoveEvent) {
		cfg.hoverElement(label)
	})
	layout.AddSpacerStatic(10, labelGutter)
}

func addEnumLabel(layout *d2gui.Layout, cfg *layoutCfg, optID optionID, text string, values []string) {
	l := layout.AddLayout(d2gui.PositionTypeHorizontal)
	l.SetSize(menuSize, 0)
	l.AddLabel(text, d2gui.FontStyle30Units)
	l.SetMouseEnterHandler(func(_ d2input.MouseMoveEvent) {
		cfg.hoverElement(l)
	})
	l.AddSpacerDynamic()
	guiLabel, _ := l.AddLabel(values[0], d2gui.FontStyle30Units)
	label := &enumLabel{
		Label:    guiLabel,
		optionID: optID,
		values:   values,
		current:  0,
	}
	l.SetMouseClickHandler(func(_ d2input.MouseEvent) {
		cfg.playSound()
		next := (label.current + 1) % len(label.values)
		label.current = next
		label.Label.SetText(label.values[label.current])
		cfg.updateValue(label.optionID, label.values[label.current])
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

func (m *EscapeMenu) playSound() {
	m.selectSound.Play()
}

func (m *EscapeMenu) showLayout(id layoutID) {
	m.playSound()

	if id == noLayoutID {
		m.Close()
		return
	}

	m.setLayout(id)
}

func (m *EscapeMenu) onHoverElement(el hoverableElement) {
	_, y := el.GetOffset()

	x, _ := m.leftPent.GetPosition()
	m.leftPent.SetPosition(x, y+10)

	x, _ = m.rightPent.GetPosition()
	m.rightPent.SetPosition(x, y+10)
	return
}

func (m *EscapeMenu) onUpdateValue(optID optionID, value string) {
	fmt.Println(fmt.Sprintf("updating value %s to %s", optID, value))
}

func (m *EscapeMenu) setLayout(id layoutID) {
	m.leftPent = m.layouts[id].leftPent
	m.rightPent = m.layouts[id].rightPent
	d2gui.SetLayout(m.layouts[id].Layout)
	m.currentLayout = id
}
