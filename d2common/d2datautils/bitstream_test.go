package d2datautils

import (
	"testing"
)

func TestBitStreamBits(t *testing.T) {
	data := []byte{0xAA}
	bitStream := CreateBitStream(data)
	shouldBeOne := 0

	for i := 0; i < 8; i++ {
		bit := bitStream.ReadBits(1)
		if bit != shouldBeOne {
			t.Fatalf("Expected %d but got %d on iteration %d", shouldBeOne, bit, i)
		}

		if shouldBeOne == 1 {
			shouldBeOne = 0
		} else {
			shouldBeOne = 1
		}
	}
}

func TestBitStreamBytes(t *testing.T) {
	data := []byte{0xAA, 0xBB, 0xCC, 0xDD, 0x12, 0x34, 0x56, 0x78}
	bitStream := CreateBitStream(data)

	for i := 0; i < 8; i++ {
		b := byte(bitStream.ReadBits(8))
		if b != data[i] {
			t.Fatalf("Expected %d but got %d on iteration %d", data[i], b, i)
		}
	}
}
