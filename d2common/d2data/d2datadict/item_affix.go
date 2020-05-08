package d2datadict

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

var MagicPrefixDictionary *d2common.DataDictionary
var MagicSuffixDictionary *d2common.DataDictionary

var MagicPrefixRecords []*ItemAffixCommonRecord
var MagicSuffixRecords []*ItemAffixCommonRecord

var AffixMagicGroups []*ItemAffixCommonGroup

var superType d2enum.ItemAffixSuperType
var subType d2enum.ItemAffixSubType

func LoadMagicPrefix(file []byte) {
	superType = d2enum.ItemAffixPrefix
	subType = d2enum.ItemAffixMagic
	loadDictionary(file, MagicPrefixDictionary, superType, subType)
}

func LoadMagicSuffix(file []byte) {
	superType = d2enum.ItemAffixSuffix
	subType = d2enum.ItemAffixMagic
	loadDictionary(file, MagicSuffixDictionary, superType, subType)
}

func getAffixString(t1 d2enum.ItemAffixSuperType, t2 d2enum.ItemAffixSubType) string {
	var name string = ""

	switch t2 {
	case d2enum.ItemAffixMagic:
		name = "Magic"
	}

	switch t1 {
	case d2enum.ItemAffixPrefix:
		name += "Prefix"
	case d2enum.ItemAffixSuffix:
		name += "Suffix"
	}

	return name

}

func loadDictionary(
	file []byte,
	dict *d2common.DataDictionary,
	superType d2enum.ItemAffixSuperType,
	subType d2enum.ItemAffixSubType,
) {
	dict = d2common.LoadDataDictionary(string(file))
	records := make([]*ItemAffixCommonRecord, 0)

	createItemAffixRecords(dict, records, superType, subType)
	name := getAffixString(superType, subType)
	log.Printf("Loaded %d %s records", len(dict.Data), name)
}

// --- column names from d2exp.mpq:/data/globa/excel/MagicPrefix.txt
// Name
// version
// spawnable
// rare
// level
// maxlevel
// levelreq
// classspecific
// class
// classlevelreq
// frequency
// group
// mod1code
// mod1param
// mod1min
// mod1max
// mod2code
// mod2param
// mod2min
// mod2max
// mod3code
// mod3param
// mod3min
// mod3max
// transform
// transformcolor
// itype1
// itype2
// itype3
// itype4
// itype5
// itype6
// itype7
// etype1
// etype2
// etype3
// etype4
// etype5
// divide
// multiply
// add

func createItemAffixRecords(
	d *d2common.DataDictionary,
	r []*ItemAffixCommonRecord,
	superType d2enum.ItemAffixSuperType,
	subType d2enum.ItemAffixSubType,
) {
	for index, _ := range d.Data {

		affix := &ItemAffixCommonRecord{
			Name:           d.GetString("Name", index),
			Version:        d.GetNumber("version", index),
			Type:           subType,
			IsPrefix:       superType == d2enum.ItemAffixPrefix,
			IsSuffix:       superType == d2enum.ItemAffixSuffix,
			Spawnable:      d.GetNumber("spawnable", index) == 1,
			Rare:           d.GetNumber("rare", index) == 1,
			Level:          d.GetNumber("level", index),
			MaxLevel:       d.GetNumber("maxlevel", index),
			LevelReq:       d.GetNumber("levelreq", index),
			Class:          d.GetString("classspecific", index),
			ClassLevelReq:  d.GetNumber("classlevelreq", index),
			Frequency:      d.GetNumber("frequency", index),
			GroupID:        d.GetNumber("group", index),
			Transform:      d.GetNumber("transform", index) == 1,
			TransformColor: d.GetString("transformcolor", index),
			PriceAdd:       d.GetNumber("add", index),
			PriceScale:     d.GetNumber("multiply", index),
		}

		// modifiers (Property references with parameters to be eval'd)
		for i := 1; i <= 3; i++ {
			codeKey := fmt.Sprintf("mod%dcode", i)
			paramKey := fmt.Sprintf("mod%dparam", i)
			minKey := fmt.Sprintf("mod%dmin", i)
			maxKey := fmt.Sprintf("mod%dmax", i)
			modifier := &ItemAffixCommonModifier{
				Code:      d.GetString(codeKey, index),
				Parameter: d.GetNumber(paramKey, index),
				Min:       d.GetNumber(minKey, index),
				Max:       d.GetNumber(maxKey, index),
			}
			affix.Modifiers = append(affix.Modifiers, modifier)
		}

		// items to include for spawning
		for i := 1; i <= 7; i++ {
			itemKey := fmt.Sprintf("itype%d", i)
			itemToken := d.GetString(itemKey, index)
			affix.ItemInclude = append(affix.ItemInclude, itemToken)
		}

		// items to exclude for spawning
		for i := 1; i <= 7; i++ {
			itemKey := fmt.Sprintf("etype%d", i)
			itemToken := d.GetString(itemKey, index)
			affix.ItemExclude = append(affix.ItemExclude, itemToken)
		}

		// affix groupis
		if ItemAffixGroups == nil {
			ItemAffixGroups = make(map[int]*ItemAffixCommonGroup)
		}

		if _, found := ItemAffixGroups[affix.GroupID]; !found {
			ItemAffixGroup := &ItemAffixCommonGroup{}
			ItemAffixGroup.ID = affix.GroupID
			ItemAffixGroups[affix.GroupID] = ItemAffixGroup
		}

		group := ItemAffixGroups[affix.GroupID]
		group.AddMember(affix)

		r = append(r, affix)
	}
}

var ItemAffixGroups map[int]*ItemAffixCommonGroup

type ItemAffixCommonGroup struct {
	ID      int
	Members map[string]*ItemAffixCommonRecord
}

func (g *ItemAffixCommonGroup) AddMember(a *ItemAffixCommonRecord) {
	if g.Members == nil {
		g.Members = make(map[string]*ItemAffixCommonRecord)
	}
	g.Members[a.Name] = a
}

func (g *ItemAffixCommonGroup) GetTotalFrequency() int {
	total := 0
	for _, affix := range g.Members {
		total += affix.Frequency
	}
	return total
}

type ItemAffixCommonModifier struct {
	Code      string
	Parameter int
	Min       int
	Max       int
}

type ItemAffixCommonRecord struct {
	Name    string
	Version int
	Type    d2enum.ItemAffixSubType

	IsPrefix bool
	IsSuffix bool

	Spawnable bool
	Rare      bool

	Level    int
	MaxLevel int

	LevelReq      int
	Class         string
	ClassLevelReq int

	Frequency int
	GroupID   int
	Group     *ItemAffixCommonGroup

	Modifiers []*ItemAffixCommonModifier

	Transform      bool
	TransformColor string

	ItemInclude []string
	ItemExclude []string

	PriceAdd   int
	PriceScale int
}

func (a *ItemAffixCommonRecord) ProbabilityToSpawn(qlvl int) float64 {
	if (qlvl > a.MaxLevel) || (qlvl < a.Level) {
		return 0.0
	}
	p := (float64)(a.Frequency) / (float64)(a.Group.GetTotalFrequency())
	return p
}
