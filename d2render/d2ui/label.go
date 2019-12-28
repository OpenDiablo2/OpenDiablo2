package d2ui

import (
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2surface"
	"github.com/hajimehoshi/ebiten"
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
	imageData *ebiten.Image
	Color     color.Color
}

// CreateLabel creates a new instance of a UI label
func CreateLabel(fontPath, palettePath string) Label {
	result := Label{
		Alignment: LabelAlignLeft,
		Color:     color.White,
		font:      GetFont(fontPath, palettePath),
	}
	return result
}

// Render draws the label on the screen
func (v *Label) Render(target *d2surface.Surface) {
	if len(v.text) == 0 {
		return
	}
	v.cacheImage()

	x, y := v.X, v.Y
	if v.Alignment == LabelAlignCenter {
		x, y = v.X-int(v.Width/2), v.Y
	} else if v.Alignment == LabelAlignRight {
		x, y = v.X-int(v.Width), v.Y
	}

	target.PushFilter(ebiten.FilterNearest)
	target.PushCompositeMode(ebiten.CompositeModeSourceAtop)
	target.PushTranslation(x, y)
	defer target.PopN(3)

	target.Render(v.imageData)
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
	if v.imageData != nil {
		return
	}
	width, height := v.font.GetTextMetrics(v.text)
	v.Width = width
	v.Height = height
	v.imageData, _ = ebiten.NewImage(int(width), int(height), ebiten.FilterNearest)
	surface := d2surface.CreateSurface(v.imageData)
	v.font.Render(0, 0, v.text, v.Color, surface)
}

// SetText sets the label's text
func (v *Label) SetText(newText string) {
	if v.text == newText {
		return
	}
	v.text = newText
	v.imageData = nil
}

// GetSize returns the size of the label
func (v Label) GetSize() (width, height int) {
	v.cacheImage()
	width = v.Width
	height = v.Height
	return
}
