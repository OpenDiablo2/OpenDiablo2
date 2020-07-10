package d2enum

// ItemEventType  used in ItemStatCost
type ItemEventType int

// Item event types
const (
	ItemEventNone             ItemEventType = iota
	ItemEventHitByMissile                   // hit By a Missile
	ItemEventDamagedInMelee                 // Damaged in Melee
	ItemEventDamagedByMissile               // Damaged By Missile
	ItemEventAttackedInMelee                // melee Attack atttempt
	ItemEventDoActive                       // do active state skill
	ItemEventDoMeleeDamage                  // do damage in melee
	ItemEventDoMissileDamage                // do missile damage
	ItemEventDoMeleeAttack                  // do melee attack
	ItemEventDoMissileAttack                // do missile attack
	ItemEventKill                           // killed something
	ItemEventKilled                         // killed By something
	ItemEventAbsorbDamage                   // dealt damage
	ItemEventLevelUp                        // gain a level
)

//nolint:gochecknoglobals // better for lookup
var itemEventsLookup = map[string]ItemEventType{
	"hitbymissile":     ItemEventHitByMissile,
	"damagedinmelee":   ItemEventDamagedInMelee,
	"damagedbymissile": ItemEventDamagedByMissile,
	"attackedinmelee":  ItemEventAttackedInMelee,
	"doactive":         ItemEventDoActive,
	"domeleedamage":    ItemEventDoMeleeDamage,
	"domissiledamage":  ItemEventDoMissileDamage,
	"domeleeattack":    ItemEventDoMeleeAttack,
	"domissileattack":  ItemEventDoMissileAttack,
	"kill":             ItemEventKill,
	"killed":           ItemEventKilled,
	"absorbdamage":     ItemEventAbsorbDamage,
	"levelup":          ItemEventLevelUp,
}

// GetItemEventType returns the ItemEventType from string, expects lowercase input
func GetItemEventType(s string) ItemEventType {
	if s == "" {
		return ItemEventNone
	}

	if v, ok := itemEventsLookup[s]; ok {
		return v
	}

	return ItemEventNone
}
