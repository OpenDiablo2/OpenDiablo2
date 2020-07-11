package d2asset

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

func ImgIndexToRGBA(indexData []byte, palette d2interface.Palette) []byte {
	bytesPerPixel := 4
	colorData := make([]byte, len(indexData)*bytesPerPixel)

	for i := 0; i < len(indexData); i++ {
		// Index zero is hardcoded transparent regardless of palette
		if indexData[i] == 0 {
			continue
		}

		c, _ := palette.GetColor(int(indexData[i]))
		colorData[i*bytesPerPixel] = c.R()
		colorData[i*bytesPerPixel+1] = c.G()
		colorData[i*bytesPerPixel+2] = c.B()
		colorData[i*bytesPerPixel+3] = c.A()
	}

	return colorData
}
