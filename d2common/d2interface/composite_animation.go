package d2interface

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

type CompositeAnimation interface {
	Advance(elapsed float64) error
	Render(target Surface) error
	GetAnimationMode() string
	GetWeaponClass() string
	SetMode(animationMode, weaponClass string) error
	Equip(equipment *[d2enum.CompositeTypeMax]string) error
	SetAnimSpeed(speed int)
	SetDirection(direction int)
	GetDirection() int
	GetPlayedCount() int
}
