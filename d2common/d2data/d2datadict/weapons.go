package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

var Weapons map[string]*ItemCommonRecord

func LoadWeapons(file []byte) {
	Weapons = LoadCommonItems(file, d2enum.InventoryItemTypeWeapon)
	log.Printf("Loaded %d weapons", len(Weapons))
}
