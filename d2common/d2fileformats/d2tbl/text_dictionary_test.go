package d2tbl

import (
	"fmt"

	"testing"
)

func exampleData() *TextDictionary {
	result := &TextDictionary{
		"abc":        "def",
		"someStr":    "Some long string",
		"teststring": "TeStxwsas123 long strin122*8:wq",
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

	fmt.Println(newTbl)

	for key, value := range *tbl {
		newValue, ok := newTbl[key]

		if !ok {
			t.Fatalf("string %s wasn't encoded to table", key)
		}

		if newValue != value {
			t.Fatal("unexpected value set")
		}
	}
}
