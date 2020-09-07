package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// Information gathered from [https://d2mods.info/forum/kb/viewarticle?a=418]

// MonsterSoundRecord represents a single line in MonSounds.txt
type MonsterSoundRecord struct {
	// ID is the identifier, used in MonStats.txt to refer to a particular sound record
	ID string

	// Melee attack sound ID, refers to a sound from Sounds.txt
	Attack1 string

	// Weapon attack sound ID, refers to a sound from Sounds.txt
	Weapon1 string

	// Delay in frames of Attack1 sound
	Attack1Delay int

	// Delay in frames of Weapon1 sound
	Weapon1Delay int

	// Probability of playing Attack1 sound instead of Weapon1
	Attack1Probability int

	// Overrides weapon volume from Sounds.txt
	Weapon1Volume int

	// Ditto, 2 sets of sounds are possible
	Attack2            string
	Weapon2            string
	Attack2Delay       int
	Weapon2Delay       int
	Attack2Probability int
	Weapon2Volume      int

	// Sound when monster takes a hit, refers to a sound from Sounds.txt
	HitSound string

	// Sound when monster dies, refers to a sound from Sounds.txt
	DeathSound string

	// Delay in frames of HitSound
	HitDelay int

	// Delay in frames of DeathSound
	DeaDelay int

	// Sound when monster enters skill mode
	Skill1 string
	Skill2 string
	Skill3 string
	Skill4 string

	// Sound played each loop of the WL animation
	Footstep string

	// Additional WL animation sound
	FootstepLayer string

	// Number of footstep sounds played (e.g. 2 for two-legged monsters)
	FootstepCount int

	// FsOff, possibly delay between footstep sounds
	FootstepOffset int

	// Probability of playing footstep sound, percentage
	FootstepProbability int

	// Sound when monster is neutral (also played when walking)
	Neutral string

	// Delay in frames between neutral sounds
	NeutralTime int

	// Sound when monster is initialized
	Init string

	// Sound when monster is encountered
	Taunt string

	// Sound when monster retreats
	Flee string

	// The following are related to skills in some way
	// Initial monster animation code (MonMode.txt)
	CvtMo1 string
	// ID of skill
	CvtSk1 string
	// End monster animation code (MonMode.txt)
	CvtTgt1 string

	CvtMo2  string
	CvtSk2  string
	CvtTgt2 string

	CvtMo3  string
	CvtSk3  string
	CvtTgt3 string
}

// MonsterSounds stores the MonsterSoundRecords
//nolint:gochecknoglobals // Currently global by design
var MonsterSounds map[string]*MonsterSoundRecord

// LoadMonsterSounds loads MonsterSoundRecords into MonsterSounds
func LoadMonsterSounds(file []byte) {
	MonsterSounds = make(map[string]*MonsterSoundRecord)

	d := d2txt.LoadDataDictionary(file)
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
		MonsterSounds[record.ID] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d MonsterUniqueModifier records", len(MonsterUniqueModifiers))
}
