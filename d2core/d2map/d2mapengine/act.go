package d2mapengine

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	// "github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type MapAct struct {
	realm  *MapRealm
	id     int
	levels map[int]*MapLevel
}

func (act *MapAct) isActive() bool {
	for _, level := range act.levels {
		if level.isActive() {
			return true
		}
	}
	return false
}

func (act *MapAct) Advance(elapsed float64) {
	if !act.isActive() {
		return
	}
	for _, level := range act.levels {
		level.Advance(elapsed)
	}
}

func (act *MapAct) Init(realm *MapRealm, actIndex int) {
	act.realm = realm
	act.levels = make(map[int]*MapLevel)
	act.id = actIndex

	actLevelRecords := d2datadict.GetLevelDetailsByActId(actIndex)

	log.Printf("Initializing Act %d", actIndex)
	for _, record := range actLevelRecords {
		level := &MapLevel{}
		levelId := record.Id
		level.Init(act, levelId)
		act.levels[levelId] = level
	}

	act.GenerateTown() // ensures that starting point is known for first player
}

func (act *MapAct) GenerateTown() {
	townId := d2datadict.GetFirstLevelIdByActId(act.id)
	act.levels[townId].GenerateMap()
}

func (act *MapAct) GenerateMap(levelId int) {
	log.Printf("Generating map in Act %d", act.id)
	act.levels[levelId].GenerateMap()
}
