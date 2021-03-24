package d2mpq

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

var _ d2interface.Archive = &MPQ{} // Static check to confirm struct conforms to interface

// MPQ represents an MPQ archive
type MPQ struct {
	filePath string
	file     *os.File
	hashes   map[uint64]*Hash
	blocks   []*Block
	header   Header
}

// PatchInfo represents patch info for the MPQ.
type PatchInfo struct {
	Length   uint32   // Length of patch info header, in bytes
	Flags    uint32   // Flags. 0x80000000 = MD5 (?)
	DataSize uint32   // Uncompressed size of the patch file
	MD5      [16]byte // MD5 of the entire patch file after decompression
}

// New loads an MPQ file and only reads the header
func New(fileName string) (*MPQ, error) {
	mpq := &MPQ{filePath: fileName}

	var err error
	if runtime.GOOS == "linux" {
		mpq.file, err = openIgnoreCase(fileName)
	} else {
		mpq.file, err = os.Open(fileName) //nolint:gosec // Will fix later
	}

	if err != nil {
		return nil, err
	}

	if err := mpq.readHeader(); err != nil {
		return nil, fmt.Errorf("failed to read reader: %v", err)
	}

	return mpq, nil
}

// FromFile loads an MPQ file and returns a MPQ structure
func FromFile(fileName string) (*MPQ, error) {
	mpq, err := New(fileName)
	if err != nil {
		return nil, err
	}

	if err := mpq.readHashTable(); err != nil {
		return nil, fmt.Errorf("failed to read hash table: %v", err)
	}

	if err := mpq.readBlockTable(); err != nil {
		return nil, fmt.Errorf("failed to read block table: %v", err)
	}

	return mpq, nil
}

// getFileBlockData gets a block table entry
func (mpq *MPQ) getFileBlockData(fileName string) (*Block, error) {
	fileEntry, ok := mpq.hashes[hashFilename(fileName)]
	if !ok {
		return nil, errors.New("file not found")
	}

	if fileEntry.BlockIndex >= uint32(len(mpq.blocks)) {
		return nil, errors.New("invalid block index")
	}

	return mpq.blocks[fileEntry.BlockIndex], nil
}

// Close closes the MPQ file
func (mpq *MPQ) Close() error {
	return mpq.file.Close()
}

// ReadFile reads a file from the MPQ and returns a memory stream
func (mpq *MPQ) ReadFile(fileName string) ([]byte, error) {
	fileBlockData, err := mpq.getFileBlockData(fileName)
	if err != nil {
		return []byte{}, err
	}

	fileBlockData.FileName = strings.ToLower(fileName)

	stream, err := CreateStream(mpq, fileBlockData, fileName)
	if err != nil {
		return []byte{}, err
	}

	buffer := make([]byte, fileBlockData.UncompressedFileSize)
	if _, err := stream.Read(buffer, 0, fileBlockData.UncompressedFileSize); err != nil {
		return []byte{}, err
	}

	return buffer, nil
}

// ReadFileStream reads the mpq file data and returns a stream
func (mpq *MPQ) ReadFileStream(fileName string) (d2interface.DataStream, error) {
	fileBlockData, err := mpq.getFileBlockData(fileName)
	if err != nil {
		return nil, err
	}

	fileBlockData.FileName = strings.ToLower(fileName)

	stream, err := CreateStream(mpq, fileBlockData, fileName)
	if err != nil {
		return nil, err
	}

	return &MpqDataStream{stream: stream}, nil
}

// ReadTextFile reads a file and returns it as a string
func (mpq *MPQ) ReadTextFile(fileName string) (string, error) {
	data, err := mpq.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Listfile returns the list of files in this MPQ
func (mpq *MPQ) Listfile() ([]string, error) {
	data, err := mpq.ReadFile("(listfile)")

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
func (mpq *MPQ) Path() string {
	return mpq.filePath
}

// Contains returns bool for whether the given filename exists in the mpq
func (mpq *MPQ) Contains(filename string) bool {
	_, ok := mpq.hashes[hashFilename(filename)]
	return ok
}

// Size returns the size of the mpq in bytes
func (mpq *MPQ) Size() uint32 {
	return mpq.header.ArchiveSize
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

	return os.Open(filepath.Join(mpqDir, mpqName)) //nolint:gosec // Will fix later
}
