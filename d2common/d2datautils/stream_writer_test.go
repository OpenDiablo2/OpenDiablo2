package d2datautils

import (
	"testing"
)

func TestStreamWriterByte(t *testing.T) {
	sr := CreateStreamWriter()
	data := []byte{0x12, 0x34, 0x56, 0x78}

	for _, d := range data {
		sr.PushByte(d)
	}

	output := sr.GetBytes()
	for i, d := range data {
		if output[i] != d {
			t.Fatalf("sr.PushByte() pushed %X, but wrote %X instead", d, output[i])
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
