package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

//MonsterSequenceRecord contains a record for a monster sequence
//Composed of multiple lines from monseq.txt with the same name in the first column.
//Information gathered from [https://d2mods.info/forum/kb/viewarticle?a=395]
type MonsterSequenceRecord struct {

	//The sequence name, refered to by monstats.txt
	Name string

	//The frames of this sequence
	Frames []*MonsterSequenceFrame
}

//MonsterSequenceFrame represents a single frame of a monster sequence
type MonsterSequenceFrame struct {
	//The animation mode for this frame (refers to MonMode.txt)
	Mode string

	//The frame of the animation mode used for this frame of the sequence
	Frame int

	//The direction of the frame, enumerated by d2enum.AnimationFrameDirection
	Direction int

	//Event trigerred by this frame
	Event int
}

//MonsterSequences contains the MonsterSequenceRecords
var MonsterSequences map[string]*MonsterSequenceRecord

//LoadMonsterSequences loads the MonsterSequenceRecords into MonsterSequences
func LoadMonsterSequences(file []byte) {
	MonsterSequences = make(map[string]*MonsterSequenceRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {

		name := d.String("sequence")
		if _, ok := MonsterSequences[name]; !ok {
			record := &MonsterSequenceRecord{
				Name:   name,
				Frames: make([]*MonsterSequenceFrame, 0),
			}
			MonsterSequences[name] = record
		}
		MonsterSequences[name].Frames = append(MonsterSequences[name].Frames, &MonsterSequenceFrame{
			Mode:      d.String("mode"),
			Frame:     d.Number("frame"),
			Direction: d.Number("dir"),
			Event:     d.Number("event"),
		})
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonsterSequence records", len(MonsterSequences))
}
