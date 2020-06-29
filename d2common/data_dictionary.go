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
		if len(strings.TrimSpace(line)) == 0 {
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

func (v *DataDictionary) GetString(fieldName string, index int) string {
	return v.Data[index][v.FieldNameLookup[fieldName]]
}

func (v *DataDictionary) GetNumber(fieldName string, index int) int {
	str := v.GetString(fieldName, index)
	str = EmptyToZero(AsterToEmpty(str))
	result, err := strconv.Atoi(str)
	if err != nil {
		log.Panic(err)
	}
	return result
}

func (v *DataDictionary) GetDelimitedList(fieldName string, index int) []string {
	unsplit := v.GetString(fieldName, index)

	// Commo delimited fields are quoted, not terribly pretty to do it here but...
	s := []byte(unsplit)
	j := 0

	for i := range s {
		if s[i] != '"' {
			s[j] = s[i]
			j++
		}
	}

	return strings.Split(string(s), ",")
}

func (v *DataDictionary) GetBool(fieldName string, index int) bool {
	n := v.GetNumber(fieldName, index)
	if n > 1 {
		log.Panic("GetBool on non-bool field")
	}
	return n == 1
}
