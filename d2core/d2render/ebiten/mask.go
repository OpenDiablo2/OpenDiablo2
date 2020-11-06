package ebiten

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Mask struct {
	drawDebug  bool
	background *ebiten.Image
	alpha      *ebiten.Image
}

func NewMask(maskWidth, maskHeight, alphaWidth, alphaHeight, alphaX, alphaY int) *Mask {
	ret := &Mask{}

	ret.background = ebiten.NewImage(maskWidth, maskHeight)
	ret.background.Fill(color.White)

	ret.alpha = ebiten.NewImageFromImage(image.NewAlpha(image.Rectangle{}))

	op := &ebiten.DrawImageOptions{}
	op.CompositeMode = ebiten.CompositeModeCopy
	op.GeoM.Translate(float64(alphaX), float64(alphaY))

	ret.background.DrawImage(ret.alpha, op)

	return ret
}

func (m *Mask) Draw(target, bg, fg *ebitenSurface) {
	op := &ebiten.DrawImageOptions{}
	op.CompositeMode = ebiten.CompositeModeSourceIn
	m.background.DrawImage(fg.image, op)
	target.image.DrawImage(bg.image, &ebiten.DrawImageOptions{})
	target.image.DrawImage(m.background, &ebiten.DrawImageOptions{})
}
