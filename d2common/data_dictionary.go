package d2common

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
)

// DataDictionary represents a data file (Excel)
type DataDictionary struct {
	lookup map[string]int
	r      *csv.Reader
	record []string
	Err    error
}

// LoadDataDictionary loads the contents of a spreadsheet style txt file
func LoadDataDictionary(buf []byte) *DataDictionary {
	cr := csv.NewReader(bytes.NewReader(buf))
	cr.Comma = '\t'
	cr.ReuseRecord = true

	fieldNames, err := cr.Read()
	if err != nil {
		panic(err)
	}

	data := &DataDictionary{
		lookup: make(map[string]int, len(fieldNames)),
		r:      cr,
	}

	for i, name := range fieldNames {
		data.lookup[name] = i
	}

	return data
}

// Next reads the next row, skips Expansion lines or
// returns false when the end of a file is reached or an error occurred
func (d *DataDictionary) Next() bool {
	var err error
	d.record, err = d.r.Read()

	if err == io.EOF {
		return false
	} else if err != nil {
		d.Err = err
		return false
	}

	if d.record[0] == "Expansion" {
		return d.Next()
	}

	return true
}

// String gets a string from the given column
func (d *DataDictionary) String(field string) string {
	return d.record[d.lookup[field]]
}

// Number gets a number for the given column
func (d *DataDictionary) Number(field string) int {
	n, err := strconv.Atoi(d.String(field))
	if err != nil {
		return 0
	}

	return n
}

// List splits a delimited list from the given column
func (d *DataDictionary) List(field string) []string {
	str := d.String(field)
	return strings.Split(str, ",")
}

// Bool gets a bool value for the given column
func (d *DataDictionary) Bool(field string) bool {
	n := d.Number(field)
	if n > 1 {
		log.Panic("Bool on non-bool field ", field)
	}

	return n == 1
}

// PopulateStruct uses reflection to fill struct fields based on their names and types.
// In this implementation, the struct field name must match the data field name exactly. Missing fields will be 0 and
// unexpected fields will throw an error.
func (d *DataDictionary) PopulateStruct(values reflect.Value) error {
	// Get the struct's Type object
	structType := values.Type()

	// Create a map of struct fields to reflect Value objects
	valueMap := make(map[string]reflect.Value)

	for i := 0; i < structType.NumField(); i++ {
		v := values.Field(i)
		fieldName := structType.Field(i).Name

		valueMap[fieldName] = v
	}

	// Iterate over the data field names, checking they exist in the struct and populating the struct fields.
	for k, _ := range d.lookup {
		// Always make the first character upper case because struct fields must be exported for reflection.
		value, exists := valueMap[strings.Title(k)]
		if !exists {
			return fmt.Errorf("data key %s is invalid for type %s", k, structType.Name())
		}

		switch value.Kind() {
		case reflect.String:
			value.SetString(d.String(k))

		case reflect.Int:
			value.SetInt(int64(d.Number(k)))
		}
	}

	return nil
}
