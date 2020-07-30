package d2audio

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
)

type SoundEnvironment struct {
	environment *d2datadict.SoundEnvironRecord
	engine      *SoundEngine
	bgm         *Sound
	ambiance    *Sound
	eventTimer  float64
}

func NewSoundEnvironment(soundEngine *SoundEngine) SoundEnvironment {
	r := SoundEnvironment{
		// Start with env NONE
		environment: d2datadict.SoundEnvirons[0],
		engine:      soundEngine,
	}

	return r
}

func (s *SoundEnvironment) SetEnv(environment int) {
	if s.environment.Index != environment {
		newEnv := d2datadict.SoundEnvirons[environment]

		if s.environment.Song != newEnv.Song {
			if s.bgm != nil {
				s.bgm.Stop()
			}

			s.bgm = s.engine.PlaySoundID(newEnv.Song)
		}

		if s.environment.DayAmbience != newEnv.DayAmbience {
			if s.ambiance != nil {
				s.ambiance.Stop()
			}

			s.ambiance = s.engine.PlaySoundID(newEnv.DayAmbience)
		}

		s.environment = newEnv
	}
}

func (s *SoundEnvironment) Advance(elapsed float64) {
	s.eventTimer -= elapsed
	if s.eventTimer < 0 {
		s.eventTimer = float64(s.environment.EventDelay) / 25
		s.engine.PlaySoundID(s.environment.DayEvent)
	}
}
