package d2enum

// used in ItemStatCost
type ItemEventType int

const (
	HitByMissile     = ItemEventType(iota) // hit By a Missile
	DamagedInMelee                         // Damaged in Melee
	DamagedByMissile                       // Damaged By Missile
	AttackedInMelee                        // melee Attack atttempt
	DoActive                               // do active state skill
	DoMeleeDamage                          // do damage in melee
	DoMissileDamage                        // do missile damage
	DoMeleeAttack                          // do melee attack
	DoMissileAttack                        // do missile attack
	Kill                                   // killed something
	Killed                                 // killed By something
	AbsorbDamage                           // dealt damage
	LevelUp                                // gain a level
)

func GetItemEventType(s string) ItemEventType {
	strLookupTable := map[string]ItemEventType{
		"HitByMissile":     HitByMissile,
		"DamagedInMelee":   DamagedInMelee,
		"DamagedByMissile": DamagedByMissile,
		"AttackedInMelee":  AttackedInMelee,
		"DoActive":         DoActive,
		"DoMeleeDamage":    DoMeleeDamage,
		"DoMissileDamage":  DoMissileDamage,
		"DoMeleeAttack":    DoMeleeAttack,
		"DoMissileAttack":  DoMissileAttack,
		"Kill":             Kill,
		"Killed":           Killed,
		"AbsorbDamage":     AbsorbDamage,
		"LevelUp":          LevelUp,
	}
	return strLookupTable[s]
}
