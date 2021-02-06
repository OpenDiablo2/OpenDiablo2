package d2datautils

import (
	"testing"
)

func TestStreamWriterBits(t *testing.T) {
	sr := CreateStreamWriter()
	data := []byte{221, 19}

	for _, i := range data {
		sr.PushBits(i, bitsPerByte)
	}

	output := sr.GetBytes()
	for i, d := range data {
		if output[i] != d {
			t.Fatalf("sr.PushBits() pushed %X, but wrote %X instead", d, output[i])
		}
	}
}

func TestStreamWriterBits16(t *testing.T) {
	sr := CreateStreamWriter()
	data := []uint16{1024, 19}

	for _, i := range data {
		sr.PushBits16(i, bitsPerByte*bytesPerint16)
	}

	output := sr.GetBytes()

	for i, d := range data {
		outputInt := uint16(output[bytesPerint16*i]) | uint16(output[bytesPerint16*i+1])<<8
		if outputInt != d {
			t.Fatalf("sr.PushBits16() pushed %X, but wrote %X instead", d, output[i])
		}
	}
}

func TestStreamWriterBits32(t *testing.T) {
	sr := CreateStreamWriter()
	data := []uint32{19324, 87}

	for _, i := range data {
		sr.PushBits32(i, bitsPerByte*bytesPerint32)
	}

	output := sr.GetBytes()

	for i, d := range data {
		outputInt := uint32(output[bytesPerint32*i]) |
			uint32(output[bytesPerint32*i+1])<<8 |
			uint32(output[bytesPerint32*i+2])<<16 |
			uint32(output[bytesPerint32*i+3])<<24

		if outputInt != d {
			t.Fatalf("sr.PushBits32() pushed %X, but wrote %X instead", d, output[i])
		}
	}
}

func TestStreamWriterByte(t *testing.T) {
	sr := CreateStreamWriter()
	data := []byte{0x12, 0x34, 0x56, 0x78}

	sr.PushBytes(data...)

	output := sr.GetBytes()
	for i, d := range data {
		if output[i] != d {
			t.Fatalf("sr.PushBytes() pushed %X, but wrote %X instead", d, output[i])
		}
	}
}

func TestStreamWriterWord(t *testing.T) {
	sr := CreateStreamWriter()
	data := []byte{0x12, 0x34, 0x56, 0x78}

	sr.PushUint16(0x3412)
	sr.PushUint16(0x7856)

	output := sr.GetBytes()
	for i, d := range data {
		if output[i] != d {
			t.Fatalf("sr.PushWord() pushed byte %X to %d, but %X was expected instead", output[i], i, d)
		}
	}
}

func TestStreamWriterDword(t *testing.T) {
	sr := CreateStreamWriter()
	data := []byte{0x12, 0x34, 0x56, 0x78}

	sr.PushUint32(0x78563412)

	output := sr.GetBytes()
	for i, d := range data {
		if output[i] != d {
			t.Fatalf("sr.PushDword() pushed byte %X to %d, but %X was expected instead", output[i], i, d)
		}
	}
}
