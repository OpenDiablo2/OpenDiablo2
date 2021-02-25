package d2tbl

import (
	"fmt"

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
			return fmt.Errorf("reading active: %v", err)
		}

		entry.IsActive = active > 0

		entry.Index, err = br.ReadUInt16()
		if err != nil {
			return fmt.Errorf("reading Index: %v", err)
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
			return fmt.Errorf("loading entry %d: %v", idx, err)
		}
	}

	return nil
}

func (td TextDictionary) loadHashEntry(idx int, hashEntry *textDictionaryHashEntry, br *d2datautils.StreamReader) error {
	br.SetPosition(uint64(hashEntry.NameString))

	nameVal, err := br.ReadBytes(int(hashEntry.NameLength - 1))
	if err != nil {
		return fmt.Errorf("reading name value: %v", err)
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

	/*
		number of indicates
		(https://d2mods.info/forum/viewtopic.php?p=202077#p202077)
		Indices ...
		An array of WORD. Each entry is an index into the hash table.
		The actual string key index in the .bin file is an index into this table.
		So to get a string from a key index ...
	*/
	numberOfElements, err := br.ReadUInt16()
	if err != nil {
		return nil, fmt.Errorf("reading number of elements: %v", err)
	}

	hashTableSize, err := br.ReadUInt32()
	if err != nil {
		return nil, fmt.Errorf("reading hash table size: %v", err)
	}

	// Version (always 0)
	version, err := br.ReadByte()
	if err != nil {
		return nil, fmt.Errorf("reading version: %v", err)
	}

	if version != 0 {
		return nil, fmt.Errorf("version isn't equal to 0")
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
			return nil, fmt.Errorf("reading element index %d: %v", i, err)
		}
	}

	hashEntries := make([]*textDictionaryHashEntry, hashTableSize)

	err = lookupTable.loadHashEntries(hashEntries, br)
	if err != nil {
		return nil, fmt.Errorf("loading has entries: %v", err)
	}

	return lookupTable, nil
}

// Marshal encodes text dictionary back into byte slice
func (td *TextDictionary) Marshal() []byte {
	sw := d2datautils.CreateStreamWriter()

	// https://github.com/OpenDiablo2/OpenDiablo2/issues/1043
	sw.PushBytes(0, 0)

	sw.PushUint16(0)

	sw.PushInt32(int32(len(*td)))

	// version (always 0)
	sw.PushBytes(0)

	// offset of start of data (unnecessary for our decoder)
	sw.PushUint32(0)

	// Max retry count for a hash hit.
	sw.PushUint32(0)

	// offset to end of data (noop)
	sw.PushUint32(0)

	// indicates (len = 0, so nothing here)

	// nolint:gomnd // 17 comes from the size of one "data-header index"
	// dataPos is a position, when we're placing data stream
	dataPos := len(sw.GetBytes()) + 17*len(*td)

	for key, value := range *td {
		// non-zero if record is used (for us, every record is used ;-)
		sw.PushBytes(1)

		// generally unused;
		// 	string key index (used in .bin)
		sw.PushUint16(0)

		// also unused in our decoder
		// 	calculated hash of the string.
		sw.PushUint32(0)

		sw.PushUint32(uint32(dataPos))
		dataPos += len(key) + 1

		sw.PushUint32(uint32(dataPos))
		dataPos += len(value) + 1

		sw.PushUint16(uint16(len(value) + 1))
	}

	// data stream: put all data in appropiate order
	for key, value := range *td {
		for _, i := range key {
			sw.PushBytes(byte(i))
		}

		// 0 as separator
		sw.PushBytes(0)

		for _, i := range value {
			sw.PushBytes(byte(i))
		}

		// 0 as separator
		sw.PushBytes(0)
	}

	return sw.GetBytes()
}
