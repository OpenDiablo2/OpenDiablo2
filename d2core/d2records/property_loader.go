package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadProperties loads gem records into a map[string]*PropertiesRecord
func propertyLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(Properties)

	for d.Next() {
		record := &PropertyRecord{
			Code:   d.String("code"),
			Active: d.String("*done"),
			Stats: [7]*PropertyStatRecord{
				{
					SetID:      d.Number("set1"),
					Value:      d.Number("val1"),
					FunctionID: d.Number("func1"),
					StatCode:   d.String("stat1"),
				},
				{
					SetID:      d.Number("set2"),
					Value:      d.Number("val2"),
					FunctionID: d.Number("func2"),
					StatCode:   d.String("stat2"),
				},
				{
					SetID:      d.Number("set3"),
					Value:      d.Number("val3"),
					FunctionID: d.Number("func3"),
					StatCode:   d.String("stat3"),
				},
				{
					SetID:      d.Number("set4"),
					Value:      d.Number("val4"),
					FunctionID: d.Number("func4"),
					StatCode:   d.String("stat4"),
				},
				{
					SetID:      d.Number("set5"),
					Value:      d.Number("val5"),
					FunctionID: d.Number("func5"),
					StatCode:   d.String("stat5"),
				},
				{
					SetID:      d.Number("set6"),
					Value:      d.Number("val6"),
					FunctionID: d.Number("func6"),
					StatCode:   d.String("stat6"),
				},
				{
					SetID:      d.Number("set7"),
					Value:      d.Number("val7"),
					FunctionID: d.Number("func7"),
					StatCode:   d.String("stat7"),
				},
			},
		}

		records[record.Code] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Properties = records

	r.Logger.Infof("Loaded %d Property records", len(records))

	return nil
}
