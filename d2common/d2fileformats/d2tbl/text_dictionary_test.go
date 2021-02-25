package d2tbl

import (
	"fmt"

	"testing"
)

func exampleData() *TextDictionary {
	result := &TextDictionary{
		"abc":       "def",
		"someStr":   "Some long string",
		"lolstring": "lol",
	}

	return result
}

func TestTBL_Marshal(t *testing.T) {
	tbl := exampleData()
	data := tbl.Marshal()
	newTbl, err := LoadTextDictionary(data)
	if err != nil {
		t.Error(err)
	}

	_, ok := newTbl["lolstring"]
	if !ok {
		t.Fatal("no string found")
	}
}
