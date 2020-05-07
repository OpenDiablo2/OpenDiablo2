package d2enum

type ItemAffixType int

const (
	ItemAffixCommon ItemAffixType = iota
	ItemAffixMagic
	ItemAffixRare
	ItemAffixUnique
)
