package d2records

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadMagicPrefix loads MagicPrefix.txt
func magicPrefixLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	superType := d2enum.ItemAffixPrefix

	subType := d2enum.ItemAffixMagic

	affixes, groups, err := loadAffixDictionary(r, d, superType, subType)
	if err != nil {
		return err
	}

	r.Item.Magic.Prefix = affixes
	r.Item.MagicPrefixGroups = groups

	return nil
}

// LoadMagicSuffix loads MagicSuffix.txt
func magicSuffixLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	superType := d2enum.ItemAffixSuffix

	subType := d2enum.ItemAffixMagic

	affixes, groups, err := loadAffixDictionary(r, d, superType, subType)
	if err != nil {
		return err
	}

	r.Item.Magic.Suffix = affixes
	r.Item.MagicSuffixGroups = groups

	return nil
}

func getAffixString(t1 d2enum.ItemAffixSuperType, t2 d2enum.ItemAffixSubType) string {
	var name = ""

	if t2 == d2enum.ItemAffixMagic {
		name = "Magic"
	}

	switch t1 {
	case d2enum.ItemAffixPrefix:
		name += "Prefix"
	case d2enum.ItemAffixSuffix:
		name += "Suffix"
	}

	return name
}

func loadAffixDictionary(
	r *RecordManager,
	d *d2txt.DataDictionary,
	superType d2enum.ItemAffixSuperType,
	subType d2enum.ItemAffixSubType,
) (map[string]*ItemAffixCommonRecord, ItemAffixGroups, error) {
	records, groups, err := createItemAffixRecords(d, superType, subType)
	if err != nil {
		return nil, nil, err
	}

	name := getAffixString(superType, subType)
	r.Logger.Infof("Loaded %d %s records", len(records), name)

	return records, groups, nil
}

func createItemAffixRecords(
	d *d2txt.DataDictionary,
	superType d2enum.ItemAffixSuperType,
	subType d2enum.ItemAffixSubType,
) (map[string]*ItemAffixCommonRecord, ItemAffixGroups, error) {
	records := make(map[string]*ItemAffixCommonRecord)
	groups := make(ItemAffixGroups)

	for d.Next() {
		affix := &ItemAffixCommonRecord{
			Name:           d.String("Name"),
			Version:        d.Number("version"),
			Type:           subType,
			IsPrefix:       superType == d2enum.ItemAffixPrefix,
			IsSuffix:       superType == d2enum.ItemAffixSuffix,
			Spawnable:      d.Bool("spawnable"),
			Rare:           d.Bool("rare"),
			Level:          d.Number("level"),
			MaxLevel:       d.Number("maxlevel"),
			LevelReq:       d.Number("levelreq"),
			Class:          d.String("classspecific"),
			ClassLevelReq:  d.Number("classlevelreq"),
			Frequency:      d.Number("frequency"),
			GroupID:        d.Number("group"),
			Transform:      d.Bool("transform"),
			TransformColor: d.String("transformcolor"),
			PriceAdd:       d.Number("add"),
			PriceScale:     d.Number("multiply"),
		}

		// modifiers (Code references with parameters to be eval'd)
		for i := 1; i <= 3; i++ {
			codeKey := fmt.Sprintf("mod%dcode", i)
			paramKey := fmt.Sprintf("mod%dparam", i)
			minKey := fmt.Sprintf("mod%dmin", i)
			maxKey := fmt.Sprintf("mod%dmax", i)
			modifier := &ItemAffixCommonModifier{
				Code:      d.String(codeKey),
				Parameter: d.String(paramKey),
				Min:       d.Number(minKey),
				Max:       d.Number(maxKey),
			}
			affix.Modifiers = append(affix.Modifiers, modifier)
		}

		// items to include for spawning
		for i := 1; i <= 7; i++ {
			itemKey := fmt.Sprintf("itype%d", i)
			itemToken := d.String(itemKey)
			affix.ItemInclude = append(affix.ItemInclude, itemToken)
		}

		// items to exclude for spawning
		for i := 1; i <= 7; i++ {
			itemKey := fmt.Sprintf("etype%d", i)
			itemToken := d.String(itemKey)
			affix.ItemExclude = append(affix.ItemExclude, itemToken)
		}

		if _, found := groups[affix.GroupID]; !found {
			ItemAffixGroup := &ItemAffixCommonGroup{}
			ItemAffixGroup.ID = affix.GroupID
			groups[affix.GroupID] = ItemAffixGroup
		}

		group := groups[affix.GroupID]
		group.AddMember(affix)

		records[affix.Name] = affix
	}

	if d.Err != nil {
		return nil, nil, d.Err
	}

	return records, groups, nil
}
