package MPQ

import (
	"encoding/binary"
	"errors"
	"log"
	"os"
	"path"
	"strings"
)

// MPQ represents an MPQ archive
type MPQ struct {
	File              *os.File
	HashTableEntries  []HashTableEntry
	BlockTableEntries []BlockTableEntry
	Data              Data
}

// Data Represents a MPQ file
type Data struct {
	Magic             [4]byte
	HeaderSize        uint32
	ArchiveSize       uint32
	FormatVersion     uint16
	BlockSize         uint16
	HashTableOffset   uint32
	BlockTableOffset  uint32
	HashTableEntries  uint32
	BlockTableEntries uint32
}

// HashTableEntry represents a hashed file entry in the MPQ file
type HashTableEntry struct { // 16 bytes
	NamePartA  uint32
	NamePartB  uint32
	Locale     uint16
	Platform   uint16
	BlockIndex uint32
}

// FileFlag represents flags for a file record in the MPQ archive
type FileFlag uint32

const (
	// MpqFileImplode - File is compressed using PKWARE Data compression library
	MpqFileImplode FileFlag = 0x00000100
	// MpqFileCompress - File is compressed using combination of compression methods
	MpqFileCompress FileFlag = 0x00000200
	// MpqFileEncrypted - The file is encrypted
	MpqFileEncrypted FileFlag = 0x00010000
	// MpqFileFixKey - The decryption key for the file is altered according to the position of the file in the archive
	MpqFileFixKey FileFlag = 0x00020000
	// MpqFilePatchFile - The file contains incremental patch for an existing file in base MPQ
	MpqFilePatchFile FileFlag = 0x00100000
	// MpqFileSingleUnit - Instead of being divided to 0x1000-bytes blocks, the file is stored as single unit
	MpqFileSingleUnit FileFlag = 0x01000000
	// FileDeleteMarker - File is a deletion marker, indicating that the file no longer exists. This is used to allow patch
	// archives to delete files present in lower-priority archives in the search chain. The file usually
	// has length of 0 or 1 byte and its name is a hash
	FileDeleteMarker FileFlag = 0x02000000
	// FileSEctorCrc - File has checksums for each sector. Ignored if file is not compressed or imploded.
	FileSEctorCrc FileFlag = 0x04000000
	// MpqFileExists - Set if file exists, reset when the file was deleted
	MpqFileExists FileFlag = 0x80000000
)

// BlockTableEntry represents an entry in the block table
type BlockTableEntry struct { // 16 bytes
	FilePosition         uint32
	CompressedFileSize   uint32
	UncompressedFileSize uint32
	Flags                FileFlag
	// Local Stuff...
	FileName       string
	EncryptionSeed uint32
}

// HasFlag returns true if the specified flag is present
func (v BlockTableEntry) HasFlag(flag FileFlag) bool {
	return (v.Flags & flag) != 0
}

// Load loads an MPQ file and returns a MPQ structure
func Load(fileName string) (MPQ, error) {
	result := MPQ{}
	file, err := os.Open(fileName)
	if err != nil {
		return MPQ{}, err
	}
	result.File = file
	err = result.readHeader()
	if err != nil {
		return MPQ{}, err
	}

	return result, nil
}

func (v *MPQ) readHeader() error {
	err := binary.Read(v.File, binary.LittleEndian, &v.Data)
	if err != nil {
		return err
	}
	if string(v.Data.Magic[:]) != "MPQ\x1A" {
		return errors.New("invalid mpq header")
	}
	v.loadHashTable()
	v.loadBlockTable()
	return nil
}

func (v *MPQ) loadHashTable() {
	v.File.Seek(int64(v.Data.HashTableOffset), 0)
	hashData := make([]uint32, v.Data.HashTableEntries*4)
	binary.Read(v.File, binary.LittleEndian, &hashData)
	decrypt(hashData, hashString("(hash table)", 3))
	for i := uint32(0); i < v.Data.HashTableEntries; i++ {
		v.HashTableEntries = append(v.HashTableEntries, HashTableEntry{
			NamePartA: hashData[i*4],
			NamePartB: hashData[(i*4)+1],
			// TODO: Verify that we're grabbing the right high/lo word for the vars below
			Locale:     uint16(hashData[(i*4)+2] >> 16),
			Platform:   uint16(hashData[(i*4)+2] & 0xFFFF),
			BlockIndex: hashData[(i*4)+3],
		})
	}
}

func (v *MPQ) loadBlockTable() {
	v.File.Seek(int64(v.Data.BlockTableOffset), 0)
	blockData := make([]uint32, v.Data.BlockTableEntries*4)
	binary.Read(v.File, binary.LittleEndian, &blockData)
	decrypt(blockData, hashString("(block table)", 3))
	for i := uint32(0); i < v.Data.BlockTableEntries; i++ {
		v.BlockTableEntries = append(v.BlockTableEntries, BlockTableEntry{
			FilePosition:         blockData[(i * 4)],
			CompressedFileSize:   blockData[(i*4)+1],
			UncompressedFileSize: blockData[(i*4)+2],
			Flags:                FileFlag(blockData[(i*4)+3]),
		})
	}
}

func decrypt(data []uint32, seed uint32) {
	seed2 := uint32(0xeeeeeeee)

	for i := 0; i < len(data); i++ {
		seed2 += CryptoBuffer[0x400+(seed&0xff)]
		result := data[i]
		result ^= seed + seed2

		seed = ((^seed << 21) + 0x11111111) | (seed >> 11)
		seed2 = result + seed2 + (seed2 << 5) + 3
		data[i] = result
	}
}

func decryptBytes(data []byte, seed uint32) {
	seed2 := uint32(0xEEEEEEEE)
	for i := 0; i < len(data)-3; i += 4 {
		seed2 += CryptoBuffer[0x400+(seed&0xFF)]
		result := binary.LittleEndian.Uint32(data[i : i+4])
		result ^= seed + seed2
		seed = ((^seed << 21) + 0x11111111) | (seed >> 11)
		seed2 = result + seed2 + (seed2 << 5) + 3

		data[i+0] = uint8(result & 0xff)
		data[i+1] = uint8((result >> 8) & 0xff)
		data[i+2] = uint8((result >> 16) & 0xff)
		data[i+3] = uint8((result >> 24) & 0xff)
	}
}

func hashString(key string, hashType uint32) uint32 {

	seed1 := uint32(0x7FED7FED)
	seed2 := uint32(0xEEEEEEEE)

	/* prepare seeds. */
	for _, char := range strings.ToUpper(key) {
		seed1 = CryptoBuffer[(hashType*0x100)+uint32(char)] ^ (seed1 + seed2)
		seed2 = uint32(char) + seed1 + seed2 + (seed2 << 5) + 3
	}
	return seed1
}

func (v MPQ) getFileHashEntry(fileName string) (HashTableEntry, error) {
	hashA := hashString(fileName, 1)
	hashB := hashString(fileName, 2)

	for idx, hashEntry := range v.HashTableEntries {
		if hashEntry.NamePartA != hashA || hashEntry.NamePartB != hashB {
			continue
		}

		return v.HashTableEntries[idx], nil
	}
	return HashTableEntry{}, errors.New("file not found")
}

// GetFileBlockData gets a block table entry
func (v MPQ) GetFileBlockData(fileName string) (BlockTableEntry, error) {
	fileEntry, err := v.getFileHashEntry(fileName)
	if err != nil {
		return BlockTableEntry{}, err
	}
	return v.BlockTableEntries[fileEntry.BlockIndex], nil
}

// Close closses the MPQ file
func (v *MPQ) Close() {
	v.File.Close()
}

// ReadFile reads a file from the MPQ and returns a memory stream
func (v MPQ) ReadFile(fileName string) ([]byte, error) {
	fileBlockData, err := v.GetFileBlockData(fileName)
	if err != nil {
		return nil, err
	}
	fileBlockData.FileName = strings.ToLower(fileName)
	fileBlockData.calculateEncryptionSeed()
	mpqStream := CreateStream(v, fileBlockData, fileName)
	buffer := make([]byte, fileBlockData.UncompressedFileSize)
	mpqStream.Read(buffer, 0, fileBlockData.UncompressedFileSize)
	return buffer, nil
}

// ReadTextFile reads a file and returns it as a string
func (v MPQ) ReadTextFile(fileName string) (string, error) {
	data, err := v.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (v *BlockTableEntry) calculateEncryptionSeed() {
	fileName := path.Base(v.FileName)
	v.EncryptionSeed = hashString(fileName, 3)
	if !v.HasFlag(MpqFileFixKey) {
		return
	}
	v.EncryptionSeed = (v.EncryptionSeed + v.FilePosition) ^ v.UncompressedFileSize
}

// GetFileList returns the list of files in this MPQ
func (v MPQ) GetFileList() ([]string, error) {
	data, err := v.ReadFile("(listfile)")
	if err != nil {
		return nil, err
	}
	log.Printf("File Contents:\n%s", strings.TrimRight(string(data), "\x00"))
	data = nil
	return []string{""}, nil
}
