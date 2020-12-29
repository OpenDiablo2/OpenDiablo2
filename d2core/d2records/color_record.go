package d2records

// Colors is a map of ColorRecords
type Colors map[string]*ColorRecord

// ColorRecord is a representation of a color transform
type ColorRecord struct {
	TransformColor string
	Code           string
}
