package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

// AutoMagic has all of the AutoMagicRecords, used for generating magic properties for spawned items
type AutoMagic []*AutoMagicRecord

// AutoMagicRecord describes rules for automatically generating magic properties when spawning
// items
type AutoMagicRecord struct {
	IncludeItemCodes      [7]string
	ModCode               [3]string
	ExcludeItemCodes      [3]string
	Name                  string
	ModParam              [3]int
	ModMin                [3]int
	ModMax                [3]int
	Class                 d2enum.Hero
	MinSpawnLevel         int
	MaxSpawnLevel         int
	LevelRequirement      int
	Version               int
	ClassLevelRequirement int
	Frequency             int
	Group                 int
	PaletteTransform      int
	CostDivide            int
	CostMultiply          int
	CostAdd               int
	Spawnable             bool
	SpawnOnRare           bool
	Transform             bool
}
