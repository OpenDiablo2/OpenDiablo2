package d2records

const (
	costDivisor = 1024.
)

// NPCs stores the NPCRecords
type NPCs map[string]*NPCRecord

// NPCRecord represents a single line in NPC.txt
// The information has been gathered from [https:// d2mods.info/forum/kb/viewarticle?a=387]
type NPCRecord struct {
	Multipliers      *costMultiplier
	QuestMultipliers map[int]*costMultiplier
	Name             string
	MaxBuy           struct {
		Normal    int
		Nightmare int
		Hell      int
	}
}

type costMultiplier struct {
	// Buy is a percentage of base item price used when an item is bought by NPC
	Buy float64

	// Sell is a percentage of base item price used when an item is sold by NPC
	Sell float64

	// Repair is a percentage of base item price used to calculate the base repair price
	Repair float64
}
