package d2mapengine

import (
	"log"
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapstamp"
)

type MapGeneratorPreset struct {
	seed   int64
	level  *MapLevel
	engine *MapEngine
}

func (m *MapGeneratorPreset) init(s int64, l *MapLevel, e *MapEngine) {
	m.seed = s
	m.level = l
	m.engine = e
}

func (m *MapGeneratorPreset) generate() {
	rand.Seed(m.seed)

	////////////////////////////////////////////////////////////////////// FIXME
	// TODO: we need to set the difficulty level of the realm in order to pull
	// the right data from level details. testing this for now with normal diff
	// NOTE: we would be setting difficulty level in the realm when a host
	// is connected (the first player)
	diffTestKey := "Normal"
	m.level.act.realm.difficulty = d2datadict.DifficultyLevels[diffTestKey] // hack
	////////////////////////////////////////////////////////////////////////////

	difficulty := m.level.act.realm.difficulty
	details := m.level.details

	tileW, tileH := 0, 0
	switch difficulty.Name {
	case "Normal":
		tileW = details.SizeXNormal
		tileH = details.SizeYNormal
	case "Nightmare":
		tileW = details.SizeXNightmare
		tileH = details.SizeYNightmare
	case "Hell":
		tileW = details.SizeXHell
		tileH = details.SizeYHell
	}

	// TODO: we shouldn't need to cast this to a RegionIdType
	// In the long run, we aren't going to be using hardcoded enumerations
	// we had initially made a list of them for testing, but not necessary now
	levelTypeId := d2enum.RegionIdType(m.level.details.LevelType)
	levelPresetId := m.level.preset.DefinitionId

	m.engine.ResetMap(levelTypeId, tileW+1, tileH+1)
	m.engine.levelType = m.level.types

	stamp := d2mapstamp.LoadStamp(levelTypeId, levelPresetId, -1)
	stampRegionPath := stamp.RegionPath()
	log.Printf("Region Path: %s", stampRegionPath)

	m.engine.PlaceStamp(stamp, 0, 0)
}
