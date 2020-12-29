package d2enum

// NPCActionType determines composite mode animations for NPC's as they move around
type NPCActionType int

// NPCAction types
// https://github.com/OpenDiablo2/OpenDiablo2/issues/811
const (
	NPCActionInvalid NPCActionType = iota
	NPCAction1
	NPCAction2
	NPCAction3
	NPCActionSkill1
)
