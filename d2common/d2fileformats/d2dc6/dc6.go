package d2dc6

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
)

const (
	endOfScanLine = 0x80
	maxRunLength  = 0x7f

	terminationSize = 4
	terminatorSize  = 3
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

// New creates a new, empty DC6
func New() *DC6 {
	result := &DC6{
		Version:            0,
		Flags:              0,
		Encoding:           0,
		Termination:        make([]byte, 4),
		Directions:         0,
		FramesPerDirection: 0,
		FramePointers:      make([]uint32, 0),
		Frames:             make([]*DC6Frame, 0),
	}

	return result
}

// Load loads a dc6 animation
func Load(data []byte) (*DC6, error) {
	d := New()
	err := d.Unmarshal(data)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// Unmarshal converts bite slice into DC6 structure
func (dc *DC6) Unmarshal(data []byte) error {
	var err error

	r := d2datautils.CreateStreamReader(data)

	err = dc.loadHeader(r)
	if err != nil {
		return err
	}

	frameCount := int(dc.Directions * dc.FramesPerDirection)

	dc.FramePointers = make([]uint32, frameCount)
	for i := 0; i < frameCount; i++ {
		dc.FramePointers[i], err = r.ReadUInt32()
		if err != nil {
			return err
		}
	}

	dc.Frames = make([]*DC6Frame, frameCount)

	if err := dc.loadFrames(r); err != nil {
		return err
	}

	return nil
}

func (d *DC6) loadHeader(r *d2datautils.StreamReader) error {
	var err error

	if d.Version, err = r.ReadInt32(); err != nil {
		return err
	}

	if d.Flags, err = r.ReadUInt32(); err != nil {
		return err
	}

	if d.Encoding, err = r.ReadUInt32(); err != nil {
		return err
	}

	if d.Termination, err = r.ReadBytes(terminationSize); err != nil {
		return err
	}

	if d.Directions, err = r.ReadUInt32(); err != nil {
		return err
	}

	if d.FramesPerDirection, err = r.ReadUInt32(); err != nil {
		return err
	}

	return nil
}

func (d *DC6) loadFrames(r *d2datautils.StreamReader) error {
	var err error

	for i := 0; i < len(d.FramePointers); i++ {
		frame := &DC6Frame{}

		if frame.Flipped, err = r.ReadUInt32(); err != nil {
			return err
		}

		if frame.Width, err = r.ReadUInt32(); err != nil {
			return err
		}

		if frame.Height, err = r.ReadUInt32(); err != nil {
			return err
		}

		if frame.OffsetX, err = r.ReadInt32(); err != nil {
			return err
		}

		if frame.OffsetY, err = r.ReadInt32(); err != nil {
			return err
		}

		if frame.Unknown, err = r.ReadUInt32(); err != nil {
			return err
		}

		if frame.NextBlock, err = r.ReadUInt32(); err != nil {
			return err
		}

		if frame.Length, err = r.ReadUInt32(); err != nil {
			return err
		}

		if frame.FrameData, err = r.ReadBytes(int(frame.Length)); err != nil {
			return err
		}

		if frame.Terminator, err = r.ReadBytes(terminatorSize); err != nil {
			return err
		}

		d.Frames[i] = frame
	}

	return nil
}

// Marshal encodes dc6 animation back into byte slice
func (d *DC6) Marshal() []byte {
	sw := d2datautils.CreateStreamWriter()

	// Encode header
	sw.PushInt32(d.Version)
	sw.PushUint32(d.Flags)
	sw.PushUint32(d.Encoding)

	sw.PushBytes(d.Termination...)

	sw.PushUint32(d.Directions)
	sw.PushUint32(d.FramesPerDirection)

	// load frames
	for _, i := range d.FramePointers {
		sw.PushUint32(i)
	}

	for i := range d.Frames {
		sw.PushUint32(d.Frames[i].Flipped)
		sw.PushUint32(d.Frames[i].Width)
		sw.PushUint32(d.Frames[i].Height)
		sw.PushInt32(d.Frames[i].OffsetX)
		sw.PushInt32(d.Frames[i].OffsetY)
		sw.PushUint32(d.Frames[i].Unknown)
		sw.PushUint32(d.Frames[i].NextBlock)
		sw.PushUint32(d.Frames[i].Length)
		sw.PushBytes(d.Frames[i].FrameData...)
		sw.PushBytes(d.Frames[i].Terminator...)
	}

	return sw.GetBytes()
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
		clone.Frames[i] = &cloneFrame
	}

	return &clone
}
