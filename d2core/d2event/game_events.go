package d2event

type GameEvent int

const (
	ScreenSet GameEvent = iota
	ScreenModeSet

	InputSetFocusMenu
	InputSetFocusWorld
	InputSetFocusPanel
	InputSetFocusTerminal

	AudioPlay
	AudioPlayPositional
	AudioFade
	AudioLoopStart
	AudioLoopStop
	AudioLoopOnRepeat

	TerminalOpen
	TerminalClose
	TerminalToggle

	PanelOpen
	PanelClose
	PanelToggle

	WorldPlayerJoin
	WorldPlayerLeave

	WorldGenerateMap
	WorldGenerateAutomap

	WorldAlterMap
	WorldAlterAutomap

	WorldSpawnEntity
	WorldSpawnItem
	WorldSpawnMonster
	WorldSpawnPlayer
	WorldSpawnMissile
	WorldSpawnMapFeature

	WorldInteractEntity
	WorldInteractItem
	WorldInteractNpc
	WorldInteractMapFeature

	WorldKillEntity
	WorldKillMonster
	WorldKillPlayer
	WorldKillNpc
	WorldKillMissile

	WorldRemoveEntity
	WorldRemoveItem
	WorldRemoveMonster
	WorldRemovePlayer
	WorldRemoveNpc
	WorldRemoveMissile
	WorldRemoveMapFeature

	WorldEntitySetPosition
	WorldEntitySetAnimationMode
	WorldEntitySetAnimationPlay
	WorldEntitySetAnimationPause
	WorldEntitySetDirection
	WorldEntitySetAlpha

	WorldEntityPathSet
	WorldEntityPathUnset

	PlayerResetAttributePoints
	PlayerResetSkillPoints
	PlayerAllocateAttributePoint
	PlayerAllocateSkillPoint

	PlayerMoveAttempt
	PlayerMoveSucceed
	PlayerMoveFailure
	PlayerMoveComplete

	PlayerEquipAttempt
	PlayerEquipSucceed
	PlayerEquipFailure
	PlayerEquipComplete

	PlayerAttackAttempt
	PlayerAttackSucceed
	PlayerAttackFailure
	PlayerAttackComplete

	PlayerSkillAttempt
	PlayerSkillSucceed
	PlayerSkillFailure
	PlayerSkillComplete

	PlayerItemStatsShow
	PlayerItemStatsHide
	PlayerItemPlaceAuto
	PlayerItemPlaceAttempt
	PlayerItemPlaceSucceed
	PlayerItemPlaceFailure
	PlayerItemPlaceComplete
	PlayerItemDrop

	PlayerItemUseAttempt
	PlayerItemUseSucceed
	PlayerItemUseFailure
	PlayerItemUseComplete

	PlayerTransmuteAttempt
	PlayerTransmuteSucceed
	PlayerTransmuteFailure
	PlayerTransmuteComplete

	PlayerQuestAttempt
	PlayerQuestSucceed
	PlayerQuestFailure
	PlayerQuestComplete

	PlayerBuyAttempt
	PlayerBuySucceed
	PlayerBuyFailure
	PlayerBuyComplete

	PlayerSellAttempt
	PlayerSellSucceed
	PlayerSellFailure
	PlayerSellComplete

	PlayerGambleAttempt
	PlayerGambleSucceed
	PlayerGambleFailure
	PlayerGambleComplete

	PlayerTradeRequest
	PlayerTradeAccept
	PlayerTradeDeny

	PlayerPartyRequest
	PlayerPartyAccept
	PlayerPartyDeny

	PlayerChatSend
	PlayerChatRecieve

	PlayerWhisperSend
	PlayerWhisperRecieve

	PlayerHostileAttempt
	PlayerHostileSucceed
	PlayerHostileFailure
	PlayerHostileComplete
	PlayerHostileEnable
	PlayerHostileDisable
	PlayerHostileToggle

	HirelingMoveAttempt
	HirelingMoveSucceed
	HirelingMoveFailure
	HirelingMoveComplete

	HirelingEquipAttempt
	HirelingEquipSucceed
	HirelingEquipFailure
	HirelingEquipComplete

	HirelingAttackAttempt
	HirelingAttackSucceed
	HirelingAttackFailure
	HirelingAttackComplete

	HirelingSkillAttempt
	HirelingSkillSucceed
	HirelingSkillFailure
	HirelingSkillComplete
)
