package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadLevelTypes loads the LevelTypeRecords
func levelTypesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(LevelTypes, 0)

	for d.Next() {
		record := &LevelTypeRecord{
			[32]string{
				d.String("File 1"),
				d.String("File 2"),
				d.String("File 3"),
				d.String("File 4"),
				d.String("File 5"),
				d.String("File 6"),
				d.String("File 7"),
				d.String("File 8"),
				d.String("File 9"),
				d.String("File 10"),
				d.String("File 11"),
				d.String("File 12"),
				d.String("File 13"),
				d.String("File 14"),
				d.String("File 15"),
				d.String("File 16"),
				d.String("File 17"),
				d.String("File 18"),
				d.String("File 19"),
				d.String("File 20"),
				d.String("File 21"),
				d.String("File 22"),
				d.String("File 23"),
				d.String("File 24"),
				d.String("File 25"),
				d.String("File 26"),
				d.String("File 27"),
				d.String("File 28"),
				d.String("File 29"),
				d.String("File 30"),
				d.String("File 31"),
				d.String("File 32"),
			},
			d.String("Name"),
			d.Number("Id"),
			d.Number("Act"),
			d.Number("Beta") > 0,
			d.Number("Expansion") > 0,
		}

		records = append(records, record)
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d LevelType records", len(records))

	r.Level.Types = records

	return nil
}
