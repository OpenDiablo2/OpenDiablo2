package d2ui

import (
	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/hajimehoshi/ebiten"
)

type Checkbox struct {
	x, y          int
	checkState    bool
	visible       bool
	width, height int
	Image         *ebiten.Image
	checkedImage  *ebiten.Image
	onClick       func()
	enabled       bool
}

func CreateCheckbox(checkState bool) Checkbox {
	result := Checkbox{
		checkState: checkState,
		visible:    true,
		width:      0,
		height:     0,
		enabled:    true,
	}
	checkboxSprite, _ := d2render.LoadSprite(d2resource.Checkbox, d2resource.PaletteFechar)
	result.width, result.height, _ = checkboxSprite.GetFrameSize(0)
	checkboxSprite.SetPosition(0, 0)

	result.Image, _ = ebiten.NewImage(int(result.width), int(result.height), ebiten.FilterNearest)
	checkboxSprite.RenderSegmented(result.Image, 1, 1, 0)

	result.checkedImage, _ = ebiten.NewImage(int(result.width), int(result.height), ebiten.FilterNearest)
	checkboxSprite.RenderSegmented(result.checkedImage, 1, 1, 1)
	return result
}

func (v Checkbox) Render(target *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{
		CompositeMode: ebiten.CompositeModeSourceAtop,
		Filter:        ebiten.FilterNearest,
	}
	opts.GeoM.Translate(float64(v.x), float64(v.y))
	if v.checkState == false {
		target.DrawImage(v.Image, opts)
	} else {
		target.DrawImage(v.checkedImage, opts)
	}
}
func (v Checkbox) GetEnabled() bool {
	return v.enabled
}

func (v *Checkbox) SetEnabled(enabled bool) {
	v.enabled = enabled
}

func (v Checkbox) SetPressed(pressed bool) {
}

func (v *Checkbox) SetCheckState(checkState bool) {
	v.checkState = checkState
}

func (v Checkbox) GetCheckState() bool {
	return v.checkState
}

func (v Checkbox) GetPressed() bool {
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

func (v Checkbox) GetPosition() (int, int) {
	return v.x, v.y
}

func (v Checkbox) GetSize() (int, int) {
	return v.width, v.height
}

func (v Checkbox) GetVisible() bool {
	return v.visible
}

func (v *Checkbox) SetPosition(x int, y int) {
	v.x = x
	v.y = y
}

func (v *Checkbox) SetVisible(visible bool) {
	v.visible = visible
}
