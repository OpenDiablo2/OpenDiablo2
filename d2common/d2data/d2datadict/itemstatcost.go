package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// refer to https://d2mods.info/forum/kb/viewarticle?a=448
type ItemStatCostRecord struct {
	Name  string
	Index int

	Signed   bool // whether the stat is signed
	KeepZero bool // prevent from going negative (assume only client side)

	// path_d2.mpq version doesnt have Ranged columne, excluding for now
	// Ranged  bool // game attempts to keep stat in a range, like strength >-1
	MinAccr int // minimum ranged value

	UpdateAnimRate bool // when altered, forces speed handler to adjust speed

	SendOther bool // whether to send to other clients
	SendBits  int  // #bits to send in stat update
	SendParam int  // #bits to send in stat update

	Saved       bool // whether this stat is saved in .d2s files
	SavedSigned bool // whether the stat is saved as signed/unsigned
	SavedBits   int  // #bits allocated to the value in .d2s file

	SaveBits      int // #bits saved to .d2s files, max == 2^SaveBits-1
	SaveAdd       int // how large the negative range is (lowers max, as well)
	SaveParamBits int // #param bits are saved (safe value is 17)

	Encode EncodingType // how the stat is encoded in .d2s files

	CallbackEnabled bool // whether callback fn is called if value changes

	// these two fields control additional cost on items
	// cost * (1 + value * multiply / 1024)) + add (...)
	CostAdd      int
	CostMultiply int
	// CostDivide // exists in txt, but division hardcoded to 1024
	// if divide is used, could we do (?):
	// cost * (1 + value * multiply / divide)) + add (...)

	ValShift int // controls how stat is stored in .d2s
	// so that you can save `+1` instead of `+256`

	OperatorType OperatorType
	OpParam      int
	OpBase       string
	OpStat1      string
	OpStat2      string
	OpStat3      string

	Direct  bool   // whether is temporary or permanent
	MaxStat string // if Direct true, will not exceed val of MaxStat

	ItemSpecific bool // prevents stacking with an existing stat on item
	// like when socketing a jewel

	DamageRelated bool // prevents stacking of stats while dual wielding

	EventID1     d2enum.ItemEventType
	EventID2     d2enum.ItemEventType
	EventFuncID1 d2enum.ItemEventFuncID
	EventFuncID2 d2enum.ItemEventFuncID

	DescPriority int // determines order when displayed
	DescFnID     d2enum.DescFuncID
	DescFn       interface{} // the sprintf func

	// Controls whenever and if so in what way the stat value is shown
	// 0 = doesn't show the value of the stat
	// 1 = shows the value of the stat infront of the description
	// 2 = shows the value of the stat after the description.
	DescVal    int
	DescStrPos string // string used when val is positive
	DescStrNeg string
	DescStr2   string // additional string used by some string funcs

	// when stats in the same group have the same value they use the
	// group func for desc (they need to be in the same affix)
	DescGroup       int
	DescGroupFuncID d2enum.DescFuncID
	DescGroupFn     interface{} // group sprintf func
	DescGroupVal    int
	DescGroupStrPos string // string used when val is positive
	DescGroupStrNeg string
	DescGroupStr2   string // additional string used by some string funcs

	// Stay far away from this column unless you really know what you're
	// doing and / or work for Blizzard, this column is used during bin-file
	// creation to generate a cache regulating the op-stat stuff and other
	// things, changing it can be futile, it works like the constants column
	// in MonUMod.txt and has no other relation to ItemStatCost.txt, the first
	// stat in the file simply must have this set or else you may break the
	// entire op stuff.
	Stuff string // ? TODO ?
}

type EncodingType int

const (
	// TODO: determine other encoding types.
	// didn't see anything about how this stuff is encoded, or the types...
	EncodeDefault = EncodingType(iota)
)

type OperatorType int // for dynamic properties

const (
	// just adds the stat to the unit directly
	OpDefault = OperatorType(iota)

	// adds opstat.base * statvalue / 100 to the opstat.
	Op1

	// adds (statvalue * basevalue) / (2 ^ param) to the opstat
	// this does not work properly with any stat other then level because of the
	// way this is updated, it is only refreshed when you re-equip the item,
	// your character is saved or you level up, similar to passive skills, just
	// because it looks like it works in the item description
	// does not mean it does, the game just recalculates the information in the
	// description every frame, while the values remain unchanged serverside.
	Op2

	// this is a percentage based version of op #2
	// look at op #2 for information about the formula behind it, just
	// remember the stat is increased by a percentage rather then by adding
	// an integer.
	Op3

	// this works the same way op #2 works, however the stat bonus is
	// added to the item and not to the player (so that +defense per level
	// properly adds the defense to the armor and not to the character
	// directly!)
	Op4

	// this works like op #4 but is percentage based, it is used for percentage
	// based increase of stats that are found on the item itself, and not stats
	// that are found on the character.
	Op5

	// like for op #7, however this adds a plain bonus to the stat, and just
	// like #7 it also doesn't work so I won't bother to explain the arithmetic
	// behind it either.
	Op6

	// this is used to increase a stat based on the current daytime of the game
	// world by a percentage, there is no need to explain the arithmetics
	// behind it because frankly enough it just doesn't work serverside, it
	// only updates clientside so this op is essentially useless.
	Op7

	// hardcoded to work only with maxmana, this will apply the proper amount
	// of mana to your character based on CharStats.txt for the amount of energy
	// the stat added (doesn't work for non characters)
	Op8

	// hardcoded to work only with maxhp and maxstamina, this will apply the
	// proper amount of maxhp and maxstamina to your character based on
	// CharStats.txt for the amount of vitality the stat added (doesn't work
	// for non characters)
	Op9

	//  doesn't do anything, this has no switch case in the op function.
	Op10

	//  adds opstat.base * statvalue / 100 similar to 1 and 13, the code just
	// does a few more checks
	Op11

	//  doesn't do anything, this has no switch case in the op function.
	Op12

	// adds opstat.base * statvalue / 100 to the value of opstat, this is
	// useable only on items it will not apply the bonus to other unit types
	// (this is why it is used for +% durability, +% level requirement,
	// +% damage, +% defense ).
	Op13
)

/* column names from path_d2.mpq/data/global/excel/ItemStatCost.txt
Stat
ID
Send Other
Signed
Send Bits
Send Param Bits
UpdateAnimRate
Saved
CSvSigned
CSvBits
CSvParam
fCallback
fMin
MinAccr
Encode
Add
Multiply
Divide
ValShift
1.09-Save Bits
1.09-Save Add
Save Bits
Save Add
Save Param Bits
keepzero
op
op param
op base
op stat1
op stat2
op stat3
direct
maxstat
itemspecific
damagerelated
itemevent1
itemeventfunc1
itemevent2
itemeventfunc2
descpriority
descfunc
descval
descstrpos
descstrneg
descstr2
dgrp
dgrpfunc
dgrpval
dgrpstrpos
dgrpstrneg
dgrpstr2
stuff
*eol
*/

var ItemStatCosts map[string]*ItemStatCostRecord

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

			Encode: EncodingType(d.GetNumber("Encode", idx)),

			CallbackEnabled: d.GetNumber("fCallback", idx) > 0,

			CostAdd:      d.GetNumber("Add", idx),
			CostMultiply: d.GetNumber("Multiply", idx),
			ValShift:     d.GetNumber("ValShift", idx),

			OperatorType: OperatorType(d.GetNumber("op", idx)),
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
