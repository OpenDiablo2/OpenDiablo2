package common

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/resourcepaths"
)

type WeaponRecord struct {
	Name string

	Type         string // base type in ItemTypes.txt
	Type2        string
	Code         string // identifies the item
	AlternateGfx string // code of the DCC used when equipped
	NameString   string // seems to be identical to code?
	Version      int    // 0 = classic, 100 = expansion
	CompactSave  bool   // if true, doesn't store any stats upon saving
	Rarity       int    // higher, the rarer
	Spawnable    bool   // if 0, cannot spawn in shops

	MinDamage          int
	MaxDamage          int
	BarbOneOrTwoHanded bool // if true, barb can wield this in one or two hands
	UsesTwoHands       bool // if true, it's a 2handed weapon
	Min2HandDamage     int
	Max2HandDamage     int
	MinMissileDamage   int // ranged damage stats
	MaxMissileDamage   int
	MissileSpeed       int // unknown, affects movement speed of wielder during ranged attacks?
	ExtraRange         int // base range = 1, if this is non-zero add this to the range
	Speed              int // affects movement speed of wielder, >0 = you move slower, <0 = you move faster
	StrengthBonus      int
	DexterityBonus     int
	// final mindam = min * str / strbonus + min * dex / dexbonus
	// same for maxdam
	RequiredStrength  int
	RequiredDexterity int
	Durability        int  // base durability 0-255
	NoDurability      bool // if true, item has no durability

	Level         int    // base item level (controls monster drops, for instance a lv20 monster cannot drop a lv30 item)
	RequiredLevel int    // required level to wield
	Cost          int    // base cost
	GambleCost    int    // for reference only, not used
	MagicLevel    int    // additional magic level (for gambling?)
	AutoPrefix    int    // prefix automatically assigned to this item on spawn, maps to group column of Automagic.txt
	OpenBetaGfx   string // unknown
	NormalCode    string
	UberCode      string
	UltraCode     string

	WeaponClass      string // what kind of attack does this weapon have (i.e. determines attack animations)
	WeaponClass2Hand string // what kind of attack when wielded with two hands
	Component        int    // corresponds to Composit.txt, player animation layer used by this
	HitClass         string // determines sounds/graphic effects when attacking
	InventoryWidth   int
	InventoryHeight  int
	Stackable        bool // can be stacked in inventory
	MinStack         int  // min size of stack when item is spawned, used if stackable
	MaxStack         int  // max size of stack when item is spawned
	SpawnStack       int  // unknown, something to do with stack size when spawned (sold maybe?)

	FlippyFile          string // DC6 file animation to play when item drops on the ground
	InventoryFile       string // DC6 file used in your inventory
	UniqueInventoryFile string // DC6 file used by the unique version of this item
	SetInventoryFile    string // DC6 file used by the set version of this item

	HasInventory bool // if true, item can store gems or runes
	GemSockets   int  // number of gems to store
	GemApplyType int  // what kind of gem effect is applied
	// 0 = weapon, 1= armor or helmet, 2 = shield

	SpecialFeature string // Just a comment

	Useable bool // can be used via right click if true
	// game knows what to do if used by item code
	DropSound    string // sfx for dropping
	DropSfxFrame int    // what frame of drop animation the sfx triggers on
	UseSound     string // sfx for using

	Unique      bool // if true, only spawns as unique
	Transparent bool // unused
	TransTable  int  // unknown, related to blending mode?
	Quivered    bool // if true, requires ammo to use
	LightRadius int  // apparently unused
	Belt        bool // seems to be unused? supposed to be whether this can go in your quick access belt

	Quest                int  // indicates that this item belongs to a given quest?
	QuestDifficultyCheck bool // if true, item only works in the difficulty it was found in

	MissileType       int // missile gfx for throwing
	DurabilityWarning int // controls what warning icon appears when durability is low
	QuantityWarning   int // controls at what quantity the low quantity warning appears
	GemOffset         int // unknown
	BitField1         int // 1 = leather item, 3 = metal

	Vendors map[string]*ItemVendorParams // controls vendor settings

	SourceArt               string // unused?
	GameArt                 string // unused?
	ColorTransform          int    // colormap to use for player's gfx
	InventoryColorTransform int    // colormap to use for inventory's gfx

	SkipName         bool   // if true, don't include the base name in the item description
	NightmareUpgrade string // upgraded in higher difficulties
	HellUpgrade      string

	Nameable      bool // if true, item can be personalized
	PermStoreItem bool // if true, vendor will always sell this
}

type ItemVendorParams struct {
	Min        int // minimum of this item they can stock
	Max        int // max they can stock
	MagicMin   int
	MagicMax   int
	MagicLevel uint8
}

func createWeaponRecord(line string) WeaponRecord {
	r := strings.Split(line, "\t")
	i := -1
	inc := func() int {
		i++
		return i
	}
	result := WeaponRecord{
		Name: r[inc()],

		Type:         r[inc()],
		Type2:        r[inc()],
		Code:         r[inc()],
		AlternateGfx: r[inc()],
		NameString:   r[inc()],
		Version:      StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		CompactSave:  StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,
		Rarity:       StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		Spawnable:    StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,

		MinDamage:          StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		MaxDamage:          StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		BarbOneOrTwoHanded: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,
		UsesTwoHands:       StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,
		Min2HandDamage:     StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		Max2HandDamage:     StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		MinMissileDamage:   StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		MaxMissileDamage:   StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		MissileSpeed:       StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		ExtraRange:         StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		Speed:              StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		StrengthBonus:      StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		DexterityBonus:     StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),

		RequiredStrength:  StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		RequiredDexterity: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		Durability:        StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		NoDurability:      StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,

		Level:         StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		RequiredLevel: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		Cost:          StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		GambleCost:    StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		MagicLevel:    StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		AutoPrefix:    StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		OpenBetaGfx:   r[inc()],
		NormalCode:    r[inc()],
		UberCode:      r[inc()],
		UltraCode:     r[inc()],

		WeaponClass:      r[inc()],
		WeaponClass2Hand: r[inc()],
		Component:        StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		HitClass:         r[inc()],
		InventoryWidth:   StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		InventoryHeight:  StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		Stackable:        StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,
		MinStack:         StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		MaxStack:         StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		SpawnStack:       StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),

		FlippyFile:          r[inc()],
		InventoryFile:       r[inc()],
		UniqueInventoryFile: r[inc()],
		SetInventoryFile:    r[inc()],

		HasInventory: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,
		GemSockets:   StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		GemApplyType: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),

		SpecialFeature: r[inc()],

		Useable: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,

		DropSound:    r[inc()],
		DropSfxFrame: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		UseSound:     r[inc()],

		Unique:      StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,
		Transparent: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,
		TransTable:  StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		Quivered:    StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,
		LightRadius: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		Belt:        StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,

		Quest:                StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		QuestDifficultyCheck: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,

		MissileType:       StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		DurabilityWarning: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		QuantityWarning:   StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		GemOffset:         StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		BitField1:         StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),

		Vendors: createWeaponVendorParams(&r, inc),

		SourceArt:               r[inc()],
		GameArt:                 r[inc()],
		ColorTransform:          StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),
		InventoryColorTransform: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))),

		SkipName:         StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,
		NightmareUpgrade: r[inc()],
		HellUpgrade:      r[inc()],

		Nameable:      StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,
		PermStoreItem: StringToInt(EmptyToZero(AsterToEmpty(r[inc()]))) == 1,
	}
	return result
}

func createWeaponVendorParams(r *[]string, inc func() int) map[string]*ItemVendorParams {
	vs := make([]string, 17)
	vs[0] = "Charsi"
	vs[1] = "Gheed"
	vs[2] = "Akara"
	vs[3] = "Fara"
	vs[4] = "Lysander"
	vs[5] = "Drognan"
	vs[6] = "Hralti"
	vs[7] = "Alkor"
	vs[8] = "Ormus"
	vs[9] = "Elzix"
	vs[10] = "Asheara"
	vs[11] = "Cain"
	vs[12] = "Halbu"
	vs[13] = "Jamella"
	vs[14] = "Larzuk"
	vs[15] = "Drehya"
	vs[16] = "Malah"

	return CreateItemVendorParams(r, inc, vs)
}

func CreateItemVendorParams(r *[]string, inc func() int, vs []string) map[string]*ItemVendorParams {
	result := make(map[string]*ItemVendorParams)
	
	for _, name := range vs {
		wvp := ItemVendorParams{
			Min:        StringToInt(EmptyToZero((*r)[inc()])),
			Max:        StringToInt(EmptyToZero((*r)[inc()])),
			MagicMin:   StringToInt(EmptyToZero((*r)[inc()])),
			MagicMax:   StringToInt(EmptyToZero((*r)[inc()])),
			MagicLevel: StringToUint8(EmptyToZero((*r)[inc()])),
		}
		result[name] = &wvp
	}
	return result
}

var Weapons map[string]*WeaponRecord

func LoadWeapons(fileProvider FileProvider) {
	Weapons = make(map[string]*WeaponRecord)
	data := strings.Split(string(fileProvider.LoadFile(resourcepaths.Weapons)), "\r\n")[1:]
	for _, line := range data {
		if len(line) == 0 {
			continue
		}
		rec := createWeaponRecord(line)
		Weapons[rec.Code] = &rec
	}
	log.Printf("Loaded %d weapons", len(Weapons))
}
