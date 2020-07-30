package d2item

import "github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"

// StatContext is anything which has a `StatList` method which yields a StatList.
// This is used for resolving stat dependencies for showing actual values, like
// stats that are based off of the current character level
type StatContext interface {
	Equipper
	BaseStatList() d2stats.StatList
	StatList() d2stats.StatList
}
