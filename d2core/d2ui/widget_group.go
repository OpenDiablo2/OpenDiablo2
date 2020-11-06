package d2ui

import (
	"sort"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// static check that WidgetGroup implements widget
var _ Widget = &WidgetGroup{}

// WidgetGroup allows the grouping of widgets to apply actions to all
// widgets at once.
type WidgetGroup struct {
	*BaseWidget
	entries  []Widget
	priority RenderPriority
}

// NewWidgetGroup creates a new widget group
func (ui *UIManager) NewWidgetGroup(priority RenderPriority) *WidgetGroup {
	base := NewBaseWidget(ui)
	base.SetRenderPriority(priority)

	group := &WidgetGroup{
		BaseWidget: base,
	}

	ui.addWidgetGroup(group)

	return group
}

// AddWidget adds a widget to the group
func (wg *WidgetGroup) AddWidget(w Widget) {
	wg.adjustSize(w)
	wg.entries = append(wg.entries, w)
	sort.SliceStable(wg.entries, func(i, j int) bool {
		return wg.entries[i].GetRenderPriority() < wg.entries[j].GetRenderPriority()
	})
}

// adjustSize recalculates the bounding box if a new widget is added
func (wg *WidgetGroup) adjustSize(w Widget) {
	x, y := w.GetPosition()
	width, height := w.GetSize()

	if x+width > wg.width {
		wg.width = x + width
	}

	if wg.x > x {
		wg.width += wg.x - x
		wg.x = x
	}

	if y+height > wg.height {
		wg.height = x + height
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
func (wg *WidgetGroup) Render(target d2interface.Surface) error {
	for _, entry := range wg.entries {
		if entry.GetVisible() {
			err := entry.Render(target)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// SetVisible sets the visibility of all widgets in the group
func (wg *WidgetGroup) SetVisible(visible bool) {
	for _, entry := range wg.entries {
		entry.SetVisible(visible)
	}
}
