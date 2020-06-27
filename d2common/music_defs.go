package d2common

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

type MusicDef struct {
	Region    d2enum.RegionIdType
	InTown    bool
	MusicFile string
}

var musicDefs = [...]MusicDef{
	{d2enum.RegionAct1Town, false, "/data/global/music/Act1/town1.wav"},
	{d2enum.RegionAct1Wilderness, false, "/data/global/music/Act1/wild.wav"},
	{d2enum.RegionAct1Cave, false, "/data/global/music/Act1/caves.wav"},
	{d2enum.RegionAct1Crypt, false, "/data/global/music/Act1/crypt.wav"},
	{d2enum.RegionAct1Monestary, false, "/data/global/music/Act1/monastery.wav"},
	{d2enum.RegionAct1Courtyard, false, "/data/global/music/Act1/wild.wav"},      // ?
	{d2enum.RegionAct1Barracks, false, "/data/global/music/Act1/wild.wav"},       // ?
	{d2enum.RegionAct1Jail, false, "/data/global/music/Act1/wild.wav"},           // ?
	{d2enum.RegionAct1Cathedral, false, "/data/global/music/Act1/monastery.wav"}, // ?
	{d2enum.RegionAct1Catacombs, false, "/data/global/music/Act1/crypt.wav"},     // ?
	{d2enum.RegionAct1Tristram, false, "/data/global/music/Act1/tristram.wav"},   // ?
	{d2enum.RegionAct2Town, false, "/data/global/music/Act2/town2.wav"},
	{d2enum.RegionAct2Sewer, false, "/data/global/music/Act2/sewer.wav"},
	{d2enum.RegionAct2Harem, false, "/data/global/music/Act2/harem.wav"},
	{d2enum.RegionAct2Basement, false, "/data/global/music/Act2/lair.wav"}, // ?
	{d2enum.RegionAct2Desert, false, "/data/global/music/Act2/desrt.wav"},
	{d2enum.RegionAct2Tomb, false, "/data/global/music/Act2/tombs.wav"},
	{d2enum.RegionAct2Lair, false, "/data/global/music/Act2/lair.wav"},
	{d2enum.RegionAct2Arcane, false, "/data/global/music/Act2/sanctuary.wav"}, // ?
	{d2enum.RegionAct3Town, false, "/data/global/music/Act3/town3.wav"},
	{d2enum.RegionAct3Jungle, false, "/data/global/music/Act3/jungle.wav"},
	{d2enum.RegionAct3Kurast, false, "/data/global/music/Act3/kurast.wav"},
	{d2enum.RegionAct3Spider, false, "/data/global/music/Act3/spider.wav"},
	{d2enum.RegionAct3Dungeon, false, "/data/global/music/Act3/kurastsewer.wav"}, // ?
	{d2enum.RegionAct3Sewer, false, "/data/global/music/Act3/kurastsewer.wav"},
	{d2enum.RegionAct4Town, false, "/data/global/music/Act4/town4.wav"},
	{d2enum.RegionAct4Mesa, false, "/data/global/music/Act4/mesa.wav"},
	{d2enum.RegionAct4Lava, false, "/data/global/music/Act4/diablo.wav"}, // ?
	{d2enum.RegonAct5Town, false, "/data/global/music/Act5/xtown.wav"},
	{d2enum.RegionAct5Siege, false, "/data/global/music/Act5/siege.wav"},
	{d2enum.RegionAct5Barricade, false, "/data/global/music/Act5/shenkmusic.wav"}, // ?
	{d2enum.RegionAct5Temple, false, "/data/global/music/Act5/xtemple.wav"},
	{d2enum.RegionAct5IceCaves, false, "/data/global/music/Act5/icecaves.wav"},
	{d2enum.RegionAct5Baal, false, "/data/global/music/Act5/baal.wav"},
	{d2enum.RegionAct5Lava, false, "/data/global/music/Act5/nihlathakmusic.wav"}, // ?
}

func GetMusicDef(regionType d2enum.RegionIdType) *MusicDef {
	for idx := range musicDefs {
		if musicDefs[idx].Region != regionType {
			continue
		}

		return &musicDefs[idx]
	}

	return &musicDefs[0]
}
