package d2records

// SoundEnvironments contains the SoundEnviron records
type SoundEnvironments map[int]*SoundEnvironRecord

// SoundEnvironRecord describes the different sound environments. Not listed on Phrozen Keep.
type SoundEnvironRecord struct {
	Handle          string
	Index           int
	Song            int
	DayAmbience     int
	NightAmbience   int
	DayEvent        int
	NightEvent      int
	EventDelay      int
	Indoors         int
	Material1       int
	Material2       int
	EAXEnviron      int
	EAXEnvSize      int
	EAXEnvDiff      int
	EAXRoomVol      int
	EAXRoomHF       int
	EAXDecayTime    int
	EAXDecayHF      int
	EAXReflect      int
	EAXReflectDelay int
	EAXReverb       int
	EAXRevDelay     int
	EAXRoomRoll     int
	EAXAirAbsorb    int
}
