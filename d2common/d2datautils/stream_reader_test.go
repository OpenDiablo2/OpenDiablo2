package d2datautils

import (
	"testing"
)

func TestStreamReaderByte(t *testing.T) {
	data := []byte{0x78, 0x56, 0x34, 0x12}
	sr := CreateStreamReader(data)

	if sr.GetPosition() != 0 {
		t.Fatal("StreamReader.GetPosition() did not start at 0")
	}

	if ss := sr.GetSize(); ss != 4 {
		t.Fatalf("StreamREader.GetSize() was expected to return %d, but returned %d instead", 4, ss)
	}

	for i := 0; i < len(data); i++ {
		ret := sr.GetByte()
		if ret != data[i] {
			t.Fatalf("StreamReader.GetDword() was expected to return %X, but returned %X instead", data[i], ret)
		}

		if pos := sr.GetPosition(); pos != uint64(i+1) {
			t.Fatalf("StreamReader.GetPosition() should be at %d, but was at %d instead", i, pos)
		}
	}
}

func TestStreamReaderWord(t *testing.T) {
	data := []byte{0x78, 0x56, 0x34, 0x12}
	sr := CreateStreamReader(data)
	ret := sr.GetUInt16()

	if ret != 0x5678 {
		t.Fatalf("StreamReader.GetDword() was expected to return %X, but returned %X instead", 0x5678, ret)
	}

	if pos := sr.GetPosition(); pos != 2 {
		t.Fatalf("StreamReader.GetPosition() should be at %d, but was at %d instead", 2, pos)
	}

	ret = sr.GetUInt16()
	if ret != 0x1234 {
		t.Fatalf("StreamReader.GetDword() was expected to return %X, but returned %X instead", 0x1234, ret)
	}

	if pos := sr.GetPosition(); pos != 4 {
		t.Fatalf("StreamReader.GetPosition() should be at %d, but was at %d instead", 4, pos)
	}
}

func TestStreamReaderDword(t *testing.T) {
	data := []byte{0x78, 0x56, 0x34, 0x12}
	sr := CreateStreamReader(data)
	ret := sr.GetUInt32()

	if ret != 0x12345678 {
		t.Fatalf("StreamReader.GetDword() was expected to return %X, but returned %X instead", 0x12345678, ret)
	}

	if pos := sr.GetPosition(); pos != 4 {
		t.Fatalf("StreamReader.GetPosition() should be at %d, but was at %d instead", 4, pos)
	}
}
