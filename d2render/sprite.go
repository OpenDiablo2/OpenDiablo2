package d2render

import (
	"image/color"
	"log"
	"sync"

	"github.com/OpenDiablo2/D2Shared/d2data/d2dc6"
	"github.com/OpenDiablo2/D2Shared/d2helper"
	"github.com/OpenDiablo2/OpenDiablo2/d2corehelper"

	"github.com/hajimehoshi/ebiten"
)

// Sprite represents a type of object in D2 that is comprised of one or more frames and directions
type Sprite struct {
	Directions         uint32
	FramesPerDirection uint32
	Frames             []SpriteFrame
	SpecialFrameTime   int
	AnimateBackwards   bool // Because why not
	StopOnLastFrame    bool
	X, Y               int
	Frame, Direction   int16
	Blend              bool
	LastFrameTime      float64
	Animate            bool
	ColorMod           color.Color
	valid              bool
}

// SpriteFrame represents a single frame of a sprite
type SpriteFrame struct {
	Flip      uint32
	Width     uint32
	Height    uint32
	OffsetX   int32
	OffsetY   int32
	Unknown   uint32
	NextBlock uint32
	Length    uint32
	FrameData []byte
	Image     *ebiten.Image
}

func CreateSpriteFromDC6(dc6 d2dc6.DC6File) Sprite {
	result := Sprite{
		X:                  50,
		Y:                  50,
		Frame:              0,
		Direction:          0,
		Blend:              false,
		ColorMod:           nil,
		Directions:         dc6.Directions,
		FramesPerDirection: dc6.FramesPerDirection,
		Animate:            false,
		LastFrameTime:      d2helper.Now(),
		SpecialFrameTime:   -1,
		StopOnLastFrame:    false,
		valid:              true,
		AnimateBackwards:   false,
	}

	result.Frames = make([]SpriteFrame, len(dc6.Frames))
	wg := sync.WaitGroup{}
	wg.Add(len(dc6.Frames))
	for i, f := range dc6.Frames {
		go func(i int, frame *d2dc6.DC6Frame) {
			defer wg.Done()

			image, err := ebiten.NewImage(int(frame.Width), int(frame.Height), ebiten.FilterNearest)
			if err != nil {
				log.Printf("failed to create new image: %v", err)
			}
			if err := image.ReplacePixels(frame.ColorData()); err != nil {
				log.Printf("failed to replace pixels: %v", err)
			}

			result.Frames[i] = SpriteFrame{
				Flip:      frame.Flipped,
				Width:     frame.Width,
				Height:    frame.Height,
				OffsetX:   frame.OffsetX,
				OffsetY:   frame.OffsetY,
				Unknown:   frame.Unknown,
				NextBlock: frame.NextBlock,
				Length:    frame.Length,
				Image:     image,
			}
		}(i, f)
	}
	wg.Wait()
	return result
}

func (v Sprite) IsValid() bool {
	return v.valid
}

// GetSize returns the size of the sprite
func (v Sprite) GetSize() (uint32, uint32) {
	frame := v.Frames[uint32(v.Frame)+(uint32(v.Direction)*v.FramesPerDirection)]
	return frame.Width, frame.Height
}

func (v *Sprite) updateAnimation() {
	if !v.Animate {
		return
	}
	var timePerFrame float64

	if v.SpecialFrameTime >= 0 {
		timePerFrame = (float64(v.SpecialFrameTime) / float64(len(v.Frames))) / 1000.0
	} else {
		timePerFrame = 1.0 / float64(len(v.Frames))
	}
	now := d2helper.Now()
	for v.LastFrameTime+timePerFrame < now {
		v.LastFrameTime += timePerFrame
		if !v.AnimateBackwards {
			v.Frame++
			if v.Frame >= int16(v.FramesPerDirection) {
				if v.StopOnLastFrame {
					v.Frame = int16(v.FramesPerDirection) - 1
				} else {
					v.Frame = 0
				}
			}
			continue
		}
		v.Frame--
		if v.Frame < 0 {
			if v.StopOnLastFrame {
				v.Frame = 0
			} else {
				v.Frame = int16(v.FramesPerDirection) - 1
			}
		}
	}
}

func (v *Sprite) ResetAnimation() {
	v.LastFrameTime = d2helper.Now()
	v.Frame = 0
}

func (v Sprite) OnLastFrame() bool {
	return v.Frame == int16(v.FramesPerDirection-1)
}

// GetFrameSize returns the size of the specific frame
func (v Sprite) GetFrameSize(frame int) (width, height uint32) {
	width = v.Frames[frame].Width
	height = v.Frames[frame].Height
	return
}

// GetTotalFrames returns the number of frames in this sprite (for all directions)
func (v Sprite) GetTotalFrames() int {
	return len(v.Frames)
}

// Draw draws the sprite onto the target
func (v *Sprite) Draw(target *ebiten.Image) {
	v.updateAnimation()
	opts := &ebiten.DrawImageOptions{}
	frame := v.Frames[uint32(v.Frame)+(uint32(v.Direction)*v.FramesPerDirection)]
	opts.GeoM.Translate(
		float64(int32(v.X)+frame.OffsetX),
		float64(int32(v.Y)-int32(frame.Height)+frame.OffsetY),
	)
	if v.Blend {
		opts.CompositeMode = ebiten.CompositeModeLighter
	} else {
		opts.CompositeMode = ebiten.CompositeModeSourceOver
	}
	if v.ColorMod != nil {
		opts.ColorM = d2corehelper.ColorToColorM(v.ColorMod)
	}
	if err := target.DrawImage(frame.Image, opts); err != nil {
		log.Panic(err.Error())
	}
}

// DrawSegments draws the sprite via a grid of segments
func (v *Sprite) DrawSegments(target *ebiten.Image, xSegments, ySegments, offset int) {
	v.updateAnimation()
	yOffset := int32(0)
	for y := 0; y < ySegments; y++ {
		xOffset := int32(0)
		biggestYOffset := int32(0)
		for x := 0; x < xSegments; x++ {
			frame := v.Frames[uint32(x+(y*xSegments)+(offset*xSegments*ySegments))]
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(
				float64(int32(v.X)+frame.OffsetX+xOffset),
				float64(int32(v.Y)+frame.OffsetY+yOffset),
			)
			if v.Blend {
				opts.CompositeMode = ebiten.CompositeModeLighter
			} else {
				opts.CompositeMode = ebiten.CompositeModeSourceOver
			}
			if v.ColorMod != nil {
				opts.ColorM = d2corehelper.ColorToColorM(v.ColorMod)
			}
			if err := target.DrawImage(frame.Image, opts); err != nil {
				log.Panic(err.Error())
			}
			xOffset += int32(frame.Width)
			biggestYOffset = d2helper.MaxInt32(biggestYOffset, int32(frame.Height))
		}
		yOffset += biggestYOffset
	}
}

// MoveTo moves the sprite to the specified coordinates
func (v *Sprite) MoveTo(x, y int) {
	v.X = x
	v.Y = y
}

// GetLocation returns the location of the sprite
func (v Sprite) GetLocation() (int, int) {
	return v.X, v.Y
}
