package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func cubeTypeLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(CubeTypes)

	for d.Next() {
		record := &CubeTypeRecord{
			Name:  d.String("cube item class"),
			Token: d.String("Code"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Item.Cube.Types = records

	r.Logger.Infof("Loaded %d Cube Type records", len(records))

	return nil
}
