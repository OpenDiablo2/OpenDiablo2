package d2animdata

import (
	"log"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	testFile, fileErr := os.Open("testdata/AnimData.d2")
	if fileErr != nil {
		t.Error("cannot open test data file")
		return
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

	err := testFile.Close()
	if err != nil {
		t.Fail()
		log.Print(err)
	}
}

func TestLoad_BadData(t *testing.T) {
	testFile, fileErr := os.Open("testdata/BadData.d2")
	if fileErr != nil {
		t.Error("cannot open test data file")
		return
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

	err := testFile.Close()
	if err != nil {
		t.Fail()
		log.Print(err)
	}
}

func TestAnimationData_GetRecordNames(t *testing.T) {
	animdata := &AnimationData{
		hashTable: hashTable{},
		blocks:    [256]*block{},
		entries: map[string][]*AnimationDataRecord{
			"a": {{}},
			"b": {{}},
			"c": {{}},
		},
	}

	names := animdata.GetRecordNames()
	if len(names) != 3 {
		t.Error("record name count mismatch")
	}
}

func TestAnimationData_GetRecords(t *testing.T) {
	animdata := &AnimationData{
		hashTable: hashTable{},
		blocks:    [256]*block{},
		entries: map[string][]*AnimationDataRecord{
			"a": {
				{name: "a", speed: 1, framesPerDirection: 1},
				{name: "a", speed: 2, framesPerDirection: 2},
				{name: "a", speed: 3, framesPerDirection: 3},
			},
		},
	}

	if len(animdata.GetRecords("a")) != 3 {
		t.Error("record count is incorrect")
	}

	if len(animdata.GetRecords("b")) > 0 {
		t.Error("retrieved records for unknown record name")
	}
}

func TestAnimationData_GetRecord(t *testing.T) {
	animdata := &AnimationData{
		hashTable: hashTable{},
		blocks:    [256]*block{},
		entries: map[string][]*AnimationDataRecord{
			"a": {
				{name: "a", speed: 1, framesPerDirection: 1},
				{name: "a", speed: 2, framesPerDirection: 2},
				{name: "a", speed: 3, framesPerDirection: 3},
			},
		},
	}

	record := animdata.GetRecord("a")
	if record.speed != 3 {
		t.Error("record returned is incorrect")
	}
}

func TestAnimationDataRecord_FPS(t *testing.T) {
	record := &AnimationDataRecord{}

	var fps float64

	record.speed = 256
	fps = record.FPS()

	if fps != float64(speedBaseFPS) {
		t.Error("incorrect fps")
	}

	record.speed = 512
	fps = record.FPS()

	if fps != float64(speedBaseFPS)*2 {
		t.Error("incorrect fps")
	}

	record.speed = 128
	fps = record.FPS()

	if fps != float64(speedBaseFPS)/2 {
		t.Error("incorrect fps")
	}
}
