package d2stats

// StatList is useful for reducing stats.
// They provide a context for stats to alter other stats or infer values
// during stat assignment/calculation
type StatList interface {
	Index(idx int) Stat
	Stats() []Stat
	SetStats([]Stat) StatList
	Clone() StatList
	ReduceStats() StatList
	RemoveStatAtIndex(idx int) Stat
	AppendStatList(other StatList) StatList
	Pop() Stat
	Push(Stat) StatList
}
