package d2records

// StorePages struct contains all store page records
type StorePages map[string]*StorePageRecord

// StorePageRecords represent a row in the storepage.txt file
type StorePageRecord struct {
	StorePage string
	Code      string
}
