package d2event

type handlerFunction interface{}

type handlerEntryMap map[GameEvent][]interface{}

var handlerEntries handlerEntryMap

type gameEventManager struct {
	entries handlerEntryMap
}

func (gem *gameEventManager) getEntries(e GameEvent) interface{} {
	return handlerEntries[e]
}

func Initialize() {
	handlerEntries = map[GameEvent]([]interface{}){
		ScreenSet:     make([]interface{}, 0),
		ScreenModeSet: make([]interface{}, 0),

		InputSetFocusMenu:     make([]interface{}, 0),
		InputSetFocusWorld:    make([]interface{}, 0),
		InputSetFocusPanel:    make([]interface{}, 0),
		InputSetFocusTerminal: make([]interface{}, 0),

		AudioPlay:           make([]interface{}, 0),
		AudioPlayPositional: make([]interface{}, 0),
		AudioFade:           make([]interface{}, 0),
		AudioLoopStart:      make([]interface{}, 0),
		AudioLoopStop:       make([]interface{}, 0),
		AudioLoopOnRepeat:   make([]interface{}, 0),

		TerminalOpen:   make([]interface{}, 0),
		TerminalClose:  make([]interface{}, 0),
		TerminalToggle: make([]interface{}, 0),

		PanelOpen:   make([]interface{}, 0),
		PanelClose:  make([]interface{}, 0),
		PanelToggle: make([]interface{}, 0),

		WorldPlayerJoin:  make([]interface{}, 0),
		WorldPlayerLeave: make([]interface{}, 0),

		WorldGenerateMap:  make([]interface{}, 0),
		WorldGenerateAuto: make([]interface{}, 0),

		WorldAlterMap:  make([]interface{}, 0),
		WorldAlterAuto: make([]interface{}, 0),

		WorldSpawnEntity:     make([]interface{}, 0),
		WorldSpawnItem:       make([]interface{}, 0),
		WorldSpawnMonster:    make([]interface{}, 0),
		WorldSpawnPlayer:     make([]interface{}, 0),
		WorldSpawnMissile:    make([]interface{}, 0),
		WorldSpawnMapFeature: make([]interface{}, 0),

		WorldInteractEntity:     make([]interface{}, 0),
		WorldInteractItem:       make([]interface{}, 0),
		WorldInteractNpc:        make([]interface{}, 0),
		WorldInteractMapFeature: make([]interface{}, 0),

		WorldKillEntity:  make([]interface{}, 0),
		WorldKillMonster: make([]interface{}, 0),
		WorldKillPlayer:  make([]interface{}, 0),
		WorldKillNpc:     make([]interface{}, 0),
		WorldKillMissile: make([]interface{}, 0),

		WorldRemoveEntity:  make([]interface{}, 0),
		WorldRemoveItem:    make([]interface{}, 0),
		WorldRemoveMonster: make([]interface{}, 0),
		WorldRemovePlayer:  make([]interface{}, 0),
		WorldRemoveNpc:     make([]interface{}, 0),
		WorldRemoveMissile: make([]interface{}, 0),GameEventHandler

		WorldEntitySetPosition:       make([]interface{}, 0),
		WorldEntitySetAnimationMode:  make([]interface{}, 0),
		WorldEntitySetAnimationPlay:  make([]interface{}, 0),
		WorldEntitySetAnimationPause: make([]interface{}, 0),
		WorldEntitySetDirection:      make([]interface{}, 0),
		WorldEntitySetAlpha:          make([]interface{}, 0),

		WorldEntityPathSet:   make([]interface{}, 0),
		WorldEntityPathUnset: make([]interface{}, 0),

		PlayerResetAttributePoints:   make([]interface{}, 0),
		PlayerResetSkillPoints:       make([]interface{}, 0),
		PlayerAllocateAttributePoint: make([]interface{}, 0),
		PlayerAllocateSkillPoint:     make([]interface{}, 0),

		PlayerMoveAttempt:  make([]interface{}, 0),
		PlayerMoveSucceed:  make([]interface{}, 0),
		PlayerMoveFailure:  make([]interface{}, 0),
		PlayerMoveComplete: make([]interface{}, 0),

		PlayerEquipAttempt:  make([]interface{}, 0),
		PlayerEquipSucceed:  make([]interface{}, 0),
		PlayerEquipFailure:  make([]interface{}, 0),
		PlayerEquipComplete: make([]interface{}, 0),

		PlayerAttackAttempt:  make([]interface{}, 0),
		PlayerAttackSucceed:  make([]interface{}, 0),
		PlayerAttackFailure:  make([]interface{}, 0),
		PlayerAttackComplete: make([]interface{}, 0),

		PlayerSkillAttempt:  make([]interface{}, 0),
		PlayerSkillSucceed:  make([]interface{}, 0),
		PlayerSkillFailure:  make([]interface{}, 0),
		PlayerSkillComplete: make([]interface{}, 0),

		PlayerItemStatsShow:     make([]interface{}, 0),
		PlayerItemStatsHide:     make([]interface{}, 0),
		PlayerItemPlaceAuto:     make([]interface{}, 0),
		PlayerItemPlaceAttempt:  make([]interface{}, 0),
		PlayerItemPlaceSucceed:  make([]interface{}, 0),
		PlayerItemPlaceFailure:  make([]interface{}, 0),
		PlayerItemPlaceComplete: make([]interface{}, 0),
		PlayerItemDrop:          make([]interface{}, 0),

		PlayerItemUseAttempt:  make([]interface{}, 0),
		PlayerItemUseSucceed:  make([]interface{}, 0),
		PlayerItemUseFailure:  make([]interface{}, 0),
		PlayerItemUseComplete: make([]interface{}, 0),

		PlayerTransmuteAttempt:  make([]interface{}, 0),
		PlayerTransmuteSucceed:  make([]interface{}, 0),
		PlayerTransmuteFailure:  make([]interface{}, 0),
		PlayerTransmuteComplete: make([]interface{}, 0),

		PlayerQuestAttempt:  make([]interface{}, 0),
		PlayerQuestSucceed:  make([]interface{}, 0),
		PlayerQuestFailure:  make([]interface{}, 0),
		PlayerQuestComplete: make([]interface{}, 0),

		PlayerBuyAttempt:  make([]interface{}, 0),
		PlayerBuySucceed:  make([]interface{}, 0),
		PlayerBuyFailure:  make([]interface{}, 0),
		PlayerBuyComplete: make([]interface{}, 0),

		PlayerSellAttempt:  make([]interface{}, 0),
		PlayerSellSucceed:  make([]interface{}, 0),
		PlayerSellFailure:  make([]interface{}, 0),
		PlayerSellComplete: make([]interface{}, 0),

		PlayerGambleAttempt:  make([]interface{}, 0),
		PlayerGambleSucceed:  make([]interface{}, 0),
		PlayerGambleFailure:  make([]interface{}, 0),
		PlayerGambleComplete: make([]interface{}, 0),

		PlayerTradeRequest: make([]interface{}, 0),
		PlayerTradeAccept:  make([]interface{}, 0),
		PlayerTradeDeny:    make([]interface{}, 0),

		PlayerPartyRequest: make([]interface{}, 0),
		PlayerPartyAccept:  make([]interface{}, 0),
		PlayerPartyDeny:    make([]interface{}, 0),

		PlayerChatSend:    make([]interface{}, 0),
		PlayerChatRecieve: make([]interface{}, 0),

		PlayerWhisperSend:    make([]interface{}, 0),
		PlayerWhisperRecieve: make([]interface{}, 0),

		PlayerHostileAttempt:  make([]interface{}, 0),
		PlayerHostileSucceed:  make([]interface{}, 0),
		PlayerHostileFailure:  make([]interface{}, 0),
		PlayerHostileComplete: make([]interface{}, 0),
		PlayerHostileEnable:   make([]interface{}, 0),
		PlayerHostileDisable:  make([]interface{}, 0),
		PlayerHostileToggle:   make([]interface{}, 0),

		HirelingMoveAttempt:  make([]interface{}, 0),
		HirelingMoveSucceed:  make([]interface{}, 0),
		HirelingMoveFailure:  make([]interface{}, 0),
		HirelingMoveComplete: make([]interface{}, 0),

		HirelingEquipAttempt:  make([]interface{}, 0),
		HirelingEquipSucceed:  make([]interface{}, 0),
		HirelingEquipFailure:  make([]interface{}, 0),
		HirelingEquipComplete: make([]interface{}, 0),

		HirelingAttackAttempt:  make([]interface{}, 0),
		HirelingAttackSucceed:  make([]interface{}, 0),
		HirelingAttackFailure:  make([]interface{}, 0),
		HirelingAttackComplete: make([]interface{}, 0),

		HirelingSkillAttempt:  make([]interface{}, 0),
		HirelingSkillSucceed:  make([]interface{}, 0),
		HirelingSkillFailure:  make([]interface{}, 0),
		HirelingSkillComplete: make([]interface{}, 0),
	}
}

func BindEventHandler(e GameEvent, fn interface{}) {

}

func UnbindEventHandler(e GameEvent, fn interface{}) {

}

func EmitEvent(event GameEvent, args ...interface{}) {
	
}
