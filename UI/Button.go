package UI

import (
	"image/color"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/Palettes"
	"github.com/essial/OpenDiablo2/ResourcePaths"
	"github.com/hajimehoshi/ebiten"
)

// Button defines a standard wide UI button
type Button struct {
	enabled       bool
	x, y          int
	width, height uint32
	visible       bool
	pressed       bool
	fileProvider  Common.FileProvider
	normalImage   *ebiten.Image
	pressedImage  *ebiten.Image
}

// CreateButton creates an instance of Button
func CreateButton(fileProvider Common.FileProvider, text string) *Button {
	result := &Button{
		fileProvider: fileProvider,
		width:        272,
		height:       35,
		visible:      true,
		enabled:      true,
		pressed:      false,
	}
	font := GetFont(ResourcePaths.FontExocet10, Palettes.Units, fileProvider)
	result.normalImage, _ = ebiten.NewImage(272, 35, ebiten.FilterNearest)
	result.pressedImage, _ = ebiten.NewImage(272, 35, ebiten.FilterNearest)
	textWidth, textHeight := font.GetTextMetrics(text)
	textX := (272 / 2) - (textWidth / 2)
	textY := (35 / 2) - (textHeight / 2) + 5
	buttonSprite := fileProvider.LoadSprite(ResourcePaths.WideButtonBlank, Palettes.Units)
	buttonSprite.MoveTo(0, 0)
	buttonSprite.Blend = true
	buttonSprite.DrawSegments(result.normalImage, 2, 1, 0)
	font.Draw(int(textX), int(textY), text, color.RGBA{100, 100, 100, 255}, result.normalImage)
	buttonSprite.DrawSegments(result.pressedImage, 2, 1, 1)
	font.Draw(int(textX-2), int(textY+2), text, color.Black, result.pressedImage)
	return result
}

// Draw renders the button
func (v *Button) Draw(target *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{
		CompositeMode: ebiten.CompositeModeSourceAtop,
		Filter:        ebiten.FilterNearest,
	}
	opts.GeoM.Translate(float64(v.x), float64(v.y))
	if v.pressed {
		target.DrawImage(v.pressedImage, opts)
		return
	}
	target.DrawImage(v.normalImage, opts)
}

// GetEnabled returns the enabled state
func (v *Button) GetEnabled() bool {
	return v.enabled
}

// SetEnabled sets the enabled state
func (v *Button) SetEnabled(enabled bool) {
	v.enabled = enabled
}

// GetSize returns the size of the button
func (v *Button) GetSize() (uint32, uint32) {
	return v.width, v.height
}

// MoveTo moves the button
func (v *Button) MoveTo(x, y int) {
	v.x = x
	v.y = y
}

// GetLocation returns the location of the button
func (v *Button) GetLocation() (x, y int) {
	return v.x, v.y
}

// GetVisible returns the visibility of the button
func (v *Button) GetVisible() bool {
	return v.visible
}

// SetVisible sets the visibility of the button
func (v *Button) SetVisible(visible bool) {
	v.visible = visible
}
