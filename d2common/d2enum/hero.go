package d2enum

import "log"

type Hero int

const (
	HeroNone        Hero = 0 //
	HeroBarbarian   Hero = 1 // Barbarian
	HeroNecromancer Hero = 2 // Necromancer
	HeroPaladin     Hero = 3 // Paladin
	HeroAssassin    Hero = 4 // Assassin
	HeroSorceress   Hero = 5 // Sorceress
	HeroAmazon      Hero = 6 // Amazon
	HeroDruid       Hero = 7 // Druid
)

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

//go:generate stringer -linecomment -type Hero
//go:generate string2enum -samepkg -linecomment -type Hero
