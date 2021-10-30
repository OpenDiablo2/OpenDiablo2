package d2records

// Gems stores all of the GemRecords
type Gems map[string]*GemRecord

// GemRecord is a representation of a single row of gems.txt
// it describes the properties of socketable items
type GemRecord struct {
	WeaponMod3Code  string
	Letter          string
	ShieldMod3Code  string
	Code            string
	ShieldMod2Code  string
	WeaponMod1Code  string
	ShieldMod1Code  string
	HelmMod2Code    string
	Name            string
	WeaponMod2Code  string
	HelmMod1Code    string
	HelmMod3Code    string
	WeaponMod3Min   int
	WeaponMod2Max   int
	WeaponMod3Param int
	WeaponMod2Min   int
	WeaponMod3Max   int
	WeaponMod2Param int
	HelmMod1Param   int
	HelmMod1Min     int
	WeaponMod1Max   int
	WeaponMod1Min   int
	HelmMod2Param   int
	HelmMod2Min     int
	HelmMod2Max     int
	HelmMod1Max     int
	HelmMod3Param   int
	HelmMod3Min     int
	HelmMod3Max     int
	WeaponMod1Param int
	ShieldMod1Param int
	ShieldMod1Min   int
	ShieldMod1Max   int
	Nummods         int
	ShieldMod2Param int
	ShieldMod2Min   int
	ShieldMod2Max   int
	Transform       int
	ShieldMod3Param int
	ShieldMod3Min   int
	ShieldMod3Max   int
}
