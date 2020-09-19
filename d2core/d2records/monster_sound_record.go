package d2records

// MonsterSounds stores the MonsterSoundRecords
type MonsterSounds map[string]*MonsterSoundRecord

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
