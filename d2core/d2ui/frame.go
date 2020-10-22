package d2ui

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type UIFrame struct {
	asset *d2asset.AssetManager
	uiManager *UIManager
	frame *Sprite
	originX int
	originY int
	frameOrientation FrameOrientation
}

type FrameOrientation = int
const(
	FrameLeft FrameOrientation = iota
	FrameRight
)

func NewUIFrame (
	asset *d2asset.AssetManager,
	uiManager *UIManager,
	frameOrientation FrameOrientation,
) *UIFrame {
	var originX, originY = 0,0

	switch frameOrientation {
	case FrameLeft:
		originX = 0
		originY = 0
	case FrameRight:
		originX = 400
		originY = 0
	}
	frame := &UIFrame {
		asset : asset,
		uiManager: uiManager,
		frameOrientation: frameOrientation,
		originX: originX,
		originY: originY,
	}
	frame.Load()
	return frame
}

func (u *UIFrame) Load() {
	sprite, err := u.uiManager.NewSprite(d2resource.Frame, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}
	u.frame = sprite
}

func (u *UIFrame) Render(target d2interface.Surface) error {
	switch u.frameOrientation {
	case FrameLeft:
		return u.renderLeft(target)
	case FrameRight:
		return u.renderRight(target)
	}
	return nil
}

func (u *UIFrame) renderLeft(target d2interface.Surface) error {
	x, y := u.originX, u.originY

	// Frame
	// Top left
	if err := u.frame.SetCurrentFrame(0); err != nil {
		return err
	}

	w, h := u.frame.GetCurrentFrameSize()

	u.frame.SetPosition(x, y+h)

	if err := u.frame.Render(target); err != nil {
		return err
	}

	x += w
	y += h

	// Top right
	if err := u.frame.SetCurrentFrame(1); err != nil {
		return err
	}

	_, h = u.frame.GetCurrentFrameSize()

	u.frame.SetPosition(x, u.originY+h)

	if err := u.frame.Render(target); err != nil {
		return err
	}

	x = u.originX

	// Right
	if err := u.frame.SetCurrentFrame(2); err != nil {
		return err
	}

	_, h = u.frame.GetCurrentFrameSize()
	u.frame.SetPosition(x, y+h)

	if err := u.frame.Render(target); err != nil {
		return err
	}

	y += h

	// Bottom left
	if err := u.frame.SetCurrentFrame(3); err != nil {
		return err
	}

	w, h = u.frame.GetCurrentFrameSize()

	u.frame.SetPosition(x, y+h)

	if err := u.frame.Render(target); err != nil {
		return err
	}

	x += w

	// Bottom right
	if err := u.frame.SetCurrentFrame(4); err != nil {
		return err
	}

	_, h = u.frame.GetCurrentFrameSize()

	u.frame.SetPosition(x, y+h)

	if err := u.frame.Render(target); err != nil {
		return err
	}
	return nil
}

func (u *UIFrame) renderRight(target d2interface.Surface) error {
	x, y := u.originX, u.originY

	// Frame
	// Top left
	if err := u.frame.SetCurrentFrame(5); err != nil {
		return err
	}

	w, h := u.frame.GetCurrentFrameSize()

	u.frame.SetPosition(x, y+h)

	if err := u.frame.Render(target); err != nil {
		return err
	}

	x += w

	// Top right
	if err := u.frame.SetCurrentFrame(6); err != nil {
		return err
	}

	w, h = u.frame.GetCurrentFrameSize()

	u.frame.SetPosition(x, y+h)

	if err := u.frame.Render(target); err != nil {
		return err
	}

	x += w
	y += h

	// Right
	if err := u.frame.SetCurrentFrame(7); err != nil {
		return err
	}

	w, h = u.frame.GetCurrentFrameSize()

	u.frame.SetPosition(x-w, y+h)

	if err := u.frame.Render(target); err != nil {
		return err
	}

	y += h

	// Bottom right
	if err := u.frame.SetCurrentFrame(8); err != nil {
		return err
	}

	w, h = u.frame.GetCurrentFrameSize()

	u.frame.SetPosition(x-w, y+h)

	if err := u.frame.Render(target); err != nil {
		return err
	}

	x -= w

	// Bottom left
	if err := u.frame.SetCurrentFrame(9); err != nil {
		return err
	}

	w, h = u.frame.GetCurrentFrameSize()

	u.frame.SetPosition(x-w, y+h)

	if err := u.frame.Render(target); err != nil {
		return err
	}
	return nil
}

func (u *UIFrame) GetFrameBounds() (width, height int) {
	return u.frame.GetFrameBounds()
}

func (u *UIFrame) GetFrameCount() int {
	return u.frame.GetFrameCount()
}
