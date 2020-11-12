package d2player

import (
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type LayoutScrollbar struct {
	maxY              int
	minY              int
	arrowSliderOffset int
	viewportSize      int
	contentSize       int

	parentLayout *d2gui.Layout
	targetLayout *d2gui.Layout
	sliderLayout *d2gui.Layout

	arrowUpLayout   *d2gui.Layout
	arrowDownLayout *d2gui.Layout

	isVisible          bool
	arrowUpClicked     bool
	arrowDownClicked   bool
	sliderClicked      bool
	clickedAtY         int
	mouseYOnSlider     int
	lastY              int
	gutterHeight       int
	sliderHeight       int
	contentToViewRatio float32

	arrowUpSprite   *d2ui.Sprite
	arrowDownSprite *d2ui.Sprite
	sliderSprites   []*d2ui.Sprite
	gutterSprites   []*d2ui.Sprite
}

const (
	textSliderPartWidth  = 12
	textSliderPartHeight = 13
)

const (
	textSliderPartArrowDownHollow int = iota + 8
	textSliderPartArrowUpHollow
	textSliderPartArrowDownFilled
	textSliderPartArrowUpFilled
	textSliderPartSquare
	textSliderPartInnerGutter
	textSliderPartFillingVariation1
)

func newLayoutScrollbar(
	parentLayout *d2gui.Layout,
	targetLayout *d2gui.Layout,
) *LayoutScrollbar {
	parentW, parentH := parentLayout.GetSize()
	_, targetH := targetLayout.GetSize()
	gutterHeight := parentH - (2 * textSliderPartHeight)
	viewportPercentage := 1.0 - (float32(targetH-parentH) / float32(targetH))
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
	ret.arrowSliderOffset = int(float32(sliderHeight) * 0.02)
	ret.viewportSize = parentH

	arrowUpLayout := parentLayout.AddLayout(d2gui.PositionTypeAbsolute)
	arrowUpLayout.SetSize(textSliderPartWidth, textSliderPartHeight)
	arrowUpLayout.SetPosition(x, 0)
	arrowUpLayout.SetMouseClickHandler(ret.onArrowUpClick)
	ret.arrowUpLayout = arrowUpLayout

	gutterLayout := parentLayout.AddLayout(d2gui.PositionTypeAbsolute)
	gutterLayout.SetSize(textSliderPartWidth, gutterHeight)
	gutterLayout.SetPosition(x, textSliderPartHeight)

	sliderLayout := parentLayout.AddLayout(d2gui.PositionTypeAbsolute)
	sliderLayout.SetPosition(x, textSliderPartHeight)
	sliderLayout.SetSize(textSliderPartWidth, sliderHeight)
	sliderLayout.SetMouseClickHandler(ret.onSliderMouseClick)

	arrowDownLayout := parentLayout.AddLayout(d2gui.PositionTypeAbsolute)
	arrowDownLayout.SetSize(textSliderPartWidth, textSliderPartHeight)
	arrowDownLayout.SetPosition(x, textSliderPartHeight+gutterHeight)
	arrowDownLayout.SetMouseClickHandler(ret.onArrowDownClick)
	ret.arrowDownLayout = arrowDownLayout

	ret.sliderLayout = sliderLayout
	ret.parentLayout = parentLayout
	ret.targetLayout = targetLayout

	ret.parentLayout.AdjustEntryPlacement()
	ret.targetLayout.AdjustEntryPlacement()
	ret.sliderLayout.AdjustEntryPlacement()
	return ret
}

func (scrollbar *LayoutScrollbar) Load(ui *d2ui.UIManager) {
	arrowUpX, arrowUpY := scrollbar.arrowUpLayout.ScreenPos()
	arrowUpSprite, _ := ui.NewSprite(d2resource.TextSlider, d2resource.PaletteSky)
	arrowUpSprite.SetCurrentFrame(textSliderPartArrowUpFilled)
	arrowUpSprite.SetPosition(arrowUpX, arrowUpY+textSliderPartHeight)
	scrollbar.arrowUpSprite = arrowUpSprite

	arrowDownX, arrowDownY := scrollbar.arrowDownLayout.ScreenPos()
	arrowDownSprite, _ := ui.NewSprite(d2resource.TextSlider, d2resource.PaletteSky)
	arrowDownSprite.SetCurrentFrame(textSliderPartArrowDownFilled)
	arrowDownSprite.SetPosition(arrowDownX, arrowDownY+textSliderPartHeight)
	scrollbar.arrowDownSprite = arrowDownSprite

	gutterParts := int(math.Ceil(float64(scrollbar.gutterHeight+(2*textSliderPartHeight)) / float64(textSliderPartHeight)))
	sliderParts := int(math.Ceil(float64(scrollbar.sliderHeight) / float64(textSliderPartHeight)))

	i := 0
	gutterX, gutterY := arrowUpX, arrowUpY+(2*textSliderPartHeight)-1
	for {
		if i >= gutterParts {
			break
		}

		f, _ := ui.NewSprite(d2resource.TextSlider, d2resource.PaletteSky)
		f.SetCurrentFrame(textSliderPartInnerGutter)
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
		f.SetCurrentFrame(textSliderPartFillingVariation1)

		scrollbar.sliderSprites = append(scrollbar.sliderSprites, f)

		i++
	}

	scrollbar.updateSliderSpritesPosition()
}

func (scrollbar *LayoutScrollbar) updateSliderSpritesPosition() {
	scrollbar.sliderLayout.AdjustEntryPlacement()
	sliderLayoutX, sliderLayoutY := scrollbar.sliderLayout.ScreenPos()

	for i, s := range scrollbar.sliderSprites {
		newY := sliderLayoutY + (i * (textSliderPartHeight - 1)) + textSliderPartHeight
		s.SetPosition(sliderLayoutX-1, newY)
	}
}

func (scrollbar *LayoutScrollbar) onSliderMouseClick(event d2interface.MouseEvent) {
	scrollbar.clickedAtY = event.Y()
	scrollbar.lastY = scrollbar.clickedAtY
	scrollbar.mouseYOnSlider = event.Y() - scrollbar.sliderLayout.Sy
}

func (scrollbar *LayoutScrollbar) onMouseMove(event d2interface.MouseMoveEvent) {
	if scrollbar.sliderClicked {
		offset := event.Y() - scrollbar.clickedAtY
		x, y := scrollbar.sliderLayout.GetPosition()
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

		scrollbar.sliderLayout.SetPosition(x, newY)
		if !outOfBoundsUp && !outOfBoundsDown {
			scrollbar.clickedAtY = scrollbar.clickedAtY + offset
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

		scrollbar.updateSliderSpritesPosition()
	}
}

func (scrollbar LayoutScrollbar) updateContentPosition() {

}

func (scrollbar *LayoutScrollbar) onArrowUpClick(event d2interface.MouseEvent) {
	x, y := scrollbar.sliderLayout.GetPosition()
	newY := y - scrollbar.arrowSliderOffset
	if newY < scrollbar.minY {
		newY = scrollbar.minY
	}

	scrollbar.sliderLayout.SetPosition(x, newY)
	scrollbar.updateSliderSpritesPosition()
}

func (scrollbar *LayoutScrollbar) onArrowDownClick(event d2interface.MouseEvent) {
	x, y := scrollbar.sliderLayout.GetPosition()
	newY := y + scrollbar.arrowSliderOffset
	if newY > scrollbar.maxY {
		newY = scrollbar.maxY
	}

	scrollbar.sliderLayout.SetPosition(x, newY)
	scrollbar.updateSliderSpritesPosition()
}

func (scrollbar *LayoutScrollbar) SetSliderClicked(value bool) {
	scrollbar.sliderClicked = value
}

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

// IsInSliderRect checks if the given point is within the overlay layout rectangle
func (scrollbar *LayoutScrollbar) IsInSliderRect(px, py int) bool {
	ww, hh := scrollbar.sliderLayout.GetSize()
	x, y := scrollbar.sliderLayout.Sx, scrollbar.sliderLayout.Sy

	if px >= x && px <= x+ww && py >= y && py <= y+hh {
		return true
	}

	return false
}
