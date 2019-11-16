package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var MiscItems map[string]*ItemCommonRecord

func LoadMiscItems(fileProvider d2interface.FileProvider) {
	MiscItems = *LoadCommonItems(fileProvider, d2resource.Misc, d2enum.InventoryItemTypeItem)
	log.Printf("Loaded %d misc items", len(MiscItems))
}
