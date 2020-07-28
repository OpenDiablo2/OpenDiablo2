package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

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

// SoundEnvirons contains the SoundEnviron records
//nolint:gochecknoglobals // Currently global by design, only written once
var SoundEnvirons map[string]*SoundEnvironRecord

// LoadSoundEnvirons loads SoundEnvirons from the supplied file
func LoadSoundEnvirons(file []byte) {
	SoundEnvirons = make(map[string]*SoundEnvironRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &SoundEnvironRecord{
			Handle:          d.String("Handle"),
			Index:           d.Number("Index"),
			Song:            d.Number("Song"),
			DayAmbience:     d.Number("Day Ambience"),
			NightAmbience:   d.Number("Night Ambience"),
			DayEvent:        d.Number("Day Event"),
			NightEvent:      d.Number("Night Event"),
			EventDelay:      d.Number("Event Delay"),
			Indoors:         d.Number("Indoors"),
			Material1:       d.Number("Material 1"),
			Material2:       d.Number("Material 2"),
			EAXEnviron:      d.Number("EAX Environ"),
			EAXEnvSize:      d.Number("EAX Env Size"),
			EAXEnvDiff:      d.Number("EAX Env Diff"),
			EAXRoomVol:      d.Number("EAX Room Vol"),
			EAXRoomHF:       d.Number("EAX Room HF"),
			EAXDecayTime:    d.Number("EAX Decay Time"),
			EAXDecayHF:      d.Number("EAX Decay HF"),
			EAXReflect:      d.Number("EAX Reflect"),
			EAXReflectDelay: d.Number("EAX Reflect Delay"),
			EAXReverb:       d.Number("EAX Reverb"),
			EAXRevDelay:     d.Number("EAX Rev Delay"),
			EAXRoomRoll:     d.Number("EAX Room Roll"),
			EAXAirAbsorb:    d.Number("EAX Air Absorb"),
		}
		SoundEnvirons[record.Handle] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d SoundEnviron records", len(SoundEnvirons))
}
