//nolint:dupl,golint,stylecheck // component declarations are supposed to look the same
package d2components

import (
	"github.com/gravestench/akara"
)

// static check that SegmentedSprite implements Component
var _ akara.Component = &SegmentedSprite{}

// SegmentedSprite represents the segmentation of a dcc or dc6 sprite.
// For example, the title screen background is a single image with 12 frames.
// The segmentation can be described as 4 x segments, 3 y segments, with a frame offset of 0.
type SegmentedSprite struct {
	Xsegments   int
	Ysegments   int
	FrameOffset int
}

// New creates a new SegmentedSprite component. By default, x and y segments are both set to 1.
func (*SegmentedSprite) New() akara.Component {
	return &SegmentedSprite{
		Xsegments: 1,
		Ysegments: 1,
	}
}

// SegmentedSpriteFactory is a wrapper for the generic component factory that returns SegmentedSprite component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a SegmentedSprite.
type SegmentedSpriteFactory struct {
	*akara.ComponentFactory
}

// Add adds a SegmentedSprite component to the given entity and returns it
func (m *SegmentedSpriteFactory) Add(id akara.EID) *SegmentedSprite {
	return m.ComponentFactory.Add(id).(*SegmentedSprite)
}

// Get returns the SegmentedSprite component for the given entity, and a bool for whether or not it exists
func (m *SegmentedSpriteFactory) Get(id akara.EID) (*SegmentedSprite, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*SegmentedSprite), found
}
