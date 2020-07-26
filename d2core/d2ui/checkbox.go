package d2ui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// Checkbox represents a checkbox UI element
type Checkbox struct {
	Image        d2interface.Surface
	checkedImage d2interface.Surface
	x            int
	y            int
	width        int
	height       int
	onClick      func()
	checkState   bool
	visible      bool
	enabled      bool
}

// CreateCheckbox creates a new instance of a checkbox
func CreateCheckbox(renderer d2interface.Renderer, checkState bool) Checkbox {
	result := Checkbox{
		checkState: checkState,
		visible:    true,
		width:      0,
		height:     0,
		enabled:    true,
	}

	animation, _ := d2asset.LoadAnimation(d2resource.Checkbox, d2resource.PaletteFechar)
	checkboxSprite, _ := LoadSprite(animation)
	result.width, result.height, _ = checkboxSprite.GetFrameSize(0)
	checkboxSprite.SetPosition(0, 0)

	result.Image, _ = renderer.NewSurface(result.width, result.height, d2enum.FilterNearest)

	_ = checkboxSprite.RenderSegmented(result.Image, 1, 1, 0)

	result.checkedImage, _ = renderer.NewSurface(result.width, result.height, d2enum.FilterNearest)

	_ = checkboxSprite.RenderSegmented(result.checkedImage, 1, 1, 1)

	return result
}

// Render renders the checkbox
func (v *Checkbox) Render(target d2interface.Surface) error {
	target.PushTranslation(v.x, v.y)
	target.PushFilter(d2enum.FilterNearest)

	defer target.PopN(2)

	if v.checkState {
		_ = target.Render(v.checkedImage)
	} else {
		_ = target.Render(v.Image)
	}

	return nil
}

// Advance does nothing for checkboxes
func (v *Checkbox) Advance(elapsed float64) {

}

// GetEnabled returns the enabled state of the checkbox
func (v *Checkbox) GetEnabled() bool {
	return v.enabled
}

// SetEnabled sets the enabled state of the checkbox
func (v *Checkbox) SetEnabled(enabled bool) {
	v.enabled = enabled
}

// SetPressed does nothing for checkboxes
func (v *Checkbox) SetPressed(_ bool) {
}

// SetCheckState sets the check state of the checkbox
func (v *Checkbox) SetCheckState(checkState bool) {
	v.checkState = checkState
}

// GetCheckState returns the check state of the checkbox
func (v *Checkbox) GetCheckState() bool {
	return v.checkState
}

// GetPressed returns the pressed state of the checkbox
func (v *Checkbox) GetPressed() bool {
	return v.checkState
}

// OnACtivated sets the callback function of the click event for the checkbox
func (v *Checkbox) OnActivated(callback func()) {
	v.onClick = callback
}

// Activate activates the checkbox
func (v *Checkbox) Activate() {
	v.checkState = !v.checkState
	if v.onClick == nil {
		return
	}
	v.onClick()
}

// GetPosition returns the position of the checkbox
func (v *Checkbox) GetPosition() (int, int) {
	return v.x, v.y
}

// GetSize returns the size of the checkbox
func (v *Checkbox) GetSize() (int, int) {
	return v.width, v.height
}

// GetVisible returns the visibility state of the checkbox
func (v *Checkbox) GetVisible() bool {
	return v.visible
}

// SetPosition sets the position of the checkbox
func (v *Checkbox) SetPosition(x int, y int) {
	v.x = x
	v.y = y
}

// SetVisible sets the visibility of the checkbox
func (v *Checkbox) SetVisible(visible bool) {
	v.visible = visible
}
