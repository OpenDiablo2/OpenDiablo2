package d2stats

import (
	"fmt"
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
)

// CreateStat creates a stat instance with the given ID and number of values
func CreateStat(record *d2datadict.ItemStatCostRecord, values []int) *Stat {
	if record == nil {
		return nil
	}

	numRequiredValues := record.NumStatValues()

	// if we have differing value counts, make a new set of values
	// and copy any of the provided values into the new slice
	if len(values) != numRequiredValues {
		newValues := make([]int, numRequiredValues)

		numToCopy := min(len(values), numRequiredValues)
		for idx := 0; idx < numToCopy; idx++ {
			newValues[idx] = values[idx]
		}

		values = newValues
	}

	stat := &Stat{
		Record: record,
		Values: values,
	}

	return stat
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Stat is an instance of a Stat, with a set of Values
type Stat struct {
	Record *d2datadict.ItemStatCostRecord
	Values []int
}

// Clone returns a deep copy of the Stat
func (s Stat) Clone() *Stat {
	clone := &Stat{
		Record: s.Record,
		Values: make([]int, len(s.Values)),
	}

	for idx := range s.Values {
		clone.Values[idx] = s.Values[idx]
	}

	return clone
}

// Description returns the formatted description string
func (s *Stat) Description() string {
	str := make([]string, len(s.Values)+1)
	for idx := range s.Values {
		str[idx] = fmt.Sprintf("%v", s.Values[idx])
	}

	toLookUp := s.Record.DescStrPos

	if len(s.Values) > 0 {
		if s.Values[0] < 0 {
			toLookUp = s.Record.DescStrNeg
		}
	}

	str[len(str)-1] = d2common.TranslateString(toLookUp)

	return s.Record.DescString(str...)
}
