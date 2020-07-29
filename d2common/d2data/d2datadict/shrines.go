package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// ShrineRecord is a representation of a row from shrines.txt
type ShrineRecord struct {
	ShrineType       string // None, Recharge, Booster, or Magic
	ShrineName       string // Name of the Shrine
	Effect           string // Effect on the player
	Code             int    // Unique identifier
	Arg0             int    // ? (0-400)
	Arg1             int    // ? (0-2000)
	DurationFrames   int    // How long the shrine lasts in frames
	ResetTimeMinutes int    // How many minutes until the shrine resets?
	Rarity           int    // 1-3
	EffectClass      int    // 0-4
	LevelMin         int    // 0-32
}

// Shrines contains the Unique Appellations
//nolint:gochecknoglobals // Currently global by design, only written once
var Shrines map[string]*ShrineRecord

// LoadShrines loads Shrines from the supplied file
func LoadShrines(file []byte) {
	Shrines = make(map[string]*ShrineRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &ShrineRecord{
			ShrineType:       d.String("Shrine Type"),
			ShrineName:       d.String("Shrine name"),
			Effect:           d.String("Effect"),
			Code:             d.Number("Code"),
			Arg0:             d.Number("Arg0"),
			Arg1:             d.Number("Arg1"),
			DurationFrames:   d.Number("Duration in frames"),
			ResetTimeMinutes: d.Number("reset time in minutes"),
			Rarity:           d.Number("rarity"),
			EffectClass:      d.Number("effectclass"),
			LevelMin:         d.Number("LevelMin"),
		}
		Shrines[record.ShrineName] = record
	}

	log.Printf("Loaded %d shrines", len(Shrines))
}
