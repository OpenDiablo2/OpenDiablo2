package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadMonsterSequences loads the MonsterSequenceRecords into MonsterSequences
func monsterSequencesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(MonsterSequences)

	for d.Next() {
		name := d.String("sequence")

		if _, ok := records[name]; !ok {
			record := &MonsterSequenceRecord{
				Name:   name,
				Frames: make([]*MonsterSequenceFrame, 0),
			}
			records[name] = record
		}

		records[name].Frames = append(records[name].Frames, &MonsterSequenceFrame{
			Mode:      d.String("mode"),
			Frame:     d.Number("frame"),
			Direction: d.Number("dir"),
			Event:     d.Number("event"),
		})
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d MonsterSequence records", len(records))

	r.Monster.Sequences = records

	return nil
}
