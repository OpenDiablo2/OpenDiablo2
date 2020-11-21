package d2player

import (
	"fmt"
	"image"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
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

	hpLabelX = 15
	hpLabelY = 487

	manaLabelX = 785
	manaLabelY = 487
)

// static check that globeWidget implements Widget
var _ d2ui.Widget = &globeWidget{}

// static check that globeWidget implements ClickableWidget
var _ d2ui.ClickableWidget = &globeWidget{}

type globeFrame struct {
	sprite  *d2ui.Sprite
	offsetX int
	offsetY int
	idx     int
	gw      *globeWidget
}

func (gf *globeFrame) setFrameIndex() {
	if err := gf.sprite.SetCurrentFrame(gf.idx); err != nil {
		gf.gw.Error(err.Error())
	}
}

func (gf *globeFrame) setPosition(x, y int) {
	gf.sprite.SetPosition(x+gf.offsetX, y+gf.offsetY)
}

func newGlobeWidget(ui *d2ui.UIManager,
	asset *d2asset.AssetManager,
	x, y int,
	gtype globeType,
	value *int, valueMax *int,
	l d2util.LogLevel) *globeWidget {
	var globe, overlap *globeFrame

	var tooltipX, tooltipY int

	var tooltipTrans string

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
		tooltipX, tooltipY = hpLabelX, hpLabelY
		tooltipTrans = "panelhealth"
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
		tooltipX, tooltipY = manaLabelX, manaLabelY
		tooltipTrans = "panelmana"
	}

	gw := &globeWidget{
		BaseWidget:      base,
		asset:           asset,
		value:           value,
		valueMax:        valueMax,
		globe:           globe,
		overlap:         overlap,
		isTooltipLocked: false,
		tooltipX:        tooltipX,
		tooltipY:        tooltipY,
		tooltipTrans:    tooltipTrans,
	}

	gw.OnHoverStart(func() {
		if !gw.isTooltipLocked {
			gw.tooltip.SetVisible(true)
		}
	})

	gw.OnHoverEnd(func() {
		if !gw.isTooltipLocked {
			gw.tooltip.SetVisible(false)
		}
	})

	gw.Logger = d2util.NewLogger()
	gw.Logger.SetLevel(l)
	gw.Logger.SetPrefix(logPrefix)

	return gw
}

type globeWidget struct {
	*d2ui.BaseWidget
	asset    *d2asset.AssetManager
	value    *int
	valueMax *int
	globe    *globeFrame
	overlap  *globeFrame
	*d2util.Logger

	pressed         bool
	isTooltipLocked bool
	tooltip         *d2ui.Tooltip
	tooltipX        int
	tooltipY        int
	tooltipTrans    string
}

func (g *globeWidget) load() {
	var err error

	g.globe.sprite, err = g.GetManager().NewSprite(d2resource.HealthManaIndicator, d2resource.PaletteSky)
	if err != nil {
		g.Error(err.Error())
	}

	g.globe.setFrameIndex()

	g.overlap.sprite, err = g.GetManager().NewSprite(d2resource.GameGlobeOverlap, d2resource.PaletteSky)
	if err != nil {
		g.Error(err.Error())
	}

	g.overlap.setFrameIndex()

	// tooltip
	g.tooltip = g.GetManager().NewTooltip(d2resource.Font16, d2resource.PaletteUnits, d2ui.TooltipXLeft, d2ui.TooltipYTop)
	g.tooltip.SetPosition(g.tooltipX, g.tooltipY)
	g.tooltip.SetBoxEnabled(false)
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

// Contains is special here as the point of origin is at the lower left corner
// in contrast to any other element which is top left.
func (g *globeWidget) Contains(px, py int) bool {
	wx, wy := g.globe.sprite.GetPosition()
	width, height := g.globe.sprite.GetSize()

	return px >= wx && px <= wx+width && py <= wy && py >= wy-height
}

func (g *globeWidget) updateTooltip() {
	// Create and format string from string lookup table.
	fmtStr := g.asset.TranslateString(g.tooltipTrans)
	strPanel := fmt.Sprintf(fmtStr, *g.value, *g.valueMax)
	g.tooltip.SetText(strPanel)
}

func (g *globeWidget) Advance(elapsed float64) error {
	g.updateTooltip()

	return nil
}

func (g *globeWidget) Activate() {
	g.isTooltipLocked = !g.isTooltipLocked
	g.tooltip.SetVisible(g.isTooltipLocked)
}

func (g *globeWidget) GetEnabled() bool {
	return true
}

func (g *globeWidget) SetEnabled(enable bool) {
	// No-op
}

func (g *globeWidget) GetPressed() bool {
	return g.pressed
}

func (g *globeWidget) SetPressed(pressed bool) {
	g.pressed = pressed
}

func (g *globeWidget) OnActivated(callback func()) {
	// No-op
}
