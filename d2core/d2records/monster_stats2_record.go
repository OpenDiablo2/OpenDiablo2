package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// MonStats2 stores all of the MonStats2Records
type MonStats2 map[string]*MonStats2Record

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
