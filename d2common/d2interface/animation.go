package d2interface

import (
	"image"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// Animation is an animation
type Animation interface {
	BindRenderer(Renderer) error
	Clone() Animation
	SetSubLoop(startFrame, EndFrame int)
	Advance(elapsed float64) error
	Render(target Surface) error
	RenderFromOrigin(target Surface, shadow bool) error
	RenderSection(sfc Surface, bound image.Rectangle) error
	GetFrameSize(frameIndex int) (int, int, error)
	GetCurrentFrameSize() (int, int)
	GetFrameBounds() (int, int)
	GetCurrentFrame() int
	GetFrameCount() int
	IsOnFirstFrame() bool
	IsOnLastFrame() bool
	GetDirectionCount() int
	SetDirection(directionIndex int) error
	GetDirection() int
	SetCurrentFrame(frameIndex int) error
	Rewind()
	PlayForward()
	PlayBackward()
	Pause()
	SetPlayLoop(loop bool)
	SetPlaySpeed(playSpeed float64)
	SetPlayLength(playLength float64)
	SetPlayLengthMs(playLengthMs int)
	SetColorMod(colorMod color.Color)
	GetPlayedCount() int
	ResetPlayedCount()
	SetEffect(effect d2enum.DrawEffect)
	SetShadow(shadow bool)
}
