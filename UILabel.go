package OpenDiablo2

import (
	"image/color"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
	"github.com/hajimehoshi/ebiten"
)

// UILabelAlignment represents a label's alignment
type UILabelAlignment uint8

const (
	// UILabelAlignLeft represents a left-aligned label
	UILabelAlignLeft UILabelAlignment = 0
	// UILabelAlignCenter represents a center-aligned label
	UILabelAlignCenter UILabelAlignment = 1
	// UILabelAlignRight represents a right-aligned label
	UILabelAlignRight UILabelAlignment = 2
)

// UILabel represents a user interface label
type UILabel struct {
	text      string
	X         int
	Y         int
	Width     uint32
	Height    uint32
	Alignment UILabelAlignment
	font      *MPQFont
	imageData *ebiten.Image
	ColorMod  color.Color
}

// CreateUILabel creates a new instance of a UI label
func CreateUILabel(engine *Engine, font string, palette Palettes.Palette) *UILabel {
	result := &UILabel{
		Alignment: UILabelAlignLeft,
		ColorMod:  nil,
		font:      engine.GetFont(font, palette),
	}

	return result
}

// Draw draws the label on the screen
func (v *UILabel) Draw(target *ebiten.Image) {
	if len(v.text) == 0 {
		return
	}
	v.cacheImage()
	opts := &ebiten.DrawImageOptions{}

	if v.Alignment == UILabelAlignCenter {
		opts.GeoM.Translate(float64(v.X-int(v.Width/2)), float64(v.Y))
	} else if v.Alignment == UILabelAlignRight {
		opts.GeoM.Translate(float64(v.X-int(v.Width)), float64(v.Y))
	} else {
		opts.GeoM.Translate(float64(v.X), float64(v.Y))
	}
	opts.CompositeMode = ebiten.CompositeModeSourceAtop
	opts.Filter = ebiten.FilterNearest
	target.DrawImage(v.imageData, opts)
}

func (v *UILabel) calculateSize() (uint32, uint32) {
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
func (v *UILabel) MoveTo(x, y int) {
	v.X = x
	v.Y = y
}

func (v *UILabel) cacheImage() {
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
func (v *UILabel) SetText(newText string) {
	if v.text == newText {
		return
	}
	v.text = newText
	v.imageData = nil
}
