package Sound

import (
	"github.com/essial/OpenDiablo2/Common"
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
		Index:     Common.StringToInt(props[1]),
		FileName:  props[2],
		Volume:    Common.StringToUint8(props[3]),
		GroupSize: Common.StringToUint8(props[4]),
		Loop:      Common.StringToUint8(props[5]) == 1,
		FadeIn:    Common.StringToUint8(props[6]),
		FadeOut:   Common.StringToUint8(props[7]),
		DeferInst: Common.StringToUint8(props[8]),
		StopInst:  Common.StringToUint8(props[9]),
		Duration:  Common.StringToUint8(props[10]),
		Compound:  Common.StringToInt8(props[11]),
		Reverb:    Common.StringToUint8(props[12]) == 1,
		Falloff:   Common.StringToUint8(props[13]),
		Cache:     Common.StringToUint8(props[14]),
		AsyncOnly: Common.StringToUint8(props[15]) == 1,
		Priority:  Common.StringToUint8(props[16]),
		Stream:    Common.StringToUint8(props[17]),
		Stereo:    Common.StringToUint8(props[18]),
		Tracking:  Common.StringToUint8(props[19]),
		Solo:      Common.StringToUint8(props[20]),
		MusicVol:  Common.StringToUint8(props[21]),
		Block1:    Common.StringToInt(props[22]),
		Block2:    Common.StringToInt(props[23]),
		Block3:    Common.StringToInt(props[24]),
	}
	return result
}
