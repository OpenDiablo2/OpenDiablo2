package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var Armors map[string]*ItemCommonRecord

func LoadArmors(fileProvider d2interface.FileProvider) {
	Armors = *LoadCommonItems(fileProvider, d2resource.Armor, d2enum.InventoryItemTypeArmor)
	log.Printf("Loaded %d armors", len(Armors))
}
