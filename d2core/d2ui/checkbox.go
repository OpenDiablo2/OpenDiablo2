package d2ui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type Checkbox struct {
	x, y          int
	checkState    bool
	visible       bool
	width, height int
	Image         d2interface.Surface
	checkedImage  d2interface.Surface
	onClick       func()
	enabled       bool
}

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

	result.Image, _ = renderer.NewSurface(result.width, result.height, d2interface.FilterNearest)

	_ = checkboxSprite.RenderSegmented(result.Image, 1, 1, 0)

	result.checkedImage, _ = renderer.NewSurface(result.width, result.height, d2interface.FilterNearest)

	_ = checkboxSprite.RenderSegmented(result.checkedImage, 1, 1, 1)
	return result
}

func (v *Checkbox) Render(target d2interface.Surface) {
	target.PushCompositeMode(d2enum.CompositeModeSourceAtop)
	target.PushTranslation(v.x, v.y)
	target.PushFilter(d2interface.FilterNearest)
	defer target.PopN(3)

	if v.checkState {
		_ = target.Render(v.checkedImage)
	} else {
		_ = target.Render(v.Image)
	}
}

func (v *Checkbox) Advance(elapsed float64) {

}

func (v *Checkbox) GetEnabled() bool {
	return v.enabled
}

func (v *Checkbox) SetEnabled(enabled bool) {
	v.enabled = enabled
}

func (v *Checkbox) SetPressed(_ bool) {
}

func (v *Checkbox) SetCheckState(checkState bool) {
	v.checkState = checkState
}

func (v *Checkbox) GetCheckState() bool {
	return v.checkState
}

func (v *Checkbox) GetPressed() bool {
	return v.checkState
}

func (v *Checkbox) OnActivated(callback func()) {
	v.onClick = callback
}

func (v *Checkbox) Activate() {
	v.checkState = !v.checkState
	if v.onClick == nil {
		return
	}
	v.onClick()
}

func (v *Checkbox) GetPosition() (int, int) {
	return v.x, v.y
}

func (v *Checkbox) GetSize() (int, int) {
	return v.width, v.height
}

func (v *Checkbox) GetVisible() bool {
	return v.visible
}

func (v *Checkbox) SetPosition(x int, y int) {
	v.x = x
	v.y = y
}

func (v *Checkbox) SetVisible(visible bool) {
	v.visible = visible
}
