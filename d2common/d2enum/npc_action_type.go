package d2enum

// NPCActionType determines composite mode animations for NPC's as they move around
type NPCActionType int

// NPCAction types
// TODO: Figure out what 1-3 are for
const (
	NPCActionInvalid NPCActionType = iota
	NPCAction1
	NPCAction2
	NPCAction3
	NPCActionSkill1
)
