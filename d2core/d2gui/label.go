package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Label struct {
	widgetBase

	text    string
	font    *d2asset.Font
	surface d2render.Surface
}

func createLabel(text string, fontStyle FontStyle) *Label {
	font, _ := loadFont(fontStyle)
	label := &Label{font: font}
	label.SetText(text)
	label.visible = true
	return label
}

func (l *Label) SetText(text string) *Label {
	l.text = text
	l.cache()
	return l
}

func (l *Label) render(target d2render.Surface) error {
	if l.surface == nil {
		return nil
	}

	return target.Render(l.surface)
}

func (l *Label) cache() error {
	l.surface = nil
	if l.font == nil {
		return nil
	}

	width, height := l.font.GetTextMetrics(l.text)

	var err error
	if l.surface, err = d2render.NewSurface(width, height, d2render.FilterNearest); err != nil {
		return err
	}

	return l.font.RenderText(l.text, l.surface)
}

func (l *Label) getSize() (int, int) {
	if l.surface == nil {
		return 0, 0
	}

	return l.surface.GetSize()
}
