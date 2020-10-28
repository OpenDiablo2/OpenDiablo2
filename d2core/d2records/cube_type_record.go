package d2records

// CubeTypes is a map of CubeTypeRecords
type CubeTypes map[string]*CubeTypeRecord

// CubeTypeRecord is a name and 3-character token for cube item types
type CubeTypeRecord struct {
	Name  string
	Token string
}
