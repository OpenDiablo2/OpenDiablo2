package d2gui

import (
	"log"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// LayoutScrollbar is a scrollbar that can be used with any layout
// and attaches to a main layout. You need to use a wrapper for your content
// as main layout in order for the scrollbar to work properly
type LayoutScrollbar struct {
	sliderSprites []*d2ui.Sprite
	gutterSprites []*d2ui.Sprite

	parentLayout *Layout
	targetLayout *Layout
	sliderLayout *Layout

	arrowUpLayout   *Layout
	arrowDownLayout *Layout

	arrowUpSprite   *d2ui.Sprite
	arrowDownSprite *d2ui.Sprite

	maxY                   int
	minY                   int
	arrowClickSliderOffset int
	viewportSize           int
	contentSize            int

	clickedAtY         int
	mouseYOnSlider     int
	lastY              int
	gutterHeight       int
	sliderHeight       int
	contentToViewRatio float32

	// isVisible bool
	arrowUpClicked   bool
	arrowDownClicked bool
	sliderClicked    bool
}

const (
	textSliderPartWidth  = 12
	textSliderPartHeight = 13

	arrrowClickContentOffsetPercentage = 0.02
	oneHundredPercent                  = 1.0
)

const (
	textSliderPartArrowDownHollow int = iota + 8
	textSliderPartArrowUpHollow
	textSliderPartArrowDownFilled int = 10
	textSliderPartArrowUpFilled   int = 11
	// textSliderPartSquare
	textSliderPartInnerGutter       int = 13
	textSliderPartFillingVariation1 int = 14
)

// NewLayoutScrollbar attaches a scrollbar to the parentLayout to control the targetLayout
func NewLayoutScrollbar(
	parentLayout *Layout,
	targetLayout *Layout,
) *LayoutScrollbar {
	parentW, parentH := parentLayout.GetSize()
	_, targetH := targetLayout.GetSize()
	gutterHeight := parentH - (2 * textSliderPartHeight)
	viewportPercentage := oneHundredPercent - (float32(targetH-parentH) / float32(targetH))
	sliderHeight := int(float32(gutterHeight) * viewportPercentage)
	x, y := parentW-textSliderPartWidth, 0

	ret := &LayoutScrollbar{
		sliderSprites: []*d2ui.Sprite{},
		gutterSprites: []*d2ui.Sprite{},
	}

	ret.contentToViewRatio = viewportPercentage
	ret.contentToViewRatio = float32(targetH) / float32(gutterHeight)
	ret.gutterHeight = gutterHeight
	ret.sliderHeight = sliderHeight
	ret.minY = y + textSliderPartHeight
	ret.maxY = (y + parentH) - (textSliderPartHeight + sliderHeight)
	ret.contentSize = targetH
	ret.arrowClickSliderOffset = int(float32(sliderHeight) * arrrowClickContentOffsetPercentage)
	ret.viewportSize = parentH

	arrowUpLayout := parentLayout.AddLayout(PositionTypeAbsolute)
	arrowUpLayout.SetSize(textSliderPartWidth, textSliderPartHeight)
	arrowUpLayout.SetPosition(x, 0)
	ret.arrowUpLayout = arrowUpLayout

	gutterLayout := parentLayout.AddLayout(PositionTypeAbsolute)
	gutterLayout.SetSize(textSliderPartWidth, gutterHeight)
	gutterLayout.SetPosition(x, textSliderPartHeight)

	sliderLayout := parentLayout.AddLayout(PositionTypeAbsolute)
	sliderLayout.SetPosition(x, textSliderPartHeight)
	sliderLayout.SetSize(textSliderPartWidth, sliderHeight)
	sliderLayout.SetMouseClickHandler(ret.OnSliderMouseClick)

	arrowDownLayout := parentLayout.AddLayout(PositionTypeAbsolute)
	arrowDownLayout.SetSize(textSliderPartWidth, textSliderPartHeight)
	arrowDownLayout.SetPosition(x, textSliderPartHeight+gutterHeight)
	ret.arrowDownLayout = arrowDownLayout

	ret.sliderLayout = sliderLayout
	ret.parentLayout = parentLayout
	ret.targetLayout = targetLayout

	ret.parentLayout.AdjustEntryPlacement()
	ret.targetLayout.AdjustEntryPlacement()
	ret.sliderLayout.AdjustEntryPlacement()

	return ret
}

// Load sets the scrollbar layouts and loads the sprites
func (scrollbar *LayoutScrollbar) Load(ui *d2ui.UIManager) error {
	arrowUpX, arrowUpY := scrollbar.arrowUpLayout.ScreenPos()
	arrowUpSprite, _ := ui.NewSprite(d2resource.TextSlider, d2resource.PaletteSky)

	if err := arrowUpSprite.SetCurrentFrame(textSliderPartArrowUpFilled); err != nil {
		return err
	}

	arrowUpSprite.SetPosition(arrowUpX, arrowUpY+textSliderPartHeight)
	scrollbar.arrowUpSprite = arrowUpSprite

	arrowDownX, arrowDownY := scrollbar.arrowDownLayout.ScreenPos()
	arrowDownSprite, _ := ui.NewSprite(d2resource.TextSlider, d2resource.PaletteSky)

	if err := arrowDownSprite.SetCurrentFrame(textSliderPartArrowDownFilled); err != nil {
		return err
	}

	arrowDownSprite.SetPosition(arrowDownX, arrowDownY+textSliderPartHeight)
	scrollbar.arrowDownSprite = arrowDownSprite

	gutterParts := int(math.Ceil(float64(scrollbar.gutterHeight+(2*textSliderPartHeight)) / float64(textSliderPartHeight)))
	sliderParts := int(math.Ceil(float64(scrollbar.sliderHeight) / float64(textSliderPartHeight)))
	gutterX, gutterY := arrowUpX, arrowUpY+(2*textSliderPartHeight)-1
	i := 0

	for {
		if i >= gutterParts {
			break
		}

		f, _ := ui.NewSprite(d2resource.TextSlider, d2resource.PaletteSky)

		if err := f.SetCurrentFrame(textSliderPartInnerGutter); err != nil {
			return err
		}

		newY := gutterY + (i * (textSliderPartHeight - 1))
		f.SetPosition(gutterX, newY)

		scrollbar.gutterSprites = append(scrollbar.gutterSprites, f)

		i++
	}

	i = 0

	for {
		if i >= sliderParts {
			break
		}

		f, _ := ui.NewSprite(d2resource.TextSlider, d2resource.PaletteSky)

		if err := f.SetCurrentFrame(textSliderPartFillingVariation1); err != nil {
			return err
		}

		scrollbar.sliderSprites = append(scrollbar.sliderSprites, f)

		i++
	}

	scrollbar.updateSliderSpritesPosition()

	return nil
}

func (scrollbar *LayoutScrollbar) updateSliderSpritesPosition() {
	scrollbar.sliderLayout.AdjustEntryPlacement()
	sliderLayoutX, sliderLayoutY := scrollbar.sliderLayout.ScreenPos()

	for i, s := range scrollbar.sliderSprites {
		newY := sliderLayoutY + (i * (textSliderPartHeight - 1)) + textSliderPartHeight
		s.SetPosition(sliderLayoutX-1, newY)
	}
}

// OnSliderMouseClick affects the state of the slider
func (scrollbar *LayoutScrollbar) OnSliderMouseClick(event d2interface.MouseEvent) {
	scrollbar.clickedAtY = event.Y()
	scrollbar.lastY = scrollbar.clickedAtY
	scrollbar.mouseYOnSlider = event.Y() - scrollbar.sliderLayout.Sy
}

func (scrollbar *LayoutScrollbar) moveScaledContentBy(offset int) int {
	_, y := scrollbar.sliderLayout.GetPosition()
	newY := y + offset

	outOfBoundsUp := false
	outOfBoundsDown := false

	if newY > scrollbar.maxY {
		newY = scrollbar.maxY
		outOfBoundsDown = true
	}

	if newY < scrollbar.minY {
		newY = scrollbar.minY
		outOfBoundsUp = true
	}

	if !outOfBoundsUp && !outOfBoundsDown {
		scrollbar.clickedAtY += offset

		if scrollbar.targetLayout != nil {
			contentX, contentY := scrollbar.targetLayout.GetPosition()
			scaledOffset := int(math.Round(float64(float32(offset) * scrollbar.contentToViewRatio)))
			newContentY := contentY - scaledOffset
			scrollbar.targetLayout.SetPosition(contentX, newContentY)
		}
	}

	if outOfBoundsDown && scrollbar.targetLayout != nil {
		newContentY := -scrollbar.contentSize + scrollbar.viewportSize
		scrollbar.targetLayout.SetPosition(0, newContentY)
	}

	if outOfBoundsUp && scrollbar.targetLayout != nil {
		scrollbar.targetLayout.SetPosition(0, 0)
	}

	return newY
}

// OnMouseMove will affect the slider and the content depending on the state fof it
func (scrollbar *LayoutScrollbar) OnMouseMove(event d2interface.MouseMoveEvent) {
	if !scrollbar.sliderClicked {
		return
	}

	sliderX, _ := scrollbar.sliderLayout.GetPosition()
	newY := scrollbar.moveScaledContentBy(event.Y() - scrollbar.clickedAtY)

	scrollbar.sliderLayout.SetPosition(sliderX, newY)
	scrollbar.updateSliderSpritesPosition()
}

// OnArrowUpClick will move the slider and the content up
func (scrollbar *LayoutScrollbar) OnArrowUpClick() {
	sliderX, _ := scrollbar.sliderLayout.GetPosition()
	newY := scrollbar.moveScaledContentBy(-scrollbar.arrowClickSliderOffset)

	scrollbar.sliderLayout.SetPosition(sliderX, newY)
	scrollbar.updateSliderSpritesPosition()
}

// OnArrowDownClick will move the slider and the content down
func (scrollbar *LayoutScrollbar) OnArrowDownClick() {
	sliderX, _ := scrollbar.sliderLayout.GetPosition()
	newY := scrollbar.moveScaledContentBy(scrollbar.arrowClickSliderOffset)

	scrollbar.sliderLayout.SetPosition(sliderX, newY)
	scrollbar.updateSliderSpritesPosition()
}

// SetSliderClicked sets the value of sliderClicked
func (scrollbar *LayoutScrollbar) SetSliderClicked(value bool) {
	scrollbar.sliderClicked = value
}

// SetArrowUpClicked sets the value of sliderClicked
func (scrollbar *LayoutScrollbar) SetArrowUpClicked(value bool) {
	var arrowSpriteFrame int

	scrollbar.arrowUpClicked = value

	if scrollbar.arrowUpClicked {
		arrowSpriteFrame = textSliderPartArrowUpHollow
	} else {
		arrowSpriteFrame = textSliderPartArrowUpFilled
	}

	if err := scrollbar.arrowUpSprite.SetCurrentFrame(arrowSpriteFrame); err != nil {
		log.Printf("unable to set arrow up sprite frame: %v", err)
	}
}

// SetArrowDownClicked sets the value of sliderClicked
func (scrollbar *LayoutScrollbar) SetArrowDownClicked(value bool) {
	var arrowSpriteFrame int

	scrollbar.arrowDownClicked = value

	if scrollbar.arrowDownClicked {
		arrowSpriteFrame = textSliderPartArrowDownHollow
	} else {
		arrowSpriteFrame = textSliderPartArrowDownFilled
	}

	if err := scrollbar.arrowDownSprite.SetCurrentFrame(arrowSpriteFrame); err != nil {
		log.Printf("unable to set arrow down sprite frame: %v", err)
	}
}

// Advance updates the layouts according to the state of the arrown
func (scrollbar *LayoutScrollbar) Advance(elapsed float64) error {
	if scrollbar.arrowDownClicked {
		scrollbar.OnArrowDownClick()
	}

	if scrollbar.arrowUpClicked {
		scrollbar.OnArrowUpClick()
	}

	return nil
}

// Render draws the scrollbar sprites on the given surface
func (scrollbar *LayoutScrollbar) Render(target d2interface.Surface) {
	for _, s := range scrollbar.gutterSprites {
		s.Render(target)
	}

	for _, s := range scrollbar.sliderSprites {
		s.Render(target)
	}

	scrollbar.arrowUpSprite.Render(target)
	scrollbar.arrowDownSprite.Render(target)
}

func (scrollbar *LayoutScrollbar) isInLayoutRect(layout *Layout, px, py int) bool {
	ww, hh := layout.GetSize()
	x, y := layout.Sx, layout.Sy

	if px >= x && px <= x+ww && py >= y && py <= y+hh {
		return true
	}

	return false
}

// IsSliderClicked returns the state of the slider
func (scrollbar *LayoutScrollbar) IsSliderClicked() bool {
	return scrollbar.sliderClicked
}

// IsArrowUpClicked returns the state of arrow up clicked
func (scrollbar *LayoutScrollbar) IsArrowUpClicked() bool {
	return scrollbar.arrowUpClicked
}

// IsArrowDownClicked returns the state of arrow down clicked
func (scrollbar *LayoutScrollbar) IsArrowDownClicked() bool {
	return scrollbar.arrowDownClicked
}

// IsInArrowUpRect checks if the given point is within the overlay layout rectangle
func (scrollbar *LayoutScrollbar) IsInArrowUpRect(px, py int) bool {
	return scrollbar.isInLayoutRect(scrollbar.arrowUpLayout, px, py)
}

// IsInArrowDownRect checks if the given point is within the overlay layout rectangle
func (scrollbar *LayoutScrollbar) IsInArrowDownRect(px, py int) bool {
	return scrollbar.isInLayoutRect(scrollbar.arrowDownLayout, px, py)
}

// IsInSliderRect checks if the given point is within the overlay layout rectangle
func (scrollbar *LayoutScrollbar) IsInSliderRect(px, py int) bool {
	return scrollbar.isInLayoutRect(scrollbar.sliderLayout, px, py)
}
