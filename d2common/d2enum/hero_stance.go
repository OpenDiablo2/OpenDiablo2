package d2enum

// HeroStance used to render hero stance
type HeroStance int

// HeroStance types
const (
	HeroStanceIdle HeroStance = iota
	HeroStanceIdleSelected
	HeroStanceApproaching
	HeroStanceSelected
	HeroStanceRetreating
)
