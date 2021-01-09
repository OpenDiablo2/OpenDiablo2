package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadItemStatCosts loads ItemStatCostRecord's from text
func itemStatCostLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(ItemStatCosts)

	for d.Next() {
		record := &ItemStatCostRecord{
			Name:  d.String("Stat"),
			Index: d.Number("ID"),

			Signed:   d.Number("Signed") > 0,
			KeepZero: d.Number("keepzero") > 0,

			// Ranged:  d.Number("Ranged") > 0,
			MinAccr: d.Number("MinAccr"),

			UpdateAnimRate: d.Number("UpdateAnimRate") > 0,

			SendOther: d.Number("Send Other") > 0,
			SendBits:  d.Number("Send Bits"),
			SendParam: d.Number("Send Param Bits"),

			Saved:         d.Number("CSvBits") > 0,
			SavedSigned:   d.Number("CSvSigned") > 0,
			SavedBits:     d.Number("CSvBits"),
			SaveBits:      d.Number("Save Bits"),
			SaveAdd:       d.Number("Save Add"),
			SaveParamBits: d.Number("Save Param Bits"),

			Encode: d2enum.EncodingType(d.Number("Encode")),

			CallbackEnabled: d.Number("fCallback") > 0,

			CostAdd:      d.Number("Add"),
			CostMultiply: d.Number("Multiply"),
			ValShift:     d.Number("ValShift"),

			OperatorType: d2enum.OperatorType(d.Number("op")),
			OpParam:      d.Number("op param"),
			OpBase:       d.String("op base"),
			OpStat1:      d.String("op stat1"),
			OpStat2:      d.String("op stat2"),
			OpStat3:      d.String("op stat3"),

			Direct:  d.Number("direct") > 0,
			MaxStat: d.String("maxstat"),

			ItemSpecific:  d.Number("itemspecific") > 0,
			DamageRelated: d.Number("damagerelated") > 0,

			EventID1:     d2enum.GetItemEventType(d.String("itemevent1")),
			EventID2:     d2enum.GetItemEventType(d.String("itemevent2")),
			EventFuncID1: d2enum.ItemEventFuncID(d.Number("itemeventfunc1")),
			EventFuncID2: d2enum.ItemEventFuncID(d.Number("itemeventfunc2")),

			DescPriority: d.Number("descpriority"),
			DescFnID:     d.Number("descfunc"),
			// DescVal:      d.Number("descval"), // needs special handling
			DescStrPos: d.String("descstrpos"),
			DescStrNeg: d.String("descstrneg"),
			DescStr2:   d.String("descstr2"),

			DescGroup:       d.Number("dgrp"),
			DescGroupFuncID: d.Number("dgrpfunc"),

			DescGroupVal:    d.Number("dgrpval"),
			DescGroupStrPos: d.String("dgrpstrpos"),
			DescGroupStrNeg: d.String("dgrpstrneg"),
			DescGroupStr2:   d.String("dgrpstr2"),

			Stuff: d.String("stuff"),
		}

		descValStr := d.String("descval")
		switch descValStr {
		case "2":
			record.DescVal = 2
		case "0":
			record.DescVal = 0
		default:
			// handle empty fields, seems like they should have been 1
			record.DescVal = 1
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d ItemStatCost records", len(records))

	r.Item.Stats = records

	return nil
}
