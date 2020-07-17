package d2stats

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"

func CreateStatList(stats ...*Stat) *StatList {
	return &StatList{stats}
}

// StatList is a list that contains stats
type StatList struct {
	stats []*Stat
}

// Clone returns a deep copy of the stat list
func (sl *StatList) Clone() *StatList {
	clone := &StatList{}
	clone.stats = make([]*Stat, len(sl.stats))

	for idx := range sl.stats {
		clone.stats[idx] = sl.stats[idx].Clone()
	}

	return clone
}

// Reduce returns a new stat list, combining like stats
func (sl *StatList) Reduce() *StatList {
	clone := sl.Clone()
	reduction := make([]*Stat, 0)

	// for quick lookups
	lookup := make(map[*d2datadict.ItemStatCostRecord]int)

	for len(clone.stats) > 0 {
		applied := false
		stat := clone.removeStat(0)

		// use lookup, may have found it already
		if idx, found := lookup[stat.Record]; found {
			if success := reduction[idx].combine(stat); success {
				continue
			}

			reduction = append(reduction, stat)
		}

		for idx := range reduction {
			if reduction[idx].combine(stat) {
				lookup[stat.Record] = idx
				applied = true

				break
			}
		}

		if !applied {
			reduction = append(reduction, stat)
		}
	}

	clone.stats = reduction

	return clone
}

func (sl *StatList) removeStat(idx int) *Stat {
	picked := sl.stats[idx]
	sl.stats[idx] = sl.stats[len(sl.stats)-1]
	sl.stats[len(sl.stats)-1] = nil
	sl.stats = sl.stats[:len(sl.stats)-1]

	return picked
}

// Append returns a new stat list, combining like stats
func (sl *StatList) Append(other *StatList) *StatList {
	clone := sl.Clone()
	clone.stats = append(clone.stats, other.stats...)

	return clone
}
