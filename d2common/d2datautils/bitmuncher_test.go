package d2datautils

import "testing"

var testData = []byte{33, 23, 4, 33, 192, 243}

func TestBitmuncherCopy(t *testing.T) {
	bm1 := CreateBitMuncher(testData, 0)
	bm2 := CopyBitMuncher(bm1)

	for i := range bm1.data {
		if bm1.data[i] != bm2.data[i] {
			t.Fatal("original bitmuncher isn't equal to copied")
		}
	}
}

func TestBitmuncherSetOffset(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)
	bm.SetOffset(5)

	if bm.Offset() != 5 {
		t.Fatal("Set Offset method didn't set offset to expected number")
	}
}

func TestBitmuncherSteBitsRead(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)
	bm.SetBitsRead(8)

	if bm.BitsRead() != 8 {
		t.Fatal("Set bits read method didn't set bits read to expected value")
	}
}

func TestBitmuncherReadBit(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	var result byte

	for i := 0; i < bitsPerByte; i++ {
		v := bm.GetBit()
		result |= byte(v) << byte(i)
	}

	if result != testData[0] {
		t.Fatal("result of rpeated 8 times get bit didn't return expected byte")
	}
}

func TestBitmuncherGetBits(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	if byte(bm.GetBits(bitsPerByte)) != testData[0] {
		t.Fatal("get bits didn't return expected value")
	}
}

func TestBitmuncherGetNoBits(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	if bm.GetBits(0) != 0 {
		t.Fatal("get bits didn't return expected value: 0")
	}
}

func TestBitmuncherGetSignedBits(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	result := bm.GetSignedBits(6)
	expected := -31
	if result != expected {
		t.Fatal("get signed bits didn't return expected value", result, expected)
	}
}

func TestBitmuncherGetNoSignedBits(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	result := bm.GetSignedBits(0)
	expected := 0
	if result != expected {
		t.Fatal("get signed bits didn't return expected value", result, expected)
	}
}

func TestBitmuncherGetOneSignedBit(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	result := bm.GetSignedBits(1)
	expected := -1
	if result != expected {
		t.Fatal("get signed bits didn't return expected value", result, expected)
	}
}

func TestBitmuncherSkipBits(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	bm.SkipBits(bitsPerByte)

	if bm.GetByte() != testData[1] {
		t.Fatal("skipping 8 bits didn't moved bit muncher's position into next byte")
	}
}

func TestBitmuncherGetInt32(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	var testInt int32

	for i := 0; i < bytesPerint32; i++ {
		testInt |= int32(testData[i]) << int32(bitsPerByte*i)
	}

	if bm.GetInt32() != testInt {
		t.Fatal("int32 value wasn't returned properly")
	}
}

func TestBitmuncherGetUint32(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	var testUint uint32

	for i := 0; i < bytesPerint32; i++ {
		testUint |= uint32(testData[i]) << uint32(bitsPerByte*i)
	}

	if bm.GetUInt32() != testUint {
		t.Fatal("uint32 value wasn't returned properly")
	}
}
