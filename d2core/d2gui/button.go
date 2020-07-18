package d2gui

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type buttonState int

const (
	buttonStateDefault buttonState = iota
	buttonStatePressed
	buttonStatePressedToggled
)

const (
	grey = 0x404040ff
)

// Button is a user actionable drawable toggle switch
type Button struct {
	widgetBase

	width    int
	height   int
	state    buttonState
	surfaces []d2interface.Surface
}

func createButton(renderer d2interface.Renderer, text string, buttonStyle ButtonStyle) (*Button, error) {
	config, ok := buttonStyleConfigs[buttonStyle]
	if !ok {
		return nil, errors.New("invalid button style")
	}

	animation, loadErr := d2asset.LoadAnimation(config.animationPath, config.palettePath)
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

	font, loadErr := loadFont(config.fontStyle)
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
		surface, surfaceErr := renderer.NewSurface(buttonWidth, buttonHeight, d2enum.FilterNearest)
		if surfaceErr != nil {
			return nil, surfaceErr
		}

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
		surfaceErr = font.RenderText(text, surface)
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

func (b *Button) onMouseButtonDown(_ d2interface.MouseEvent) bool {
	b.state = buttonStatePressed
	return false
}

func (b *Button) onMouseButtonUp(_ d2interface.MouseEvent) bool {
	b.state = buttonStateDefault
	return false
}

func (b *Button) onMouseLeave(_ d2interface.MouseMoveEvent) bool {
	b.state = buttonStateDefault
	return false
}

func (b *Button) render(target d2interface.Surface) error {
	return target.Render(b.surfaces[b.state])
}

func (b *Button) getSize() (width, height int) {
	return b.width, b.height
}
