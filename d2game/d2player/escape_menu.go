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
	// UI
	labelGutter    = 10
	sidePanelsSize = 80
	pentSize       = 52
	menuSize       = 500

	// layouts
	noLayoutID                layoutID = -1
	mainLayoutID                       = 0
	optionsLayoutID                    = 1
	soundOptionsLayoutID               = 2
	videoOptionsLayoutID               = 3
	automapOptionsLayoutID             = 4
	configureControlsLayoutID          = 5

	// options
	optAudioSoundVolume          optionID = 0 // audio
	optAudioMusicVolume                   = 1
	optAudio3dSound                       = 2
	optAudioHardwareAcceleration          = 3
	optAudioEnvEffects                    = 4
	optAudioNpcSpeech                     = 5
	optVideoResolution                    = 6 // video
	optVideoLightingQuality               = 7
	optVideoBlendedShadows                = 8
	optVideoPerspective                   = 9
	optVideoGamma                         = 10
	optVideoContrast                      = 11
	optAutomapSize                        = 12 // automap
	optAutomapFade                        = 13
	optAutomapCenterWhenCleared           = 14
	optAutomapShowParty                   = 15
	optAutomapShowNames                   = 16
)

type EscapeMenu struct {
	isOpen        bool
	selectSound   d2audio.SoundEffect
	currentLayout layoutID

	leftPent  *d2gui.AnimatedSprite
	rightPent *d2gui.AnimatedSprite
	layouts   []*layout
}

type layout struct {
	*d2gui.Layout
	leftPent          *d2gui.AnimatedSprite
	rightPent         *d2gui.AnimatedSprite
	current           int
	hoverableElements []hoverableElement
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

type layoutCfg struct {
	playSound    func()
	showLayout   func(id layoutID)
	wrapLayout   func(func(*layout)) *layout
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
		mainLayoutID:              newMainLayout(cfg),
		optionsLayoutID:           newOptionsLayout(cfg),
		soundOptionsLayoutID:      newSoundOptionsLayout(cfg),
		videoOptionsLayoutID:      newVideoOptionsLayout(cfg),
		automapOptionsLayoutID:    newAutomapOptionsLayout(cfg),
		configureControlsLayoutID: newConfigureControlsLayout(cfg),
	}
	return m
}

func (m *EscapeMenu) wrapLayout(fn func(*layout)) *layout {
	wrapper := d2gui.CreateLayout(d2gui.PositionTypeHorizontal)
	wrapper.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	wrapper.AddSpacerDynamic()

	center := wrapper.AddLayout(d2gui.PositionTypeHorizontal)
	center.SetSize(menuSize, 0)

	left := center.AddLayout(d2gui.PositionTypeHorizontal)
	left.SetSize(sidePanelsSize, 0)
	leftPent, _ := left.AddAnimatedSprite(d2resource.PentSpin, d2resource.PaletteUnits, d2gui.DirectionForward)
	m.leftPent = leftPent

	// wrap the base layout so we can pass values around more easily
	base := &layout{}
	baseLayout := center.AddLayout(d2gui.PositionTypeVertical)
	baseLayout.SetHorizontalAlign(d2gui.HorizontalAlignCenter)
	base.Layout = baseLayout
	fn(base)

	right := center.AddLayout(d2gui.PositionTypeHorizontal)
	// For some reason, aligning the panel to the right won't align the pentagram, so we need to add a static spacer.
	right.AddSpacerStatic(sidePanelsSize-pentSize, 0)
	right.SetSize(sidePanelsSize, 0)
	rightPent, _ := right.AddAnimatedSprite(d2resource.PentSpin, d2resource.PaletteUnits, d2gui.DirectionBackward)
	m.rightPent = rightPent

	wrapper.AddSpacerDynamic()
	return &layout{
		Layout:            wrapper,
		leftPent:          leftPent,
		rightPent:         rightPent,
		hoverableElements: base.hoverableElements,
	}
}

func newMainLayout(cfg *layoutCfg) *layout {
	return cfg.wrapLayout(func(base *layout) {
		addBigSelectionLabel(base, cfg, "OPTIONS", optionsLayoutID)
		addBigSelectionLabel(base, cfg, "SAVE AND EXIT GAME", noLayoutID)
		addBigSelectionLabel(base, cfg, "RETURN TO GAME", noLayoutID)
	})
}

func newOptionsLayout(cfg *layoutCfg) *layout {
	return cfg.wrapLayout(func(base *layout) {
		addBigSelectionLabel(base, cfg, "SOUND OPTIONS", soundOptionsLayoutID)
		addBigSelectionLabel(base, cfg, "VIDEO OPTIONS", videoOptionsLayoutID)
		addBigSelectionLabel(base, cfg, "AUTOMAP OPTIONS", automapOptionsLayoutID)
		addBigSelectionLabel(base, cfg, "CONFIGURE CONTROLS", configureControlsLayoutID)
		addBigSelectionLabel(base, cfg, "PREVIOUS MENU", mainLayoutID)
	})
}

func newSoundOptionsLayout(cfg *layoutCfg) *layout {
	return cfg.wrapLayout(func(base *layout) {
		addTitle(base, "SOUND OPTIONS")
		addEnumLabel(base, cfg, optAudioSoundVolume, "SOUND VOLUME", []string{"TODO"})
		addEnumLabel(base, cfg, optAudioMusicVolume, "MUSIC VOLUME", []string{"TODO"})
		addEnumLabel(base, cfg, optAudio3dSound, "3D BIAS", []string{"TODO"})
		addEnumLabel(base, cfg, optAudioHardwareAcceleration, "HARDWARE ACCELERATION", []string{"ON", "OFF"})
		addEnumLabel(base, cfg, optAudioEnvEffects, "ENVIRONMENTAL EFFECTS", []string{"ON", "OFF"})
		addEnumLabel(base, cfg, optAudioNpcSpeech, "NPC SPEECH", []string{"AUDIO AND TEXT", "AUDIO ONLY", "TEXT ONLY"})
		addPreviousMenuLabel(base, cfg, optionsLayoutID)
	})
}

func newVideoOptionsLayout(cfg *layoutCfg) *layout {
	return cfg.wrapLayout(func(base *layout) {
		addTitle(base, "VIDEO OPTIONS")
		addEnumLabel(base, cfg, optVideoResolution, "VIDEO RESOLUTION", []string{"800X600", "1024X768"})
		addEnumLabel(base, cfg, optVideoLightingQuality, "LIGHTING QUALITY", []string{"LOW", "HIGH"})
		addEnumLabel(base, cfg, optVideoBlendedShadows, "BLENDED SHADOWS", []string{"ON", "OFF"})
		addEnumLabel(base, cfg, optVideoPerspective, "PERSPECTIVE", []string{"ON", "OFF"})
		addEnumLabel(base, cfg, optVideoGamma, "GAMMA", []string{"TODO"})
		addEnumLabel(base, cfg, optVideoContrast, "CONTRAST", []string{"TODO"})
		addPreviousMenuLabel(base, cfg, optionsLayoutID)
	})
}

func newAutomapOptionsLayout(cfg *layoutCfg) *layout {
	return cfg.wrapLayout(func(base *layout) {
		addTitle(base, "AUTOMAP OPTIONS")
		addEnumLabel(base, cfg, optAutomapSize, "AUTOMAP SIZE", []string{"FULL SCREEN"})
		addEnumLabel(base, cfg, optAutomapFade, "FADE", []string{"YES", "NO"})
		addEnumLabel(base, cfg, optAutomapCenterWhenCleared, "CENTER WHEN CLEARED", []string{"YES", "NO"})
		addEnumLabel(base, cfg, optAutomapShowParty, "SHOW PARTY", []string{"YES", "NO"})
		addEnumLabel(base, cfg, optAutomapShowNames, "SHOW NAMES", []string{"YES", "NO"})
		addPreviousMenuLabel(base, cfg, optionsLayoutID)
	})
}

func newConfigureControlsLayout(cfg *layoutCfg) *layout {
	return cfg.wrapLayout(func(base *layout) {
		addTitle(base, "CONFIGURE CONTROLS")
		addPreviousMenuLabel(base, cfg, optionsLayoutID)
	})
}

func addTitle(layout *layout, text string) {
	layout.AddLabel(text, d2gui.FontStyle42Units)
	layout.AddSpacerStatic(10, labelGutter)
}

func addBigSelectionLabel(layout *layout, cfg *layoutCfg, text string, targetLayout layoutID) {
	label, _ := layout.AddLabel(text, d2gui.FontStyle42Units)
	label.SetMouseClickHandler(func(_ d2input.MouseEvent) {
		cfg.showLayout(targetLayout)
	})
	label.SetMouseEnterHandler(func(_ d2input.MouseMoveEvent) {
		cfg.hoverElement(label)
	})
	layout.AddSpacerStatic(10, labelGutter)
	layout.hoverableElements = append(layout.hoverableElements, label)
}

func addPreviousMenuLabel(layout *layout, cfg *layoutCfg, targetLayout layoutID) {
	layout.AddSpacerStatic(10, labelGutter)
	label, _ := layout.AddLabel("PREVIOUS MENU", d2gui.FontStyle30Units)
	label.SetMouseClickHandler(func(_ d2input.MouseEvent) {
		cfg.showLayout(targetLayout)
	})
	label.SetMouseEnterHandler(func(_ d2input.MouseMoveEvent) {
		cfg.hoverElement(label)
	})
	layout.hoverableElements = append(layout.hoverableElements, label)
}

func addEnumLabel(layout *layout, cfg *layoutCfg, optID optionID, text string, values []string) {
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
	layout.hoverableElements = append(layout.hoverableElements, l)
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
	case soundOptionsLayoutID,
		videoOptionsLayoutID,
		automapOptionsLayoutID,
		configureControlsLayoutID:
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
	fmt.Println("hover", x, y)
	return
}

func (m *EscapeMenu) onUpdateValue(optID optionID, value string) {
	fmt.Println(fmt.Sprintf("updating value %d to %s", int(optID), value))
}

func (m *EscapeMenu) setLayout(id layoutID) {
	m.leftPent = m.layouts[id].leftPent
	m.rightPent = m.layouts[id].rightPent
	m.currentLayout = id
	m.layouts[id].current = 0
	d2gui.SetLayout(m.layouts[id].Layout)
}

func (m *EscapeMenu) OnUpKey() {
	if !m.isOpen {
		return
	}
	fmt.Println("before up", m.layouts[m.currentLayout].current, len(m.layouts[m.currentLayout].hoverableElements))
	if m.layouts[m.currentLayout].current == 0 {
		fmt.Println("return")
		return
	}
	m.layouts[m.currentLayout].current--
	fmt.Println("down", m.layouts[m.currentLayout].current)
	m.onHoverElement(m.layouts[m.currentLayout].hoverableElements[m.layouts[m.currentLayout].current])
}

func (m *EscapeMenu) OnDownKey() {
	if !m.isOpen {
		return
	}
	fmt.Println("before down", m.layouts[m.currentLayout].current, len(m.layouts[m.currentLayout].hoverableElements))
	if m.layouts[m.currentLayout].current == len(m.layouts[m.currentLayout].hoverableElements)-1 {
		fmt.Println("return")
		return
	}
	fmt.Println("down", m.layouts[m.currentLayout].current)
	m.layouts[m.currentLayout].current++
	m.onHoverElement(m.layouts[m.currentLayout].hoverableElements[m.layouts[m.currentLayout].current])
}

func (m *EscapeMenu) OnEnterKey() {
	if !m.isOpen {
		return
	}
	//m.layouts[m.currentLayout].hoverableElements[m.layouts[m.currentLayout].current]
}
