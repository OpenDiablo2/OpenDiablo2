package d2enum

type GameEvent int

const (
	// ToggleGameMenu will display the game menu
	ToggleGameMenu GameEvent = iota + 1

	// ToggleCharacterPanel toggles the player's character sheet
	ToggleCharacterPanel
	// ToggleInventoryPanel toggles the player's inventory
	ToggleInventoryPanel
	// TogglePartyPanel toggles the party manager
	TogglePartyPanel
	// ToggleSkillTreePanel toggles the skill tree
	ToggleSkillTreePanel
	// ToggleHirelingPanel toggles the hirelings manager
	ToggleHirelingPanel
	// ToggleQuestLog toggles the quest log
	ToggleQuestLog
	//ToggleHelpScreen toggles the help screen
	ToggleHelpScreen
	//ToggleChatOverlay toggles the chat overlay
	ToggleChatOverlay
	// ToggleMessageLog toggles the server's message logs
	ToggleMessageLog

	// ToggleAutomap toggles the automap
	ToggleAutomap
	// CenterAutomap recenters the automap when opened
	CenterAutomap
	// FadeAutomap reduces the brightness of the map
	// (not the players/npcs)
	FadeAutomap
	// TogglePartyOnAutomap toggles the display of the party
	// members on the automap
	TogglePartyOnAutomap
	// ToggleNamesOnAutomap toggles the display of the names for
	// party members and npcs on the automap
	ToggleNamesOnAutomap

	// ToggleRightSkillSelector toggles the right hand skill
	// selector overlay
	ToggleRightSkillSelector

	UseSkill1
	UseSkill2
	UseSkill3
	UseSkill4
	UseSkill5
	UseSkill6
	UseSkill7
	UseSkill8
	UseSkill9
	UseSkill10
	UseSkill11
	UseSkill12
	UseSkill13
	UseSkill14
	UseSkill15
	UseSkill16
	UseSkill17
	UseSkill18
	SelectPreviousSkill
	SelectNextSkill

	// ToggleBelts toggles the display of the different level for
	// the currently equipped belt
	ToggleBelts
	UseBeltSlot1
	UseBeltSlot2
	UseBeltSlot3
	UseBeltSlot4

	// SwapWeapons swaps the player's equipped weapons
	SwapWeapons
	// ToggleRunWalk toggles the run/walk movement
	ToggleRunWalk
	// HoldRun is fired while the player holds the
	// run key and and stopped when the key is released
	HoldRun
	// HoldStandStill prevents the player from moving while
	// the key is held
	HoldStandStill
	// HoldShowGroundItems displays a tooltip above the items
	// on the ground with their name
	HoldShowGroundItems
	HoldShowPortraits

	// ClearScreen closes all the panels currently open
	ClearScreen
)
