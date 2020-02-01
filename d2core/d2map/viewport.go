package d2map

import (
	"math"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type worldTrans struct {
	x float64
	y float64
}

type Viewport struct {
	screenRect   d2common.Rectangle
	transStack   []worldTrans
	transCurrent worldTrans
	camera       *Camera
}

func NewViewport(x, y, width, height int) *Viewport {
	return &Viewport{
		screenRect: d2common.Rectangle{
			Left:   x,
			Top:    y,
			Width:  width,
			Height: height,
		},
	}
}

func (v *Viewport) SetCamera(camera *Camera) {
	v.camera = camera
}

func (v *Viewport) WorldToScreen(x, y float64) (int, int) {
	return v.OrthoToScreen(v.WorldToOrtho(x, y))
}

func (v *Viewport) ScreenToWorld(x, y int) (float64, float64) {
	return v.OrthoToWorld(v.ScreenToOrtho(x, y))
}

func (v *Viewport) OrthoToWorld(x, y float64) (float64, float64) {
	worldX := (x/80 + y/40) / 2
	worldY := (y/40 - x/80) / 2
	return worldX, worldY
}

func (v *Viewport) WorldToOrtho(x, y float64) (float64, float64) {
	orthoX := (x - y) * 80
	orthoY := (x + y) * 40
	return orthoX, orthoY
}

func (v *Viewport) ScreenToOrtho(x, y int) (float64, float64) {
	camX, camY := v.getCameraOffset()
	screenX := float64(x) + camX - float64(v.screenRect.Left)
	screenY := float64(y) + camY - float64(v.screenRect.Top)
	return screenX, screenY
}

func (v *Viewport) OrthoToScreen(x, y float64) (int, int) {
	camOrthoX, camOrthoY := v.getCameraOffset()
	orthoX := int(math.Floor(x - camOrthoX + float64(v.screenRect.Left)))
	orthoY := int(math.Floor(y - camOrthoY + float64(v.screenRect.Top)))
	return orthoX, orthoY
}

func (v *Viewport) IsTileVisible(x, y float64) bool {
	orthoX1, orthoY1 := v.WorldToOrtho(x-3, y)
	orthoX2, orthoY2 := v.WorldToOrtho(x+3, y)
	return v.IsOrthoRectVisible(orthoX1, orthoY1, orthoX2, orthoY2)
}

func (v *Viewport) IsTileRectVisible(rect d2common.Rectangle) bool {
	left := float64((rect.Left - rect.Bottom()) * 80)
	top := float64((rect.Left + rect.Top) * 40)
	right := float64((rect.Right() - rect.Top) * 80)
	bottom := float64((rect.Right() + rect.Bottom()) * 40)
	return v.IsOrthoRectVisible(left, top, right, bottom)
}

func (v *Viewport) IsOrthoRectVisible(x1, y1, x2, y2 float64) bool {
	screenX1, screenY1 := v.OrthoToScreen(x1, y1)
	screenX2, screenY2 := v.OrthoToScreen(x2, y2)
	return !(screenX1 >= v.screenRect.Width || screenX2 < 0 || screenY1 >= v.screenRect.Height || screenY2 < 0)
}

func (v *Viewport) GetTranslationOrtho() (float64, float64) {
	return v.transCurrent.x, v.transCurrent.y
}

func (v *Viewport) GetTranslationScreen() (int, int) {
	return v.OrthoToScreen(v.transCurrent.x, v.transCurrent.y)
}

func (v *Viewport) PushTranslationOrtho(x, y float64) {
	v.transStack = append(v.transStack, v.transCurrent)
	v.transCurrent.x += x
	v.transCurrent.y += y
}

func (v *Viewport) PushTranslationWorld(x, y float64) {
	v.PushTranslationOrtho(v.WorldToOrtho(x, y))
}

func (v *Viewport) PushTranslationScreen(x, y int) {
	v.PushTranslationOrtho(v.ScreenToOrtho(x, y))
}

func (v *Viewport) PopTranslation() {
	count := len(v.transStack)
	if count == 0 {
		panic("empty stack")
	}

	v.transCurrent = v.transStack[count-1]
	v.transStack = v.transStack[:count-1]
}

func (v *Viewport) getCameraOffset() (float64, float64) {
	var camX, camY float64
	if v.camera != nil {
		camX, camY = v.camera.GetPosition()
	}

	camX -= float64(v.screenRect.Width / 2)
	camY -= float64(v.screenRect.Height / 2)

	return camX, camY
}
