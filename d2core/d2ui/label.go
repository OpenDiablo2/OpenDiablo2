package d2ui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"image/color"
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// Label represents a user interface label
type Label struct {
	text      string
	X         int
	Y         int
	Alignment d2gui.HorizontalAlign
	font      d2interface.Font
	Color     color.Color
}

// CreateLabel creates a new instance of a UI label
func CreateLabel(fontPath, palettePath string) Label {
	font, _ := d2asset.LoadFont(fontPath+".tbl", fontPath+".dc6", palettePath)
	result := Label{
		Alignment: d2gui.HorizontalAlignLeft,
		Color:     color.White,
		font:      font,
	}

	return result
}

// Render draws the label on the screen, respliting the lines to allow for other alignments
func (v *Label) Render(target d2interface.Surface) {
	v.font.SetColor(v.Color)
	target.PushTranslation(v.X, v.Y)

	lines := strings.Split(v.text, "\n")
	yOffset := 0

	for _, line := range lines {
		lw, lh := v.GetTextMetrics(line)
		target.PushTranslation(v.getAlignOffset(lw), yOffset)

		_ = v.font.RenderText(line, target)

		yOffset += lh

		target.Pop()
	}

	target.Pop()
}

// SetPosition moves the label to the specified location
func (v *Label) SetPosition(x, y int) {
	v.X = x
	v.Y = y
}

// GetSize returns the size of the label
func (v *Label) GetSize() (width, height int) {
	return v.font.GetTextMetrics(v.text)
}

// GetTextMetrics returns the width and height of the enclosing rectangle in Pixels.
func (v *Label) GetTextMetrics(text string) (width, height int) {
	return v.font.GetTextMetrics(text)
}

// SetText sets the label's text
func (v *Label) SetText(newText string) {
	v.text = newText
}

func (v *Label) getAlignOffset(textWidth int) int {
	switch v.Alignment {
	case d2gui.HorizontalAlignLeft:
		return 0
	case d2gui.HorizontalAlignCenter:
		return -textWidth / 2
	case d2gui.HorizontalAlignRight:
		return -textWidth
	default:
		log.Fatal("Invalid Alignment")
		return 0
	}
}
