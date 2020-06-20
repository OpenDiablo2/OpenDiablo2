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
	dict := d2common.LoadDataDictionary(string(file))
	numRecords := len(dict.Data)
	MonStats2 = make(map[string]*MonStats2Record, numRecords)

	for idx := range dict.Data {
		record := &MonStats2Record{
			Key:                dict.GetString("Id", idx),
			Height:             dict.GetNumber("Height", idx),
			OverlayHeight:      dict.GetNumber("OverlayHeight", idx),
			PixelHeight:        dict.GetNumber("pixHeight", idx),
			SizeX:              dict.GetNumber("SizeX", idx),
			SizeY:              dict.GetNumber("SizeY", idx),
			SpawnMethod:        dict.GetNumber("spawnCol", idx),
			MeleeRng:           dict.GetNumber("MeleeRng", idx),
			BaseWeaponClass:    dict.GetString("BaseW", idx),
			HitClass:           dict.GetNumber("HitClass", idx),
			HDv:                dict.GetDelimitedList("HDv", idx),
			TRv:                dict.GetDelimitedList("TRv", idx),
			LGv:                dict.GetDelimitedList("LGv", idx),
			Rav:                dict.GetDelimitedList("Rav", idx),
			Lav:                dict.GetDelimitedList("Lav", idx),
			RHv:                dict.GetDelimitedList("RHv", idx),
			LHv:                dict.GetDelimitedList("LHv", idx),
			SHv:                dict.GetDelimitedList("SHv", idx),
			S1v:                dict.GetDelimitedList("S1v", idx),
			S2v:                dict.GetDelimitedList("S2v", idx),
			S3v:                dict.GetDelimitedList("S3v", idx),
			S4v:                dict.GetDelimitedList("S4v", idx),
			S5v:                dict.GetDelimitedList("S5v", idx),
			S6v:                dict.GetDelimitedList("S6v", idx),
			S7v:                dict.GetDelimitedList("S7v", idx),
			S8v:                dict.GetDelimitedList("S8v", idx),
			HD:                 dict.GetBool("HD", idx),
			TR:                 dict.GetBool("TR", idx),
			LG:                 dict.GetBool("LG", idx),
			RA:                 dict.GetBool("RA", idx),
			LA:                 dict.GetBool("LA", idx),
			RH:                 dict.GetBool("RH", idx),
			LH:                 dict.GetBool("LH", idx),
			SH:                 dict.GetBool("SH", idx),
			S1:                 dict.GetBool("S1", idx),
			S2:                 dict.GetBool("S2", idx),
			S3:                 dict.GetBool("S3", idx),
			S4:                 dict.GetBool("S4", idx),
			S5:                 dict.GetBool("S5", idx),
			S6:                 dict.GetBool("S6", idx),
			S7:                 dict.GetBool("S7", idx),
			S8:                 dict.GetBool("S8", idx),
			TotalPieces:        dict.GetNumber("TotalPieces", idx),
			mDT:                dict.GetBool("mDT", idx),
			mNU:                dict.GetBool("mNU", idx),
			mWL:                dict.GetBool("mWL", idx),
			mGH:                dict.GetBool("mGH", idx),
			mA1:                dict.GetBool("mA1", idx),
			mA2:                dict.GetBool("mA2", idx),
			mBL:                dict.GetBool("mBL", idx),
			mSC:                dict.GetBool("mSC", idx),
			mS1:                dict.GetBool("mS1", idx),
			mS2:                dict.GetBool("mS2", idx),
			mS3:                dict.GetBool("mS3", idx),
			mS4:                dict.GetBool("mS4", idx),
			mDD:                dict.GetBool("mDD", idx),
			mKB:                dict.GetBool("mKB", idx),
			mSQ:                dict.GetBool("mSQ", idx),
			mRN:                dict.GetBool("mRN", idx),
			dDT:                dict.GetNumber("mDT", idx),
			dNU:                dict.GetNumber("mNU", idx),
			dWL:                dict.GetNumber("mWL", idx),
			dGH:                dict.GetNumber("mGH", idx),
			dA1:                dict.GetNumber("mA1", idx),
			dA2:                dict.GetNumber("mA2", idx),
			dBL:                dict.GetNumber("mBL", idx),
			dSC:                dict.GetNumber("mSC", idx),
			dS1:                dict.GetNumber("mS1", idx),
			dS2:                dict.GetNumber("mS2", idx),
			dS3:                dict.GetNumber("mS3", idx),
			dS4:                dict.GetNumber("mS4", idx),
			dDD:                dict.GetNumber("mDD", idx),
			dKB:                dict.GetNumber("mKB", idx),
			dSQ:                dict.GetNumber("mSQ", idx),
			dRN:                dict.GetNumber("mRN", idx),
			A1mv:               dict.GetBool("A1mv", idx),
			A2mv:               dict.GetBool("A2mv", idx),
			SCmv:               dict.GetBool("SCmv", idx),
			S1mv:               dict.GetBool("S1mv", idx),
			S2mv:               dict.GetBool("S2mv", idx),
			S3mv:               dict.GetBool("S3mv", idx),
			S4mv:               dict.GetBool("S4mv", idx),
			NoGfxHitTest:       dict.GetBool("noGfxHitTest", idx),
			BoxTop:             dict.GetNumber("htTop", idx),
			BoxLeft:            dict.GetNumber("htLeft", idx),
			BoxWidth:           dict.GetNumber("htWidth", idx),
			BoxHeight:          dict.GetNumber("htHeight", idx),
			Restore:            dict.GetNumber("restore", idx),
			AutomapCel:         dict.GetNumber("automapCel", idx),
			NoMap:              dict.GetBool("noMap", idx),
			NoOvly:             dict.GetBool("noOvly", idx),
			IsSelectable:       dict.GetBool("isSel", idx),
			AllySelectable:     dict.GetBool("alSel", idx),
			shiftSel:           dict.GetBool("shiftSel", idx),
			NotSelectable:      dict.GetBool("noSel", idx),
			IsCorpseSelectable: dict.GetBool("corpseSel", idx),
			IsAttackable:       dict.GetBool("isAtt", idx),
			IsRevivable:        dict.GetBool("revive", idx),
			IsCritter:          dict.GetBool("critter", idx),
			IsSmall:            dict.GetBool("small", idx),
			IsLarge:            dict.GetBool("large", idx),
			IsSoft:             dict.GetBool("soft", idx),
			IsInert:            dict.GetBool("inert", idx),
			objCol:             dict.GetBool("objCol", idx),
			IsCorpseCollidable: dict.GetBool("deadCol", idx),
			IsCorpseWalkable:   dict.GetBool("unflatDead", idx),
			HasShadow:          dict.GetBool("Shadow", idx),
			NoUniqueShift:      dict.GetBool("noUniqueShift", idx),
			CompositeDeath:     dict.GetBool("compositeDeath", idx),
			LocalBlood:         dict.GetNumber("localBlood", idx),
			Bleed:              dict.GetNumber("Bleed", idx),
			Light:              dict.GetNumber("Light", idx),
			LightR:             dict.GetNumber("light-r", idx),
			LightG:             dict.GetNumber("light-g", idx),
			lightB:             dict.GetNumber("light-b", idx),
			NormalPalette:      dict.GetNumber("Utrans", idx),
			NightmarePalette:   dict.GetNumber("Utrans(N)", idx),
			HellPalatte:        dict.GetNumber("Utrans(H)", idx),
			Heart:              dict.GetString("Heart", idx),
			BodyPart:           dict.GetString("BodyPart", idx),
			InfernoLen:         dict.GetNumber("InfernoLen", idx),
			InfernoAnim:        dict.GetNumber("InfernoAnim", idx),
			InfernoRollback:    dict.GetNumber("InfernoRollback", idx),
			ResurrectMode:      d2enum.MonsterAnimationModeFromString(dict.GetString("ResurrectMode", idx)),
			ResurrectSkill:     dict.GetString("ResurrectSkill", idx),
		}
		MonStats2[record.Key] = record
	}

	log.Printf("Loaded %d MonStats2 records", len(MonStats2))
}
