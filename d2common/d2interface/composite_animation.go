package d2interface

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"


type AnimationMode interface {
	String() string
}

type CompositeAnimation interface {
	Advance(elapsed float64) error
	Render(target Surface) error
	GetAnimationMode() AnimationMode
	GetWeaponClass() string
	SetMode(mode AnimationMode, weaponClass string) error
	Equip(equipment *[d2enum.CompositeTypeMax]string) error
	SetAnimSpeed(speed int)
	SetDirection(direction int)
	GetDirection() int
	GetPlayedCount() int
	SetSubLoop(start, end int)
	SetPlayLoop(bool)
	SetCurrentFrame(int)
}
