package d2gui

import (
	"errors"
	"image/color"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

func loadFont(fontStyle FontStyle) (*d2asset.Font, error) {
	config := getFontStyleConfig(fontStyle)
	if config == nil {
		return nil, errors.New("invalid font style")
	}

	return singleton.asset.LoadFont(config.fontBasePath+".tbl", config.fontBasePath+".dc6",
		config.palettePath)
}

func renderSegmented(animation d2interface.Animation, segmentsX, segmentsY, frameOffset int,
	target d2interface.Surface) error {
	var currentY int

	for y := 0; y < segmentsY; y++ {
		var currentX, maxHeight int

		for x := 0; x < segmentsX; x++ {
			if err := animation.SetCurrentFrame(x + y*segmentsX + frameOffset*segmentsX*segmentsY); err != nil {
				return err
			}

			target.PushTranslation(x+currentX, y+currentY)
			err := animation.Render(target)
			target.Pop()

			if err != nil {
				return err
			}

			width, height := animation.GetCurrentFrameSize()
			maxHeight = d2math.MaxInt(maxHeight, height)
			currentX += width
		}

		currentY += maxHeight
	}

	return nil
}

func half(n int) int {
	return n / 2
}

func rgbaColor(rgba uint32) color.RGBA {
	result := color.RGBA{}
	a, b, g, r := 0, 1, 2, 3
	byteWidth := 8
	byteMask := 0xff

	for idx := 0; idx < 4; idx++ {
		shift := idx * byteWidth
		component := uint8(rgba>>shift) & uint8(byteMask)

		switch idx {
		case a:
			result.A = component
		case b:
			result.B = component
		case g:
			result.G = component
		case r:
			result.R = component
		}
	}

	return result
}
