package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func soundEnvironmentLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(SoundEnvironments)

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

		records[record.Index] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Sound.Environment = records

	r.Logger.Infof("Loaded %d SoundEnviron records", len(records))

	return nil
}
