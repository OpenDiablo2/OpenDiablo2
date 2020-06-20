package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// ObjectLookupRecord is a representation of a row from objects.txt
type ObjectLookupRecord struct {
	Act           int
	Type          d2enum.ObjectType
	Id            int //nolint:golint Id is the right key
	Name          string
	Description   string
	ObjectsTxtId  int //nolint:golint Id is the right key
	MonstatsTxtId int //nolint:golint Id is the right key
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

// LookupObject looks up an object record
func LookupObject(act, typ, id int) *ObjectLookupRecord {
	object := lookupObject(act, typ, id, indexedObjects)
	if object == nil {
		log.Panicf("Failed to look up object Act: %d, Type: %d, Id: %d", act, typ, id)
	}

	return object
}

func lookupObject(act, typ, id int, objects [][][]*ObjectLookupRecord) *ObjectLookupRecord {
	if len(objects) < act {
		return nil
	}

	if len(objects[act]) < typ {
		return nil
	}

	if len(objects[act][typ]) < id {
		return nil
	}

	return objects[act][typ][id]
}

// InitObjectRecords loads ObjectRecords
func InitObjectRecords() {
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

// IndexedObjects is a slice of object records for quick lookups.
// nil checks should be done for uninitialized values at each level.
// [Act 1-5][Type 1-2][Id 0-855]
//nolint:gochecknoglobals // Currently global by design
var indexedObjects [][][]*ObjectLookupRecord
