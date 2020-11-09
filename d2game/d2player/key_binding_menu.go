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

	asset         *d2asset.AssetManager
	renderer      d2interface.Renderer
	ui            *d2ui.UIManager
	guiManager    *d2gui.GuiManager
	keyMap        *KeyMap
	mainLayout    *d2gui.Layout
	contentLayout *d2gui.Layout
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
		keyMap:        keyMap,
		asset:         asset,
		ui:            ui,
		guiManager:    guiManager,
		renderer:      renderer,
		mainLayout:    mainLayout,
		contentLayout: contentLayout,
	}

	ret.Box = NewBox(asset, renderer, ui, guiManager, ret.mainLayout, 620, 375, 90, 65, "")

	return ret
}

func (menu *KeyBindingMenu) Load() error {
	menu.Box.Load()

	// _, mainLayoutY := menu.mainLayout.GetPosition()
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
	newLayoutScrollbar(bindingWrapper, bindingLayout)

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
				label:   menu.asset.TranslateString("CfgBelt0"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseBeltSlot1),
			},
			{
				label:   menu.asset.TranslateString("CfgBelt1"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseBeltSlot2),
			},
			{
				label:   menu.asset.TranslateString("CfgBelt2"),
				binding: menu.keyMap.GetKeysForGameEvent(d2enum.UseBeltSlot3),
			},
			{
				label:   menu.asset.TranslateString("CfgBelt3"),
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
			settingLayout := groupLayout.AddLayout(d2gui.PositionTypeHorizontal)
			settingLayout.AddSpacerStatic(26, 0)
			descLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
			descLabelWrapper.SetSize(190, 0)
			descLabelWrapper.AddLabel(setting.label, d2gui.FontStyleFormal11Units)

			if binding := setting.binding; binding != nil {
				primaryStr := d2util.KeyToString(binding.Primary, menu.asset)
				secondaryStr := d2util.KeyToString(binding.Secondary, menu.asset)
				primaryKeyLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
				primaryKeyLabelWrapper.SetSize(190, 0)
				primaryKeyLabelWrapper.AddLabelWithColor(primaryStr, d2gui.FontStyleFormal11Units, menu.getKeyColor(binding.Primary))

				secondaryKeyLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
				secondaryKeyLabelWrapper.SetSize(90, 0)
				secondaryKeyLabelWrapper.AddLabelWithColor(secondaryStr, d2gui.FontStyleFormal11Units, menu.getKeyColor(binding.Secondary))
			}
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
		return d2util.Color(0x555555)
	default:
		return d2util.Color(0xA1925DFF)
	}
}
