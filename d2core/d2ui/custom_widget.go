package d2ui

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

// static check that CustomWidget implements widget
var _ Widget = &CustomWidget{}

// CustomWidget is a widget with a fully custom render function
type CustomWidget struct {
	*BaseWidget
	renderFunc func(target d2interface.Surface)
	cached     bool
	cachedImg  *d2interface.Surface
	tooltip    *Tooltip
}

// NewCustomWidgetCached creates a new widget and caches anything rendered via the
// renderFunc into a static image to be displayed
func (ui *UIManager) NewCustomWidgetCached(renderFunc func(target d2interface.Surface), width, height int) *CustomWidget {
	c := ui.NewCustomWidget(renderFunc, width, height)
	c.cached = true

	// render using the renderFunc to a cache
	surface := ui.Renderer().NewSurface(width, height)
	c.cachedImg = &surface
	renderFunc(*c.cachedImg)

	return c
}

// NewCustomWidget creates a new widget with custom render function
func (ui *UIManager) NewCustomWidget(renderFunc func(target d2interface.Surface), width, height int) *CustomWidget {
	base := NewBaseWidget(ui)
	base.width = width
	base.height = height

	return &CustomWidget{
		BaseWidget: base,
		renderFunc: renderFunc,
	}
}

// Render draws the custom widget
func (c *CustomWidget) Render(target d2interface.Surface) {
	if c.cached {
		target.PushTranslation(c.GetPosition())
		target.Render(*c.cachedImg)
		target.Pop()
	} else {
		c.renderFunc(target)
	}
}

// SetTooltip gives this widget a Tooltip that is displayed if the widget is hovered
func (c *CustomWidget) SetTooltip(t *Tooltip) {
	c.tooltip = t
	c.OnHoverStart(func() { c.tooltip.SetVisible(true) })
	c.OnHoverEnd(func() { c.tooltip.SetVisible(false) })
}

// Advance is a no-op
func (c *CustomWidget) Advance(elapsed float64) error {
	return nil
}
