package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var Weapons map[string]*ItemCommonRecord

func LoadWeapons(fileProvider d2interface.FileProvider) {
	Weapons = *LoadCommonItems(fileProvider, d2resource.Weapons, d2enum.InventoryItemTypeWeapon)
	log.Printf("Loaded %d weapons", len(Weapons))
}
