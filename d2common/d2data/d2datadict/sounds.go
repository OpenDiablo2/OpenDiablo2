package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// SoundEntry represents a sound entry
type SoundEntry struct {
	Handle    string
	FileName  string
	Index     int
	Volume    int
	GroupSize int
	FadeIn    int
	FadeOut   int
	Duration  int
	Compound  int
	Reverb    int
	Falloff   int
	Priority  int
	Block1    int
	Block2    int
	Block3    int
	Loop      bool
	DeferInst bool
	StopInst  bool
	Cache     bool
	AsyncOnly bool
	Stream    bool
	Stereo    bool
	Tracking  bool
	Solo      bool
	MusicVol  bool
}

// Sounds stores all of the SoundEntries
//nolint:gochecknoglobals // Currently global by design, only written once
var Sounds map[string]*SoundEntry

// LoadSounds loads SoundEntries from sounds.txt
func LoadSounds(file []byte) {
	Sounds = make(map[string]*SoundEntry)

	d := d2txt.LoadDataDictionary(file)
	for d.Next() {
		entry := &SoundEntry{
			Handle:    d.String("Sound"),
			Index:     d.Number("Index"),
			FileName:  d.String("FileName"),
			Volume:    d.Number("Volume"),
			GroupSize: d.Number("Group Size"),
			Loop:      d.Bool("Loop"),
			FadeIn:    d.Number("Fade In"),
			FadeOut:   d.Number("Fade Out"),
			DeferInst: d.Bool("Defer Inst"),
			StopInst:  d.Bool("Stop Inst"),
			Duration:  d.Number("Duration"),
			Compound:  d.Number("Compound"),
			Reverb:    d.Number("Reverb"),
			Falloff:   d.Number("Falloff"),
			Cache:     d.Bool("Cache"),
			AsyncOnly: d.Bool("Async Only"),
			Priority:  d.Number("Priority"),
			Stream:    d.Bool("Stream"),
			Stereo:    d.Bool("Stereo"),
			Tracking:  d.Bool("Tracking"),
			Solo:      d.Bool("Solo"),
			MusicVol:  d.Bool("Music Vol"),
			Block1:    d.Number("Block 1"),
			Block2:    d.Number("Block 2"),
			Block3:    d.Number("Block 3"),
		}
		Sounds[entry.Handle] = entry
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d sound definitions", len(Sounds))
}

// SelectSoundByIndex selects a sound by its ID
func SelectSoundByIndex(index int) *SoundEntry {
	for idx := range Sounds {
		if Sounds[idx].Index == index {
			return Sounds[idx]
		}
	}

	return nil
}
