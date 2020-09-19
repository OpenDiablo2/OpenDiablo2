package d2records

// PlayerClasses stores the PlayerClassRecords
type PlayerClasses map[string]*PlayerClassRecord

// PlayerClassRecord represents a single line from PlayerClass.txt
// Lookup table for class codes
type PlayerClassRecord struct {
	// Name of the player class
	Name string

	// Code for the player class
	Code string
}
