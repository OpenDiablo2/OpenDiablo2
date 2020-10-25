package d2player

import (
	"fmt"
	"log"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
)

type (
	layoutID int
	optionID int
)

const (
	// UI
	labelGutter    = 10
	sidePanelsSize = 80
	pentSize       = 54
	menuSize       = 500
	spacerWidth    = 10

	// layouts
	noLayoutID layoutID = iota - 2
	saveLayoutID
	mainLayoutID
	optionsLayoutID
	soundOptionsLayoutID
	videoOptionsLayoutID
	automapOptionsLayoutID
	configureControlsLayoutID

	// audio
	optAudioSoundVolume optionID = iota
	optAudioMusicVolume
	optAudio3dSound
	optAudioHardwareAcceleration
	optAudioEnvEffects
	optAudioNpcSpeech
	// video
	optVideoResolution
	optVideoLightingQuality
	optVideoBlendedShadows
	optVideoPerspective
	optVideoGamma
	optVideoContrast
	// automap
	optAutomapSize
	optAutomapFade
	optAutomapCenterWhenCleared
	optAutomapShowParty
	optAutomapShowNames
)

const (
	singleFrame = time.Millisecond * 16
)

// EscapeMenu represents the in-game menu that shows up when the esc key is pressed
type EscapeMenu struct {
	isOpen        bool
	selectSound   d2interface.SoundEffect
	currentLayout layoutID

	// leftPent and rightPent are generated once and shared between the layouts
	leftPent  *d2gui.AnimatedSprite
	rightPent *d2gui.AnimatedSprite
	layouts   []*layout

	renderer      d2interface.Renderer
	audioProvider d2interface.AudioProvider
	navigator     d2interface.Navigator
	guiManager    *d2gui.GuiManager
	assetManager  *d2asset.AssetManager
}

type layout struct {
	*d2gui.Layout
	leftPent           *d2gui.AnimatedSprite
	rightPent          *d2gui.AnimatedSprite
	currentEl          int
	rendered           bool
	actionableElements []actionableElement
}

func (l *layout) Trigger() {
	// noop
}

type showLayoutLabel struct {
	*d2gui.Label
	target     layoutID
	showLayout func(id layoutID)
}

func (l *showLayoutLabel) Trigger() {
	l.showLayout(l.target)
}

type enumLabel struct {
	*d2gui.Layout
	textChangingLabel *d2gui.Label
	optionID          optionID
	values            []string
	current           int
	playSound         func()
	updateValue       func(optID optionID, value string)
}

func (l *enumLabel) Trigger() {
	l.playSound()
	next := (l.current + 1) % len(l.values)
	l.current = next

	currentValue := l.values[l.current]
	if err := l.textChangingLabel.SetText(currentValue); err != nil {
		fmt.Printf("could not change the label text to: %s\n", currentValue)
	}

	l.updateValue(l.optionID, currentValue)
}

type actionableElement interface {
	GetOffset() (int, int)
	Trigger()
}

// NewEscapeMenu creates a new escape menu
func NewEscapeMenu(navigator d2interface.Navigator,
	renderer d2interface.Renderer,
	audioProvider d2interface.AudioProvider,
	guiManager *d2gui.GuiManager,
	assetManager *d2asset.AssetManager,
) *EscapeMenu {
	m := &EscapeMenu{
		audioProvider: audioProvider,
		renderer:      renderer,
		navigator:     navigator,
		guiManager:    guiManager,
		assetManager:  assetManager,
	}

	m.layouts = []*layout{
		mainLayoutID:              m.newMainLayout(),
		optionsLayoutID:           m.newOptionsLayout(),
		soundOptionsLayoutID:      m.newSoundOptionsLayout(),
		videoOptionsLayoutID:      m.newVideoOptionsLayout(),
		automapOptionsLayoutID:    m.newAutomapOptionsLayout(),
		configureControlsLayoutID: m.newConfigureControlsLayout(),
	}

	return m
}

func (m *EscapeMenu) newMainLayout() *layout {
	return m.wrapLayout(func(l *layout) {
		m.addBigSelectionLabel(l, "OPTIONS", optionsLayoutID)
		m.addBigSelectionLabel(l, "SAVE AND EXIT GAME", saveLayoutID)
		m.addBigSelectionLabel(l, "RETURN TO GAME", noLayoutID)
	})
}

func (m *EscapeMenu) newOptionsLayout() *layout {
	return m.wrapLayout(func(l *layout) {
		m.addBigSelectionLabel(l, "SOUND OPTIONS", soundOptionsLayoutID)
		m.addBigSelectionLabel(l, "VIDEO OPTIONS", videoOptionsLayoutID)
		m.addBigSelectionLabel(l, "AUTOMAP OPTIONS", automapOptionsLayoutID)
		m.addBigSelectionLabel(l, "CONFIGURE CONTROLS", configureControlsLayoutID)
		m.addBigSelectionLabel(l, "PREVIOUS MENU", mainLayoutID)
	})
}

func (m *EscapeMenu) newSoundOptionsLayout() *layout {
	return m.wrapLayout(func(l *layout) {
		m.addTitle(l, "SOUND OPTIONS")
		m.addEnumLabel(l, optAudioSoundVolume, "SOUND VOLUME", []string{"TODO"})
		m.addEnumLabel(l, optAudioMusicVolume, "MUSIC VOLUME", []string{"TODO"})
		m.addEnumLabel(l, optAudio3dSound, "3D BIAS", []string{"TODO"})
		m.addEnumLabel(l, optAudioHardwareAcceleration, "HARDWARE ACCELERATION", []string{"ON", "OFF"})
		m.addEnumLabel(l, optAudioEnvEffects, "ENVIRONMENTAL EFFECTS", []string{"ON", "OFF"})
		m.addEnumLabel(l, optAudioNpcSpeech, "NPC SPEECH", []string{"AUDIO AND TEXT", "AUDIO ONLY", "TEXT ONLY"})
		m.addPreviousMenuLabel(l)
	})
}

func (m *EscapeMenu) newVideoOptionsLayout() *layout {
	return m.wrapLayout(func(l *layout) {
		m.addTitle(l, "VIDEO OPTIONS")
		m.addEnumLabel(l, optVideoResolution, "VIDEO RESOLUTION", []string{"800X600", "1024X768"})
		m.addEnumLabel(l, optVideoLightingQuality, "LIGHTING QUALITY", []string{"LOW", "HIGH"})
		m.addEnumLabel(l, optVideoBlendedShadows, "BLENDED SHADOWS", []string{"ON", "OFF"})
		m.addEnumLabel(l, optVideoPerspective, "PERSPECTIVE", []string{"ON", "OFF"})
		m.addEnumLabel(l, optVideoGamma, "GAMMA", []string{"TODO"})
		m.addEnumLabel(l, optVideoContrast, "CONTRAST", []string{"TODO"})
		m.addPreviousMenuLabel(l)
	})
}

func (m *EscapeMenu) newAutomapOptionsLayout() *layout {
	return m.wrapLayout(func(l *layout) {
		m.addTitle(l, "AUTOMAP OPTIONS")
		m.addEnumLabel(l, optAutomapSize, "AUTOMAP SIZE", []string{"FULL SCREEN"})
		m.addEnumLabel(l, optAutomapFade, "FADE", []string{"YES", "NO"})
		m.addEnumLabel(l, optAutomapCenterWhenCleared, "CENTER WHEN CLEARED", []string{"YES", "NO"})
		m.addEnumLabel(l, optAutomapShowParty, "SHOW PARTY", []string{"YES", "NO"})
		m.addEnumLabel(l, optAutomapShowNames, "SHOW NAMES", []string{"YES", "NO"})
		m.addPreviousMenuLabel(l)
	})
}

func (m *EscapeMenu) newConfigureControlsLayout() *layout {
	return m.wrapLayout(func(l *layout) {
		m.addTitle(l, "CONFIGURE CONTROLS")
		m.addPreviousMenuLabel(l)
	})
}

func (m *EscapeMenu) wrapLayout(fn func(*layout)) *layout {
	wrapper := d2gui.CreateLayout(m.renderer, d2gui.PositionTypeHorizontal, m.assetManager)
	wrapper.SetVerticalAlign(d2gui.VerticalAlignMiddle)
	wrapper.AddSpacerDynamic()

	center := wrapper.AddLayout(d2gui.PositionTypeHorizontal)
	center.SetSize(menuSize, 0)

	left := center.AddLayout(d2gui.PositionTypeVertical)
	left.SetSize(sidePanelsSize, pentSize)
	leftPent, err := left.AddAnimatedSprite(d2resource.PentSpin, d2resource.PaletteUnits, d2gui.DirectionBackward)

	if err != nil {
		log.Print(err)
		return nil
	}

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
	right.SetSize(sidePanelsSize, pentSize)
	rightPent, err := right.AddAnimatedSprite(d2resource.PentSpin, d2resource.PaletteUnits, d2gui.DirectionForward)

	if err != nil {
		log.Print(err)
		return nil
	}

	m.rightPent = rightPent

	wrapper.AddSpacerDynamic()

	return &layout{
		Layout:             wrapper,
		leftPent:           leftPent,
		rightPent:          rightPent,
		actionableElements: base.actionableElements,
	}
}

func (m *EscapeMenu) addTitle(l *layout, text string) {
	_, err := l.AddLabel(text, d2gui.FontStyle42Units)
	if err != nil {
		fmt.Printf("could not add label: %s to the escape menu\n", text)
	}

	l.AddSpacerStatic(spacerWidth, labelGutter)
}

func (m *EscapeMenu) addBigSelectionLabel(l *layout, text string, targetLayout layoutID) {
	guiLabel, err := l.AddLabel(text, d2gui.FontStyle42Units)
	if err != nil {
		log.Print(err)
	}

	label := &showLayoutLabel{Label: guiLabel, target: targetLayout, showLayout: m.showLayout}
	label.SetMouseClickHandler(func(_ d2interface.MouseEvent) {
		label.Trigger()
	})

	elID := len(l.actionableElements)

	label.SetMouseEnterHandler(func(_ d2interface.MouseMoveEvent) {
		m.onHoverElement(elID)
	})

	l.AddSpacerStatic(spacerWidth, labelGutter)
	l.actionableElements = append(l.actionableElements, label)
}

func (m *EscapeMenu) addPreviousMenuLabel(l *layout) {
	l.AddSpacerStatic(spacerWidth, labelGutter)

	guiLabel, err := l.AddLabel("PREVIOUS MENU", d2gui.FontStyle30Units)
	if err != nil {
		log.Print(err)
	}

	label := &showLayoutLabel{Label: guiLabel, target: optionsLayoutID, showLayout: m.showLayout}
	label.SetMouseClickHandler(func(_ d2interface.MouseEvent) {
		label.Trigger()
	})

	elID := len(l.actionableElements)

	label.SetMouseEnterHandler(func(_ d2interface.MouseMoveEvent) {
		m.onHoverElement(elID)
	})

	l.actionableElements = append(l.actionableElements, label)
}

func (m *EscapeMenu) addEnumLabel(l *layout, optID optionID, text string, values []string) {
	guiLayout := l.AddLayout(d2gui.PositionTypeHorizontal)
	layout := &layout{Layout: guiLayout}
	layout.SetSize(menuSize, 0)

	_, err := layout.AddLabel(text, d2gui.FontStyle30Units)
	if err != nil {
		fmt.Printf("could not add label: %s to the escape menu\n", text)
	}

	elID := len(l.actionableElements)

	layout.SetMouseEnterHandler(func(_ d2interface.MouseMoveEvent) {
		m.onHoverElement(elID)
	})

	layout.AddSpacerDynamic()

	guiLabel, err := layout.AddLabel(values[0], d2gui.FontStyle30Units)
	if err != nil {
		log.Print(err)
	}

	label := &enumLabel{
		Layout:            guiLayout,
		textChangingLabel: guiLabel,
		optionID:          optID,
		values:            values,
		current:           0,
		playSound:         m.playSound,
		updateValue:       m.onUpdateValue,
	}

	layout.SetMouseClickHandler(func(_ d2interface.MouseEvent) {
		label.Trigger()
	})
	l.AddSpacerStatic(spacerWidth, labelGutter)

	l.actionableElements = append(l.actionableElements, label)
}

// OnLoad loads the necessary files for the escape menu
func (m *EscapeMenu) OnLoad() {
	var err error

	m.selectSound, err = m.audioProvider.LoadSound(d2resource.SFXCursorSelect, false, false)
	if err != nil {
		log.Print(err)
	}
}

// OnEscKey is called when the escape key is pressed
func (m *EscapeMenu) OnEscKey() {
	// note: original D2 returns straight to the game from however deep in the menu we are
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

	m.close()
}

func (m *EscapeMenu) close() {
	m.isOpen = false

	m.guiManager.SetLayout(nil)
}

func (m *EscapeMenu) open() {
	m.isOpen = true
	m.setLayout(mainLayoutID)
}

func (m *EscapeMenu) playSound() {
	m.selectSound.Play()
}

func (m *EscapeMenu) showLayout(id layoutID) {
	m.playSound()

	if id == noLayoutID {
		m.close()
		return
	}

	if id == saveLayoutID {
		m.navigator.ToMainMenu()
		return
	}

	m.setLayout(id)
}

func (m *EscapeMenu) onHoverElement(id int) {
	_, y := m.layouts[m.currentLayout].actionableElements[id].GetOffset()
	m.layouts[m.currentLayout].currentEl = id

	x, _ := m.leftPent.GetPosition()
	m.leftPent.SetPosition(x, y+spacerWidth)

	x, _ = m.rightPent.GetPosition()
	m.rightPent.SetPosition(x, y+spacerWidth)
}

func (m *EscapeMenu) onUpdateValue(optID optionID, value string) {
	fmt.Printf("updating value %d with %s\n", optID, value)
}

func (m *EscapeMenu) setLayout(id layoutID) {
	m.leftPent = m.layouts[id].leftPent
	m.rightPent = m.layouts[id].rightPent
	m.currentLayout = id
	m.layouts[id].currentEl = len(m.layouts[id].actionableElements) - 1 // default to Previous Menu

	m.guiManager.SetLayout(m.layouts[id].Layout)

	// when first rendering a layout, widgets don't have offsets so we hide pentagrams for a frame
	if !m.layouts[id].rendered {
		m.layouts[id].rendered = true
		m.leftPent.SetVisible(false)
		m.rightPent.SetVisible(false)

		go func() {
			time.Sleep(singleFrame)
			m.onHoverElement(m.layouts[id].currentEl)
			m.leftPent.SetVisible(true)
			m.rightPent.SetVisible(true)
		}()
	} else {
		m.onHoverElement(m.layouts[id].currentEl)
	}
}

func (m *EscapeMenu) onUpKey() {
	if !m.isOpen {
		return
	}

	if m.layouts[m.currentLayout].currentEl == 0 {
		return
	}

	m.layouts[m.currentLayout].currentEl--
	m.onHoverElement(m.layouts[m.currentLayout].currentEl)
}

func (m *EscapeMenu) onDownKey() {
	if !m.isOpen {
		return
	}

	if m.layouts[m.currentLayout].currentEl == len(m.layouts[m.currentLayout].actionableElements)-1 {
		return
	}

	m.layouts[m.currentLayout].currentEl++
	m.onHoverElement(m.layouts[m.currentLayout].currentEl)
}

func (m *EscapeMenu) onEnterKey() {
	if !m.isOpen {
		return
	}

	m.layouts[m.currentLayout].actionableElements[m.layouts[m.currentLayout].currentEl].Trigger()
}

// IsOpen returns whether the escape menu is open (visible) or not
func (m *EscapeMenu) IsOpen() bool {
	return m.isOpen
}

// OnKeyDown defines the actions of the Escape Menu when a key is pressed
func (m *EscapeMenu) OnKeyDown(event d2interface.KeyEvent) bool {
	switch event.Key() {
	case d2enum.KeyUp:
		m.onUpKey()
	case d2enum.KeyDown:
		m.onDownKey()
	case d2enum.KeyEnter:
		m.onEnterKey()
	default:
		return false
	}

	return true
}
