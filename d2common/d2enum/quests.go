package d2enum

const (
	// NormalActQuestsNumber is number of quests in standard act
	NormalActQuestsNumber = 6
	// HalfQuestsNumber is number of quests in act 4
	HalfQuestsNumber = 3
)

// ActsNumber is number of acts in game
const ActsNumber = 5

const (
	// Act1 is act 1 in game
	Act1 = iota + 1
	// Act2 is act 2 in game
	Act2
	// Act3 is act 3 in game
	Act3
	// Act4 is act 4 in game
	Act4
	// Act5 is act 4 in game
	Act5
)

/* I think, It should looks like that:
   each quest has its own position in questStatus map
   which should come from save file.
   quests status values:
           - -2 - done
           - -1 - done, need to play animation
           -  0 - not started yet
           - and after that we have "in progress status"
             so for status (from 1 to n) we have appropriate
             quest descriptions and we'll have appropriate
             actions
*/
const (
	QuestStatusCompleted  = iota - 2 // quest completed
	QuestStatusCompleting            // quest completed (need to play animation)
	QuestStatusNotStarted            // quest not started yet
	QuestStatusInProgress            // quest is in progress
)

const (
	// QuestNone describes "no selected quest" status
	QuestNone = iota
	// Quest1 describes quest field 1
	Quest1
	// Quest2 describes quest field 2
	Quest2
	// Quest3 describes quest field 3
	Quest3
	// Quest4 describes quest field 4
	Quest4
	// Quest5 describes quest field 5
	Quest5
	// Quest6 describes quest field 6
	Quest6
)
