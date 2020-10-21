package d2records

type ObjectModes map[string]*ObjectModeRecord

type ObjectModeRecord struct {
	Name  string
	Token string
}
