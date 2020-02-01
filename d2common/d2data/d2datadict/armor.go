package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

var Armors map[string]*ItemCommonRecord

func LoadArmors(file []byte) {
	Armors = *LoadCommonItems(file, d2enum.InventoryItemTypeArmor)
	log.Printf("Loaded %d armors", len(Armors))
}
