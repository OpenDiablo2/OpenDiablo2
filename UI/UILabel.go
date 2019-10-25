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
	ColorMod  color.Color
}

// CreateLabel creates a new instance of a UI label
func CreateLabel(provider Common.FileProvider, font string, palette Palettes.Palette) *Label {
	result := &Label{
		Alignment: LabelAlignLeft,
		ColorMod:  nil,
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

func (v *Label) calculateSize() (uint32, uint32) {
	width := uint32(0)
	height := uint32(0)
	for _, ch := range v.text {
		metric := v.font.Metrics[uint8(ch)]
		width += uint32(metric.Width)
		height = Common.Max(height, uint32(metric.Height))
	}
	return width, height
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
	width, height := v.calculateSize()
	v.Width = width
	v.Height = height
	v.imageData, _ = ebiten.NewImage(int(width), int(height), ebiten.FilterNearest)
	x := uint32(0)
	v.font.FontSprite.ColorMod = v.ColorMod
	for _, ch := range v.text {
		char := uint8(ch)
		metric := v.font.Metrics[char]
		v.font.FontSprite.Frame = char
		v.font.FontSprite.MoveTo(int(x), int(height))
		v.font.FontSprite.Draw(v.imageData)
		x += uint32(metric.Width)
	}
}

// SetText sets the label's text
func (v *Label) SetText(newText string) {
	if v.text == newText {
		return
	}
	v.text = newText
	v.imageData = nil
}
