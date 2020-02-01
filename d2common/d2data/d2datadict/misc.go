package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

var MiscItems map[string]*ItemCommonRecord

func LoadMiscItems(file []byte) {
	MiscItems = *LoadCommonItems(file, d2enum.InventoryItemTypeItem)
	log.Printf("Loaded %d misc items", len(MiscItems))
}
