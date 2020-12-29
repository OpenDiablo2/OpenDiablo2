package d2dt1

const (
	blockDataLength = 256
)

// DecodeTileGfxData decodes tile graphics data for a slice of dt1 blocks
func DecodeTileGfxData(blocks []Block, pixels *[]byte, tileYOffset, tileWidth int32) {
	for _, block := range blocks {
		if block.Format == BlockFormatIsometric {
			// 3D isometric decoding
			xjump := []int32{14, 12, 10, 8, 6, 4, 2, 0, 2, 4, 6, 8, 10, 12, 14}
			nbpix := []int32{4, 8, 12, 16, 20, 24, 28, 32, 28, 24, 20, 16, 12, 8, 4}
			blockX := int32(block.X)
			blockY := int32(block.Y)
			length := int32(blockDataLength)
			x := int32(0)
			y := int32(0)
			idx := 0

			for length > 0 {
				x = xjump[y]
				n := nbpix[y]
				length -= n

				for n > 0 {
					offset := ((blockY + y + tileYOffset) * tileWidth) + (blockX + x)
					(*pixels)[offset] = block.EncodedData[idx]
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
				offset := ((blockY + y + tileYOffset) * tileWidth) + (blockX + x)
				(*pixels)[offset] = block.EncodedData[idx]
				idx++
				x++
				b2--
			}
		}
	}
}
