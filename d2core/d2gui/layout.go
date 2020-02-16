package d2gui

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type layoutEntry struct {
	widget widget

	mouseOver bool
	mouseDown [3]bool
}

type Layout struct {
	widgetBase
	entries []*layoutEntry
}

func createLayout() *Layout {
	layout := new(Layout)
	layout.visible = true
	return layout
}

func (l *Layout) render(target d2render.Surface) error {
	for _, entry := range l.entries {
		if entry.widget.isVisible() {
			l.renderWidget(entry.widget, target)
			l.renderWidgetDebug(entry.widget, target)
		}
	}

	return nil
}

func (l *Layout) advance(elapsed float64) error {
	for _, entry := range l.entries {
		if entry.widget.isVisible() {
			if err := entry.widget.advance(elapsed); err != nil {
				return err
			}
		}
	}

	return nil
}

func (l *Layout) renderWidget(widget widget, target d2render.Surface) {
	target.PushTranslation(widget.getPosition())
	defer target.Pop()

	widget.render(target)
}

func (l *Layout) renderWidgetDebug(widget widget, target d2render.Surface) {
	target.PushTranslation(widget.getPosition())
	defer target.Pop()

	drawColor := color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xb0}

	width, height := widget.getSize()
	target.DrawLine(width, 0, drawColor)
	target.DrawLine(0, height, drawColor)

	target.PushTranslation(width, 0)
	target.DrawLine(0, height, drawColor)
	target.Pop()

	target.PushTranslation(0, height)
	target.DrawLine(width, 0, drawColor)
	target.Pop()
}

func (l *Layout) getSize() (int, int) {
	return 0, 0
}

func (l *Layout) onMouseButtonDown(event d2input.MouseEvent) bool {
	for _, entry := range l.entries {
		eventLocal := event

		if l.adjustEventCoords(entry.widget, &eventLocal.X, &eventLocal.Y) {
			entry.mouseDown[event.Button] = true
		}
	}

	return false
}

func (l *Layout) onMouseButtonUp(event d2input.MouseEvent) bool {
	for _, entry := range l.entries {
		eventLocal := event

		if l.adjustEventCoords(entry.widget, &eventLocal.X, &eventLocal.Y) {
			if entry.mouseDown[event.Button] {
				entry.widget.onMouseClick(eventLocal)
			}
		}

		entry.mouseDown[event.Button] = false
	}

	return false
}

func (l *Layout) onMouseMove(event d2input.MouseMoveEvent) bool {
	for _, entry := range l.entries {
		eventLocal := event

		if l.adjustEventCoords(entry.widget, &eventLocal.X, &eventLocal.Y) {
			if entry.mouseOver {
				entry.widget.onMouseOver(eventLocal)
			} else {
				entry.widget.onMouseEnter(eventLocal)
			}
			entry.mouseOver = true
		} else if entry.mouseOver {
			entry.widget.onMouseLeave(eventLocal)
			entry.mouseOver = false
		}
	}

	return false
}

func (l *Layout) adjustEventCoords(widget widget, eventX, eventY *int) bool {
	x, y := widget.getPosition()
	width, height := widget.getSize()

	*eventX -= x
	*eventY -= y

	if *eventX < 0 || *eventY < 0 || *eventX >= width || *eventY >= height {
		return false
	}

	return true
}

func (l *Layout) addLayout() *Layout {
	layout := createLayout()
	l.entries = append(l.entries, &layoutEntry{widget: layout})
	return layout
}

func (l *Layout) addSprite(imagePath, palettePath string) *Sprite {
	sprite := createSprite(imagePath, palettePath)
	l.entries = append(l.entries, &layoutEntry{widget: sprite})
	return sprite
}

func (l *Layout) addLabel(text string, fontStyle FontStyle) *Label {
	label := createLabel(text, fontStyle)
	l.entries = append(l.entries, &layoutEntry{widget: label})
	return label
}

func (l *Layout) clear() {
	l.entries = nil
}
