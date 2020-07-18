package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
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

const layoutDebug = false // turns on debug rendering stuff for layouts

const (
	white   = 0xffffff_ff
	magenta = 0xff00ff_ff
	grey2   = 0x808080_ff
	green   = 0x0000ff_ff
	yellow  = 0xffff00_ff
)

// VerticalAlign type, determines alignment along y-axis within a layout
type VerticalAlign int

// VerticalAlign types
const (
	VerticalAlignTop VerticalAlign = iota
	VerticalAlignMiddle
	VerticalAlignBottom
)

// HorizontalAlign type, determines alignment along x-axis within a layout
type HorizontalAlign int

// Horizontal alignment types
const (
	HorizontalAlignLeft HorizontalAlign = iota
	HorizontalAlignCenter
	HorizontalAlignRight
)

// PositionType determines layout positioning
type PositionType int

// Positioning types
const (
	PositionTypeAbsolute PositionType = iota
	PositionTypeVertical
	PositionTypeHorizontal
)

// Layout is a gui element container which will automatically position/align gui elements.
// Layouts are gui elements as well, so they can be nested in other layouts.
type Layout struct {
	widgetBase

	renderer d2interface.Renderer

	width           int
	height          int
	verticalAlign   VerticalAlign
	horizontalAlign HorizontalAlign
	positionType    PositionType
	entries         []*layoutEntry
}

func createLayout(renderer d2interface.Renderer, positionType PositionType) *Layout {
	layout := &Layout{
		renderer:     renderer,
		positionType: positionType,
	}

	layout.SetVisible(true)

	return layout
}

// SetSize sets the size of the layout
func (l *Layout) SetSize(width, height int) {
	l.width = width
	l.height = height
}

// SetVerticalAlign sets the vertical alignment type of the layout
func (l *Layout) SetVerticalAlign(verticalAlign VerticalAlign) {
	l.verticalAlign = verticalAlign
}

// SetHorizontalAlign sets the horizontal alignment type of the layout
func (l *Layout) SetHorizontalAlign(horizontalAlign HorizontalAlign) {
	l.horizontalAlign = horizontalAlign
}

// AddLayout adds a nested layout to this layout, given a position type.
// Returns a pointer to the nested layout
func (l *Layout) AddLayout(positionType PositionType) *Layout {
	layout := createLayout(l.renderer, positionType)
	l.entries = append(l.entries, &layoutEntry{widget: layout})

	return layout
}

// AddSpacerStatic adds a spacer with explicitly defined height and width
func (l *Layout) AddSpacerStatic(width, height int) *SpacerStatic {
	spacer := createSpacerStatic(width, height)

	l.entries = append(l.entries, &layoutEntry{widget: spacer})

	return spacer
}

// AddSpacerDynamic adds a spacer which has dynamic width and height. The width
// and height are computed based off of the position/alignment type of the layout
// and the dimensions/positions of the layout entries.
func (l *Layout) AddSpacerDynamic() *SpacerDynamic {
	spacer := createSpacerDynamic()

	l.entries = append(l.entries, &layoutEntry{widget: spacer})

	return spacer
}

// AddSprite given a path and palette, adds a Sprite as a layout entry
func (l *Layout) AddSprite(imagePath, palettePath string) (*Sprite, error) {
	sprite, err := createSprite(imagePath, palettePath)
	if err != nil {
		return nil, err
	}

	l.entries = append(l.entries, &layoutEntry{widget: sprite})

	return sprite, nil
}

// AddAnimatedSprite given a path, palette, and direction will add an animated
// sprite as a layout entry
func (l *Layout) AddAnimatedSprite(imagePath, palettePath string, direction AnimationDirection) (*AnimatedSprite, error) {
	sprite, err := createAnimatedSprite(imagePath, palettePath, direction)
	if err != nil {
		return nil, err
	}

	l.entries = append(l.entries, &layoutEntry{widget: sprite})

	return sprite, nil
}

// AddLabel given a string and a FontStyle, adds a text label as a layout entry
func (l *Layout) AddLabel(text string, fontStyle FontStyle) (*Label, error) {
	label, err := createLabel(l.renderer, text, fontStyle)
	if err != nil {
		return nil, err
	}

	l.entries = append(l.entries, &layoutEntry{widget: label})

	return label, nil
}

// AddButton given a string and ButtonStyle, adds a button as a layout entry
func (l *Layout) AddButton(text string, buttonStyle ButtonStyle) (*Button, error) {
	button, err := createButton(l.renderer, text, buttonStyle)
	if err != nil {
		return nil, err
	}

	l.entries = append(l.entries, &layoutEntry{widget: button})

	return button, nil
}

// Clear removes all layout entries
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

		if layoutDebug {
			if err := l.renderEntryDebug(entry, target); err != nil {
				return err
			}
		}
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

	drawColor := rgbaColor(white)
	switch entry.widget.(type) {
	case *Layout:
		drawColor = rgbaColor(magenta)
	case *SpacerStatic, *SpacerDynamic:
		drawColor = rgbaColor(grey2)
	case *Label:
		drawColor = rgbaColor(green)
	case *Button:
		drawColor = rgbaColor(yellow)
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

func (l *Layout) getContentSize() (width, height int) {
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

func (l *Layout) getSize() (width, height int) {
	width, height = l.getContentSize()
	return d2common.MaxInt(width, l.width), d2common.MaxInt(height, l.height)
}

func (l *Layout) onMouseButtonDown(event d2interface.MouseEvent) bool {
	for _, entry := range l.entries {
		if entry.IsIn(event) {
			entry.widget.onMouseButtonDown(event)
			entry.mouseDown[event.Button()] = true
		}
	}

	return false
}

func (l *Layout) onMouseButtonUp(event d2interface.MouseEvent) bool {
	for _, entry := range l.entries {
		if entry.IsIn(event) {
			if entry.mouseDown[event.Button()] {
				entry.widget.onMouseButtonClick(event)
				entry.widget.onMouseButtonUp(event)
			}
		}

		entry.mouseDown[event.Button()] = false
	}

	return false
}

func (l *Layout) onMouseMove(event d2interface.MouseMoveEvent) bool {
	for _, entry := range l.entries {
		if entry.IsIn(event) {
			entry.widget.onMouseMove(event)

			if entry.mouseOver {
				entry.widget.onMouseOver(event)
			} else {
				entry.widget.onMouseEnter(event)
			}

			entry.mouseOver = true
		} else if entry.mouseOver {
			entry.widget.onMouseLeave(event)
			entry.mouseOver = false
		}
	}

	return false
}

// AdjustEntryPlacement calculates and sets the position for all layout entries.
// This is based on the position/horizontal/vertical alignment type set, as well as the
// expansion types of spacers.
func (l *Layout) AdjustEntryPlacement() {
	width, height := l.getSize()

	var expanderCount, expanderWidth, expanderHeight int

	for _, entry := range l.entries {
		if entry.widget.isVisible() && entry.widget.isExpanding() {
			expanderCount++
		}
	}

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

	for idx := range l.entries {
		entry := l.entries[idx]
		if !entry.widget.isVisible() {
			continue
		}

		if entry.widget.isExpanding() {
			entry.width, entry.height = expanderWidth, expanderHeight
		} else {
			entry.width, entry.height = entry.widget.getSize()
		}

		l.handleEntryPosition(offsetX, offsetY, entry)

		switch l.positionType {
		case PositionTypeVertical:
			offsetY += entry.height
		case PositionTypeHorizontal:
			offsetX += entry.width
		}

		sx, sy := l.ScreenPos()
		entry.widget.SetScreenPos(entry.x+sx, entry.y+sy)
		entry.widget.setOffset(offsetX, offsetY)
	}
}

func (l *Layout) handleEntryPosition(offsetX, offsetY int, entry *layoutEntry) {
	width, height := l.getSize()

	switch l.positionType {
	case PositionTypeVertical:
		entry.y = offsetY
		l.handleEntryVerticalAlign(width, entry)

	case PositionTypeHorizontal:
		entry.x = offsetX
		l.handleEntryHorizontalAlign(height, entry)

	case PositionTypeAbsolute:
		entry.x, entry.y = entry.widget.getPosition()
	}
}

func (l *Layout) handleEntryHorizontalAlign(height int, entry *layoutEntry) {
	switch l.verticalAlign {
	case VerticalAlignTop:
		entry.y = 0
	case VerticalAlignMiddle:
		entry.y = half(height) - half(entry.height)
	case VerticalAlignBottom:
		entry.y = height - entry.height
	}
}

func (l *Layout) handleEntryVerticalAlign(width int, entry *layoutEntry) {
	switch l.horizontalAlign {
	case HorizontalAlignLeft:
		entry.x = 0
	case HorizontalAlignCenter:
		entry.x = half(width) - half(entry.width)
	case HorizontalAlignRight:
		entry.x = width - entry.width
	}
}

// IsIn layout entry, spc. of an event.
func (l *layoutEntry) IsIn(event d2interface.HandlerEvent) bool {
	sx, sy := l.widget.ScreenPos()
	rect := d2common.Rectangle{Left: sx, Top: sy, Width: l.width, Height: l.height}

	return rect.IsInRect(event.X(), event.Y())
}
