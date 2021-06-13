package d2records

// MonsterSounds stores the MonsterSoundRecords
type MonsterSounds map[string]*MonsterSoundRecord

// MonsterSoundRecord represents a single line in MonSounds.txt
type MonsterSoundRecord struct {
	Neutral             string
	Attack1             string
	Weapon1             string
	CvtSk3              string
	CvtMo3              string
	CvtTgt2             string
	CvtSk2              string
	Attack2             string
	Weapon2             string
	CvtMo2              string
	CvtTgt1             string
	CvtSk1              string
	CvtMo1              string
	HitSound            string
	DeathSound          string
	Flee                string
	Taunt               string
	Skill1              string
	Skill2              string
	Skill3              string
	Skill4              string
	Footstep            string
	FootstepLayer       string
	Init                string
	ID                  string
	CvtTgt3             string
	Attack2Probability  int
	NeutralTime         int
	FootstepCount       int
	DeaDelay            int
	HitDelay            int
	Weapon2Volume       int
	FootstepOffset      int
	Weapon2Delay        int
	Attack2Delay        int
	Weapon1Volume       int
	Attack1Probability  int
	Weapon1Delay        int
	Attack1Delay        int
	FootstepProbability int
}
