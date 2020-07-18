package d2dc6

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

const (
	endOfScanLine = 0x80
	maxRunLength  = 0x7f
)

// DC6 represents a DC6 file.
type DC6 struct {
	Version            int32
	Flags              uint32
	Encoding           uint32
	Termination        []byte // 4 bytes
	Directions         uint32
	FramesPerDirection uint32
	FramePointers      []uint32    // size is Directions*FramesPerDirection
	Frames             []*DC6Frame // size is Directions*FramesPerDirection
}

// Load uses restruct to read the binary dc6 data into structs then parses image data from the frame data.
func Load(data []byte) (*DC6, error) {
	const (
		terminationSize = 4
		terminatorSize  = 3
	)

	r := d2common.CreateStreamReader(data)

	var dc DC6
	dc.Version = r.GetInt32()
	dc.Flags = r.GetUInt32()
	dc.Encoding = r.GetUInt32()
	dc.Termination = r.ReadBytes(terminationSize)
	dc.Directions = r.GetUInt32()
	dc.FramesPerDirection = r.GetUInt32()

	frameCount := int(dc.Directions * dc.FramesPerDirection)

	dc.FramePointers = make([]uint32, frameCount)
	for i := 0; i < frameCount; i++ {
		dc.FramePointers[i] = r.GetUInt32()
	}

	dc.Frames = make([]*DC6Frame, frameCount)

	for i := 0; i < frameCount; i++ {
		frame := &DC6Frame{
			Flipped:   r.GetUInt32(),
			Width:     r.GetUInt32(),
			Height:    r.GetUInt32(),
			OffsetX:   r.GetInt32(),
			OffsetY:   r.GetInt32(),
			Unknown:   r.GetUInt32(),
			NextBlock: r.GetUInt32(),
			Length:    r.GetUInt32(),
		}
		frame.FrameData = r.ReadBytes(int(frame.Length))
		frame.Terminator = r.ReadBytes(terminatorSize)
		dc.Frames[i] = frame
	}

	return &dc, nil
}

// DecodeFrame decodes the given frame to an indexed color texture
func (d *DC6) DecodeFrame(frameIndex int) []byte {
	frame := d.Frames[frameIndex]

	indexData := make([]byte, frame.Width*frame.Height)
	x := 0
	y := int(frame.Height) - 1
	offset := 0

	for {
		b := int(frame.FrameData[offset])
		offset++

		if b == endOfScanLine {
			if y == 0 {
				break
			}

			y--

			x = 0
		} else if b&endOfScanLine > 0 {
			transparentPixels := b & maxRunLength
			x += transparentPixels
		} else {
			for i := 0; i < b; i++ {
				indexData[x+y*int(frame.Width)+i] = frame.FrameData[offset]
				offset++
			}

			x += b
		}
	}

	return indexData
}
