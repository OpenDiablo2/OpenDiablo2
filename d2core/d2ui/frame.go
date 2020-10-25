package d2ui

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type frameOrientation = int

// Frame orientations
const (
	FrameLeft frameOrientation = iota
	FrameRight
)

// UIFrame is a representation of a ui panel that occupies the left or right half of the screen
// when it is visible.
type UIFrame struct {
	asset            *d2asset.AssetManager
	uiManager        *UIManager
	frame            *Sprite
	originX          int
	originY          int
	frameOrientation frameOrientation
}

// frame indices into dc6 images for panels
const (
	leftFrameTopLeft = iota
	leftFrameTopRight
	leftFrameMiddleRight
	leftFrameBottomLeft
	leftFrameBottomRight
	rightFrameTopLeft
	rightFrameTopRight
	rightFrameMiddleRight
	rightFrameBottomRight
	rightFrameBottomLeft
)

// NewUIFrame creates a new Frame instance
func NewUIFrame(
	asset *d2asset.AssetManager,
	uiManager *UIManager,
	frameOrientation frameOrientation,
) *UIFrame {
	var originX, originY = 0, 0

	switch frameOrientation {
	case FrameLeft:
		originX = 0
		originY = 0
	case FrameRight:
		originX = 400
		originY = 0
	}

	frame := &UIFrame{
		asset:            asset,
		uiManager:        uiManager,
		frameOrientation: frameOrientation,
		originX:          originX,
		originY:          originY,
	}
	frame.Load()

	return frame
}

// Load the necessary frame resources
func (u *UIFrame) Load() {
	sprite, err := u.uiManager.NewSprite(d2resource.Frame, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}

	u.frame = sprite
}

// Render the frame to the target surface
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
	// the frame pieces we are interested in.
	framePieces := []int{
		leftFrameTopLeft,
		leftFrameTopRight,
		leftFrameMiddleRight,
		leftFrameBottomLeft,
		leftFrameBottomRight,
	}

	// the frame coordinates
	coord := make(map[int]*struct{ x, y int })

	startX, startY := u.originX, u.originY
	currentX, currentY := startX, startY

	// first determine the coordinates for each frame
	// the order that we check is important
	for _, piece := range framePieces {
		width, height, err := u.frame.GetFrameSize(piece)
		if err != nil {
			return err
		}

		c := &struct{ x, y int }{}

		switch piece {
		case leftFrameTopLeft:
			c.x, c.y = currentX, currentY+height
			currentX, currentY = currentX+width, currentY+height
		case leftFrameTopRight:
			c.x, c.y = currentX, startY+height
			currentX = startX
		case leftFrameMiddleRight:
			c.x, c.y = currentX, currentY+height
			currentY += height
		case leftFrameBottomLeft:
			c.x, c.y = currentX, currentY+height
			currentX += width
		case leftFrameBottomRight:
			c.x, c.y = currentX, currentY+height
		}

		coord[piece] = c
	}

	// now render the pieces with the coordinates
	for idx, c := range coord {
		err := u.renderFramePiece(target, c.x, c.y, idx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *UIFrame) renderRight(target d2interface.Surface) error {
	// the frame pieces we are interested in.
	framePieces := []int{
		rightFrameTopLeft,
		rightFrameTopRight,
		rightFrameMiddleRight,
		rightFrameBottomRight,
		rightFrameBottomLeft,
	}

	// the frame coordinates
	coord := make(map[int]*struct{ x, y int })

	startX, startY := u.originX, u.originY
	currentX, currentY := startX, startY

	// first determine the coordinates for each frame
	// the order that we check is important
	for _, piece := range framePieces {
		width, height, err := u.frame.GetFrameSize(piece)
		if err != nil {
			return err
		}

		c := &struct{ x, y int }{}

		switch piece {
		case rightFrameTopLeft:
			c.x, c.y = currentX, currentY+height
			currentX += width
		case rightFrameTopRight:
			c.x, c.y = currentX, currentY+height
			currentX += width
			currentY += height
		case rightFrameMiddleRight:
			c.x, c.y = currentX-width, currentY+height
			currentY += height
		case rightFrameBottomRight:
			c.x, c.y = currentX-width, currentY+height
			currentX -= width
		case rightFrameBottomLeft:
			c.x, c.y = currentX-width, currentY+height
			currentX += width
		}

		coord[piece] = c
	}

	// now render the pieces with the coordinates
	for idx, c := range coord {
		err := u.renderFramePiece(target, c.x, c.y, idx)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetFrameBounds returns the maximum width and height of all frames in sprite.
func (u *UIFrame) GetFrameBounds() (width, height int) {
	return u.frame.GetFrameBounds()
}

// GetFrameCount returns the number of frames in the sprite
func (u *UIFrame) GetFrameCount() int {
	return u.frame.GetFrameCount()
}

func (u *UIFrame) renderFramePiece(sfc d2interface.Surface, x, y, idx int) error {
	if err := u.frame.SetCurrentFrame(idx); err != nil {
		return err
	}

	u.frame.SetPosition(x, y)

	if err := u.frame.Render(sfc); err != nil {
		return err
	}

	return nil
}
