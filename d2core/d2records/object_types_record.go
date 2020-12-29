package d2records

// ObjectTypes contains the name and token for objects
type ObjectTypes []ObjectTypeRecord

// ObjectTypeRecord is a representation of a row from objtype.txt
type ObjectTypeRecord struct {
	Name  string
	Token string
}
