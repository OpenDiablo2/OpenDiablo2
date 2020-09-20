package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation"

// Missiles stores all of the MissileRecords
type Missiles map[int]*MissileRecord

type missilesByName map[string]*MissileRecord

// MissileCalcParam is a calculation parameter for a missile
type MissileCalcParam struct {
	Param int
	Desc  string
}

// MissileCalc is a calculation for a missile
type MissileCalc struct {
	Calc   d2calculation.CalcString
	Desc   string
	Params []MissileCalcParam
}

// MissileLight has the parameters for missile lighting
type MissileLight struct {
	Diameter int
	Flicker  int
	Red      uint8
	Green    uint8
	Blue     uint8
}

// MissileAnimation stores parameters for a missile animation
type MissileAnimation struct {
	CelFileName        string
	StepsBeforeVisible int
	StepsBeforeActive  int
	AnimationRate      int // seems to do nothing
	AnimationLength    int
	AnimationSpeed     int
	StartingFrame      int // called "RandFrame"
	SubStartingFrame   int
	SubEndingFrame     int
	LoopAnimation      bool
	HasSubLoop         bool // runs after first animation ends
}

// MissileCollision  parameters for missile collision
type MissileCollision struct {
	CollisionType int // controls the kind of collision
	// 0 = none, 1 = units only, 3 = normal (units, walls),
	// 6 = walls only, 8 = walls, units, and floors
	TimerFrames            int // how many frames to persist
	DestroyedUponCollision bool
	FriendlyFire           bool
	LastCollide            bool // unknown
	Collision              bool // unknown
	ClientCollision        bool // unknown
	ClientSend             bool // unclear
	UseCollisionTimer      bool // after hit, use timer before dying
}

// MissileDamage parameters for calculating missile physical damage
type MissileDamage struct {
	MinDamage      int
	MaxDamage      int
	MinLevelDamage [5]int // additional damage per missile level
	// [0]: lvs 2-8, [1]: lvs 9-16, [2]: lvs 17-22, [3]: lvs 23-28, [4]: lv 29+
	MaxLevelDamage       [5]int                   // see above
	DamageSynergyPerCalc d2calculation.CalcString // works like synergy in skills.txt, not clear
}

// MissileElementalDamage parameters for calculating missile elemental damage
type MissileElementalDamage struct {
	Damage        MissileDamage
	ElementType   string
	Duration      int    // frames, 25 = 1 second
	LevelDuration [3]int // 0,1,2, unknown level intervals, bonus duration per level
}

// MissileRecord is a representation of a row from missiles.txt
type MissileRecord struct {
	ServerMovementCalc  MissileCalc
	ClientMovementCalc  MissileCalc
	ServerCollisionCalc MissileCalc
	ClientCollisionCalc MissileCalc
	ServerDamageCalc    MissileCalc
	Light               MissileLight
	Animation           MissileAnimation
	Collision           MissileCollision
	Damage              MissileDamage
	ElementalDamage     MissileElementalDamage
	SubMissile          [3]string // 0,1,2 name of missiles spawned by movement function
	HitSubMissile       [4]string // 0,1,2 name of missiles spawned by collision function
	ClientSubMissile    [3]string // see above, but for client only
	ClientHitSubMissile [4]string // see above, but for client only
	Name                string

	SkillName string // if not empty, the missile will refer to this skill instead of its own data for the following:
	// "ResultFlags, HitFlags, HitShift, HitClass, SrcDamage (SrcDam in skills.txt!),
	// MinDam, MinLevDam1-5, MaxDam, MaxLevDam1-5, DmgSymPerCalc, EType, EMin, EMinLev1-5,
	// EMax, EMaxLev1-5, EDmgSymPerCalc, ELen, ELenLev1-3, ELenSymPerCalc"

	TravelSound string // name of sound to play during lifetime
	// whether or not it loops depends on the specific sound's settings?
	// if it doesn't loop, it's just a on-spawn sound effect
	HitSound  string // sound plays upon collision
	ProgSound string // plays at "special events", like a mariachi band

	ProgOverlay      string // name of an overlay from overlays.txt to use at special events
	ExplosionMissile string // name of a missile from missiles.txt that is created upon collision
	// or anytime it is destroyed if AlwaysExplode is true

	Id int //nolint:golint,stylecheck // ID is the correct key

	ClientMovementFunc  int
	ClientCollisionFunc int
	ServerMovementFunc  int
	ServerCollisionFunc int
	ServerDamageFunc    int

	Velocity           int
	MaxVelocity        int
	LevelVelocityBonus int
	Accel              int
	Range              int
	LevelRangeBonus    int

	XOffset int
	YOffset int
	ZOffset int
	Size    int // diameter

	DestroyedByTPFrame int // see above, for client side, (this is a guess) which frame it vanishes on

	HolyFilterType   int // specifies what this missile can hit
	KnockbackPercent int // chance of knocking the target back, 0-100

	TransparencyMode int // controls rendering
	// 0 = normal, 1 = alpha blending (darker = more transparent)
	// 2 = special (black and white?)

	ResultFlags int // unknown
	// 4 = normal missiles, 5 = explosions, 8 = non-damaging missiles
	HitFlags int // unknown
	// 2 = explosions, 5 = freezing arrow

	HitShift int // damage is measured in 256s
	// the actual damage is [damage] * 2 ^ [hitshift]
	// e.g. 100 damage, 8 hitshift = 100 * 2 ^ 8 = 100 * 256 = 25600
	// (visually, the damage is this result / 256)

	SourceDamage     int // 0-128, 128 is 100%
	SourceMissDamage int // 0-128, 128 is 100%
	// unknown, only used for poison clouds.

	HitClass int // controls clientside aesthetic effects for collisions
	// particularly sound effects that are played on a hit
	NumDirections int // count of dirs in the DCC loaded by CelFile
	// apparently this value is no longer needed in D2
	LocalBlood int // blood effects?
	// 0 = no blood, 1 = blood, 2 = blood and affected by open wounds
	DamageReductionRate int // how many frames between applications of the
	// magic_damage_reduced stat, so for instance on a 0 this stat applies every frame
	// on a 3, only every 4th frame has damage reduced

	DestroyedByTP bool // if true, destroyed when source player teleports to town
	CanDestroy    bool // unknown

	UseAttackRating bool // if true, uses 'attack rating' to determine if it hits or misses
	// if false, has a 95% chance to hit.
	AlwaysExplode bool // if true, always calls its collision function when it is destroyed,
	// even if it doesn't hit anything
	// note that some collision functions (lightning fury)
	// seem to ignore this and always explode regardless of setting (requires investigation)

	ClientExplosion bool // if true, does not really exist
	// is only aesthetic / client side
	TownSafe bool // if true, doesn't vanish when spawned in town
	// if false, vanishes when spawned in town
	IgnoreBossModifiers bool // if true, doesn't get bonuses from boss mods
	IgnoreMultishot     bool // if true, can't gain the mulitshot modifier
	// 0 = all units, 1 = undead only, 2 = demons only, 3 = all units (again?)
	CanBeSlowed       bool // if true, is affected by skill_handofathena
	TriggersHitEvents bool // if true, triggers events that happen "upon getting hit" on targets
	TriggersGetHit    bool // if true, can cause target to enter hit recovery mode
	SoftHit           bool // unknown

	UseQuantity bool // if true, uses quantity
	// not clear what this means. Also apparently requires a special starting function in skills.txt
	AffectedByPierce bool // if true, affected by the pierce modifier and the Pierce skill
	SpecialSetup     bool // unknown, only true for potions

	MissileSkill bool // if true, applies elemental damage from items to the splash radius instead of normal damage modifiers

	ApplyMastery bool // unknown
	// percentage of source units attack properties to apply to the missile?
	// not only affects damage but also other modifiers like lifesteal and manasteal (need a complete list)
	// setting this to -1 "gets rid of SrcDmg from skills.txt", not clear what that means
	HalfDamageForTwoHander bool // if true, damage is halved when a two-handed weapon is used

}
