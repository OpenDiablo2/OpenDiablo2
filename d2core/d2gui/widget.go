package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
)

type MouseHandler func(d2input.MouseEvent)
type MouseMoveHandler func(d2input.MouseMoveEvent)

type widget interface {
	render(target d2interface.Surface) error
	advance(elapsed float64) error

	onMouseMove(event d2input.MouseMoveEvent) bool
	onMouseEnter(event d2input.MouseMoveEvent) bool
	onMouseLeave(event d2input.MouseMoveEvent) bool
	onMouseOver(event d2input.MouseMoveEvent) bool
	onMouseButtonDown(event d2input.MouseEvent) bool
	onMouseButtonUp(event d2input.MouseEvent) bool
	onMouseButtonClick(event d2input.MouseEvent) bool

	getPosition() (int, int)
	setOffset(x, y int)
	getSize() (int, int)
	getLayer() int
	isVisible() bool
	isExpanding() bool
}

type widgetBase struct {
	x         int
	y         int
	layer     int
	visible   bool
	expanding bool

	offsetX int
	offsetY int

	mouseEnterHandler MouseMoveHandler
	mouseLeaveHandler MouseMoveHandler
	mouseClickHandler MouseHandler
}

func (w *widgetBase) SetPosition(x, y int) {
	w.x = x
	w.y = y
}

func (w *widgetBase) GetPosition() (int, int) {
	return w.x, w.y
}

func (w *widgetBase) GetOffset() (int, int) {
	return w.offsetX, w.offsetY
}

func (w *widgetBase) setOffset(x, y int) {
	w.offsetX = x
	w.offsetY = y
}

func (w *widgetBase) SetLayer(layer int) {
	w.layer = layer
}

func (w *widgetBase) SetVisible(visible bool) {
	w.visible = visible
}

func (w *widgetBase) SetExpanding(expanding bool) {
	w.expanding = expanding
}

func (w *widgetBase) SetMouseEnterHandler(handler MouseMoveHandler) {
	w.mouseEnterHandler = handler
}

func (w *widgetBase) SetMouseLeaveHandler(handler MouseMoveHandler) {
	w.mouseLeaveHandler = handler
}

func (w *widgetBase) SetMouseClickHandler(handler MouseHandler) {
	w.mouseClickHandler = handler
}

func (w *widgetBase) getPosition() (int, int) {
	return w.x, w.y
}

func (w *widgetBase) getSize() (int, int) {
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

func (w *widgetBase) render(target d2interface.Surface) error {
	return nil
}

func (w *widgetBase) advance(elapsed float64) error {
	return nil
}

func (w *widgetBase) onMouseEnter(event d2input.MouseMoveEvent) bool {
	if w.mouseEnterHandler != nil {
		w.mouseEnterHandler(event)
	}

	return false
}

func (w *widgetBase) onMouseLeave(event d2input.MouseMoveEvent) bool {
	if w.mouseLeaveHandler != nil {
		w.mouseLeaveHandler(event)
	}

	return false
}

func (w *widgetBase) onMouseButtonClick(event d2input.MouseEvent) bool {
	if w.mouseClickHandler != nil {
		w.mouseClickHandler(event)
	}

	return false
}

func (w *widgetBase) onMouseMove(event d2input.MouseMoveEvent) bool {
	return false
}

func (w *widgetBase) onMouseOver(event d2input.MouseMoveEvent) bool {
	return false
}

func (w *widgetBase) onMouseButtonDown(event d2input.MouseEvent) bool {
	return false
}

func (w *widgetBase) onMouseButtonUp(event d2input.MouseEvent) bool {
	return false
}
