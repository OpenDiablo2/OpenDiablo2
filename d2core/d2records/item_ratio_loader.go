package d2records

import (
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadItemRatios loads all of the ItemRatioRecords from ItemRatio.txt
func itemRatioLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(map[string]*ItemRatioRecord)

	for d.Next() {
		record := &ItemRatioRecord{
			Function:      d.String("Function"),
			Version:       d.Bool("Version"),
			Uber:          d.Bool("Uber"),
			ClassSpecific: d.Bool("Class Specific"),
			UniqueDropInfo: DropRatioInfo{
				Frequency:  d.Number("Unique"),
				Divisor:    d.Number("UniqueDivisor"),
				DivisorMin: d.Number("UniqueMin"),
			},
			RareDropInfo: DropRatioInfo{
				Frequency:  d.Number("Rare"),
				Divisor:    d.Number("RareDivisor"),
				DivisorMin: d.Number("RareMin"),
			},
			SetDropInfo: DropRatioInfo{
				Frequency:  d.Number("Set"),
				Divisor:    d.Number("SetDivisor"),
				DivisorMin: d.Number("SetMin"),
			},
			MagicDropInfo: DropRatioInfo{
				Frequency:  d.Number("Magic"),
				Divisor:    d.Number("MagicDivisor"),
				DivisorMin: d.Number("MagicMin"),
			},
			HiQualityDropInfo: DropRatioInfo{
				Frequency:  d.Number("HiQuality"),
				Divisor:    d.Number("HiQualityDivisor"),
				DivisorMin: 0,
			},
			NormalDropInfo: DropRatioInfo{
				Frequency:  d.Number("Normal"),
				Divisor:    d.Number("NormalDivisor"),
				DivisorMin: 0,
			},
		}

		records[record.Function+strconv.FormatBool(record.Version)] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d ItemRatio records", len(records))

	r.Item.Ratios = records

	return nil
}
