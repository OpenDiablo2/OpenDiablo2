package d2gui

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

func loadFont(fontStyle FontStyle) (*d2asset.Font, error) {
	config, ok := fontStyleConfigs[fontStyle]
	if !ok {
		return nil, errors.New("invalid font style")
	}

	return d2asset.LoadFont(config.fontBasePath+".tbl", config.fontBasePath+".dc6", config.palettePath)
}

func renderSegmented(animation *d2asset.Animation, segmentsX, segmentsY, frameOffset int, target d2interface.Surface) error {
	var currentY int
	for y := 0; y < segmentsY; y++ {
		var currentX int
		var maxHeight int
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
			maxHeight = d2common.MaxInt(maxHeight, height)
			currentX += width
		}

		currentY += maxHeight
	}

	return nil
}
