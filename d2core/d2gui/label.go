package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Label struct {
	widgetBase

	surface d2render.Surface
}

func createLabel(text string, fontStyle FontStyle) (*Label, error) {
	font, err := loadFont(fontStyle)
	if err != nil {
		return nil, err
	}

	width, height := font.GetTextMetrics(text)
	surface, err := d2render.NewSurface(width, height, d2render.FilterNearest)
	if err != nil {
		return nil, err
	}

	if err := font.RenderText(text, surface); err != nil {
		return nil, err
	}

	label := &Label{surface: surface}
	label.SetVisible(true)

	return label, nil
}

func (l *Label) render(target d2render.Surface) error {
	return target.Render(l.surface)
}

func (l *Label) getSize() (int, int) {
	return l.surface.GetSize()
}
