package d2records

import (
	"fmt"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

const (
	maxTreasuresPerRecord = 10
	treasureItemFmt       = "Item%d"
	treasureProbFmt       = "Prob%d"
)

//nolint:funlen // Makes no sense to split
func treasureClassLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(TreasureClass)

	for d.Next() {
		record := &TreasureClassRecord{
			Name:       d.String("Treasure Class"),
			Group:      d.Number("group"),
			Level:      d.Number("level"),
			NumPicks:   d.Number("Picks"),
			FreqUnique: d.Number("Unique"),
			FreqSet:    d.Number("Set"),
			FreqRare:   d.Number("Rare"),
			FreqMagic:  d.Number("Magic"),
			FreqNoDrop: d.Number("NoDrop"),
		}

		if record.Name == "" {
			continue
		}

		for treasureIdx := 0; treasureIdx < maxTreasuresPerRecord; treasureIdx++ {
			treasureColumnKey := fmt.Sprintf(treasureItemFmt, treasureIdx+1)
			probColumnKey := fmt.Sprintf(treasureProbFmt, treasureIdx+1)

			treasureName := d.String(treasureColumnKey)
			if treasureName == "" {
				continue
			}

			prob := d.Number(probColumnKey)

			treasure := &Treasure{
				Code:        treasureName,
				Probability: prob,
			}

			if record.Treasures == nil {
				record.Treasures = []*Treasure{treasure}
			} else {
				record.Treasures = append(record.Treasures, treasure)
			}
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Item.TreasureClass = records

	log.Printf("Loaded %d records records", len(records))

	return nil
}
