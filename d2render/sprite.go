package d2render

import (
	"encoding/binary"
	"image"
	"image/color"
	"log"
	"sync"

	"github.com/OpenDiablo2/OpenDiablo2/d2helper"

	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"

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
	ImageData []int16
	FrameData []byte
	Image     *ebiten.Image
}

// CreateSprite creates an instance of a sprite
func CreateSprite(data []byte, palette d2datadict.PaletteRec) Sprite {
	result := Sprite{
		X:                  50,
		Y:                  50,
		Frame:              0,
		Direction:          0,
		Blend:              false,
		ColorMod:           nil,
		Directions:         binary.LittleEndian.Uint32(data[16:20]),
		FramesPerDirection: binary.LittleEndian.Uint32(data[20:24]),
		Animate:            false,
		LastFrameTime:      d2helper.Now(),
		SpecialFrameTime:   -1,
		StopOnLastFrame:    false,
		valid:              false,
		AnimateBackwards:   false,
	}
	dataPointer := uint32(24)
	totalFrames := result.Directions * result.FramesPerDirection
	framePointers := make([]uint32, totalFrames)
	for i := uint32(0); i < totalFrames; i++ {
		framePointers[i] = binary.LittleEndian.Uint32(data[dataPointer : dataPointer+4])
		dataPointer += 4
	}
	result.Frames = make([]SpriteFrame, totalFrames)
	wg := sync.WaitGroup{}
	wg.Add(int(totalFrames))
	for i := uint32(0); i < totalFrames; i++ {
		go func(i uint32) {
			defer wg.Done()
			dataPointer := framePointers[i]
			result.Frames[i] = SpriteFrame{}
			result.Frames[i].Flip = binary.LittleEndian.Uint32(data[dataPointer : dataPointer+4])
			dataPointer += 4
			result.Frames[i].Width = binary.LittleEndian.Uint32(data[dataPointer : dataPointer+4])
			dataPointer += 4
			result.Frames[i].Height = binary.LittleEndian.Uint32(data[dataPointer : dataPointer+4])
			dataPointer += 4
			result.Frames[i].OffsetX = d2helper.BytesToInt32(data[dataPointer : dataPointer+4])
			dataPointer += 4
			result.Frames[i].OffsetY = d2helper.BytesToInt32(data[dataPointer : dataPointer+4])
			dataPointer += 4
			result.Frames[i].Unknown = binary.LittleEndian.Uint32(data[dataPointer : dataPointer+4])
			dataPointer += 4
			result.Frames[i].NextBlock = binary.LittleEndian.Uint32(data[dataPointer : dataPointer+4])
			dataPointer += 4
			result.Frames[i].Length = binary.LittleEndian.Uint32(data[dataPointer : dataPointer+4])
			dataPointer += 4
			result.Frames[i].ImageData = make([]int16, result.Frames[i].Width*result.Frames[i].Height)
			for fi := range result.Frames[i].ImageData {
				result.Frames[i].ImageData[fi] = -1
			}

			x := uint32(0)
			y := result.Frames[i].Height - 1
			for {
				b := data[dataPointer]
				dataPointer++
				if b == 0x80 {
					if y == 0 {
						break
					}
					y--
					x = 0
				} else if (b & 0x80) > 0 {
					transparentPixels := b & 0x7F
					for ti := byte(0); ti < transparentPixels; ti++ {
						result.Frames[i].ImageData[x+(y*result.Frames[i].Width)+uint32(ti)] = -1
					}
					x += uint32(transparentPixels)
				} else {
					for bi := 0; bi < int(b); bi++ {
						result.Frames[i].ImageData[x+(y*result.Frames[i].Width)+uint32(bi)] = int16(data[dataPointer])
						dataPointer++
					}
					x += uint32(b)
				}
			}
			var img = image.NewRGBA(image.Rect(0, 0, int(result.Frames[i].Width), int(result.Frames[i].Height)))
			for ii := uint32(0); ii < result.Frames[i].Width*result.Frames[i].Height; ii++ {
				if result.Frames[i].ImageData[ii] < 1 { // TODO: Is this == -1 or < 1?
					continue
				}
				img.Pix[ii*4] = palette.Colors[result.Frames[i].ImageData[ii]].R
				img.Pix[(ii*4)+1] = palette.Colors[result.Frames[i].ImageData[ii]].G
				img.Pix[(ii*4)+2] = palette.Colors[result.Frames[i].ImageData[ii]].B
				img.Pix[(ii*4)+3] = 0xFF
			}
			newImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterNearest)
			result.Frames[i].Image = newImage
			img = nil
		}(i)
	}
	wg.Wait()
	result.valid = true
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
		opts.ColorM = d2helper.ColorToColorM(v.ColorMod)
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
				opts.ColorM = d2helper.ColorToColorM(v.ColorMod)
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
