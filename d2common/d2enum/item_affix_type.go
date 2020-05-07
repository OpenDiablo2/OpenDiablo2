package d2enum

type AffixSuperType int
type AffixSubType int

const (
	AffixPrefix AffixSuperType = iota
	AffixSuffix
)

const (
	AffixCommon AffixSubType = iota
	AffixMagic
	AffixRare
	AffixUnique
)
