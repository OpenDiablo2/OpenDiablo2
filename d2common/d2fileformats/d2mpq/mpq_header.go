package d2mpq

import (
	"encoding/binary"
	"errors"
	"io"
)

// Header Represents a MPQ file
type Header struct {
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

func (mpq *MPQ) readHeader() error {
	if _, err := mpq.file.Seek(0, io.SeekStart); err != nil {
		return err
	}

	if err := binary.Read(mpq.file, binary.LittleEndian, &mpq.header); err != nil {
		return err
	}

	if string(mpq.header.Magic[:]) != "MPQ\x1A" {
		return errors.New("invalid mpq header")
	}

	return nil
}
