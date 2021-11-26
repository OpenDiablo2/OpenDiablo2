package d2datautils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []byte{33, 23, 4, 33, 192, 243} //nolint:gochecknoglobals // just a test

func TestBitmuncherCopy(t *testing.T) {
	bm1 := CreateBitMuncher(testData, 0)
	bm2 := CopyBitMuncher(bm1)

	for i := range bm1.data {
		assert.Equal(t, bm1.data[i], bm2.data[i], "original bitmuncher isn't equal to copied")
	}
}

func TestBitmuncherSetOffset(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)
	bm.SetOffset(5)

	assert.Equal(t, bm.Offset(), 5, "Set Offset method didn't set offset to expected number")
}

func TestBitmuncherSteBitsRead(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)
	bm.SetBitsRead(8)

	assert.Equal(t, bm.BitsRead(), 8, "Set bits read method didn't set bits read to expected value")
}

func TestBitmuncherReadBit(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	var result byte

	for i := 0; i < byteLen; i++ {
		v := bm.GetBit()
		result |= byte(v) << byte(i)
	}

	assert.Equal(t, result, testData[0], "result of rpeated 8 times get bit didn't return expected byte")
}

func TestBitmuncherGetBits(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	assert.Equal(t, byte(bm.GetBits(byteLen)), testData[0], "get bits didn't return expected value")
}

func TestBitmuncherGetNoBits(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	assert.Equal(t, bm.GetBits(0), uint32(0), "get bits didn't return expected value: 0")
}

func TestBitmuncherGetSignedBits(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	assert.Equal(t, bm.GetSignedBits(6), -31, "get signed bits didn't return expected value")
}

func TestBitmuncherGetNoSignedBits(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	assert.Equal(t, bm.GetSignedBits(0), 0, "get signed bits didn't return expected value")
}

func TestBitmuncherGetOneSignedBit(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	assert.Equal(t, bm.GetSignedBits(1), -1, "get signed bits didn't return expected value")
}

func TestBitmuncherSkipBits(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	bm.SkipBits(byteLen)

	assert.Equal(t, bm.GetByte(), testData[1], "skipping 8 bits didn't moved bit muncher's position into next byte")
}

func TestBitmuncherGetInt32(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	var testInt int32

	for i := 0; i < bytesPerint32; i++ {
		testInt |= int32(testData[i]) << int32(byteLen*i)
	}

	assert.Equal(t, bm.GetInt32(), testInt, "int32 value wasn't returned properly")
}

func TestBitmuncherGetUint32(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	var testUint uint32

	for i := 0; i < bytesPerint32; i++ {
		testUint |= uint32(testData[i]) << uint32(byteLen*i)
	}

	assert.Equal(t, bm.GetUInt32(), testUint, "uint32 value wasn't returned properly")
}

func TestBitMuncherEnsureBits(t *testing.T) {
	bm := CreateBitMuncher(testData, 0)

	assert.Equal(t, true, bm.EnsureBits(byteLen*len(testData)), "unexpected value returned by EnsureBits")
	assert.Equal(t, false, bm.EnsureBits(byteLen*len(testData)+1), "unexpected value returned by EnsureBits")

	bm.SkipBits(5)

	assert.Equal(t, true, bm.EnsureBits(byteLen*len(testData)-5), "unexpected value returned by EnsureBits")
	assert.Equal(t, false, bm.EnsureBits(byteLen*len(testData)-4), "unexpected value returned by EnsureBits")
}
