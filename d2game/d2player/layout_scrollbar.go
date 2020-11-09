package d2player

import (
	"fmt"
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
)

type LayoutScrollbar struct {
	x                 int
	y                 int
	maxY              int
	minY              int
	arrowSliderOffset int
	viewportSize      int
	contentSize       int

	parentLayout *d2gui.Layout
	targetLayout *d2gui.Layout
	sliderLayout *d2gui.Layout

	sliderClicked      bool
	clickedAtY         int
	mouseYOnSlider     int
	lastY              int
	gutterHeight       int
	contentToViewRatio float32
}

const (
	textSliderPartWidth  = 12
	textSliderPartHeight = 13
)

type textSliderPart int

const (
	textSliderPartArrowDownHollow textSliderPart = 8
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
	fmt.Println("gutterHeight", gutterHeight)
	viewportPercentage := 1.0 - (float32(targetH-parentH) / float32(targetH))
	sliderHeight := int(float32(gutterHeight) * viewportPercentage)
	x, y := parentW-textSliderPartWidth, 0

	ret := &LayoutScrollbar{}
	ret.contentToViewRatio = viewportPercentage
	ret.contentToViewRatio = float32(targetH) / float32(gutterHeight)
	ret.gutterHeight = gutterHeight
	ret.minY = y + textSliderPartHeight
	ret.maxY = (y + parentH) - (textSliderPartHeight + sliderHeight)
	ret.contentSize = targetH
	ret.arrowSliderOffset = int(float32(sliderHeight) * 0.02)
	ret.viewportSize = parentH

	arrowUpLayout := parentLayout.AddLayout(d2gui.PositionTypeAbsolute)
	arrowUpLayout.SetSize(textSliderPartWidth, textSliderPartHeight)
	arrowUpLayout.SetPosition(x, 0)
	arrowUpLayout.SetMouseClickHandler(ret.onArrowUpClick)

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

	ret.sliderLayout = sliderLayout
	ret.parentLayout = parentLayout
	ret.targetLayout = targetLayout

	return ret
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
	}
}

func (scrollbar *LayoutScrollbar) onArrowUpClick(event d2interface.MouseEvent) {
	x, y := scrollbar.sliderLayout.GetPosition()
	newY := y - scrollbar.arrowSliderOffset
	if newY < scrollbar.minY {
		newY = scrollbar.minY
	}

	scrollbar.sliderLayout.SetPosition(x, newY)
}

func (scrollbar *LayoutScrollbar) onArrowDownClick(event d2interface.MouseEvent) {
	x, y := scrollbar.sliderLayout.GetPosition()
	newY := y + scrollbar.arrowSliderOffset
	if newY > scrollbar.maxY {
		newY = scrollbar.maxY
	}

	scrollbar.sliderLayout.SetPosition(x, newY)
}

func (scrollbar *LayoutScrollbar) SetSliderClicked(value bool) {
	scrollbar.sliderClicked = value
}

// IsInSliderRect checks if the given point is within the overlay layout rectangle
func (scrollbar *LayoutScrollbar) IsInSliderRect(px, py int) bool {
	ww, hh := scrollbar.sliderLayout.GetSize()
	x, y := scrollbar.sliderLayout.Sx, scrollbar.sliderLayout.Sy

	fmt.Printf("Mouse click at (%d,%d), slider at (%d,%d) size: %dx%d\n", px, py, x, y, ww, hh)

	if px >= x && px <= x+ww && py >= y && py <= y+hh {
		fmt.Println("Slider HIT")
		return true
	}

	return false
}
