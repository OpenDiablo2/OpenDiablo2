package d2datadict

import (
	"log"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

//ItemRatioRecord encapsulates information found in ItemRatio.txt
//The information has been gathered from [https://d2mods.info/forum/kb/viewarticle?a=387]
type ItemRatioRecord struct {
	Function string
	// 0 for classic, 1 for LoD
	Version bool

	// 0 for normal, 1 for exceptional
	Uber          bool
	ClassSpecific bool

	// All following fields are used in item drop calculation
	Unique           int
	UniqueDivisor    int
	UniqueMin        int
	Rare             int
	RareDivisor      int
	RareMin          int
	Set              int
	SetDivisor       int
	SetMin           int
	Magic            int
	MagicDivisor     int
	MagicMin         int
	HiQuality        int
	HiQualityDivisor int
	Normal           int
	NormalDivisor    int
}

var ItemRatios map[string]*ItemRatioRecord

func LoadItemRatios(file []byte) {
	ItemRatios = make(map[string]*ItemRatioRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &ItemRatioRecord{
			Function:         d.String("Function"),
			Version:          d.Bool("Version"),
			Uber:             d.Bool("Uber"),
			ClassSpecific:    d.Bool("Class Specific"),
			Unique:           d.Number("Unique"),
			UniqueDivisor:    d.Number("UniqueDivisor"),
			UniqueMin:        d.Number("UniqueMin"),
			Rare:             d.Number("Rare"),
			RareDivisor:      d.Number("RareDivisor"),
			RareMin:          d.Number("RareMin"),
			Set:              d.Number("Set"),
			SetDivisor:       d.Number("SetDivisor"),
			SetMin:           d.Number("SetMin"),
			Magic:            d.Number("Magic"),
			MagicMin:         d.Number("MagicMin"),
			MagicDivisor:     d.Number("MagicDivisor"),
			HiQuality:        d.Number("HiQuality"),
			HiQualityDivisor: d.Number("HiQualityDivisor"),
			Normal:           d.Number("Normal"),
			NormalDivisor:    d.Number("NormalDivisor"),
		}
		ItemRatios[record.Function+strconv.FormatBool(record.Version)] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d ItemRatio records", len(ItemRatios))
}
