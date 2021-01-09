package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func npcLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(NPCs)

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

		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.NPCs = records

	r.Logger.Infof("Loaded %d NPC records", len(records))

	return nil
}
