package d2ui

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

// static check that CustomWidget implements widget
var _ Widget = &CustomWidget{}

// CustomWidget is a widget with a fully custom render function
type CustomWidget struct {
	*BaseWidget
	renderFunc func(target d2interface.Surface) error
}

// NewCustomWidget creates a new widget with custom render function
func (ui *UIManager) NewCustomWidget(renderFunc func(target d2interface.Surface) error) *CustomWidget {
	base := NewBaseWidget(ui)

	return &CustomWidget{
		BaseWidget: base,
		renderFunc: renderFunc,
	}
}

// Render draws the custom widget
func (c *CustomWidget) Render(target d2interface.Surface) error {
	return c.renderFunc(target)
}

// Advance is a no-op
func (c *CustomWidget) Advance(elapsed float64) error {
	return nil
}
