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
		"hitbymissile":     HitByMissile,
		"damagedinmelee":   DamagedInMelee,
		"damagedbymissile": DamagedByMissile,
		"attackedinmelee":  AttackedInMelee,
		"doactive":         DoActive,
		"domeleedamage":    DoMeleeDamage,
		"domissiledamage":  DoMissileDamage,
		"domeleeattack":    DoMeleeAttack,
		"domissileattack":  DoMissileAttack,
		"kill":             Kill,
		"killed":           Killed,
		"absorbdamage":     AbsorbDamage,
		"levelup":          LevelUp,
	}
	return strLookupTable[s]
}
