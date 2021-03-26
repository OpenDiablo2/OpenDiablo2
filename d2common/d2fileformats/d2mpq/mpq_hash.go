package d2mpq

import "io"

// Hash represents a hashed file entry in the MPQ file
type Hash struct { // 16 bytes
	A          uint32
	B          uint32
	Locale     uint16
	Platform   uint16
	BlockIndex uint32
}

// Name64 returns part A and B as uint64
func (h *Hash) Name64() uint64 {
	return uint64(h.A)<<32 | uint64(h.B)
}

//nolint:gomnd // number
func (mpq *MPQ) readHashTable() error {
	if _, err := mpq.file.Seek(int64(mpq.header.HashTableOffset), io.SeekStart); err != nil {
		return err
	}

	hashData, err := decryptTable(mpq.file, mpq.header.HashTableEntries, "(hash table)")
	if err != nil {
		return err
	}

	mpq.hashes = make(map[uint64]*Hash)

	for n, i := uint32(0), uint32(0); i < mpq.header.HashTableEntries; n, i = n+4, i+1 {
		e := &Hash{
			A: hashData[n],
			B: hashData[n+1],
			// https://github.com/OpenDiablo2/OpenDiablo2/issues/812
			Locale:     uint16(hashData[n+2] >> 16),    //nolint:gomnd // // binary data
			Platform:   uint16(hashData[n+2] & 0xFFFF), //nolint:gomnd // // binary data
			BlockIndex: hashData[n+3],
		}
		mpq.hashes[e.Name64()] = e
	}

	return nil
}
