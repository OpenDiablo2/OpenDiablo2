package d2common

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

// MusicDef stores the music definitions of a region
type MusicDef struct {
	Region    d2enum.RegionIdType
	InTown    bool
	MusicFile string
}

func getMusicDefs() []MusicDef {
	return []MusicDef{
		{d2enum.RegionAct1Town, false, d2resource.BGMAct1Town1},
		{d2enum.RegionAct1Wilderness, false, d2resource.BGMAct1Wild},
		{d2enum.RegionAct1Cave, false, d2resource.BGMAct1Caves},
		{d2enum.RegionAct1Crypt, false, d2resource.BGMAct1Crypt},
		{d2enum.RegionAct1Monestary, false, d2resource.BGMAct1Monastery},
		{d2enum.RegionAct1Courtyard, false, d2resource.BGMAct1Monastery},
		{d2enum.RegionAct1Barracks, false, d2resource.BGMAct1Monastery},
		{d2enum.RegionAct1Jail, false, d2resource.BGMAct1Monastery},
		{d2enum.RegionAct1Cathedral, false, d2resource.BGMAct1Monastery},
		{d2enum.RegionAct1Catacombs, false, d2resource.BGMAct1Monastery},
		{d2enum.RegionAct1Tristram, false, d2resource.BGMAct1Tristram},
		{d2enum.RegionAct2Town, false, d2resource.BGMAct2Town2},
		{d2enum.RegionAct2Sewer, false, d2resource.BGMAct2Sewer},
		{d2enum.RegionAct2Harem, false, d2resource.BGMAct2Harem},
		{d2enum.RegionAct2Basement, false, d2resource.BGMAct2Harem},
		{d2enum.RegionAct2Desert, false, d2resource.BGMAct2Desert},
		{d2enum.RegionAct2Tomb, false, d2resource.BGMAct2Tombs},
		{d2enum.RegionAct2Lair, false, d2resource.BGMAct2Lair},
		{d2enum.RegionAct2Arcane, false, d2resource.BGMAct2Sanctuary},
		{d2enum.RegionAct3Town, false, d2resource.BGMAct3Town3},
		{d2enum.RegionAct3Jungle, false, d2resource.BGMAct3Jungle},
		{d2enum.RegionAct3Kurast, false, d2resource.BGMAct3Kurast},
		{d2enum.RegionAct3Spider, false, d2resource.BGMAct3Spider},
		{d2enum.RegionAct3Dungeon, false, d2resource.BGMAct3KurastSewer},
		{d2enum.RegionAct3Sewer, false, d2resource.BGMAct3KurastSewer},
		{d2enum.RegionAct4Town, false, d2resource.BGMAct4Town4},
		{d2enum.RegionAct4Mesa, false, d2resource.BGMAct4Mesa},
		{d2enum.RegionAct4Lava, false, d2resource.BGMAct4Mesa},
		{d2enum.RegonAct5Town, false, d2resource.BGMAct5XTown},
		{d2enum.RegionAct5Siege, false, d2resource.BGMAct5Siege},
		{d2enum.RegionAct5Barricade, false, d2resource.BGMAct5Siege}, // ?
		{d2enum.RegionAct5Temple, false, d2resource.BGMAct5XTemple},
		{d2enum.RegionAct5IceCaves, false, d2resource.BGMAct5IceCaves},
		{d2enum.RegionAct5Baal, false, d2resource.BGMAct5Baal},
		{d2enum.RegionAct5Lava, false, d2resource.BGMAct5Nihlathak}, // ?
	}
}

// GetMusicDef returns the MusicDef of the given region
func GetMusicDef(regionType d2enum.RegionIdType) *MusicDef {
	musicDefs := getMusicDefs()
	for idx := range musicDefs {
		if musicDefs[idx].Region != regionType {
			continue
		}

		return &musicDefs[idx]
	}

	return &musicDefs[0]
}
