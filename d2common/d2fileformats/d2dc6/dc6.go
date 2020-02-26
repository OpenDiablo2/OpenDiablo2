package d2dc6

import (
	"encoding/binary"

	"github.com/go-restruct/restruct"
)

type DC6File struct {
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

type DC6Header struct {
	Version            int32  `struct:"int32"`
	Flags              uint32 `struct:"uint32"`
	Encoding           uint32 `struct:"uint32"`
	Termination        []byte `struct:"[4]byte"`
	Directions         int32  `struct:"int32"`
	FramesPerDirection int32  `struct:"int32"`
}

type DC6FrameHeader struct {
	Flipped   int32  `struct:"int32"`
	Width     int32  `struct:"int32"`
	Height    int32  `struct:"int32"`
	OffsetX   int32  `struct:"int32"`
	OffsetY   int32  `struct:"int32"`
	Unknown   uint32 `struct:"uint32"`
	NextBlock uint32 `struct:"uint32"`
	Length    uint32 `struct:"uint32"`
}

type DC6Frame struct {
	Flipped    uint32 `struct:"uint32"`
	Width      uint32 `struct:"uint32"`
	Height     uint32 `struct:"uint32"`
	OffsetX    int32  `struct:"int32"`
	OffsetY    int32  `struct:"int32"`
	Unknown    uint32 `struct:"uint32"`
	NextBlock  uint32 `struct:"uint32"`
	Length     uint32 `struct:"uint32,sizeof=FrameData"`
	FrameData  []byte
	Terminator []byte `struct:"[3]byte"`
}

// LoadDC6 uses restruct to read the binary dc6 data into structs then parses image data from the frame data.
func LoadDC6(data []byte) (*DC6File, error) {
	result := &DC6File{}

	restruct.EnableExprBeta()
	err := restruct.Unpack(data, binary.LittleEndian, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}
