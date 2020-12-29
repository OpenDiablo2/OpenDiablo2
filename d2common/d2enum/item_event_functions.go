package d2enum

// ItemEventFuncID represents a item event function
type ItemEventFuncID int

// Item event functions
const (
	// shoots a missile at the owner of a missile that has just hit you
	// (Chilling Armor uses this)
	ReflectMissile ItemEventFuncID = iota

	// freezes the attacker for a set duration the attacker
	// (Frozen Armor uses this)
	FreezeAttacker

	// does cold damage to and chills the attacker (Shiver Armor uses this)
	FreezeChillAttacker

	// % of damage taken is done to the attacker
	// (Iron Maiden, thorns uses a hardcoded stat)
	ReflectPercentDamage

	// % of damage done added to life, bypassing the targets resistance
	// (used by Life Tap)
	DamageDealtToHealth

	// attacker takes physical damage of #
	AttackerTakesPhysical

	// knocks the target back
	Knockback

	// induces fear in the target making it run away
	InduceFear

	// applies Dim Vision to the target (it casts the actual curse on the
	// monster)
	BlindTarget

	// attacker takes lightning damage of #
	AttackerTakesLightning

	// attacker takes fire damage of #
	AttackerTakesFire

	// attacker takes cold damage of #
	AttackerTakesCold

	// % damage taken is added to mana
	DamageTakenToMana

	// freezes the target
	FreezeTarget

	// causes the target to bleed and lose life (negative life regen)
	OpenWounds

	// crushing blow against the target
	CrushingBlow

	// mana after killing a monster
	ManaOnKillMonster

	// life after killing a demon
	LifeOnKillDemon

	// slows the target
	SlowTarget

	// casts a skill against the defender
	CastSkillAgainstDefender

	// casts a skill against the attacker
	CastSkillAgainstAttacker

	// absorbs physical damage taken (used by Bone Armor)
	AbsorbPhysical

	// transfers damage done from the summon to the owner (used by Blood Golem)
	TakeSummonDamage

	// used by Energy Shield to absorb damage and shift it from life to mana
	ManaAbsorbsDamage

	// absorbs elemental damage taken (used by Cyclone Armor)
	AbsorbElementalDamage

	// transfers damage taken from the summon to the owner (used by Blood Golem)
	TakeSummonDamage2

	// used to slow the attacker if he hits a unit that has the slow target stat
	// (used by Clay Golem)
	TargetSlowsTarget

	// life after killing a monster
	LifeOnKillMonster

	// destroys the corpse of a killed monster (rest in peace effect)
	RestInPeace

	// cast a skill when the event occurs, without a target
	CastSkillWithoutTarget

	// reanimate the target as the specified monster
	ReanimateTargetAsMonster
)
