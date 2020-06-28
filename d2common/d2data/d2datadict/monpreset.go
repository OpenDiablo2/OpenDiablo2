package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type MonPresetRecord struct {
	Act int
	Place string
}

var MonPresets []MonPresetRecord

func LoadMonPresets(file []byte) {
	dict := d2common.LoadDataDictionary(string(file))
	numRecords := len(dict.Data)
	MonPresets = make([]MonPresetRecord, numRecords)
	for idx := range dict.Data {
		record := MonPresetRecord{
			Act: dict.GetNumber("Act", idx),
			Place: dict.GetString("Place", idx),
		}
		MonPresets[idx] = record
	}
	log.Printf("Loaded %d MonPreset records", len(MonPresets))
}
