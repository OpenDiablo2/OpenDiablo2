package diablo2stats

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

// NewStat creates a stat instance with the given record and values
func NewStat(record *d2datadict.ItemStatCostRecord, values ...d2stats.StatValue) d2stats.Stat {
	if record == nil {
		return nil
	}

	stat := &Diablo2Stat{
		record: record,
		values: values,
	}

	return stat
}

// NewStatList creates a stat list
func NewStatList(stats ...d2stats.Stat) d2stats.StatList {
	return &Diablo2StatList{stats}
}

// NewStatValue creates a stat value of the given type
func NewStatValue(t d2stats.StatValueType) d2stats.StatValue {
	sv := &Diablo2StatValue{_type: t}

	switch t {
	case d2stats.StatValueFloat:
		sv._stringer = stringerUnsignedFloat
	case d2stats.StatValueInt:
		sv._stringer = stringerUnsignedInt
	default:
		sv._stringer = stringerEmpty
	}

	return sv
}

func intVal(i int) d2stats.StatValue {
	return NewStatValue(d2stats.StatValueInt).SetInt(i)
}
