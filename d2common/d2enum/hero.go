package d2enum

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

//go:generate stringer -linecomment -type Hero
//go:generate string2enum -samepkg -linecomment -type Hero
