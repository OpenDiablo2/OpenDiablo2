package d2player

import (
	"image"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type globeType = int

const (
	typeHealthGlobe globeType = iota
	typeManaGlobe
)

const (
	globeHeight = 80
	globeWidth  = 80

	globeSpriteOffsetX = 28
	globeSpriteOffsetY = -5

	healthStatusOffsetX = 30
	healthStatusOffsetY = -13

	manaStatusOffsetX = 7
	manaStatusOffsetY = -12

	manaGlobeScreenOffsetX = 117
)

// static check that globeWidget implements Widget
var _ d2ui.Widget = &globeWidget{}

type globeFrame struct {
	sprite  *d2ui.Sprite
	offsetX int
	offsetY int
	idx     int
}

func (gf *globeFrame) setFrameIndex() error {
	if err := gf.sprite.SetCurrentFrame(gf.idx); err != nil {
		return err
	}

	return nil
}

func (gf *globeFrame) setPosition(x, y int) {
	gf.sprite.SetPosition(x+gf.offsetX, y+gf.offsetY)
}

func (gf *globeFrame) getSize() (x, y int) {
	w, h := gf.sprite.GetSize()
	return w + gf.offsetX, h + gf.offsetY
}

type globeWidget struct {
	*d2ui.BaseWidget
	value    *int
	valueMax *int
	globe    *globeFrame
	overlap  *globeFrame
}

func newGlobeWidget(ui *d2ui.UIManager, x, y int, gtype globeType, value, valueMax *int) *globeWidget {
	var globe, overlap *globeFrame

	base := d2ui.NewBaseWidget(ui)
	base.SetPosition(x, y)

	if gtype == typeHealthGlobe {
		globe = &globeFrame{
			offsetX: healthStatusOffsetX,
			offsetY: healthStatusOffsetY,
			idx:     frameHealthStatus,
		}
		overlap = &globeFrame{
			offsetX: globeSpriteOffsetX,
			offsetY: globeSpriteOffsetY,
			idx:     frameHealthStatus,
		}
	} else if gtype == typeManaGlobe {
		globe = &globeFrame{
			offsetX: manaStatusOffsetX,
			offsetY: manaStatusOffsetY,
			idx:     frameManaStatus,
		}
		overlap = &globeFrame{
			offsetX: rightGlobeOffsetX,
			offsetY: rightGlobeOffsetY,
			idx:     frameRightGlobe,
		}
	}

	return &globeWidget{
		BaseWidget: base,
		value:      value,
		valueMax:   valueMax,
		globe:      globe,
		overlap:    overlap,
	}
}

func (g *globeWidget) load() error {
	var err error

	g.globe.sprite, err = g.GetManager().NewSprite(d2resource.HealthManaIndicator, d2resource.PaletteSky)
	if err != nil {
		return err
	}

	err = g.globe.setFrameIndex()
	if err != nil {
		return err
	}

	g.overlap.sprite, err = g.GetManager().NewSprite(d2resource.GameGlobeOverlap, d2resource.PaletteSky)
	if err != nil {
		return err
	}

	err = g.overlap.setFrameIndex()
	if err != nil {
		return err
	}

	return nil
}

// Render draws the widget to the screen
func (g *globeWidget) Render(target d2interface.Surface) {
	valuePercent := float64(*g.value) / float64(*g.valueMax)
	barHeight := int(valuePercent * float64(globeHeight))

	maskRect := image.Rect(0, globeHeight-barHeight, globeWidth, globeHeight)

	g.globe.setPosition(g.GetPosition())
	g.globe.sprite.RenderSection(target, maskRect)

	g.overlap.setPosition(g.GetPosition())
	g.overlap.sprite.Render(target)
}

func (g *globeWidget) GetSize() (x, y int) {
	return g.overlap.getSize()
}

func (g *globeWidget) Advance(elapsed float64) error {
	return nil
}
