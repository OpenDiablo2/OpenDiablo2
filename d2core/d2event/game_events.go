package d2event

type gameEvent int

const (
	GamePause				gameEvent = iota
	GameResume
	GameSaveAndExit

	ViewToggleTerminal
	ViewShowTerminal
	ViewHideTerminal

	ViewToggleMainMenu
	ViewShowMainMenu
	ViewHideMainMenu

	ViewToggleCharacter
	ViewShowCharacter
	ViewHideCharacter

	ViewToggleInventory
	ViewShowInventory
	ViewHideInventory

	ViewToggleParty
	ViewShowParty
	ViewHideParty

	ViewToggleHireling
	ViewShowHireling
	ViewHideHireling

	ViewToggleMessage
	ViewShowMessage
	ViewHideMessage

	ViewToggleQuests
	ViewShowQuests
	ViewHideQuests

	ViewToggleSkillTree
	ViewShowSkillTree
	ViewHideSkillTree

	ViewToggleSkillBarLeft
	ViewShowSkillBarLeft
	ViewHideSkillBarLeft

	ViewToggleSkillBarRight
	ViewShowSkillBarRight
	ViewHideSkillBarRight

	ViewToggleMap
	ViewShowMap
	ViewHideMap

	ViewToggleChat
	ViewShowChat
	ViewHideChat

	ViewToggleItems
	ViewShowItems
	ViewHideItems

	ViewTogglePortraits
	ViewShowPortraits
	ViewHidePortraits

	CharExpAdd
	CharExpRemove

	CharLvlSet

	CharPointStatGive
	CharPointStatAllocate
	CharPointStatReset

	CharPointSkillGive
	CharPointSkillAllocate
	CharPointSkillReset

	SkillTreeEnable
	SkillTreeDisable

	EquipHelm
	EquipArmor
	EquipBelt
	EquipGloves
	EquipBoots
	EquipWieldLeft
	EquipWieldRight
	EquipRingLeft
	EquipRingRight
	EquipAmulet
	EquipOffhandLeft
	EquipOffhandRight
)

