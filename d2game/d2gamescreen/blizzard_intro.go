package d2gamescreen

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2video"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
)

// BlizzardIntro represents the Blizzard Intro screen
type BlizzardIntro struct {
	videoDecoder *d2video.BinkDecoder
}

// CreateBlizzardIntro creates a Blizzard Intro screen
func CreateBlizzardIntro() *BlizzardIntro {
	return &BlizzardIntro{}
}

// OnLoad loads the resources for the Blizzard Intro screen
func (v *BlizzardIntro) OnLoad(loading d2screen.LoadingState) {
	videoBytes, err := d2asset.LoadFile("/data/local/video/BlizNorth640x480.bik")
	if err != nil {
		loading.Error(err)
		return
	}
	
	loading.Progress(fiftyPercent)

	v.videoDecoder = d2video.CreateBinkDecoder(videoBytes)
}
