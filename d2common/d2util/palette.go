package d2util

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

// ImgIndexToRGBA converts the given indices byte slice and palette into
// a byte slice of RGBA values
func ImgIndexToRGBA(indexData []byte, palette d2interface.Palette) []byte {
	bytesPerPixel := 4
	colorData := make([]byte, len(indexData)*bytesPerPixel)

	for i := 0; i < len(indexData); i++ {
		// Index zero is hardcoded transparent regardless of palette
		if indexData[i] == 0 {
			continue
		}

		c, err := palette.GetColor(int(indexData[i]))
		if err != nil {
			log.Print(err)
		}

		colorData[i*bytesPerPixel] = c.R()
		colorData[i*bytesPerPixel+1] = c.G()
		colorData[i*bytesPerPixel+2] = c.B()
		colorData[i*bytesPerPixel+3] = c.A()
	}

	return colorData
}
