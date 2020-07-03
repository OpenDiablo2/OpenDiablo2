package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type Label struct {
	widgetBase

	renderer d2interface.Renderer
	text     string
	font     *d2asset.Font
	surface  d2interface.Surface
}

func createLabel(renderer d2interface.Renderer, text string, fontStyle FontStyle) (*Label, error) {
	font, err := loadFont(fontStyle)
	if err != nil {
		return nil, err
	}

	label := &Label{
		font:     font,
		renderer: renderer,
	}

	_ = label.setText(text)
	label.SetVisible(true)

	return label, nil
}

func (l *Label) render(target d2interface.Surface) error {
	return target.Render(l.surface)
}

func (l *Label) getSize() (int, int) {
	return l.surface.GetSize()
}

func (l *Label) GetText() string {
	return l.text
}

func (l *Label) SetText(text string) error {
	if text == l.text {
		return nil
	}
	return l.setText(text)
}

func (l *Label) setText(text string) error {
	width, height := l.font.GetTextMetrics(text)
	surface, err := l.renderer.NewSurface(width, height, d2interface.FilterNearest)
	if err != nil {
		return err
	}
	if err := l.font.RenderText(text, surface); err != nil {
		return err
	}
	l.surface = surface
	l.text = text
	return nil
}
