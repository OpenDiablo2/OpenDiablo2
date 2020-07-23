package diablo2stats

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2stats"
)

const (
	monsterNotFound = "{Monster not found!}"
)

func getHeroMap() map[int]d2enum.Hero {
	return map[int]d2enum.Hero{
		int(d2enum.HeroAmazon):      d2enum.HeroAmazon,
		int(d2enum.HeroSorceress):   d2enum.HeroSorceress,
		int(d2enum.HeroNecromancer): d2enum.HeroNecromancer,
		int(d2enum.HeroPaladin):     d2enum.HeroPaladin,
		int(d2enum.HeroBarbarian):   d2enum.HeroBarbarian,
		int(d2enum.HeroDruid):       d2enum.HeroDruid,
		int(d2enum.HeroAssassin):    d2enum.HeroAssassin,
	}
}

func stringerIntSigned(sv d2stats.StatValue) string {
	return fmt.Sprintf("%+d", sv.Int())
}

func stringerIntPercentageSigned(sv d2stats.StatValue) string {
	return fmt.Sprintf("%+d%%", sv.Int())
}

func stringerIntPercentageUnsigned(sv d2stats.StatValue) string {
	return fmt.Sprintf("%d%%", sv.Int())
}

func stringerClassAllSkills(sv d2stats.StatValue) string {
	heroIndex := sv.Int()

	heroMap := getHeroMap()
	classRecord := d2datadict.CharStats[heroMap[heroIndex]]

	return d2common.TranslateString(classRecord.SkillStrAll)
}

func stringerClassOnly(sv d2stats.StatValue) string {
	heroMap := getHeroMap()
	heroIndex := sv.Int()
	classRecord := d2datadict.CharStats[heroMap[heroIndex]]
	classOnlyKey := classRecord.SkillStrClassOnly

	return d2common.TranslateString(classOnlyKey)
}

func stringerSkillName(sv d2stats.StatValue) string {
	skillRecord := d2datadict.SkillDetails[sv.Int()]
	return skillRecord.Skill
}

func stringerMonsterName(sv d2stats.StatValue) string {
	for key := range d2datadict.MonStats {
		if d2datadict.MonStats[key].Id == sv.Int() {
			return d2datadict.MonStats[key].NameString
		}
	}

	return monsterNotFound
}
