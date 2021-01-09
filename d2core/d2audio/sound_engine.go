package d2audio

import (
	"fmt"
	"math/rand"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2records"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
)

type envState int

const (
	logPrefix = "Sound Engine"
)

const (
	envAttack  = 0
	envSustain = 1
	envRelease = 2
	envStopped = 3
)

const volMax float64 = 255
const originalFPS float64 = 25

// A Sound that can be started and stopped
type Sound struct {
	effect  d2interface.SoundEffect
	entry   *d2records.SoundDetailsRecord
	volume  float64
	vTarget float64
	vRate   float64
	state   envState
	// panning float64 // lets forget about this for now

	*d2util.Logger
}

func (s *Sound) update(elapsed float64) {
	// attack
	if s.state == envAttack {
		s.volume += s.vRate * elapsed
		if s.volume > s.vTarget {
			s.volume = s.vTarget
			s.state = envSustain
		}

		s.effect.SetVolume(s.volume)
	}

	// release
	if s.state == envRelease {
		s.volume -= s.vRate * elapsed
		if s.volume < 0 {
			s.effect.Stop()
			s.volume = 0
			s.state = envStopped
		}

		s.effect.SetVolume(s.volume)
	}
}

// SetPan sets the stereo pan, range -1 to 1
func (s *Sound) SetPan(pan float64) {
	s.effect.SetPan(pan)
}

// Play the sound
func (s *Sound) Play() {
	s.Info("starting sound" + s.entry.Handle)
	s.effect.Play()

	if s.entry.FadeIn != 0 {
		s.effect.SetVolume(0)
		s.volume = 0
		s.state = envAttack
		s.vTarget = float64(s.entry.Volume) / volMax
		s.vRate = s.vTarget / (float64(s.entry.FadeIn) / originalFPS)
	} else {
		s.volume = float64(s.entry.Volume) / volMax
		s.effect.SetVolume(s.volume)
		s.state = envSustain
	}
}

// Stop the sound, only required for looping sounds
func (s *Sound) Stop() {
	if s.entry.FadeOut != 0 {
		s.state = envRelease
		s.vTarget = 0
		s.vRate = s.volume / (float64(s.entry.FadeOut) / originalFPS)
	} else {
		s.state = envStopped
		s.volume = 0
		s.effect.SetVolume(s.volume)
		s.effect.Stop()
	}
}

// SoundEngine provides functions for playing sounds
type SoundEngine struct {
	asset    *d2asset.AssetManager
	provider d2interface.AudioProvider
	timer    float64
	accTime  float64
	sounds   map[*Sound]struct{}

	*d2util.Logger
}

// NewSoundEngine creates a new sound engine
func NewSoundEngine(provider d2interface.AudioProvider,
	asset *d2asset.AssetManager, l d2util.LogLevel, term d2interface.Terminal) *SoundEngine {
	r := SoundEngine{
		asset:    asset,
		provider: provider,
		sounds:   map[*Sound]struct{}{},
		timer:    1,
	}

	r.Logger = d2util.NewLogger()
	r.Logger.SetPrefix(logPrefix)
	r.Logger.SetLevel(l)

	err := term.BindAction("playsoundid", "plays the sound for a given id", func(id int) {
		r.PlaySoundID(id)
	})
	if err != nil {
		r.Error(err.Error())
		return nil
	}

	err = term.BindAction("playsound", "plays the sound for a given handle string", func(handle string) {
		r.PlaySoundHandle(handle)
	})
	if err != nil {
		r.Error(err.Error())
		return nil
	}

	err = term.BindAction("activesounds", "list currently active sounds", func() {
		for s := range r.sounds {
			if err != nil {
				r.Error(err.Error())
				return
			}

			r.Info(fmt.Sprint(s))
		}
	})

	err = term.BindAction("killsounds", "kill active sounds", func() {
		for s := range r.sounds {
			if err != nil {
				r.Error(err.Error())
				return
			}

			s.Stop()
		}
	})

	return &r
}

// Advance updates sound engine state, triggering events and envelopes
func (s *SoundEngine) Advance(elapsed float64) {
	s.timer -= elapsed
	s.accTime += elapsed

	if s.timer < 0 {
		for sound := range s.sounds {
			sound.update(s.accTime)

			// Clean up finished non-looping effects
			if !sound.effect.IsPlaying() {
				delete(s.sounds, sound)
			}

			// Clean up stopped looping effects
			if sound.state == envStopped {
				delete(s.sounds, sound)
			}
		}

		s.timer = 0.2
		s.accTime = 0
	}
}

// Reset stop all sounds and reset state
func (s *SoundEngine) Reset() {
	for snd := range s.sounds {
		snd.effect.Stop()
		delete(s.sounds, snd)
	}
}

// PlaySoundID plays a sound by sounds.txt index, returning the sound here is kinda ugly
// now we could have a situation where someone holds onto the sound after the sound engine is done with it
// someone needs to be in charge of deciding when to stopping looping sounds though...
func (s *SoundEngine) PlaySoundID(id int) *Sound {
	if id == 0 {
		return nil
	}

	entry := s.asset.Records.SelectSoundByIndex(id)

	if entry.GroupSize > 0 {
		// nolint:gosec // this is client-only, no big deal if rand index isn't securely generated
		indexOffset := rand.Intn(entry.GroupSize)
		entry = s.asset.Records.SelectSoundByIndex(entry.Index + indexOffset)
	}

	effect, err := s.provider.LoadSound(entry.FileName, entry.Loop, entry.MusicVol)
	if err != nil {
		s.Error(err.Error())
		return nil
	}

	snd := Sound{
		entry:  entry,
		effect: effect,
		Logger: s.Logger,
	}

	s.sounds[&snd] = struct{}{}

	snd.Play()

	return &snd
}

// PlaySoundHandle plays a sound by sounds.txt handle
func (s *SoundEngine) PlaySoundHandle(handle string) *Sound {
	sound := s.asset.Records.Sound.Details[handle].Index
	return s.PlaySoundID(sound)
}
