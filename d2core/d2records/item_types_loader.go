package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadItemTypes loads ItemType records
func itemTypesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(ItemTypes)

	charCodeMap := map[string]d2enum.Hero{
		"ama": d2enum.HeroAmazon,
		"ass": d2enum.HeroAssassin,
		"bar": d2enum.HeroBarbarian,
		"dru": d2enum.HeroDruid,
		"nec": d2enum.HeroNecromancer,
		"pal": d2enum.HeroPaladin,
		"sor": d2enum.HeroSorceress,
	}

	for d.Next() {
		if d.String("*eol") == "" {
			continue
		}

		itemType := &ItemTypeRecord{
			Name:          d.String("ItemType"),
			Code:          d.String("Code"),
			Equiv1:        d.String("Equiv1"),
			Equiv2:        d.String("Equiv2"),
			Repair:        d.Number("Repair") > 0,
			Body:          d.Number("Body") > 0,
			BodyLoc1:      d.Number("BodyLoc1"),
			BodyLoc2:      d.Number("BodyLoc2"),
			Shoots:        d.String("Shoots"),
			Quiver:        d.String("Quiver"),
			Throwable:     d.Number("Throwable") > 0,
			Reload:        d.Number("Reload") > 0,
			ReEquip:       d.Number("ReEquip") > 0,
			AutoStack:     d.Number("AutoStack") > 0,
			Magic:         d.Number("Magic") > 0,
			Rare:          d.Number("Rare") > 0,
			Normal:        d.Number("Normal") > 0,
			Charm:         d.Number("Charm") > 0,
			Gem:           d.Number("Gem") > 0,
			Beltable:      d.Number("Beltable") > 0,
			MaxSock1:      d.Number("MaxSock1"),
			MaxSock25:     d.Number("MaxSock25"),
			MaxSock40:     d.Number("MaxSock40"),
			TreasureClass: d.Number("TreasureClass"),
			Rarity:        d.Number("Rarity"),
			StaffMods:     charCodeMap[d.String("StaffMods")],
			CostFormula:   d.Number("CostFormula"),
			Class:         charCodeMap[d.String("Class")],
			VarInvGfx:     d.Number("VarInvGfx"),
			InvGfx1:       d.String("InvGfx1"),
			InvGfx2:       d.String("InvGfx2"),
			InvGfx3:       d.String("InvGfx3"),
			InvGfx4:       d.String("InvGfx4"),
			InvGfx5:       d.String("InvGfx5"),
			InvGfx6:       d.String("InvGfx6"),
			StorePage:     d.String("StorePage"),
		}

		records[itemType.Code] = itemType
	}

	equivMap := LoadItemEquivalencies(r.Item.All, records)

	for idx := range records {
		records[idx].EquivalentItems = equivMap[records[idx].Code]
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d ItemType records", len(records))

	r.Item.Types = records
	r.Item.Equivalency = equivMap

	return nil
}

// LoadItemEquivalencies loads a map of ItemType string codes to slices of ItemCommonRecord pointers
func LoadItemEquivalencies(allItems CommonItems, allTypes ItemTypes) ItemEquivalenceMap {
	equivMap := make(ItemEquivalenceMap)

	for typeCode := range allTypes {
		code := []string{
			typeCode,
			allTypes[typeCode].Equiv1,
			allTypes[typeCode].Equiv2,
		}

		for _, str := range code {
			if str == "" {
				continue
			}

			if equivMap[str] == nil {
				equivMap[str] = make([]*ItemCommonRecord, 0)
			}
		}
	}

	for icrCode := range allItems {
		commonItem := allItems[icrCode]
		updateEquivalencies(allTypes, equivMap, commonItem, allTypes[commonItem.Type], nil)

		if commonItem.Type2 != "" { // some items (like gems) have a secondary type
			updateEquivalencies(allTypes, equivMap, commonItem, allTypes[commonItem.Type2], nil)
		}
	}

	return equivMap
}

func updateEquivalencies(
	allTypes ItemTypes,
	equivMap ItemEquivalenceMap,
	icr *ItemCommonRecord,
	itemType *ItemTypeRecord,
	checked []string,
) {
	if itemType.Code == "" {
		return
	}

	if checked == nil {
		checked = make([]string, 0)
	}

	checked = append(checked, itemType.Code)

	if !itemEquivPresent(icr, equivMap[itemType.Code]) {
		equivMap[itemType.Code] = append(equivMap[itemType.Code], icr)
	}

	if itemType.Equiv1 != "" {
		updateEquivalencies(allTypes, equivMap, icr, allTypes[itemType.Equiv1], checked)
	}

	if itemType.Equiv2 != "" {
		updateEquivalencies(allTypes, equivMap, icr, allTypes[itemType.Equiv2], checked)
	}
}

func itemEquivPresent(icr *ItemCommonRecord, list []*ItemCommonRecord) bool {
	for idx := range list {
		if list[idx] == icr {
			return true
		}
	}

	return false
}
