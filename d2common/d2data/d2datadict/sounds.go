package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// SoundEntry represents a sound entry
type SoundEntry struct {
	Handle    string
	Index     int
	FileName  string
	Volume    int
	GroupSize int
	Loop      bool
	FadeIn    int
	FadeOut   int
	DeferInst bool
	StopInst  bool
	Duration  int
	Compound  int
	Reverb    int
	Falloff   int
	Cache     bool
	AsyncOnly bool
	Priority  int
	Stream    bool
	Stereo    bool
	Tracking  bool
	Solo      bool
	MusicVol  bool
	Block1    int
	Block2    int
	Block3    int
}

// Sounds stores all of the SoundEntries
//nolint:gochecknoglobals // Currently global by design, only written once
var Sounds map[string]SoundEntry

// LoadSounds loads SoundEntries from sounds.txt
func LoadSounds(file []byte) {
	Sounds = make(map[string]SoundEntry)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		entry := SoundEntry{
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
	for _, el := range Sounds {
		if el.Index == index {
			return &el
		}
	}

	return nil
}
