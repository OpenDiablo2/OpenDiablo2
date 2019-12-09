package d2mapengine

import (
	"math"

	"github.com/OpenDiablo2/D2Shared/d2common"
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

func (v *Viewport) IsoToScreen(x, y float64) (int, int) {
	return v.WorldToScreen(v.IsoToWorld(x, y))
}

func (v *Viewport) ScreenToIso(x, y int) (float64, float64) {
	return v.WorldToIso(v.ScreenToWorld(x, y))
}

func (v *Viewport) WorldToIso(x, y float64) (float64, float64) {
	isoX := (x/80 + y/40) / 2
	isoY := (y/40 - x/80) / 2
	return isoX, isoY
}

func (v *Viewport) IsoToWorld(x, y float64) (float64, float64) {
	worldX := (x - y) * 80
	worldY := (x + y) * 40
	return worldX, worldY
}

func (v *Viewport) ScreenToWorld(x, y int) (float64, float64) {
	camX, camY := v.getCameraOffset()
	screenX := float64(x) + camX - float64(v.screenRect.Left)
	screenY := float64(y) + camY - float64(v.screenRect.Top)
	return screenX, screenY
}

func (v *Viewport) WorldToScreen(x, y float64) (int, int) {
	camX, camY := v.getCameraOffset()
	worldX := int(math.Floor(x - camX + float64(v.screenRect.Left)))
	worldY := int(math.Floor(y - camY + float64(v.screenRect.Top)))
	return worldX, worldY
}

func (v *Viewport) IsWorldTileVisbile(x, y float64) bool {
	worldX1, worldY1 := v.IsoToWorld(x-2, y)
	worldX2, worldY2 := v.IsoToWorld(x+2, y)
	return v.IsWorldRectVisible(worldX1, worldY1, worldX2, worldY2)
}

func (v *Viewport) IsWorldPointVisible(x, y float64) bool {
	screenX, screenY := v.WorldToScreen(x, y)
	return screenX >= 0 && screenX < v.screenRect.Width && screenY >= 0 && screenY < v.screenRect.Height
}

func (v *Viewport) IsWorldRectVisible(x1, y1, x2, y2 float64) bool {
	screenX1, screenY1 := v.WorldToScreen(x1, y1)
	screenX2, screenY2 := v.WorldToScreen(x2, y2)
	return !(screenX1 >= v.screenRect.Width || screenX2 < 0 || screenY1 >= v.screenRect.Height || screenY2 < 0)
}

func (v *Viewport) GetTranslation() (float64, float64) {
	return v.transCurrent.x, v.transCurrent.y
}

func (v *Viewport) PushTranslation(x, y float64) {
	v.transStack = append(v.transStack, v.transCurrent)
	v.transCurrent.x += x
	v.transCurrent.y += y
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
