package d2dc6

import (
	"encoding/binary"

	"github.com/go-restruct/restruct"
)

// DC6 represents a DC6 file.
type DC6 struct {
	// Header
	Version            int32  `struct:"int32"`
	Flags              uint32 `struct:"uint32"`
	Encoding           uint32 `struct:"uint32"`
	Termination        []byte `struct:"[4]byte"`
	Directions         uint32 `struct:"uint32"`
	FramesPerDirection uint32 `struct:"uint32"`

	FramePointers []uint32    `struct:"[]uint32,size=Directions*FramesPerDirection"`
	Frames        []*DC6Frame `struct-size:"Directions*FramesPerDirection"`
}

// Load uses restruct to read the binary dc6 data into structs then parses image data from the frame data.
func Load(data []byte) (*DC6, error) {
	result := &DC6{}

	restruct.EnableExprBeta()
	err := restruct.Unpack(data, binary.LittleEndian, &result)

	if err != nil {
		return nil, err
	}

	return result, err
}

// Decodes the given frame to an indexed color texture
func (d *DC6) DecodeFrame(frameIndex int) []byte {
	frame := d.Frames[frameIndex]

	indexData := make([]byte, frame.Width*frame.Height)
	x := 0
	y := int(frame.Height) - 1
	offset := 0

	for {
		b := int(frame.FrameData[offset])
		offset++

		if b == 0x80 {
			if y == 0 {
				break
			}

			y--

			x = 0
		} else if b&0x80 > 0 {
			transparentPixels := b & 0x7f
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
