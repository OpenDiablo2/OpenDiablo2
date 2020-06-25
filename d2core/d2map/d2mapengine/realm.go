package d2mapengine

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
)

/*
	A MapRealm represents the state of the maps/levels/quests for a server

	A MapRealm has MapActs
	A MapAct has MapLevels
	A MapLevel has:
		a MapEngine
		a MapGenerator for the level
		data records from the txt files for the level

	The MapRealm is created by the game server

	The first player to connect to the realm becomes the host
	The host determines the difficulty and which quests are completed

	The Realm, Acts, and Levels do not advance unless they are `active`
	Nothing happens in a realm unless it is active
	Levels do not generate maps until the level becomes `active`

	A Level is active if a player is within it OR in an adjacent level
	An Act is active if one of its levels is active
	The Realm is active if and only if one of its Acts is active
*/
type MapRealm struct {
	seed       int64
	difficulty *d2datadict.DifficultyLevelRecord
	acts       map[int]*MapAct
	players    map[string]string
	host       string
}

// Checks if the realm is in an active state
func (realm *MapRealm) isActive() bool {
	return realm.hasActiveActs()
}

// Checks if there is an active act
func (realm *MapRealm) hasActiveActs() bool {
	for _, act := range realm.acts {
		if act.isActive() {
			return true
		}
	}
	return false
}

// Advances the realm, which advances the acts, which advances the levels...
func (realm *MapRealm) Advance(elapsed float64) {
	if !realm.isActive() {
		return
	}
	for _, act := range realm.acts {
		act.Advance(elapsed)
	}
}

// Sets the host of the realm, which determines quest availability for players
func (realm *MapRealm) SetHost(id string) {
	if player, found := realm.players[id]; found {
		realm.host = player
		log.Printf("Host is now %s", id)
	}
}

// Adds a player to the realm
func (realm *MapRealm) AddPlayer(id string, actId int) {
	realm.players[id] = id
	if realm.host == "" {
		realm.SetHost(id)
	}
}

// Removes a player from the realm
func (realm *MapRealm) RemovePlayer(id string) {
	delete(realm.players, id)
}

// Initialize the realm
func (realm *MapRealm) Init(seed int64, engine *MapEngine) {
	// realm.playerStates = make(map[string]*d2mapentitiy.Player)

	log.Printf("Initializing Realm...")
	realm.seed = seed
	engine.SetSeed(seed)
	actIds := d2datadict.GetActIds()
	realm.acts = make(map[int]*MapAct)

	for _, actId := range actIds {
		act := &MapAct{}
		realm.acts[actId] = act

		act.Init(realm, actId, engine)
	}
}

func (realm *MapRealm) GenerateMap(actId, levelId int) {
	realm.acts[actId].GenerateMap(levelId)
}
