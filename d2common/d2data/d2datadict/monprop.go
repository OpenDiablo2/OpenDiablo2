package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// MonPropRecord is a representation of a single row of monprop.txt
type MonPropRecord struct {
	ID string

	Prop1       string
	Chance1     int
	Parameter1  string
	Min1        int
	Max1        int
	Prop2       string
	Chance2     int
	Parameter2  string
	Min2        int
	Max2        int
	Prop3       string
	Chance3     int
	Parameter3  string
	Min3        int
	Max3        int
	Prop4       string
	Chance4     int
	Parameter4  string
	Min4        int
	Max4        int
	Prop5       string
	Chance5     int
	Parameter5  string
	Min5        int
	Max5        int
	Prop6       string
	Chance6     int
	Parameter6  string
	Min6        int
	Max6        int
	Prop1N      string
	Chance1N    int
	Parameter1N string
	Min1N       int
	Max1N       int
	Prop2N      string
	Chance2N    int
	Parameter2N string
	Min2N       int
	Max2N       int
	Prop3N      string
	Chance3N    int
	Parameter3N string
	Min3N       int
	Max3N       int
	Prop4N      string
	Chance4N    int
	Parameter4N string
	Min4N       int
	Max4N       int
	Prop5N      string
	Chance5N    int
	Parameter5N string
	Min5N       int
	Max5N       int
	Prop6N      string
	Chance6N    int
	Parameter6N string
	Min6N       int
	Max6N       int
	Prop1H      string
	Chance1H    int
	Parameter1H string
	Min1H       int
	Max1H       int
	Prop2H      string
	Chance2H    int
	Parameter2H string
	Min2H       int
	Max2H       int
	Prop3H      string
	Chance3H    int
	Parameter3H string
	Min3H       int
	Max3H       int
	Prop4H      string
	Chance4H    int
	Parameter4H string
	Min4H       int
	Max4H       int
	Prop5H      string
	Chance5H    int
	Parameter5H string
	Min5H       int
	Max5H       int
	Prop6H      string
	Chance6H    int
	Parameter6H string
	Min6H       int
	Max6H       int

	// EOL int // unused
}

// MonProps stores all of the MonPropRecords
var MonProps map[string]*MonPropRecord //nolint:gochecknoglobals // Currently global by design, only written once

// LoadMonProps loads Monprop records into a map[string]*MonPropRecord
func LoadMonProps(file []byte) {
	MonProps = make(map[string]*MonPropRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &MonPropRecord{
			ID: d.String("Id"),

			Prop1:       d.String("prop1"),
			Chance1:     d.Number("chance1"),
			Parameter1:  d.String("par1"),
			Min1:        d.Number("min1"),
			Max1:        d.Number("max1"),
			Prop2:       d.String("prop2"),
			Chance2:     d.Number("chance2"),
			Parameter2:  d.String("par2"),
			Min2:        d.Number("min2"),
			Max2:        d.Number("max2"),
			Prop3:       d.String("prop3"),
			Chance3:     d.Number("chance3"),
			Parameter3:  d.String("par3"),
			Min3:        d.Number("min3"),
			Max3:        d.Number("max3"),
			Prop4:       d.String("prop4"),
			Chance4:     d.Number("chance4"),
			Parameter4:  d.String("par4"),
			Min4:        d.Number("min4"),
			Max4:        d.Number("max4"),
			Prop5:       d.String("prop5"),
			Chance5:     d.Number("chance5"),
			Parameter5:  d.String("par5"),
			Min5:        d.Number("min5"),
			Max5:        d.Number("max5"),
			Prop6:       d.String("prop6"),
			Chance6:     d.Number("chance6"),
			Parameter6:  d.String("par6"),
			Min6:        d.Number("min6"),
			Max6:        d.Number("max6"),
			Prop1N:      d.String("prop1 (N)"),
			Chance1N:    d.Number("chance1 (N)"),
			Parameter1N: d.String("par1 (N)"),
			Min1N:       d.Number("min1 (N)"),
			Max1N:       d.Number("max1 (N)"),
			Prop2N:      d.String("prop2 (N)"),
			Chance2N:    d.Number("chance2 (N)"),
			Parameter2N: d.String("par2 (N)"),
			Min2N:       d.Number("min2 (N)"),
			Max2N:       d.Number("max2 (N)"),
			Prop3N:      d.String("prop3 (N)"),
			Chance3N:    d.Number("chance3 (N)"),
			Parameter3N: d.String("par3 (N)"),
			Min3N:       d.Number("min3 (N)"),
			Max3N:       d.Number("max3 (N)"),
			Prop4N:      d.String("prop4 (N)"),
			Chance4N:    d.Number("chance4 (N)"),
			Parameter4N: d.String("par4 (N)"),
			Min4N:       d.Number("min4 (N)"),
			Max4N:       d.Number("max4 (N)"),
			Prop5N:      d.String("prop5 (N)"),
			Chance5N:    d.Number("chance5 (N)"),
			Parameter5N: d.String("par5 (N)"),
			Min5N:       d.Number("min5 (N)"),
			Max5N:       d.Number("max5 (N)"),
			Prop6N:      d.String("prop6 (N)"),
			Chance6N:    d.Number("chance6 (N)"),
			Parameter6N: d.String("par6 (N)"),
			Min6N:       d.Number("min6 (N)"),
			Max6N:       d.Number("max6 (N)"),
			Prop1H:      d.String("prop1 (H)"),
			Chance1H:    d.Number("chance1 (H)"),
			Parameter1H: d.String("par1 (H)"),
			Min1H:       d.Number("min1 (H)"),
			Max1H:       d.Number("max1 (H)"),
			Prop2H:      d.String("prop2 (H)"),
			Chance2H:    d.Number("chance2 (H)"),
			Parameter2H: d.String("par2 (H)"),
			Min2H:       d.Number("min2 (H)"),
			Max2H:       d.Number("max2 (H)"),
			Prop3H:      d.String("prop3 (H)"),
			Chance3H:    d.Number("chance3 (H)"),
			Parameter3H: d.String("par3 (H)"),
			Min3H:       d.Number("min3 (H)"),
			Max3H:       d.Number("max3 (H)"),
			Prop4H:      d.String("prop4 (H)"),
			Chance4H:    d.Number("chance4 (H)"),
			Parameter4H: d.String("par4 (H)"),
			Min4H:       d.Number("min4 (H)"),
			Max4H:       d.Number("max4 (H)"),
			Prop5H:      d.String("prop5 (H)"),
			Chance5H:    d.Number("chance5 (H)"),
			Parameter5H: d.String("par5 (H)"),
			Min5H:       d.Number("min5 (H)"),
			Max5H:       d.Number("max5 (H)"),
			Prop6H:      d.String("prop6 (H)"),
			Chance6H:    d.Number("chance6 (H)"),
			Parameter6H: d.String("par6 (H)"),
			Min6H:       d.Number("min6 (H)"),
			Max6H:       d.Number("max6 (H)"),
		}
		MonProps[record.ID] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonProp records", len(MonProps))
}
