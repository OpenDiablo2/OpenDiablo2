package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

type LevelTypeRecord struct {
	Name      string   // Name
	Id        int      // Id
	Files     []string // File 1 -- File 32
	Beta      bool     // Beta
	Act       int      // Act
	Expansion bool     // Expansion
}

var LevelTypes map[d2enum.RegionIdType]*LevelTypeRecord

func LoadLevelTypes(file []byte) {
	LevelTypes = make(map[d2enum.RegionIdType]*LevelTypeRecord)
	dict := d2common.LoadDataDictionary(string(file))
	for idx := range dict.Data {
		record := &LevelTypeRecord{
			Name: dict.GetString("Name", idx),
			Id:   dict.GetNumber("Id", idx),
			Files: []string{
				dict.GetString("File 1", idx),
				dict.GetString("File 2", idx),
				dict.GetString("File 3", idx),
				dict.GetString("File 4", idx),
				dict.GetString("File 5", idx),
				dict.GetString("File 6", idx),
				dict.GetString("File 7", idx),
				dict.GetString("File 8", idx),
				dict.GetString("File 9", idx),
				dict.GetString("File 10", idx),
				dict.GetString("File 11", idx),
				dict.GetString("File 12", idx),
				dict.GetString("File 13", idx),
				dict.GetString("File 14", idx),
				dict.GetString("File 15", idx),
				dict.GetString("File 16", idx),
				dict.GetString("File 17", idx),
				dict.GetString("File 18", idx),
				dict.GetString("File 19", idx),
				dict.GetString("File 20", idx),
				dict.GetString("File 21", idx),
				dict.GetString("File 22", idx),
				dict.GetString("File 23", idx),
				dict.GetString("File 24", idx),
				dict.GetString("File 25", idx),
				dict.GetString("File 26", idx),
				dict.GetString("File 27", idx),
				dict.GetString("File 28", idx),
				dict.GetString("File 29", idx),
				dict.GetString("File 30", idx),
				dict.GetString("File 31", idx),
				dict.GetString("File 32", idx),
			},
			Beta:      dict.GetNumber("Beta", idx) > 0,
			Act:       dict.GetNumber("Act", idx),
			Expansion: dict.GetNumber("Expansion", idx) > 0,
		}
		LevelTypes[d2enum.RegionIdType(record.Id)] = record
	}
	log.Printf("Loaded %d LevelType records", len(LevelTypes))
}
