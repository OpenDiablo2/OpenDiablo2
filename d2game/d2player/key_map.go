package d2player

import (
	"sync"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// KeyMap represents the key mappings of the game. Each game event
// can be associated to 2 different keys. A key of -1 means none
type KeyMap struct {
	mutex    sync.RWMutex
	mapping  map[d2enum.Key]d2enum.GameEvent
	controls map[d2enum.GameEvent]*KeyBinding
}

type KeyBindingType int

const (
	KeyBindingTypePrimary KeyBindingType = iota
	KeyBindingTypeSecondary
)

// NewKeyMap returns a new instance of a KeyMap
func NewKeyMap() *KeyMap {
	return &KeyMap{
		mapping:  make(map[d2enum.Key]d2enum.GameEvent),
		controls: make(map[d2enum.GameEvent]*KeyBinding),
	}
}

func (km *KeyMap) checkOverwrite(key d2enum.Key) (*KeyBinding, KeyBindingType) {
	var overwrittenBinding *KeyBinding
	var overwrittenBindingType KeyBindingType

	for _, binding := range km.controls {
		if binding.Primary == key {
			binding.Primary = -1
			overwrittenBinding = binding
			overwrittenBindingType = KeyBindingTypePrimary
		}
		if binding.Secondary == key {
			binding.Secondary = -1
			overwrittenBinding = binding
			overwrittenBindingType = KeyBindingTypeSecondary
		}
	}

	return overwrittenBinding, overwrittenBindingType
}

// SetPrimaryBinding binds the first key for gameEvent
func (km *KeyMap) SetPrimaryBinding(gameEvent d2enum.GameEvent, key d2enum.Key) (*KeyBinding, KeyBindingType) {
	if key == d2enum.KeyEscape {
		return nil, -1
	}

	km.mutex.Lock()
	defer km.mutex.Unlock()

	if km.controls[gameEvent] == nil {
		km.controls[gameEvent] = &KeyBinding{}
	}

	overwrittenBinding, overwrittenBindingType := km.checkOverwrite(key)

	currentKey := km.controls[gameEvent].Primary
	delete(km.mapping, currentKey)
	km.mapping[key] = gameEvent

	km.controls[gameEvent].Primary = key

	return overwrittenBinding, overwrittenBindingType
}

// SetSecondaryBinding binds the second key for gameEvent
func (km *KeyMap) SetSecondaryBinding(gameEvent d2enum.GameEvent, key d2enum.Key) (*KeyBinding, KeyBindingType) {
	if key == d2enum.KeyEscape {
		return nil, -1
	}

	km.mutex.Lock()
	defer km.mutex.Unlock()

	if km.controls[gameEvent] == nil {
		km.controls[gameEvent] = &KeyBinding{}
	}
	overwrittenBinding, overwrittenBindingType := km.checkOverwrite(key)

	currentKey := km.controls[gameEvent].Secondary
	delete(km.mapping, currentKey)
	km.mapping[key] = gameEvent

	if km.controls[gameEvent].Primary == key {
		km.controls[gameEvent].Primary = d2enum.Key(-1)
	}

	km.controls[gameEvent].Secondary = key

	return overwrittenBinding, overwrittenBindingType
}

func (km *KeyMap) getGameEvent(key d2enum.Key) d2enum.GameEvent {
	km.mutex.RLock()
	defer km.mutex.RUnlock()

	return km.mapping[key]
}

// GetKeysForGameEvent returns the bindings for a givent game event
func (km *KeyMap) GetKeysForGameEvent(gameEvent d2enum.GameEvent) *KeyBinding {
	km.mutex.RLock()
	defer km.mutex.RUnlock()

	return km.controls[gameEvent]
}

// GetBindingByKey returns the bindings for a givent game event
func (km *KeyMap) GetBindingByKey(key d2enum.Key) (*KeyBinding, d2enum.GameEvent, KeyBindingType) {
	km.mutex.RLock()
	defer km.mutex.RUnlock()

	for gameEvent, binding := range km.controls {
		if binding.Primary == key {
			return binding, gameEvent, KeyBindingTypePrimary
		}
		if binding.Secondary == key {
			return binding, gameEvent, KeyBindingTypeSecondary
		}
	}

	return nil, -1, -1
}

// KeyBinding holds the primary and secondary keys assigned to a GameEvent
type KeyBinding struct {
	Primary   d2enum.Key
	Secondary d2enum.Key
}

func (b KeyBinding) IsEmpty() bool {
	return b.Primary == -1 && b.Secondary == -1
}

func GetDefaultKeyMap() *KeyMap {
	keyMap := NewKeyMap()

	defaultControls := map[d2enum.GameEvent]KeyBinding{
		d2enum.ToggleCharacterPanel: {d2enum.KeyA, d2enum.KeyC},
		d2enum.ToggleInventoryPanel: {d2enum.KeyB, d2enum.KeyI},
		d2enum.TogglePartyPanel:     {d2enum.KeyP, -1},
		d2enum.ToggleHirelingPanel:  {d2enum.KeyO, -1},
		d2enum.ToggleMessageLog:     {d2enum.KeyM, -1},
		d2enum.ToggleQuestLog:       {d2enum.KeyQ, -1},
		d2enum.ToggleHelpScreen:     {d2enum.KeyH, -1},

		d2enum.ToggleSkillTreePanel:     {d2enum.KeyT, -1},
		d2enum.ToggleRightSkillSelector: {d2enum.KeyS, -1},
		d2enum.UseSkill1:                {d2enum.KeyF1, -1},
		d2enum.UseSkill2:                {d2enum.KeyF2, -1},
		d2enum.UseSkill3:                {d2enum.KeyF3, -1},
		d2enum.UseSkill4:                {d2enum.KeyF4, -1},
		d2enum.UseSkill5:                {d2enum.KeyF5, -1},
		d2enum.UseSkill6:                {d2enum.KeyF6, -1},
		d2enum.UseSkill7:                {d2enum.KeyF7, -1},
		d2enum.UseSkill8:                {d2enum.KeyF8, -1},
		d2enum.UseSkill9:                {-1, -1},
		d2enum.UseSkill10:               {-1, -1},
		d2enum.UseSkill11:               {-1, -1},
		d2enum.UseSkill12:               {-1, -1},
		d2enum.UseSkill13:               {-1, -1},
		d2enum.UseSkill14:               {-1, -1},
		d2enum.UseSkill15:               {-1, -1},
		d2enum.UseSkill16:               {-1, -1},
		d2enum.SelectPreviousSkill:      {d2enum.KeyMouseWheelUp, -1},
		d2enum.SelectNextSkill:          {d2enum.KeyMouseWheelDown, -1},

		d2enum.ToggleBelts:  {d2enum.KeyTilde, -1},
		d2enum.UseBeltSlot1: {d2enum.Key1, -1},
		d2enum.UseBeltSlot2: {d2enum.Key2, -1},
		d2enum.UseBeltSlot3: {d2enum.Key3, -1},
		d2enum.UseBeltSlot4: {d2enum.Key4, -1},
		d2enum.SwapWeapons:  {d2enum.KeyW, -1},

		d2enum.ToggleChatOverlay:   {d2enum.KeyEnter, -1},
		d2enum.HoldRun:             {d2enum.KeyControl, -1},
		d2enum.ToggleRunWalk:       {d2enum.KeyR, -1},
		d2enum.HoldStandStill:      {d2enum.KeyShift, -1},
		d2enum.HoldShowGroundItems: {d2enum.KeyAlt, -1},
		d2enum.HoldShowPortraits:   {d2enum.KeyZ, -1},

		d2enum.ToggleAutomap:        {d2enum.KeyTab, -1},
		d2enum.CenterAutomap:        {d2enum.KeyHome, -1},
		d2enum.TogglePartyOnAutomap: {d2enum.KeyF11, -1},
		d2enum.ToggleNamesOnAutomap: {d2enum.KeyF12, -1},
		d2enum.ToggleMiniMap:        {d2enum.KeyV, -1},

		d2enum.SayHelp:         {d2enum.KeyKP0, -1},
		d2enum.SayFollowMe:     {d2enum.KeyKP1, -1},
		d2enum.SayThisIsForYou: {d2enum.KeyKP2, -1},
		d2enum.SayThanks:       {d2enum.KeyKP3, -1},
		d2enum.SaySorry:        {d2enum.KeyKP4, -1},
		d2enum.SayBye:          {d2enum.KeyKP5, -1},
		d2enum.SayNowYouDie:    {d2enum.KeyKP6, -1},
		d2enum.SayRetreat:      {d2enum.KeyKP7, -1},

		d2enum.TakeScreenShot: {d2enum.KeyPrintScreen, -1},
		d2enum.ClearScreen:    {d2enum.KeySpace, -1},
		d2enum.ClearMessages:  {d2enum.KeyN, -1},
	}
	for gameEvent, keys := range defaultControls {
		keyMap.SetPrimaryBinding(gameEvent, keys.Primary)
		keyMap.SetSecondaryBinding(gameEvent, keys.Secondary)
	}

	return keyMap
}
