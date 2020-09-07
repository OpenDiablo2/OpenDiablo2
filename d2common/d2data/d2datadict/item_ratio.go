package d2datadict

import (
	"log"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// A helper type for item drop calculation
type dropRatioInfo struct {
	frequency  int
	divisor    int
	divisorMin int
}

// ItemRatioRecord encapsulates information found in ItemRatio.txt, it specifies drop ratios
// for various types of items
// The information has been gathered from [https://d2mods.info/forum/kb/viewarticle?a=387]
type ItemRatioRecord struct {
	Function string
	// 0 for classic, 1 for LoD
	Version bool

	// 0 for normal, 1 for exceptional
	Uber          bool
	ClassSpecific bool

	// All following fields are used in item drop calculation
	UniqueDropInfo    dropRatioInfo
	RareDropInfo      dropRatioInfo
	SetDropInfo       dropRatioInfo
	MagicDropInfo     dropRatioInfo
	HiQualityDropInfo dropRatioInfo
	NormalDropInfo    dropRatioInfo
}

// ItemRatios holds all of the ItemRatioRecords from ItemRatio.txt
var ItemRatios map[string]*ItemRatioRecord //nolint:gochecknoglobals // Currently global by design

// LoadItemRatios loads all of the ItemRatioRecords from ItemRatio.txt
func LoadItemRatios(file []byte) {
	ItemRatios = make(map[string]*ItemRatioRecord)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		record := &ItemRatioRecord{
			Function:      d.String("Function"),
			Version:       d.Bool("Version"),
			Uber:          d.Bool("Uber"),
			ClassSpecific: d.Bool("Class Specific"),
			UniqueDropInfo: dropRatioInfo{
				frequency:  d.Number("Unique"),
				divisor:    d.Number("UniqueDivisor"),
				divisorMin: d.Number("UniqueMin"),
			},
			RareDropInfo: dropRatioInfo{
				frequency:  d.Number("Rare"),
				divisor:    d.Number("RareDivisor"),
				divisorMin: d.Number("RareMin"),
			},
			SetDropInfo: dropRatioInfo{
				frequency:  d.Number("Set"),
				divisor:    d.Number("SetDivisor"),
				divisorMin: d.Number("SetMin"),
			},
			MagicDropInfo: dropRatioInfo{
				frequency:  d.Number("Magic"),
				divisor:    d.Number("MagicDivisor"),
				divisorMin: d.Number("MagicMin"),
			},
			HiQualityDropInfo: dropRatioInfo{
				frequency:  d.Number("HiQuality"),
				divisor:    d.Number("HiQualityDivisor"),
				divisorMin: 0,
			},
			NormalDropInfo: dropRatioInfo{
				frequency:  d.Number("Normal"),
				divisor:    d.Number("NormalDivisor"),
				divisorMin: 0,
			},
		}
		ItemRatios[record.Function+strconv.FormatBool(record.Version)] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d ItemRatio records", len(ItemRatios))
}
