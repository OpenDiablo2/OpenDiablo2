package d2enum

// from levels.txt, field `Teleport`
// https://d2mods.info/forum/kb/viewarticle?a=301

// TeleportFlag Controls if teleport is allowed in that level.
// 0 = Teleport not allowed
// 1 = Teleport allowed
// 2 = Teleport allowed, but not able to use teleport throu walls/objects
// (maybe for objects)
type TeleportFlag int

// Teleport flag types
const (
	TeleportDisabled TeleportFlag = iota
	TeleportEnabled
	TeleportEnabledLimited
)
