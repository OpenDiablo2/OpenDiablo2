package d2datadict

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// MagicPrefix + MagicSuffix store item affix records
var MagicPrefix []*ItemAffixCommonRecord //nolint:gochecknoglobals // Currently global by design
var MagicSuffix []*ItemAffixCommonRecord //nolint:gochecknoglobals // Currently global by design

// LoadMagicPrefix loads MagicPrefix.txt
func LoadMagicPrefix(file []byte) {
	superType := d2enum.ItemAffixPrefix

	subType := d2enum.ItemAffixMagic

	MagicPrefix = loadDictionary(file, superType, subType)
}

// LoadMagicSuffix loads MagicSuffix.txt
func LoadMagicSuffix(file []byte) {
	superType := d2enum.ItemAffixSuffix

	subType := d2enum.ItemAffixMagic

	MagicSuffix = loadDictionary(file, superType, subType)
}

func getAffixString(t1 d2enum.ItemAffixSuperType, t2 d2enum.ItemAffixSubType) string {
	var name = ""

	if t2 == d2enum.ItemAffixMagic {
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
	superType d2enum.ItemAffixSuperType,
	subType d2enum.ItemAffixSubType,
) []*ItemAffixCommonRecord {
	dict := d2common.LoadDataDictionary(string(file))
	records := createItemAffixRecords(dict, superType, subType)
	name := getAffixString(superType, subType)
	log.Printf("Loaded %d %s records", len(dict.Data), name)

	return records
}

func createItemAffixRecords(
	d *d2common.DataDictionary,
	superType d2enum.ItemAffixSuperType,
	subType d2enum.ItemAffixSubType,
) []*ItemAffixCommonRecord {
	records := make([]*ItemAffixCommonRecord, 0)

	for index := range d.Data {
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
		group.addMember(affix)

		records = append(records, affix)
	}

	return records
}

// ItemAffixGroups are groups of MagicPrefix/Suffixes
var ItemAffixGroups map[int]*ItemAffixCommonGroup //nolint:gochecknoglobals // Currently global by design

// ItemAffixCommonGroup is a grouping that is common between prefix/suffix
type ItemAffixCommonGroup struct {
	ID      int
	Members map[string]*ItemAffixCommonRecord
}

func (g *ItemAffixCommonGroup) addMember(a *ItemAffixCommonRecord) {
	if g.Members == nil {
		g.Members = make(map[string]*ItemAffixCommonRecord)
	}

	g.Members[a.Name] = a
}

func (g *ItemAffixCommonGroup) getTotalFrequency() int {
	total := 0

	for _, affix := range g.Members {
		total += affix.Frequency
	}

	return total
}

// ItemAffixCommonModifier is the generic modifier form that prefix/suffix shares
// modifiers are like dynamic properties, they have a key that points to a property
// a parameter for the property, and a min/max value
type ItemAffixCommonModifier struct {
	Code      string
	Parameter int
	Min       int
	Max       int
}

// ItemAffixCommonRecord is a common definition that both prefix and suffix use
type ItemAffixCommonRecord struct {
	Group     *ItemAffixCommonGroup
	Modifiers []*ItemAffixCommonModifier

	ItemInclude []string
	ItemExclude []string

	Name           string
	Class          string
	TransformColor string

	Version int
	Type    d2enum.ItemAffixSubType

	Level    int
	MaxLevel int

	LevelReq      int
	ClassLevelReq int

	Frequency int
	GroupID   int

	PriceAdd   int
	PriceScale int

	IsPrefix bool
	IsSuffix bool

	Spawnable bool
	Rare      bool
	Transform bool
}

// ProbabilityToSpawn returns the chance of the affix spawning on an
// item with a given quality level
func (a *ItemAffixCommonRecord) ProbabilityToSpawn(qlvl int) float64 {
	if (qlvl > a.MaxLevel) || (qlvl < a.Level) {
		return 0.0
	}

	p := float64(a.Frequency) / float64(a.Group.getTotalFrequency())

	return p
}
