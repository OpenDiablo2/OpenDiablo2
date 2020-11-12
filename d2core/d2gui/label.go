package d2gui

import (
	"image/color"
	"log"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// Label is renderable text
type Label struct {
	widgetBase

	renderer    d2interface.Renderer
	text        string
	font        *d2asset.Font
	surface     d2interface.Surface
	color       color.RGBA
	hoverColor  color.RGBA
	isHovered   bool
	isBlinking  bool
	isDisplayed bool
	blinkTimer  time.Time
}

func createLabel(renderer d2interface.Renderer, text string, font *d2asset.Font, col color.RGBA) *Label {
	label := &Label{
		font:       font,
		renderer:   renderer,
		color:      col,
		hoverColor: col,
	}

	err := label.setText(text)
	if err != nil {
		log.Print(err)
		return nil
	}

	label.SetVisible(true)

	return label
}

func (l *Label) SetHoverColor(col color.RGBA) {
	l.hoverColor = col
}

func (l *Label) SetIsBlinking(isBlinking bool) {
	l.isBlinking = isBlinking
}
func (l *Label) SetIsHovered(isHovered bool) {
	l.isHovered = isHovered
	l.setText(l.text)
}

func (l *Label) render(target d2interface.Surface) {
	if l.isBlinking && time.Since(l.blinkTimer) >= 200*time.Millisecond {
		l.isDisplayed = !l.isDisplayed
		l.blinkTimer = time.Now()
	}

	if l.isBlinking && !l.isDisplayed {
		return
	}

	target.Render(l.surface)
}

func (l *Label) getSize() (width, height int) {
	return l.surface.GetSize()
}

// GetText returns the label text
func (l *Label) GetText() string {
	return l.text
}

// SetColor sets the label text
func (l *Label) SetColor(col color.RGBA) error {
	l.color = col
	return l.setText(l.text)
}

// SetText sets the label text
func (l *Label) SetText(text string) error {
	if text == l.text {
		return nil
	}

	return l.setText(text)
}

func (l *Label) setText(text string) error {
	width, height := l.font.GetTextMetrics(text)

	surface := l.renderer.NewSurface(width, height)

	col := l.color
	if l.isHovered {
		col = l.hoverColor
	}

	l.font.SetColor(col)
	if err := l.font.RenderText(text, surface); err != nil {
		return err
	}

	l.surface = surface
	l.text = text

	return nil
}
