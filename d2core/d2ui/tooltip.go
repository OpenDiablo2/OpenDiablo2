package d2ui

import (
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

const (
	blackAlpha70 = 0x000000C8
	screenWidth  = 800
	screenHeight = 600
)

// static check that Tooltip implements widget
var _ Widget = &Tooltip{}

// Tooltip contains a label containing text with a transparent, black background
type Tooltip struct {
	*BaseWidget
	lines           []string
	label           *Label
	backgroundColor int
	originX         tooltipXOrigin
	originY         tooltipYOrigin
	boxEnabled      bool
}

type tooltipXOrigin = int
type tooltipYOrigin = int

const (
	// TooltipYTop sets the Y origin of the tooltip to the top
	TooltipYTop tooltipYOrigin = iota
	// TooltipYCenter sets the Y origin of the tooltip to the center
	TooltipYCenter
	// TooltipYBottom sets the Y origin of the tooltip to the bottom
	TooltipYBottom
)

const (
	// TooltipXLeft sets the X origin of the tooltip to the left
	TooltipXLeft tooltipXOrigin = iota
	// TooltipXCenter sets the X origin of the tooltip to the center
	TooltipXCenter
	// TooltipXRight sets the X origin of the tooltip to the right
	TooltipXRight
)

// NewTooltip creates a tooltip instance. Note here, that we need to define the
// orign point of the tooltip rect using tooltipXOrigin and tooltinYOrigin
func (ui *UIManager) NewTooltip(font,
	palette string,
	originX tooltipXOrigin,
	originY tooltipYOrigin) *Tooltip {
	label := ui.NewLabel(font, palette)
	label.Alignment = HorizontalAlignCenter

	base := NewBaseWidget(ui)
	base.SetVisible(false)

	res := &Tooltip{
		BaseWidget:      base,
		backgroundColor: blackAlpha70,
		label:           label,
		originX:         originX,
		originY:         originY,
		boxEnabled:      true,
	}
	res.manager = ui
	ui.addTooltip(res)

	return res
}

// SetTextLines sets the tooltip text in the form of an array of strings
func (t *Tooltip) SetTextLines(lines []string) {
	t.lines = lines
}

// SetText sets the tooltip text and splits \n into lines
func (t *Tooltip) SetText(text string) {
	t.lines = strings.Split(text, "\n")
}

// SetBoxEnabled determines whether a black box is drawn around the text
func (t *Tooltip) SetBoxEnabled(enable bool) {
	t.boxEnabled = enable
}

func (t *Tooltip) adjustCoordinatesToScreen(maxW, maxH, halfW, halfH int) (rx, ry int) {
	var xOffset, yOffset int

	switch t.originX {
	case TooltipXLeft:
		xOffset = maxW
	case TooltipXCenter:
		xOffset = halfW
	case TooltipXRight:
		xOffset = 0
	}

	renderX := t.x
	if (t.x + xOffset) > screenWidth {
		renderX = screenWidth - xOffset
	}

	switch t.originY {
	case TooltipYTop:
		yOffset = 0
	case TooltipYCenter:
		yOffset = halfH
	case TooltipYBottom:
		yOffset = maxH
	}

	renderY := t.y
	if (t.y + yOffset) > screenHeight {
		renderY = screenHeight - yOffset
	}

	return renderX, renderY
}

// GetSize returns the size of the tooltip
func (t *Tooltip) GetSize() (sx, sy int) {
	maxW, maxH := 0, 0

	for i := range t.lines {
		w, h := t.label.GetTextMetrics(t.lines[i])

		if maxW < w {
			maxW = w
		}

		maxH += h
	}

	return maxW, maxH
}

// Render draws the tooltip
func (t *Tooltip) Render(target d2interface.Surface) {
	maxW, maxH := t.GetSize()

	// nolint:gomnd // no magic numbers, their meaning is obvious
	halfW, halfH := maxW/2, maxH/2

	renderX, renderY := t.adjustCoordinatesToScreen(maxW, maxH, halfW, halfH)

	target.PushTranslation(renderX, renderY)
	defer target.Pop()

	// adjust starting point of the background rect based on the origin point
	// as we always draw a rect from top left
	switch t.originX {
	case TooltipXLeft:
		target.PushTranslation(0, 0)
	case TooltipXCenter:
		target.PushTranslation(-halfW, 0)
	case TooltipXRight:
		target.PushTranslation(-maxW, 0)
	}

	defer target.Pop()

	switch t.originY {
	case TooltipYTop:
		target.PushTranslation(0, 0)
	case TooltipYCenter:
		target.PushTranslation(0, -halfH)
	case TooltipXRight:
		target.PushTranslation(0, -maxH)
	}

	defer target.Pop()

	// tooltip background
	if t.boxEnabled {
		target.DrawRect(maxW, maxH, d2util.Color(blackAlpha70))
	}

	// text
	target.PushTranslation(halfW, 0) // text is centered, our box is not
	defer target.Pop()

	for i := range t.lines {
		t.label.SetText(t.lines[i])
		_, h := t.label.GetTextMetrics(t.lines[i])
		t.label.Render(target)
		target.PushTranslation(0, h)
	}

	target.PopN(len(t.lines))
}

// Advance is a no-op
func (t *Tooltip) Advance(elapsed float64) error {
	return nil
}
