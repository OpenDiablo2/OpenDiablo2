package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type MouseHandler func(d2input.MouseEvent)
type MouseMoveHandler func(d2input.MouseMoveEvent)

type widget interface {
	render(target d2render.Surface) error
	advance(elapsed float64) error

	onMouseEnter(event d2input.MouseMoveEvent)
	onMouseLeave(event d2input.MouseMoveEvent)
	onMouseOver(event d2input.MouseMoveEvent)
	onMouseClick(event d2input.MouseEvent)

	getPosition() (int, int)
	getSize() (int, int)
	getLayer() int
	isVisible() bool
}

type widgetBase struct {
	x       int
	y       int
	layer   int
	visible bool

	mouseEnterHandler MouseMoveHandler
	mouseLeaveHandler MouseMoveHandler
	mouseMoveHandler  MouseMoveHandler
	mouseClickHandler MouseHandler
}

func (w *widgetBase) SetPosition(x, y int) {
	w.x = x
	w.y = y
}

func (w *widgetBase) SetLayer(layer int) {
	w.layer = layer
}

func (w *widgetBase) SetVisible(visible bool) {
	w.visible = visible
}

func (w *widgetBase) SetMouseEnterHandler(handler MouseMoveHandler) {
	w.mouseEnterHandler = handler
}

func (w *widgetBase) SetMouseLeaveHandler(handler MouseMoveHandler) {
	w.mouseLeaveHandler = handler
}

func (w *widgetBase) SetMouseMoveHandler(handler MouseMoveHandler) {
	w.mouseMoveHandler = handler
}

func (w *widgetBase) SetMouseClickHandler(handler MouseHandler) {
	w.mouseClickHandler = handler
}

func (w *widgetBase) getPosition() (int, int) {
	return w.x, w.y
}

func (w *widgetBase) getLayer() int {
	return w.layer
}

func (w *widgetBase) isVisible() bool {
	return w.visible
}

func (w *widgetBase) advance(elapsed float64) error {
	return nil
}

func (w *widgetBase) onMouseEnter(event d2input.MouseMoveEvent) {
	if w.mouseEnterHandler != nil {
		w.mouseEnterHandler(event)
	}
}

func (w *widgetBase) onMouseLeave(event d2input.MouseMoveEvent) {
	if w.mouseLeaveHandler != nil {
		w.mouseLeaveHandler(event)
	}
}

func (w *widgetBase) onMouseOver(event d2input.MouseMoveEvent) {
	if w.mouseMoveHandler != nil {
		w.mouseMoveHandler(event)
	}
}

func (w *widgetBase) onMouseClick(event d2input.MouseEvent) {
	if w.mouseClickHandler != nil {
		w.mouseClickHandler(event)
	}
}
