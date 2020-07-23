package diablo2stats

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

func TestDiablo2StatList_Index(t *testing.T) {
	record := d2datadict.ItemStatCosts["strength"]
	strength := NewStat(record, intVal(10))

	list1 := &Diablo2StatList{stats: []d2stats.Stat{strength}}
	if list1.Index(0) != strength {
		t.Error("list should contain a stat")
	}
}

func TestStatList_Clone(t *testing.T) {
	record := d2datadict.ItemStatCosts["strength"]
	strength := NewStat(record, intVal(10))

	list1 := &Diablo2StatList{}
	list1.Push(strength)

	list2 := list1.Clone()
	str1 := list1.Index(0).String()
	str2 := list2.Index(0).String()

	if str1 != str2 {
		t.Errorf("Stats of cloned stat list should be identitcal")
	}

	list2.Index(0).Values()[0].SetInt(0)

	if list1.Index(0).String() == list2.Index(0).String() {
		t.Errorf("Stats of cloned stat list should be different")
	}
}

func TestStatList_Reduce(t *testing.T) {
	records := []*d2datadict.ItemStatCostRecord{
		d2datadict.ItemStatCosts["strength"],
		d2datadict.ItemStatCosts["energy"],
		d2datadict.ItemStatCosts["dexterity"],
		d2datadict.ItemStatCosts["vitality"],
	}

	stats := []d2stats.Stat{
		NewStat(records[0], intVal(1)),
		NewStat(records[0], intVal(1)),
		NewStat(records[0], intVal(1)),
		NewStat(records[0], intVal(1)),
	}

	list := NewStatList(stats...)
	reduction := list.ReduceStats()

	if len(reduction.Stats()) != 1 || reduction.Index(0).String() != "+4 to Strength" {
		t.Errorf("Diablo2Stat reduction failed")
	}

	stats = []d2stats.Stat{
		NewStat(records[0], intVal(1)),
		NewStat(records[1], intVal(1)),
		NewStat(records[2], intVal(1)),
		NewStat(records[3], intVal(1)),
	}

	list = NewStatList(stats...)
	reduction = list.ReduceStats()

	if len(reduction.Stats()) != 4 {
		t.Errorf("Diablo2Stat reduction failed")
	}
}

func TestStatList_Append(t *testing.T) {
	records := []*d2datadict.ItemStatCostRecord{
		d2datadict.ItemStatCosts["strength"],
		d2datadict.ItemStatCosts["energy"],
		d2datadict.ItemStatCosts["dexterity"],
		d2datadict.ItemStatCosts["vitality"],
	}

	list1 := &Diablo2StatList{
		[]d2stats.Stat{
			NewStat(records[0], intVal(1)),
			NewStat(records[1], intVal(1)),
			NewStat(records[2], intVal(1)),
			NewStat(records[3], intVal(1)),
		},
	}
	list2 := list1.Clone()

	list3 := list1.AppendStatList(list2)

	if len(list3.Stats()) != 8 {
		t.Errorf("Diablo2Stat append failed")
	}

	if len(list3.ReduceStats().Stats()) != 4 {
		t.Errorf("Diablo2Stat append failed")
	}
}
