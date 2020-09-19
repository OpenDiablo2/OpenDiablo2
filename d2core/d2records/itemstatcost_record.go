package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// ItemStatCostRecord represents a row from itemstatcost.txt
// these records describe the stat values and costs (in shops) of items
// refer to https://d2mods.info/forum/kb/viewarticle?a=448
type ItemStatCostRecord struct {
	Name    string
	OpBase  string
	OpStat1 string
	OpStat2 string
	OpStat3 string

	MaxStat         string // if Direct true, will not exceed val of MaxStat
	DescStrPos      string // string used when val is positive
	DescStrNeg      string
	DescStr2        string // additional string used by some string funcs
	DescGroupStrPos string // string used when val is positive
	DescGroupStrNeg string
	DescGroupStr2   string // additional string used by some string funcs

	// Stuff
	// Stay far away from this column unless you really know what you're
	// doing and / or work for Blizzard, this column is used during bin-file
	// creation to generate a cache regulating the op-stat stuff and other
	// things, changing it can be futile, it works like the constants column
	// in MonUMod.txt and has no other relation to ItemStatCost.txt, the first
	// stat in the file simply must have this set or else you may break the
	// entire op stuff.
	Stuff string

	Index int

	// path_d2.mpq version doesnt have Ranged columne, excluding for now
	// Ranged  bool // game attempts to keep stat in a range, like strength >-1
	MinAccr int // minimum ranged value

	SendBits  int // #bits to send in stat update
	SendParam int // #bits to send in stat update

	SavedBits int // #bits allocated to the value in .d2s file

	SaveBits      int // #bits saved to .d2s files, max == 2^SaveBits-1
	SaveAdd       int // how large the negative range is (lowers max, as well)
	SaveParamBits int // #param bits are saved (safe value is 17)

	Encode d2enum.EncodingType // how the stat is encoded in .d2s files

	// these two fields control additional cost on items
	// cost * (1 + value * multiply / 1024)) + add (...)
	CostAdd      int
	CostMultiply int
	// CostDivide // exists in txt, but division hardcoded to 1024
	// if divide is used, could we do (?):
	// cost * (1 + value * multiply / divide)) + add (...)

	ValShift int // controls how stat is stored in .d2s
	// so that you can save `+1` instead of `+256`

	OperatorType d2enum.OperatorType
	OpParam      int

	EventID1     d2enum.ItemEventType
	EventID2     d2enum.ItemEventType
	EventFuncID1 d2enum.ItemEventFuncID
	EventFuncID2 d2enum.ItemEventFuncID

	DescPriority int // determines order when displayed
	DescFnID     int

	// Controls whenever and if so in what way the stat value is shown
	// 0 = doesn't show the value of the stat
	// 1 = shows the value of the stat infront of the description
	// 2 = shows the value of the stat after the description.
	DescVal int

	// when stats in the same group have the same value they use the
	// group func for desc (they need to be in the same affix)
	DescGroup       int
	DescGroupVal    int
	DescGroupFuncID int

	CallbackEnabled bool // whether callback fn is called if value changes
	Signed          bool // whether the stat is signed
	KeepZero        bool // prevent from going negative (assume only client side)
	UpdateAnimRate  bool // when altered, forces speed handler to adjust speed
	SendOther       bool // whether to send to other clients
	Saved           bool // whether this stat is saved in .d2s files
	SavedSigned     bool // whether the stat is saved as signed/unsigned
	Direct          bool // whether is temporary or permanent
	ItemSpecific    bool // prevents stacking with an existing stat on item
	// like when socketing a jewel

	DamageRelated bool // prevents stacking of stats while dual wielding
}

// ItemStatCosts stores all of the ItemStatCostRecords
type ItemStatCosts map[string]*ItemStatCostRecord
