package d2enum

type ItemAffixSuperType int
type ItemAffixSubType int

const (
	ItemAffixPrefix ItemAffixSuperType = iota
	ItemAffixSuffix
)

const (
	ItemAffixCommon ItemAffixSubType = iota
	ItemAffixMagic
)
