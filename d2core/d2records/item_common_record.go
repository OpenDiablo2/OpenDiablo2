package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// CommonItems stores all ItemCommonRecords
type CommonItems map[string]*ItemCommonRecord

// ItemCommonRecord is a representation of entries from armor.txt, weapons.txt, and misc.txt
type ItemCommonRecord struct {
	UsageStats             [3]ItemUsageStat         // stat boosts applied upon usage
	CureOverlayStates      [2]string                // name of the overlay states that are removed upon use of this item
	OverlayState           string                   // name of the overlay state to be applied upon use of this item
	SpellDescriptionString string                   // points to a string containing the description
	BetterGem              string                   // 3 char code pointing to the gem this upgrades to (non if not applicable)
	SpellDescriptionCalc   d2calculation.CalcString // a calc string what value to display
	WeaponClass            string                   // what kind of attack does this weapon have (i.e. determines attack animations)
	WeaponClass2Hand       string                   // what kind of attack when wielded with two hands
	HitClass               string                   // determines sounds/graphic effects when attacking
	SpecialFeature         string                   // Just a comment
	FlavorText             string                   // unknown, probably just for reference
	TransmogCode           string                   // the 3 char code representing the item this becomes via transmog
	NightmareUpgrade       string                   // upgraded in higher difficulties
	HellUpgrade            string
	SourceArt              string                       // unused?
	GameArt                string                       // unused?
	Vendors                map[string]*ItemVendorParams // controls vendor settings
	Type                   string                       // base type in ItemTypes.txt
	Type2                  string
	DropSound              string // sfx for dropping
	UseSound               string // sfx for using
	FlippyFile             string // DC6 file animation to play when item drops on the ground
	InventoryFile          string // DC6 file used in your inventory
	UniqueInventoryFile    string // DC6 file used by the unique version of this item
	SetInventoryFile       string // DC6 file used by the set version of this item
	Code                   string // identifies the item
	NameString             string // seems to be identical to code?
	AlternateGfx           string // code of the DCC used when equipped
	OpenBetaGfx            string // unknown
	NormalCode             string
	UberCode               string
	UltraCode              string
	Name                   string
	Source                 d2enum.InventoryItemType

	Version          int // 0 = classic, 100 = expansion
	Rarity           int // higher, the rarer
	MinAC            int
	MaxAC            int
	Absorbs          int // unused?
	Speed            int // affects movement speed of wielder, >0 = you move slower, <0 = you move faster
	RequiredStrength int
	Block            int // chance to block, capped at 75%
	Durability       int // base durability 0-255
	Level            int // base item level (controls monster drops, for instance a lv20 monster cannot drop a lv30 item)
	RequiredLevel    int // required level to wield
	Cost             int // base cost
	GambleCost       int // for reference only, not used
	MagicLevel       int // additional magic level (for gambling?)
	AutoPrefix       int // prefix automatically assigned to this item on spawn, maps to group column of Automagic.txt
	SpellOffset      int // unknown
	Component        int // corresponds to Composit.txt, player animation layer used by this
	InventoryWidth   int
	InventoryHeight  int
	GemSockets       int // number of gems to store
	GemApplyType     int // what kind of gem effect is applied
	// 0 = weapon, 1= armor or helmet, 2 = shield

	// these represent how player animations and graphics change upon wearing this
	// these come from ArmType.txt
	AnimRightArm         int
	AnimLeftArm          int
	AnimTorso            int
	AnimLegs             int
	AnimRightShoulderPad int
	AnimLeftShoulderPad  int

	MinStack          int // min size of stack when item is spawned, used if stackable
	MaxStack          int // max size of stack when item is spawned
	DropSfxFrame      int // what frame of drop animation the sfx triggers on
	TransTable        int // unknown, related to blending mode?
	LightRadius       int // apparently unused
	Quest             int // indicates that this item belongs to a given quest?
	MissileType       int // missile gfx for throwing
	DurabilityWarning int // controls what warning icon appears when durability is low
	QuantityWarning   int // controls at what quantity the low quantity warning appears
	MinDamage         int
	MaxDamage         int
	StrengthBonus     int
	DexterityBonus    int
	// final mindam = min * str / strbonus + min * dex / dexbonus
	// same for maxdam

	GemOffset               int // unknown
	BitField1               int // 1 = leather item, 3 = metal
	ColorTransform          int // colormap to use for player's gfx
	InventoryColorTransform int // colormap to use for inventory's gfx
	Min2HandDamage          int
	Max2HandDamage          int
	MinMissileDamage        int // ranged damage stats
	MaxMissileDamage        int
	MissileSpeed            int // unknown, affects movement speed of wielder during ranged attacks?
	ExtraRange              int // base range = 1, if this is non-zero add this to the range
	// final mindam = min * str / strbonus + min * dex / dexbonus
	// same for maxdam
	RequiredDexterity    int
	SpawnStack           int // unknown, something to do with stack size when spawned (sold maybe?)
	TransmogMin          int // min amount of the transmog item to create
	TransmogMax          int // max ''
	SpellIcon            int // which icon to display when used? Is this always -1?
	SpellType            int // determines what kind of function is used when you use this item
	EffectLength         int // timer for timed usage effects
	SpellDescriptionType int // specifies how to format the usage description
	// 0 = none, 1 = use desc string, 2 = use desc string + calc value

	AutoBelt     bool // if true, item is put into your belt when picked up
	HasInventory bool // if true, item can store gems or runes
	CompactSave  bool // if true, doesn't store any stats upon saving
	Spawnable    bool // if 0, cannot spawn in shops
	NoDurability bool // if true, item has no durability
	Useable      bool // can be used via right click if true
	// game knows what to do if used by item code
	Throwable            bool
	Stackable            bool // can be stacked in inventory
	Unique               bool // if true, only spawns as unique
	Transparent          bool // unused
	Quivered             bool // if true, requires ammo to use
	Belt                 bool // tells what kind of belt this item is
	SkipName             bool // if true, don't include the base name in the item description
	Nameable             bool // if true, item can be personalized
	BarbOneOrTwoHanded   bool // if true, barb can wield this in one or two hands
	UsesTwoHands         bool // if true, it's a 2handed weapon
	QuestDifficultyCheck bool // if true, item only works in the difficulty it was found in
	PermStoreItem        bool // if true, vendor will always sell this
	Transmogrify         bool // if true, can be turned into another item via right click
	Multibuy             bool // if true, when you buy via right click + shift it will fill your belt automatically
}

// ItemUsageStat the stat that gets applied when the item is used
type ItemUsageStat struct {
	Stat string                   // name of the stat to add to
	Calc d2calculation.CalcString // calc string representing the amount to add
}

// ItemVendorParams are parameters that vendors use
type ItemVendorParams struct {
	Min        int // minimum of this item they can stock
	Max        int // max they can stock
	MagicMin   int
	MagicMax   int
	MagicLevel int
}
