package d2records

// Gems stores all of the GemsRecords
type Gems map[string]*GemsRecord

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
