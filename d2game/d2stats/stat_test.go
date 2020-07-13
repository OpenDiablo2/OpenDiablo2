package d2stats

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2config"
	"testing"
)

func TestStat_AssetInit(t *testing.T) {
	d2config.Load()
	d2asset.Initialize(nil, nil)

	tablePaths := []string{
		d2resource.PatchStringTable,
		d2resource.ExpansionStringTable,
		d2resource.StringTable,
	}

	for _, tablePath := range tablePaths {
		data, err := d2asset.LoadFile(tablePath)
		if err != nil {
			panic(err)
		}

		d2common.LoadTextDictionary(data)
	}

	data, err := d2asset.LoadFile(d2resource.ItemStatCost)
	if err != nil {
		panic(err)
	}
	d2datadict.LoadItemStatCosts(data)
}


func TestStat_Clone(t *testing.T) {
	r := d2datadict.ItemStatCosts["strength"]
	s1 := &Stat{Record: r, Values: []int{5}}

	s1.Values[0] = 5
	s2 := s1.Clone()

	// make sure the stats are distinct
	if &s1 == &s2 {
		t.Errorf("stats share the same pointer %d == %d", &s1, &s2)
	}

	// make sure the stat values are unique
	vs1, vs2 := s1.Values, s2.Values
	if &vs1 == &vs2 {
		t.Errorf("stat values share the same pointer %d == %d", &s1, &s2)
	}

	s2.Values[0] = 6
	v1, v2 := s1.Values[0], s2.Values[0]

	// make sure the value ranges are distinct
	if v1 == v2 {
		t.Errorf("stat value ranges should not be equal")
	}
}

func TestStat_Description(t *testing.T) {
	r := d2datadict.ItemStatCosts["strength"]

	s1 := &Stat{Record: r, Values: []int{5}}
	desc := s1.Description()

	if desc != "+5 to Strength" {
		t.Errorf("unexpected description string: %s", desc)
	}

	s1.Values[0] = -5
	desc = s1.Description()

	if desc != "-5 to Strength" {
		t.Errorf("unexpected description string: %s", desc)
	}
}
