package d2asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dc6"
	d2iface "github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// DC6Animation is an animation made from a DC6 file
type DC6Animation struct {
	Animation
}

// CreateDC6Animation creates an Animation from d2dc6.DC6 and d2dat.DATPalette
func CreateDC6Animation(renderer d2iface.Renderer, dc6 *d2dc6.DC6,
	palette d2iface.Palette) (d2iface.Animation, error) {
	animation := Animation{
		directions:     make([]animationDirection, dc6.Directions),
		playLength:     defaultPlayLength,
		playLoop:       true,
		originAtBottom: true,
	}

	DC6 := DC6Animation{Animation: animation}

	for frameIndex, dc6Frame := range dc6.Frames {
		sfc, err := renderer.NewSurface(int(dc6Frame.Width), int(dc6Frame.Height),
			d2iface.FilterNearest)
		if err != nil {
			return nil, err
		}

		indexData := make([]int, dc6Frame.Width*dc6Frame.Height)
		for i := range indexData {
			indexData[i] = -1
		}

		x := 0
		y := int(dc6Frame.Height) - 1
		offset := 0

		for {
			b := int(dc6Frame.FrameData[offset])
			offset++

			if b == 0x80 {
				if y == 0 {
					break
				}
				y--
				x = 0
			} else if b&0x80 > 0 {
				transparentPixels := b & 0x7f
				for i := 0; i < transparentPixels; i++ {
					indexData[x+y*int(dc6Frame.Width)+i] = -1
				}
				x += transparentPixels
			} else {
				for i := 0; i < b; i++ {
					indexData[x+y*int(dc6Frame.Width)+i] = int(dc6Frame.FrameData[offset])
					offset++
				}
				x += b
			}
		}

		bytesPerPixel := 4
		colorData := make([]byte, int(dc6Frame.Width)*int(dc6Frame.Height)*bytesPerPixel)

		for i := 0; i < int(dc6Frame.Width*dc6Frame.Height); i++ {
			// TODO: Is this == -1 or < 1?
			if indexData[i] < 1 {
				continue
			}

			c := palette.GetColors()[indexData[i]]
			colorData[i*bytesPerPixel] = c.R()
			colorData[i*bytesPerPixel+1] = c.G()
			colorData[i*bytesPerPixel+2] = c.B()
			colorData[i*bytesPerPixel+3] = c.A()
		}

		if err := sfc.ReplacePixels(colorData); err != nil {
			return nil, err
		}

		directionIndex := frameIndex / int(dc6.FramesPerDirection)

		// Is this required?
		//if directionIndex >= len(animation.directions) {
		//	animation.directions = append(animation.directions, new(animationDirection))
		//}

		direction := &animation.directions[directionIndex]
		direction.frames = append(direction.frames, &animationFrame{
			width:   int(dc6Frame.Width),
			height:  int(dc6Frame.Height),
			offsetX: int(dc6Frame.OffsetX),
			offsetY: int(dc6Frame.OffsetY),
			image:   sfc,
		})
	}

	return &DC6, nil
}

// Clone creates a copy of the animation
func (a *Animation) Clone() d2iface.Animation {
	animation := *a
	return &animation
}
