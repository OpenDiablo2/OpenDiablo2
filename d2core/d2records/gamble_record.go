package d2records

type Gamble map[string]*GambleRecord

type GambleRecord struct {
	Name string
	Code string
}
