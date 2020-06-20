package d2ui

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

// LabelAlignment represents a label's alignment
type LabelAlignment uint8

const (
	// LabelAlignLeft represents a left-aligned label
	LabelAlignLeft LabelAlignment = 0
	// LabelAlignCenter represents a center-aligned label
	LabelAlignCenter LabelAlignment = 1
	// LabelAlignRight represents a right-aligned label
	LabelAlignRight LabelAlignment = 2
)

// Label represents a user interface label
type Label struct {
	text      string
	X         int
	Y         int
	Width     int
	Height    int
	Alignment LabelAlignment
	font      *Font
	imageData map[string]d2interface.Surface
	Color     color.Color
}

// CreateLabel creates a new instance of a UI label
func CreateLabel(fontPath, palettePath string) Label {
	result := Label{
		Alignment: LabelAlignLeft,
		Color:     color.White,
		font:      GetFont(fontPath, palettePath),
		imageData: make(map[string]d2interface.Surface),
	}

	return result
}

// Render draws the label on the screen
func (v *Label) Render(target d2interface.Surface) {
	if len(v.text) == 0 {
		return
	}
	v.cacheImage()

	x, y := v.X, v.Y
	if v.Alignment == LabelAlignCenter {
		x, y = v.X-v.Width/2, v.Y
	} else if v.Alignment == LabelAlignRight {
		x, y = v.X-v.Width, v.Y
	}

	target.PushFilter(d2interface.FilterNearest)
	target.PushCompositeMode(d2enum.CompositeModeSourceAtop)
	target.PushTranslation(x, y)
	defer target.PopN(3)

	target.Render(v.imageData[v.text])
}

// SetPosition moves the label to the specified location
func (v *Label) SetPosition(x, y int) {
	v.X = x
	v.Y = y
}

func (v *Label) GetTextMetrics(text string) (width, height int) {
	return v.font.GetTextMetrics(text)
}

func (v *Label) cacheImage() {
	if v.imageData[v.text] != nil {
		return
	}
	width, height := v.font.GetTextMetrics(v.text)
	v.Width = width
	v.Height = height
	v.imageData[v.text], _ = d2render.NewSurface(width, height, d2interface.FilterNearest)
	surface, _ := d2render.CreateSurface(v.imageData[v.text])
	v.font.Render(0, 0, v.text, v.Color, surface)
}

// SetText sets the label's text
func (v *Label) SetText(newText string) {
	if v.text == newText {
		return
	}
	v.text = newText
}

// Size returns the size of the label
func (v Label) GetSize() (width, height int) {
	v.cacheImage()
	width = v.Width
	height = v.Height
	return
}
