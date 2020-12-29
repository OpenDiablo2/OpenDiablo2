package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

// IndexedObjects is a slice of object records for quick lookups.
// nil checks should be done for uninitialized values at each level.
// [Act 1-5][Type 1-2][ID 0-855]
//nolint:gochecknoglobals // Currently global by design
type IndexedObjects [][][]*ObjectLookupRecord

// ObjectLookupRecord is a representation of a row from objectLookups.txt
type ObjectLookupRecord struct {
	Act           int
	Type          d2enum.ObjectType
	Id            int //nolint:golint,stylecheck // ID is the right key
	Name          string
	Description   string
	ObjectsTxtId  int //nolint:golint,stylecheck // ID is the right key
	MonstatsTxtId int //nolint:golint,stylecheck // ID is the right key
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
