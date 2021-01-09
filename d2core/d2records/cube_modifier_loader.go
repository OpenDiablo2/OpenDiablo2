package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func cubeModifierLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(CubeModifiers)

	for d.Next() {
		record := &CubeModifierRecord{
			Name:  d.String("cube modifier type"),
			Token: d.String("Code"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Item.Cube.Modifiers = records

	r.Logger.Infof("Loaded %d Cube Modifier records", len(records))

	return nil
}
