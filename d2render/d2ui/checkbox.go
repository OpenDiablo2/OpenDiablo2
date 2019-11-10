package d2ui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/hajimehoshi/ebiten"
)

type Checkbox struct {
	x, y          int
	pressed       bool
	visible       bool
	width, height uint32
	Image         *ebiten.Image
	checkedImage  *ebiten.Image
	onClick       func()
	enabled       bool
}

func CreateCheckbox(fileProvider d2interface.FileProvider, toggleState bool) Checkbox {
	result := Checkbox{
		pressed: toggleState,
		visible: true,
		width:   0,
		height:  0,
		enabled: true,
	}
	checkboxSprite := d2render.CreateSprite(fileProvider.LoadFile(d2resource.Checkbox), d2datadict.Palettes[d2enum.Fechar])
	result.width, result.height = checkboxSprite.GetFrameSize(0)
	checkboxSprite.MoveTo(0, 0)

	result.Image, _ = ebiten.NewImage(int(result.width), int(result.height), ebiten.FilterNearest)
	checkboxSprite.DrawSegments(result.Image, 1, 1, 0)

	result.checkedImage, _ = ebiten.NewImage(int(result.width), int(result.height), ebiten.FilterNearest)
	checkboxSprite.DrawSegments(result.checkedImage, 1, 1, 1)
	return result
}

func (v Checkbox) Draw(target *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{
		CompositeMode: ebiten.CompositeModeSourceAtop,
		Filter:        ebiten.FilterNearest,
	}
	opts.GeoM.Translate(float64(v.x), float64(v.y))
	if v.pressed == false {
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

func (v *Checkbox) SetPressed(pressed bool) {
	v.pressed = pressed
}

func (v Checkbox) GetPressed() bool {
	return v.pressed
}

func (v *Checkbox) OnActivated(callback func()) {
	v.onClick = callback
}

func (v *Checkbox) Activate() {
	v.pressed = !v.pressed
	if v.onClick == nil {
		return
	}
	v.onClick()
}

func (v Checkbox) GetLocation() (int, int) {
	return v.x, v.y
}

func (v Checkbox) GetSize() (uint32, uint32) {
	return v.width, v.height
}

func (v Checkbox) GetVisible() bool {
	return v.visible
}

func (v *Checkbox) MoveTo(x int, y int) {
	v.x = x
	v.y = y
}

func (v *Checkbox) SetVisible(visible bool) {
	v.visible = visible
}
