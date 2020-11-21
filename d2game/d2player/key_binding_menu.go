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

const (
	selectionBackgroundColor = 0x000000d0
	defaultGameEvent         = -1

	keyBindingMenuWidth  = 620
	keyBindingMenuHeight = 375
	keyBindingMenuX      = 90
	keyBindingMenuY      = 75

	keyBindingMenuPaddingX    = 17
	keyBindingSettingPaddingY = 19

	keyBindingMenuHeaderHeight  = 24
	keyBindingMenuHeaderSpacer1 = 131
	keyBindingMenuHeaderSpacer2 = 86

	keyBindingMenuBindingSpacerBetween   = 25
	keyBindingMenuBindingSpacerLeft      = 17
	keyBindingMenuBindingDescWidth       = 190
	keyBindingMenuBindingDescHeight      = 0
	keyBindingMenuBindingPrimaryWidth    = 190
	keyBindingMenuBindingPrimaryHeight   = 0
	keyBindingMenuBindingSecondaryWidth  = 90
	keyBindingMenuBindingSecondaryHeight = 0
)

type bindingChange struct {
	target    *KeyBinding
	primary   d2enum.Key
	secondary d2enum.Key
}

// NewKeyBindingMenu generates a new instance of the "Configure Keys"
// menu found in the options
func NewKeyBindingMenu(
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	ui *d2ui.UIManager,
	guiManager *d2gui.GuiManager,
	keyMap *KeyMap,
	l d2util.LogLevel,
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

	ret.Logger = d2util.NewLogger()
	ret.Logger.SetLevel(l)
	ret.Logger.SetPrefix(logPrefix)

	ret.Box = d2gui.NewBox(
		asset, renderer, ui, ret.mainLayout,
		keyBindingMenuWidth, keyBindingMenuHeight,
		keyBindingMenuX, keyBindingMenuY, l, "",
	)

	ret.Box.SetPadding(keyBindingMenuPaddingX, keyBindingSettingPaddingY)

	ret.Box.SetOptions([]*d2gui.LabelButton{
		d2gui.NewLabelButton(0, 0, "Cancel", d2util.Color(d2gui.ColorRed), d2util.LogLevelDefault, func() {
			if err := ret.onCancelClicked(); err != nil {
				ret.Errorf("error while clicking option Cancel: %v", err.Error())
			}
		}),
		d2gui.NewLabelButton(0, 0, "Default", d2util.Color(d2gui.ColorBlue), d2util.LogLevelDefault, func() {
			if err := ret.onDefaultClicked(); err != nil {
				ret.Errorf("error while clicking option Default: %v", err)
			}
		}),
		d2gui.NewLabelButton(0, 0, "Accept", d2util.Color(d2gui.ColorGreen), d2util.LogLevelDefault, func() {
			if err := ret.onAcceptClicked(); err != nil {
				ret.Errorf("error while clicking option Accept: %v", err)
			}
		}),
	})

	return ret
}

// KeyBindingMenu represents the menu to view/edit the
// key bindings
type KeyBindingMenu struct {
	*d2gui.Box

	asset      *d2asset.AssetManager
	renderer   d2interface.Renderer
	ui         *d2ui.UIManager
	guiManager *d2gui.GuiManager
	keyMap     *KeyMap
	escapeMenu *EscapeMenu

	mainLayout       *d2gui.Layout
	contentLayout    *d2gui.Layout
	scrollbar        *d2gui.LayoutScrollbar
	bindingLayouts   []*bindingLayout
	changesToBeSaved map[d2enum.GameEvent]*bindingChange

	isAwaitingKeyDown          bool
	currentBindingModifierType KeyBindingType
	currentBindingModifier     d2enum.GameEvent
	currentBindingLayout       *bindingLayout
	lastBindingLayout          *bindingLayout

	*d2util.Logger
}

// Close will disable the render of the menu and clear
// the current selection
func (menu *KeyBindingMenu) Close() error {
	menu.Box.Close()

	if err := menu.clearSelection(); err != nil {
		return err
	}

	return nil
}

// Load will setup the layouts of the menu
func (menu *KeyBindingMenu) Load() error {
	if err := menu.Box.Load(); err != nil {
		return err
	}

	mainLayoutW, mainLayoutH := menu.mainLayout.GetSize()

	headerLayout := menu.contentLayout.AddLayout(d2gui.PositionTypeHorizontal)
	headerLayout.SetSize(mainLayoutW, keyBindingMenuHeaderHeight)

	if _, err := headerLayout.AddLabelWithColor(
		menu.asset.TranslateString("CfgFunction"),
		d2gui.FontStyleFormal11Units,
		d2util.Color(d2gui.ColorBrown),
	); err != nil {
		return err
	}

	headerLayout.AddSpacerStatic(keyBindingMenuHeaderSpacer1, keyBindingMenuHeaderHeight)

	if _, err := headerLayout.AddLabelWithColor(
		menu.asset.TranslateString("CfgPrimaryKey"),
		d2gui.FontStyleFormal11Units,
		d2util.Color(d2gui.ColorBrown),
	); err != nil {
		return err
	}

	headerLayout.AddSpacerStatic(keyBindingMenuHeaderSpacer2, 1)

	if _, err := headerLayout.AddLabelWithColor(
		menu.asset.TranslateString("CfgSecondaryKey"),
		d2gui.FontStyleFormal11Units,
		d2util.Color(d2gui.ColorBrown),
	); err != nil {
		return err
	}

	headerLayout.SetVerticalAlign(d2gui.VerticalAlignMiddle)

	bindingWrapper := menu.contentLayout.AddLayout(d2gui.PositionTypeAbsolute)
	bindingWrapper.SetPosition(0, keyBindingMenuHeaderHeight)
	bindingWrapper.SetSize(mainLayoutW, mainLayoutH-keyBindingMenuHeaderHeight)

	bindingLayout := menu.generateLayout()

	menu.Box.GetLayout().AdjustEntryPlacement()
	menu.mainLayout.AdjustEntryPlacement()
	menu.contentLayout.AdjustEntryPlacement()

	menu.scrollbar = d2gui.NewLayoutScrollbar(bindingWrapper, bindingLayout)

	if err := menu.scrollbar.Load(menu.ui); err != nil {
		return err
	}

	bindingWrapper.AddLayoutFromSource(bindingLayout)
	bindingWrapper.AdjustEntryPlacement()

	return nil
}

type keyBindingSetting struct {
	label     string
	gameEvent d2enum.GameEvent
}

func (menu *KeyBindingMenu) getBindingGroups() [][]keyBindingSetting {
	return [][]keyBindingSetting{
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
}

func (menu *KeyBindingMenu) generateLayout() *d2gui.Layout {
	groups := menu.getBindingGroups()

	wrapper := d2gui.CreateLayout(menu.renderer, d2gui.PositionTypeAbsolute, menu.asset)
	layout := wrapper.AddLayout(d2gui.PositionTypeVertical)

	for i, settingsGroup := range groups {
		groupLayout := layout.AddLayout(d2gui.PositionTypeVertical)

		for _, setting := range settingsGroup {
			bl := bindingLayout{}

			settingLayout := groupLayout.AddLayout(d2gui.PositionTypeHorizontal)
			settingLayout.AddSpacerStatic(keyBindingMenuBindingSpacerLeft, 0)
			descLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
			descLabelWrapper.SetSize(keyBindingMenuBindingDescWidth, keyBindingMenuBindingDescHeight)

			descLabel, _ := descLabelWrapper.AddLabel(setting.label, d2gui.FontStyleFormal11Units)
			descLabel.SetHoverColor(d2util.Color(d2gui.ColorBlue))

			bl.wrapperLayout = settingLayout
			bl.descLabel = descLabel
			bl.descLayout = descLabelWrapper

			if binding := menu.keyMap.GetKeysForGameEvent(setting.gameEvent); binding != nil {
				primaryStr := menu.keyMap.KeyToString(binding.Primary)
				secondaryStr := menu.keyMap.KeyToString(binding.Secondary)
				primaryCol := menu.getKeyColor(binding.Primary)
				secondaryCol := menu.getKeyColor(binding.Secondary)

				if binding.IsEmpty() {
					primaryCol = d2util.Color(d2gui.ColorRed)
					secondaryCol = d2util.Color(d2gui.ColorRed)
				}

				primaryKeyLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
				primaryKeyLabelWrapper.SetSize(keyBindingMenuBindingPrimaryWidth, keyBindingMenuBindingPrimaryHeight)
				primaryLabel, _ := primaryKeyLabelWrapper.AddLabelWithColor(primaryStr, d2gui.FontStyleFormal11Units, primaryCol)
				primaryLabel.SetHoverColor(d2util.Color(d2gui.ColorBlue))

				bl.primaryLabel = primaryLabel
				bl.primaryLayout = primaryKeyLabelWrapper
				bl.gameEvent = setting.gameEvent

				secondaryKeyLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
				secondaryKeyLabelWrapper.SetSize(keyBindingMenuBindingSecondaryWidth, keyBindingMenuBindingSecondaryHeight)
				secondaryLabel, _ := secondaryKeyLabelWrapper.AddLabelWithColor(secondaryStr, d2gui.FontStyleFormal11Units, secondaryCol)
				secondaryLabel.SetHoverColor(d2util.Color(d2gui.ColorBlue))

				bl.secondaryLabel = secondaryLabel
				bl.secondaryLayout = secondaryKeyLabelWrapper
				bl.binding = binding
			}

			menu.bindingLayouts = append(menu.bindingLayouts, &bl)
		}

		if i < len(groups)-1 {
			layout.AddSpacerStatic(0, keyBindingMenuBindingSpacerBetween)
		}
	}

	return wrapper
}

func (menu *KeyBindingMenu) getKeyColor(key d2enum.Key) color.RGBA {
	switch key {
	case -1:
		return d2util.Color(d2gui.ColorGrey)
	default:
		return d2util.Color(d2gui.ColorBrown)
	}
}

func (menu *KeyBindingMenu) setSelection(bl *bindingLayout, bindingType KeyBindingType, gameEvent d2enum.GameEvent) error {
	if menu.currentBindingLayout != nil {
		menu.lastBindingLayout = menu.currentBindingLayout
		if err := menu.currentBindingLayout.Reset(); err != nil {
			return err
		}
	}

	menu.currentBindingModifier = gameEvent
	menu.currentBindingLayout = bl

	if bindingType == KeyBindingTypePrimary {
		menu.currentBindingLayout.primaryLabel.SetIsBlinking(true)
	} else if bindingType == KeyBindingTypeSecondary {
		menu.currentBindingLayout.secondaryLabel.SetIsBlinking(true)
	}

	menu.currentBindingModifierType = bindingType
	menu.isAwaitingKeyDown = true

	if err := bl.descLabel.SetIsHovered(true); err != nil {
		return err
	}

	if err := bl.primaryLabel.SetIsHovered(true); err != nil {
		return err
	}

	if err := bl.secondaryLabel.SetIsHovered(true); err != nil {
		return err
	}

	return nil
}

func (menu *KeyBindingMenu) onMouseButtonDown(event d2interface.MouseEvent) error {
	if !menu.IsOpen() {
		return nil
	}

	menu.Box.OnMouseButtonDown(event)

	if menu.scrollbar != nil {
		if menu.scrollbar.IsInSliderRect(event.X(), event.Y()) {
			menu.scrollbar.SetSliderClicked(true)
			menu.scrollbar.OnSliderMouseClick(event)

			return nil
		}

		if menu.scrollbar.IsInArrowUpRect(event.X(), event.Y()) {
			if !menu.scrollbar.IsArrowUpClicked() {
				menu.scrollbar.SetArrowUpClicked(true)
			}

			menu.scrollbar.OnArrowUpClick()

			return nil
		}

		if menu.scrollbar.IsInArrowDownRect(event.X(), event.Y()) {
			if !menu.scrollbar.IsArrowDownClicked() {
				menu.scrollbar.SetArrowDownClicked(true)
			}

			menu.scrollbar.OnArrowDownClick()

			return nil
		}
	}

	for _, bl := range menu.bindingLayouts {
		gameEvent, typ := bl.GetPointedLayoutAndLabel(event.X(), event.Y())

		if gameEvent != -1 {
			if err := menu.setSelection(bl, typ, gameEvent); err != nil {
				return err
			}

			break
		} else if menu.currentBindingLayout != nil {
			if err := menu.clearSelection(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (menu *KeyBindingMenu) onMouseMove(event d2interface.MouseMoveEvent) {
	if !menu.IsOpen() {
		return
	}

	if menu.scrollbar != nil && menu.scrollbar.IsSliderClicked() {
		menu.scrollbar.OnMouseMove(event)
	}
}

func (menu *KeyBindingMenu) onMouseButtonUp() {
	if !menu.IsOpen() {
		return
	}

	if menu.scrollbar != nil {
		menu.scrollbar.SetSliderClicked(false)
		menu.scrollbar.SetArrowDownClicked(false)
		menu.scrollbar.SetArrowUpClicked(false)
	}
}

func (menu *KeyBindingMenu) getPendingChangeByKey(key d2enum.Key) (*bindingChange, *KeyBinding, d2enum.GameEvent, KeyBindingType) {
	var (
		existingBinding *KeyBinding
		gameEvent       d2enum.GameEvent
		bindingType     KeyBindingType
	)

	for ge, existingChange := range menu.changesToBeSaved {
		if existingChange.primary == key {
			bindingType = KeyBindingTypePrimary
		} else if existingChange.secondary == key {
			bindingType = KeyBindingTypeSecondary
		}

		if bindingType != -1 {
			existingBinding = existingChange.target
			gameEvent = ge

			return existingChange, existingBinding, gameEvent, bindingType
		}
	}

	return nil, nil, -1, KeyBindingTypeNone
}

func (menu *KeyBindingMenu) saveKeyChange(key d2enum.Key) error {
	changeExisting, existingBinding, gameEvent, bindingType := menu.getPendingChangeByKey(key)

	if changeExisting == nil {
		existingBinding, gameEvent, bindingType = menu.keyMap.GetBindingByKey(key)
	}

	if existingBinding != nil && changeExisting == nil {
		changeExisting = &bindingChange{
			target:    existingBinding,
			primary:   existingBinding.Primary,
			secondary: existingBinding.Secondary,
		}

		menu.changesToBeSaved[gameEvent] = changeExisting
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
	case KeyBindingTypeSecondary:
		changeCurrent.secondary = key
	}

	if changeExisting != nil {
		if bindingType == KeyBindingTypePrimary {
			changeExisting.primary = -1
		}

		if bindingType == KeyBindingTypeSecondary {
			changeExisting.secondary = -1
		}
	}

	if err := menu.setBindingLabels(
		changeCurrent.primary,
		changeCurrent.secondary,
		menu.currentBindingLayout,
	); err != nil {
		return err
	}

	if changeExisting != nil {
		for _, bindingLayout := range menu.bindingLayouts {
			if bindingLayout.binding == changeExisting.target {
				if err := menu.setBindingLabels(changeExisting.primary, changeExisting.secondary, bindingLayout); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (menu *KeyBindingMenu) setBindingLabels(primary, secondary d2enum.Key, bl *bindingLayout) error {
	noneStr := menu.keyMap.KeyToString(-1)

	if primary != -1 {
		if err := bl.SetPrimaryBindingTextAndColor(menu.keyMap.KeyToString(primary), d2util.Color(d2gui.ColorBrown)); err != nil {
			return err
		}
	} else {
		if err := bl.SetPrimaryBindingTextAndColor(noneStr, d2util.Color(d2gui.ColorGrey)); err != nil {
			return err
		}
	}

	if secondary != -1 {
		if err := bl.SetSecondaryBindingTextAndColor(menu.keyMap.KeyToString(secondary), d2util.Color(d2gui.ColorBrown)); err != nil {
			return err
		}
	} else {
		if err := bl.SetSecondaryBindingTextAndColor(noneStr, d2util.Color(d2gui.ColorGrey)); err != nil {
			return err
		}
	}

	if primary == -1 && secondary == -1 {
		if err := bl.primaryLabel.SetColor(d2util.Color(d2gui.ColorRed)); err != nil {
			return err
		}

		if err := bl.secondaryLabel.SetColor(d2util.Color(d2gui.ColorRed)); err != nil {
			return err
		}
	}

	return nil
}

func (menu *KeyBindingMenu) onCancelClicked() error {
	for gameEvent := range menu.changesToBeSaved {
		for _, bindingLayout := range menu.bindingLayouts {
			if bindingLayout.gameEvent == gameEvent {
				if err := menu.setBindingLabels(bindingLayout.binding.Primary, bindingLayout.binding.Secondary, bindingLayout); err != nil {
					return err
				}
			}
		}
	}

	menu.changesToBeSaved = make(map[d2enum.GameEvent]*bindingChange)
	if menu.currentBindingLayout != nil {
		if err := menu.clearSelection(); err != nil {
			return err
		}
	}

	if err := menu.Close(); err != nil {
		return err
	}

	menu.escapeMenu.showLayout(optionsLayoutID)

	return nil
}

func (menu *KeyBindingMenu) reload() error {
	for _, bl := range menu.bindingLayouts {
		if bl.binding != nil {
			if err := menu.setBindingLabels(bl.binding.Primary, bl.binding.Secondary, bl); err != nil {
				return err
			}
		}
	}

	return nil
}

func (menu *KeyBindingMenu) clearSelection() error {
	if menu.currentBindingLayout != nil {
		if err := menu.currentBindingLayout.Reset(); err != nil {
			return err
		}

		menu.lastBindingLayout = menu.currentBindingLayout
		menu.currentBindingLayout = nil
		menu.currentBindingModifier = -1
		menu.currentBindingModifierType = -1
	}

	return nil
}

func (menu *KeyBindingMenu) onDefaultClicked() error {
	menu.keyMap.ResetToDefault()

	if err := menu.reload(); err != nil {
		return err
	}

	menu.changesToBeSaved = make(map[d2enum.GameEvent]*bindingChange)

	return menu.clearSelection()
}

func (menu *KeyBindingMenu) onAcceptClicked() error {
	for gameEvent, change := range menu.changesToBeSaved {
		menu.keyMap.SetPrimaryBinding(gameEvent, change.primary)
		menu.keyMap.SetSecondaryBinding(gameEvent, change.primary)
	}

	menu.changesToBeSaved = make(map[d2enum.GameEvent]*bindingChange)

	return menu.clearSelection()
}

// OnKeyDown will assign the new key to the selected binding if any
func (menu *KeyBindingMenu) OnKeyDown(event d2interface.KeyEvent) error {
	if menu.isAwaitingKeyDown {
		key := event.Key()

		if key == d2enum.KeyEscape {
			if menu.currentBindingLayout != nil {
				menu.lastBindingLayout = menu.currentBindingLayout

				if err := menu.currentBindingLayout.Reset(); err != nil {
					return err
				}

				if err := menu.clearSelection(); err != nil {
					return err
				}
			}
		} else {
			if err := menu.saveKeyChange(key); err != nil {
				return err
			}
		}

		menu.isAwaitingKeyDown = false
	}

	return nil
}

// Advance computes the state of the elements of the menu overtime
func (menu *KeyBindingMenu) Advance(elapsed float64) error {
	if menu.scrollbar != nil {
		if err := menu.scrollbar.Advance(elapsed); err != nil {
			return err
		}
	}

	return nil
}

// Render draws the different element of the menu on the target surface
func (menu *KeyBindingMenu) Render(target d2interface.Surface) error {
	if menu.IsOpen() {
		if err := menu.Box.Render(target); err != nil {
			return err
		}

		if menu.scrollbar != nil {
			menu.scrollbar.Render(target)
		}

		if menu.currentBindingLayout != nil {
			x, y := menu.currentBindingLayout.wrapperLayout.Sx, menu.currentBindingLayout.wrapperLayout.Sy
			w, h := menu.currentBindingLayout.wrapperLayout.GetSize()

			target.PushTranslation(x, y)
			target.DrawRect(w, h, d2util.Color(selectionBackgroundColor))
			target.Pop()
		}
	}

	return nil
}
