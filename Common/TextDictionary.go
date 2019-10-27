package Common

import (
	"log"
	"strconv"

	"github.com/essial/OpenDiablo2/ResourcePaths"
)

type textDictionaryHashEntry struct {
	IsActive    bool
	Index       uint16
	HashValue   uint32
	IndexString uint32
	NameString  uint32
	NameLength  uint16
}

var lookupTable map[string]string

func TranslateString(key string) string {
	result, ok := lookupTable[key]
	if !ok {
		log.Panic("Could not find a string for the key '%s'", key)
	}
	return result
}

func LoadTextDictionary(fileProvider FileProvider) {
	lookupTable = make(map[string]string)

	loadDictionary(fileProvider, ResourcePaths.PatchStringTable)
	loadDictionary(fileProvider, ResourcePaths.ExpansionStringTable)
	loadDictionary(fileProvider, ResourcePaths.StringTable)
	log.Printf("Loaded %d entries from the string table", len(lookupTable))
}

func loadDictionary(fileProvider FileProvider, dictionaryName string) {
	dictionaryData := fileProvider.LoadFile(dictionaryName)
	br := CreateStreamReader(dictionaryData)
	br.ReadBytes(2) // CRC
	numberOfElements := br.GetUInt16()
	hashTableSize := br.GetUInt32()
	br.ReadByte()  // Version (always 0)
	br.GetUInt32() // StringOffset
	br.GetUInt32() // When the number of times you have missed a match with a hash key equals this value, you give up because it is not there.
	br.GetUInt32() // FileSize

	elementIndex := make([]uint16, numberOfElements)
	for i := 0; i < int(numberOfElements); i++ {
		elementIndex[i] = br.GetUInt16()
	}

	hashEntries := make([]textDictionaryHashEntry, hashTableSize)
	for i := 0; i < int(hashTableSize); i++ {
		hashEntries[i] = textDictionaryHashEntry{
			br.GetByte() == 1,
			br.GetUInt16(),
			br.GetUInt32(),
			br.GetUInt32(),
			br.GetUInt32(),
			br.GetUInt16(),
		}
	}

	for idx, hashEntry := range hashEntries {
		if !hashEntry.IsActive {
			continue
		}
		br.SetPosition(uint64(hashEntry.NameString))
		nameVal, _ := br.ReadBytes(int(hashEntry.NameLength - 1))
		value := string(nameVal)
		br.SetPosition(uint64(hashEntry.IndexString))
		key := ""
		for true {
			b := br.GetByte()
			if b == 0 {
				break
			}
			key += string(b)
		}
		if key == "x" || key == "X" {
			key = "#" + strconv.Itoa(idx)
		}
		_, exists := lookupTable[key]
		if !exists {
			lookupTable[key] = value

		}
		// Use the following code to write out the values
		/*
			f, err := os.OpenFile(`C:\Users\lunat\Desktop\D2\langdict.txt`,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString("\n[" + key + "] " + value); err != nil {
				log.Println(err)
			}
		*/
	}
}
