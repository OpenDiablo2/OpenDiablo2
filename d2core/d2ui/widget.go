package d2ui

// RenderPriority determines in which order ui elements are drawn.
// The higher the number the later an element is drawn.
type RenderPriority int

// Render priorities that determine the order in which widgets/widgetgroups are
// rendered. The higher the later it is rendered
const (
	RenderPriorityBackground RenderPriority = iota
	RenderPrioritySkilltree
	RenderPrioritySkilltreeIcon
	RenderPriorityHeroStatsPanel
	RenderPriorityQuestLog
	RenderPriorityInventory
	RenderPriorityHUDPanel
	RenderPriorityMinipanel
	RenderPriorityHelpPanel
	RenderPriorityForeground
)

// Widget defines an object that is a UI widget
type Widget interface {
	Drawable
	bindManager(ui *UIManager)
	GetManager() (ui *UIManager)
	OnMouseMove(x int, y int)
	OnHoverStart(callback func())
	OnHoverEnd(callback func())
	isHovered() bool
	hoverStart()
	hoverEnd()
	Contains(x, y int) (contained bool)
}

// ClickableWidget defines an object that can be clicked
type ClickableWidget interface {
	Widget
	SetEnabled(enabled bool)
	SetPressed(pressed bool)
	GetEnabled() bool
	GetPressed() bool
	OnActivated(callback func())
	Activate()
}

// BaseWidget contains default functionality that all widgets share
type BaseWidget struct {
	manager        *UIManager
	x              int
	y              int
	width          int
	height         int
	renderPriority RenderPriority
	visible        bool

	hovered        bool
	onHoverStartCb func()
	onHoverEndCb   func()
}

// NewBaseWidget creates a new BaseWidget with defaults
func NewBaseWidget(manager *UIManager) *BaseWidget {
	return &BaseWidget{
		manager:        manager,
		x:              0,
		y:              0,
		width:          0,
		height:         0,
		visible:        true,
		renderPriority: RenderPriorityBackground,
	}
}

func (b *BaseWidget) bindManager(manager *UIManager) {
	b.manager = manager
}

// GetSize returns the size of the widget
func (b *BaseWidget) GetSize() (width, height int) {
	return b.width, b.height
}

// SetPosition sets the position of the widget
func (b *BaseWidget) SetPosition(x, y int) {
	b.x, b.y = x, y
}

// OffsetPosition moves the widget by x and y
func (b *BaseWidget) OffsetPosition(x, y int) {
	b.x += x
	b.y += y
}

// GetPosition returns the position of the widget
func (b *BaseWidget) GetPosition() (x, y int) {
	return b.x, b.y
}

// GetVisible returns whether the widget is visible
func (b *BaseWidget) GetVisible() (visible bool) {
	return b.visible
}

// SetVisible make the widget visible, not visible
func (b *BaseWidget) SetVisible(visible bool) {
	b.visible = visible
}

// GetRenderPriority returns the order in which this widget is rendered
func (b *BaseWidget) GetRenderPriority() (prio RenderPriority) {
	return b.renderPriority
}

// SetRenderPriority sets the order in which this widget is rendered
func (b *BaseWidget) SetRenderPriority(prio RenderPriority) {
	b.renderPriority = prio
}

// OnHoverStart sets a function that is called if the hovering of the widget starts
func (b *BaseWidget) OnHoverStart(callback func()) {
	b.onHoverStartCb = callback
}

// HoverStart is called when the hovering of the widget starts
func (b *BaseWidget) hoverStart() {
	b.hovered = true
	if b.onHoverStartCb != nil {
		b.onHoverStartCb()
	}
}

// OnHoverEnd sets a function that is called if the hovering of the widget ends
func (b *BaseWidget) OnHoverEnd(callback func()) {
	b.onHoverEndCb = callback
}

// hoverEnd is called when the widget hovering ends
func (b *BaseWidget) hoverEnd() {
	b.hovered = false
	if b.onHoverEndCb != nil {
		b.onHoverEndCb()
	}
}

func (b *BaseWidget) isHovered() bool {
	return b.hovered
}

// Contains determines whether a given x,y coordinate lands within a Widget
func (b *BaseWidget) Contains(x, y int) bool {
	wx, wy := b.GetPosition()
	ww, wh := b.GetSize()

	return x >= wx && x <= wx+ww && y >= wy && y <= wy+wh
}

// GetManager returns the uiManager
func (b *BaseWidget) GetManager() (ui *UIManager) {
	return b.manager
}

// OnMouseMove is called when the mouse is moved
func (b *BaseWidget) OnMouseMove(x, y int) {
	if b.Contains(x, y) {
		if !b.isHovered() {
			b.hoverStart()
		}
	} else if b.isHovered() {
		b.hoverEnd()
	}
}
