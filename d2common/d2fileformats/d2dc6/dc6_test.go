package d2dc6

import (
	"testing"
)

func TestDC6New(t *testing.T) {
	dc6 := New()

	if dc6 == nil {
		t.Error("d2dc6.New() method returned nil")
	}
}

func getExampleDC6() *DC6 {
	exampleDC6 := &DC6{
		Version:            6,
		Flags:              1,
		Encoding:           0,
		Termination:        []byte{238, 238, 238, 238},
		Directions:         1,
		FramesPerDirection: 1,
		FramePointers:      []uint32{56},
		Frames: []*DC6Frame{
			{
				Flipped:    0,
				Width:      32,
				Height:     26,
				OffsetX:    45,
				OffsetY:    24,
				Unknown:    0,
				NextBlock:  50,
				Length:     10,
				FrameData:  []byte{2, 23, 34, 128, 53, 64, 39, 43, 123, 12},
				Terminator: []byte{2, 8, 5},
			},
		},
	}

	return exampleDC6
}

func TestDC6Unmarshal(t *testing.T) {
	exampleDC6 := getExampleDC6()

	data := exampleDC6.Marshal()

	extractedDC6, err := Load(data)
	if err != nil {
		t.Error(err)
	}

	if exampleDC6.Version != extractedDC6.Version ||
		len(exampleDC6.Frames) != len(extractedDC6.Frames) ||
		exampleDC6.Frames[0].NextBlock != extractedDC6.Frames[0].NextBlock {
		t.Fatal("encoded and decoded DC6 isn't the same")
	}
}

func TestDC6Clone(t *testing.T) {
	exampleDC6 := getExampleDC6()
	clonedDC6 := exampleDC6.Clone()

	if exampleDC6.Termination[0] != clonedDC6.Termination[0] ||
		len(exampleDC6.Frames) != len(clonedDC6.Frames) ||
		exampleDC6.Frames[0].NextBlock != clonedDC6.Frames[0].NextBlock {
		t.Fatal("cloned dc6 isn't equal to original")
	}
}
