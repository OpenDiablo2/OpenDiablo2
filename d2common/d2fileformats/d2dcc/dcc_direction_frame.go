package d2dcc

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
)

// DCCDirectionFrame represents a direction frame for a DCC.
type DCCDirectionFrame struct {
	Box                   d2geom.Rectangle
	Cells                 []DCCCell
	PixelData             []byte
	Width                 int
	Height                int
	XOffset               int
	YOffset               int
	NumberOfOptionalBytes int
	NumberOfCodedBytes    int
	HorizontalCellCount   int
	VerticalCellCount     int
	FrameIsBottomUp       bool
	valid                 bool
}

// CreateDCCDirectionFrame Creates a DCCDirectionFrame for a DCC.
func CreateDCCDirectionFrame(bits *d2datautils.BitMuncher, direction *DCCDirection) *DCCDirectionFrame {
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
		result.Box = d2geom.Rectangle{
			Left:   result.XOffset,
			Top:    result.YOffset - result.Height + 1,
			Width:  result.Width,
			Height: result.Height,
		}
	}

	result.valid = true

	return result
}

func (v *DCCDirectionFrame) recalculateCells(direction *DCCDirection) {
	var w = 4 - ((v.Box.Left - direction.Box.Left) % 4) // Width of the first column (in pixels)

	if (v.Width - w) <= 1 {
		v.HorizontalCellCount = 1
	} else {
		tmp := v.Width - w - 1
		v.HorizontalCellCount = 2 + (tmp / 4) //nolint:gomnd // magic math
		if (tmp % 4) == 0 {
			v.HorizontalCellCount--
		}
	}

	// Height of the first column (in pixels)
	h := 4 - ((v.Box.Top - direction.Box.Top) % 4) //nolint:gomnd // data decode

	if (v.Height - h) <= 1 {
		v.VerticalCellCount = 1
	} else {
		tmp := v.Height - h - 1
		v.VerticalCellCount = 2 + (tmp / 4) //nolint:gomnd // data decode
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
