package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// Information gathered from [https://d2mods.info/forum/kb/viewarticle?a=418]

// LoadMonsterSounds loads MonsterSoundRecords into MonsterSounds
func monsterSoundsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(MonsterSounds)

	for d.Next() {
		record := &MonsterSoundRecord{
			ID:                  d.String("Id"),
			Attack1:             d.String("Attack1"),
			Weapon1:             d.String("Weapon1"),
			Attack1Delay:        d.Number("Att1Del"),
			Weapon1Delay:        d.Number("Wea1Del"),
			Attack1Probability:  d.Number("Att1Prb"),
			Weapon1Volume:       d.Number("Wea1Vol"),
			Attack2:             d.String("Attack2"),
			Weapon2:             d.String("Weapon2"),
			Attack2Delay:        d.Number("Att2Del"),
			Weapon2Delay:        d.Number("Wea2Del"),
			Attack2Probability:  d.Number("Att2Prb"),
			Weapon2Volume:       d.Number("Wea2Vol"),
			Skill1:              d.String("Skill1"),
			Skill2:              d.String("Skill2"),
			Skill3:              d.String("Skill3"),
			Skill4:              d.String("Skill4"),
			Footstep:            d.String("Footstep"),
			FootstepLayer:       d.String("FootstepLayer"),
			FootstepCount:       d.Number("FsCnt"),
			FootstepOffset:      d.Number("FsOff"),
			FootstepProbability: d.Number("FsPrb"),
			Neutral:             d.String("Neutral"),
			NeutralTime:         d.Number("NeuTime"),
			Init:                d.String("Init"),
			Taunt:               d.String("Taunt"),
			Flee:                d.String("Flee"),
			CvtMo1:              d.String("CvtMo1"),
			CvtMo2:              d.String("CvtMo2"),
			CvtMo3:              d.String("CvtMo3"),
			CvtSk1:              d.String("CvtSk1"),
			CvtSk2:              d.String("CvtSk2"),
			CvtSk3:              d.String("CvtSk3"),
			CvtTgt1:             d.String("CvtTgt1"),
			CvtTgt2:             d.String("CvtTgt2"),
			CvtTgt3:             d.String("CvtTgt3"),
		}

		records[record.ID] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d Monster Sound records", len(records))

	r.Monster.Sounds = records

	return nil
}
