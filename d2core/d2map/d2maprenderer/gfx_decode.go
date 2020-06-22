package d2maprenderer

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2dt1"

func (mr *MapRenderer) decodeTileGfxData(blocks []d2dt1.Block, pixels *[]byte, tileYOffset int32, tileWidth int32) {
	for _, block := range blocks {
		if block.Format == d2dt1.BlockFormatIsometric {
			// 3D isometric decoding
			xjump := []int32{14, 12, 10, 8, 6, 4, 2, 0, 2, 4, 6, 8, 10, 12, 14}
			nbpix := []int32{4, 8, 12, 16, 20, 24, 28, 32, 28, 24, 20, 16, 12, 8, 4}
			blockX := int32(block.X)
			blockY := int32(block.Y)
			length := int32(256)
			x := int32(0)
			y := int32(0)
			idx := 0
			for length > 0 {
				x = xjump[y]
				n := nbpix[y]
				length -= n
				for n > 0 {
					colorIndex := block.EncodedData[idx]
					if colorIndex != 0 {
						pixelColor := mr.palette.Colors[colorIndex]
						offset := 4 * (((blockY + y + tileYOffset) * tileWidth) + (blockX + x))
						(*pixels)[offset] = pixelColor.R
						(*pixels)[offset+1] = pixelColor.G
						(*pixels)[offset+2] = pixelColor.B
						(*pixels)[offset+3] = 255
					}
					x++
					n--
					idx++
				}
				y++
			}
			continue
		}
		// RLE Encoding
		blockX := int32(block.X)
		blockY := int32(block.Y)
		x := int32(0)
		y := int32(0)
		idx := 0
		length := block.Length
		for length > 0 {
			b1 := block.EncodedData[idx]
			b2 := block.EncodedData[idx+1]
			idx += 2
			length -= 2
			if (b1 | b2) == 0 {
				x = 0
				y++
				continue
			}
			x += int32(b1)
			length -= int32(b2)
			for b2 > 0 {
				colorIndex := block.EncodedData[idx]
				if colorIndex != 0 {
					pixelColor := mr.palette.Colors[colorIndex]

					offset := 4 * (((blockY + y + tileYOffset) * tileWidth) + (blockX + x))
					(*pixels)[offset] = pixelColor.R
					(*pixels)[offset+1] = pixelColor.G
					(*pixels)[offset+2] = pixelColor.B
					(*pixels)[offset+3] = 255

				}
				idx++
				x++
				b2--
			}
		}
	}
}
