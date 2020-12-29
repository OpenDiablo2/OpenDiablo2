package d2mapstamp

import (
	"math"
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2ds1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

const logPrefix = "Map Stamp"

// NewStampFactory creates a MapStamp factory instance
func NewStampFactory(asset *d2asset.AssetManager, l d2util.LogLevel, entity *d2mapentity.MapEntityFactory) *StampFactory {
	result := &StampFactory{
		asset:  asset,
		entity: entity,
	}

	result.Logger = d2util.NewLogger()
	result.Logger.SetLevel(l)
	result.Logger.SetPrefix(logPrefix)

	return result
}

// StampFactory is responsible for loading map stamps. A stamp can be thought of like a
// preset map configuration, like the various configurations of Act 1 town.
type StampFactory struct {
	asset  *d2asset.AssetManager
	entity *d2mapentity.MapEntityFactory

	*d2util.Logger
}

// LoadStamp loads the Stamp data from file, using the given level type, level preset index, and
// level file index.
func (f *StampFactory) LoadStamp(levelType d2enum.RegionIdType, levelPreset, fileIndex int) *Stamp {
	stamp := &Stamp{
		factory:     f,
		entity:      f.entity,
		regionID:    levelType,
		levelType:   *f.asset.Records.Level.Types[levelType],
		levelPreset: f.asset.Records.Level.Presets[levelPreset],
	}

	for _, levelTypeDt1 := range &stamp.levelType.Files {
		if levelTypeDt1 == "" || levelTypeDt1 == "0" {
			continue
		}

		fileData, err := f.asset.LoadFile("/data/global/tiles/" + levelTypeDt1)
		if err != nil {
			panic(err)
		}

		dt1, err := d2dt1.LoadDT1(fileData)
		if err != nil {
			f.Error(err.Error())
			return nil
		}

		stamp.tiles = append(stamp.tiles, dt1.Tiles...)
	}

	var levelFilesToPick []string

	for _, fileRecord := range stamp.levelPreset.Files {
		if fileRecord != "" && fileRecord != "0" {
			levelFilesToPick = append(levelFilesToPick, fileRecord)
		}
	}

	// nolint:gosec // not a big deal for now
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

	stamp.ds1, err = d2ds1.LoadDS1(fileData)
	if err != nil {
		f.Error(err.Error())
		return nil
	}

	return stamp
}
