package d2datadict

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

type ObjectTypeRecord struct {
	Name  string
	Token string
}

var ObjectTypes []ObjectTypeRecord

func LoadObjectTypes(objectTypeData []byte) {
	streamReader := d2common.CreateStreamReader(objectTypeData)
	count := streamReader.GetInt32()
	ObjectTypes = make([]ObjectTypeRecord, count)
	for i := range ObjectTypes {
		nameBytes, _ := streamReader.ReadBytes(32)
		tokenBytes, _ := streamReader.ReadBytes(20)
		ObjectTypes[i] = ObjectTypeRecord{
			Name:  strings.TrimSpace(strings.ReplaceAll(string(nameBytes), string(0), "")),
			Token: strings.TrimSpace(strings.ReplaceAll(string(tokenBytes), string(0), "")),
		}
	}
	log.Printf("Loaded %d object types", len(ObjectTypes))
}
