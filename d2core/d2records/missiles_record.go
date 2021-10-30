package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation"

// Missiles stores all of the MissileRecords
type Missiles map[int]*MissileRecord

type missilesByName map[string]*MissileRecord

// MissileCalcParam is a calculation parameter for a missile
type MissileCalcParam struct {
	Desc  string
	Param int
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
	DamageSynergyPerCalc d2calculation.CalcString
	MinLevelDamage       [5]int
	MaxLevelDamage       [5]int
	MinDamage            int
	MaxDamage            int
}

// MissileElementalDamage parameters for calculating missile elemental damage
type MissileElementalDamage struct {
	ElementType   string
	Damage        MissileDamage
	LevelDuration [3]int
	Duration      int
}

// MissileRecord is a representation of a row from missiles.txt
type MissileRecord struct {
	ClientHitSubMissile    [4]string
	HitSubMissile          [4]string
	ClientSubMissile       [3]string
	SubMissile             [3]string
	ProgSound              string
	ExplosionMissile       string
	HitSound               string
	Name                   string
	ProgOverlay            string
	TravelSound            string
	SkillName              string
	ServerCollisionCalc    MissileCalc
	ServerMovementCalc     MissileCalc
	ClientMovementCalc     MissileCalc
	ServerDamageCalc       MissileCalc
	ClientCollisionCalc    MissileCalc
	Animation              MissileAnimation
	Damage                 MissileDamage
	ElementalDamage        MissileElementalDamage
	Collision              MissileCollision
	Light                  MissileLight
	ServerMovementFunc     int
	ClientMovementFunc     int
	ClientCollisionFunc    int
	ID                     int
	ServerCollisionFunc    int
	ServerDamageFunc       int
	Velocity               int
	MaxVelocity            int
	LevelVelocityBonus     int
	Accel                  int
	Range                  int
	LevelRangeBonus        int
	XOffset                int
	DamageReductionRate    int
	ZOffset                int
	Size                   int
	DestroyedByTPFrame     int
	HolyFilterType         int
	KnockbackPercent       int
	TransparencyMode       int
	ResultFlags            int
	HitFlags               int
	HitShift               int
	SourceDamage           int
	SourceMissDamage       int
	HitClass               int
	NumDirections          int
	LocalBlood             int
	YOffset                int
	DestroyedByTP          bool
	CanDestroy             bool
	UseAttackRating        bool
	AlwaysExplode          bool
	ClientExplosion        bool
	TownSafe               bool
	IgnoreBossModifiers    bool
	IgnoreMultishot        bool
	CanBeSlowed            bool
	TriggersHitEvents      bool
	TriggersGetHit         bool
	SoftHit                bool
	UseQuantity            bool
	AffectedByPierce       bool
	SpecialSetup           bool
	MissileSkill           bool
	ApplyMastery           bool
	HalfDamageForTwoHander bool
}
