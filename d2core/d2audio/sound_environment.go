package d2audio

import (
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"
)

const assumedFPS = 25

// SoundEnvironment represents the audio environment for map areas
type SoundEnvironment struct {
	environment *d2records.SoundEnvironRecord
	engine      *SoundEngine
	bgm         *Sound
	ambiance    *Sound
	eventTimer  float64
}

// NewSoundEnvironment creates a SoundEnvironment using the given SoundEngine
func NewSoundEnvironment(soundEngine *SoundEngine) SoundEnvironment {
	r := SoundEnvironment{
		// Start with env NONE
		environment: soundEngine.asset.Records.Sound.Environment[0],
		engine:      soundEngine,
	}

	return r
}

// SetEnv sets the sound environment using the given record index
func (s *SoundEnvironment) SetEnv(environmentIdx int) {
	if s.environment.Index != environmentIdx {
		newEnv := s.engine.asset.Records.Sound.Environment[environmentIdx]

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
			// nolint:gosec // client-side, no big deal if rand number isn't securely generated
			pan := (rand.Float64() * 2) - 1
			snd.SetPan(pan)
		}
	}
}
