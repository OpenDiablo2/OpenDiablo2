package OpenDiablo2

import (
	"strings"
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
func CreateSoundEntry(soundLine string) SoundEntry {
	props := strings.Split(soundLine, "\t")
	result := SoundEntry{
		Handle:    props[0],
		Index:     StringToInt(props[1]),
		FileName:  props[2],
		Volume:    StringToUint8(props[3]),
		GroupSize: StringToUint8(props[4]),
		Loop:      StringToUint8(props[5]) == 1,
		FadeIn:    StringToUint8(props[6]),
		FadeOut:   StringToUint8(props[7]),
		DeferInst: StringToUint8(props[8]),
		StopInst:  StringToUint8(props[9]),
		Duration:  StringToUint8(props[10]),
		Compound:  StringToInt8(props[11]),
		Reverb:    StringToUint8(props[12]) == 1,
		Falloff:   StringToUint8(props[13]),
		Cache:     StringToUint8(props[14]),
		AsyncOnly: StringToUint8(props[15]) == 1,
		Priority:  StringToUint8(props[16]),
		Stream:    StringToUint8(props[17]),
		Stereo:    StringToUint8(props[18]),
		Tracking:  StringToUint8(props[19]),
		Solo:      StringToUint8(props[20]),
		MusicVol:  StringToUint8(props[21]),
		Block1:    StringToInt(props[22]),
		Block2:    StringToInt(props[23]),
		Block3:    StringToInt(props[24]),
	}
	return result
}
