package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadGems loads gem records into a map[string]*GemsRecord
func gemsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(Gems)

	for d.Next() {
		gem := &GemsRecord{
			Name:            d.String("name"),
			Letter:          d.String("letter"),
			Transform:       d.Number("transform"),
			Code:            d.String("code"),
			Nummods:         d.Number("nummods"),
			WeaponMod1Code:  d.String("weaponMod1Code"),
			WeaponMod1Param: d.Number("weaponMod1Param"),
			WeaponMod1Min:   d.Number("weaponMod1Min"),
			WeaponMod1Max:   d.Number("weaponMod1Max"),
			WeaponMod2Code:  d.String("weaponMod2Code"),
			WeaponMod2Param: d.Number("weaponMod2Param"),
			WeaponMod2Min:   d.Number("weaponMod2Min"),
			WeaponMod2Max:   d.Number("weaponMod2Max"),
			WeaponMod3Code:  d.String("weaponMod3Code"),
			WeaponMod3Param: d.Number("weaponMod3Param"),
			WeaponMod3Min:   d.Number("weaponMod3Min"),
			WeaponMod3Max:   d.Number("weaponMod3Max"),
			HelmMod1Code:    d.String("helmMod1Code"),
			HelmMod1Param:   d.Number("helmMod1Param"),
			HelmMod1Min:     d.Number("helmMod1Min"),
			HelmMod1Max:     d.Number("helmMod1Max"),
			HelmMod2Code:    d.String("helmMod2Code"),
			HelmMod2Param:   d.Number("helmMod2Param"),
			HelmMod2Min:     d.Number("helmMod2Min"),
			HelmMod2Max:     d.Number("helmMod2Max"),
			HelmMod3Code:    d.String("helmMod3Code"),
			HelmMod3Param:   d.Number("helmMod3Param"),
			HelmMod3Min:     d.Number("helmMod3Min"),
			HelmMod3Max:     d.Number("helmMod3Max"),
			ShieldMod1Code:  d.String("shieldMod1Code"),
			ShieldMod1Param: d.Number("shieldMod1Param"),
			ShieldMod1Min:   d.Number("shieldMod1Min"),
			ShieldMod1Max:   d.Number("shieldMod1Max"),
			ShieldMod2Code:  d.String("shieldMod2Code"),
			ShieldMod2Param: d.Number("shieldMod2Param"),
			ShieldMod2Min:   d.Number("shieldMod2Min"),
			ShieldMod2Max:   d.Number("shieldMod2Max"),
			ShieldMod3Code:  d.String("shieldMod3Code"),
			ShieldMod3Param: d.Number("shieldMod3Param"),
			ShieldMod3Min:   d.Number("shieldMod3Min"),
			ShieldMod3Max:   d.Number("shieldMod3Max"),
		}

		records[gem.Name] = gem
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d Gems records", len(records))

	r.Item.Gems = records

	return nil
}
