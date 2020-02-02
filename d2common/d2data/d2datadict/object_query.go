package d2datadict

import (
	"log"
)

type ObjectType int

const (
	ObjectTypeCharacter ObjectType = 1
	ObjectTypeItem      ObjectType = 2
)

type ObjectLookupRecord struct {
	Act           int
	Type          ObjectType
	Id            int
	Description   string
	ObjectsTxtId  int
	MonstatsTxtId int
	Direction     int
	Base          string
	Token         string
	Mode          string
	Class         string
	HD            string
	TR            string
	LG            string
	RA            string
	LA            string
	RH            string
	LH            string
	SH            string
	S1            string
	S2            string
	S3            string
	S4            string
	S5            string
	S6            string
	S7            string
	S8            string
	ColorMap      string
	Index         int
}

func LookupObject(act, typ, id int) *ObjectLookupRecord {
	object := lookupObject(act, typ, id, indexedObjects)
	if object == nil {
		log.Panicf("Failed to look up object Act: %d, Type: %d, Id: %d", act, typ, id)
	}
	return object
}

func lookupObject(act, typ, id int, objects [][][]*ObjectLookupRecord) *ObjectLookupRecord {
	if objects[act] != nil && objects[act][typ] != nil && objects[act][typ][id] != nil {
		return objects[act][typ][id]
	}
	return nil
}

func init() {
	indexedObjects = indexObjects(objectLookups)
}

func indexObjects(objects []ObjectLookupRecord) [][][]*ObjectLookupRecord {
	// Allocating 6 to allow Acts 1-5 without requiring a -1 at every read.
	indexedObjects = make([][][]*ObjectLookupRecord, 6)
	for i := range objects {
		record := &objects[i]
		if indexedObjects[record.Act] == nil {
			// Likewise allocating 3 so a -1 isn't necessary.
			indexedObjects[record.Act] = make([][]*ObjectLookupRecord, 3)
		}

		if indexedObjects[record.Act][record.Type] == nil {
			// For simplicity, allocating with length 1000 then filling the values in by index.
			// If ids in the dictionary ever surpass 1000, raise this number.
			indexedObjects[record.Act][record.Type] = make([]*ObjectLookupRecord, 1000)
		}

		indexedObjects[record.Act][record.Type][record.Id] = record
	}

	return indexedObjects
}

// Indexed slice of object records for quick lookups.
// nil checks should be done for uninitialized values at each level.
// [Act 1-5][Type 1-2][Id 0-855]
var indexedObjects [][][]*ObjectLookupRecord
