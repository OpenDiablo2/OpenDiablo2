package Common

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/ResourcePaths"
)

type ObjectTypeRecord struct {
	Name  string
	Token string
}

var ObjectTypes []ObjectTypeRecord

func LoadObjectTypes(fileProvider FileProvider) {
	objectTypeData := fileProvider.LoadFile(ResourcePaths.ObjectType)
	streamReader := CreateStreamReader(objectTypeData)
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
