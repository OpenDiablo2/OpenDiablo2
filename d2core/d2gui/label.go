package d2gui

import (
	"image/color"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// Constants defining the main shades of basic colors
// found in the game
const (
	ColorWhite = 0xffffffff
	ColorRed   = 0xdb3f3dff
	ColorGreen = 0x00d000ff
	ColorBlue  = 0x5450d1ff
	ColorBrown = 0xa1925dff
	ColorGrey  = 0x555555ff
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

func createLabel(renderer d2interface.Renderer, text string, font *d2asset.Font, col color.RGBA) (*Label, error) {
	label := &Label{
		font:       font,
		renderer:   renderer,
		color:      col,
		hoverColor: col,
	}

	err := label.setText(text)
	if err != nil {
		return nil, err
	}

	label.SetVisible(true)

	return label, nil
}

// SetHoverColor will set the value of hoverColor
func (l *Label) SetHoverColor(col color.RGBA) {
	l.hoverColor = col
}

// SetIsBlinking will set the isBlinking value
func (l *Label) SetIsBlinking(isBlinking bool) {
	l.isBlinking = isBlinking
}

// SetIsHovered will set the isHovered value
func (l *Label) SetIsHovered(isHovered bool) error {
	l.isHovered = isHovered

	return l.setText(l.text)
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
