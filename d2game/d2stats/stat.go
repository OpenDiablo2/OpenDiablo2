package d2stats

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
)

// CreateStat creates a stat instance with the given ID and number of values
func CreateStat(record *d2datadict.ItemStatCostRecord, values ...int) *Stat {
	if record == nil {
		return nil
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
	return s.Record.DescString(s.Values...)
}

