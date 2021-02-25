package d2tbl

import (
	"testing"
)

func exampleData() *TextDictionary {
	result := &TextDictionary{
		crcBytes: make([]byte, crcByteCount),
	}
}
