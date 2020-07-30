package d2mapentity

import (
	"errors"
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math/d2vector"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2item/diablo2item"
)

const (
	errInvalidItemCodes = "invalid item codes supplied"
)

type Item struct {
	*AnimatedEntity
	Item *diablo2item.Item
}

func (i *Item) GetPosition() d2vector.Position {
	return i.AnimatedEntity.Position
}

func (i *Item) GetVelocity() d2vector.Vector {
	return i.AnimatedEntity.velocity
}

func CreateItem(x, y int, codes ...string) (*Item, error) {
	item := diablo2item.NewItem(codes...)

	if item == nil {
		return nil, errors.New(errInvalidItemCodes)
	}

	animation, err := d2asset.LoadAnimation(
		fmt.Sprintf("%s/%s.DC6", d2resource.ItemGraphics, item.CommonRecord().FlippyFile),
		d2resource.PaletteUnits,
	)

	if err != nil {
		return nil, err
	}

	animation.PlayForward()
	animation.SetPlayLoop(false)
	entity := CreateAnimatedEntity(x*5, y*5, animation)

	result := &Item{
		AnimatedEntity: entity,
		Item:           item,
	}

	return result, nil
}
