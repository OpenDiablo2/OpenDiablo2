package diablo2stats

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

// NewStat creates a stat instance with the given record and values
func NewStat(key string, values ...float64) d2stats.Stat {
	record := d2datadict.ItemStatCosts[key]

	if record == nil {
		return nil
	}

	stat := &diablo2Stat{
		record: record,
	}

	stat.init(values...) // init stat values, value types, and value combination rules

	return stat
}

// NewStatList creates a stat list
func NewStatList(stats ...d2stats.Stat) d2stats.StatList {
	return &Diablo2StatList{stats}
}

// NewValue creates a stat value of the given type
func NewValue(t d2stats.StatNumberType, c d2stats.ValueCombineType) d2stats.StatValue {
	sv := &Diablo2StatValue{
		numberType:  t,
		combineType: c,
	}

	switch t {
	case d2stats.StatValueFloat:
		sv.stringerFn = stringerUnsignedFloat
	case d2stats.StatValueInt:
		sv.stringerFn = stringerUnsignedInt
	default:
		sv.stringerFn = stringerEmpty
	}

	return sv
}
