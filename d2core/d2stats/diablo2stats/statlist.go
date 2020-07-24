package diablo2stats

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

// static check that diablo2Stat implements Stat
var _ d2stats.StatList = &Diablo2StatList{}

// Diablo2StatList is a diablo 2 implementation of a stat list
type Diablo2StatList struct {
	stats []d2stats.Stat
}

// Index returns a stat given with the given index
func (sl *Diablo2StatList) Index(idx int) d2stats.Stat {
	if idx < 0 || idx > len(sl.stats) {
		return nil
	}

	return sl.stats[idx]
}

// Stats returns a slice of stats
func (sl *Diablo2StatList) Stats() []d2stats.Stat {
	return sl.stats
}

// SetStats sets the stats, given a slice of stats
func (sl *Diablo2StatList) SetStats(stats []d2stats.Stat) d2stats.StatList {
	sl.stats = stats
	return sl
}

// Pop removes the last stat from the stat list
func (sl *Diablo2StatList) Pop() d2stats.Stat {
	num := len(sl.stats)
	if num < 1 {
		return nil
	}

	idx := num - 1
	last := sl.stats[idx]
	sl.stats = sl.stats[:idx]

	return last
}

// Push adds a stat at the end of the stat list
func (sl *Diablo2StatList) Push(stat d2stats.Stat) d2stats.StatList {
	sl.stats = append(sl.stats, stat)

	return sl
}

// Clone returns a deep copy of the stat list
func (sl *Diablo2StatList) Clone() d2stats.StatList {
	clone := &Diablo2StatList{}
	clone.stats = make([]d2stats.Stat, len(sl.stats))

	for idx := range sl.stats {
		if stat := sl.Index(idx); stat != nil {
			clone.stats[idx] = stat.Clone()
		}
	}

	return clone
}

// ReduceStats combines like stats (does not alter this stat list, returns clone)
func (sl *Diablo2StatList) ReduceStats() d2stats.StatList {
	clone := sl.Clone()
	reduction := make([]d2stats.Stat, 0)

	// for quick lookups
	lookup := make(map[string]int)

	for len(clone.Stats()) > 0 {
		stat := clone.Pop()

		// if we find it in the lookup, immediately try to combine
		// if it doesn't combine, we append to the reduction
		if idx, found := lookup[stat.Name()]; found {
			if result, err := reduction[idx].Combine(stat); err == nil {
				reduction[idx] = result
				continue
			}
		}

		// we didnt find it in the lookup, so we will try to combine with other stats
		for idx := range reduction {
			if _, err := reduction[idx].Combine(stat); err == nil {
				continue
			}
		}

		lookup[stat.Name()] = len(lookup)

		reduction = append(reduction, stat)
	}

	return clone.SetStats(reduction)
}

// RemoveStatAtIndex removes the stat from the stat list, returns the stat
func (sl *Diablo2StatList) RemoveStatAtIndex(idx int) d2stats.Stat {
	picked := sl.stats[idx]
	sl.stats[idx] = sl.stats[len(sl.stats)-1]
	sl.stats[len(sl.stats)-1] = nil
	sl.stats = sl.stats[:len(sl.stats)-1]

	return picked
}

// AppendStatList adds the stats from the other stat list to this stat list
func (sl *Diablo2StatList) AppendStatList(other d2stats.StatList) d2stats.StatList {
	sl.stats = append(sl.stats, other.Stats()...)

	return sl
}
