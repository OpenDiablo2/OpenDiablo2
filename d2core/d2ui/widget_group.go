package d2ui

import (
	"image/color"
	"sort"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

const widgetGroupDebug = false // turns on debug rendering stuff for groups

// static check that WidgetGroup implements widget
var _ Widget = &WidgetGroup{}

// WidgetGroup allows the grouping of widgets to apply actions to all
// widgets at once.
type WidgetGroup struct {
	*BaseWidget
	entries []Widget
}

// NewWidgetGroup creates a new widget group
func (ui *UIManager) NewWidgetGroup(priority RenderPriority) *WidgetGroup {
	base := NewBaseWidget(ui)
	base.SetRenderPriority(priority)

	group := &WidgetGroup{
		BaseWidget: base,
	}

	ui.addWidget(group)

	return group
}

// AddWidget adds a widget to the group
func (wg *WidgetGroup) AddWidget(w Widget) {
	wg.adjustSize(w)
	wg.entries = append(wg.entries, w)
	sort.SliceStable(wg.entries, func(i, j int) bool {
		return wg.entries[i].GetRenderPriority() < wg.entries[j].GetRenderPriority()
	})

	if clickable, ok := w.(ClickableWidget); ok {
		wg.manager.addClickable(clickable)
	}
}

// adjustSize recalculates the bounding box if a new widget is added
func (wg *WidgetGroup) adjustSize(w Widget) {
	x, y := w.GetPosition()
	width, height := w.GetSize()

	if x+width > wg.x+wg.width {
		wg.width += (x + width) - (wg.x + wg.width)
	}

	if wg.x > x {
		wg.width += wg.x - x
		wg.x = x
	}

	if y+height > wg.y+wg.height {
		wg.height += (y + height) - (wg.y + wg.height)
	}

	if wg.y > y {
		wg.height += wg.y - y
		wg.y = y
	}
}

// Advance is a no-op here
func (wg *WidgetGroup) Advance(elapsed float64) error {
	// No-op
	return nil
}

// Render draw the widgets to the screen
func (wg *WidgetGroup) Render(target d2interface.Surface) {
	for _, entry := range wg.entries {
		if entry.GetVisible() {
			entry.Render(target)
		}
	}

	if widgetGroupDebug && wg.GetVisible() {
		wg.renderDebug(target)
	}
}

func (wg *WidgetGroup) renderDebug(target d2interface.Surface) {
	target.PushTranslation(wg.GetPosition())
	defer target.Pop()
	target.DrawLine(wg.width, 0, color.White)
	target.DrawLine(0, wg.height, color.White)

	target.PushTranslation(wg.width, wg.height)
	target.DrawLine(-wg.width, 0, color.White)
	target.DrawLine(0, -wg.height, color.White)
	target.Pop()
}

// SetVisible sets the visibility of all widgets in the group
func (wg *WidgetGroup) SetVisible(visible bool) {
	wg.BaseWidget.SetVisible(visible)

	for _, entry := range wg.entries {
		entry.SetVisible(visible)
	}
}

// OffsetPosition moves all widgets by x and y
func (wg *WidgetGroup) OffsetPosition(x, y int) {
	wg.BaseWidget.OffsetPosition(x, y)

	for _, entry := range wg.entries {
		entry.OffsetPosition(x, y)
	}
}

// OnMouseMove handles mouse move events
func (wg *WidgetGroup) OnMouseMove(x, y int) {
	for _, entry := range wg.entries {
		if entry.Contains(x, y) && entry.GetVisible() {
			if !entry.isHovered() {
				entry.hoverStart()
			}
		} else if entry.isHovered() {
			entry.hoverEnd()
		}
	}
}

// SetEnabled sets enable on all clickable widgets of this group
func (wg *WidgetGroup) SetEnabled(enabled bool) {
	for _, entry := range wg.entries {
		if v, ok := entry.(ClickableWidget); ok {
			v.SetEnabled(enabled)
		}
	}
}
