package d2enum

// ItemAffixSuperType represents a item affix super type
type ItemAffixSuperType int

// Super types
const (
	ItemAffixPrefix ItemAffixSuperType = iota
	ItemAffixSuffix
)

// ItemAffixSubType represents a item affix sub type
type ItemAffixSubType int

// Sub types
const (
	ItemAffixCommon ItemAffixSubType = iota
	ItemAffixMagic
)
