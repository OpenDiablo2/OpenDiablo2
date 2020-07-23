package d2enum

// RegionIdType represents a region ID
type RegionIdType int //nolint:golint,stylecheck // many changed needed when changing to ID

// Regions
const (
	RegionNone RegionIdType = iota
	RegionAct1Town
	RegionAct1Wilderness
	RegionAct1Cave
	RegionAct1Crypt
	RegionAct1Monestary
	RegionAct1Courtyard
	RegionAct1Barracks
	RegionAct1Jail
	RegionAct1Cathedral
	RegionAct1Catacombs
	RegionAct1Tristram
	RegionAct2Town
	RegionAct2Sewer
	RegionAct2Harem
	RegionAct2Basement
	RegionAct2Desert
	RegionAct2Tomb
	RegionAct2Lair
	RegionAct2Arcane
	RegionAct3Town
	RegionAct3Jungle
	RegionAct3Kurast
	RegionAct3Spider
	RegionAct3Dungeon
	RegionAct3Sewer
	RegionAct4Town
	RegionAct4Mesa
	RegionAct4Lava
	RegonAct5Town
	RegionAct5Siege
	RegionAct5Barricade
	RegionAct5Temple
	RegionAct5IceCaves
	RegionAct5Baal
	RegionAct5Lava
)
