package d2dc6

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
)

const (
	endOfScanLine = 0x80
	maxRunLength  = 0x7f
)

type scanlineState int

const (
	endOfLine scanlineState = iota
	runOfTransparentPixels
	runOfOpaquePixels
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

	r := d2datautils.CreateStreamReader(data)

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

loop: // this is a label for the loop, so the switch can break the loop (and not the switch)
	for {
		b := int(frame.FrameData[offset])
		offset++

		switch scanlineType(b) {
		case endOfLine:
			if y == 0 {
				break loop
			}

			y--

			x = 0
		case runOfTransparentPixels:
			transparentPixels := b & maxRunLength
			x += transparentPixels
		case runOfOpaquePixels:
			for i := 0; i < b; i++ {
				indexData[x+y*int(frame.Width)+i] = frame.FrameData[offset]
				offset++
			}

			x += b
		}
	}

	return indexData
}

func scanlineType(b int) scanlineState {
	if b == endOfScanLine {
		return endOfLine
	}

	if (b & endOfScanLine) > 0 {
		return runOfTransparentPixels
	}

	return runOfOpaquePixels
}

// Clone creates a copy of the DC6
func (d *DC6) Clone() *DC6 {
	clone := *d
	copy(clone.Termination, d.Termination)
	copy(clone.FramePointers, d.FramePointers)
	clone.Frames = make([]*DC6Frame, len(d.Frames))

	for i := range d.Frames {
		cloneFrame := *d.Frames[i]
		clone.Frames = append(clone.Frames, &cloneFrame)
	}

	return &clone
}
