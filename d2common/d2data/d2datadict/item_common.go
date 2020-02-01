package d2datadict

import (
	"strconv"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type ItemCommonRecord struct {
	Source d2enum.InventoryItemType

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

	Nameable bool // if true, item can be personalized

	// weapon params
	BarbOneOrTwoHanded bool // if true, barb can wield this in one or two hands
	UsesTwoHands       bool // if true, it's a 2handed weapon
	Min2HandDamage     int
	Max2HandDamage     int
	MinMissileDamage   int // ranged damage stats
	MaxMissileDamage   int
	MissileSpeed       int // unknown, affects movement speed of wielder during ranged attacks?
	ExtraRange         int // base range = 1, if this is non-zero add this to the range
	// final mindam = min * str / strbonus + min * dex / dexbonus
	// same for maxdam
	RequiredDexterity int

	WeaponClass      string // what kind of attack does this weapon have (i.e. determines attack animations)
	WeaponClass2Hand string // what kind of attack when wielded with two hands
	HitClass         string // determines sounds/graphic effects when attacking
	SpawnStack       int    // unknown, something to do with stack size when spawned (sold maybe?)

	SpecialFeature string // Just a comment

	QuestDifficultyCheck bool // if true, item only works in the difficulty it was found in

	PermStoreItem bool // if true, vendor will always sell this

	// misc params
	FlavorText string // unknown, probably just for reference

	Transmogrify bool   // if true, can be turned into another item via right click
	TransmogCode string // the 3 char code representing the item this becomes via transmog
	TransmogMin  int    // min amount of the transmog item to create
	TransmogMax  int    // max ''

	AutoBelt bool // if true, item is put into your belt when picked up

	SpellIcon         int              // which icon to display when used? Is this always -1?
	SpellType         int              // determines what kind of function is used when you use this item
	OverlayState      string           // name of the overlay state to be applied upon use of this item
	CureOverlayStates [2]string        // name of the overlay states that are removed upon use of this item
	EffectLength      int              // timer for timed usage effects
	UsageStats        [3]ItemUsageStat // stat boosts applied upon usage

	SpellDescriptionType int // specifies how to format the usage description
	// 0 = none, 1 = use desc string, 2 = use desc string + calc value
	SpellDescriptionString string              // points to a string containing the description
	SpellDescriptionCalc   d2common.CalcString // a calc string what value to display

	BetterGem string // 3 char code pointing to the gem this upgrades to (non if not applicable)

	Multibuy bool // if true, when you buy via right click + shift it will fill your belt automatically
}

type ItemUsageStat struct {
	Stat string              // name of the stat to add to
	Calc d2common.CalcString // calc string representing the amount to add
}

type ItemVendorParams struct {
	Min        int // minimum of this item they can stock
	Max        int // max they can stock
	MagicMin   int
	MagicMax   int
	MagicLevel uint8
}

// Loading Functions
var CommonItems map[string]*ItemCommonRecord

func LoadCommonItems(file []byte, source d2enum.InventoryItemType) *map[string]*ItemCommonRecord {
	if CommonItems == nil {
		CommonItems = make(map[string]*ItemCommonRecord)
	}
	items := make(map[string]*ItemCommonRecord)
	data := strings.Split(string(file), "\r\n")
	mapping := MapHeaders(data[0])
	for lineno, line := range data {
		if lineno == 0 {
			continue
		}
		if len(line) == 0 {
			continue
		}
		rec := createCommonItemRecord(line, &mapping, source)
		items[rec.Code] = &rec
		CommonItems[rec.Code] = &rec
	}
	return &items
}

func createCommonItemRecord(line string, mapping *map[string]int, source d2enum.InventoryItemType) ItemCommonRecord {
	r := strings.Split(line, "\t")
	result := ItemCommonRecord{
		Source: source,

		Name: MapLoadString(&r, mapping, "name"),

		Version:     MapLoadInt(&r, mapping, "version"),
		CompactSave: MapLoadBool(&r, mapping, "compactsave"),
		Rarity:      MapLoadInt(&r, mapping, "rarity"),
		Spawnable:   MapLoadBool(&r, mapping, "spawnable"),

		MinAC:            MapLoadInt(&r, mapping, "minac"),
		MaxAC:            MapLoadInt(&r, mapping, "maxac"),
		Absorbs:          MapLoadInt(&r, mapping, "absorbs"),
		Speed:            MapLoadInt(&r, mapping, "speed"),
		RequiredStrength: MapLoadInt(&r, mapping, "reqstr"),
		Block:            MapLoadInt(&r, mapping, "block"),
		Durability:       MapLoadInt(&r, mapping, "durability"),
		NoDurability:     MapLoadBool(&r, mapping, "nodurability"),

		Level:         MapLoadInt(&r, mapping, "level"),
		RequiredLevel: MapLoadInt(&r, mapping, "levelreq"),
		Cost:          MapLoadInt(&r, mapping, "cost"),
		GambleCost:    MapLoadInt(&r, mapping, "gamble cost"),
		Code:          MapLoadString(&r, mapping, "code"),
		NameString:    MapLoadString(&r, mapping, "namestr"),
		MagicLevel:    MapLoadInt(&r, mapping, "magic lvl"),
		AutoPrefix:    MapLoadInt(&r, mapping, "auto prefix"),

		AlternateGfx: MapLoadString(&r, mapping, "alternategfx"),
		OpenBetaGfx:  MapLoadString(&r, mapping, "OpenBetaGfx"),
		NormalCode:   MapLoadString(&r, mapping, "normcode"),
		UberCode:     MapLoadString(&r, mapping, "ubercode"),
		UltraCode:    MapLoadString(&r, mapping, "ultracode"),

		SpellOffset: MapLoadInt(&r, mapping, "spelloffset"),

		Component:       MapLoadInt(&r, mapping, "component"),
		InventoryWidth:  MapLoadInt(&r, mapping, "invwidth"),
		InventoryHeight: MapLoadInt(&r, mapping, "invheight"),
		HasInventory:    MapLoadBool(&r, mapping, "hasinv"),
		GemSockets:      MapLoadInt(&r, mapping, "gemsockets"),
		GemApplyType:    MapLoadInt(&r, mapping, "gemapplytype"),

		FlippyFile:          MapLoadString(&r, mapping, "flippyfile"),
		InventoryFile:       MapLoadString(&r, mapping, "invfile"),
		UniqueInventoryFile: MapLoadString(&r, mapping, "uniqueinvfile"),
		SetInventoryFile:    MapLoadString(&r, mapping, "setinvfile"),

		AnimRightArm:         MapLoadInt(&r, mapping, "rArm"),
		AnimLeftArm:          MapLoadInt(&r, mapping, "lArm"),
		AnimTorso:            MapLoadInt(&r, mapping, "Torso"),
		AnimLegs:             MapLoadInt(&r, mapping, "Legs"),
		AnimRightShoulderPad: MapLoadInt(&r, mapping, "rSPad"),
		AnimLeftShoulderPad:  MapLoadInt(&r, mapping, "lSPad"),

		Useable: MapLoadBool(&r, mapping, "useable"),

		Throwable: MapLoadBool(&r, mapping, "throwable"),
		Stackable: MapLoadBool(&r, mapping, "stackable"),
		MinStack:  MapLoadInt(&r, mapping, "minstack"),
		MaxStack:  MapLoadInt(&r, mapping, "maxstack"),

		Type:  MapLoadString(&r, mapping, "type"),
		Type2: MapLoadString(&r, mapping, "type2"),

		DropSound:    MapLoadString(&r, mapping, "dropsound"),
		DropSfxFrame: MapLoadInt(&r, mapping, "dropsfxframe"),
		UseSound:     MapLoadString(&r, mapping, "usesound"),

		Unique:      MapLoadBool(&r, mapping, "unique"),
		Transparent: MapLoadBool(&r, mapping, "transparent"),
		TransTable:  MapLoadInt(&r, mapping, "transtbl"),
		Quivered:    MapLoadBool(&r, mapping, "quivered"),
		LightRadius: MapLoadInt(&r, mapping, "lightradius"),
		Belt:        MapLoadBool(&r, mapping, "belt"),

		Quest: MapLoadInt(&r, mapping, "quest"),

		MissileType:       MapLoadInt(&r, mapping, "missiletype"),
		DurabilityWarning: MapLoadInt(&r, mapping, "durwarning"),
		QuantityWarning:   MapLoadInt(&r, mapping, "qntwarning"),

		MinDamage:      MapLoadInt(&r, mapping, "mindam"),
		MaxDamage:      MapLoadInt(&r, mapping, "maxdam"),
		StrengthBonus:  MapLoadInt(&r, mapping, "StrBonus"),
		DexterityBonus: MapLoadInt(&r, mapping, "DexBonus"),

		GemOffset: MapLoadInt(&r, mapping, "gemoffset"),
		BitField1: MapLoadInt(&r, mapping, "bitfield1"),

		Vendors: createItemVendorParams(&r, mapping),

		SourceArt:               MapLoadString(&r, mapping, "Source Art"),
		GameArt:                 MapLoadString(&r, mapping, "Game Art"),
		ColorTransform:          MapLoadInt(&r, mapping, "Transform"),
		InventoryColorTransform: MapLoadInt(&r, mapping, "InvTrans"),

		SkipName:         MapLoadBool(&r, mapping, "SkipName"),
		NightmareUpgrade: MapLoadString(&r, mapping, "NightmareUpgrade"),
		HellUpgrade:      MapLoadString(&r, mapping, "HellUpgrade"),

		Nameable: MapLoadBool(&r, mapping, "Nameable"),

		// weapon params
		BarbOneOrTwoHanded: MapLoadBool(&r, mapping, "1or2handed"),
		UsesTwoHands:       MapLoadBool(&r, mapping, "2handed"),
		Min2HandDamage:     MapLoadInt(&r, mapping, "2handmindam"),
		Max2HandDamage:     MapLoadInt(&r, mapping, "2handmaxdam"),
		MinMissileDamage:   MapLoadInt(&r, mapping, "minmisdam"),
		MaxMissileDamage:   MapLoadInt(&r, mapping, "maxmisdam"),
		MissileSpeed:       MapLoadInt(&r, mapping, "misspeed"),
		ExtraRange:         MapLoadInt(&r, mapping, "rangeadder"),

		RequiredDexterity: MapLoadInt(&r, mapping, "reqdex"),

		WeaponClass:      MapLoadString(&r, mapping, "wclass"),
		WeaponClass2Hand: MapLoadString(&r, mapping, "2handedwclass"),

		HitClass:   MapLoadString(&r, mapping, "hit class"),
		SpawnStack: MapLoadInt(&r, mapping, "spawnstack"),

		SpecialFeature: MapLoadString(&r, mapping, "special"),

		QuestDifficultyCheck: MapLoadBool(&r, mapping, "questdiffcheck"),

		PermStoreItem: MapLoadBool(&r, mapping, "PermStoreItem"),

		// misc params
		FlavorText: MapLoadString(&r, mapping, "szFlavorText"),

		Transmogrify: MapLoadBool(&r, mapping, "Transmogrify"),
		TransmogCode: MapLoadString(&r, mapping, "TMogType"),
		TransmogMin:  MapLoadInt(&r, mapping, "TMogMin"),
		TransmogMax:  MapLoadInt(&r, mapping, "TMogMax"),

		AutoBelt: MapLoadBool(&r, mapping, "autobelt"),

		SpellIcon:    MapLoadInt(&r, mapping, "spellicon"),
		SpellType:    MapLoadInt(&r, mapping, "pSpell"),
		OverlayState: MapLoadString(&r, mapping, "state"),
		CureOverlayStates: [2]string{
			MapLoadString(&r, mapping, "cstate1"),
			MapLoadString(&r, mapping, "cstate2"),
		},
		EffectLength: MapLoadInt(&r, mapping, "len"),
		UsageStats:   createItemUsageStats(&r, mapping),

		SpellDescriptionType: MapLoadInt(&r, mapping, "spelldesc"),
		// 0 = none, 1 = use desc string, 2 = use desc string + calc value
		SpellDescriptionString: MapLoadString(&r, mapping, "spelldescstr"),
		SpellDescriptionCalc:   d2common.CalcString(MapLoadString(&r, mapping, "spelldesccalc")),

		BetterGem: MapLoadString(&r, mapping, "BetterGem"),

		Multibuy: MapLoadBool(&r, mapping, "multibuy"),
	}
	return result
}

func createItemVendorParams(r *[]string, mapping *map[string]int) map[string]*ItemVendorParams {
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

	result := make(map[string]*ItemVendorParams)

	for _, name := range vs {
		wvp := ItemVendorParams{
			Min:        MapLoadInt(r, mapping, name+"Min"),
			Max:        MapLoadInt(r, mapping, name+"Max"),
			MagicMin:   MapLoadInt(r, mapping, name+"MagicMin"),
			MagicMax:   MapLoadInt(r, mapping, name+"MagicMax"),
			MagicLevel: MapLoadUint8(r, mapping, name+"MagicLvl"),
		}
		result[name] = &wvp
	}
	return result
}

func createItemUsageStats(r *[]string, mapping *map[string]int) [3]ItemUsageStat {
	result := [3]ItemUsageStat{}
	for i := 0; i < 3; i++ {
		result[i].Stat = MapLoadString(r, mapping, "stat"+strconv.Itoa(i))
		result[i].Calc = d2common.CalcString(MapLoadString(r, mapping, "calc"+strconv.Itoa(i)))
	}
	return result
}
