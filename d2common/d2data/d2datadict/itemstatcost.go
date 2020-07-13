package d2datadict

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"log"
	"regexp"
	"strings"
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

//nolint:gochecknoglobals // better for lookup
var StatDescriptionFormatStrings = []string{
	"%v %s",
	"%v%% %s",
	"%v %s",
	"%v%% %s",
	"%v%% %s",
	"%v %s %s",
	"%v%% %v %s",
	"%v%% %s %s",
	"%v %s %s",
	"%v %s %s",
	"Repairs 1 Durability In %v Seconds",
	"%v +%v",
	"+%v to %s",
	"+%v to %s %s",
	"%v%% %s",
	"Level %v %s",
	"%v %s (Increases near %v)",
	"%v%% %s (Increases near %v)",
	"",
	"%v%% %s",
	"%v %s",
	"%v%% %s %s",
	"",
	"%v%% %s %s",
	"Level %v %s (%v/%v Charges)",
	"",
	"",
	"%v to %s (%s Only)",
	"%v to %s",
}

// these are the problem children
func stringHack(s string) string {
	// in the lookup table above, `+%v %s` always puts a `+` in front
	// but when the stat value is negative it comes out as `+-`
	s = strings.ReplaceAll(s, "+-", "-")

	// Pesky chance to cast
	s = strings.ReplaceAll(s, "%d%% ", "")

	// oddbal fnid 1 case for min/max damage
	s = strings.ReplaceAll(s, " +%d", "")

	s = strings.Trim(s, " ")

	return s
}

var statValueCountLookup map[int]int //nolint:gochecknoglobals // lookup

// DescString return a string based on the DescFnID
func (r *ItemStatCostRecord) DescString(values ...int) string {
	if r.DescFnID < 0 || r.DescFnID > len(StatDescriptionFormatStrings) {
		return ""
	}

	var result string
	switch r.DescFnID {
	case 0:
		result = r.descFn1(values...)
	case 1:
		result = r.descFn2(values...)
	case 2:
		result = r.descFn3(values...)
	case 3:
		result = r.descFn4(values...)
	case 4:
		result = r.descFn5(values...)
	case 5:
		result = r.descFn6(values...)
	case 6:
		result = r.descFn7(values...)
	case 7:
		result = r.descFn8(values...)
	case 8:
		result = r.descFn9(values...)
	case 10:
		result = r.descFn11(values...)
	case 11:
		result = r.descFn12(values...)
	case 12:
		result = r.descFn13(values...)
	case 13:
		result = r.descFn14(values...)
	case 14:
		result = r.descFn15(values...)
	case 15:
		result = r.descFn16(values...)
	case 16:
		result = r.descFn17(values...)
	case 17:
		result = r.descFn18(values...)
	case 19:
		result = r.descFn20(values...)
	case 21:
		result = r.descFn22(values...)
	case 22:
		result = r.descFn23(values...)
	case 23:
		result = r.descFn24(values...)
	case 26:
		result = r.descFn27(values...)
	case 27:
		result = r.descFn28(values...)
	}

	return result
}

func makeStrings(values ...int) []string {
	strlist := make([]string, len(values)+1)
	for idx := range values {
		strlist[idx] = fmt.Sprintf("%v", values[idx])
	}

	return strlist
}

func (r *ItemStatCostRecord) descFn1(values ...int) string {
	format := StatDescriptionFormatStrings[r.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = r.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey = r.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	var result string

	switch r.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableString)
	case 1:
		result = fmt.Sprintf(format, value, stringTableString)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableString, value)
	default:
		result = ""
	}

	// bugs
	result = strings.Replace(result, "+-", "-", -1)
	result = strings.Replace(result, " +%d", "", -1)

	return result
}

func (r *ItemStatCostRecord) descFn2(values ...int) string {
	format := StatDescriptionFormatStrings[r.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = r.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey = r.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	var result string

	switch r.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableString)
	case 1:
		result = fmt.Sprintf(format, value, stringTableString)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableString, value)
	default:
		result = ""
	}

	// bugs
	result = strings.Replace(result, " +%d", "", -1)

	return result
}

func (r *ItemStatCostRecord) descFn3(values ...int) string {
	format := StatDescriptionFormatStrings[r.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = r.DescStrNeg
	} else {
		stringTableKey = r.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	var result string

	switch r.DescVal {
	case 0:
		format = strings.Split(format, " ")[0]
		result = fmt.Sprintf(format, stringTableString)
	case 1:
		result = fmt.Sprintf(format, value, stringTableString)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableString, value)
	default:
		result = ""
	}

	// bugs
	result = strings.Replace(result, "+-", "-", -1)
	result = strings.Replace(result, " +%d", "", -1)

	return result
}

func (r *ItemStatCostRecord) descFn4(values ...int) string {
	format := StatDescriptionFormatStrings[r.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = r.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey = r.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	var result string

	switch r.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableString)
	case 1:
		result = fmt.Sprintf(format, value, stringTableString)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableString, value)
	default:
		result = ""
	}

	// bugs
	result = strings.Replace(result, " +%d", "", -1)

	return result
}

func (r *ItemStatCostRecord) descFn5(values ...int) string {
	format := StatDescriptionFormatStrings[r.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = r.DescStrNeg
	} else {
		stringTableKey = r.DescStrPos
	}

	stringTableString := d2common.TranslateString(stringTableKey)

	var result string

	switch r.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableString)
	case 1:
		result = fmt.Sprintf(format, value, stringTableString)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableString, value)
	default:
		result = ""
	}

	// bugs
	result = strings.Replace(result, " +%d", "", -1)

	return result
}

func (r *ItemStatCostRecord) descFn6(values ...int) string {
	format := StatDescriptionFormatStrings[r.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey1 string
	if value < 0 {
		stringTableKey1 = r.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey1 = r.DescStrPos
	}

	stringTableStr1 := d2common.TranslateString(stringTableKey1)

	// this stat has an additional string (Based on Character Level)
	stringTableStr2 := d2common.TranslateString(r.DescStr2)

	var result string

	switch r.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableStr1, stringTableStr2)
	case 1:
		result = fmt.Sprintf(format, value, stringTableStr1, stringTableStr2)
	case 2:
		formatSplit := strings.Split(format, " ")
		format = strings.Join(reverseStringSlice(formatSplit), " ")
		result = fmt.Sprintf(format, stringTableStr1, value)
	default:
		result = ""
	}

	// bugs
	result = strings.Replace(result, "+-", "-", -1)
	result = strings.Replace(result, " +%d", "", -1)

	return result
}

func (r *ItemStatCostRecord) descFn7(values ...int) string {
	format := StatDescriptionFormatStrings[r.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey string
	if value < 0 {
		stringTableKey = r.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey = r.DescStrPos
	}

	stringTableStr1 := d2common.TranslateString(stringTableKey)

	// this stat has an additional string (Based on Character Level)
	stringTableStr2 := d2common.TranslateString(r.DescStr2)

	var result string

	switch r.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableStr1, stringTableStr2)
	case 1:
		result = fmt.Sprintf(format, value, stringTableStr1, stringTableStr2)
	case 2:
		formatSplit := strings.Split(format, " ")
		// formatSplit = reverseStringSlice(formatSplit)
		formatSplit[0], formatSplit[1] = formatSplit[1], formatSplit[0]
		format = strings.Join(formatSplit, " ")
		result = fmt.Sprintf(format, stringTableStr1, value, stringTableStr2)
	default:
		result = ""
	}

	return result
}

func (r *ItemStatCostRecord) descFn8(values ...int) string {
	format := StatDescriptionFormatStrings[r.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey1 string
	if value < 0 {
		stringTableKey1 = r.DescStrNeg
	} else {
		format = strings.Join([]string{"+", format}, "")
		stringTableKey1 = r.DescStrPos
	}

	stringTableStr1 := d2common.TranslateString(stringTableKey1)

	// this stat has an additional string (Based on Character Level)
	stringTableStr2 := d2common.TranslateString(r.DescStr2)

	var result string

	switch r.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableStr1, stringTableStr2)
	case 1:
		result = fmt.Sprintf(format, value, stringTableStr1, stringTableStr2)
	case 2:
		formatSplit := strings.Split(format, " ")
		formatSplit[0], formatSplit[1] = formatSplit[1], formatSplit[0]
		format = strings.Join(formatSplit, " ")
		result = fmt.Sprintf(format, stringTableStr1, value, stringTableStr2)
	default:
		result = ""
	}

	return result
}

func (r *ItemStatCostRecord) descFn9(values ...int) string {
	format := StatDescriptionFormatStrings[r.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	var stringTableKey1 string
	if value < 0 {
		stringTableKey1 = r.DescStrNeg
	} else {
		stringTableKey1 = r.DescStrPos
	}

	stringTableStr1 := d2common.TranslateString(stringTableKey1)

	// this stat has an additional string (Based on Character Level)
	stringTableStr2 := d2common.TranslateString(r.DescStr2)

	var result string

	switch r.DescVal {
	case 0:
		result = fmt.Sprintf(format, stringTableStr1, stringTableStr2)
	case 1:
		result = fmt.Sprintf(format, value, stringTableStr1, stringTableStr2)
	case 2:
		formatSplit := strings.Split(format, " ")
		formatSplit[0], formatSplit[1] = formatSplit[1], formatSplit[0]
		format = strings.Join(formatSplit, " ")
		result = fmt.Sprintf(format, stringTableStr1, value, stringTableStr2)
	default:
		result = ""
	}

	return result
}

func (r *ItemStatCostRecord) descFn11(values ...int) string {
	// we know there is only one value for this stat
	value := values[0]

	// the only stat to use this fn is "Repairs durability in X seconds"
	format := d2common.TranslateString(r.DescStrPos)

	return fmt.Sprintf(format, value)
}

func (r *ItemStatCostRecord) descFn12(values ...int) string {
	format := StatDescriptionFormatStrings[r.DescFnID]

	// we know there is only one value for this stat
	value := values[0]

	str1 := d2common.TranslateString(r.DescStrPos)

	return fmt.Sprintf(format, str1, value)
}

func (r *ItemStatCostRecord) descFn13(values ...int) string {
	format := StatDescriptionFormatStrings[r.DescFnID]
	numSkills, heroIndex := values[0], values[1]

	heroMap := map[int]d2enum.Hero{
		int(d2enum.HeroAmazon): d2enum.HeroAmazon,
		int(d2enum.HeroSorceress): d2enum.HeroSorceress,
		int(d2enum.HeroNecromancer): d2enum.HeroNecromancer,
		int(d2enum.HeroPaladin): d2enum.HeroPaladin,
		int(d2enum.HeroBarbarian): d2enum.HeroBarbarian,
		int(d2enum.HeroDruid): d2enum.HeroDruid,
		int(d2enum.HeroAssassin): d2enum.HeroAssassin,
	}

	classRecord := CharStats[heroMap[heroIndex]]
	descStr1 := d2common.TranslateString(classRecord.SkillStrAll)
	result := fmt.Sprintf(format, numSkills, descStr1)

	// bugs
	result = strings.Replace(result, "+-", "-", -1)

	return result
}

func (r *ItemStatCostRecord) descFn14(values ...int) string {
	numSkills, heroIndex, skillTabIndex:= values[0], values[1], values[2]

	if skillTabIndex > 2 || skillTabIndex < 0 {
		skillTabIndex = 0
	}

	heroMap := map[int]d2enum.Hero{
		int(d2enum.HeroAmazon): d2enum.HeroAmazon,
		int(d2enum.HeroSorceress): d2enum.HeroSorceress,
		int(d2enum.HeroNecromancer): d2enum.HeroNecromancer,
		int(d2enum.HeroPaladin): d2enum.HeroPaladin,
		int(d2enum.HeroBarbarian): d2enum.HeroBarbarian,
		int(d2enum.HeroDruid): d2enum.HeroDruid,
		int(d2enum.HeroAssassin): d2enum.HeroAssassin,
	}

	classRecord := CharStats[heroMap[heroIndex]]
	skillTabKey := classRecord.SkillStrTab[skillTabIndex]
	classOnlyKey := classRecord.SkillStrClassOnly

	skillTabStr := d2common.TranslateString(skillTabKey) + " %v"
	skillTabStr = strings.ReplaceAll(skillTabStr, "%d", "%v")
	classOnlyStr := d2common.TranslateString(classOnlyKey)
	result := fmt.Sprintf(skillTabStr, numSkills, classOnlyStr)

	// bugs
	result = strings.Replace(result, "+-", "-", -1)

	return result
}

func (r *ItemStatCostRecord) descFn15(values ...int) string {
	return ""
}

func (r *ItemStatCostRecord) descFn16(values ...int) string {
	return ""
}

func (r *ItemStatCostRecord) descFn17(values ...int) string {
	return ""
}

func (r *ItemStatCostRecord) descFn18(values ...int) string {
	return ""
}

func (r *ItemStatCostRecord) descFn20(values ...int) string {
	return ""
}

func (r *ItemStatCostRecord) descFn22(values ...int) string {
	return ""
}

func (r *ItemStatCostRecord) descFn23(values ...int) string {
	return ""
}

func (r *ItemStatCostRecord) descFn24(values ...int) string {
	return ""
}

func (r *ItemStatCostRecord) descFn27(values ...int) string {
	return ""
}

func (r *ItemStatCostRecord) descFn28(values ...int) string {
	return ""
}

// DescGroupString return a string based on the DescGroupFuncID
func (r *ItemStatCostRecord) DescGroupString(a ...interface{}) string {
	if r.DescGroupFuncID < 0 || r.DescGroupFuncID > len(StatDescriptionFormatStrings) {
		return ""
	}

	format := StatDescriptionFormatStrings[r.DescGroupFuncID]

	return fmt.Sprintf(format, a...)
}

// NumStatValues returns the number of values a stat instance for this
// record should have
func (r *ItemStatCostRecord) NumStatValues() int {
	if num, found := statValueCountLookup[r.DescGroupFuncID]; found {
		return num
	}

	if statValueCountLookup == nil {
		statValueCountLookup = make(map[int]int)
	}

	format := StatDescriptionFormatStrings[r.DescGroupFuncID]
	pattern := regexp.MustCompile("%v")
	matches := pattern.FindAllStringIndex(format, -1)
	num := len(matches)
	statValueCountLookup[r.DescGroupFuncID] = num

	return num
}

func reverseStringSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

// ItemStatCosts stores all of the ItemStatCostRecords
//nolint:gochecknoglobals // Currently global by design
var ItemStatCosts map[string]*ItemStatCostRecord

// LoadItemStatCosts loads ItemStatCostRecord's from text
func LoadItemStatCosts(file []byte) {
	ItemStatCosts = make(map[string]*ItemStatCostRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &ItemStatCostRecord{
			Name:  d.String("Stat"),
			Index: d.Number("ID"),

			Signed:   d.Number("Signed") > 0,
			KeepZero: d.Number("keepzero") > 0,

			// Ranged:  d.Number("Ranged") > 0,
			MinAccr: d.Number("MinAccr"),

			UpdateAnimRate: d.Number("UpdateAnimRate") > 0,

			SendOther: d.Number("Send Other") > 0,
			SendBits:  d.Number("Send Bits"),
			SendParam: d.Number("Send Param Bits"),

			Saved:         d.Number("CSvBits") > 0,
			SavedSigned:   d.Number("CSvSigned") > 0,
			SavedBits:     d.Number("CSvBits"),
			SaveBits:      d.Number("Save Bits"),
			SaveAdd:       d.Number("Save Add"),
			SaveParamBits: d.Number("Save Param Bits"),

			Encode: d2enum.EncodingType(d.Number("Encode")),

			CallbackEnabled: d.Number("fCallback") > 0,

			CostAdd:      d.Number("Add"),
			CostMultiply: d.Number("Multiply"),
			ValShift:     d.Number("ValShift"),

			OperatorType: d2enum.OperatorType(d.Number("op")),
			OpParam:      d.Number("op param"),
			OpBase:       d.String("op base"),
			OpStat1:      d.String("op stat1"),
			OpStat2:      d.String("op stat2"),
			OpStat3:      d.String("op stat3"),

			Direct:  d.Number("direct") > 0,
			MaxStat: d.String("maxstat"),

			ItemSpecific:  d.Number("itemspecific") > 0,
			DamageRelated: d.Number("damagerelated") > 0,

			EventID1:     d2enum.GetItemEventType(d.String("itemevent1")),
			EventID2:     d2enum.GetItemEventType(d.String("itemevent2")),
			EventFuncID1: d2enum.ItemEventFuncID(d.Number("itemeventfunc1")),
			EventFuncID2: d2enum.ItemEventFuncID(d.Number("itemeventfunc2")),

			DescPriority: d.Number("descpriority"),
			DescFnID:     d.Number("descfunc") - 1,
			// DescVal:      d.Number("descval"), // needs special handling
			DescStrPos: d.String("descstrpos"),
			DescStrNeg: d.String("descstrneg"),
			DescStr2:   d.String("descstr2"),

			DescGroup:       d.Number("dgrp"),
			DescGroupFuncID: d.Number("dgrpfunc"),

			DescGroupVal:    d.Number("dgrpval"),
			DescGroupStrPos: d.String("dgrpstrpos"),
			DescGroupStrNeg: d.String("dgrpstrneg"),
			DescGroupStr2:   d.String("dgrpstr2"),

			Stuff: d.String("stuff"),
		}

		descValStr := d.String("descval")
		switch descValStr {
		case "2":
			record.DescVal = 2
		case "0":
			record.DescVal = 0
		default:
			// handle empty fields, seems like they should have been 1
			record.DescVal = 1
		}

		ItemStatCosts[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d ItemStatCost records", len(ItemStatCosts))
}
