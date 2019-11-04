package Common

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/ResourcePaths"
)

type MissileCalcParam struct {
	Param int
	Desc  string
}

type MissileCalc struct {
	Calc   string
	Desc   string
	Params []MissileCalcParam
}

type MissileLight struct {
	Diameter int
	Flicker  int
	Red      uint8
	Green    uint8
	Blue     uint8
}

type MissileAnimation struct {
	StepsBeforeVisible int
	StepsBeforeActive  int
	LoopAnimation      bool
	CelFileName        string
	AnimationRate      int // seems to do nothing
	AnimationLength    int
	AnimationSpeed     int
	StartingFrame      int // called "RandFrame"
	HasSubLoop         bool // runs after first animation ends
	SubStartingFrame   int
	SubEndingFrame     int
}

type MissileCollision struct {
	CollisionType          int // controls the kind of collision
	// 0 = none, 1 = units only, 3 = normal (units, walls), 
	// 6 = walls only, 8 = walls, units, and floors  
	DestroyedUponCollision bool
	FriendlyFire           bool
	LastCollide            bool // unknown
	Collision              bool // unknown
	ClientCollision        bool // unknown
	ClientSend             bool // unclear
	UseCollisionTimer      bool // after hit, use timer before dying
	TimerFrames            int // how many frames to persist
}

type MissileDamage struct {
	MinDamage              int
	MaxDamage              int
	MinLevelDamage         [5]int // additional damage per missile level
	// [0]: lvs 2-8, [1]: lvs 9-16, [2]: lvs 17-22, [3]: lvs 23-28, [4]: lv 29+
	MaxLevelDamage         [5]int // see above
	DamageSynergyPerCalc   string // works like synergy in skills.txt, not clear
}

type MissileElementalDamage struct {
	Damage        MissileDamage
	ElementType   string
	Duration      int // frames, 25 = 1 second
	LevelDuration [3]int // 0,1,2, unknown level intervals, bonus duration per level
}

type MissileRecord struct {
	Name                string
	Id                     int

	ClientMovementFunc     int
	ClientCollisionFunc    int
	ServerMovementFunc     int
	ServerCollisionFunc    int
	ServerDamageFunc       int
	ServerMovementCalc     MissileCalc
	ClientMovementCalc     MissileCalc
	ServerCollisionCalc    MissileCalc
	ClientCollisionCalc    MissileCalc
	ServerDamageCalc       MissileCalc

	Velocity               int
	MaxVelocity            int
	LevelVelocityBonus     int
	Accel                  int
	Range                  int
	LevelRangeBonus        int

	Light                  MissileLight

	Animation              MissileAnimation

	Collision              MissileCollision

	XOffset                int
	YOffset                int
	ZOffset                int
	Size                   int // diameter

	DestroyedByTP          bool // if true, destroyed when source player teleports to town
	DestroyedByTPFrame     int // see above, for client side, (this is a guess) which frame it vanishes on
	CanDestroy             bool // unknown

	UseAttackRating        bool // if true, uses 'attack rating' to determine if it hits or misses
	// if false, has a 95% chance to hit.
	AlwaysExplode          bool // if true, always calls its collision function when it is destroyed, even if it doesn't hit anything
	// note that some collision functions (lightning fury) seem to ignore this and always explode regardless of setting (requires investigation)

	ClientExplosion        bool // if true, does not really exist
	// is only aesthetic / client side
	TownSafe               bool // if true, doesn't vanish when spawned in town
	// if false, vanishes when spawned in town
	IgnoreBossModifiers    bool // if true, doesn't get bonuses from boss mods
	IgnoreMultishot        bool // if true, can't gain the mulitshot modifier
	HolyFilterType         int // specifies what this missile can hit
	// 0 = all units, 1 = undead only, 2 = demons only, 3 = all units (again?)
	CanBeSlowed            bool // if true, is affected by skill_handofathena
	TriggersHitEvents      bool // if true, triggers events that happen "upon getting hit" on targets
	TriggersGetHit         bool // if true, can cause target to enter hit recovery mode
	SoftHit                bool // unknown
	KnockbackPercent       int // chance of knocking the target back, 0-100

	TransparencyMode       int // controls rendering
	// 0 = normal, 1 = alpha blending (darker = more transparent)
	// 2 = special (black and white?)

	UseQuantity            bool // if true, uses quantity
	// not clear what this means. Also apparently requires a special starting function in skills.txt
	AffectedByPierce       bool // if true, affected by the pierce modifier and the Pierce skill
	SpecialSetup           bool // unknown, only true for potions

	MissileSkill           bool // if true, applies elemental damage from items to the splash radius instead of normal damage modifiers
	SkillName              string // if not empty, the missile will refer to this skill instead of its own data for the following:
	// "ResultFlags, HitFlags, HitShift, HitClass, SrcDamage (SrcDam in skills.txt!),
	// MinDam, MinLevDam1-5, MaxDam, MaxLevDam1-5, DmgSymPerCalc, EType, EMin, EMinLev1-5,
	// EMax, EMaxLev1-5, EDmgSymPerCalc, ELen, ELenLev1-3, ELenSymPerCalc"
	
	ResultFlags            int // unknown
	// 4 = normal missiles, 5 = explosions, 8 = non-damaging missiles
	HitFlags               int // unknown
	// 2 = explosions, 5 = freezing arrow

	HitShift               int // damage is measured in 256s
	// the actual damage is [damage] * 2 ^ [hitshift]
	// e.g. 100 damage, 8 hitshift = 100 * 2 ^ 8 = 100 * 256 = 25600
	// (visually, the damage is this result / 256)
	ApplyMastery           bool // unknown
	SourceDamage           int // 0-128, 128 is 100%
	// percentage of source units attack properties to apply to the missile?
	// not only affects damage but also other modifiers like lifesteal and manasteal (need a complete list)
	// setting this to -1 "gets rid of SrcDmg from skills.txt", not clear what that means
	HalfDamageForTwoHander bool // if true, damage is halved when a two-handed weapon is used
	SourceMissDamage       int // 0-128, 128 is 100%
	// unknown, only used for poison clouds.

	Damage                 MissileDamage
	ElementalDamage        MissileElementalDamage

	HitClass               int // controls clientside aesthetic effects for collisions
	// particularly sound effects that are played on a hit
	NumDirections          int // count of dirs in the DCC loaded by CelFile
	// apparently this value is no longer needed in D2
	LocalBlood             int // blood effects?
	// 0 = no blood, 1 = blood, 2 = blood and affected by open wounds
	DamageReductionRate    int // how many frames between applications of the 
	// magic_damage_reduced stat, so for instance on a 0 this stat applies every frame
	// on a 3, only every 4th frame has damage reduced

	TravelSound            string // name of sound to play during lifetime
	// whether or not it loops depends on the specific sound's settings?
	// if it doesn't loop, it's just a on-spawn sound effect
	HitSound               string // sound plays upon collision
	ProgSound              string // plays at "special events", like a mariachi band

	ProgOverlay            string // name of an overlay from overlays.txt to use at special events
	ExplosionMissile       string // name of a missile from missiles.txt that is created upon collision
	// or anytime it is destroyed if AlwaysExplode is true

	SubMissile             [3]string // 0,1,2 name of missiles spawned by movement function
	HitSubMissile          [4]string // 0,1,2 name of missiles spawned by collision function
	ClientSubMissile       [3]string // see above, but for client only
	ClientHitSubMissile    [4]string // see above, but for client only
}

func createMissileRecord(line string) MissileRecord {
	r := strings.Split(line, "\t")
	i := -1
	inc := func() int {
		i++
		return i
	}
	// note: in this file, empties are equivalent to zero, so all numerical conversions should
	//       be wrapped in an EmptyToZero transform
	result := MissileRecord{
		Name: r[inc()],
		Id: StringToInt(EmptyToZero(r[inc()])),

		ClientMovementFunc: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		ClientCollisionFunc: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		ServerMovementFunc: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		ServerCollisionFunc: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		ServerDamageFunc: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),

		ServerMovementCalc: loadMissileCalc(&r, inc, 5),
		ClientMovementCalc: loadMissileCalc(&r, inc, 5),
		ServerCollisionCalc: loadMissileCalc(&r, inc, 3),
		ClientCollisionCalc: loadMissileCalc(&r, inc, 3),
		ServerDamageCalc: loadMissileCalc(&r, inc, 2),
		
		Velocity: StringToInt(EmptyToZero(r[inc()])),
		MaxVelocity: StringToInt(EmptyToZero(r[inc()])),
		LevelVelocityBonus: StringToInt(EmptyToZero(r[inc()])),
		Accel: StringToInt(EmptyToZero(r[inc()])),
		Range: StringToInt(EmptyToZero(r[inc()])),
		LevelRangeBonus: StringToInt(EmptyToZero(r[inc()])),

		Light: loadMissileLight(&r, inc),

		Animation: loadMissileAnimation(&r, inc),

		Collision: loadMissileCollision(&r, inc),

		XOffset: StringToInt(EmptyToZero(r[inc()])),
		YOffset: StringToInt(EmptyToZero(r[inc()])),
		ZOffset: StringToInt(EmptyToZero(r[inc()])),
		Size: StringToInt(EmptyToZero(r[inc()])),

		DestroyedByTP: StringToInt(EmptyToZero(r[inc()])) == 1,
		DestroyedByTPFrame: StringToInt(EmptyToZero(r[inc()])),
		CanDestroy: StringToInt(EmptyToZero(r[inc()])) == 1,

		UseAttackRating: StringToInt(EmptyToZero(r[inc()])) == 1,
		AlwaysExplode: StringToInt(EmptyToZero(r[inc()])) == 1,

		ClientExplosion: StringToInt(EmptyToZero(r[inc()])) == 1,
		TownSafe: StringToInt(EmptyToZero(r[inc()])) == 1,
		IgnoreBossModifiers: StringToInt(EmptyToZero(r[inc()])) == 1,
		IgnoreMultishot: StringToInt(EmptyToZero(r[inc()])) == 1,
		HolyFilterType: StringToInt(EmptyToZero(r[inc()])),
		CanBeSlowed: StringToInt(EmptyToZero(r[inc()])) == 1,
		TriggersHitEvents: StringToInt(EmptyToZero(r[inc()])) == 1,
		TriggersGetHit: StringToInt(EmptyToZero(r[inc()])) == 1,
		SoftHit: StringToInt(EmptyToZero(r[inc()])) == 1,
		KnockbackPercent: StringToInt(EmptyToZero(r[inc()])),

		TransparencyMode: StringToInt(EmptyToZero(r[inc()])),

		UseQuantity: StringToInt(EmptyToZero(r[inc()])) == 1,
		AffectedByPierce: StringToInt(EmptyToZero(r[inc()])) == 1,
		SpecialSetup: StringToInt(EmptyToZero(r[inc()])) == 1,

		MissileSkill: StringToInt(EmptyToZero(r[inc()])) == 1,
		SkillName: r[inc()],

		ResultFlags: StringToInt(EmptyToZero(r[inc()])),
		HitFlags: StringToInt(EmptyToZero(r[inc()])),

		HitShift: StringToInt(EmptyToZero(r[inc()])),
		ApplyMastery: StringToInt(EmptyToZero(r[inc()])) == 1,
		SourceDamage: StringToInt(EmptyToZero(r[inc()])),
		HalfDamageForTwoHander: StringToInt(EmptyToZero(r[inc()])) == 1,
		SourceMissDamage: StringToInt(EmptyToZero(r[inc()])),

		Damage: loadMissileDamage(&r, inc),
		ElementalDamage: loadMissileElementalDamage(&r, inc),

		HitClass: StringToInt(EmptyToZero(r[inc()])),
		NumDirections: StringToInt(EmptyToZero(r[inc()])),
		LocalBlood: StringToInt(EmptyToZero(r[inc()])),
		DamageReductionRate: StringToInt(EmptyToZero(r[inc()])),

		TravelSound: r[inc()],
		HitSound: r[inc()],
		ProgSound: r[inc()],
		ProgOverlay: r[inc()],
		ExplosionMissile: r[inc()],

		SubMissile: [3]string{r[inc()], r[inc()], r[inc()]},
		HitSubMissile: [4]string{r[inc()], r[inc()], r[inc()], r[inc()]},
		ClientSubMissile: [3]string{r[inc()], r[inc()], r[inc()]},
		ClientHitSubMissile: [4]string{r[inc()], r[inc()], r[inc()], r[inc()]},
	}
	return result
}

var Missiles map[int]*MissileRecord

func LoadMissiles(fileProvider FileProvider) {
	Missiles = make(map[int]*MissileRecord)
	data := strings.Split(string(fileProvider.LoadFile(ResourcePaths.Missiles)), "\r\n")[1:]
	for _, line := range data {
		if len(line) == 0 {
			continue
		}
		rec := createMissileRecord(line)
		Missiles[rec.Id] = &rec
	}
	log.Printf("Loaded %d missiles", len(Missiles))
}

func loadMissileCalcParam(r *[]string, inc func() int) MissileCalcParam {
	result := MissileCalcParam{
		Param: StringToInt(EmptyToZero((*r)[inc()])),
		Desc: (*r)[inc()],
	}
	return result
}

func loadMissileCalc(r *[]string, inc func() int, params int) MissileCalc {
	result := MissileCalc{
		Calc: (*r)[inc()],
		Desc: (*r)[inc()],
	}
	result.Params = make([]MissileCalcParam, params)
	for p := 0; p < params; p++ {
		result.Params[p] = loadMissileCalcParam(r, inc);
	}
	return result
}

func loadMissileLight(r *[]string, inc func() int) MissileLight {
	result := MissileLight{
		Diameter: StringToInt(EmptyToZero((*r)[inc()])),
		Flicker: StringToInt(EmptyToZero((*r)[inc()])),
		Red: StringToUint8(EmptyToZero((*r)[inc()])),
		Green: StringToUint8(EmptyToZero((*r)[inc()])),
		Blue: StringToUint8(EmptyToZero((*r)[inc()])),
	}
	return result
}

func loadMissileAnimation(r *[]string, inc func() int) MissileAnimation {
	result := MissileAnimation{
		StepsBeforeVisible: StringToInt(EmptyToZero((*r)[inc()])),
		StepsBeforeActive: StringToInt(EmptyToZero((*r)[inc()])),
		LoopAnimation: StringToInt(EmptyToZero((*r)[inc()])) == 1,
		CelFileName: (*r)[inc()],
		AnimationRate: StringToInt(EmptyToZero((*r)[inc()])), 
		AnimationLength: StringToInt(EmptyToZero((*r)[inc()])),
		AnimationSpeed: StringToInt(EmptyToZero((*r)[inc()])),
		StartingFrame: StringToInt(EmptyToZero((*r)[inc()])), 
		HasSubLoop: StringToInt(EmptyToZero((*r)[inc()])) == 1, 
		SubStartingFrame: StringToInt(EmptyToZero((*r)[inc()])),
		SubEndingFrame: StringToInt(EmptyToZero((*r)[inc()])),
	}
	return result
}

func loadMissileCollision(r *[]string, inc func() int) MissileCollision {
	result := MissileCollision{
		CollisionType: StringToInt(EmptyToZero((*r)[inc()])), 
		DestroyedUponCollision: StringToInt(EmptyToZero((*r)[inc()])) == 1,
		FriendlyFire: StringToInt(EmptyToZero((*r)[inc()])) == 1,
		LastCollide: StringToInt(EmptyToZero((*r)[inc()])) == 1, 
		Collision: StringToInt(EmptyToZero((*r)[inc()])) == 1, 
		ClientCollision: StringToInt(EmptyToZero((*r)[inc()])) == 1, 
		ClientSend: StringToInt(EmptyToZero((*r)[inc()])) == 1, 
		UseCollisionTimer: StringToInt(EmptyToZero((*r)[inc()])) == 1, 
		TimerFrames: StringToInt(EmptyToZero((*r)[inc()])), 
	}
	return result
}

func loadMissileDamage(r *[]string, inc func() int) MissileDamage {
	result := MissileDamage{
		MinDamage: StringToInt(EmptyToZero((*r)[inc()])), 
		MinLevelDamage: [5]int{
			StringToInt(EmptyToZero((*r)[inc()])), 
			StringToInt(EmptyToZero((*r)[inc()])), 
			StringToInt(EmptyToZero((*r)[inc()])), 
			StringToInt(EmptyToZero((*r)[inc()])), 
			StringToInt(EmptyToZero((*r)[inc()])), 
		},
		MaxDamage: StringToInt(EmptyToZero((*r)[inc()])), 
		MaxLevelDamage: [5]int{
			StringToInt(EmptyToZero((*r)[inc()])), 
			StringToInt(EmptyToZero((*r)[inc()])), 
			StringToInt(EmptyToZero((*r)[inc()])), 
			StringToInt(EmptyToZero((*r)[inc()])), 
			StringToInt(EmptyToZero((*r)[inc()])), 
		},
		DamageSynergyPerCalc: (*r)[inc()], 
	}
	return result
}

func loadMissileElementalDamage(r *[]string, inc func() int) MissileElementalDamage {
	result := MissileElementalDamage{
		ElementType: (*r)[inc()],
		Damage: loadMissileDamage(r, inc),
		Duration: StringToInt(EmptyToZero((*r)[inc()])), 
		LevelDuration: [3]int{
			StringToInt(EmptyToZero((*r)[inc()])), 
			StringToInt(EmptyToZero((*r)[inc()])), 
			StringToInt(EmptyToZero((*r)[inc()])), 
		},
	}
	return result
}