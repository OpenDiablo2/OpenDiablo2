package d2player

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type bindingChange struct {
	target    *KeyBinding
	primary   d2enum.Key
	secondary d2enum.Key
}

// KeyBindingMenu represents the menu to view/edit the
// key bindings
type KeyBindingMenu struct {
	*Box

	asset      *d2asset.AssetManager
	renderer   d2interface.Renderer
	ui         *d2ui.UIManager
	guiManager *d2gui.GuiManager
	keyMap     *KeyMap
	escapeMenu *EscapeMenu

	mainLayout       *d2gui.Layout
	contentLayout    *d2gui.Layout
	scrollbar        *LayoutScrollbar
	bindingLayouts   []*bindingLayout
	changesToBeSaved map[d2enum.GameEvent]*bindingChange

	isAwaitingKeyDown          bool
	currentBindingModifierType KeyBindingType
	currentBindingModifier     d2enum.GameEvent
	currentBindingLayout       *bindingLayout
	lastBindingLayout          *bindingLayout
}

type bindingLayout struct {
	wrapperLayout   *d2gui.Layout
	descLayout      *d2gui.Layout
	descLabel       *d2gui.Label
	primaryLayout   *d2gui.Layout
	primaryLabel    *d2gui.Label
	secondaryLayout *d2gui.Layout
	secondaryLabel  *d2gui.Label

	binding   *KeyBinding
	gameEvent d2enum.GameEvent
}

func (l *bindingLayout) Reset() {
	l.descLabel.SetIsHovered(false)
	l.primaryLabel.SetIsHovered(false)
	l.secondaryLabel.SetIsHovered(false)
	l.primaryLabel.SetIsBlinking(false)
	l.secondaryLabel.SetIsBlinking(false)
}

func (l *bindingLayout) isInLayoutRect(x, y int, targetLayout *d2gui.Layout) bool {
	targetW, targetH := targetLayout.GetSize()
	targetX, targetY := targetLayout.Sx, targetLayout.Sy

	if x >= targetX && x <= targetX+targetW && y >= targetY && y <= targetY+targetH {
		return true
	}

	return false
}

func (l *bindingLayout) GetPointedLayoutAndLabel(x, y int) (d2enum.GameEvent, KeyBindingType) {
	if l.isInLayoutRect(x, y, l.descLayout) {
		return l.gameEvent, KeyBindingTypePrimary
	}

	if l.primaryLayout != nil {
		if l.isInLayoutRect(x, y, l.primaryLayout) {
			return l.gameEvent, KeyBindingTypePrimary
		}
	}

	if l.secondaryLayout != nil {
		if l.isInLayoutRect(x, y, l.secondaryLayout) {
			return l.gameEvent, KeyBindingTypeSecondary
		}
	}

	return -1, -1
}

func NewKeyBindingMenu(
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	ui *d2ui.UIManager,
	guiManager *d2gui.GuiManager,
	keyMap *KeyMap,
	escapeMenu *EscapeMenu,
) *KeyBindingMenu {
	mainLayout := d2gui.CreateLayout(renderer, d2gui.PositionTypeAbsolute, asset)
	contentLayout := mainLayout.AddLayout(d2gui.PositionTypeAbsolute)

	ret := &KeyBindingMenu{
		keyMap:           keyMap,
		asset:            asset,
		ui:               ui,
		guiManager:       guiManager,
		renderer:         renderer,
		mainLayout:       mainLayout,
		contentLayout:    contentLayout,
		bindingLayouts:   []*bindingLayout{},
		changesToBeSaved: make(map[d2enum.GameEvent]*bindingChange),
		escapeMenu:       escapeMenu,
	}

	ret.Box = NewBox(asset, renderer, ui, guiManager, ret.mainLayout, 620, 375, 90, 65, "")
	ret.Box.SetPadding(19, 14)
	ret.Box.SetOptions([]*LabelButton{
		NewLabelButton(0, 0, "Cancel", d2util.Color(0xD03C39FF), func() { ret.onCancelClicked() }),
		NewLabelButton(0, 0, "Default", d2util.Color(0x5450D1FF), func() { ret.onDefaultClicked() }),
		NewLabelButton(0, 0, "Accept", d2util.Color(0x00D000FF), func() { ret.onAcceptClicked() }),
	})

	return ret
}

func (menu *KeyBindingMenu) Close() {
	menu.Box.Close()
	menu.currentBindingLayout = nil
	menu.currentBindingModifier = -1
	menu.currentBindingModifierType = -1
}

func (menu *KeyBindingMenu) Load() error {
	menu.Box.Load()

	mainLayoutW, mainLayoutH := menu.mainLayout.GetSize()

	headerLayout := menu.contentLayout.AddLayout(d2gui.PositionTypeHorizontal)
	headerLayout.SetSize(mainLayoutW, 24)
	headerLayout.AddLabelWithColor(
		menu.asset.TranslateString("CfgFunction"),
		d2gui.FontStyleFormal11Units,
		d2util.Color(0xA1925DFF),
	)
	headerLayout.AddSpacerStatic(131, 1)
	headerLayout.AddLabelWithColor(
		menu.asset.TranslateString("CfgPrimaryKey"),
		d2gui.FontStyleFormal11Units,
		d2util.Color(0xA1925DFF),
	)
	headerLayout.AddSpacerStatic(86, 1)
	headerLayout.AddLabelWithColor(
		menu.asset.TranslateString("CfgSecondaryKey"),
		d2gui.FontStyleFormal11Units,
		d2util.Color(0xA1925DFF),
	)
	headerLayout.SetVerticalAlign(d2gui.VerticalAlignMiddle)

	bindingWrapper := menu.contentLayout.AddLayout(d2gui.PositionTypeAbsolute)
	bindingWrapper.SetPosition(0, 24)
	bindingWrapper.SetSize(mainLayoutW, mainLayoutH-24)
	bindingLayout := menu.generateLayout()

	menu.Box.layout.AdjustEntryPlacement()
	menu.mainLayout.AdjustEntryPlacement()
	menu.contentLayout.AdjustEntryPlacement()

	menu.scrollbar = newLayoutScrollbar(bindingWrapper, bindingLayout)
	menu.scrollbar.Load(menu.ui)

	bindingWrapper.AddLayoutFromSource(bindingLayout)

	bindingWrapper.AdjustEntryPlacement()

	return nil
}

type keyBindingSetting struct {
	label     string
	gameEvent d2enum.GameEvent
}

func (menu *KeyBindingMenu) generateLayout() *d2gui.Layout {
	groups := [][]keyBindingSetting{
		{
			{menu.asset.TranslateString("CfgCharacter"), d2enum.ToggleCharacterPanel},
			{menu.asset.TranslateString("CfgInventory"), d2enum.ToggleInventoryPanel},
			{menu.asset.TranslateString("CfgParty"), d2enum.TogglePartyPanel},
			{menu.asset.TranslateString("Cfghireling"), d2enum.ToggleHirelingPanel},
			{menu.asset.TranslateString("CfgMessageLog"), d2enum.ToggleMessageLog},
			{menu.asset.TranslateString("CfgQuestLog"), d2enum.ToggleQuestLog},
			{menu.asset.TranslateString("CfgHelp"), d2enum.ToggleHelpScreen},
		},
		{
			{menu.asset.TranslateString("CfgSkillTree"), d2enum.ToggleSkillTreePanel},
			{menu.asset.TranslateString("CfgSkillPick"), d2enum.ToggleRightSkillSelector},
			{menu.asset.TranslateString("CfgSkill1"), d2enum.UseSkill1},
			{menu.asset.TranslateString("CfgSkill2"), d2enum.UseSkill2},
			{menu.asset.TranslateString("CfgSkill3"), d2enum.UseSkill3},
			{menu.asset.TranslateString("CfgSkill4"), d2enum.UseSkill4},
			{menu.asset.TranslateString("CfgSkill5"), d2enum.UseSkill5},
			{menu.asset.TranslateString("CfgSkill6"), d2enum.UseSkill6},
			{menu.asset.TranslateString("CfgSkill7"), d2enum.UseSkill7},
			{menu.asset.TranslateString("CfgSkill8"), d2enum.UseSkill8},
			{menu.asset.TranslateString("CfgSkill9"), d2enum.UseSkill9},
			{menu.asset.TranslateString("CfgSkill10"), d2enum.UseSkill10},
			{menu.asset.TranslateString("CfgSkill11"), d2enum.UseSkill11},
			{menu.asset.TranslateString("CfgSkill12"), d2enum.UseSkill12},
			{menu.asset.TranslateString("CfgSkill13"), d2enum.UseSkill13},
			{menu.asset.TranslateString("CfgSkill14"), d2enum.UseSkill14},
			{menu.asset.TranslateString("CfgSkill15"), d2enum.UseSkill15},
			{menu.asset.TranslateString("CfgSkill16"), d2enum.UseSkill16},
			{menu.asset.TranslateString("Cfgskillup"), d2enum.SelectPreviousSkill},
			{menu.asset.TranslateString("Cfgskilldown"), d2enum.SelectNextSkill},
		},
		{
			{menu.asset.TranslateString("CfgBeltShow"), d2enum.ToggleBelts},
			{menu.asset.TranslateString("CfgBelt1"), d2enum.UseBeltSlot1},
			{menu.asset.TranslateString("CfgBelt2"), d2enum.UseBeltSlot2},
			{menu.asset.TranslateString("CfgBelt3"), d2enum.UseBeltSlot3},
			{menu.asset.TranslateString("CfgBelt4"), d2enum.UseBeltSlot4},
			{menu.asset.TranslateString("Cfgswapweapons"), d2enum.SwapWeapons},
		},
		{
			{menu.asset.TranslateString("CfgChat"), d2enum.ToggleChatBox},
			{menu.asset.TranslateString("CfgRun"), d2enum.HoldRun},
			{menu.asset.TranslateString("CfgRunLock"), d2enum.ToggleRunWalk},
			{menu.asset.TranslateString("CfgStandStill"), d2enum.HoldStandStill},
			{menu.asset.TranslateString("CfgShowItems"), d2enum.HoldShowGroundItems},
			{menu.asset.TranslateString("CfgTogglePortraits"), d2enum.HoldShowPortraits},
		},
		{
			{menu.asset.TranslateString("CfgAutoMap"), d2enum.ToggleAutomap},
			{menu.asset.TranslateString("CfgAutoMapCenter"), d2enum.CenterAutomap},
			{menu.asset.TranslateString("CfgAutoMapParty"), d2enum.TogglePartyOnAutomap},
			{menu.asset.TranslateString("CfgAutoMapNames"), d2enum.ToggleNamesOnAutomap},
			{menu.asset.TranslateString("CfgToggleminimap"), d2enum.ToggleMiniMap},
		},
		{
			{menu.asset.TranslateString("CfgSay0"), d2enum.SayHelp},
			{menu.asset.TranslateString("CfgSay1"), d2enum.SayFollowMe},
			{menu.asset.TranslateString("CfgSay2"), d2enum.SayThisIsForYou},
			{menu.asset.TranslateString("CfgSay3"), d2enum.SayThanks},
			{menu.asset.TranslateString("CfgSay4"), d2enum.SaySorry},
			{menu.asset.TranslateString("CfgSay5"), d2enum.SayBye},
			{menu.asset.TranslateString("CfgSay6"), d2enum.SayNowYouDie},
			{menu.asset.TranslateString("CfgSay7"), d2enum.SayNowYouDie},
		},
		{
			{menu.asset.TranslateString("CfgSnapshot"), d2enum.TakeScreenShot},
			{menu.asset.TranslateString("CfgClearScreen"), d2enum.ClearScreen},
			{menu.asset.TranslateString("Cfgcleartextmsg"), d2enum.ClearMessages},
		},
	}

	wrapper := d2gui.CreateLayout(menu.renderer, d2gui.PositionTypeAbsolute, menu.asset)
	layout := wrapper.AddLayout(d2gui.PositionTypeVertical)
	for i, settingsGroup := range groups {
		groupLayout := layout.AddLayout(d2gui.PositionTypeVertical)

		for _, setting := range settingsGroup {
			bl := bindingLayout{}

			settingLayout := groupLayout.AddLayout(d2gui.PositionTypeHorizontal)
			settingLayout.AddSpacerStatic(17, 0)
			descLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
			descLabelWrapper.SetSize(190, 0)

			descLabel, _ := descLabelWrapper.AddLabel(setting.label, d2gui.FontStyleFormal11Units)
			descLabel.SetHoverColor(d2util.Color(0x5450D1FF))

			bl.wrapperLayout = settingLayout
			bl.descLabel = descLabel
			bl.descLayout = descLabelWrapper
			if binding := menu.keyMap.GetKeysForGameEvent(setting.gameEvent); binding != nil {
				primaryStr := KeyToString(binding.Primary, menu.asset)
				secondaryStr := KeyToString(binding.Secondary, menu.asset)
				primaryCol := menu.getKeyColor(binding.Primary)
				secondaryCol := menu.getKeyColor(binding.Secondary)

				if binding.IsEmpty() {
					primaryCol = d2util.Color(0xDB3F3DFF)
					secondaryCol = d2util.Color(0xDB3F3DFF)
				}

				primaryKeyLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
				primaryKeyLabelWrapper.SetSize(190, 0)
				primaryLabel, _ := primaryKeyLabelWrapper.AddLabelWithColor(primaryStr, d2gui.FontStyleFormal11Units, primaryCol)
				primaryLabel.SetHoverColor(d2util.Color(0x5450D1FF))

				bl.primaryLabel = primaryLabel
				bl.primaryLayout = primaryKeyLabelWrapper
				bl.gameEvent = setting.gameEvent

				secondaryKeyLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
				secondaryKeyLabelWrapper.SetSize(90, 0)
				secondaryLabel, _ := secondaryKeyLabelWrapper.AddLabelWithColor(secondaryStr, d2gui.FontStyleFormal11Units, secondaryCol)
				secondaryLabel.SetHoverColor(d2util.Color(0x5450D1FF))

				bl.secondaryLabel = secondaryLabel
				bl.secondaryLayout = secondaryKeyLabelWrapper
				bl.binding = binding
			}

			menu.bindingLayouts = append(menu.bindingLayouts, &bl)
		}

		if i < len(groups)-1 {
			layout.AddSpacerStatic(0, 25)
		}
	}

	return wrapper
}

func (menu *KeyBindingMenu) getKeyColor(key d2enum.Key) color.RGBA {
	switch key {
	case -1:
		return d2util.Color(0x555555FF)
	default:
		return d2util.Color(0xA1925DFF)
	}
}

func (menu *KeyBindingMenu) OnMouseButtonDown(event d2interface.MouseEvent) {
	if menu.Box.OnMouseButtonDown(event) {
		return
	}

	if menu.scrollbar != nil && menu.scrollbar.IsInSliderRect(event.X(), event.Y()) {
		menu.scrollbar.SetSliderClicked(true)
		menu.scrollbar.onSliderMouseClick(event)
		return
	}

	for _, bl := range menu.bindingLayouts {
		gameEvent, typ := bl.GetPointedLayoutAndLabel(event.X(), event.Y())
		if gameEvent != -1 {
			if menu.currentBindingLayout != nil {
				menu.lastBindingLayout = menu.currentBindingLayout
				menu.currentBindingLayout.Reset()
			}

			menu.currentBindingModifier = gameEvent
			menu.currentBindingLayout = bl
			if typ == KeyBindingTypePrimary {
				menu.currentBindingLayout.primaryLabel.SetIsBlinking(true)
			} else if typ == KeyBindingTypeSecondary {
				menu.currentBindingLayout.secondaryLabel.SetIsBlinking(true)
			}
			menu.currentBindingModifierType = typ
			menu.isAwaitingKeyDown = true

			bl.descLabel.SetIsHovered(true)
			bl.primaryLabel.SetIsHovered(true)
			bl.secondaryLabel.SetIsHovered(true)

			return
		} else if menu.currentBindingLayout != nil {
			menu.currentBindingLayout.Reset()
			menu.currentBindingLayout = nil
			menu.isAwaitingKeyDown = false
		}
	}
}

func (menu *KeyBindingMenu) OnMouseMove(event d2interface.MouseMoveEvent) {
	if menu.scrollbar != nil {
		menu.scrollbar.onMouseMove(event)
	}
}

func (menu *KeyBindingMenu) OnMouseButtonUp(event d2interface.MouseEvent) {
	if menu.scrollbar != nil {
		menu.scrollbar.SetSliderClicked(false)
	}
}

func (menu *KeyBindingMenu) saveKeyChange(key d2enum.Key) {
	var (
		existingBinding *KeyBinding
		gameEvent       d2enum.GameEvent
		bindingType     KeyBindingType = -1

		changeExisting *bindingChange
	)

	for ge, existingChange := range menu.changesToBeSaved {
		if existingChange.primary == key {
			bindingType = KeyBindingTypePrimary
		} else if existingChange.secondary == key {
			bindingType = KeyBindingTypeSecondary
		}

		if bindingType != -1 {
			existingBinding = existingChange.target
			changeExisting = existingChange
			gameEvent = ge

			break
		}
	}

	if existingBinding == nil {
		existingBinding, gameEvent, bindingType = menu.keyMap.GetBindingByKey(key)
	}

	if changeExisting == nil {
		changeExisting = menu.changesToBeSaved[gameEvent]
		if existingBinding != nil && changeExisting == nil {
			changeExisting = &bindingChange{
				target:    existingBinding,
				primary:   existingBinding.Primary,
				secondary: existingBinding.Secondary,
			}

			menu.changesToBeSaved[gameEvent] = changeExisting
		}
	}

	changeCurrent := menu.changesToBeSaved[menu.currentBindingLayout.gameEvent]
	if changeCurrent == nil {
		changeCurrent = &bindingChange{
			target:    menu.currentBindingLayout.binding,
			primary:   menu.currentBindingLayout.binding.Primary,
			secondary: menu.currentBindingLayout.binding.Secondary,
		}

		menu.changesToBeSaved[menu.currentBindingLayout.gameEvent] = changeCurrent
	}

	switch menu.currentBindingModifierType {
	case KeyBindingTypePrimary:
		changeCurrent.primary = key
		break
	case KeyBindingTypeSecondary:
		changeCurrent.secondary = key
		break
	}

	if changeExisting != nil {
		if bindingType == KeyBindingTypePrimary {
			changeExisting.primary = -1
		}

		if bindingType == KeyBindingTypeSecondary {
			changeExisting.secondary = -1
		}
	}

	menu.setBindingLabels(changeCurrent.primary, changeCurrent.secondary, menu.currentBindingLayout)

	if changeExisting != nil {
		for _, bindingLayout := range menu.bindingLayouts {
			if bindingLayout.binding == changeExisting.target {

				menu.setBindingLabels(changeExisting.primary, changeExisting.secondary, bindingLayout)
			}
		}
	}
}

func (menu *KeyBindingMenu) setBindingLabels(primary, secondary d2enum.Key, bl *bindingLayout) {
	noneStr := KeyToString(-1, menu.asset)

	if primary != -1 {
		bl.primaryLabel.SetText(KeyToString(primary, menu.asset))
		bl.primaryLabel.SetColor(d2util.Color(0xA1925DFF))
	} else {
		bl.primaryLabel.SetText(noneStr)
		bl.primaryLabel.SetColor(d2util.Color(0x555555FF))
	}

	if secondary != -1 {
		bl.secondaryLabel.SetText(KeyToString(secondary, menu.asset))
		bl.secondaryLabel.SetColor(d2util.Color(0xA1925DFF))
	} else {
		bl.secondaryLabel.SetText(noneStr)
		bl.secondaryLabel.SetColor(d2util.Color(0x555555FF))
	}

	if primary == -1 && secondary == -1 {
		bl.primaryLabel.SetColor(d2util.Color(0xDB3F3DFF))
		bl.secondaryLabel.SetColor(d2util.Color(0xDB3F3DFF))
	}
}

func (menu *KeyBindingMenu) onCancelClicked() {
	for gameEvent := range menu.changesToBeSaved {
		for _, bindingLayout := range menu.bindingLayouts {
			if bindingLayout.gameEvent == gameEvent {
				menu.setBindingLabels(bindingLayout.binding.Primary, bindingLayout.binding.Secondary, bindingLayout)
			}
		}
	}

	menu.changesToBeSaved = make(map[d2enum.GameEvent]*bindingChange)
	if menu.currentBindingLayout != nil {
		menu.currentBindingLayout.Reset()
		menu.lastBindingLayout = nil
		menu.currentBindingLayout = nil
		menu.currentBindingModifier = -1
		menu.currentBindingModifierType = -1
	}

	menu.Close()
	menu.escapeMenu.showLayout(optionsLayoutID)
}

func (menu *KeyBindingMenu) reload() {
	for _, bl := range menu.bindingLayouts {
		if bl.binding != nil {
			menu.setBindingLabels(bl.binding.Primary, bl.binding.Secondary, bl)
		}
	}
}

func (menu *KeyBindingMenu) onDefaultClicked() {
	menu.keyMap.ResetToDefault()
	menu.reload()

	menu.changesToBeSaved = make(map[d2enum.GameEvent]*bindingChange)
	if menu.currentBindingLayout != nil {
		menu.currentBindingLayout.Reset()
		menu.lastBindingLayout = nil
		menu.currentBindingLayout = nil
		menu.currentBindingModifier = -1
		menu.currentBindingModifierType = -1
	}
}

func (menu *KeyBindingMenu) onAcceptClicked() {
	for gameEvent, change := range menu.changesToBeSaved {
		menu.keyMap.SetPrimaryBinding(gameEvent, change.primary)
		menu.keyMap.SetSecondaryBinding(gameEvent, change.primary)
	}

	menu.changesToBeSaved = make(map[d2enum.GameEvent]*bindingChange)
	if menu.currentBindingLayout != nil {
		menu.currentBindingLayout.Reset()
		menu.lastBindingLayout = nil
		menu.currentBindingLayout = nil
		menu.currentBindingModifier = -1
		menu.currentBindingModifierType = -1
	}
}

func (menu *KeyBindingMenu) OnKeyDown(event d2interface.KeyEvent) {
	if menu.isAwaitingKeyDown {
		key := event.Key()

		if key == d2enum.KeyEscape {
			if menu.currentBindingLayout != nil {
				menu.lastBindingLayout = menu.currentBindingLayout
				menu.currentBindingLayout.Reset()
				menu.currentBindingLayout = nil
			}
		} else {
			menu.saveKeyChange(key)
		}

		menu.isAwaitingKeyDown = false
	}
}

func (menu *KeyBindingMenu) Render(target d2interface.Surface) {
	if menu.isOpen {
		menu.Box.Render(target)
		if menu.scrollbar != nil {
			menu.scrollbar.Render(target)
		}

		if menu.currentBindingLayout != nil {
			x, y := menu.currentBindingLayout.wrapperLayout.Sx, menu.currentBindingLayout.wrapperLayout.Sy
			w, h := menu.currentBindingLayout.wrapperLayout.GetSize()
			target.PushTranslation(x, y)
			target.DrawRect(w, h, d2util.Color(0x000000D0))
			target.Pop()
		}
	}
}
