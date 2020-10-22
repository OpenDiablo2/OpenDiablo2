package d2enum

import "log"

type SkillClass int

const (
	// SkillClassGeneric is ""
	SkillClassGeneric SkillClass = iota
	SkillClassBarbarian
	SkillClassNecromancer
	SkillClassPaladin
	SkillClassAssassin
	SkillClassSorceress
	SkillClassAmazon
	SkillClassDruid
)

// FromToken returns the enum which corresponds to the given class token
func (sc *SkillClass) FromToken(classToken string) SkillClass {
	resource := SkillClassGeneric

	switch classToken {
	case "":
		return SkillClassGeneric
	case "bar":
		return SkillClassBarbarian
	case "nec":
		return SkillClassNecromancer
	case "pal":
		return SkillClassPaladin
	case "ass":
		return SkillClassAssassin
	case "sor":
		return SkillClassSorceress
	case "ama":
		return SkillClassAmazon
	case "dru":
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
