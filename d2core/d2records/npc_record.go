package d2records

const (
	costDivisor = 1024.
)

// NPCs stores the NPCRecords
type NPCs map[string]*NPCRecord

// NPCRecord represents a single line in NPC.txt
// The information has been gathered from [https:// d2mods.info/forum/kb/viewarticle?a=387]
type NPCRecord struct {
	// Name is an ID pointer to row of this npc in monstats.txt
	Name string

	Multipliers *costMultiplier

	QuestMultipliers map[int]*costMultiplier

	// MaxBuy is the maximum amount of gold an NPC will pay for an item for the corresponding
	// difficulty
	MaxBuy struct {
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
