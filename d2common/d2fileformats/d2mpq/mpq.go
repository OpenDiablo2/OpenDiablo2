package d2mpq

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var _ d2interface.Archive = &MPQ{} // Static check to confirm struct conforms to interface

// MPQ represents an MPQ archive
type MPQ struct {
	filePath          string
	file              *os.File
	hashEntryMap      HashEntryMap
	blockTableEntries []BlockTableEntry
	data              Data
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

// PatchInfo represents patch info for the MPQ.
type PatchInfo struct {
	Length   uint32   // Length of patch info header, in bytes
	Flags    uint32   // Flags. 0x80000000 = MD5 (?)
	DataSize uint32   // Uncompressed size of the patch file
	Md5      [16]byte // MD5 of the entire patch file after decompression
}

// FileFlag represents flags for a file record in the MPQ archive
type FileFlag uint32

const (
	// FileImplode - File is compressed using PKWARE Data compression library
	FileImplode FileFlag = 0x00000100
	// FileCompress - File is compressed using combination of compression methods
	FileCompress FileFlag = 0x00000200
	// FileEncrypted - The file is encrypted
	FileEncrypted FileFlag = 0x00010000
	// FileFixKey - The decryption key for the file is altered according to the position of the file in the archive
	FileFixKey FileFlag = 0x00020000
	// FilePatchFile - The file contains incremental patch for an existing file in base MPQ
	FilePatchFile FileFlag = 0x00100000
	// FileSingleUnit - Instead of being divided to 0x1000-bytes blocks, the file is stored as single unit
	FileSingleUnit FileFlag = 0x01000000
	// FileDeleteMarker - File is a deletion marker, indicating that the file no longer exists. This is used to allow patch
	// archives to delete files present in lower-priority archives in the search chain. The file usually
	// has length of 0 or 1 byte and its name is a hash
	FileDeleteMarker FileFlag = 0x02000000
	// FileSectorCrc - File has checksums for each sector. Ignored if file is not compressed or imploded.
	FileSectorCrc FileFlag = 0x04000000
	// FileExists - Set if file exists, reset when the file was deleted
	FileExists FileFlag = 0x80000000
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
func Load(fileName string) (d2interface.Archive, error) {
	result := &MPQ{filePath: fileName}

	var err error
	if runtime.GOOS == "linux" {
		result.file, err = openIgnoreCase(fileName)
	} else {
		result.file, err = os.Open(fileName) //nolint:gosec // Will fix later
	}

	if err != nil {
		return nil, err
	}

	if err := result.readHeader(); err != nil {
		return nil, err
	}

	return result, nil
}

func openIgnoreCase(mpqPath string) (*os.File, error) {
	// First see if file exists with specified case
	mpqFile, err := os.Open(mpqPath) //nolint:gosec // Will fix later
	if err == nil {
		return mpqFile, err
	}

	mpqName := filepath.Base(mpqPath)
	mpqDir := filepath.Dir(mpqPath)

	files, err := ioutil.ReadDir(mpqDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.EqualFold(file.Name(), mpqName) {
			mpqName = file.Name()
			break
		}
	}

	file, err := os.Open(path.Join(mpqDir, mpqName)) //nolint:gosec // Will fix later

	return file, err
}

func (v *MPQ) readHeader() error {
	err := binary.Read(v.file, binary.LittleEndian, &v.data)

	if err != nil {
		return err
	}

	if string(v.data.Magic[:]) != "MPQ\x1A" {
		return errors.New("invalid mpq header")
	}

	err = v.loadHashTable()
	if err != nil {
		return err
	}

	v.loadBlockTable()

	return nil
}

func (v *MPQ) loadHashTable() error {
	_, err := v.file.Seek(int64(v.data.HashTableOffset), 0)
	if err != nil {
		log.Panic(err)
	}

	hashData := make([]uint32, v.data.HashTableEntries*4) //nolint:gomnd // // Decryption magic
	hash := make([]byte, 4)

	for i := range hashData {
		_, err := v.file.Read(hash)
		if err != nil {
			log.Print(err)
		}

		hashData[i] = binary.LittleEndian.Uint32(hash)
	}

	decrypt(hashData, hashString("(hash table)", 3))

	for i := uint32(0); i < v.data.HashTableEntries; i++ {
		v.hashEntryMap.Insert(&HashTableEntry{
			NamePartA: hashData[i*4],
			NamePartB: hashData[(i*4)+1],
			// https://github.com/OpenDiablo2/OpenDiablo2/issues/812
			Locale:     uint16(hashData[(i*4)+2] >> 16),    //nolint:gomnd // // binary data
			Platform:   uint16(hashData[(i*4)+2] & 0xFFFF), //nolint:gomnd // // binary data
			BlockIndex: hashData[(i*4)+3],
		})
	}

	return nil
}

func (v *MPQ) loadBlockTable() {
	_, err := v.file.Seek(int64(v.data.BlockTableOffset), 0)
	if err != nil {
		log.Panic(err)
	}

	blockData := make([]uint32, v.data.BlockTableEntries*4) //nolint:gomnd // // binary data
	hash := make([]byte, 4)

	for i := range blockData {
		_, err = v.file.Read(hash) //nolint:errcheck // Will fix later
		if err != nil {
			log.Print(err)
		}

		blockData[i] = binary.LittleEndian.Uint32(hash)
	}

	decrypt(blockData, hashString("(block table)", 3))

	for i := uint32(0); i < v.data.BlockTableEntries; i++ {
		v.blockTableEntries = append(v.blockTableEntries, BlockTableEntry{
			FilePosition:         blockData[(i * 4)],
			CompressedFileSize:   blockData[(i*4)+1],
			UncompressedFileSize: blockData[(i*4)+2],
			Flags:                FileFlag(blockData[(i*4)+3]),
		})
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

// GetFileBlockData gets a block table entry
func (v *MPQ) getFileBlockData(fileName string) (BlockTableEntry, error) {
	fileEntry, found := v.hashEntryMap.Find(fileName)

	if !found || fileEntry.BlockIndex >= uint32(len(v.blockTableEntries)) {
		return BlockTableEntry{}, errors.New("file not found")
	}

	return v.blockTableEntries[fileEntry.BlockIndex], nil
}

// Close closes the MPQ file
func (v *MPQ) Close() {
	err := v.file.Close()
	if err != nil {
		log.Panic(err)
	}
}

// FileExists checks the mpq to see if the file exists
func (v *MPQ) FileExists(fileName string) bool {
	return v.hashEntryMap.Contains(fileName)
}

// ReadFile reads a file from the MPQ and returns a memory stream
func (v *MPQ) ReadFile(fileName string) ([]byte, error) {
	fileBlockData, err := v.getFileBlockData(fileName)
	if err != nil {
		return []byte{}, err
	}

	fileBlockData.FileName = strings.ToLower(fileName)

	fileBlockData.calculateEncryptionSeed()
	mpqStream, err := CreateStream(v, fileBlockData, fileName)

	if err != nil {
		return []byte{}, err
	}

	buffer := make([]byte, fileBlockData.UncompressedFileSize)
	mpqStream.Read(buffer, 0, fileBlockData.UncompressedFileSize)

	return buffer, nil
}

// ReadFileStream reads the mpq file data and returns a stream
func (v *MPQ) ReadFileStream(fileName string) (d2interface.DataStream, error) {
	fileBlockData, err := v.getFileBlockData(fileName)

	if err != nil {
		return nil, err
	}

	fileBlockData.FileName = strings.ToLower(fileName)
	fileBlockData.calculateEncryptionSeed()

	mpqStream, err := CreateStream(v, fileBlockData, fileName)
	if err != nil {
		return nil, err
	}

	return &MpqDataStream{stream: mpqStream}, nil
}

// ReadTextFile reads a file and returns it as a string
func (v *MPQ) ReadTextFile(fileName string) (string, error) {
	data, err := v.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (v *BlockTableEntry) calculateEncryptionSeed() {
	fileName := path.Base(v.FileName)
	v.EncryptionSeed = hashString(fileName, 3)

	if !v.HasFlag(FileFixKey) {
		return
	}

	v.EncryptionSeed = (v.EncryptionSeed + v.FilePosition) ^ v.UncompressedFileSize
}

// GetFileList returns the list of files in this MPQ
func (v *MPQ) GetFileList() ([]string, error) {
	data, err := v.ReadFile("(listfile)")

	if err != nil {
		return nil, err
	}

	raw := strings.TrimRight(string(data), "\x00")
	s := bufio.NewScanner(strings.NewReader(raw))

	var filePaths []string

	for s.Scan() {
		filePath := s.Text()
		filePaths = append(filePaths, filePath)
	}

	return filePaths, nil
}

// Path returns the MPQ file path
func (v *MPQ) Path() string {
	return v.filePath
}

// Contains returns bool for whether the given filename exists in the mpq
func (v *MPQ) Contains(filename string) bool {
	return v.hashEntryMap.Contains(filename)
}

// Size returns the size of the mpq in bytes
func (v *MPQ) Size() uint32 {
	return v.data.ArchiveSize
}
