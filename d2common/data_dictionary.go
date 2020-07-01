package d2common

import (
	"log"
	"strconv"
	"strings"
)

// DataDictionary represents a data file (Excel)
type DataDictionary struct {
	FieldNameLookup map[string]int
	Data            [][]string
}

// LoadDataDictionary loads the contents of a spreadsheet style txt file
func LoadDataDictionary(text string) *DataDictionary {
	result := &DataDictionary{}
	lines := strings.Split(text, "\r\n")
	fileNames := strings.Split(lines[0], "\t")
	result.FieldNameLookup = make(map[string]int)

	for i, fieldName := range fileNames {
		result.FieldNameLookup[fieldName] = i
	}

	result.Data = make([][]string, len(lines)-2)

	for i, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			continue
		}

		values := strings.Split(line, "\t")
		if len(values) != len(result.FieldNameLookup) {
			continue
		}

		result.Data[i] = values
	}

	return result
}

// GetString gets a string from the given column and row
func (v *DataDictionary) GetString(fieldName string, index int) string {
	return v.Data[index][v.FieldNameLookup[fieldName]]
}

// GetNumber gets a number for the given column and row
func (v *DataDictionary) GetNumber(fieldName string, index int) int {
	str := v.GetString(fieldName, index)
	str = EmptyToZero(AsterToEmpty(str))

	result, err := strconv.Atoi(str)
	if err != nil {
		log.Panic(err)
	}

	return result
}

// GetDelimitedList splits a delimited list from the given column and row
func (v *DataDictionary) GetDelimitedList(fieldName string, index int) []string {
	unsplit := v.GetString(fieldName, index)

	// Comma delimited fields are quoted, not terribly pretty to fix that here but...
	unsplit = strings.TrimRight(unsplit, "\"")
	unsplit = strings.TrimLeft(unsplit, "\"")

	return strings.Split(unsplit, ",")
}

// GetBool gets a bool value for the given column and row
func (v *DataDictionary) GetBool(fieldName string, index int) bool {
	n := v.GetNumber(fieldName, index)
	if n > 1 {
		log.Panic("GetBool on non-bool field")
	}

	return n == 1
}
