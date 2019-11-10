package d2enum

type HeroStance int

const (
	HeroStanceIdle         HeroStance = 0
	HeroStanceIdleSelected HeroStance = 1
	HeroStanceApproaching  HeroStance = 2
	HeroStanceSelected     HeroStance = 3
	HeroStanceRetreating   HeroStance = 4
)
