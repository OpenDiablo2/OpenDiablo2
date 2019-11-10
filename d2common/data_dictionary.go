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
	result.Data = make([][]string, len(lines)-1)
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
	result, err := strconv.Atoi(v.GetString(fieldName, index))
	if err != nil {
		log.Panic(err)
	}
	return result
}
