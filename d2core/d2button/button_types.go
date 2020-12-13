package d2button

// ButtonType defines the type of button
type ButtonType int

// ButtonType constants
const (
	ButtonTypeWide     ButtonType = 1
	ButtonTypeMedium   ButtonType = 2
	ButtonTypeNarrow   ButtonType = 3
	ButtonTypeCancel   ButtonType = 4
	ButtonTypeTall     ButtonType = 5
	ButtonTypeShort    ButtonType = 6
	ButtonTypeOkCancel ButtonType = 7

	// Game UI

	ButtonTypeSkill              ButtonType = 7
	ButtonTypeRun                ButtonType = 8
	ButtonTypeMenu               ButtonType = 9
	ButtonTypeGoldCoin           ButtonType = 10
	ButtonTypeClose              ButtonType = 11
	ButtonTypeSecondaryInvHand   ButtonType = 12
	ButtonTypeMinipanelCharacter ButtonType = 13
	ButtonTypeMinipanelInventory ButtonType = 14
	ButtonTypeMinipanelSkill     ButtonType = 15
	ButtonTypeMinipanelAutomap   ButtonType = 16
	ButtonTypeMinipanelMessage   ButtonType = 17
	ButtonTypeMinipanelQuest     ButtonType = 18
	ButtonTypeMinipanelMen       ButtonType = 19
	ButtonTypeSquareClose        ButtonType = 20
	ButtonTypeSquareOk           ButtonType = 21
	ButtonTypeSkillTreeTab       ButtonType = 22
	ButtonTypeQuestDescr         ButtonType = 23
	ButtonTypeMinipanelOpenClose ButtonType = 24
	ButtonTypeMinipanelParty     ButtonType = 25
	ButtonTypeBuy                ButtonType = 26
	ButtonTypeSell               ButtonType = 27
	ButtonTypeRepair             ButtonType = 28
	ButtonTypeRepairAll          ButtonType = 29
	ButtonTypeLeftArrow          ButtonType = 30
	ButtonTypeRightArrow         ButtonType = 31
	ButtonTypeQuery              ButtonType = 32
	ButtonTypeSquelchChat        ButtonType = 33
	ButtonTypeTabBlank           ButtonType = 34
	ButtonTypeBlankQuestBtn      ButtonType = 35

	ButtonNoFixedWidth  int = -1
	ButtonNoFixedHeight int = -1
)
