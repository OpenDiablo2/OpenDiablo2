package common

import (
	"log"
)

type DCCPixelBufferEntry struct {
	Value          []byte
	Frame          int
	FrameCellIndex int
}

type DCCCell struct {
	Width       int
	Height      int
	XOffset     int
	YOffset     int
	LastWidth   int
	LastHeight  int
	LastXOffset int
	LastYOffset int
}

type DCCDirectionFrame struct {
	Width                 int
	Height                int
	XOffset               int
	YOffset               int
	NumberOfOptionalBytes int
	NumberOfCodedBytes    int
	FrameIsBottomUp       bool
	Box                   Rectangle
	Cells                 []DCCCell
	PixelData             []byte
	HorizontalCellCount   int
	VerticalCellCount     int
}

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
	Box                        Rectangle
	Cells                      []*DCCCell
	PixelData                  []byte
	HorizontalCellCount        int
	VerticalCellCount          int
	PixelBuffer                []*DCCPixelBufferEntry
}

type DCC struct {
	Signature          int
	Version            int
	NumberOfDirections int
	FramesPerDirection int
	Directions         []*DCCDirection
}

var crazyBitTable = []byte{0, 1, 2, 4, 6, 8, 10, 12, 14, 16, 20, 24, 26, 28, 30, 32}
var pixelMaskLookup = []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4}
var dccDir4 = []byte{0, 1, 2, 3}
var dccDir8 = []byte{4, 0, 5, 1, 6, 2, 7, 3}
var dccDir16 = []byte{4, 8, 0, 9, 5, 10, 1, 11, 6, 12, 2, 13, 7, 14, 3, 15}
var dccDir32 = []byte{4, 16, 8, 17, 0, 18, 9, 19, 5, 20, 10, 21, 1, 22, 11, 23,
	6, 24, 12, 25, 2, 26, 13, 27, 7, 28, 14, 29, 3, 30, 15, 31}

func CreateDCCDirectionFrame(bits *BitMuncher, direction *DCCDirection) *DCCDirectionFrame {
	result := &DCCDirectionFrame{}
	bits.GetBits(direction.Variable0Bits) // Variable0
	result.Width = int(bits.GetBits(direction.WidthBits))
	result.Height = int(bits.GetBits(direction.HeightBits))
	result.XOffset = bits.GetSignedBits(direction.XOffsetBits)
	result.YOffset = bits.GetSignedBits(direction.YOffsetBits)
	result.NumberOfOptionalBytes = int(bits.GetBits(direction.OptionalDataBits))
	result.NumberOfCodedBytes = int(bits.GetBits(direction.CodedBytesBits))
	result.FrameIsBottomUp = bits.GetBit() == 1
	if result.FrameIsBottomUp {
		log.Panic("Bottom up frames are not implemented.")
	} else {
		result.Box = Rectangle{
			result.XOffset,
			result.YOffset - result.Height + 1,
			result.Width,
			result.Height,
		}
	}
	return result
}

func (v *DCCDirectionFrame) CalculateCells(direction *DCCDirection) {
	var w = 4 - ((v.Box.Left - direction.Box.Left) % 4) // Width of the first column (in pixels)
	if (v.Width - w) <= 1 {
		v.HorizontalCellCount = 1
	} else {
		tmp := v.Width - w - 1
		v.HorizontalCellCount = 2 + (tmp / 4)
		if (tmp % 4) == 0 {
			v.HorizontalCellCount--
		}
	}
	h := 4 - ((v.Box.Top - direction.Box.Top) % 4) // Height of the first column (in pixels)
	if (v.Height - h) <= 1 {
		v.VerticalCellCount = 1
	} else {
		tmp := v.Height - h - 1
		v.VerticalCellCount = 2 + (tmp / 4)
		if (tmp % 4) == 0 {
			v.VerticalCellCount--
		}
	}
	// Calculate the cell widths and heights
	cellWidths := make([]int, v.HorizontalCellCount)
	if v.HorizontalCellCount == 1 {
		cellWidths[0] = v.Width
	} else {
		cellWidths[0] = w
		for i := 1; i < (v.HorizontalCellCount - 1); i++ {
			cellWidths[i] = 4
		}
		cellWidths[v.HorizontalCellCount-1] = v.Width - w - (4 * (v.HorizontalCellCount - 2))
	}

	cellHeights := make([]int, v.VerticalCellCount)
	if v.VerticalCellCount == 1 {
		cellHeights[0] = v.Height
	} else {
		cellHeights[0] = h
		for i := 1; i < (v.VerticalCellCount - 1); i++ {
			cellHeights[i] = 4
		}
		cellHeights[v.VerticalCellCount-1] = v.Height - h - (4 * (v.VerticalCellCount - 2))
	}

	v.Cells = make([]DCCCell, v.HorizontalCellCount*v.VerticalCellCount)
	offsetY := v.Box.Top - direction.Box.Top
	for y := 0; y < v.VerticalCellCount; y++ {
		offsetX := v.Box.Left - direction.Box.Left
		for x := 0; x < v.HorizontalCellCount; x++ {
			v.Cells[x+(y*v.HorizontalCellCount)] = DCCCell{
				XOffset: offsetX,
				YOffset: offsetY,
				Width:   cellWidths[x],
				Height:  cellHeights[y],
			}
			offsetX += cellWidths[x]
		}
		offsetY += cellHeights[y]
	}
}

func CreateDCCDirection(bm *BitMuncher, file *DCC) *DCCDirection {
	result := &DCCDirection{}
	result.OutSizeCoded = int(bm.GetUInt32())
	result.CompressionFlags = int(bm.GetBits(2))
	result.Variable0Bits = int(crazyBitTable[bm.GetBits(4)])
	result.WidthBits = int(crazyBitTable[bm.GetBits(4)])
	result.HeightBits = int(crazyBitTable[bm.GetBits(4)])
	result.XOffsetBits = int(crazyBitTable[bm.GetBits(4)])
	result.YOffsetBits = int(crazyBitTable[bm.GetBits(4)])
	result.OptionalDataBits = int(crazyBitTable[bm.GetBits(4)])
	result.CodedBytesBits = int(crazyBitTable[bm.GetBits(4)])
	result.Frames = make([]*DCCDirectionFrame, file.FramesPerDirection)
	minx := 100000
	miny := 100000
	maxx := -100000
	maxy := -100000
	// Load the frame headers
	for frameIdx := 0; frameIdx < file.FramesPerDirection; frameIdx++ {
		result.Frames[frameIdx] = CreateDCCDirectionFrame(bm, result)
		minx = int(MinInt32(int32(result.Frames[frameIdx].Box.Left), int32(minx)))
		miny = int(MinInt32(int32(result.Frames[frameIdx].Box.Top), int32(miny)))
		maxx = int(MaxInt32(int32(result.Frames[frameIdx].Box.Right()), int32(maxx)))
		maxy = int(MaxInt32(int32(result.Frames[frameIdx].Box.Bottom()), int32(maxy)))
	}
	result.Box = Rectangle{minx, miny, (maxx - minx), (maxy - miny)}
	if result.OptionalDataBits > 0 {
		log.Panic("Optional bits in DCC data is not currently supported.")
	}
	if (result.CompressionFlags & 0x2) > 0 {
		result.EqualCellsBitstreamSize = int(bm.GetBits(20))
	}
	result.PixelMaskBitstreamSize = int(bm.GetBits(20))
	if (result.CompressionFlags & 0x1) > 0 {
		result.EncodingTypeBitsreamSize = int(bm.GetBits(20))
		result.RawPixelCodesBitstreamSize = int(bm.GetBits(20))
	}
	// PixelValuesKey
	paletteEntryCount := 0
	for i := 0; i < 256; i++ {
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
	equalCellsBitstream := CopyBitMuncher(bm)
	bm.SkipBits(result.EqualCellsBitstreamSize)
	pixelMaskBitstream := CopyBitMuncher(bm)
	bm.SkipBits(result.PixelMaskBitstreamSize)
	encodingTypeBitsream := CopyBitMuncher(bm)
	bm.SkipBits(result.EncodingTypeBitsreamSize)
	rawPixelCodesBitstream := CopyBitMuncher(bm)
	bm.SkipBits(result.RawPixelCodesBitstreamSize)
	pixelCodeandDisplacement := CopyBitMuncher(bm)
	// Calculate the cells for the direction
	result.CalculateCells()
	// Calculate the cells for each of the frames
	for _, frame := range result.Frames {
		frame.CalculateCells(result)
	}
	// Fill in the pixel buffer
	result.FillPixelBuffer(pixelCodeandDisplacement, equalCellsBitstream, pixelMaskBitstream, encodingTypeBitsream, rawPixelCodesBitstream)
	// Generate the actual frame pixel data
	result.GenerateFrames(pixelCodeandDisplacement)
	result.PixelBuffer = nil
	// Verify that everything we expected to read was actually read (sanity check)...
	if equalCellsBitstream.BitsRead != result.EqualCellsBitstreamSize {
		log.Panic("Did not read the correct number of bits!")
	}
	if pixelMaskBitstream.BitsRead != result.PixelMaskBitstreamSize {
		log.Panic("Did not read the correct number of bits!")
	}
	if encodingTypeBitsream.BitsRead != result.EncodingTypeBitsreamSize {
		log.Panic("Did not read the correct number of bits!")
	}
	if rawPixelCodesBitstream.BitsRead != result.RawPixelCodesBitstreamSize {
		log.Panic("Did not read the correct number of bits!")
	}
	bm.SkipBits(pixelCodeandDisplacement.BitsRead)
	return result
}

func (v *DCCDirection) GenerateFrames(pcd *BitMuncher) {
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
			cellX := cell.XOffset / 4
			cellY := cell.YOffset / 4
			cellIndex := cellX + (cellY * v.HorizontalCellCount)
			bufferCell := v.Cells[cellIndex]
			pbe := v.PixelBuffer[pbIdx]
			if (pbe.Frame != frameIndex) || (pbe.FrameCellIndex != c) {
				// This buffer cell has an EqualCell bit set to 1, so copy the frame cell or clear it
				if (cell.Width != bufferCell.LastWidth) || (cell.Height != bufferCell.LastHeight) {
					// Different sizes
					/// TODO: Clear the pixels of the frame cell
					for y := 0; y < cell.Height; y++ {
						for x := 0; x < cell.Width; x++ {
							v.PixelData[x+cell.XOffset+((y+cell.YOffset)*frame.Width)] = 0
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
							v.PixelData[fx+cell.XOffset+((fy+cell.YOffset)*v.Box.Width)] = v.PixelData[fx+bufferCell.LastXOffset+((fy+bufferCell.LastYOffset)*v.Box.Width)]
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
					//cell.PixelData = new byte[cell.Width * cell.Height];
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
						//blit(cell->bmp, frm_bmp, 0, 0, cell->x0, cell->y0, cell->w, cell->h );
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

func (v *DCCDirection) FillPixelBuffer(pcd, ec, pm, et, rp *BitMuncher) {
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
	v.PixelBuffer = make([]*DCCPixelBufferEntry, maxCellX*maxCellY)
	for i := 0; i < maxCellX*maxCellY; i++ {
		v.PixelBuffer[i] = &DCCPixelBufferEntry{
			Value:          make([]byte, 4),
			Frame:          -1,
			FrameCellIndex: -1,
		}
	}
	cellBuffer := make([]*DCCPixelBufferEntry, v.HorizontalCellCount*v.VerticalCellCount)
	frameIndex := -1
	pbIndex := -1
	pixelMask := uint32(0x00)
	for _, frame := range v.Frames {
		frameIndex++
		originCellX := (frame.Box.Left - v.Box.Left) / 4
		originCellY := (frame.Box.Top - v.Box.Top) / 4
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
						pixelMask = pm.GetBits(4)
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
						pixelStack[i] = rp.GetBits(8)
					} else {
						pixelStack[i] = lastPixel
						pixelDisplacement := pcd.GetBits(4)
						pixelStack[i] += pixelDisplacement
						for pixelDisplacement == 15 {
							pixelDisplacement = pcd.GetBits(4)
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
				newEntry := v.PixelBuffer[pbIndex]
				curIdx := decodedPixel - 1
				for i := 0; i < 4; i++ {
					if (pixelMask & (1 << uint(i))) != 0 {
						if curIdx >= 0 {
							newEntry.Value[i] = byte(pixelStack[curIdx])
							curIdx--
						} else {
							newEntry.Value[i] = 0
						}
					} else {
						newEntry.Value[i] = oldEntry.Value[i]
					}
				}
				cellBuffer[currentCell] = newEntry
				newEntry.Frame = frameIndex
				newEntry.FrameCellIndex = cellX + (cellY * frame.HorizontalCellCount)
			}
		}
	}
	cellBuffer = nil
	// Convert the palette entry index into actual palette entries
	for i := 0; i <= pbIndex; i++ {
		for x := 0; x < 4; x++ {
			v.PixelBuffer[i].Value[x] = v.PaletteEntries[v.PixelBuffer[i].Value[x]]
		}
	}
}

func (v *DCCDirection) CalculateCells() {
	// Calculate the number of vertical and horizontal cells we need
	v.HorizontalCellCount = 1 + (v.Box.Width-1)/4
	v.VerticalCellCount = 1 + (v.Box.Height-1)/4
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

func LoadDCC(path string, fileProvider FileProvider) *DCC {
	result := &DCC{}
	fileData := fileProvider.LoadFile(path)
	var bm = CreateBitMuncher(fileData, 0)
	result.Signature = int(bm.GetByte())
	if result.Signature != 0x74 {
		log.Fatal("Signature expected to be 0x74 but it is not.")
	}
	result.Version = int(bm.GetByte())
	result.NumberOfDirections = int(bm.GetByte())
	result.FramesPerDirection = int(bm.GetInt32())
	if bm.GetInt32() != 1 {
		log.Fatal("This value isn't 1. It has to be 1.")
	}
	bm.GetInt32() // TotalSizeCoded
	directionOffsets := make([]int, result.NumberOfDirections)
	for i := 0; i < result.NumberOfDirections; i++ {
		directionOffsets[i] = int(bm.GetInt32())
	}
	result.Directions = make([]*DCCDirection, result.NumberOfDirections)
	for i := 0; i < result.NumberOfDirections; i++ {
		dir := byte(0)
		switch result.NumberOfDirections {
		case 1:
			dir = 0
		case 4:
			dir = dccDir4[i]
		case 8:
			dir = dccDir8[i]
		case 16:
			dir = dccDir16[i]
		case 32:
			dir = dccDir32[i]
		}
		result.Directions[dir] = CreateDCCDirection(CreateBitMuncher(fileData, directionOffsets[i]*8), result)
	}

	return result
}
