package d2systems

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
	ebiten_input "github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/ebiten"
)

const (
	logPrefixInputSystem = "Input System"
)

// static check that InputSystem implements the System interface
var _ akara.System = &InputSystem{}

// InputSystem is responsible for handling interactive entities
type InputSystem struct {
	akara.BaseSubscriberSystem
	*d2util.Logger
	d2interface.InputService
	configs      *akara.Subscription
	interactives *akara.Subscription
	inputState *d2input.InputVector
	Components struct {
		GameConfig d2components.GameConfigFactory
		Interactive d2components.InteractiveFactory
	}
}

// Init initializes the system with the given world, injecting the necessary components
func (m *InputSystem) Init(world *akara.World) {
	m.World = world

	m.InputService = ebiten_input.InputService{}

	m.setupLogger()

	m.Debug("initializing ...")

	m.setupFactories()
	m.setupSubscriptions()

	m.inputState = d2input.NewInputVector()
}

func (m *InputSystem) setupLogger() {
	m.Logger = d2util.NewLogger()
	m.SetPrefix(logPrefixInputSystem)
}

func (m *InputSystem) setupFactories() {
	m.Debug("setting up component factories")

	m.InjectComponent(&d2components.GameConfig{}, &m.Components.GameConfig.ComponentFactory)
	m.InjectComponent(&d2components.Interactive{}, &m.Components.Interactive.ComponentFactory)
}

func (m *InputSystem) setupSubscriptions() {
	m.Debug("setting up component subscriptions")

	interactives := m.NewComponentFilter().
		Require(&d2components.Interactive{}).
		Build()

	gameConfigs := m.NewComponentFilter().
		Require(&d2components.GameConfig{}).
		Build()

	m.interactives = m.AddSubscription(interactives)
	m.configs = m.AddSubscription(gameConfigs)
}

// Update will iterate over interactive entities
func (m *InputSystem) Update() {
	m.updateInputState()

	for _, id := range m.interactives.GetEntities() {
		preventPropagation := m.applyInputState(id)
		if preventPropagation {
			break
		}
	}
}

func (m *InputSystem) updateInputState() {
	m.inputState.Clear()

	var keysToCheck = []d2input.Key{
		d2input.Key0, d2input.Key1, d2input.Key2, d2input.Key3, d2input.Key4, d2input.Key5, d2input.Key6,
		d2input.Key7, d2input.Key8, d2input.Key9, d2input.KeyA, d2input.KeyB, d2input.KeyC, d2input.KeyD,
		d2input.KeyE, d2input.KeyF, d2input.KeyG, d2input.KeyH, d2input.KeyI, d2input.KeyJ, d2input.KeyK,
		d2input.KeyL, d2input.KeyM, d2input.KeyN, d2input.KeyO, d2input.KeyP, d2input.KeyQ, d2input.KeyR,
		d2input.KeyS, d2input.KeyT, d2input.KeyU, d2input.KeyV, d2input.KeyW, d2input.KeyX, d2input.KeyY,
		d2input.KeyZ, d2input.KeyApostrophe, d2input.KeyBackslash, d2input.KeyBackspace,
		d2input.KeyCapsLock, d2input.KeyComma, d2input.KeyDelete, d2input.KeyDown,
		d2input.KeyEnd, d2input.KeyEnter, d2input.KeyEqual, d2input.KeyEscape,
		d2input.KeyF1, d2input.KeyF2, d2input.KeyF3, d2input.KeyF4, d2input.KeyF5, d2input.KeyF6,
		d2input.KeyF7, d2input.KeyF8, d2input.KeyF9, d2input.KeyF10, d2input.KeyF11, d2input.KeyF12,
		d2input.KeyGraveAccent, d2input.KeyHome, d2input.KeyInsert, d2input.KeyKP0,
		d2input.KeyKP1, d2input.KeyKP2, d2input.KeyKP3, d2input.KeyKP4, d2input.KeyKP5,
		d2input.KeyKP6, d2input.KeyKP7, d2input.KeyKP8, d2input.KeyKP9,
		d2input.KeyKPAdd, d2input.KeyKPDecimal, d2input.KeyKPDivide, d2input.KeyKPEnter,
		d2input.KeyKPEqual, d2input.KeyKPMultiply, d2input.KeyKPSubtract, d2input.KeyLeft,
		d2input.KeyLeftBracket, d2input.KeyMenu, d2input.KeyMinus, d2input.KeyNumLock,
		d2input.KeyPageDown, d2input.KeyPageUp, d2input.KeyPause, d2input.KeyPeriod,
		d2input.KeyPrintScreen, d2input.KeyRight, d2input.KeyRightBracket,
		d2input.KeyScrollLock, d2input.KeySemicolon, d2input.KeySlash,
		d2input.KeySpace, d2input.KeyTab, d2input.KeyUp,
	}

	var modifiersToCheck = []d2input.Modifier{
		d2input.ModAlt, d2input.ModControl, d2input.ModShift,
	}

	var buttonsToCheck = []d2input.MouseButton{
		d2input.MouseButtonLeft, d2input.MouseButtonMiddle, d2input.MouseButtonRight,
	}

	for _, key := range keysToCheck {
		truth := m.InputService.IsKeyJustPressed(d2enum.Key(key))
		m.inputState.KeyVector.Set(key, truth)
	}

	for _, mod := range modifiersToCheck {
		truth := m.InputService.IsKeyJustPressed(d2enum.Key(mod))
		m.inputState.ModifierVector.Set(mod, truth)
	}

	for _, btn := range buttonsToCheck {
		truth := m.InputService.IsMouseButtonJustPressed(d2enum.MouseButton(btn))
		m.inputState.MouseButtonVector.Set(btn, truth)
	}
}

func (m *InputSystem) applyInputState(id akara.EID) (preventPropagation bool) {
	v, found := m.Components.Interactive.Get(id)
	if !found {
		return false
	}

	if !v.Enabled || !m.inputState.Contains(v.InputVector) {
		return false
	}

	return v.Callback()
}
