package d2dcc

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2math"
)

const cellsPerRow = 4

// DCCDirection represents a DCCDirection file.
type DCCDirection struct {
	OutSizeCoded               int
	CompressionFlags           int
	Variable0Bits              int
	WidthBits                  int
	HeightBits                 int
	XOffsetBits                int
	YOffsetBits                int
	OptionalDataBits           int
	CodedBytesBits             int
	EqualCellsBitstreamSize    int
	PixelMaskBitstreamSize     int
	EncodingTypeBitsreamSize   int
	RawPixelCodesBitstreamSize int
	Frames                     []*DCCDirectionFrame
	PaletteEntries             [256]byte
	Box                        d2geom.Rectangle
	Cells                      []*DCCCell
	PixelData                  []byte
	HorizontalCellCount        int
	VerticalCellCount          int
	PixelBuffer                []DCCPixelBufferEntry
}

// CreateDCCDirection creates an instance of a DCCDirection.
func CreateDCCDirection(bm *d2datautils.BitMuncher, file *DCC) *DCCDirection {
	var crazyBitTable = []byte{0, 1, 2, 4, 6, 8, 10, 12, 14, 16, 20, 24, 26, 28, 30, 32}

	result := &DCCDirection{
		OutSizeCoded:     int(bm.GetUInt32()),
		CompressionFlags: int(bm.GetBits(2)),                //nolint:gomnd // binary data
		Variable0Bits:    int(crazyBitTable[bm.GetBits(4)]), //nolint:gomnd // binary data
		WidthBits:        int(crazyBitTable[bm.GetBits(4)]), //nolint:gomnd // binary data
		HeightBits:       int(crazyBitTable[bm.GetBits(4)]), //nolint:gomnd // binary data
		XOffsetBits:      int(crazyBitTable[bm.GetBits(4)]), //nolint:gomnd // binary data
		YOffsetBits:      int(crazyBitTable[bm.GetBits(4)]), //nolint:gomnd // binary data
		OptionalDataBits: int(crazyBitTable[bm.GetBits(4)]), //nolint:gomnd // binary data
		CodedBytesBits:   int(crazyBitTable[bm.GetBits(4)]), //nolint:gomnd // binary data
		Frames:           make([]*DCCDirectionFrame, file.FramesPerDirection),
	}

	minx := 100000
	miny := 100000
	maxx := -100000
	maxy := -100000

	// Load the frame headers
	for frameIdx := 0; frameIdx < file.FramesPerDirection; frameIdx++ {
		result.Frames[frameIdx] = CreateDCCDirectionFrame(bm, result)
		minx = int(d2math.MinInt32(int32(result.Frames[frameIdx].Box.Left), int32(minx)))
		miny = int(d2math.MinInt32(int32(result.Frames[frameIdx].Box.Top), int32(miny)))
		maxx = int(d2math.MaxInt32(int32(result.Frames[frameIdx].Box.Right()), int32(maxx)))
		maxy = int(d2math.MaxInt32(int32(result.Frames[frameIdx].Box.Bottom()), int32(maxy)))
	}

	result.Box = d2geom.Rectangle{Left: minx, Top: miny, Width: maxx - minx, Height: maxy - miny}

	if result.OptionalDataBits > 0 {
		log.Panic("Optional bits in DCC data is not currently supported.")
	}

	if (result.CompressionFlags & 0x2) > 0 {
		result.EqualCellsBitstreamSize = int(bm.GetBits(20)) //nolint:gomnd // binary data
	}

	result.PixelMaskBitstreamSize = int(bm.GetBits(20)) //nolint:gomnd // binary data

	if (result.CompressionFlags & 0x1) > 0 {
		result.EncodingTypeBitsreamSize = int(bm.GetBits(20))   //nolint:gomnd // binary data
		result.RawPixelCodesBitstreamSize = int(bm.GetBits(20)) //nolint:gomnd // binary data
	}

	for paletteEntryCount, i := 0, 0; i < 256; i++ {
		valid := bm.GetBit() != 0
		if valid {
			result.PaletteEntries[paletteEntryCount] = byte(i)
			paletteEntryCount++
		}
	}

	// HERE BE GIANTS:
	// Because of the way this thing mashes bits together, BIT offset matters
	// here. For example, if you are on byte offset 3, bit offset 6, and
	// the EqualCellsBitstreamSize is 20 bytes, then the next bit stream
	// will be located at byte 23, bit offset 6!
	equalCellsBitstream := d2datautils.CopyBitMuncher(bm)

	bm.SkipBits(result.EqualCellsBitstreamSize)

	pixelMaskBitstream := d2datautils.CopyBitMuncher(bm)

	bm.SkipBits(result.PixelMaskBitstreamSize)

	encodingTypeBitsream := d2datautils.CopyBitMuncher(bm)

	bm.SkipBits(result.EncodingTypeBitsreamSize)

	rawPixelCodesBitstream := d2datautils.CopyBitMuncher(bm)

	bm.SkipBits(result.RawPixelCodesBitstreamSize)

	pixelCodeandDisplacement := d2datautils.CopyBitMuncher(bm)

	// Calculate the cells for the direction
	result.calculateCells()

	// Calculate the cells for each of the frames
	for _, frame := range result.Frames {
		frame.recalculateCells(result)
	}

	// Fill in the pixel buffer
	result.fillPixelBuffer(pixelCodeandDisplacement, equalCellsBitstream, pixelMaskBitstream, encodingTypeBitsream, rawPixelCodesBitstream)

	// Generate the actual frame pixel data
	result.generateFrames(pixelCodeandDisplacement)

	result.PixelBuffer = nil

	// Verify that everything we expected to read was actually read (sanity check)...
	result.verify(equalCellsBitstream, pixelMaskBitstream, encodingTypeBitsream, rawPixelCodesBitstream)

	bm.SkipBits(pixelCodeandDisplacement.BitsRead())

	return result
}

func (v *DCCDirection) verify(
	equalCellsBitstream,
	pixelMaskBitstream,
	encodingTypeBitstream,
	rawPixelCodesBitstream *d2datautils.BitMuncher,
) {
	if equalCellsBitstream.BitsRead() != v.EqualCellsBitstreamSize {
		log.Panic("Did not read the correct number of bits!")
	}

	if pixelMaskBitstream.BitsRead() != v.PixelMaskBitstreamSize {
		log.Panic("Did not read the correct number of bits!")
	}

	if encodingTypeBitstream.BitsRead() != v.EncodingTypeBitsreamSize {
		log.Panic("Did not read the correct number of bits!")
	}

	if rawPixelCodesBitstream.BitsRead() != v.RawPixelCodesBitstreamSize {
		log.Panic("Did not read the correct number of bits!")
	}
}

// nolint:gocognit,gocyclo // Can't reduce
func (v *DCCDirection) generateFrames(pcd *d2datautils.BitMuncher) {
	pbIdx := 0

	for _, cell := range v.Cells {
		cell.LastWidth = -1
		cell.LastHeight = -1
	}

	v.PixelData = make([]byte, v.Box.Width*v.Box.Height)
	frameIndex := -1

	for _, frame := range v.Frames {
		frameIndex++

		frame.PixelData = make([]byte, v.Box.Width*v.Box.Height)
		c := -1

		for _, cell := range frame.Cells {
			c++

			cellX := cell.XOffset / cellsPerRow
			cellY := cell.YOffset / cellsPerRow
			cellIndex := cellX + (cellY * v.HorizontalCellCount)
			bufferCell := v.Cells[cellIndex]
			pbe := v.PixelBuffer[pbIdx]

			if (pbe.Frame != frameIndex) || (pbe.FrameCellIndex != c) {
				// This buffer cell has an EqualCell bit set to 1, so copy the frame cell or clear it
				if (cell.Width != bufferCell.LastWidth) || (cell.Height != bufferCell.LastHeight) {
					// Different sizes
					for y := 0; y < cell.Height; y++ {
						for x := 0; x < cell.Width; x++ {
							v.PixelData[x+cell.XOffset+((y+cell.YOffset)*v.Box.Width)] = 0
						}
					}
				} else {
					// Same sizes
					// Copy the old frame cell into the new position
					for fy := 0; fy < cell.Height; fy++ {
						for fx := 0; fx < cell.Width; fx++ {
							// Frame (buff.lastx, buff.lasty) -> Frame (cell.offx, cell.offy)
							// Cell (0, 0,) ->
							// blit(dir->bmp, dir->bmp, buff_cell->last_x0, buff_cell->last_y0, cell->x0, cell->y0, cell->w, cell->h );
							v.PixelData[fx+cell.XOffset+((fy+cell.YOffset)*v.Box.Width)] =
								v.PixelData[fx+bufferCell.LastXOffset+((fy+bufferCell.LastYOffset)*v.Box.Width)]
						}
					}
					// Copy it again into the final frame image
					for fy := 0; fy < cell.Height; fy++ {
						for fx := 0; fx < cell.Width; fx++ {
							// blit(cell->bmp, frm_bmp, 0, 0, cell->x0, cell->y0, cell->w, cell->h );
							frame.PixelData[fx+cell.XOffset+((fy+cell.YOffset)*v.Box.Width)] = v.PixelData[cell.XOffset+fx+((cell.YOffset+fy)*v.Box.Width)]
						}
					}
				}
			} else {
				if pbe.Value[0] == pbe.Value[1] {
					// Clear the frame
					for y := 0; y < cell.Height; y++ {
						for x := 0; x < cell.Width; x++ {
							v.PixelData[x+cell.XOffset+((y+cell.YOffset)*v.Box.Width)] = pbe.Value[0]
						}
					}
				} else {
					// Fill the frame cell with the pixels
					bitsToRead := 1
					if pbe.Value[1] != pbe.Value[2] {
						bitsToRead = 2
					}
					for y := 0; y < cell.Height; y++ {
						for x := 0; x < cell.Width; x++ {
							paletteIndex := pcd.GetBits(bitsToRead)
							v.PixelData[x+cell.XOffset+((y+cell.YOffset)*v.Box.Width)] = pbe.Value[paletteIndex]
						}
					}
				}

				// Copy the frame cell into the frame
				for fy := 0; fy < cell.Height; fy++ {
					for fx := 0; fx < cell.Width; fx++ {
						// blit(cell->bmp, frm_bmp, 0, 0, cell->x0, cell->y0, cell->w, cell->h );
						frame.PixelData[fx+cell.XOffset+((fy+cell.YOffset)*v.Box.Width)] = v.PixelData[fx+cell.XOffset+((fy+cell.YOffset)*v.Box.Width)]
					}
				}
				pbIdx++
			}

			bufferCell.LastWidth = cell.Width
			bufferCell.LastHeight = cell.Height
			bufferCell.LastXOffset = cell.XOffset
			bufferCell.LastYOffset = cell.YOffset
		}

		// Free up the stuff we no longer need
		frame.Cells = nil
	}

	v.Cells = nil
	v.PixelData = nil
	v.PixelBuffer = nil
}

//nolint:funlen,gocognit,gocyclo // can't reduce
func (v *DCCDirection) fillPixelBuffer(pcd, ec, pm, et, rp *d2datautils.BitMuncher) {
	var pixelMaskLookup = []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}

	lastPixel := uint32(0)
	maxCellX := 0
	maxCellY := 0

	for _, frame := range v.Frames {
		if frame == nil {
			continue
		}

		maxCellX += frame.HorizontalCellCount
		maxCellY += frame.VerticalCellCount
	}

	v.PixelBuffer = make([]DCCPixelBufferEntry, maxCellX*maxCellY)

	for i := 0; i < maxCellX*maxCellY; i++ {
		v.PixelBuffer[i].Frame = -1
		v.PixelBuffer[i].FrameCellIndex = -1
	}

	cellBuffer := make([]*DCCPixelBufferEntry, v.HorizontalCellCount*v.VerticalCellCount)
	frameIndex := -1
	pbIndex := -1

	var pixelMask uint32

	for _, frame := range v.Frames {
		frameIndex++

		originCellX := (frame.Box.Left - v.Box.Left) / cellsPerRow
		originCellY := (frame.Box.Top - v.Box.Top) / cellsPerRow

		for cellY := 0; cellY < frame.VerticalCellCount; cellY++ {
			currentCellY := cellY + originCellY

			for cellX := 0; cellX < frame.HorizontalCellCount; cellX++ {
				currentCell := originCellX + cellX + (currentCellY * v.HorizontalCellCount)
				nextCell := false
				tmp := 0

				if cellBuffer[currentCell] != nil {
					if v.EqualCellsBitstreamSize > 0 {
						tmp = int(ec.GetBit())
					} else {
						tmp = 0
					}

					if tmp == 0 {
						pixelMask = pm.GetBits(4) //nolint:gomnd // binary data
					} else {
						nextCell = true
					}
				} else {
					pixelMask = 0x0F
				}

				if nextCell {
					continue
				}

				// Decode the pixels
				var pixelStack [4]uint32

				lastPixel = 0
				numberOfPixelBits := pixelMaskLookup[pixelMask]
				encodingType := 0

				if (numberOfPixelBits != 0) && (v.EncodingTypeBitsreamSize > 0) {
					encodingType = int(et.GetBit())
				} else {
					encodingType = 0
				}

				decodedPixel := 0

				for i := 0; i < numberOfPixelBits; i++ {
					if encodingType != 0 {
						pixelStack[i] = rp.GetBits(8) //nolint:gomnd // binary data
					} else {
						pixelStack[i] = lastPixel
						pixelDisplacement := pcd.GetBits(4) //nolint:gomnd // binary data
						pixelStack[i] += pixelDisplacement
						for pixelDisplacement == 15 {
							pixelDisplacement = pcd.GetBits(4) //nolint:gomnd // binary data
							pixelStack[i] += pixelDisplacement
						}
					}

					if pixelStack[i] == lastPixel {
						pixelStack[i] = 0
						break
					} else {
						lastPixel = pixelStack[i]
						decodedPixel++
					}
				}

				oldEntry := cellBuffer[currentCell]

				pbIndex++

				curIdx := decodedPixel - 1

				for i := 0; i < 4; i++ {
					if (pixelMask & (1 << uint(i))) != 0 {
						if curIdx >= 0 {
							v.PixelBuffer[pbIndex].Value[i] = byte(pixelStack[curIdx])
							curIdx--
						} else {
							v.PixelBuffer[pbIndex].Value[i] = 0
						}
					} else {
						v.PixelBuffer[pbIndex].Value[i] = oldEntry.Value[i]
					}
				}

				cellBuffer[currentCell] = &v.PixelBuffer[pbIndex]
				v.PixelBuffer[pbIndex].Frame = frameIndex
				v.PixelBuffer[pbIndex].FrameCellIndex = cellX + (cellY * frame.HorizontalCellCount)
			}
		}
	}

	// Convert the palette entry index into actual palette entries
	for i := 0; i <= pbIndex; i++ {
		for x := 0; x < 4; x++ {
			v.PixelBuffer[i].Value[x] = v.PaletteEntries[v.PixelBuffer[i].Value[x]]
		}
	}
}

func (v *DCCDirection) calculateCells() {
	// Calculate the number of vertical and horizontal cells we need
	v.HorizontalCellCount = 1 + (v.Box.Width-1)/cellsPerRow
	v.VerticalCellCount = 1 + (v.Box.Height-1)/cellsPerRow
	// Calculate the cell widths
	cellWidths := make([]int, v.HorizontalCellCount)
	if v.HorizontalCellCount == 1 {
		cellWidths[0] = v.Box.Width
	} else {
		for i := 0; i < v.HorizontalCellCount-1; i++ {
			cellWidths[i] = 4
		}
		cellWidths[v.HorizontalCellCount-1] = v.Box.Width - (4 * (v.HorizontalCellCount - 1))
	}
	// Calculate the cell heights
	cellHeights := make([]int, v.VerticalCellCount)
	if v.VerticalCellCount == 1 {
		cellHeights[0] = v.Box.Height
	} else {
		for i := 0; i < v.VerticalCellCount-1; i++ {
			cellHeights[i] = 4
		}
		cellHeights[v.VerticalCellCount-1] = v.Box.Height - (4 * (v.VerticalCellCount - 1))
	}
	// Set the cell widths and heights in the cell buffer
	v.Cells = make([]*DCCCell, v.VerticalCellCount*v.HorizontalCellCount)
	yOffset := 0

	for y := 0; y < v.VerticalCellCount; y++ {
		xOffset := 0

		for x := 0; x < v.HorizontalCellCount; x++ {
			v.Cells[x+(y*v.HorizontalCellCount)] = &DCCCell{
				Width:   cellWidths[x],
				Height:  cellHeights[y],
				XOffset: xOffset,
				YOffset: yOffset,
			}

			xOffset += 4
		}

		yOffset += 4
	}
}
