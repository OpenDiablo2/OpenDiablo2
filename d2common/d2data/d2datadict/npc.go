package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

//NPCRecord represents a single line in NPC.txt
//The information has been gathered from [https://d2mods.info/forum/kb/viewarticle?a=387]
type NPCRecord struct {
	//ID pointer to row of this npc in monstats.txt
	NPC string

	//Percentage of base item price used when an item is bought by NPC
	//Actual value in file is fraction of 1024
	//Notice: BuyMultiplier used when selling an item to an NPC
	BuyMultiplier float64

	//Percentage of base item price used when an item is sold by NPC
	//Actual value in file is fraction of 1024
	//Notice: SellMultiplier used when buying an item from an NPC
	SellMultiplier float64

	//Percentage of base item price used to calculate the base repair price
	//Actual value in file is fraction of 1024
	RepairMultiplier float64

	//The quest flag that must be checked to activate the corresponding QuestBuyMultiplier
	QuestFlagA int

	//Percentage of change of multiplier when the corresponding quest flag is checked
	//Example: Quest flag A is checked, so the item price will be: (base price) * BuyMultiplier * QuestBuyMultiplierA
	//Actual value in file is fraction of 1024
	QuestBuyMultiplierA float64

	//Percentage of change of multiplier when the corresponding quest flag is checked
	//Actual value in file is fraction of 1024
	QuestSellMultiplierA float64

	//Percentage of change of multiplier when the corresponding quest flag is checked
	//Actual value in file is fraction of 1024
	QuestRepairMultiplierA float64

	//ditto, three quests maximum
	QuestFlagB             int
	QuestBuyMultiplierB    float64
	QuestSellMultiplierB   float64
	QuestRepairMultiplierB float64

	QuestFlagC             int
	QuestBuyMultiplierC    float64
	QuestSellMultiplierC   float64
	QuestRepairMultiplierC float64

	//The maximum amount of gold an NPC will pay for an item for the corresponding difficulty
	MaxBuy          int
	MaxBuyNightmare int
	MaxBuyHell      int
}

//NPCs stores the NPCRecords
var NPCs map[string]*NPCRecord //nolint:gochecknoglobals // Currently global by design

//LoadNPCs loads NPCRecords into NPCs
func LoadNPCs(file []byte) {
	NPCs = make(map[string]*NPCRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &NPCRecord{
			NPC:                    d.String("npc"),
			BuyMultiplier:          float64(d.Number("buy mult")) / 1024.,
			SellMultiplier:         float64(d.Number("sell mult")) / 1024.,
			RepairMultiplier:       float64(d.Number("rep mult")) / 1024.,
			QuestFlagA:             d.Number("questflag A"),
			QuestBuyMultiplierA:    float64(d.Number("questbuymult A")) / 1024.,
			QuestSellMultiplierA:   float64(d.Number("questsellmult A")) / 1024.,
			QuestRepairMultiplierA: float64(d.Number("questrepmult A")) / 1024.,
			QuestFlagB:             d.Number("questflag B"),
			QuestBuyMultiplierB:    float64(d.Number("questbuymult B")) / 1024.,
			QuestSellMultiplierB:   float64(d.Number("questsellmult B")) / 1024.,
			QuestRepairMultiplierB: float64(d.Number("questrepmult B")) / 1024.,
			QuestFlagC:             d.Number("questflag C"),
			QuestBuyMultiplierC:    float64(d.Number("questbuymult C")) / 1024.,
			QuestSellMultiplierC:   float64(d.Number("questsellmult C")) / 1024.,
			QuestRepairMultiplierC: float64(d.Number("questrepmult C")) / 1024.,
			MaxBuy:                 d.Number("max buy"),
			MaxBuyNightmare:        d.Number("max buy (N)"),
			MaxBuyHell:             d.Number("max buy (H)"),
		}
		NPCs[record.NPC] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d NPC records", len(NPCs))
}
