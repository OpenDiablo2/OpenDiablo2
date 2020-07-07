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
	Gems = make(map[string]*GemsRecord)

	d := d2common.LoadDataDictionary(file)
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
		Gems[gem.Name] = gem
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Gems records", len(Gems))
}
