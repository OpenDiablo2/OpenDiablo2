package d2gui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

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
			animation.Render(target)
			target.Pop()

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
