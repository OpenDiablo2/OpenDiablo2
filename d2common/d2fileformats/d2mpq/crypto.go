package d2mpq

import (
	"encoding/binary"
	"io"
	"strings"
)

var cryptoBuffer [0x500]uint32 //nolint:gochecknoglobals // will fix later..
var cryptoBufferReady bool     //nolint:gochecknoglobals // will fix later..

func cryptoLookup(index uint32) uint32 {
	if !cryptoBufferReady {
		cryptoInitialize()

		cryptoBufferReady = true
	}

	return cryptoBuffer[index]
}

//nolint:gomnd // magic cryptographic stuff here...
func cryptoInitialize() {
	seed := uint32(0x00100001)

	for index1 := 0; index1 < 0x100; index1++ {
		index2 := index1

		for i := 0; i < 5; i++ {
			seed = (seed*125 + 3) % 0x2AAAAB
			temp1 := (seed & 0xFFFF) << 0x10
			seed = (seed*125 + 3) % 0x2AAAAB
			temp2 := seed & 0xFFFF
			cryptoBuffer[index2] = temp1 | temp2
			index2 += 0x100
		}
	}
}

func decrypt(data []uint32, seed uint32) {
	seed2 := uint32(0xeeeeeeee) //nolint:gomnd // Decryption magic

	for i := 0; i < len(data); i++ {
		seed2 += cryptoLookup(0x400 + (seed & 0xff)) //nolint:gomnd // Decryption magic
		result := data[i]
		result ^= seed + seed2

		seed = ((^seed << 21) + 0x11111111) | (seed >> 11)
		seed2 = result + seed2 + (seed2 << 5) + 3 //nolint:gomnd // Decryption magic
		data[i] = result
	}
}

func decryptBytes(data []byte, seed uint32) {
	seed2 := uint32(0xEEEEEEEE) //nolint:gomnd // Decryption magic
	for i := 0; i < len(data)-3; i += 4 {
		seed2 += cryptoLookup(0x400 + (seed & 0xFF)) //nolint:gomnd // Decryption magic
		result := binary.LittleEndian.Uint32(data[i : i+4])
		result ^= seed + seed2
		seed = ((^seed << 21) + 0x11111111) | (seed >> 11)
		seed2 = result + seed2 + (seed2 << 5) + 3 //nolint:gomnd // Decryption magic

		data[i+0] = uint8(result & 0xff)         //nolint:gomnd // Decryption magic
		data[i+1] = uint8((result >> 8) & 0xff)  //nolint:gomnd // Decryption magic
		data[i+2] = uint8((result >> 16) & 0xff) //nolint:gomnd // Decryption magic
		data[i+3] = uint8((result >> 24) & 0xff) //nolint:gomnd // Decryption magic
	}
}

func decryptTable(r io.Reader, size uint32, name string) ([]uint32, error) {
	seed := hashString(name, 3)
	seed2 := uint32(0xEEEEEEEE) //nolint:gomnd // Decryption magic
	size *= 4

	table := make([]uint32, size)
	buf := make([]byte, 4)

	for i := uint32(0); i < size; i++ {
		seed2 += cryptoBuffer[0x400+(seed&0xff)] //nolint:gomnd // Decryption magic

		if _, err := r.Read(buf); err != nil {
			return table, err
		}

		result := binary.LittleEndian.Uint32(buf)
		result ^= seed + seed2

		seed = ((^seed << 21) + 0x11111111) | (seed >> 11)
		seed2 = result + seed2 + (seed2 << 5) + 3 //nolint:gomnd // Decryption magic
		table[i] = result
	}

	return table, nil
}

func hashFilename(key string) uint64 {
	a, b := hashString(key, 1), hashString(key, 2)
	return uint64(a)<<32 | uint64(b)
}

func hashString(key string, hashType uint32) uint32 {
	seed1 := uint32(0x7FED7FED) //nolint:gomnd // Decryption magic
	seed2 := uint32(0xEEEEEEEE) //nolint:gomnd // Decryption magic

	/* prepare seeds. */
	for _, char := range strings.ToUpper(key) {
		seed1 = cryptoLookup((hashType*0x100)+uint32(char)) ^ (seed1 + seed2)
		seed2 = uint32(char) + seed1 + seed2 + (seed2 << 5) + 3 //nolint:gomnd // Decryption magic
	}

	return seed1
}

//nolint:unused,deadcode // will use this for creating mpq's
func encrypt(data []uint32, seed uint32) {
	seed2 := uint32(0xeeeeeeee) //nolint:gomnd // Decryption magic

	for i := 0; i < len(data); i++ {
		seed2 += cryptoLookup(0x400 + (seed & 0xff)) //nolint:gomnd // Decryption magic
		result := data[i]
		result ^= seed + seed2

		seed = ((^seed << 21) + 0x11111111) | (seed >> 11)
		seed2 = data[i] + seed2 + (seed2 << 5) + 3 //nolint:gomnd // Decryption magic
		data[i] = result
	}
}
