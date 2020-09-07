package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"log"
	"strings"
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
	streamReader := d2datautils.CreateStreamReader(objectTypeData)
	count := streamReader.GetInt32()
	ObjectTypes = make([]ObjectTypeRecord, count)

	const (
		nameSize  = 32
		tokenSize = 20
	)

	for i := range ObjectTypes {
		nameBytes := streamReader.ReadBytes(nameSize)
		tokenBytes := streamReader.ReadBytes(tokenSize)
		ObjectTypes[i] = ObjectTypeRecord{
			Name:  strings.TrimSpace(strings.ReplaceAll(string(nameBytes), string(byte(0)), "")),
			Token: strings.TrimSpace(strings.ReplaceAll(string(tokenBytes), string(byte(0)), "")),
		}
	}

	log.Printf("Loaded %d object types", len(ObjectTypes))
}
