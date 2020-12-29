package d2records

// CubeModifiers is a map of CubeModifierRecords
type CubeModifiers map[string]*CubeModifierRecord

// CubeModifierRecord is a name and 3-character token for cube modifier codes and gem types
type CubeModifierRecord struct {
	Name  string
	Token string
}
