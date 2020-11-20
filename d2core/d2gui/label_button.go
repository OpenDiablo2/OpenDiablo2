package d2gui

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// LabelButton is a label that can change when hovered and has
// a callback function that can be called when clicked
type LabelButton struct {
	label      string
	callback   func()
	hoverColor color.RGBA
	canHover   bool
	isHovered  bool
	layout     *Layout
	x, y       int

	*d2util.Logger
}

// NewLabelButton generates a new instance of LabelButton
func NewLabelButton(x, y int, text string, col color.RGBA, l d2util.LogLevel, callback func()) *LabelButton {
	lb := &LabelButton{
		x:          x,
		y:          y,
		hoverColor: col,
		label:      text,
		callback:   callback,
		canHover:   true,
	}

	lb.Logger = d2util.NewLogger()
	lb.Logger.SetLevel(l)
	lb.Logger.SetPrefix(logPrefix)

	return lb
}

// IsInRect checks if the given point is within the overlay layout rectangle
func (lb *LabelButton) IsInRect(px, py int) bool {
	if lb.layout == nil {
		return false
	}

	ww, hh := lb.layout.GetSize()
	x, y := lb.layout.Sx, lb.layout.Sy

	if px >= x && px <= x+ww && py >= y && py <= y+hh {
		return true
	}

	return false
}

// Load sets the button handlers and sets the layouts
func (lb *LabelButton) Load(renderer d2interface.Renderer, asset *d2asset.AssetManager) {
	mainLayout := CreateLayout(renderer, PositionTypeAbsolute, asset)
	l, _ := mainLayout.AddLabelWithColor(lb.label, FontStyleFormal11Units, d2util.Color(ColorBrown))

	if lb.canHover {
		l.SetHoverColor(lb.hoverColor)
	}

	mainLayout.SetMouseEnterHandler(func(event d2interface.MouseMoveEvent) {
		if err := l.SetIsHovered(true); err != nil {
			lb.Errorf("could not change label to hover state: %v", err)
		}
	})

	mainLayout.SetMouseLeaveHandler(func(event d2interface.MouseMoveEvent) {
		if err := l.SetIsHovered(false); err != nil {
			lb.Errorf("could not change label to hover state: %v", err)
		}
	})

	lb.layout = mainLayout
}

// SetLabel sets the text of label label
func (lb *LabelButton) SetLabel(val string) {
	lb.label = val
}

// SetHoverColor sets the hover color of the Label
func (lb *LabelButton) SetHoverColor(col color.RGBA) {
	lb.hoverColor = col
}

// SetCanHover sets the value of canHover
func (lb *LabelButton) SetCanHover(val bool) {
	lb.canHover = val
}

// IsHovered returns the value of isHovered
func (lb *LabelButton) IsHovered() bool {
	return lb.isHovered
}

// GetLayout returns the laout of the label
func (lb *LabelButton) GetLayout() *Layout {
	return lb.layout
}
