package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// MonStats2Record is a representation of a row from monstats2.txt
type MonStats2Record struct {
	// Key, the object ID MonStatEx feild from MonStat
	Key string

	// These follow three are apparently unused
	Height        int
	OverlayHeight int
	PixelHeight   int

	// Diameter in subtiles
	SizeX int
	SizeY int

	// This specifies if the size values get used for collision detection
	NoGfxHitTest bool

	// Bounding box
	BoxTop    int
	BoxLeft   int
	BoxWidth  int
	BoxHeight int

	// Spawn method used
	SpawnMethod int

	// Melee radius
	MeleeRng int

	// base weaponclass?
	BaseWeaponClass string
	HitClass        int

	// Available options for equipment
	// randomly selected from
	HDv []string
	TRv []string
	LGv []string
	Rav []string
	Lav []string
	RHv []string
	LHv []string
	SHv []string
	S1v []string
	S2v []string
	S3v []string
	S4v []string
	S5v []string
	S6v []string
	S7v []string
	S8v []string

	// Does the unit have this component
	HD bool
	TR bool
	LG bool
	RA bool
	LA bool
	RH bool
	LH bool
	SH bool
	S1 bool
	S2 bool
	S3 bool
	S4 bool
	S5 bool
	S6 bool
	S7 bool
	S8 bool

	// Sum of available components
	TotalPieces int

	// Available animation modes
	mDT bool
	mNU bool
	mWL bool
	mGH bool
	mA1 bool
	mA2 bool
	mBL bool
	mSC bool
	mS1 bool
	mS2 bool
	mS3 bool
	mS4 bool
	mDD bool
	mKB bool
	mSQ bool
	mRN bool

	// Number of directions for each mode
	dDT int
	dNU int
	dWL int
	dGH int
	dA1 int
	dA2 int
	dBL int
	dSC int
	dS1 int
	dS2 int
	dS3 int
	dS4 int
	dDD int
	dKB int
	dSQ int
	dRN int

	// Available modes while moving aside from WL and RN
	A1mv bool
	A2mv bool
	SCmv bool
	S1mv bool
	S2mv bool
	S3mv bool
	S4mv bool

	// If the units is restored on map reload
	Restore int

	// What maximap index is used for the automap
	AutomapCel int

	// true of unit uses an automap entry
	NoMap bool

	// If the units can use overlays
	NoOvly bool

	// If unit is selectable
	IsSelectable bool

	// If unit is selectable by allies
	AllySelectable bool

	// If unit is not selectable
	NotSelectable bool

	// Kinda unk, used for bonewalls etc that are not properly selectable
	shiftSel bool

	// if the units corpse is selectable
	IsCorpseSelectable bool

	// If the unit is attackable
	IsAttackable bool

	// If the unit is revivable
	IsRevivable bool

	// If the unit is a critter
	IsCritter bool

	// If the unit is Small, Small units can be knocked back with 100% efficiency
	IsSmall bool

	// Large units can be knocked back at 25% efficincy
	IsLarge bool

	// Possibly to do with sound, usually set for creatures without flesh
	IsSoft bool

	// Aggressive or harmless, usually NPC's
	IsInert bool

	// Unknown
	objCol bool

	// Enables collision on corpse for units
	IsCorpseCollidable bool

	// Can the corpse be walked through
	IsCorpseWalkable bool

	// If the unit casts a shadow
	HasShadow bool

	// If unique palettes should not be used
	NoUniqueShift bool

	// If multiple layers should be used on death (otherwise only TR)
	CompositeDeath bool

	// Blood offset?
	LocalBlood int

	// 0 = don't bleed, 1 = small blood missile, 2 = small and large, > 3 other missiles?
	Bleed int

	// If the unit is lights up the area
	Light int

	// Light color
	LightR int
	LightG int
	lightB int

	// Palettes per difficulty
	NormalPalette    int
	NightmarePalette int
	HellPalatte      int

	// These two are useless as of 1.07
	Heart    string
	BodyPart string

	// Inferno animation stuff
	InfernoLen      int
	InfernoAnim     int
	InfernoRollback int

	// Which mode is used after resurrection
	ResurrectMode d2enum.MonsterAnimationMode

	// Which skill is used for resurrection
	ResurrectSkill string
}

// MonStats2 stores all of the MonStats2Records
//nolint:gochecknoglobals // Current design issue
var MonStats2 map[string]*MonStats2Record

// LoadMonStats2 loads MonStats2Records from monstats2.txt
//nolint:funlen //just a big data loader
func LoadMonStats2(file []byte) {
	MonStats2 = make(map[string]*MonStats2Record, 0)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &MonStats2Record{
			Key:                d.String("Id"),
			Height:             d.Number("Height"),
			OverlayHeight:      d.Number("OverlayHeight"),
			PixelHeight:        d.Number("pixHeight"),
			SizeX:              d.Number("SizeX"),
			SizeY:              d.Number("SizeY"),
			SpawnMethod:        d.Number("spawnCol"),
			MeleeRng:           d.Number("MeleeRng"),
			BaseWeaponClass:    d.String("BaseW"),
			HitClass:           d.Number("HitClass"),
			HDv:                d.List("HDv"),
			TRv:                d.List("TRv"),
			LGv:                d.List("LGv"),
			Rav:                d.List("Rav"),
			Lav:                d.List("Lav"),
			RHv:                d.List("RHv"),
			LHv:                d.List("LHv"),
			SHv:                d.List("SHv"),
			S1v:                d.List("S1v"),
			S2v:                d.List("S2v"),
			S3v:                d.List("S3v"),
			S4v:                d.List("S4v"),
			S5v:                d.List("S5v"),
			S6v:                d.List("S6v"),
			S7v:                d.List("S7v"),
			S8v:                d.List("S8v"),
			HD:                 d.Bool("HD"),
			TR:                 d.Bool("TR"),
			LG:                 d.Bool("LG"),
			RA:                 d.Bool("RA"),
			LA:                 d.Bool("LA"),
			RH:                 d.Bool("RH"),
			LH:                 d.Bool("LH"),
			SH:                 d.Bool("SH"),
			S1:                 d.Bool("S1"),
			S2:                 d.Bool("S2"),
			S3:                 d.Bool("S3"),
			S4:                 d.Bool("S4"),
			S5:                 d.Bool("S5"),
			S6:                 d.Bool("S6"),
			S7:                 d.Bool("S7"),
			S8:                 d.Bool("S8"),
			TotalPieces:        d.Number("TotalPieces"),
			mDT:                d.Bool("mDT"),
			mNU:                d.Bool("mNU"),
			mWL:                d.Bool("mWL"),
			mGH:                d.Bool("mGH"),
			mA1:                d.Bool("mA1"),
			mA2:                d.Bool("mA2"),
			mBL:                d.Bool("mBL"),
			mSC:                d.Bool("mSC"),
			mS1:                d.Bool("mS1"),
			mS2:                d.Bool("mS2"),
			mS3:                d.Bool("mS3"),
			mS4:                d.Bool("mS4"),
			mDD:                d.Bool("mDD"),
			mKB:                d.Bool("mKB"),
			mSQ:                d.Bool("mSQ"),
			mRN:                d.Bool("mRN"),
			dDT:                d.Number("mDT"),
			dNU:                d.Number("mNU"),
			dWL:                d.Number("mWL"),
			dGH:                d.Number("mGH"),
			dA1:                d.Number("mA1"),
			dA2:                d.Number("mA2"),
			dBL:                d.Number("mBL"),
			dSC:                d.Number("mSC"),
			dS1:                d.Number("mS1"),
			dS2:                d.Number("mS2"),
			dS3:                d.Number("mS3"),
			dS4:                d.Number("mS4"),
			dDD:                d.Number("mDD"),
			dKB:                d.Number("mKB"),
			dSQ:                d.Number("mSQ"),
			dRN:                d.Number("mRN"),
			A1mv:               d.Bool("A1mv"),
			A2mv:               d.Bool("A2mv"),
			SCmv:               d.Bool("SCmv"),
			S1mv:               d.Bool("S1mv"),
			S2mv:               d.Bool("S2mv"),
			S3mv:               d.Bool("S3mv"),
			S4mv:               d.Bool("S4mv"),
			NoGfxHitTest:       d.Bool("noGfxHitTest"),
			BoxTop:             d.Number("htTop"),
			BoxLeft:            d.Number("htLeft"),
			BoxWidth:           d.Number("htWidth"),
			BoxHeight:          d.Number("htHeight"),
			Restore:            d.Number("restore"),
			AutomapCel:         d.Number("automapCel"),
			NoMap:              d.Bool("noMap"),
			NoOvly:             d.Bool("noOvly"),
			IsSelectable:       d.Bool("isSel"),
			AllySelectable:     d.Bool("alSel"),
			shiftSel:           d.Bool("shiftSel"),
			NotSelectable:      d.Bool("noSel"),
			IsCorpseSelectable: d.Bool("corpseSel"),
			IsAttackable:       d.Bool("isAtt"),
			IsRevivable:        d.Bool("revive"),
			IsCritter:          d.Bool("critter"),
			IsSmall:            d.Bool("small"),
			IsLarge:            d.Bool("large"),
			IsSoft:             d.Bool("soft"),
			IsInert:            d.Bool("inert"),
			objCol:             d.Bool("objCol"),
			IsCorpseCollidable: d.Bool("deadCol"),
			IsCorpseWalkable:   d.Bool("unflatDead"),
			HasShadow:          d.Bool("Shadow"),
			NoUniqueShift:      d.Bool("noUniqueShift"),
			CompositeDeath:     d.Bool("compositeDeath"),
			LocalBlood:         d.Number("localBlood"),
			Bleed:              d.Number("Bleed"),
			Light:              d.Number("Light"),
			LightR:             d.Number("light-r"),
			LightG:             d.Number("light-g"),
			lightB:             d.Number("light-b"),
			NormalPalette:      d.Number("Utrans"),
			NightmarePalette:   d.Number("Utrans(N)"),
			HellPalatte:        d.Number("Utrans(H)"),
			Heart:              d.String("Heart"),
			BodyPart:           d.String("BodyPart"),
			InfernoLen:         d.Number("InfernoLen"),
			InfernoAnim:        d.Number("InfernoAnim"),
			InfernoRollback:    d.Number("InfernoRollback"),
			ResurrectMode:      d2enum.MonsterAnimationModeFromString(d.String("ResurrectMode")),
			ResurrectSkill:     d.String("ResurrectSkill"),
		}
		MonStats2[record.Key] = record
	}
	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonStats2 records", len(MonStats2))
}
