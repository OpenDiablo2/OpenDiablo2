package d2enum

import "log"

//go:generate stringer -linecomment -type Hero
//go:generate string2enum -samepkg -linecomment -type Hero

// Hero is used for different types of hero's
type Hero int

// Heroes
const (
	HeroNone        Hero = iota //
	HeroBarbarian               // Barbarian
	HeroNecromancer             // Necromancer
	HeroPaladin                 // Paladin
	HeroAssassin                // Assassin
	HeroSorceress               // Sorceress
	HeroAmazon                  // Amazon
	HeroDruid                   // Druid
)

// GetToken returns a 2 letter token
func (h Hero) GetToken() string {
	switch h {
	case HeroBarbarian:
		return "BA"
	case HeroNecromancer:
		return "NE"
	case HeroPaladin:
		return "PA"
	case HeroAssassin:
		return "AI"
	case HeroSorceress:
		return "SO"
	case HeroAmazon:
		return "AM"
	case HeroDruid:
		return "DZ"
	default:
		log.Fatalf("Unknown hero token: %d", h)
	}

	return ""
}

// GetToken3 returns a 3 letter token
func (h Hero) GetToken3() string {
	switch h {
	case HeroBarbarian:
		return "BAR"
	case HeroNecromancer:
		return "NEC"
	case HeroPaladin:
		return "PAL"
	case HeroAssassin:
		return "ASS"
	case HeroSorceress:
		return "SOR"
	case HeroAmazon:
		return "AMA"
	case HeroDruid:
		return "DRU"
	default:
		log.Fatalf("Unknown hero token: %d", h)
	}

	return ""
}
