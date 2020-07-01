package d2datadict

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// ObjectTypeRecord is a representation of a row from objtype.txt
type ObjectTypeRecord struct {
	Name  string
	Token string
}

// ObjectTypes contains the name and token for objects
//nolint:gochecknoglobals // Currently global by design, only written once
var ObjectTypes []ObjectTypeRecord

// LoadObjectTypes loads ObjectTypeRecords from objtype.txt
func LoadObjectTypes(objectTypeData []byte) {
	streamReader := d2common.CreateStreamReader(objectTypeData)
	count := streamReader.GetInt32()
	ObjectTypes = make([]ObjectTypeRecord, count)

	for i := range ObjectTypes {
		nameBytes := streamReader.ReadBytes(32)
		tokenBytes := streamReader.ReadBytes(20)
		ObjectTypes[i] = ObjectTypeRecord{
			Name:  strings.TrimSpace(strings.ReplaceAll(string(nameBytes), string(0), "")),
			Token: strings.TrimSpace(strings.ReplaceAll(string(tokenBytes), string(0), "")),
		}
	}

	log.Printf("Loaded %d object types", len(ObjectTypes))
}
