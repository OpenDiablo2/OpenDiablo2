package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

//PlrModeRecord represents a single line in PlrMode.txt
type PlrModeRecord struct {
	//Player animation mode name
	Name string

	//Player animation mode token
	Token string
}

//PlrModes stores the PlrModeRecords
var PlrModes map[string]*PlrModeRecord //nolint:gochecknoglobals // Currently global by design

//LoadPlrModes loads PlrModeRecords into PlrModes
func LoadPlrModes(file []byte) {
	PlrModes = make(map[string]*PlrModeRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &PlrModeRecord{
			Name:  d.String("Name"),
			Token: d.String("Token"),
		}
		PlrModes[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d PlrMode records", len(PlrModes))
}
