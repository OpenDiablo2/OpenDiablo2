package d2datadict

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

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

func createMissileRecord(line string) MissileRecord {
	r := strings.Split(line, "\t")
	i := -1
	inc := func() int {
		i++
		return i
	}
	// note: in this file, empties are equivalent to zero, so all numerical conversions should
	//       be wrapped in an d2common.EmptyToZero transform
	result := MissileRecord{
		Name: r[inc()],
		Id:   d2util.StringToInt(d2util.EmptyToZero(r[inc()])),

		ClientMovementFunc:  d2util.StringToInt(d2util.EmptyToZero(d2util.AsterToEmpty(r[inc()]))),
		ClientCollisionFunc: d2util.StringToInt(d2util.EmptyToZero(d2util.AsterToEmpty(r[inc()]))),
		ServerMovementFunc:  d2util.StringToInt(d2util.EmptyToZero(d2util.AsterToEmpty(r[inc()]))),
		ServerCollisionFunc: d2util.StringToInt(d2util.EmptyToZero(d2util.AsterToEmpty(r[inc()]))),
		ServerDamageFunc:    d2util.StringToInt(d2util.EmptyToZero(d2util.AsterToEmpty(r[inc()]))),

		ServerMovementCalc:  loadMissileCalc(&r, inc, 5),
		ClientMovementCalc:  loadMissileCalc(&r, inc, 5),
		ServerCollisionCalc: loadMissileCalc(&r, inc, 3),
		ClientCollisionCalc: loadMissileCalc(&r, inc, 3),
		ServerDamageCalc:    loadMissileCalc(&r, inc, 2),

		Velocity:           d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		MaxVelocity:        d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		LevelVelocityBonus: d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		Accel:              d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		Range:              d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		LevelRangeBonus:    d2util.StringToInt(d2util.EmptyToZero(r[inc()])),

		Light: loadMissileLight(&r, inc),

		Animation: loadMissileAnimation(&r, inc),

		Collision: loadMissileCollision(&r, inc),

		XOffset: d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		YOffset: d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		ZOffset: d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		Size:    d2util.StringToInt(d2util.EmptyToZero(r[inc()])),

		DestroyedByTP:      d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		DestroyedByTPFrame: d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		CanDestroy:         d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,

		UseAttackRating: d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		AlwaysExplode:   d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,

		ClientExplosion:     d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		TownSafe:            d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		IgnoreBossModifiers: d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		IgnoreMultishot:     d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		HolyFilterType:      d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		CanBeSlowed:         d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		TriggersHitEvents:   d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		TriggersGetHit:      d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		SoftHit:             d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		KnockbackPercent:    d2util.StringToInt(d2util.EmptyToZero(r[inc()])),

		TransparencyMode: d2util.StringToInt(d2util.EmptyToZero(r[inc()])),

		UseQuantity:      d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		AffectedByPierce: d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		SpecialSetup:     d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,

		MissileSkill: d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		SkillName:    r[inc()],

		ResultFlags: d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		HitFlags:    d2util.StringToInt(d2util.EmptyToZero(r[inc()])),

		HitShift:               d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		ApplyMastery:           d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		SourceDamage:           d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		HalfDamageForTwoHander: d2util.StringToInt(d2util.EmptyToZero(r[inc()])) == 1,
		SourceMissDamage:       d2util.StringToInt(d2util.EmptyToZero(r[inc()])),

		Damage:          loadMissileDamage(&r, inc),
		ElementalDamage: loadMissileElementalDamage(&r, inc),

		HitClass:            d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		NumDirections:       d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		LocalBlood:          d2util.StringToInt(d2util.EmptyToZero(r[inc()])),
		DamageReductionRate: d2util.StringToInt(d2util.EmptyToZero(r[inc()])),

		TravelSound:      r[inc()],
		HitSound:         r[inc()],
		ProgSound:        r[inc()],
		ProgOverlay:      r[inc()],
		ExplosionMissile: r[inc()],

		SubMissile:          [3]string{r[inc()], r[inc()], r[inc()]},
		HitSubMissile:       [4]string{r[inc()], r[inc()], r[inc()], r[inc()]},
		ClientSubMissile:    [3]string{r[inc()], r[inc()], r[inc()]},
		ClientHitSubMissile: [4]string{r[inc()], r[inc()], r[inc()], r[inc()]},
	}

	return result
}

// Missiles stores all of the MissileRecords
//nolint:gochecknoglobals // Currently global by design, only written once
var Missiles map[int]*MissileRecord

// LoadMissiles loads MissileRecords from missiles.txt
func LoadMissiles(file []byte) {
	Missiles = make(map[int]*MissileRecord)
	data := strings.Split(string(file), "\r\n")[1:]

	for _, line := range data {
		if line == "" {
			continue
		}

		rec := createMissileRecord(line)
		Missiles[rec.Id] = &rec
	}

	log.Printf("Loaded %d missiles", len(Missiles))
}

func loadMissileCalcParam(r *[]string, inc func() int) MissileCalcParam {
	result := MissileCalcParam{
		Param: d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		Desc:  (*r)[inc()],
	}

	return result
}

func loadMissileCalc(r *[]string, inc func() int, params int) MissileCalc {
	result := MissileCalc{
		Calc: d2calculation.CalcString((*r)[inc()]),
		Desc: (*r)[inc()],
	}
	result.Params = make([]MissileCalcParam, params)

	for p := 0; p < params; p++ {
		result.Params[p] = loadMissileCalcParam(r, inc)
	}

	return result
}

func loadMissileLight(r *[]string, inc func() int) MissileLight {
	result := MissileLight{
		Diameter: d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		Flicker:  d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		Red:      d2util.StringToUint8(d2util.EmptyToZero((*r)[inc()])),
		Green:    d2util.StringToUint8(d2util.EmptyToZero((*r)[inc()])),
		Blue:     d2util.StringToUint8(d2util.EmptyToZero((*r)[inc()])),
	}

	return result
}

func loadMissileAnimation(r *[]string, inc func() int) MissileAnimation {
	result := MissileAnimation{
		StepsBeforeVisible: d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		StepsBeforeActive:  d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		LoopAnimation:      d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])) == 1,
		CelFileName:        (*r)[inc()],
		AnimationRate:      d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		AnimationLength:    d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		AnimationSpeed:     d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		StartingFrame:      d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		HasSubLoop:         d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])) == 1,
		SubStartingFrame:   d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		SubEndingFrame:     d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
	}

	return result
}

func loadMissileCollision(r *[]string, inc func() int) MissileCollision {
	result := MissileCollision{
		CollisionType:          d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		DestroyedUponCollision: d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])) == 1,
		FriendlyFire:           d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])) == 1,
		LastCollide:            d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])) == 1,
		Collision:              d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])) == 1,
		ClientCollision:        d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])) == 1,
		ClientSend:             d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])) == 1,
		UseCollisionTimer:      d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])) == 1,
		TimerFrames:            d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
	}

	return result
}

func loadMissileDamage(r *[]string, inc func() int) MissileDamage {
	result := MissileDamage{
		MinDamage: d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		MinLevelDamage: [5]int{
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		},
		MaxDamage: d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		MaxLevelDamage: [5]int{
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		},
		DamageSynergyPerCalc: d2calculation.CalcString((*r)[inc()]),
	}

	return result
}

func loadMissileElementalDamage(r *[]string, inc func() int) MissileElementalDamage {
	result := MissileElementalDamage{
		ElementType: (*r)[inc()],
		Damage:      loadMissileDamage(r, inc),
		Duration:    d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		LevelDuration: [3]int{
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
			d2util.StringToInt(d2util.EmptyToZero((*r)[inc()])),
		},
	}

	return result
}
