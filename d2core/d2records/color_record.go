package d2records

type Colors map[string]*ColorRecord

type ColorRecord struct {
	TransformColor string
	Code           string
}
