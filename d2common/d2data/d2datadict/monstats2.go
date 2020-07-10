package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// MonStats2Record is a representation of a row from monstats2.txt
type MonStats2Record struct {
	// Available options for equipment
	// randomly selected from
	EquipmentOptions [16][]string

	Key             string // Key, the object ID MonStatEx feild from MonStat
	BaseWeaponClass string
	ResurrectSkill  string
	Heart           string
	BodyPart        string

	// These follow three are apparently unused
	Height        int
	OverlayHeight int
	PixelHeight   int

	// Diameter in subtiles
	SizeX int
	SizeY int

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
	HitClass int

	// Sum of available components
	TotalPieces int

	// Number of directions for each mode
	DirectionsPerMode [16]int

	// If the units is restored on map reload
	Restore int

	// What maximap index is used for the automap
	AutomapCel int

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

	// Inferno animation stuff
	InfernoLen      int
	InfernoAnim     int
	InfernoRollback int
	// Which mode is used after resurrection
	ResurrectMode d2enum.MonsterAnimationMode

	// This specifies if the size values get used for collision detection
	NoGfxHitTest bool

	// Does the unit have this component
	HasComponent [16]bool

	// Available animation modes
	HasAnimationMode [16]bool

	// Available modes while moving aside from WL and RN
	A1mv bool
	A2mv bool
	SCmv bool
	S1mv bool
	S2mv bool
	S3mv bool
	S4mv bool

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

	// Which skill is used for resurrection

}

// MonStats2 stores all of the MonStats2Records
//nolint:gochecknoglobals // Current design issue
var MonStats2 map[string]*MonStats2Record

// LoadMonStats2 loads MonStats2Records from monstats2.txt
//nolint:funlen //just a big data loader
func LoadMonStats2(file []byte) {
	MonStats2 = make(map[string]*MonStats2Record)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &MonStats2Record{
			Key:             d.String("Id"),
			Height:          d.Number("Height"),
			OverlayHeight:   d.Number("OverlayHeight"),
			PixelHeight:     d.Number("pixHeight"),
			SizeX:           d.Number("SizeX"),
			SizeY:           d.Number("SizeY"),
			SpawnMethod:     d.Number("spawnCol"),
			MeleeRng:        d.Number("MeleeRng"),
			BaseWeaponClass: d.String("BaseW"),
			HitClass:        d.Number("HitClass"),
			EquipmentOptions: [16][]string{
				d.List("HDv"),
				d.List("TRv"),
				d.List("LGv"),
				d.List("Rav"),
				d.List("Lav"),
				d.List("RHv"),
				d.List("LHv"),
				d.List("SHv"),
				d.List("S1v"),
				d.List("S2v"),
				d.List("S3v"),
				d.List("S4v"),
				d.List("S5v"),
				d.List("S6v"),
				d.List("S7v"),
				d.List("S8v"),
			},
			HasComponent: [16]bool{
				d.Bool("HD"),
				d.Bool("TR"),
				d.Bool("LG"),
				d.Bool("RA"),
				d.Bool("LA"),
				d.Bool("RH"),
				d.Bool("LH"),
				d.Bool("SH"),
				d.Bool("S1"),
				d.Bool("S2"),
				d.Bool("S3"),
				d.Bool("S4"),
				d.Bool("S5"),
				d.Bool("S6"),
				d.Bool("S7"),
				d.Bool("S8"),
			},
			TotalPieces: d.Number("TotalPieces"),
			HasAnimationMode: [16]bool{
				d.Bool("mDT"),
				d.Bool("mNU"),
				d.Bool("mWL"),
				d.Bool("mGH"),
				d.Bool("mA1"),
				d.Bool("mA2"),
				d.Bool("mBL"),
				d.Bool("mSC"),
				d.Bool("mS1"),
				d.Bool("mS2"),
				d.Bool("mS3"),
				d.Bool("mS4"),
				d.Bool("mDD"),
				d.Bool("mKB"),
				d.Bool("mSQ"),
				d.Bool("mRN"),
			},
			DirectionsPerMode: [16]int{
				d.Number("dDT"),
				d.Number("dNU"),
				d.Number("dWL"),
				d.Number("dGH"),
				d.Number("dA1"),
				d.Number("dA2"),
				d.Number("dBL"),
				d.Number("dSC"),
				d.Number("dS1"),
				d.Number("dS2"),
				d.Number("dS3"),
				d.Number("dS4"),
				d.Number("dDD"),
				d.Number("dKB"),
				d.Number("dSQ"),
				d.Number("dRN"),
			},
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
			ResurrectMode:      monsterAnimationModeFromString(d.String("ResurrectMode")),
			ResurrectSkill:     d.String("ResurrectSkill"),
		}
		MonStats2[record.Key] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonStats2 records", len(MonStats2))
}

//nolint:gochecknoglobals // better for lookup
var monsterAnimationModeLookup = map[string]d2enum.MonsterAnimationMode{
	d2enum.MonsterAnimationModeNeutral.String():  d2enum.MonsterAnimationModeNeutral,
	d2enum.MonsterAnimationModeSkill1.String():   d2enum.MonsterAnimationModeSkill1,
	d2enum.MonsterAnimationModeSequence.String(): d2enum.MonsterAnimationModeSequence,
}

func monsterAnimationModeFromString(s string) d2enum.MonsterAnimationMode {
	v, ok := monsterAnimationModeLookup[s]
	if !ok {
		log.Fatalf("unhandled MonsterAnimationMode %q", s)
		return d2enum.MonsterAnimationModeNeutral
	}

	return v
}
