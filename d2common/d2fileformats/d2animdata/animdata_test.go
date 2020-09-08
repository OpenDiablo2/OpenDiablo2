package d2animdata

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	testFile, fileErr := os.Open("testdata/AnimData.d2")
	if fileErr != nil {
		t.Error("cannot open test data file")
	}

	data := make([]byte, 0)
	buf := make([]byte, 16)

	for {
		numRead, err := testFile.Read(buf)

		data = append(data, buf[:numRead]...)

		if err != nil {
			break
		}
	}

	_, loadErr := Load(data)
	if loadErr != nil {
		t.Error(loadErr)
	}

	testFile.Close()
}

func TestLoad_BadData(t *testing.T) {
	testFile, fileErr := os.Open("testdata/BadData.d2")
	if fileErr != nil {
		t.Error("cannot open test data file")
	}

	data := make([]byte, 0)
	buf := make([]byte, 16)

	for {
		numRead, err := testFile.Read(buf)

		data = append(data, buf[:numRead]...)

		if err != nil {
			break
		}
	}

	_, loadErr := Load(data)
	if loadErr == nil {
		t.Error("bad data file should not be parsed")
	}

	testFile.Close()
}
