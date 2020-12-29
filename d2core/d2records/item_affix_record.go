package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// MagicPrefix stores all of the magic prefix records
type MagicPrefix map[string]*ItemAffixCommonRecord

// MagicSuffix stores all of the magic suffix records
type MagicSuffix map[string]*ItemAffixCommonRecord

// ItemAffixGroups are groups of MagicPrefix/Suffixes
type ItemAffixGroups map[int]*ItemAffixCommonGroup

// ItemAffixCommonModifier is the generic modifier form that prefix/suffix shares
// modifiers are like dynamic properties, they have a key that points to a property
// a parameter for the property, and a min/max value
type ItemAffixCommonModifier = PropertyDescriptor

// ItemAffixCommonGroup is a grouping that is common between prefix/suffix
type ItemAffixCommonGroup struct {
	ID      int
	Members map[string]*ItemAffixCommonRecord
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

// AddMember adds an affix to the group
func (g *ItemAffixCommonGroup) AddMember(a *ItemAffixCommonRecord) {
	if g.Members == nil {
		g.Members = make(map[string]*ItemAffixCommonRecord)
	}

	g.Members[a.Name] = a
}

// GetTotalFrequency returns the cumulative frequency of the affix group
func (g *ItemAffixCommonGroup) GetTotalFrequency() int {
	total := 0

	for _, affix := range g.Members {
		total += affix.Frequency
	}

	return total
}

// ProbabilityToSpawn returns the chance of the affix spawning on an
// item with a given quality level
func (a *ItemAffixCommonRecord) ProbabilityToSpawn(qlvl int) float64 {
	if (qlvl > a.MaxLevel) || (qlvl < a.Level) {
		return 0
	}

	p := float64(a.Frequency) / float64(a.Group.GetTotalFrequency())

	return p
}
