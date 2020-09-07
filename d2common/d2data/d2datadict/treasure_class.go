package d2datadict

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

// TreasureClassRecord represents a rule for item drops in diablo 2
type TreasureClassRecord struct {
	Name       string
	Group      int
	Level      int
	NumPicks   int
	FreqUnique int
	FreqSet    int
	FreqRare   int
	FreqMagic  int
	FreqNoDrop int
	Treasures  []*Treasure
}

// Treasure describes a treasure to drop
// the Name is either a reference to an item, or to another treasure class
type Treasure struct {
	Code        string
	Probability int
}

// TreasureClass contains all of the TreasureClassRecords
var TreasureClass map[string]*TreasureClassRecord //nolint:gochecknoglobals // Currently global by design

// LoadTreasureClassRecords loads treasure class records from TreasureClassEx.txt
//nolint:funlen // Makes no sense to split
func LoadTreasureClassRecords(file []byte) {
	TreasureClass = make(map[string]*TreasureClassRecord)

	d := d2txt.LoadDataDictionary(file)

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

		TreasureClass[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d TreasureClass records", len(TreasureClass))
}
