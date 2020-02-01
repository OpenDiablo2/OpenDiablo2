package d2dc6

import (
	"encoding/binary"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
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
	valid         bool
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

	colorData []byte
	palette   d2datadict.PaletteRec
	valid     bool
}

func (frame *DC6Frame) ColorData() []byte {
	if frame.colorData == nil {
		frame.completeLoad()
	}

	return frame.colorData
}

func (frame *DC6Frame) completeLoad() {
	frame.valid = true

	indexData := make([]int16, frame.Width*frame.Height)
	for fi := range indexData {
		indexData[fi] = -1
	}

	x := uint32(0)
	y := frame.Height - 1
	dataPointer := 0
	for {
		b := frame.FrameData[dataPointer]
		dataPointer++
		if b == 0x80 {
			if y == 0 {
				break
			}
			y--
			x = 0
		} else if (b & 0x80) > 0 {
			transparentPixels := b & 0x7F
			for ti := byte(0); ti < transparentPixels; ti++ {
				indexData[x+(y*frame.Width)+uint32(ti)] = -1
			}
			x += uint32(transparentPixels)
		} else {
			for bi := 0; bi < int(b); bi++ {
				indexData[x+(y*frame.Width)+uint32(bi)] = int16(frame.FrameData[dataPointer])
				dataPointer++
			}
			x += uint32(b)
		}
	}

	// Probably don't need this data again
	frame.FrameData = nil

	frame.colorData = make([]byte, int(frame.Width*frame.Height)*4)
	for i := uint32(0); i < frame.Width*frame.Height; i++ {
		if indexData[i] < 1 { // TODO: Is this == -1 or < 1?
			continue
		}
		frame.colorData[i*4] = frame.palette.Colors[indexData[i]].R
		frame.colorData[(i*4)+1] = frame.palette.Colors[indexData[i]].G
		frame.colorData[(i*4)+2] = frame.palette.Colors[indexData[i]].B
		frame.colorData[(i*4)+3] = 0xFF
	}
}

// LoadDC6 uses restruct to read the binary dc6 data into structs then parses image data from the frame data.
func LoadDC6(data []byte, palette d2datadict.PaletteRec) (DC6File, error) {
	result := DC6File{valid: true}

	restruct.EnableExprBeta()
	err := restruct.Unpack(data, binary.LittleEndian, &result)
	if err != nil {
		result.valid = false
		log.Printf("failed to read dc6: %v", err)
	}

	for _, frame := range result.Frames {
		frame.palette = palette
	}

	return result, err
}
