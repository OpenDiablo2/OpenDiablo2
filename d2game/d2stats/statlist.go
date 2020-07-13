package d2stats

// StatList is a list that contains stats
type StatList struct {
	stats []*Stat
}

// Clone returns a deep copy of the stat list
func (sl StatList) Clone() *StatList {
	clone := &StatList{}
	clone.stats = make([]*Stat, len(sl.stats))

	for idx := range sl.stats {
		clone.stats[idx] = sl.stats[idx].Clone()
	}

	return clone
}
