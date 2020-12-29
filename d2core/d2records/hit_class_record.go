package d2records

// HitClasses is a map of HitClassRecords
type HitClasses map[string]*HitClassRecord

// HitClassRecord is used for changing character animation modes.
type HitClassRecord struct {
	Name  string
	Token string
}
