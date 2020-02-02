package d2datadict

import (
	"log"
	"strings"

	dh "github.com/OpenDiablo2/OpenDiablo2/d2common"
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
		Index:     dh.StringToInt(props[inc()]),
		FileName:  props[inc()],
		Volume:    dh.StringToUint8(props[inc()]),
		GroupSize: dh.StringToUint8(props[inc()]),
		Loop:      dh.StringToUint8(props[inc()]) == 1,
		FadeIn:    dh.StringToUint8(props[inc()]),
		FadeOut:   dh.StringToUint8(props[inc()]),
		DeferInst: dh.StringToUint8(props[inc()]),
		StopInst:  dh.StringToUint8(props[inc()]),
		Duration:  dh.StringToUint8(props[inc()]),
		Compound:  dh.StringToInt8(props[inc()]),
		Reverb:    dh.StringToUint8(props[inc()]) == 1,
		Falloff:   dh.StringToUint8(props[inc()]),
		Cache:     dh.StringToUint8(props[inc()]),
		AsyncOnly: dh.StringToUint8(props[inc()]) == 1,
		Priority:  dh.StringToUint8(props[inc()]),
		Stream:    dh.StringToUint8(props[inc()]),
		Stereo:    dh.StringToUint8(props[inc()]),
		Tracking:  dh.StringToUint8(props[inc()]),
		Solo:      dh.StringToUint8(props[inc()]),
		MusicVol:  dh.StringToUint8(props[inc()]),
		Block1:    dh.StringToInt(props[inc()]),
		Block2:    dh.StringToInt(props[inc()]),
		Block3:    dh.StringToInt(props[inc()]),
	}
	return result
}

var Sounds map[string]SoundEntry

func LoadSounds(file []byte) {
	Sounds = make(map[string]SoundEntry)
	soundData := strings.Split(string(file), "\r\n")[1:]
	for _, line := range soundData {
		if len(line) == 0 {
			continue
		}
		soundEntry := createSoundEntry(line)
		soundEntry.FileName = "/data/global/sfx/" + strings.ReplaceAll(soundEntry.FileName, `\`, "/")
		Sounds[soundEntry.Handle] = soundEntry
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
	}
	log.Printf("Loaded %d sound definitions", len(Sounds))
}
