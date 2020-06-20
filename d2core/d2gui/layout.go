package d2gui

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
)

type layoutEntry struct {
	widget widget

	x      int
	y      int
	width  int
	height int

	mouseOver bool
	mouseDown [3]bool
}

type VerticalAlign int

const (
	VerticalAlignTop VerticalAlign = iota
	VerticalAlignMiddle
	VerticalAlignBottom
)

type HorizontalAlign int

const (
	HorizontalAlignLeft HorizontalAlign = iota
	HorizontalAlignCenter
	HorizontalAlignRight
)

type PositionType int

const (
	PositionTypeAbsolute PositionType = iota
	PositionTypeVertical
	PositionTypeHorizontal
)

type Layout struct {
	widgetBase

	width           int
	height          int
	verticalAlign   VerticalAlign
	horizontalAlign HorizontalAlign
	positionType    PositionType
	entries         []*layoutEntry
}

func createLayout(positionType PositionType) *Layout {
	layout := &Layout{positionType: positionType}
	layout.SetVisible(true)

	return layout
}

func (l *Layout) SetSize(width, height int) {
	l.width = width
	l.height = height
}

func (l *Layout) SetVerticalAlign(verticalAlign VerticalAlign) {
	l.verticalAlign = verticalAlign
}

func (l *Layout) SetHorizontalAlign(horizontalAlign HorizontalAlign) {
	l.horizontalAlign = horizontalAlign
}

func (l *Layout) AddLayout(positionType PositionType) *Layout {
	layout := createLayout(positionType)
	l.entries = append(l.entries, &layoutEntry{widget: layout})
	return layout
}

func (l *Layout) AddSpacerStatic(width, height int) *SpacerStatic {
	spacer := createSpacerStatic(width, height)
	l.entries = append(l.entries, &layoutEntry{widget: spacer})
	return spacer
}

func (l *Layout) AddSpacerDynamic() *SpacerDynamic {
	spacer := createSpacerDynamic()
	l.entries = append(l.entries, &layoutEntry{widget: spacer})
	return spacer
}

func (l *Layout) AddSprite(imagePath, palettePath string) (*Sprite, error) {
	sprite, err := createSprite(imagePath, palettePath)
	if err != nil {
		return nil, err
	}

	l.entries = append(l.entries, &layoutEntry{widget: sprite})
	return sprite, nil
}

func (l *Layout) AddAnimatedSprite(imagePath, palettePath string, direction AnimationDirection) (*AnimatedSprite, error) {
	sprite, err := createAnimatedSprite(imagePath, palettePath, direction)
	if err != nil {
		return nil, err
	}

	l.entries = append(l.entries, &layoutEntry{widget: sprite})
	return sprite, nil
}

func (l *Layout) AddLabel(text string, fontStyle FontStyle) (*Label, error) {
	label, err := createLabel(text, fontStyle)
	if err != nil {
		return nil, err
	}

	l.entries = append(l.entries, &layoutEntry{widget: label})
	return label, nil
}

func (l *Layout) AddButton(text string, buttonStyle ButtonStyle) (*Button, error) {
	button, err := createButton(text, buttonStyle)
	if err != nil {
		return nil, err
	}

	l.entries = append(l.entries, &layoutEntry{widget: button})
	return button, nil
}

func (l *Layout) Clear() {
	l.entries = nil
}

func (l *Layout) render(target d2interface.Surface) error {
	l.AdjustEntryPlacement()

	for _, entry := range l.entries {
		if !entry.widget.isVisible() {
			continue
		}

		if err := l.renderEntry(entry, target); err != nil {
			return err
		}

		// uncomment to see debug boxes
		//if err := l.renderEntryDebug(entry, target); err != nil {
		//	return err
		//}
	}

	return nil
}

func (l *Layout) advance(elapsed float64) error {
	for _, entry := range l.entries {
		if err := entry.widget.advance(elapsed); err != nil {
			return err
		}
	}

	return nil
}

func (l *Layout) renderEntry(entry *layoutEntry, target d2interface.Surface) error {
	target.PushTranslation(entry.x, entry.y)
	defer target.Pop()

	return entry.widget.render(target)
}

func (l *Layout) renderEntryDebug(entry *layoutEntry, target d2interface.Surface) error {
	target.PushTranslation(entry.x, entry.y)
	defer target.Pop()

	drawColor := color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	switch entry.widget.(type) {
	case *Layout:
		drawColor = color.RGBA{R: 0xff, G: 0x00, B: 0xff, A: 0xff}
	case *SpacerStatic, *SpacerDynamic:
		drawColor = color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xff}
	case *Label:
		drawColor = color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}
	case *Button:
		drawColor = color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
	}

	target.DrawLine(entry.width, 0, drawColor)
	target.DrawLine(0, entry.height, drawColor)

	target.PushTranslation(entry.width, 0)
	target.DrawLine(0, entry.height, drawColor)
	target.Pop()

	target.PushTranslation(0, entry.height)
	target.DrawLine(entry.width, 0, drawColor)
	target.Pop()

	return nil
}

func (l *Layout) getContentSize() (int, int) {
	var width, height int

	for _, entry := range l.entries {
		x, y := entry.widget.getPosition()
		w, h := entry.widget.getSize()

		switch l.positionType {
		case PositionTypeVertical:
			width = d2common.MaxInt(width, w)
			height += h
		case PositionTypeHorizontal:
			width += w
			height = d2common.MaxInt(height, h)
		case PositionTypeAbsolute:
			width = d2common.MaxInt(width, x+w)
			height = d2common.MaxInt(height, y+h)
		}
	}

	return width, height
}

func (l *Layout) getSize() (int, int) {
	width, height := l.getContentSize()
	return d2common.MaxInt(width, l.width), d2common.MaxInt(height, l.height)
}

func (l *Layout) onMouseButtonDown(event d2input.MouseEvent) bool {
	for _, entry := range l.entries {
		eventLocal := event

		if l.adjustEntryEvent(entry, &eventLocal.X, &eventLocal.Y) {
			entry.widget.onMouseButtonDown(eventLocal)
			entry.mouseDown[event.Button] = true
		}
	}

	return false
}

func (l *Layout) onMouseButtonUp(event d2input.MouseEvent) bool {
	for _, entry := range l.entries {
		eventLocal := event

		if l.adjustEntryEvent(entry, &eventLocal.X, &eventLocal.Y) {
			if entry.mouseDown[event.Button] {
				entry.widget.onMouseButtonClick(eventLocal)
				entry.widget.onMouseButtonUp(eventLocal)
			}
		}

		entry.mouseDown[event.Button] = false
	}

	return false
}

func (l *Layout) onMouseMove(event d2input.MouseMoveEvent) bool {
	for _, entry := range l.entries {
		eventLocal := event

		if l.adjustEntryEvent(entry, &eventLocal.X, &eventLocal.Y) {
			entry.widget.onMouseMove(eventLocal)
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

func (l *Layout) adjustEntryEvent(entry *layoutEntry, eventX, eventY *int) bool {
	*eventX -= entry.x
	*eventY -= entry.y

	if *eventX < 0 || *eventY < 0 || *eventX >= entry.width || *eventY >= entry.height {
		return false
	}

	return true
}

func (l *Layout) AdjustEntryPlacement() {
	width, height := l.getSize()

	var expanderCount int
	for _, entry := range l.entries {
		if entry.widget.isVisible() && entry.widget.isExpanding() {
			expanderCount++
		}
	}

	var expanderWidth, expanderHeight int
	if expanderCount > 0 {
		contentWidth, contentHeight := l.getContentSize()

		switch l.positionType {
		case PositionTypeVertical:
			expanderHeight = (height - contentHeight) / expanderCount
		case PositionTypeHorizontal:
			expanderWidth = (width - contentWidth) / expanderCount
		}

		expanderWidth = d2common.MaxInt(0, expanderWidth)
		expanderHeight = d2common.MaxInt(0, expanderHeight)
	}

	var offsetX, offsetY int
	for _, entry := range l.entries {
		if !entry.widget.isVisible() {
			continue
		}

		if entry.widget.isExpanding() {
			entry.width, entry.height = expanderWidth, expanderHeight
		} else {
			entry.width, entry.height = entry.widget.getSize()
		}

		switch l.positionType {
		case PositionTypeVertical:
			entry.y = offsetY
			offsetY += entry.height
			switch l.horizontalAlign {
			case HorizontalAlignLeft:
				entry.x = 0
			case HorizontalAlignCenter:
				entry.x = width/2 - entry.width/2
			case HorizontalAlignRight:
				entry.x = width - entry.width
			}
		case PositionTypeHorizontal:
			entry.x = offsetX
			offsetX += entry.width
			switch l.verticalAlign {
			case VerticalAlignTop:
				entry.y = 0
			case VerticalAlignMiddle:
				entry.y = height/2 - entry.height/2
			case VerticalAlignBottom:
				entry.y = height - entry.height
			}
		case PositionTypeAbsolute:
			entry.x, entry.y = entry.widget.getPosition()
		}

		entry.widget.setOffset(offsetX, offsetY)
	}
}
