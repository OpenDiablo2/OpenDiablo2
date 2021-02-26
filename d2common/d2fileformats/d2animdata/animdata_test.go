package d2animdata

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func exampleData() *AnimationData {
	testEntries := []*AnimationDataRecord{
		&AnimationDataRecord{
			"TST",
			5, 8,
			map[int]AnimationEvent{
				1: AnimationEventNone,
			},
		},
		&AnimationDataRecord{
			"TST",
			8, 3,
			map[int]AnimationEvent{
				2: AnimationEventNone,
			},
		},
	}

	testEntries2 := []*AnimationDataRecord{
		&AnimationDataRecord{
			"TTT",
			7, 8,
			map[int]AnimationEvent{
				1: AnimationEventNone,
			},
		},
		&AnimationDataRecord{
			"TTT",
			8, 9,
			map[int]AnimationEvent{
				8: AnimationEventNone,
			},
		},
	}

	result := &AnimationData{
		entries: map[string][]*AnimationDataRecord{
			"TST": testEntries,
			"TTT": testEntries2,
		},
	}

	return result
}

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

func TestAnimationDataRecord_Marshal(t *testing.T) {
	file, fileErr := os.Open("testdata/AnimData.d2")
	if fileErr != nil {
		t.Error("cannot open test data file")
		return
	}

	data := make([]byte, 0)
	buf := make([]byte, 16)

	for {
		numRead, err := file.Read(buf)

		data = append(data, buf[:numRead]...)

		if err != nil {
			break
		}
	}

	ad, err := Load(data)
	if err != nil {
		t.Error(err)
	}
	newData := ad.Marshal()
	newAd, err := Load(newData)
	if err != nil {
		t.Error(err)
	}

	keys1 := make([]string, 0)
	for i := range ad.entries {
		keys1 = append(keys1, i)
	}
	keys2 := make([]string, 0)
	for i := range newAd.entries {
		keys2 = append(keys2, i)
	}

	if len(keys1) != len(keys2) {
		t.Fatalf("unexpected length of keys in first and second dict: %d, %d", len(keys1), len(keys2))
	}

	fmt.Println(len(ad.entries["TST"]))
	for key := range newAd.entries {
		for n, i := range newAd.entries[key] {
			fmt.Println(i.speed, ad.entries[key][n].speed)
			if i.speed != ad.entries[key][n].speed {
				t.Fatal("unexpected record set")
			}
		}
	}
}
