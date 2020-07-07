package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// Weapons stores all of the WeaponRecords
var Weapons map[string]*ItemCommonRecord //nolint:gochecknoglobals // Currently global by design, only written once

// LoadWeapons loads weapon records
func LoadWeapons(file []byte) {
	Weapons = LoadCommonItems(file, d2enum.InventoryItemTypeWeapon)
	log.Printf("Loaded %d weapons", len(Weapons))
}
