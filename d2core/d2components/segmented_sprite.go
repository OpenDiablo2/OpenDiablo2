//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that SegmentedSpriteComponent implements Component
var _ akara.Component = &SegmentedSpriteComponent{}

// static check that SegmentedSpriteMap implements ComponentMap
var _ akara.ComponentMap = &SegmentedSpriteMap{}

// SegmentedSpriteComponent represents an entities x,y axis scale as a vector
type SegmentedSpriteComponent struct {
	*akara.BaseComponent
	Xsegments   int
	Ysegments   int
	FrameOffset int
}

// SegmentedSpriteMap is a map of entity ID's to SegmentedSprite
type SegmentedSpriteMap struct {
	*akara.BaseComponentMap
}

// AddSegmentedSprite adds a new SegmentedSpriteComponent for the given entity id and returns it.
// this is a convenience method for the generic Add method, as it returns a
// *SegmentedSpriteComponent instead of an akara.Component
func (cm *SegmentedSpriteMap) AddSegmentedSprite(id akara.EID) *SegmentedSpriteComponent {
	c := cm.Add(id).(*SegmentedSpriteComponent)

	c.Xsegments = 1
	c.Ysegments = 1

	return c
}

// GetSegmentedSprite returns the SegmentedSpriteComponent associated with the given entity id
func (cm *SegmentedSpriteMap) GetSegmentedSprite(id akara.EID) (*SegmentedSpriteComponent, bool) {
	entry, found := cm.Get(id)
	if entry == nil {
		return nil, false
	}

	return entry.(*SegmentedSpriteComponent), found
}

// SegmentedSprite is a convenient reference to be used as a component identifier
var SegmentedSprite = newSegmentedSprite() // nolint:gochecknoglobals // global by design

func newSegmentedSprite() akara.Component {
	return &SegmentedSpriteComponent{
		BaseComponent: akara.NewBaseComponent(SegmentedSpriteCID, newSegmentedSprite, newSegmentedSpriteMap),
	}
}

func newSegmentedSpriteMap() akara.ComponentMap {
	baseComponent := akara.NewBaseComponent(SegmentedSpriteCID, newSegmentedSprite, newSegmentedSpriteMap)
	baseMap := akara.NewBaseComponentMap(baseComponent)

	cm := &SegmentedSpriteMap{
		BaseComponentMap: baseMap,
	}

	return cm
}
