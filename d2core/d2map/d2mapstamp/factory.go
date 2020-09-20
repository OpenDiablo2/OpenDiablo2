package d2mapstamp

import (
	"math"
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

func NewStampFactory(asset *d2asset.AssetManager, entity *d2mapentity.MapEntityFactory) *StampFactory {
	return &StampFactory{asset, entity}
}

type StampFactory struct {
	asset  *d2asset.AssetManager
	entity *d2mapentity.MapEntityFactory
}

// LoadStamp loads the Stamp data from file.
func (f *StampFactory) LoadStamp(levelType d2enum.RegionIdType, levelPreset, fileIndex int) *Stamp {
	stamp := &Stamp{
		entity:      f.entity,
		regionID:    levelType,
		levelType:   *f.asset.Records.Level.Types[levelType],
		levelPreset: f.asset.Records.Level.Presets[levelPreset],
	}

	for _, levelTypeDt1 := range &stamp.levelType.Files {
		if levelTypeDt1 != "" && levelTypeDt1 != "0" {
			fileData, err := f.asset.LoadFile("/data/global/tiles/" + levelTypeDt1)
			if err != nil {
				panic(err)
			}

			dt1, _ := d2dt1.LoadDT1(fileData)

			stamp.tiles = append(stamp.tiles, dt1.Tiles...)
		}
	}

	var levelFilesToPick []string

	for _, fileRecord := range stamp.levelPreset.Files {
		if fileRecord != "" && fileRecord != "0" {
			levelFilesToPick = append(levelFilesToPick, fileRecord)
		}
	}

	levelIndex := int(math.Round(float64(len(levelFilesToPick)-1) * rand.Float64()))
	if fileIndex >= 0 && fileIndex < len(levelFilesToPick) {
		levelIndex = fileIndex
	}

	if levelFilesToPick == nil {
		panic("no level files to pick from")
	}

	stamp.regionPath = levelFilesToPick[levelIndex]
	fileData, err := f.asset.LoadFile("/data/global/tiles/" + stamp.regionPath)

	if err != nil {
		panic(err)
	}

	stamp.ds1, _ = d2ds1.LoadDS1(fileData)

	return stamp
}
