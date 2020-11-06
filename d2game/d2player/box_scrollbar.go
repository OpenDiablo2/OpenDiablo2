package d2player

import (
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
)

type BoxScrollbar struct {
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
	arrowUpClicked     bool
	arrowDownClicked   bool
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

func newBoxScrollBar(
	parentLayout *d2gui.Layout,
	targetLayout *d2gui.Layout,
) *BoxScrollbar {
	parentW, parentH := parentLayout.GetSize()
	_, targetH := targetLayout.GetSize()
	gutterHeight := parentH - (2 * textSliderPartHeight)
	viewportPercentage := 1.0 - (float32(targetH-parentH) / float32(targetH))
	sliderHeight := int(float32(gutterHeight) * viewportPercentage)
	x, y := parentW-textSliderPartWidth, 0

	ret := &BoxScrollbar{}
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

func (boxScroll *BoxScrollbar) onSliderMouseClick(event d2interface.MouseEvent) {
	if event.Button() == d2enum.MouseButtonLeft {
		boxScroll.sliderClicked = !boxScroll.sliderClicked
		if boxScroll.sliderClicked {
			boxScroll.clickedAtY = event.Y()
			boxScroll.lastY = boxScroll.clickedAtY
			boxScroll.mouseYOnSlider = event.Y() - boxScroll.sliderLayout.Sy
		}
	}
}

func (boxScroll *BoxScrollbar) onArrowUpClick(event d2interface.MouseEvent) {
	if event.Button() == d2enum.MouseButtonLeft {
		boxScroll.arrowUpClicked = !boxScroll.arrowUpClicked
	}
}

func (boxScroll *BoxScrollbar) onArrowDownClick(event d2interface.MouseEvent) {
	if event.Button() == d2enum.MouseButtonLeft {
		boxScroll.arrowDownClicked = !boxScroll.arrowDownClicked
	}
}

func (boxScroll *BoxScrollbar) update(cursorPosY int) {
	if boxScroll.arrowDownClicked {
		x, y := boxScroll.sliderLayout.GetPosition()
		newY := y + boxScroll.arrowSliderOffset
		if newY > boxScroll.maxY {
			newY = boxScroll.maxY
		}

		boxScroll.sliderLayout.SetPosition(x, newY)
	}

	if boxScroll.arrowUpClicked {
		x, y := boxScroll.sliderLayout.GetPosition()
		newY := y - boxScroll.arrowSliderOffset
		if newY < boxScroll.minY {
			newY = boxScroll.minY
		}

		boxScroll.sliderLayout.SetPosition(x, newY)
	}

	if boxScroll.sliderClicked {
		offset := cursorPosY - boxScroll.clickedAtY
		x, y := boxScroll.sliderLayout.GetPosition()
		newY := y + offset
		outOfBoundsUp := false
		outOfBoundsDown := false

		if newY > boxScroll.maxY {
			newY = boxScroll.maxY
			outOfBoundsDown = true
		}

		if newY < boxScroll.minY {
			newY = boxScroll.minY
			outOfBoundsUp = true
		}

		boxScroll.sliderLayout.SetPosition(x, newY)
		if !outOfBoundsUp && !outOfBoundsDown {
			boxScroll.clickedAtY = boxScroll.clickedAtY + offset
			if boxScroll.targetLayout != nil {
				contentX, contentY := boxScroll.targetLayout.GetPosition()
				scaledOffset := int(math.Round(float64(float32(offset) * boxScroll.contentToViewRatio)))
				newContentY := contentY - scaledOffset
				boxScroll.targetLayout.SetPosition(contentX, newContentY)
			}
		}

		if outOfBoundsDown && boxScroll.targetLayout != nil {
			newContentY := -boxScroll.contentSize + boxScroll.viewportSize
			boxScroll.targetLayout.SetPosition(0, newContentY)
		}

		if outOfBoundsUp && boxScroll.targetLayout != nil {
			boxScroll.targetLayout.SetPosition(0, 0)
		}
	}
}
