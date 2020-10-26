package d2enum

import "log"

// SkillClass represents the skills for a character class
type SkillClass int

// Skill classes
const (
	SkillClassGeneric SkillClass = iota
	SkillClassBarbarian
	SkillClassNecromancer
	SkillClassPaladin
	SkillClassAssassin
	SkillClassSorceress
	SkillClassAmazon
	SkillClassDruid
)

// Skill class tokens
const (
	SkillClassTokenGeneric     = ""
	SkillClassTokenBarbarian   = "bar"
	SkillClassTokenNecromancer = "nec"
	SkillClassTokenPaladin     = "pal"
	SkillClassTokenAssassin    = "ass"
	SkillClassTokenSorceress   = "sor"
	SkillClassTokenAmazon      = "ama"
	SkillClassTokenDruid       = "dru"
)

// FromToken returns the enum which corresponds to the given class token
func (sc *SkillClass) FromToken(classToken string) SkillClass {
	resource := SkillClassGeneric

	switch classToken {
	case SkillClassTokenGeneric:
		return SkillClassGeneric
	case SkillClassTokenBarbarian:
		return SkillClassBarbarian
	case SkillClassTokenNecromancer:
		return SkillClassNecromancer
	case SkillClassTokenPaladin:
		return SkillClassPaladin
	case SkillClassTokenAssassin:
		return SkillClassAssassin
	case SkillClassTokenSorceress:
		return SkillClassSorceress
	case SkillClassTokenAmazon:
		return SkillClassAmazon
	case SkillClassTokenDruid:
		return SkillClassDruid
	default:
		log.Fatalf("Unknown skill class token: '%s'", classToken)
	}

	// should not be reached
	return resource
}

// GetToken returns a string token for the enum
func (sc SkillClass) GetToken() string {
	switch sc {
	case SkillClassGeneric:
		return ""
	case SkillClassBarbarian:
		return "bar"
	case SkillClassNecromancer:
		return "nec"
	case SkillClassPaladin:
		return "pal"
	case SkillClassAssassin:
		return "ass"
	case SkillClassSorceress:
		return "sor"
	case SkillClassAmazon:
		return "ama"
	case SkillClassDruid:
		return "dru"
	default:
		log.Fatalf("Unknown skill class token: %v", sc)
	}

	// should not be reached
	return ""
}
