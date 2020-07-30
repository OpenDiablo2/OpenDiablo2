package d2datadict

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// MagicPrefix stores all of the magic prefix records
var MagicPrefix map[string]*ItemAffixCommonRecord //nolint:gochecknoglobals // Currently global by
// design
// MagicSuffix stores all of the magic suffix records
var MagicSuffix map[string]*ItemAffixCommonRecord //nolint:gochecknoglobals // Currently global by
// design

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
) map[string]*ItemAffixCommonRecord {
	d := d2common.LoadDataDictionary(file)
	records := createItemAffixRecords(d, superType, subType)
	name := getAffixString(superType, subType)
	log.Printf("Loaded %d %s records", len(records), name)

	return records
}

func createItemAffixRecords(
	d *d2common.DataDictionary,
	superType d2enum.ItemAffixSuperType,
	subType d2enum.ItemAffixSubType,
) map[string]*ItemAffixCommonRecord {
	records := make(map[string]*ItemAffixCommonRecord)

	for d.Next() {
		affix := &ItemAffixCommonRecord{
			Name:           d.String("Name"),
			Version:        d.Number("version"),
			Type:           subType,
			IsPrefix:       superType == d2enum.ItemAffixPrefix,
			IsSuffix:       superType == d2enum.ItemAffixSuffix,
			Spawnable:      d.Bool("spawnable"),
			Rare:           d.Bool("rare"),
			Level:          d.Number("level"),
			MaxLevel:       d.Number("maxlevel"),
			LevelReq:       d.Number("levelreq"),
			Class:          d.String("classspecific"),
			ClassLevelReq:  d.Number("classlevelreq"),
			Frequency:      d.Number("frequency"),
			GroupID:        d.Number("group"),
			Transform:      d.Bool("transform"),
			TransformColor: d.String("transformcolor"),
			PriceAdd:       d.Number("add"),
			PriceScale:     d.Number("multiply"),
		}

		// modifiers (Code references with parameters to be eval'd)
		for i := 1; i <= 3; i++ {
			codeKey := fmt.Sprintf("mod%dcode", i)
			paramKey := fmt.Sprintf("mod%dparam", i)
			minKey := fmt.Sprintf("mod%dmin", i)
			maxKey := fmt.Sprintf("mod%dmax", i)
			modifier := &ItemAffixCommonModifier{
				Code:      d.String(codeKey),
				Parameter: d.Number(paramKey),
				Min:       d.Number(minKey),
				Max:       d.Number(maxKey),
			}
			affix.Modifiers = append(affix.Modifiers, modifier)
		}

		// items to include for spawning
		for i := 1; i <= 7; i++ {
			itemKey := fmt.Sprintf("itype%d", i)
			itemToken := d.String(itemKey)
			affix.ItemInclude = append(affix.ItemInclude, itemToken)
		}

		// items to exclude for spawning
		for i := 1; i <= 7; i++ {
			itemKey := fmt.Sprintf("etype%d", i)
			itemToken := d.String(itemKey)
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

		records[affix.Name] = affix
	}
	if d.Err != nil {
		panic(d.Err)
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
