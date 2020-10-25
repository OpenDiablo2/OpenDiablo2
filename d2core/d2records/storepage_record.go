package d2records

// StorePages is a map of all store page records
type StorePages map[string]*StorePageRecord

// StorePageRecord represent a row in the storepage.txt file
type StorePageRecord struct {
	StorePage string
	Code      string
}
