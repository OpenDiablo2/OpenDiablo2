package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// MouseHandler is a handler function for mouse events
type MouseHandler func(d2interface.MouseEvent)

// MouseMoveHandler is a handler function for mouse movement events
type MouseMoveHandler func(d2interface.MouseMoveEvent)

type widget interface {
	render(target d2interface.Surface)
	advance(elapsed float64) error

	onMouseMove(event d2interface.MouseMoveEvent) bool
	onMouseEnter(event d2interface.MouseMoveEvent) bool
	onMouseLeave(event d2interface.MouseMoveEvent) bool
	onMouseOver(event d2interface.MouseMoveEvent) bool
	onMouseButtonDown(event d2interface.MouseEvent) bool
	onMouseButtonUp(event d2interface.MouseEvent) bool
	onMouseButtonClick(event d2interface.MouseEvent) bool

	getPosition() (int, int)
	setOffset(x, y int)
	SetScreenPos(x, y int)
	ScreenPos() (x, y int)
	getSize() (int, int)
	getLayer() int
	isVisible() bool
	isExpanding() bool
}

type widgetBase struct {
	x         int
	y         int
	Sx        int
	Sy        int
	layer     int
	visible   bool
	expanding bool

	offsetX int
	offsetY int

	mouseEnterHandler MouseMoveHandler
	mouseLeaveHandler MouseMoveHandler
	mouseClickHandler MouseHandler
}

// SetPosition sets the widget position
func (w *widgetBase) SetPosition(x, y int) {
	w.x = x
	w.y = y
}

// GetPosition returns the widget position
func (w *widgetBase) GetPosition() (x, y int) {
	return w.x, w.y
}

// GetOffset gets the widget offset
func (w *widgetBase) GetOffset() (x, y int) {
	return w.offsetX, w.offsetY
}

// SetScreenPos sets the screen position
func (w *widgetBase) SetScreenPos(x, y int) {
	w.Sx, w.Sy = x, y
}

// ScreenPos returns the screen position
func (w *widgetBase) ScreenPos() (x, y int) {
	return w.Sx, w.Sy
}

func (w *widgetBase) setOffset(x, y int) {
	w.offsetX = x
	w.offsetY = y
}

// SetLayer sets the widget layer (for rendering order)
func (w *widgetBase) SetLayer(layer int) {
	w.layer = layer
}

// SetVisible sets the widget's visibility
func (w *widgetBase) SetVisible(visible bool) {
	w.visible = visible
}

// SetExpanding tells the widget to expand (or not) when inside of a layout
func (w *widgetBase) SetExpanding(expanding bool) {
	w.expanding = expanding
}

// SetMouseEnterHandler sets the handler function for mouse-enter events
func (w *widgetBase) SetMouseEnterHandler(handler MouseMoveHandler) {
	w.mouseEnterHandler = handler
}

// SetMouseLeaveHandler sets the handler function for mouse-leave events
func (w *widgetBase) SetMouseLeaveHandler(handler MouseMoveHandler) {
	w.mouseLeaveHandler = handler
}

// SetMouseClickHandler sets the handler function for mouse-click events
func (w *widgetBase) SetMouseClickHandler(handler MouseHandler) {
	w.mouseClickHandler = handler
}

func (w *widgetBase) getPosition() (x, y int) {
	return w.x, w.y
}

func (w *widgetBase) getSize() (width, height int) {
	return 0, 0
}

func (w *widgetBase) getLayer() int {
	return w.layer
}

func (w *widgetBase) isVisible() bool {
	return w.visible
}

func (w *widgetBase) isExpanding() bool {
	return w.expanding
}

func (w *widgetBase) render(_ d2interface.Surface) { /* NOOP */ }

func (w *widgetBase) advance(_ float64) error {
	return nil
}

func (w *widgetBase) onMouseEnter(event d2interface.MouseMoveEvent) bool {
	if w.mouseEnterHandler != nil {
		w.mouseEnterHandler(event)
	}

	return false
}

func (w *widgetBase) onMouseLeave(event d2interface.MouseMoveEvent) bool {
	if w.mouseLeaveHandler != nil {
		w.mouseLeaveHandler(event)
	}

	return false
}

func (w *widgetBase) onMouseButtonClick(event d2interface.MouseEvent) bool {
	if w.mouseClickHandler != nil {
		w.mouseClickHandler(event)
	}

	return false
}

func (w *widgetBase) onMouseMove(_ d2interface.MouseMoveEvent) bool {
	return false
}

func (w *widgetBase) onMouseOver(_ d2interface.MouseMoveEvent) bool {
	return false
}

func (w *widgetBase) onMouseButtonDown(_ d2interface.MouseEvent) bool {
	return false
}

func (w *widgetBase) onMouseButtonUp(_ d2interface.MouseEvent) bool {
	return false
}
