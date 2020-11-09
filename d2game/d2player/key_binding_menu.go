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

// KeyBindingMenu represents the menu to view/edit the
// key bindings
type KeyBindingMenu struct {
	*Box

	asset          *d2asset.AssetManager
	renderer       d2interface.Renderer
	ui             *d2ui.UIManager
	guiManager     *d2gui.GuiManager
	keyMap         *KeyMap
	mainLayout     *d2gui.Layout
	contentLayout  *d2gui.Layout
	scrollbar      *LayoutScrollbar
	bindingLayouts []*bindingLayout

	isAwaitingKeyDown          bool
	currentBindingModifierType awaitingKeyDownType
	currentBindingModifier     d2enum.GameEvent
	currentBindingLayout       *bindingLayout
	lastBindingLayout          *bindingLayout
}

type awaitingKeyDownType int

const (
	awaitingKeyDownTypePrimary awaitingKeyDownType = iota
	awaitingKeyDownTypeSecondary
)

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
}

func (l *bindingLayout) GetPointedLayoutAndLabel(x, y int) (d2enum.GameEvent, awaitingKeyDownType) {
	ww, hh := l.descLayout.GetSize()
	xx, yy := l.descLayout.Sx, l.descLayout.Sy
	if x >= xx && x <= xx+ww && y >= yy && y <= yy+hh {
		return l.gameEvent, awaitingKeyDownTypePrimary
	}

	if l.primaryLayout != nil {
		ww, hh = l.primaryLayout.GetSize()
		xx, yy = l.primaryLayout.Sx, l.primaryLayout.Sy
		if x >= xx && x <= xx+ww && y >= yy && y <= yy+hh {
			return l.gameEvent, awaitingKeyDownTypePrimary
		}
	}

	if l.secondaryLayout != nil {
		ww, hh = l.secondaryLayout.GetSize()
		xx, yy = l.secondaryLayout.Sx, l.secondaryLayout.Sy
		if x >= xx && x <= xx+ww && y >= yy && y <= yy+hh {
			return l.gameEvent, awaitingKeyDownTypeSecondary
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
) *KeyBindingMenu {
	mainLayout := d2gui.CreateLayout(renderer, d2gui.PositionTypeAbsolute, asset)
	contentLayout := mainLayout.AddLayout(d2gui.PositionTypeAbsolute)

	ret := &KeyBindingMenu{
		keyMap:         keyMap,
		asset:          asset,
		ui:             ui,
		guiManager:     guiManager,
		renderer:       renderer,
		mainLayout:     mainLayout,
		contentLayout:  contentLayout,
		bindingLayouts: []*bindingLayout{},
	}

	ret.Box = NewBox(asset, renderer, ui, guiManager, ret.mainLayout, 620, 375, 90, 65, "")

	return ret
}

func (menu *KeyBindingMenu) Load() error {
	menu.Box.Load()

	mainLayoutW, mainLayoutH := menu.mainLayout.GetSize()

	headerLayout := menu.mainLayout.AddLayout(d2gui.PositionTypeHorizontal)
	headerLayout.SetSize(mainLayoutW, 24)
	headerLayout.AddSpacerStatic(13, 1)
	headerLayout.AddLabelWithColor("Function", d2gui.FontStyleFormal11Units, d2util.Color(0xA1925DFF))
	headerLayout.AddSpacerStatic(131, 1)
	headerLayout.AddLabelWithColor("Key/Button One", d2gui.FontStyleFormal11Units, d2util.Color(0xA1925DFF))
	headerLayout.AddSpacerStatic(86, 1)
	headerLayout.AddLabelWithColor("Key/Button Two", d2gui.FontStyleFormal11Units, d2util.Color(0xA1925DFF))
	headerLayout.SetVerticalAlign(d2gui.VerticalAlignMiddle)

	bindingWrapper := menu.mainLayout.AddLayout(d2gui.PositionTypeAbsolute)
	bindingWrapper.SetPosition(0, 24)
	bindingWrapper.SetSize(mainLayoutW, mainLayoutH-24)
	bindingLayout := menu.generateLayout()
	menu.scrollbar = newLayoutScrollbar(bindingWrapper, bindingLayout)

	bindingWrapper.AddLayoutFromSource(bindingLayout)

	return nil
}

type keyBindingSetting struct {
	label   string
	binding *KeyBinding
}

func (menu *KeyBindingMenu) generateLayout() *d2gui.Layout {
	groups := [][]keyBindingSetting{
		{
			{
				label:   menu.asset.TranslateString("CfgCharacter"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleCharacterPanel),
			},
			{
				label:   menu.asset.TranslateString("CfgInventory"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleInventoryPanel),
			},
			{
				label:   menu.asset.TranslateString("CfgParty"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.TogglePartyPanel),
			},
			{
				label:   menu.asset.TranslateString("Cfghireling"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleHirelingPanel),
			},
			{
				label:   menu.asset.TranslateString("CfgMessageLog"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleMessageLog),
			},
			{
				label:   menu.asset.TranslateString("CfgQuestLog"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleQuestLog),
			},
			{
				label:   menu.asset.TranslateString("CfgHelp"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleHelpScreen),
			},
		},
		{
			{
				label:   menu.asset.TranslateString("CfgSkillTree"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleSkillTreePanel),
			},
			{
				label:   menu.asset.TranslateString("CfgSkillPick"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleLeftSkillSelector),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill1"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill1),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill2"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill2),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill3"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill3),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill4"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill4),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill5"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill5),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill6"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill6),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill7"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill7),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill8"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill8),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill9"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill9),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill10"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill10),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill11"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill11),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill12"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill12),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill13"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill13),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill14"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill14),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill15"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill15),
			},
			{
				label:   menu.asset.TranslateString("CfgSkill16"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseSkill16),
			},
			{
				label:   menu.asset.TranslateString("Cfgskillup"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.SelectPreviousSkill),
			},
			{
				label:   menu.asset.TranslateString("Cfgskilldown"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.SelectNextSkill),
			},
		},
		{
			{
				label:   menu.asset.TranslateString("CfgBeltShow"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleBelts),
			},
			{
				label:   menu.asset.TranslateString("CfgBelt1"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseBeltSlot1),
			},
			{
				label:   menu.asset.TranslateString("CfgBelt2"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseBeltSlot2),
			},
			{
				label:   menu.asset.TranslateString("CfgBelt3"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseBeltSlot3),
			},
			{
				label:   menu.asset.TranslateString("CfgBelt4"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseBeltSlot4),
			},
			{
				label:   menu.asset.TranslateString("Cfgswapweapons"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.SwapWeapons),
			},
		},
		{
			{
				label:   menu.asset.TranslateString("Cfgchat"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleChatBox),
			},
			{
				label:   menu.asset.TranslateString("CfgRun"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.HoldRun),
			},
			{
				label:   menu.asset.TranslateString("CfgRunLock"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleRunWalk),
			},
			{
				label:   menu.asset.TranslateString("CfgStandStill"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.HoldStandStill),
			},
			{
				label:   menu.asset.TranslateString("CfgShowItems"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.HoldShowGroundItems),
			},
			{
				label:   menu.asset.TranslateString("CfgTogglePortraits"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.HoldShowPortraits),
			},
		},
		{
			{
				label:   menu.asset.TranslateString("CfgAutoMap"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleAutomap),
			},
			{
				label:   menu.asset.TranslateString("CfgAutoMapCenter"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.CenterAutomap),
			},
			{
				label:   menu.asset.TranslateString("CfgAutoMapParty"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.TogglePartyOnAutomap),
			},
			{
				label:   menu.asset.TranslateString("CfgAutoMapNames"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleNamesOnAutomap),
			},
			{
				label:   menu.asset.TranslateString("CfgToggleminimap"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ToggleMiniMap),
			},
		},
		{
			{
				label:   menu.asset.TranslateString("CfgSay0"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.SayHelp),
			},
			{
				label:   menu.asset.TranslateString("CfgSay1"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.SayFollowMe),
			},
			{
				label:   menu.asset.TranslateString("CfgSay2"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.SayThisIsForYou),
			},
			{
				label:   menu.asset.TranslateString("CfgSay3"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.SayThanks),
			},
			{
				label:   menu.asset.TranslateString("CfgSay4"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.SaySorry),
			},
			{
				label:   menu.asset.TranslateString("CfgSay5"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.SayBye),
			},
			{
				label:   menu.asset.TranslateString("CfgSay6"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.SayNowYouDie),
			},
			{
				label:   menu.asset.TranslateString("CfgSay7"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.SayNowYouDie),
			},
		},
		{
			{
				label:   menu.asset.TranslateString("CfgSnapshot"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.TakeScreenShot),
			},
			{
				label:   menu.asset.TranslateString("CfgClearScreen"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ClearScreen),
			},
			{
				label:   menu.asset.TranslateString("Cfgcleartextmsg"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.ClearMessages),
			},
		},
	}

	wrapper := d2gui.CreateLayout(menu.renderer, d2gui.PositionTypeAbsolute, menu.asset)
	layout := wrapper.AddLayout(d2gui.PositionTypeVertical)
	for i, settingsGroup := range groups {
		groupLayout := layout.AddLayout(d2gui.PositionTypeVertical)

		for _, setting := range settingsGroup {
			bl := bindingLayout{}

			settingLayout := groupLayout.AddLayout(d2gui.PositionTypeHorizontal)
			settingLayout.AddSpacerStatic(26, 0)
			descLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
			descLabelWrapper.SetSize(190, 0)

			descLabel, _ := descLabelWrapper.AddLabel(setting.label, d2gui.FontStyleFormal11Units)
			descLabel.SetHoverColor(d2util.Color(0x5450D1FF))

			bl.wrapperLayout = settingLayout
			bl.descLabel = descLabel
			bl.descLayout = descLabelWrapper
			if binding := setting.binding; binding != nil {
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

func (menu *KeyBindingMenu) OnKeyDown(event d2interface.KeyEvent) {
	if menu.isAwaitingKeyDown {
		key := event.Key()
		if key != d2enum.KeyEscape {
			var (
				overwrittenBinding *KeyBinding
				bindingType        KeyBindingType = -1
			)

			switch menu.currentBindingModifierType {
			case awaitingKeyDownTypePrimary:
				overwrittenBinding, bindingType = menu.keyMap.SetPrimaryBinding(menu.currentBindingModifier, key)
				menu.currentBindingLayout.primaryLabel.SetText(KeyToString(key, menu.asset))
				menu.currentBindingLayout.primaryLabel.SetColor(d2util.Color(0xA1925DFF))
				break
			case awaitingKeyDownTypeSecondary:
				overwrittenBinding, bindingType = menu.keyMap.SetSecondaryBinding(menu.currentBindingModifier, key)
				menu.currentBindingLayout.secondaryLabel.SetText(KeyToString(key, menu.asset))
				menu.currentBindingLayout.secondaryLabel.SetColor(d2util.Color(0xA1925DFF))
				break
			}

			if overwrittenBinding != nil {
				for _, bindingLayout := range menu.bindingLayouts {
					if bindingLayout.binding == overwrittenBinding {
						noneStr := KeyToString(-1, menu.asset)

						switch bindingType {
						case KeyBindingTypePrimary:
							bindingLayout.primaryLabel.SetText(noneStr)
							bindingLayout.primaryLabel.SetColor(d2util.Color(0x555555FF))
							break
						case KeyBindingTypeSecondary:
							bindingLayout.secondaryLabel.SetText(noneStr)
							bindingLayout.secondaryLabel.SetColor(d2util.Color(0x555555FF))
							break
						}

						if bindingLayout.binding.IsEmpty() {
							bindingLayout.primaryLabel.SetColor(d2util.Color(0xDB3F3DFF))
							bindingLayout.secondaryLabel.SetColor(d2util.Color(0xDB3F3DFF))
						}
					}
				}
			}
		}

		menu.isAwaitingKeyDown = false
		menu.lastBindingLayout = menu.currentBindingLayout
		menu.currentBindingLayout.Reset()
		menu.currentBindingLayout = nil
	}
}

func (menu *KeyBindingMenu) Render(target d2interface.Surface) {
	menu.Box.Render(target)
	if menu.currentBindingLayout != nil {
		x, y := menu.currentBindingLayout.wrapperLayout.Sx, menu.currentBindingLayout.wrapperLayout.Sy
		w, h := menu.currentBindingLayout.wrapperLayout.GetSize()
		target.PushTranslation(x, y)
		target.DrawRect(w, h, d2util.Color(0x000000D0))
		target.Pop()
	}
}
