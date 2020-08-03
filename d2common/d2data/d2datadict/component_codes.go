package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type ComponentCodeRecord struct {
	Component string
	Code      string
}

var ComponentCodes map[string]*ComponentCodeRecord

func LoadComponentCodes(file []byte) {
	ComponentCodes = make(map[string]*ComponentCodeRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &ComponentCodeRecord{
			Component: d.String("component"),
			Code:      d.String("code"),
		}
		ComponentCodes[record.Component] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d ComponentCode records", len(ComponentCodes))

}
