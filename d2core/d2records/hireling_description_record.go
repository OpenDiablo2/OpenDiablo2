package d2records

// HirelingDescriptions is a lookup table for hireling subtype codes
type HirelingDescriptions map[string]*HirelingDescriptionRecord

// HirelingDescriptionRecord represents is a hireling subtype
type HirelingDescriptionRecord struct {
	Name  string
	Token string
}
