package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type LevelDetailsRecord struct {
	Name      string
	Id        int
	LevelName string
	LevelType int
}

var LevelDetails []LevelDetailsRecord

func LoadLevelDetails(file []byte) {
	dict := d2common.LoadDataDictionary(string(file))
	numRecords := len(dict.Data)
	LevelDetails = make([]LevelDetailsRecord, numRecords)

	for idx := range dict.Data {
		// A row(on line 36) "Expansion" is used to separate between base levels and expansion levels.
		if dict.GetString("Name", idx) == "Expansion" {
			continue
		}

		record := LevelDetailsRecord{
			Id:        dict.GetNumber("Id", idx),
			Name:      dict.GetString("Name", idx),
			LevelName: dict.GetString("LevelName", idx),
			LevelType: dict.GetNumber("LevelType", idx),
		}

		// TODO: Mapping is incorrect(added only temporary for testing purposes),
		// LevelType can't be used for indexing since multiple records have the same level types. Only the latest record of that level type
		// will be available in the array, since the others get overriden.
		LevelDetails[record.LevelType] = record
	}

	log.Printf("Loaded %d level details records", len(LevelWarps))
}
