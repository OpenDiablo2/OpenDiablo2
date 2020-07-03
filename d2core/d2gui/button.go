package d2gui

import (
	"errors"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type buttonState int

const (
	buttonStateDefault buttonState = iota
	buttonStatePressed
	buttonStateToggled
	buttonStatePressedToggled
)

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

	animation, err := d2asset.LoadAnimation(config.animationPath, config.palettePath)
	if err != nil {
		return nil, err
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

	font, err := loadFont(config.fontStyle)
	if err != nil {
		return nil, err
	}

	textColor := color.RGBA{R: 0x64, G: 0x64, B: 0x64, A: 0xff}
	textWidth, textHeight := font.GetTextMetrics(text)
	textX := buttonWidth/2 - textWidth/2
	textY := buttonHeight/2 - textHeight/2 + config.textOffset

	surfaceCount := animation.GetFrameCount() / (config.segmentsX * config.segmentsY)
	surfaces := make([]d2interface.Surface, surfaceCount)
	for i := 0; i < surfaceCount; i++ {
		surface, err := renderer.NewSurface(buttonWidth, buttonHeight, d2interface.FilterNearest)
		if err != nil {
			return nil, err
		}

		if err := renderSegmented(animation, config.segmentsX, config.segmentsY, i, surface); err != nil {
			return nil, err
		}

		font.SetColor(textColor)

		var textOffsetX, textOffsetY int
		switch buttonState(i) {
		case buttonStatePressed, buttonStatePressedToggled:
			textOffsetX = -2
			textOffsetY = 2
		}

		surface.PushTranslation(textX+textOffsetX, textY+textOffsetY)
		err = font.RenderText(text, surface)
		surface.Pop()

		if err != nil {
			return nil, err
		}

		surfaces[i] = surface
	}

	button := &Button{width: buttonWidth, height: buttonHeight, surfaces: surfaces}
	button.SetVisible(true)

	return button, nil
}

func (b *Button) onMouseButtonDown(event d2interface.MouseEvent) bool {
	b.state = buttonStatePressed
	return false
}

func (b *Button) onMouseButtonUp(event d2interface.MouseEvent) bool {
	b.state = buttonStateDefault
	return false
}

func (b *Button) onMouseLeave(event d2interface.MouseMoveEvent) bool {
	b.state = buttonStateDefault
	return false
}

func (b *Button) render(target d2interface.Surface) error {
	return target.Render(b.surfaces[b.state])
}

func (b *Button) getSize() (int, int) {
	return b.width, b.height
}
