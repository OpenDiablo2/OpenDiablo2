package d2enum

// from levels.txt

type TeleportFlag int

const (
	TeleportDisabled TeleportFlag = iota
	TeleportEnabled
	TeleportEnabledLimited
)
