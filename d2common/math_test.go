package d2common

import (
	"testing"
)

type TestRecord struct {
	source, adjust, max, expectedResult, expectedRemain float64
}

func TestSomething(t *testing.T) {
	var testValues = []TestRecord{
		{100, 10, 100.2, 100.2, 9.8},
	}
	for _, test := range testValues {
		res, remain := AdjustWithRemainder(test.source, test.adjust, test.max)
		if res != test.expectedResult {
			t.Errorf("Expected result of %f but got %f", test.expectedResult, res)
		}
		if remain != test.expectedRemain {
			t.Errorf("Expected result of %f but got %f", test.expectedRemain, remain)
		}
	}

}
