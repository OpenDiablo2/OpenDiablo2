package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"

// recordLoader represents something that can load a data dictionary and
// handles placing it in the record manager exports
type recordLoader func(r *RecordManager, dictionary *d2txt.DataDictionary) error
