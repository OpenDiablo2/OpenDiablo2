package diablo2stats

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

// nolint:gochecknoglobals // just a test
var testAssetManager *d2asset.AssetManager

// nolint:gochecknoglobals // just a test
var testStatFactory *StatFactory

func TestSetup_StatList(t *testing.T) {
	testAssetManager = &d2asset.AssetManager{}
	testAssetManager.Records = &d2records.RecordManager{}

	testStatFactory, _ = NewStatFactory(testAssetManager)

	testAssetManager.Records.Item.Stats = itemStatCosts
	testAssetManager.Records.Character.Stats = charStats
	testAssetManager.Records.Skill.Details = skillDetails
	testAssetManager.Records.Monster.Stats = monStats
}

func TestDiablo2StatList_Index(t *testing.T) {
	strength := testStatFactory.NewStat("strength", 10)

	list1 := &Diablo2StatList{stats: []d2stats.Stat{strength}}
	if list1.Index(0) != strength {
		t.Error("list should contain a stat")
	}
}

func TestStatList_Clone(t *testing.T) {
	strength := testStatFactory.NewStat("strength", 10)

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
	stats := []d2stats.Stat{
		testStatFactory.NewStat("strength", 1),
		testStatFactory.NewStat("strength", 1),
		testStatFactory.NewStat("strength", 1),
		testStatFactory.NewStat("strength", 1),
	}

	list := testStatFactory.NewStatList(stats...)
	reduction := list.ReduceStats()

	if len(reduction.Stats()) != 1 || reduction.Index(0).String() != "+4 to Strength" {
		t.Errorf("diablo2Stat reduction failed")
	}

	stats = []d2stats.Stat{
		testStatFactory.NewStat("strength", 1),
		testStatFactory.NewStat("energy", 1),
		testStatFactory.NewStat("dexterity", 1),
		testStatFactory.NewStat("vitality", 1),
	}

	list = testStatFactory.NewStatList(stats...)
	reduction = list.ReduceStats()

	if len(reduction.Stats()) != 4 {
		t.Errorf("diablo2Stat reduction failed")
	}
}

func TestStatList_Append(t *testing.T) {
	list1 := &Diablo2StatList{
		[]d2stats.Stat{
			testStatFactory.NewStat("strength", 1),
			testStatFactory.NewStat("energy", 1),
			testStatFactory.NewStat("dexterity", 1),
			testStatFactory.NewStat("vitality", 1),
		},
	}
	list2 := list1.Clone()

	list3 := list1.AppendStatList(list2)

	if len(list3.Stats()) != 8 {
		t.Errorf("diablo2Stat append failed")
	}

	if len(list3.ReduceStats().Stats()) != 4 {
		t.Errorf("diablo2Stat append failed")
	}
}
