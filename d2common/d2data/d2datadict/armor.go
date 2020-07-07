package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// Armors stores all of the ArmorRecords
//nolint:gochecknoglobals // Currently global by design, only written once
var Armors map[string]*ItemCommonRecord

// LoadArmors loads entries from armor.txt as ItemCommonRecords
func LoadArmors(file []byte) {
	Armors = LoadCommonItems(file, d2enum.InventoryItemTypeArmor)
	log.Printf("Loaded %d armors", len(Armors))
}
