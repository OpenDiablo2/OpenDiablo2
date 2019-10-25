package Common

import (
	"encoding/binary"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Sprite represents a type of object in D2 that is comprised of one or more frames and directions
type Sprite struct {
	Directions         uint32
	FramesPerDirection uint32
	Frames             []SpriteFrame
	X, Y               int
	Frame, Direction   uint8
	Blend              bool
	LastFrameTime      time.Time
	Animate            bool
	ColorMod           color.Color
	visible            bool
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
	Image     *ebiten.Image
}

// CreateSprite creates an instance of a sprite
func CreateSprite(data []byte, palette Palette) *Sprite {
	result := &Sprite{
		X:                  50,
		Y:                  50,
		Frame:              0,
		Direction:          0,
		Blend:              false,
		ColorMod:           nil,
		Directions:         binary.LittleEndian.Uint32(data[16:20]),
		FramesPerDirection: binary.LittleEndian.Uint32(data[20:24]),
		Animate:            false,
		LastFrameTime:      time.Now(),
	}
	dataPointer := uint32(24)
	totalFrames := result.Directions * result.FramesPerDirection
	framePointers := make([]uint32, totalFrames)
	for i := uint32(0); i < totalFrames; i++ {
		framePointers[i] = binary.LittleEndian.Uint32(data[dataPointer : dataPointer+4])
		dataPointer += 4
	}
	result.Frames = make([]SpriteFrame, totalFrames)
	for i := uint32(0); i < totalFrames; i++ {
		dataPointer = framePointers[i]

		result.Frames[i].Flip = binary.LittleEndian.Uint32(data[dataPointer : dataPointer+4])
		dataPointer += 4
		result.Frames[i].Width = binary.LittleEndian.Uint32(data[dataPointer : dataPointer+4])
		dataPointer += 4
		result.Frames[i].Height = binary.LittleEndian.Uint32(data[dataPointer : dataPointer+4])
		dataPointer += 4
		result.Frames[i].OffsetX = BytesToInt32(data[dataPointer : dataPointer+4])
		dataPointer += 4
		result.Frames[i].OffsetY = BytesToInt32(data[dataPointer : dataPointer+4])
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
		y := uint32(result.Frames[i].Height - 1)
		go func(ix, dataPointer uint32) {
			for true {
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
						result.Frames[ix].ImageData[x+(y*result.Frames[ix].Width)+uint32(ti)] = -1
					}
					x += uint32(transparentPixels)
				} else {
					for bi := 0; bi < int(b); bi++ {
						result.Frames[ix].ImageData[x+(y*result.Frames[ix].Width)+uint32(bi)] = int16(data[dataPointer])
						dataPointer++
					}
					x += uint32(b)
				}
			}
			result.Frames[ix].Image, _ = ebiten.NewImage(int(result.Frames[ix].Width), int(result.Frames[ix].Height), ebiten.FilterNearest)
			newData := make([]byte, result.Frames[ix].Width*result.Frames[ix].Height*4)
			for ii := uint32(0); ii < result.Frames[ix].Width*result.Frames[ix].Height; ii++ {
				if result.Frames[ix].ImageData[ii] == -1 {
					continue
				}
				newData[ii*4] = palette.Colors[result.Frames[ix].ImageData[ii]].R
				newData[(ii*4)+1] = palette.Colors[result.Frames[ix].ImageData[ii]].G
				newData[(ii*4)+2] = palette.Colors[result.Frames[ix].ImageData[ii]].B
				newData[(ii*4)+3] = 0xFF
			}
			result.Frames[ix].Image.ReplacePixels(newData)
		}(i, dataPointer)
	}
	return result
}

// GetSize returns the size of the sprite
func (v *Sprite) GetSize() (uint32, uint32) {
	frame := v.Frames[uint32(v.Frame)+(uint32(v.Direction)*v.FramesPerDirection)]
	return frame.Width, frame.Height
}

func (v *Sprite) updateAnimation() {
	if !v.Animate {
		return
	}
	tNow := time.Now()
	if v.LastFrameTime.Add(time.Millisecond * 25).After(tNow) {
		return
	}
	v.LastFrameTime = tNow
	v.Frame++
	if v.Frame >= uint8(v.FramesPerDirection) {
		v.Frame = 0
	}
}

// Draw draws the sprite onto the target
func (v *Sprite) Draw(target *ebiten.Image) {
	v.updateAnimation()
	opts := &ebiten.DrawImageOptions{}
	frame := v.Frames[uint32(v.Frame)+(uint32(v.Direction)*v.FramesPerDirection)]
	opts.GeoM.Translate(
		float64(int32(v.X)+frame.OffsetX),
		float64((int32(v.Y) - int32(frame.Height) + frame.OffsetY)),
	)
	if v.Blend {
		opts.CompositeMode = ebiten.CompositeModeLighter
	}
	if v.ColorMod != nil {
		opts.ColorM = ColorToColorM(v.ColorMod)
	}
	target.DrawImage(frame.Image, opts)
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
			}
			if v.ColorMod != nil {
				opts.ColorM = ColorToColorM(v.ColorMod)
			}
			target.DrawImage(frame.Image, opts)
			xOffset += int32(frame.Width)
			biggestYOffset = MaxInt32(biggestYOffset, int32(frame.Height))
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
func (v *Sprite) GetLocation() (int, int) {
	return v.X, v.Y
}
