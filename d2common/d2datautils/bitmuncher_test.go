package d2datautils

import (
	"testing"
)

func getTestData() []byte {
	result := []byte{33, 23, 4, 33, 192, 243}

	return result
}

func TestBitmuncherCopy(t *testing.T) {
	bm1 := CreateBitMuncher(getTestData(), 0)
	bm2 := CopyBitMuncher(bm1)

	for i := range bm1.data {
		if bm1.data[i] != bm2.data[i] {
			t.Fatal("original bitmuncher isn't equal to copied")
		}
	}
}

func TestBitmuncherSetOffset(t *testing.T) {
	bm := CreateBitMuncher(getTestData(), 0)
	bm.SetOffset(5)

	if bm.Offset() != 5 {
		t.Fatal("Set Offset method didn't set offset to expected number")
	}
}

func TestBitmuncherSteBitsRead(t *testing.T) {
	bm := CreateBitMuncher(getTestData(), 0)
	bm.SetBitsRead(8)

	if bm.BitsRead() != 8 {
		t.Fatal("Set bits read method didn't set bits read to expected value")
	}
}

func TestBitmuncherReadBit(t *testing.T) {
	td := getTestData()
	bm := CreateBitMuncher(td, 0)
	var result byte

	for i := 0; i < bitsPerByte; i++ {
		v := bm.GetBit()
		result |= byte(v) << byte(i)
	}

	if result != td[0] {
		t.Fatal("result of rpeated 8 times get bit didn't return expected byte")
	}
}

func TestBitmuncherGetBits(t *testing.T) {
	td := getTestData()
	bm := CreateBitMuncher(td, 0)

	res := bm.GetBits(bitsPerByte)

	if byte(res) != td[0] {
		t.Fatal("get bits didn't return expected value")
	}
}

func TestBitmuncherSkipBits(t *testing.T) {
	td := getTestData()
	bm := CreateBitMuncher(td, 0)

	bm.SkipBits(bitsPerByte)

	if bm.GetByte() != td[1] {
		t.Fatal("skipping 8 bits didn't moved bit muncher's position into next byte")
	}
}

func TestBitmuncherGetInt32(t *testing.T) {
	td := getTestData()
	bm := CreateBitMuncher(td, 0)

	var testInt int32

	for i := 0; i < bytesPerint32; i++ {
		testInt |= int32(td[i]) << int32(bitsPerByte*i)
	}

	bmInt := bm.GetInt32()

	if bmInt != testInt {
		t.Fatal("int32 value wasn't returned properly")
	}
}
