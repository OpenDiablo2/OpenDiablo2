package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func storePagesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(StorePages)

	for d.Next() {
		record := &StorePageRecord{
			StorePage: d.String("Store Page"),
			Code:      d.String("Code"),
		}
		records[record.StorePage] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Item.StorePages = records

	r.Logger.Infof("Loaded %d StorePage records", len(records))

	return nil
}
