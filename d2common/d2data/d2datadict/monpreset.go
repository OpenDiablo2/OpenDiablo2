package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
	"log"
)

// MonPresets stores monster presets
//nolint:gochecknoglobals // Currently global by design, only written once
var MonPresets map[int32][]string

// LoadMonPresets loads monster presets from monpresets.txt
func LoadMonPresets(file []byte) {
	MonPresets = make(map[int32][]string)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		act := int32(d.Number("Act"))
		if _, ok := MonPresets[act]; !ok {
			MonPresets[act] = make([]string, 0)
		}

		MonPresets[act] = append(MonPresets[act], d.String("Place"))
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonPreset records", len(MonPresets))
}
