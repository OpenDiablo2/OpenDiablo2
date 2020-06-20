package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// AutoMapRecord represents one row from d2data.mpq/AutoMap.txt.
type AutoMapRecord struct {
	//LevelName
	//TileName
	//Style
	//StartSequence
	//EndSequence
	//Type1
	//Cel1
	//Type2
	//Cel2
	//Type2
	//Cel3
	//Type4
	//Cel4
}

// AutoMaps contains all rows AutoMap.txt.
var AutoMaps []*AutoMapRecord

// LoadAutoMaps populates AutoMaps with the data from AutoMap.txt.
func LoadAutoMaps(file []byte) {
	// Load data
	d := d2common.LoadDataDictionary(string(file))

	// Create slice
	AutoMaps = make([]*AutoMapRecord, len(d.Data))

	// Populate slice items
	for idx := range d.Data {
		AutoMaps[idx] = &AutoMapRecord{
			//LevelName:
			//TileName:
			//Style:
			//StartSequence:
			//EndSequence:
			//Type1:
			//Cel1:
			//Type2:
			//Cel2:
			//Type2:
			//Cel3:
			//Type4:
			//Cel4:
		}
	}

	log.Printf( /*"Loaded %d AutoMapRecord records"*/ "LoadAutoMaps ran - %d", len(AutoMaps))
}
