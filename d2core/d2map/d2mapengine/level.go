package d2mapengine

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type MapLevel struct {
	act           *MapAct
	details       *d2datadict.LevelDetailsRecord
	preset        *d2datadict.LevelPresetRecord
	warps         []*d2datadict.LevelWarpRecord
	substitutions *d2datadict.LevelSubstitutionRecord
	types         *d2datadict.LevelTypeRecord
	generator     MapGenerator
	engine        *MapEngine
	isInit        bool
	isGenerated   bool
}

func (level *MapLevel) isActive() bool {
	return true // todo determine where players are
}

func (level *MapLevel) Advance(elapsed float64) {
	if !level.isActive() {
		return
	}
	level.engine.Advance(elapsed)
}

func (level *MapLevel) Init(act *MapAct, levelId int, engine *MapEngine) {
	if level.isInit {
		return
	}
	if levelId < 1 {
		levelId = 1 // there is a Nonetype map at index 0 in levels.txt
	}
	level.act = act
	level.details = d2datadict.GetLevelDetailsByLevelId(levelId)
	level.preset = d2datadict.GetLevelPresetByLevelId(levelId)
	level.warps = d2datadict.GetLevelWarpsByLevelId(levelId)
	level.substitutions = d2datadict.LevelSubstitutions[level.details.SubType]
	level.types = d2datadict.LevelTypes[d2enum.RegionIdType(level.details.LevelType)]
	level.isInit = true
	level.engine = engine

	switch level.details.LevelGenerationType {
	case d2enum.LevelTypeNone:
		level.generator = nil
	case d2enum.LevelTypeRandomMaze:
		level.generator = &MapGeneratorMaze{}
	case d2enum.LevelTypeWilderness:
		level.generator = &MapGeneratorWilderness{}
	case d2enum.LevelTypePreset:
		level.generator = &MapGeneratorPreset{}
	}

	seed := act.realm.seed
	if level.generator != nil {
		log.Printf("Initializing Level: %s", level.details.Name)
		level.generator.init(seed, level, engine)
	}
}

func (level *MapLevel) GenerateMap() {
	if level.isGenerated {
		return
	}
	log.Printf("Generating Level: %s", level.details.Name)
	level.generator.generate()
	level.isGenerated = true
}
