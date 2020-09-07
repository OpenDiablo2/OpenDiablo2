package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
	"log"
)

const (
	costDivisor = 1024.
)

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

// NPCs stores the NPCRecords
var NPCs map[string]*NPCRecord // nolint:gochecknoglobals // Currently global by design

// LoadNPCs loads NPCRecords into NPCs
func LoadNPCs(file []byte) {
	NPCs = make(map[string]*NPCRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &NPCRecord{
			Name: d.String("npc"),
			Multipliers: &costMultiplier{
				Buy:    float64(d.Number("buy mult")) / costDivisor,
				Sell:   float64(d.Number("sell mult")) / costDivisor,
				Repair: float64(d.Number("rep mult")) / costDivisor,
			},
			MaxBuy: struct {
				Normal    int
				Nightmare int
				Hell      int
			}{
				Normal:    d.Number("max buy"),
				Nightmare: d.Number("max buy (N)"),
				Hell:      d.Number("max buy (H)"),
			},
		}

		record.QuestMultipliers = make(map[int]*costMultiplier)

		if flagStr := d.String("questflag A"); flagStr != "" {
			flag := d.Number("questflag A")
			record.QuestMultipliers[flag] = &costMultiplier{
				float64(d.Number("questbuymult A")) / costDivisor,
				float64(d.Number("questsellmult A")) / costDivisor,
				float64(d.Number("questrepmult A")) / costDivisor,
			}
		}

		if flagStr := d.String("questflag B"); flagStr != "" {
			flag := d.Number("questflag B")
			record.QuestMultipliers[flag] = &costMultiplier{
				float64(d.Number("questbuymult B")) / costDivisor,
				float64(d.Number("questsellmult B")) / costDivisor,
				float64(d.Number("questrepmult B")) / costDivisor,
			}
		}

		if flagStr := d.String("questflag C"); flagStr != "" {
			flag := d.Number("questflag C")
			record.QuestMultipliers[flag] = &costMultiplier{
				float64(d.Number("questbuymult C")) / costDivisor,
				float64(d.Number("questsellmult C")) / costDivisor,
				float64(d.Number("questrepmult C")) / costDivisor,
			}
		}

		NPCs[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d NPC records", len(NPCs))
}
