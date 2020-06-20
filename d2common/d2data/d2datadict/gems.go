package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// GemsRecord is a representation of a single row of gems.txt
// it describes the properties of socketable items
type GemsRecord struct {
	Name            string
	Letter          string
	Transform       int
	Code            string
	Nummods         int
	WeaponMod1Code  string
	WeaponMod1Param int
	WeaponMod1Min   int
	WeaponMod1Max   int
	WeaponMod2Code  string
	WeaponMod2Param int
	WeaponMod2Min   int
	WeaponMod2Max   int
	WeaponMod3Code  string
	WeaponMod3Param int
	WeaponMod3Min   int
	WeaponMod3Max   int
	HelmMod1Code    string
	HelmMod1Param   int
	HelmMod1Min     int
	HelmMod1Max     int
	HelmMod2Code    string
	HelmMod2Param   int
	HelmMod2Min     int
	HelmMod2Max     int
	HelmMod3Code    string
	HelmMod3Param   int
	HelmMod3Min     int
	HelmMod3Max     int
	ShieldMod1Code  string
	ShieldMod1Param int
	ShieldMod1Min   int
	ShieldMod1Max   int
	ShieldMod2Code  string
	ShieldMod2Param int
	ShieldMod2Min   int
	ShieldMod2Max   int
	ShieldMod3Code  string
	ShieldMod3Param int
	ShieldMod3Min   int
	ShieldMod3Max   int
}

// Gems stores all of the GemsRecords
var Gems map[string]*GemsRecord //nolint:gochecknoglobals // Currently global by design, only written once

// LoadGems loads gem records into a map[string]*GemsRecord
func LoadGems(file []byte) {
	d := d2common.LoadDataDictionary(string(file))

	Gems = make(map[string]*GemsRecord, len(d.Data))

	for idx := range d.Data {
		if d.GetString("name", idx) != expansion {
			/*
				"Expansion" is the only field in line 36 of /data/global/excel/gems.txt and is only used to visually
				separate base-game gems and expansion runes.
			*/
			gem := &GemsRecord{
				Name:            d.GetString("name", idx),
				Letter:          d.GetString("letter", idx),
				Transform:       d.GetNumber("transform", idx),
				Code:            d.GetString("code", idx),
				Nummods:         d.GetNumber("nummods", idx),
				WeaponMod1Code:  d.GetString("weaponMod1Code", idx),
				WeaponMod1Param: d.GetNumber("weaponMod1Param", idx),
				WeaponMod1Min:   d.GetNumber("weaponMod1Min", idx),
				WeaponMod1Max:   d.GetNumber("weaponMod1Max", idx),
				WeaponMod2Code:  d.GetString("weaponMod2Code", idx),
				WeaponMod2Param: d.GetNumber("weaponMod2Param", idx),
				WeaponMod2Min:   d.GetNumber("weaponMod2Min", idx),
				WeaponMod2Max:   d.GetNumber("weaponMod2Max", idx),
				WeaponMod3Code:  d.GetString("weaponMod3Code", idx),
				WeaponMod3Param: d.GetNumber("weaponMod3Param", idx),
				WeaponMod3Min:   d.GetNumber("weaponMod3Min", idx),
				WeaponMod3Max:   d.GetNumber("weaponMod3Max", idx),
				HelmMod1Code:    d.GetString("helmMod1Code", idx),
				HelmMod1Param:   d.GetNumber("helmMod1Param", idx),
				HelmMod1Min:     d.GetNumber("helmMod1Min", idx),
				HelmMod1Max:     d.GetNumber("helmMod1Max", idx),
				HelmMod2Code:    d.GetString("helmMod2Code", idx),
				HelmMod2Param:   d.GetNumber("helmMod2Param", idx),
				HelmMod2Min:     d.GetNumber("helmMod2Min", idx),
				HelmMod2Max:     d.GetNumber("helmMod2Max", idx),
				HelmMod3Code:    d.GetString("helmMod3Code", idx),
				HelmMod3Param:   d.GetNumber("helmMod3Param", idx),
				HelmMod3Min:     d.GetNumber("helmMod3Min", idx),
				HelmMod3Max:     d.GetNumber("helmMod3Max", idx),
				ShieldMod1Code:  d.GetString("shieldMod1Code", idx),
				ShieldMod1Param: d.GetNumber("shieldMod1Param", idx),
				ShieldMod1Min:   d.GetNumber("shieldMod1Min", idx),
				ShieldMod1Max:   d.GetNumber("shieldMod1Max", idx),
				ShieldMod2Code:  d.GetString("shieldMod2Code", idx),
				ShieldMod2Param: d.GetNumber("shieldMod2Param", idx),
				ShieldMod2Min:   d.GetNumber("shieldMod2Min", idx),
				ShieldMod2Max:   d.GetNumber("shieldMod2Max", idx),
				ShieldMod3Code:  d.GetString("shieldMod3Code", idx),
				ShieldMod3Param: d.GetNumber("shieldMod3Param", idx),
				ShieldMod3Min:   d.GetNumber("shieldMod3Min", idx),
				ShieldMod3Max:   d.GetNumber("shieldMod3Max", idx),
			}
			Gems[gem.Name] = gem
		}
	}

	log.Printf("Loaded %d Gems records", len(Gems))
}
