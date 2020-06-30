package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// MonPresets stores monster presets
//nolint:gochecknoglobals // Currently global by design, only written once
var MonPresets [][]string

// LoadMonPresets loads monster presets from monpresets.txt
func LoadMonPresets(file []byte) {
	dict := d2common.LoadDataDictionary(string(file))
	numRecords := len(dict.Data)
	MonPresets = make([][]string, numRecords)

	for idx := range MonPresets {
		MonPresets[idx] = make([]string, numRecords)
	}

	lastAct := 0
	placeIdx := 0

	for dictIdx := range dict.Data {
		act := dict.GetNumber("Act", dictIdx)
		if act != lastAct {
			placeIdx = 0
		}

		MonPresets[act][placeIdx] = dict.GetString("Place", dictIdx)
		lastAct = act

		placeIdx++
	}

	log.Printf("Loaded %d MonPreset records", len(MonPresets))
}
