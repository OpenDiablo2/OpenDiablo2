package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

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

	DescFn      interface{} // the sprintf func
	DescGroupFn interface{} // group sprintf func

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
	DescFnID     d2enum.DescFuncID

	// Controls whenever and if so in what way the stat value is shown
	// 0 = doesn't show the value of the stat
	// 1 = shows the value of the stat infront of the description
	// 2 = shows the value of the stat after the description.
	DescVal int

	// when stats in the same group have the same value they use the
	// group func for desc (they need to be in the same affix)
	DescGroup       int
	DescGroupVal    int
	DescGroupFuncID d2enum.DescFuncID

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
//nolint:gochecknoglobals // Currently global by design
var ItemStatCosts map[string]*ItemStatCostRecord

// LoadItemStatCosts loads ItemStatCostRecord's from text
func LoadItemStatCosts(file []byte) {
	d := d2common.LoadDataDictionary(string(file))
	numRecords := len(d.Data)
	ItemStatCosts = make(map[string]*ItemStatCostRecord, numRecords)

	for idx := range d.Data {
		record := &ItemStatCostRecord{
			Name:  d.GetString("Stat", idx),
			Index: d.GetNumber("ID", idx),

			Signed:   d.GetNumber("Signed", idx) > 0,
			KeepZero: d.GetNumber("keepzero", idx) > 0,

			// Ranged:  d.GetNumber("Ranged", idx) > 0,
			MinAccr: d.GetNumber("MinAccr", idx),

			UpdateAnimRate: d.GetNumber("UpdateAnimRate", idx) > 0,

			SendOther: d.GetNumber("Send Other", idx) > 0,
			SendBits:  d.GetNumber("Send Bits", idx),
			SendParam: d.GetNumber("Send Param Bits", idx),

			Saved:         d.GetNumber("CSvBits", idx) > 0,
			SavedSigned:   d.GetNumber("CSvSigned", idx) > 0,
			SavedBits:     d.GetNumber("CSvBits", idx),
			SaveBits:      d.GetNumber("Save Bits", idx),
			SaveAdd:       d.GetNumber("Save Add", idx),
			SaveParamBits: d.GetNumber("Save Param Bits", idx),

			Encode: d2enum.EncodingType(d.GetNumber("Encode", idx)),

			CallbackEnabled: d.GetNumber("fCallback", idx) > 0,

			CostAdd:      d.GetNumber("Add", idx),
			CostMultiply: d.GetNumber("Multiply", idx),
			ValShift:     d.GetNumber("ValShift", idx),

			OperatorType: d2enum.OperatorType(d.GetNumber("op", idx)),
			OpParam:      d.GetNumber("op param", idx),
			OpBase:       d.GetString("op base", idx),
			OpStat1:      d.GetString("op stat1", idx),
			OpStat2:      d.GetString("op stat2", idx),
			OpStat3:      d.GetString("op stat3", idx),

			Direct:  d.GetNumber("direct", idx) > 0,
			MaxStat: d.GetString("maxstat", idx),

			ItemSpecific:  d.GetNumber("itemspecific", idx) > 0,
			DamageRelated: d.GetNumber("damagerelated", idx) > 0,

			EventID1:     d2enum.GetItemEventType(d.GetString("itemevent1", idx)),
			EventID2:     d2enum.GetItemEventType(d.GetString("itemevent2", idx)),
			EventFuncID1: d2enum.GetItemEventFuncID(d.GetNumber("itemeventfunc1", idx)),
			EventFuncID2: d2enum.GetItemEventFuncID(d.GetNumber("itemeventfunc2", idx)),

			DescPriority: d.GetNumber("descpriority", idx),
			DescFnID:     d2enum.DescFuncID(d.GetNumber("descfunc", idx)),
			DescFn:       d2enum.GetDescFunction(d2enum.DescFuncID(d.GetNumber("descfunc", idx))),
			DescVal:      d.GetNumber("descval", idx),
			DescStrPos:   d.GetString("descstrpos", idx),
			DescStrNeg:   d.GetString("descstrneg", idx),
			DescStr2:     d.GetString("descstr2", idx),

			DescGroup:       d.GetNumber("dgrp", idx),
			DescGroupFuncID: d2enum.DescFuncID(d.GetNumber("dgrpfunc", idx)),
			DescGroupFn:     d2enum.GetDescFunction(d2enum.DescFuncID(d.GetNumber("dgrpfunc", idx))),
			DescGroupVal:    d.GetNumber("dgrpval", idx),
			DescGroupStrPos: d.GetString("dgrpstrpos", idx),
			DescGroupStrNeg: d.GetString("dgrpstrneg", idx),
			DescGroupStr2:   d.GetString("dgrpstr2", idx),

			Stuff: d.GetString("stuff", idx),
		}

		ItemStatCosts[record.Name] = record
	}

	log.Printf("Loaded %d ItemStatCost records", len(ItemStatCosts))
}
