package d2records

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadMonPresets loads monster presets from monpresets.txt
func monsterPresetLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(MonPresets)

	for d.Next() {
		act := int32(d.Number("Act"))
		if _, ok := records[act]; !ok {
			records[act] = make([]string, 0)
		}

		records[act] = append(records[act], d.String("Place"))
	}

	if d.Err != nil {
		return d.Err
	}

	log.Printf("Loaded %d MonPreset records", len(records))

	r.Monster.Presets = records

	return nil
}
