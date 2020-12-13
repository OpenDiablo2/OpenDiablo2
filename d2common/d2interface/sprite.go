package d2interface

import (
	"image"
	"image/color"
	"time"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// Sprite is an Sprite
type Sprite interface {
	BindRenderer(Renderer)
	Clone() Sprite
	SetSubLoop(startFrame, EndFrame int)
	Advance(time.Duration) error
	GetCurrentFrameSurface() Surface
	Render(target Surface)
	RenderFromOrigin(target Surface, shadow bool)
	RenderSection(sfc Surface, bound image.Rectangle)
	GetFrameSize(frameIndex int) (int, int, error)
	GetCurrentFrameSize() (int, int)
	GetCurrentFrameOffset() (int, int)
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
	SetPlayFPS(fps float64)
	SetPlaySpeed(playSpeed time.Duration)
	SetPlayLength(playLength time.Duration)
	SetColorMod(colorMod color.Color)
	GetPlayedCount() int
	ResetPlayedCount()
	SetEffect(effect d2enum.DrawEffect)
	SetShadow(shadow bool)
}
