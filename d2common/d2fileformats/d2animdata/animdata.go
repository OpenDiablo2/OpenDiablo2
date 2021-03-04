package d2animdata

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
)

const (
	numBlocks             = 256
	maxRecordsPerBlock    = 67
	byteCountName         = 8
	byteCountSpeedPadding = 2
	numEvents             = 144
	speedDivisor          = 256
	speedBaseFPS          = 25
	milliseconds          = 1000
)

// AnimationData is a representation of the binary data from `data/global/AnimData.d2`
type AnimationData struct {
	hashTable
	blocks  [numBlocks]*block
	entries map[string][]*AnimationDataRecord
}

// GetRecordNames returns a slice of all record name strings
func (ad *AnimationData) GetRecordNames() []string {
	result := make([]string, 0)

	for name := range ad.entries {
		result = append(result, name)
	}

	return result
}

// GetRecord returns a single AnimationDataRecord with the given name string. If there is more
// than one record with the given name string, the last record entry will be returned.
func (ad *AnimationData) GetRecord(name string) *AnimationDataRecord {
	records := ad.GetRecords(name)
	numRecords := len(records)

	if numRecords < 1 {
		return nil
	}

	return records[numRecords-1]
}

// GetRecords returns all records that have the given name string. The AnimData.d2 files have
// multiple records with the same name, but other values in the record are different.
func (ad *AnimationData) GetRecords(name string) []*AnimationDataRecord {
	return ad.entries[name]
}

// GetRecordsCount returns number of animation data records
func (ad *AnimationData) GetRecordsCount() int {
	return len(ad.entries)
}

// PushRecord adds a new record to entry named 'name'
func (ad *AnimationData) PushRecord(name string) {
	ad.entries[name] = append(
		ad.entries[name],
		&AnimationDataRecord{
			name: name,
		},
	)
}

// DeleteRecord teletes specified index from specified entry
func (ad *AnimationData) DeleteRecord(name string, recordIdx int) error {
	newRecords := make([]*AnimationDataRecord, 0)

	for n, i := range ad.entries[name] {
		if n == recordIdx {
			continue
		}

		newRecords = append(newRecords, i)
	}

	if len(ad.entries[name]) == len(newRecords) {
		return fmt.Errorf("index %d not found", recordIdx)
	}

	ad.entries[name] = newRecords

	return nil
}

// AddEntry adds a new animation entry with name given
func (ad *AnimationData) AddEntry(name string) error {
	_, found := ad.entries[name]
	if found {
		return fmt.Errorf("entry of name %s already exist", name)
	}

	ad.entries[name] = make([]*AnimationDataRecord, 0)

	return nil
}

// DeleteEntry deltees entry with specified name
func (ad *AnimationData) DeleteEntry(name string) error {
	_, found := ad.entries[name]
	if !found {
		return fmt.Errorf("entry named %s doesn't exist", name)
	}

	delete(ad.entries, name)

	return nil
}

// Load loads the data into an AnimationData struct
//nolint:gocognit,funlen // can't reduce
func Load(data []byte) (*AnimationData, error) {
	reader := d2datautils.CreateStreamReader(data)
	animdata := &AnimationData{}
	hashIdx := 0

	animdata.entries = make(map[string][]*AnimationDataRecord)

	for blockIdx := range animdata.blocks {
		recordCount, err := reader.ReadUInt32()
		if err != nil {
			return nil, err
		}

		if recordCount > maxRecordsPerBlock {
			return nil, fmt.Errorf("more than %d records in block", maxRecordsPerBlock)
		}

		records := make([]*AnimationDataRecord, recordCount)

		for recordIdx := uint32(0); recordIdx < recordCount; recordIdx++ {
			nameBytes, err := reader.ReadBytes(byteCountName)
			if err != nil {
				return nil, err
			}

			if nameBytes[byteCountName-1] != byte(0) {
				return nil, errors.New("animdata AnimationDataRecord name missing null terminator byte")
			}

			name := string(nameBytes)
			name = strings.ReplaceAll(name, string(byte(0)), "")

			animdata.hashTable[hashIdx] = hashName(name)

			frames, err := reader.ReadUInt32()
			if err != nil {
				return nil, err
			}

			speed, err := reader.ReadUInt16()
			if err != nil {
				return nil, err
			}

			reader.SkipBytes(byteCountSpeedPadding)

			events := make(map[int]AnimationEvent)

			for eventIdx := 0; eventIdx < numEvents; eventIdx++ {
				eventByte, err := reader.ReadByte()
				if err != nil {
					return nil, err
				}

				event := AnimationEvent(eventByte)
				if event != AnimationEventNone {
					events[eventIdx] = event
				}
			}

			r := &AnimationDataRecord{
				name,
				frames,
				speed,
				events,
			}

			records[recordIdx] = r

			if _, found := animdata.entries[r.name]; !found {
				animdata.entries[r.name] = make([]*AnimationDataRecord, 0)
			}

			animdata.entries[r.name] = append(animdata.entries[r.name], r)
		}

		b := &block{
			recordCount,
			records,
		}

		animdata.blocks[blockIdx] = b
	}

	if reader.Position() != uint64(len(data)) {
		return nil, fmt.Errorf("unable to parse animation data: %d != %d", reader.Position(), len(data))
	}

	return animdata, nil
}

// Marshal encodes animation data back into byte slice
// basing on AnimationData.records
func (ad *AnimationData) Marshal() []byte {
	sw := d2datautils.CreateStreamWriter()

	// keys - all entries in animationData
	keys := make([]string, len(ad.entries))

	// we must manually add index
	idx := 0

	for i := range ad.entries {
		keys[idx] = i
		idx++
	}

	// name terminates current name
	name := 0

	// recordIdx determinates current record index
	recordIdx := 0

	// numberOfEntries is a number of entries in all map indexes
	var numberOfEntries int = 0

	for i := 0; i < len(keys); i++ {
		numberOfEntries += len(ad.entries[keys[i]])
	}

	for idx := 0; idx < numBlocks; idx++ {
		// number of records (max is maxRecordsPerObject)
		l := 0

		switch {
		// first condition: end up with all this and push 0 to dhe end
		case numberOfEntries == 0:
			sw.PushUint32(0)
			continue
		case numberOfEntries < maxRecordsPerBlock:
			// second condition - if number of entries left is smaller than
			// maxRecordsPerBlock, push...
			l = numberOfEntries
			sw.PushUint32(uint32(l))
		default:
			// else use maxRecordsPerBlock
			l = maxRecordsPerBlock
			sw.PushUint32(maxRecordsPerBlock)
		}

		for currentRecordIdx := 0; currentRecordIdx < l; currentRecordIdx++ {
			numberOfEntries--

			if recordIdx == len(ad.entries[keys[name]]) {
				recordIdx = 0
				name++
			}

			animationRecord := ad.entries[keys[name]][recordIdx]
			recordIdx++

			name := animationRecord.name
			missingZeroBytes := byteCountName - len(name)
			sw.PushBytes([]byte(name)...)

			for i := 0; i < missingZeroBytes; i++ {
				sw.PushBytes(0)
			}

			sw.PushUint32(animationRecord.framesPerDirection)
			sw.PushUint16(animationRecord.speed)

			for i := 0; i < byteCountSpeedPadding; i++ {
				sw.PushBytes(0)
			}

			for event := 0; event < numEvents; event++ {
				sw.PushBytes(byte(animationRecord.events[event]))
			}
		}
	}

	return sw.GetBytes()
}
