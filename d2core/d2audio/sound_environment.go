package d2audio

import (
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2datadict"
)

const assumedFPS = 25

// SoundEnvironment represents the audio environment for map areas
type SoundEnvironment struct {
	environment *d2datadict.SoundEnvironRecord
	engine      *SoundEngine
	bgm         *Sound
	ambiance    *Sound
	eventTimer  float64
}

// NewSoundEnvironment creates a SoundEnvironment using the given SoundEngine
func NewSoundEnvironment(soundEngine *SoundEngine) SoundEnvironment {
	r := SoundEnvironment{
		// Start with env NONE
		environment: d2datadict.SoundEnvirons[0],
		engine:      soundEngine,
	}

	return r
}

// SetEnv sets the sound environment using the given record index
func (s *SoundEnvironment) SetEnv(environmentIdx int) {
	if s.environment.Index != environmentIdx {
		newEnv := d2datadict.SoundEnvirons[environmentIdx]

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

// Advance advances the sound engine and plays sounds when necessary
func (s *SoundEnvironment) Advance(elapsed float64) {
	s.eventTimer -= elapsed
	if s.eventTimer < 0 {
		s.eventTimer = float64(s.environment.EventDelay) / assumedFPS

		snd := s.engine.PlaySoundID(s.environment.DayEvent)
		if snd != nil {
			pan := (rand.Float64() * 2) - 1
			snd.SetPan(pan)
		}
	}
}
