package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type PropertyStatRecord struct {
	SetID      int
	Value      int
	FunctionID int
	StatCode   string
}

// PropertyRecord is a representation of a single row of properties.txt
type PropertyRecord struct {
	Code   string
	Active string
	Stats  [7]*PropertyStatRecord
}

// Properties stores all of the PropertyRecords
var Properties map[string]*PropertyRecord //nolint:gochecknoglobals // Currently global by design, 
// only written once

// LoadProperties loads gem records into a map[string]*PropertiesRecord
func LoadProperties(file []byte) {
	Properties = make(map[string]*PropertyRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		prop := &PropertyRecord{
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
		Properties[prop.Code] = prop
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Property records", len(Properties))
}
