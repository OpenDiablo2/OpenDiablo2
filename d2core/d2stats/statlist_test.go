package d2stats

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
)

func TestStatList_Clone(t *testing.T) {
	record := d2datadict.ItemStatCosts["strength"]
	strength := CreateStat(record, 10)

	list1 := CreateStatList(strength)
	list2 := list1.Clone()

	if list1.stats[0].Description() != list2.stats[0].Description() {
		t.Errorf("Stats of cloned stat list should be identitcal")
	}

	list2.stats[0].Values[0] = 0
	if list1.stats[0].Description() == list2.stats[0].Description() {
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

	stats := []*Stat{
		CreateStat(records[0], 1),
		CreateStat(records[0], 1),
		CreateStat(records[0], 1),
		CreateStat(records[0], 1),
	}

	list := CreateStatList(stats...)
	reduction := list.Reduce()

	if len(reduction.stats) != 1 || reduction.stats[0].Description() != "+4 to Strength" {
		t.Errorf("Stat reduction failed")
	}

	stats = []*Stat{
		CreateStat(records[0], 1),
		CreateStat(records[1], 1),
		CreateStat(records[2], 1),
		CreateStat(records[3], 1),
	}

	list = CreateStatList(stats...)
	reduction = list.Reduce()

	if len(reduction.stats) != 4 {
		t.Errorf("Stat reduction failed")
	}
}

func TestStatList_Append(t *testing.T) {
	records := []*d2datadict.ItemStatCostRecord{
		d2datadict.ItemStatCosts["strength"],
		d2datadict.ItemStatCosts["energy"],
		d2datadict.ItemStatCosts["dexterity"],
		d2datadict.ItemStatCosts["vitality"],
	}

	list1 := &StatList{
		[]*Stat{
			CreateStat(records[0], 1),
			CreateStat(records[1], 1),
			CreateStat(records[2], 1),
			CreateStat(records[3], 1),
		},
	}
	list2 := list1.Clone()

	list3 := list1.Append(list2)

	if len(list3.stats) != 8 {
		t.Errorf("Stat append failed")
	}

	if len(list3.Reduce().stats) != 4 {
		t.Errorf("Stat append failed")
	}
}
