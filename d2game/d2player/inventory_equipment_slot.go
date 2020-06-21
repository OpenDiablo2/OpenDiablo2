package d2player

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
	"log"
)

type EquipmentSlot struct {
	item   InventoryItem
	x      int
	y      int
	sprite *d2ui.Sprite
}

func (e *EquipmentSlot) Load(item InventoryItem) {
	var itemSprite *d2ui.Sprite
	// TODO: Put the pattern into D2Shared
	animation, err := d2asset.LoadAnimation(
		fmt.Sprintf("/data/global/items/inv%s.dc6", item.GetItemCode()),
		d2resource.PaletteSky,
	)
	if err != nil {
		log.Printf("failed to load sprite for item (%s): %v", item.GetItemCode(), err)
	}
	itemSprite, err = d2ui.LoadSprite(animation)
	e.sprite = itemSprite
}
