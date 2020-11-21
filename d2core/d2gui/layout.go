package d2gui

import (
	"errors"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

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

// static check that Layout implements a widget
var _ widget = &Layout{}

// Layout is a gui element container which will automatically position/align gui elements.
// Layouts are gui elements as well, so they can be nested in other layouts.
type Layout struct {
	widgetBase

	renderer     d2interface.Renderer
	assetManager *d2asset.AssetManager

	width           int
	height          int
	verticalAlign   VerticalAlign
	horizontalAlign HorizontalAlign
	positionType    PositionType
	entries         []*layoutEntry
}

// CreateLayout creates a new  GUI layout
func CreateLayout(renderer d2interface.Renderer, positionType PositionType, assetManager *d2asset.AssetManager) *Layout {
	layout := &Layout{
		renderer:     renderer,
		positionType: positionType,
		assetManager: assetManager,
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
	layout := CreateLayout(l.renderer, positionType, l.assetManager)
	l.entries = append(l.entries, &layoutEntry{widget: layout})

	return layout
}

// AddLayoutFromSource adds a nested layout to this layout, given a position type.
// Returns a pointer to the nested layout
func (l *Layout) AddLayoutFromSource(source *Layout) *Layout {
	l.entries = append(l.entries, &layoutEntry{widget: source})

	return source
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
	sprite, err := createSprite(imagePath, palettePath, l.assetManager)
	if err != nil {
		return nil, err
	}

	l.entries = append(l.entries, &layoutEntry{widget: sprite})

	return sprite, nil
}

// AddAnimatedSprite given a path, palette, and direction will add an animated
// sprite as a layout entry
func (l *Layout) AddAnimatedSprite(imagePath, palettePath string, direction AnimationDirection) (*AnimatedSprite, error) {
	sprite, err := createAnimatedSprite(imagePath, palettePath, direction, l.assetManager)
	if err != nil {
		return nil, err
	}

	l.entries = append(l.entries, &layoutEntry{widget: sprite})

	return sprite, nil
}

// AddLabel given a string and a FontStyle, adds a text label as a layout entry
func (l *Layout) AddLabel(text string, fontStyle FontStyle) (*Label, error) {
	font, err := l.loadFont(fontStyle)
	if err != nil {
		return nil, err
	}

	label, err := createLabel(l.renderer, text, font, d2util.Color(ColorWhite))
	if err != nil {
		return nil, err
	}

	l.entries = append(l.entries, &layoutEntry{widget: label})

	return label, nil
}

// AddLabelWithColor given a string and a FontStyle and a Color, adds a text label as a layout entry
func (l *Layout) AddLabelWithColor(text string, fontStyle FontStyle, col color.RGBA) (*Label, error) {
	font, err := l.loadFont(fontStyle)
	if err != nil {
		return nil, err
	}

	label, err := createLabel(l.renderer, text, font, col)
	if err != nil {
		return nil, err
	}

	l.entries = append(l.entries, &layoutEntry{widget: label})

	return label, nil
}

// AddButton given a string and ButtonStyle, adds a button as a layout entry
func (l *Layout) AddButton(text string, buttonStyle ButtonStyle) (*Button, error) {
	button, err := l.createButton(l.renderer, text, buttonStyle)
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

func (l *Layout) render(target d2interface.Surface) {
	l.AdjustEntryPlacement()

	for _, entry := range l.entries {
		if !entry.widget.isVisible() {
			continue
		}

		l.renderEntry(entry, target)

		if layoutDebug {
			l.renderEntryDebug(entry, target)
		}
	}
}

func (l *Layout) advance(elapsed float64) error {
	for _, entry := range l.entries {
		if err := entry.widget.advance(elapsed); err != nil {
			return err
		}
	}

	return nil
}

func (l *Layout) renderEntry(entry *layoutEntry, target d2interface.Surface) {
	target.PushTranslation(entry.x, entry.y)
	defer target.Pop()

	entry.widget.render(target)
}

func (l *Layout) renderEntryDebug(entry *layoutEntry, target d2interface.Surface) {
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
}

func (l *Layout) getContentSize() (width, height int) {
	for _, entry := range l.entries {
		x, y := entry.widget.getPosition()
		w, h := entry.widget.getSize()

		switch l.positionType {
		case PositionTypeVertical:
			width = d2math.MaxInt(width, w)
			height += h
		case PositionTypeHorizontal:
			width += w
			height = d2math.MaxInt(height, h)
		case PositionTypeAbsolute:
			width = d2math.MaxInt(width, x+w)
			height = d2math.MaxInt(height, y+h)
		}
	}

	return width, height
}

// GetSize returns the layout width and height
func (l *Layout) GetSize() (width, height int) {
	return l.getSize()
}

func (l *Layout) getSize() (width, height int) {
	width, height = l.getContentSize()
	return d2math.MaxInt(width, l.width), d2math.MaxInt(height, l.height)
}

func (l *Layout) onMouseButtonDown(event d2interface.MouseEvent) bool {
	for _, entry := range l.entries {
		if entry.IsIn(event) {
			entry.widget.onMouseButtonClick(event)
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

		expanderWidth = d2math.MaxInt(0, expanderWidth)
		expanderHeight = d2math.MaxInt(0, expanderHeight)
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

func (l *Layout) createButton(renderer d2interface.Renderer, text string,
	buttonStyle ButtonStyle) (*Button,
	error) {
	config := getButtonStyleConfig(buttonStyle)
	if config == nil {
		return nil, errors.New("invalid button style")
	}

	animation, loadErr := l.assetManager.LoadAnimation(config.animationPath, config.palettePath)
	if loadErr != nil {
		return nil, loadErr
	}

	var buttonWidth int

	for i := 0; i < config.segmentsX; i++ {
		w, _, err := animation.GetFrameSize(i)
		if err != nil {
			return nil, err
		}

		buttonWidth += w
	}

	var buttonHeight int

	for i := 0; i < config.segmentsY; i++ {
		_, h, err := animation.GetFrameSize(i * config.segmentsY)
		if err != nil {
			return nil, err
		}

		buttonHeight += h
	}

	font, loadErr := l.loadFont(config.fontStyle)
	if loadErr != nil {
		return nil, loadErr
	}

	textColor := rgbaColor(grey)
	textWidth, textHeight := font.GetTextMetrics(text)
	textX := half(buttonWidth) - half(textWidth)
	textY := half(buttonHeight) - half(textHeight) + config.textOffset

	surfaceCount := animation.GetFrameCount() / (config.segmentsX * config.segmentsY)
	surfaces := make([]d2interface.Surface, surfaceCount)

	for i := 0; i < surfaceCount; i++ {
		surface := renderer.NewSurface(buttonWidth, buttonHeight)

		segX, segY, frame := config.segmentsX, config.segmentsY, i
		if segErr := renderSegmented(animation, segX, segY, frame, surface); segErr != nil {
			return nil, segErr
		}

		font.SetColor(textColor)

		var textOffsetX, textOffsetY int

		switch buttonState(i) {
		case buttonStatePressed, buttonStatePressedToggled:
			textOffsetX = -2
			textOffsetY = 2
		}

		surface.PushTranslation(textX+textOffsetX, textY+textOffsetY)
		surfaceErr := font.RenderText(text, surface)
		surface.Pop()

		if surfaceErr != nil {
			return nil, surfaceErr
		}

		surfaces[i] = surface
	}

	button := &Button{width: buttonWidth, height: buttonHeight, surfaces: surfaces}
	button.SetVisible(true)

	return button, nil
}

func (l *Layout) loadFont(fontStyle FontStyle) (*d2asset.Font, error) {
	config := getFontStyleConfig(fontStyle)
	if config == nil {
		return nil, errors.New("invalid font style")
	}

	return l.assetManager.LoadFont(config.fontBasePath+".tbl", config.fontBasePath+".dc6", config.palettePath)
}
