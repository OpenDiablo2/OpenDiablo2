package d2tbl

import (
	"errors"
	"sort"
	"strconv"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
)

// TextDictionary is a string map
type TextDictionary struct {
	crcBytes      []byte
	elementIndex  []uint16
	hashTableSize uint32
	version       byte
	stringOffset  uint32
	unknown1      uint32
	fileSize      uint32
	hashEntries   []*textDictionaryHashEntry
	Entries       map[string]string
}

func (td *TextDictionary) loadHashEntries(br *d2datautils.StreamReader) error {
	var err error

	for i := 0; i < len(td.hashEntries); i++ {
		entry := textDictionaryHashEntry{}

		entry.active, err = br.ReadByte()
		if err != nil {
			return err
		}

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

		td.hashEntries[i] = &entry
	}

	for idx := range td.hashEntries {
		if !td.hashEntries[idx].IsActive() {
			continue
		}

		if err := td.loadHashEntry(idx, td.hashEntries[idx], br); err != nil {
			return err
		}
	}

	return nil
}

func (td *TextDictionary) loadHashEntry(idx int, hashEntry *textDictionaryHashEntry, br *d2datautils.StreamReader) error {
	var err error

	br.SetPosition(uint64(hashEntry.NameString))

	nameVal, err := br.ReadBytes(int(hashEntry.NameLength - 1))
	if err != nil {
		return err
	}

	hashEntry.name = string(nameVal)

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

	hashEntry.key = key

	if hashEntry.key == "x" || hashEntry.key == "X" {
		key = "#" + strconv.Itoa(idx)
	}

	_, exists := td.Entries[key]
	if !exists {
		td.Entries[key] = hashEntry.name
	}

	return nil
}

type textDictionaryHashEntry struct {
	name        string
	key         string
	HashValue   uint32
	IndexString uint32
	NameString  uint32
	Index       uint16
	NameLength  uint16
	active      byte
}

func (t *textDictionaryHashEntry) IsActive() bool {
	return t.active > 0
}

const (
	crcByteCount = 2
)

// LoadTextDictionary loads the text dictionary from the given data
func LoadTextDictionary(dictionaryData []byte) (*TextDictionary, error) {
	var err error

	lookupTable := &TextDictionary{
		Entries: make(map[string]string),
	}

	br := d2datautils.CreateStreamReader(dictionaryData)

	// skip past the CRC
	lookupTable.crcBytes, err = br.ReadBytes(crcByteCount)
	if err != nil {
		return nil, err
	}

	numberOfElements, err := br.ReadUInt16()
	if err != nil {
		return nil, err
	}

	lookupTable.hashTableSize, err = br.ReadUInt32()
	if err != nil {
		return nil, err
	}

	// Version (always 0)
	if lookupTable.version, err = br.ReadByte(); err != nil {
		return nil, errors.New("error reading Version record")
	}

	// StringOffset
	lookupTable.stringOffset, err = br.ReadUInt32()
	if err != nil {
		return nil, errors.New("error reading string offset")
	}

	// When the number of times you have missed a match with a
	// hash key equals this value, you give up because it is not there.
	lookupTable.unknown1, err = br.ReadUInt32()
	if err != nil {
		return nil, err
	}

	// FileSize
	lookupTable.fileSize, err = br.ReadUInt32()
	if err != nil {
		return nil, err
	}

	elementIndex := make([]uint16, numberOfElements)
	for i := 0; i < int(numberOfElements); i++ {
		elementIndex[i], err = br.ReadUInt16()
		if err != nil {
			return nil, err
		}
	}

	lookupTable.elementIndex = elementIndex

	lookupTable.hashEntries = make([]*textDictionaryHashEntry, lookupTable.hashTableSize)

	err = lookupTable.loadHashEntries(br)
	if err != nil {
		return nil, err
	}

	return lookupTable, nil
}

// Marshal encodes text dictionary back to byte slice
func (td *TextDictionary) Marshal() []byte {
	// create stream writter
	sw := d2datautils.CreateStreamWriter()

	sw.PushBytes(td.crcBytes...)
	sw.PushUint16(uint16(len(td.elementIndex)))
	sw.PushUint32(td.hashTableSize)
	sw.PushBytes(td.version)
	sw.PushUint32(td.stringOffset)
	sw.PushUint32(td.unknown1)

	sw.PushUint32(td.fileSize)

	for _, i := range td.elementIndex {
		sw.PushUint16(i)
	}

	for i := 0; i < len(td.hashEntries); i++ {
		sw.PushBytes(td.hashEntries[i].active)
		sw.PushUint16(td.hashEntries[i].Index)
		sw.PushUint32(td.hashEntries[i].HashValue)
		sw.PushUint32(td.hashEntries[i].IndexString)
		sw.PushUint32(td.hashEntries[i].NameString)
		sw.PushUint16(td.hashEntries[i].NameLength)
	}

	// values are table entries data (key & values)
	var values map[int]string = make(map[int]string)
	// valuesSorted are sorted values
	var valuesSorted map[int]string = make(map[int]string)

	// add values key / names to map
	for _, i := range td.hashEntries {
		values[int(i.IndexString)] = i.key
		values[int(i.NameString)] = i.name
	}

	// added map keys
	keys := make([]int, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}

	// sort keys
	sort.Ints(keys)

	// create sorted values map
	for _, k := range keys {
		valuesSorted[k] = values[k]
	}

	// add first value (without 0-byte separator)
	sw.PushBytes([]byte(valuesSorted[keys[0]])...)

	// adds values to result
	for i := 1; i < len(valuesSorted); i++ {
		sw.PushBytes([]byte(valuesSorted[keys[i]])...)
		sw.PushBytes(0)
	}

	return sw.GetBytes()
}
