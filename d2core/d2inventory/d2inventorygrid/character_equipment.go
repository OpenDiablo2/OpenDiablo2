package d2inventorygrid

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

func NewCharacterEquipmentGrid(hero d2enum.Hero) *CharacterEquipment {
	return &CharacterEquipment{

	}
}

type CharacterEquipment struct{
	head InventoryGrid
	neck InventoryGrid
	torso InventoryGrid
	wieldLeft InventoryGrid
	wieldRight InventoryGrid
	wieldLeftAlt InventoryGrid
	wieldRightAlt InventoryGrid
	ringLeft InventoryGrid
	ringRight InventoryGrid
}

func (ce *CharacterEquipment) Head() *InventoryGrid {
	panic("implement me")
}

func (ce *CharacterEquipment) Neck() *InventoryGrid {
	panic("implement me")
}

func (ce *CharacterEquipment) Torso() *InventoryGrid {
	panic("implement me")
}

func (ce *CharacterEquipment) WieldLeft() *InventoryGrid {
	panic("implement me")
}

func (ce *CharacterEquipment) WieldRight() *InventoryGrid {
	panic("implement me")
}

func (ce *CharacterEquipment) WieldLeftAlt() *InventoryGrid {
	panic("implement me")
}

func (ce *CharacterEquipment) WieldRightAlt() *InventoryGrid {
	panic("implement me")
}

func (ce *CharacterEquipment) RingLeft() *InventoryGrid {
	panic("implement me")
}

func (ce *CharacterEquipment) RingRight() *InventoryGrid {
	panic("implement me")
}

func (ce *CharacterEquipment) Hands() *InventoryGrid {
	panic("implement me")
}

func (ce *CharacterEquipment) Feet() *InventoryGrid {
	panic("implement me")
}

func (ce *CharacterEquipment) Items() []*GridItem {
	panic("implement me")
}

