package diablo2stats

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

// CreateStat creates a stat instance with the given record and values
func CreateStat(record *d2datadict.ItemStatCostRecord, values ...d2stats.StatValue) d2stats.Stat {
	if record == nil {
		return nil
	}

	stat := &Diablo2Stat{
		record: record,
		values: values,
	}

	return stat
}

// CreateStatList creates a stat list
func CreateStatList(stats ...d2stats.Stat) d2stats.StatList {
	return &Diablo2StatList{stats}
}

// CreateStatValue creates a stat value of the given type
func CreateStatValue(t d2stats.StatValueType) d2stats.StatValue {
	sv := &Diablo2StatValue{_type: t}

	switch t {
	case d2stats.StatValueFloat:
		sv._stringer = defaultFloatStringer
	case d2stats.StatValueInt:
		sv._stringer = defaultIntStringer
	default:
		sv._stringer = defaultStringer
	}

	return sv
}

func intVal(i int) d2stats.StatValue {
	return CreateStatValue(d2stats.StatValueInt).SetInt(i)
}

func floatVal(f float64) d2stats.StatValue {
	return CreateStatValue(d2stats.StatValueFloat).SetFloat(f)
}

func defaultIntStringer(sv d2stats.StatValue) string {
	return fmt.Sprintf("%d", sv.Int())
}

func defaultFloatStringer(sv d2stats.StatValue) string {
	return fmt.Sprintf("%.2f", sv.Float())
}

func defaultStringer(_ d2stats.StatValue) string {
	return ""
}
