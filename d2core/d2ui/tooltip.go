package d2ui

import (
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
)

const (
	blackAlpha70 = 0x000000C8
	screenWidth  = 800
	screenHeight = 600
)

type Tooltip struct {
	manager         *UIManager
	lines           []string
	label           *Label
	backgroundColor int
	x, y            int
	originX         TooltipXOrigin
	originY         TooltipYOrigin
}

type TooltipXOrigin = int
type TooltipYOrigin = int
const (
	TooltipYTop TooltipYOrigin = iota
	TooltipYCenter
	TooltipYBottom
)

const (
	TooltipXLeft TooltipXOrigin = iota
	TooltipXCenter
	TooltipXRight
)


func (ui *UIManager) NewToolTip(font string,
	palette string,
	originX TooltipXOrigin,
	originY TooltipYOrigin) *Tooltip {
	label := ui.NewLabel(font, palette)
	label.Alignment = d2gui.HorizontalAlignCenter

	res := &Tooltip {
		backgroundColor: blackAlpha70,
		label: label,
		x: 0,
		y: 0,
		originX: originX,
		originY: originY,
	}
	res.manager = ui
	return res
}

func (t *Tooltip) SetPosition(x int, y int) {
	t.x = x
	t.y = y
}

func (t *Tooltip) SetTextLines(lines []string) {
	t.lines = lines
}

func (t *Tooltip) SetText(text string) {
	t.lines = strings.Split(text, "\n")
}

func (t *Tooltip) adjustCoordinatesToScreen(maxW int, maxH int, halfW int, halfH int) (int, int){
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

func (t *Tooltip) getTextSize() (int, int){
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

func (t *Tooltip) Render(target d2interface.Surface) {
	maxW, maxH := t.getTextSize()
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
		target.PushTranslation(0 , -maxH)
	}
	defer target.Pop()

	// tooltip background
	target.DrawRect(maxW, maxH, d2util.Color(blackAlpha70))

	// text
	target.PushTranslation(halfW, 0) // text is centered, our box is not
	defer target.Pop()
	for i:= range t.lines {
		t.label.SetText(t.lines[i])
		_, h := t.label.GetTextMetrics(t.lines[i])
		t.label.Render(target)
		target.PushTranslation(0, h)
	}
	target.PopN(len(t.lines))
}
