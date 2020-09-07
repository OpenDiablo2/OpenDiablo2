package d2resource

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// MusicDef stores the music definitions of a region
type MusicDef struct {
	Region    d2enum.RegionIdType
	InTown    bool
	MusicFile string
}

func getMusicDefs() []MusicDef {
	return []MusicDef{
		{d2enum.RegionAct1Town, false, BGMAct1Town1},
		{d2enum.RegionAct1Wilderness, false, BGMAct1Wild},
		{d2enum.RegionAct1Cave, false, BGMAct1Caves},
		{d2enum.RegionAct1Crypt, false, BGMAct1Crypt},
		{d2enum.RegionAct1Monestary, false, BGMAct1Monastery},
		{d2enum.RegionAct1Courtyard, false, BGMAct1Monastery},
		{d2enum.RegionAct1Barracks, false, BGMAct1Monastery},
		{d2enum.RegionAct1Jail, false, BGMAct1Monastery},
		{d2enum.RegionAct1Cathedral, false, BGMAct1Monastery},
		{d2enum.RegionAct1Catacombs, false, BGMAct1Monastery},
		{d2enum.RegionAct1Tristram, false, BGMAct1Tristram},
		{d2enum.RegionAct2Town, false, BGMAct2Town2},
		{d2enum.RegionAct2Sewer, false, BGMAct2Sewer},
		{d2enum.RegionAct2Harem, false, BGMAct2Harem},
		{d2enum.RegionAct2Basement, false, BGMAct2Harem},
		{d2enum.RegionAct2Desert, false, BGMAct2Desert},
		{d2enum.RegionAct2Tomb, false, BGMAct2Tombs},
		{d2enum.RegionAct2Lair, false, BGMAct2Lair},
		{d2enum.RegionAct2Arcane, false, BGMAct2Sanctuary},
		{d2enum.RegionAct3Town, false, BGMAct3Town3},
		{d2enum.RegionAct3Jungle, false, BGMAct3Jungle},
		{d2enum.RegionAct3Kurast, false, BGMAct3Kurast},
		{d2enum.RegionAct3Spider, false, BGMAct3Spider},
		{d2enum.RegionAct3Dungeon, false, BGMAct3KurastSewer},
		{d2enum.RegionAct3Sewer, false, BGMAct3KurastSewer},
		{d2enum.RegionAct4Town, false, BGMAct4Town4},
		{d2enum.RegionAct4Mesa, false, BGMAct4Mesa},
		{d2enum.RegionAct4Lava, false, BGMAct4Mesa},
		{d2enum.RegonAct5Town, false, BGMAct5XTown},
		{d2enum.RegionAct5Siege, false, BGMAct5Siege},
		{d2enum.RegionAct5Barricade, false, BGMAct5Siege}, // ?
		{d2enum.RegionAct5Temple, false, BGMAct5XTemple},
		{d2enum.RegionAct5IceCaves, false, BGMAct5IceCaves},
		{d2enum.RegionAct5Baal, false, BGMAct5Baal},
		{d2enum.RegionAct5Lava, false, BGMAct5Nihlathak}, // ?
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
