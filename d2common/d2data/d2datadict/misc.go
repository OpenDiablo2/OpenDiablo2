package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// MiscItems stores all of the ItemCommonRecords for misc.txt
var MiscItems map[string]*ItemCommonRecord //nolint:gochecknoglobals // Currently global by design

// LoadMiscItems loads ItemCommonRecords from misc.txt
func LoadMiscItems(file []byte) {
	MiscItems = LoadCommonItems(file, d2enum.InventoryItemTypeItem)
	log.Printf("Loaded %d misc items", len(MiscItems))
}
