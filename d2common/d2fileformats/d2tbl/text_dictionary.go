package d2tbl

import (
	"errors"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
)

// TextDictionary is a string map
type TextDictionary map[string]string

func (td TextDictionary) loadHashEntries(hashEntries []*textDictionaryHashEntry, br *d2datautils.StreamReader) error {
	for i := 0; i < len(hashEntries); i++ {
		entry := textDictionaryHashEntry{}

		active, err := br.ReadByte()
		if err != nil {
			return err
		}

		entry.IsActive = active > 0

		entry.Index, err = br.ReadUInt16()
		if err != nil {
			return err
		}

		entry.HashValue, err = br.ReadUInt32()
		if err != nil {
			return err
		}

		entry.IndexString, err = br.ReadUInt32()
		if err != nil {
			return err
		}

		entry.NameString, err = br.ReadUInt32()
		if err != nil {
			return err
		}

		entry.NameLength, err = br.ReadUInt16()
		if err != nil {
			return err
		}

		hashEntries[i] = &entry
	}

	for idx := range hashEntries {
		if !hashEntries[idx].IsActive {
			continue
		}

		if err := td.loadHashEntry(idx, hashEntries[idx], br); err != nil {
			return err
		}
	}

	return nil
}

func (td TextDictionary) loadHashEntry(idx int, hashEntry *textDictionaryHashEntry, br *d2datautils.StreamReader) error {
	br.SetPosition(uint64(hashEntry.NameString))

	nameVal, err := br.ReadBytes(int(hashEntry.NameLength - 1))
	if err != nil {
		return err
	}

	value := string(nameVal)

	br.SetPosition(uint64(hashEntry.IndexString))

	key := ""

	for {
		b, err := br.ReadByte()
		if b == 0 {
			break
		}

		if err != nil {
			return err
		}

		key += string(b)
	}

	if key == "x" || key == "X" {
		key = "#" + strconv.Itoa(idx)
	}

	_, exists := td[key]
	if !exists {
		td[key] = value
	}

	return nil
}

type textDictionaryHashEntry struct {
	IsActive    bool
	Index       uint16
	HashValue   uint32
	IndexString uint32
	NameString  uint32
	NameLength  uint16
}

const (
	crcByteCount = 2
)

// LoadTextDictionary loads the text dictionary from the given data
func LoadTextDictionary(dictionaryData []byte) (TextDictionary, error) {
	lookupTable := make(TextDictionary)

	br := d2datautils.CreateStreamReader(dictionaryData)

	// skip past the CRC
	_, _ = br.ReadBytes(crcByteCount)

	var err error

	numberOfElements, err := br.ReadUInt16()
	if err != nil {
		return nil, err
	}

	hashTableSize, err := br.ReadUInt32()
	if err != nil {
		return nil, err
	}

	// Version (always 0)
	if _, err = br.ReadByte(); err != nil {
		return nil, errors.New("error reading Version record")
	}

	_, _ = br.ReadUInt32() // StringOffset

	// When the number of times you have missed a match with a
	// hash key equals this value, you give up because it is not there.
	_, _ = br.ReadUInt32()

	_, _ = br.ReadUInt32() // FileSize

	elementIndex := make([]uint16, numberOfElements)
	for i := 0; i < int(numberOfElements); i++ {
		elementIndex[i], err = br.ReadUInt16()
		if err != nil {
			return nil, err
		}
	}

	hashEntries := make([]*textDictionaryHashEntry, hashTableSize)

	err = lookupTable.loadHashEntries(hashEntries, br)
	if err != nil {
		return nil, err
	}

	return lookupTable, nil
}
