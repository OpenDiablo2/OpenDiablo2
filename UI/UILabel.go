package UI

import (
	"image/color"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
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
	Width     uint32
	Height    uint32
	Alignment LabelAlignment
	font      *Font
	imageData *ebiten.Image
	Color     color.Color
}

// CreateLabel creates a new instance of a UI label
func CreateLabel(provider Common.FileProvider, font string, palette Palettes.Palette) *Label {
	result := &Label{
		Alignment: LabelAlignLeft,
		Color:     color.White,
		font:      GetFont(font, palette, provider),
	}

	return result
}

// Draw draws the label on the screen
func (v *Label) Draw(target *ebiten.Image) {
	if len(v.text) == 0 {
		return
	}
	v.cacheImage()
	opts := &ebiten.DrawImageOptions{}

	if v.Alignment == LabelAlignCenter {
		opts.GeoM.Translate(float64(v.X-int(v.Width/2)), float64(v.Y))
	} else if v.Alignment == LabelAlignRight {
		opts.GeoM.Translate(float64(v.X-int(v.Width)), float64(v.Y))
	} else {
		opts.GeoM.Translate(float64(v.X), float64(v.Y))
	}
	opts.CompositeMode = ebiten.CompositeModeSourceAtop
	opts.Filter = ebiten.FilterNearest
	target.DrawImage(v.imageData, opts)
}

// MoveTo moves the label to the specified location
func (v *Label) MoveTo(x, y int) {
	v.X = x
	v.Y = y
}

func (v *Label) cacheImage() {
	if v.imageData != nil {
		return
	}
	width, height := v.font.GetTextMetrics(v.text)
	v.Width = width
	v.Height = height
	v.imageData, _ = ebiten.NewImage(int(width), int(height), ebiten.FilterNearest)
	v.font.Draw(0, 0, v.text, v.Color, v.imageData)
}

// SetText sets the label's text
func (v *Label) SetText(newText string) {
	if v.text == newText {
		return
	}
	v.text = newText
	v.imageData = nil
}
