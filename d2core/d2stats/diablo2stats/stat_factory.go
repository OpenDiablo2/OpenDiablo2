package diablo2stats

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

// NewStatFactory creates a new stat factory instance
func NewStatFactory(asset *d2asset.AssetManager) (*StatFactory, error) {
	factory := &StatFactory{asset: asset}

	return factory, nil
}

// StatFactory is responsible for creating stats
type StatFactory struct {
	asset *d2asset.AssetManager
}

// NewStat creates a stat instance with the given record and values
func (f *StatFactory) NewStat(key string, values ...float64) d2stats.Stat {
	record := f.asset.Records.Item.Stats[key]

	if record == nil {
		return nil
	}

	stat := &diablo2Stat{
		factory: f,
		record:  record,
	}

	stat.init(values...) // init stat values, value types, and value combination rules

	return stat
}

// NewStatList creates a stat list
func (f *StatFactory) NewStatList(stats ...d2stats.Stat) d2stats.StatList {
	return &Diablo2StatList{stats}
}

// NewValue creates a stat value of the given type
func (f *StatFactory) NewValue(t d2stats.StatNumberType, c d2stats.ValueCombineType) d2stats.StatValue {
	sv := &Diablo2StatValue{
		numberType:  t,
		combineType: c,
	}

	switch t {
	case d2stats.StatValueFloat:
		sv.stringerFn = f.stringerUnsignedFloat
	case d2stats.StatValueInt:
		sv.stringerFn = f.stringerUnsignedInt
	default:
		sv.stringerFn = f.stringerEmpty
	}

	return sv
}

const (
	monsterNotFound = "{Monster not found!}"
)

func (f *StatFactory) getHeroMap() []d2enum.Hero {
	return []d2enum.Hero{
		d2enum.HeroAmazon,
		d2enum.HeroSorceress,
		d2enum.HeroNecromancer,
		d2enum.HeroPaladin,
		d2enum.HeroBarbarian,
		d2enum.HeroDruid,
		d2enum.HeroAssassin,
	}
}

func (f *StatFactory) stringerUnsignedInt(sv d2stats.StatValue) string {
	return fmt.Sprintf("%d", sv.Int())
}

func (f *StatFactory) stringerUnsignedFloat(sv d2stats.StatValue) string {
	return fmt.Sprintf("%.2f", sv.Float())
}

func (f *StatFactory) stringerEmpty(_ d2stats.StatValue) string {
	return ""
}

func (f *StatFactory) stringerIntSigned(sv d2stats.StatValue) string {
	return fmt.Sprintf("%+d", sv.Int())
}

func (f *StatFactory) stringerIntPercentageSigned(sv d2stats.StatValue) string {
	return fmt.Sprintf("%+d%%", sv.Int())
}

func (f *StatFactory) stringerIntPercentageUnsigned(sv d2stats.StatValue) string {
	return fmt.Sprintf("%d%%", sv.Int())
}

func (f *StatFactory) stringerClassAllSkills(sv d2stats.StatValue) string {
	heroIndex := sv.Int()

	heroMap := f.getHeroMap()
	classRecord := f.asset.Records.Character.Stats[heroMap[heroIndex]]

	return f.asset.TranslateString(classRecord.SkillStrAll)
}

func (f *StatFactory) stringerClassOnly(sv d2stats.StatValue) string {
	heroMap := f.getHeroMap()
	heroIndex := sv.Int()
	classRecord := f.asset.Records.Character.Stats[heroMap[heroIndex]]
	classOnlyKey := classRecord.SkillStrClassOnly

	return f.asset.TranslateString(classOnlyKey)
}

func (f *StatFactory) stringerSkillName(sv d2stats.StatValue) string {
	skillRecord := f.asset.Records.Skill.Details[sv.Int()]
	return skillRecord.Skill
}

func (f *StatFactory) stringerMonsterName(sv d2stats.StatValue) string {
	for key := range f.asset.Records.Monster.Stats {
		if f.asset.Records.Monster.Stats[key].ID == sv.Int() {
			return f.asset.Records.Monster.Stats[key].NameString
		}
	}

	return monsterNotFound
}
