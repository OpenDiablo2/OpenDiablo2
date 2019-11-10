package datadict

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	dh "github.com/OpenDiablo2/OpenDiablo2/d2helper"
)

type ArmorRecord struct {
	Name string

	Version     int  // 0 = classic, 100 = expansion
	CompactSave bool // if true, doesn't store any stats upon saving
	Rarity      int  // higher, the rarer
	Spawnable   bool // if 0, cannot spawn in shops

	MinAC            int
	MaxAC            int
	Absorbs          int // unused?
	Speed            int // affects movement speed of wielder, >0 = you move slower, <0 = you move faster
	RequiredStrength int
	Block            int  // chance to block, capped at 75%
	Durability       int  // base durability 0-255
	NoDurability     bool // if true, item has no durability

	Level         int    // base item level (controls monster drops, for instance a lv20 monster cannot drop a lv30 item)
	RequiredLevel int    // required level to wield
	Cost          int    // base cost
	GambleCost    int    // for reference only, not used
	Code          string // identifies the item
	NameString    string // seems to be identical to code?
	MagicLevel    int    // additional magic level (for gambling?)
	AutoPrefix    int    // prefix automatically assigned to this item on spawn, maps to group column of Automagic.txt

	AlternateGfx string // code of the DCC used when equipped
	OpenBetaGfx  string // unknown
	NormalCode   string
	UberCode     string
	UltraCode    string

	SpellOffset int // unknown

	Component       int // corresponds to Composit.txt, player animation layer used by this
	InventoryWidth  int
	InventoryHeight int
	HasInventory    bool // if true, item can store gems or runes
	GemSockets      int  // number of gems to store
	GemApplyType    int  // what kind of gem effect is applied
	// 0 = weapon, 1= armor or helmet, 2 = shield

	FlippyFile          string // DC6 file animation to play when item drops on the ground
	InventoryFile       string // DC6 file used in your inventory
	UniqueInventoryFile string // DC6 file used by the unique version of this item
	SetInventoryFile    string // DC6 file used by the set version of this item

	// these represent how player animations and graphics change upon wearing this
	// these come from ArmType.txt
	AnimRightArm         int
	AnimLeftArm          int
	AnimTorso            int
	AnimLegs             int
	AnimRightShoulderPad int
	AnimLeftShoulderPad  int

	Useable bool // can be used via right click if true
	// game knows what to do if used by item code
	Throwable bool
	Stackable bool // can be stacked in inventory
	MinStack  int  // min size of stack when item is spawned, used if stackable
	MaxStack  int  // max size of stack when item is spawned

	Type  string // base type in ItemTypes.txt
	Type2 string

	DropSound    string // sfx for dropping
	DropSfxFrame int    // what frame of drop animation the sfx triggers on
	UseSound     string // sfx for using

	Unique      bool // if true, only spawns as unique
	Transparent bool // unused
	TransTable  int  // unknown, related to blending mode?
	Quivered    bool // if true, requires ammo to use
	LightRadius int  // apparently unused
	Belt        bool // tells what kind of belt this item is

	Quest int // indicates that this item belongs to a given quest?

	MissileType       int // missile gfx for throwing
	DurabilityWarning int // controls what warning icon appears when durability is low
	QuantityWarning   int // controls at what quantity the low quantity warning appears

	MinDamage      int
	MaxDamage      int
	StrengthBonus  int
	DexterityBonus int
	// final mindam = min * str / strbonus + min * dex / dexbonus
	// same for maxdam

	GemOffset int // unknown
	BitField1 int // 1 = leather item, 3 = metal

	Vendors map[string]*ItemVendorParams // controls vendor settings

	SourceArt               string // unused?
	GameArt                 string // unused?
	ColorTransform          int    // colormap to use for player's gfx
	InventoryColorTransform int    // colormap to use for inventory's gfx

	SkipName         bool   // if true, don't include the base name in the item description
	NightmareUpgrade string // upgraded in higher difficulties
	HellUpgrade      string

	UnusedMinDamage int
	UnusedMaxDamage int

	Nameable bool // if true, item can be personalized
}

type ArmorVendorParams struct {
	Min        int // minimum of this item they can stock
	Max        int // max they can stock
	MagicMin   int
	MagicMax   int
	MagicLevel uint8
}

func createArmorRecord(line string) ArmorRecord {
	r := strings.Split(line, "\t")
	i := -1
	inc := func() int {
		i++
		return i
	}
	result := ArmorRecord{
		Name: r[inc()],

		Version:     dh.StringToInt(dh.EmptyToZero(r[inc()])),
		CompactSave: dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,
		Rarity:      dh.StringToInt(dh.EmptyToZero(r[inc()])),
		Spawnable:   dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,

		MinAC:            dh.StringToInt(dh.EmptyToZero(r[inc()])),
		MaxAC:            dh.StringToInt(dh.EmptyToZero(r[inc()])),
		Absorbs:          dh.StringToInt(dh.EmptyToZero(r[inc()])),
		Speed:            dh.StringToInt(dh.EmptyToZero(r[inc()])),
		RequiredStrength: dh.StringToInt(dh.EmptyToZero(r[inc()])),
		Block:            dh.StringToInt(dh.EmptyToZero(r[inc()])),
		Durability:       dh.StringToInt(dh.EmptyToZero(r[inc()])),
		NoDurability:     dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,

		Level:         dh.StringToInt(dh.EmptyToZero(r[inc()])),
		RequiredLevel: dh.StringToInt(dh.EmptyToZero(r[inc()])),
		Cost:          dh.StringToInt(dh.EmptyToZero(r[inc()])),
		GambleCost:    dh.StringToInt(dh.EmptyToZero(r[inc()])),
		Code:          r[inc()],
		NameString:    r[inc()],
		MagicLevel:    dh.StringToInt(dh.EmptyToZero(r[inc()])),
		AutoPrefix:    dh.StringToInt(dh.EmptyToZero(r[inc()])),

		AlternateGfx: r[inc()],
		OpenBetaGfx:  r[inc()],
		NormalCode:   r[inc()],
		UberCode:     r[inc()],
		UltraCode:    r[inc()],

		SpellOffset: dh.StringToInt(dh.EmptyToZero(r[inc()])),

		Component:       dh.StringToInt(dh.EmptyToZero(r[inc()])),
		InventoryWidth:  dh.StringToInt(dh.EmptyToZero(r[inc()])),
		InventoryHeight: dh.StringToInt(dh.EmptyToZero(r[inc()])),
		HasInventory:    dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,
		GemSockets:      dh.StringToInt(dh.EmptyToZero(r[inc()])),
		GemApplyType:    dh.StringToInt(dh.EmptyToZero(r[inc()])),

		FlippyFile:          r[inc()],
		InventoryFile:       r[inc()],
		UniqueInventoryFile: r[inc()],
		SetInventoryFile:    r[inc()],

		AnimRightArm:         dh.StringToInt(dh.EmptyToZero(r[inc()])),
		AnimLeftArm:          dh.StringToInt(dh.EmptyToZero(r[inc()])),
		AnimTorso:            dh.StringToInt(dh.EmptyToZero(r[inc()])),
		AnimLegs:             dh.StringToInt(dh.EmptyToZero(r[inc()])),
		AnimRightShoulderPad: dh.StringToInt(dh.EmptyToZero(r[inc()])),
		AnimLeftShoulderPad:  dh.StringToInt(dh.EmptyToZero(r[inc()])),

		Useable: dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,

		Throwable: dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,
		Stackable: dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,
		MinStack:  dh.StringToInt(dh.EmptyToZero(r[inc()])),
		MaxStack:  dh.StringToInt(dh.EmptyToZero(r[inc()])),

		Type:  r[inc()],
		Type2: r[inc()],

		DropSound:    r[inc()],
		DropSfxFrame: dh.StringToInt(dh.EmptyToZero(r[inc()])),
		UseSound:     r[inc()],

		Unique:      dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,
		Transparent: dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,
		TransTable:  dh.StringToInt(dh.EmptyToZero(r[inc()])),
		Quivered:    dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,
		LightRadius: dh.StringToInt(dh.EmptyToZero(r[inc()])),
		Belt:        dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,

		Quest: dh.StringToInt(dh.EmptyToZero(r[inc()])),

		MissileType:       dh.StringToInt(dh.EmptyToZero(r[inc()])),
		DurabilityWarning: dh.StringToInt(dh.EmptyToZero(r[inc()])),
		QuantityWarning:   dh.StringToInt(dh.EmptyToZero(r[inc()])),

		MinDamage:      dh.StringToInt(dh.EmptyToZero(r[inc()])),
		MaxDamage:      dh.StringToInt(dh.EmptyToZero(r[inc()])),
		StrengthBonus:  dh.StringToInt(dh.EmptyToZero(r[inc()])),
		DexterityBonus: dh.StringToInt(dh.EmptyToZero(r[inc()])),

		GemOffset: dh.StringToInt(dh.EmptyToZero(r[inc()])),
		BitField1: dh.StringToInt(dh.EmptyToZero(r[inc()])),

		Vendors: createArmorVendorParams(&r, inc),

		SourceArt:               r[inc()],
		GameArt:                 r[inc()],
		ColorTransform:          dh.StringToInt(dh.EmptyToZero(r[inc()])),
		InventoryColorTransform: dh.StringToInt(dh.EmptyToZero(r[inc()])),

		SkipName:         dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,
		NightmareUpgrade: r[inc()],
		HellUpgrade:      r[inc()],

		UnusedMinDamage: dh.StringToInt(dh.EmptyToZero(r[inc()])),
		UnusedMaxDamage: dh.StringToInt(dh.EmptyToZero(r[inc()])),

		Nameable: dh.StringToInt(dh.EmptyToZero(r[inc()])) == 1,
	}
	return result
}

func createArmorVendorParams(r *[]string, inc func() int) map[string]*ItemVendorParams {
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
	vs[15] = "Malah"
	vs[16] = "Drehya"

	return CreateItemVendorParams(r, inc, vs)
}

var Armors map[string]*ArmorRecord

func LoadArmors(fileProvider d2interface.FileProvider) {
	Armors = make(map[string]*ArmorRecord)
	data := strings.Split(string(fileProvider.LoadFile(d2common.Armor)), "\r\n")[1:]
	for _, line := range data {
		if len(line) == 0 {
			continue
		}
		rec := createArmorRecord(line)
		Armors[rec.Code] = &rec
	}
	log.Printf("Loaded %d armors", len(Armors))
}
