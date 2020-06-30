package d2datadict

import (
	"log"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// SoundEntry represents a sound entry
type SoundEntry struct {
	Handle    string
	Index     int
	FileName  string
	Volume    byte
	GroupSize uint8
	Loop      bool
	FadeIn    uint8
	FadeOut   uint8
	DeferInst uint8
	StopInst  uint8
	Duration  uint8
	Compound  int8
	Reverb    bool
	Falloff   uint8
	Cache     uint8
	AsyncOnly bool
	Priority  uint8
	Stream    uint8
	Stereo    uint8
	Tracking  uint8
	Solo      uint8
	MusicVol  uint8
	Block1    int
	Block2    int
	Block3    int
}

// CreateSoundEntry creates a sound entry based on a sound row on sounds.txt
func createSoundEntry(soundLine string) SoundEntry {
	props := strings.Split(soundLine, "\t")
	i := -1
	inc := func() int {
		i++
		return i
	}
	result := SoundEntry{
		Handle:    props[inc()],
		Index:     d2common.StringToInt(props[inc()]),
		FileName:  props[inc()],
		Volume:    d2common.StringToUint8(props[inc()]),
		GroupSize: d2common.StringToUint8(props[inc()]),
		Loop:      d2common.StringToUint8(props[inc()]) == 1,
		FadeIn:    d2common.StringToUint8(props[inc()]),
		FadeOut:   d2common.StringToUint8(props[inc()]),
		DeferInst: d2common.StringToUint8(props[inc()]),
		StopInst:  d2common.StringToUint8(props[inc()]),
		Duration:  d2common.StringToUint8(props[inc()]),
		Compound:  d2common.StringToInt8(props[inc()]),
		Reverb:    d2common.StringToUint8(props[inc()]) == 1,
		Falloff:   d2common.StringToUint8(props[inc()]),
		Cache:     d2common.StringToUint8(props[inc()]),
		AsyncOnly: d2common.StringToUint8(props[inc()]) == 1,
		Priority:  d2common.StringToUint8(props[inc()]),
		Stream:    d2common.StringToUint8(props[inc()]),
		Stereo:    d2common.StringToUint8(props[inc()]),
		Tracking:  d2common.StringToUint8(props[inc()]),
		Solo:      d2common.StringToUint8(props[inc()]),
		MusicVol:  d2common.StringToUint8(props[inc()]),
		Block1:    d2common.StringToInt(props[inc()]),
		Block2:    d2common.StringToInt(props[inc()]),
		Block3:    d2common.StringToInt(props[inc()]),
	}

	return result
}

// Sounds stores all of the SoundEntries
//nolint:gochecknoglobals // Currently global by design, only written once
var Sounds map[string]SoundEntry

// LoadSounds loads SoundEntries from sounds.txt
func LoadSounds(file []byte) {
	Sounds = make(map[string]SoundEntry)
	soundData := strings.Split(string(file), "\r\n")[1:]

	for _, line := range soundData {
		if line == "" {
			continue
		}

		soundEntry := createSoundEntry(line)
		soundEntry.FileName = "/data/global/sfx/" + strings.ReplaceAll(soundEntry.FileName, `\`, "/")
		Sounds[soundEntry.Handle] = soundEntry

		//nolint:gocritic // Debug util code
		/*
			// Use the following code to write out the values
			f, err := os.OpenFile(`C:\Users\lunat\Desktop\D2\sounds.txt`,
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString("\n[" + soundEntry.Handle + "] " + soundEntry.FileName); err != nil {
				log.Println(err)
			}
		*/
	} //nolint:wsl // Debug util code

	log.Printf("Loaded %d sound definitions", len(Sounds))
}
