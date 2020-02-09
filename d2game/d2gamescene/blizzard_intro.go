package d2gamescene

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2data/d2video"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

type BlizzardIntro struct {
	videoDecoder *d2video.BinkDecoder
}

func CreateBlizzardIntro() *BlizzardIntro {
	return &BlizzardIntro{}
}

func (v *BlizzardIntro) OnLoad() error {
	videoBytes, err := d2asset.LoadFile("/data/local/video/BlizNorth640x480.bik")
	if err != nil {
		return err
	}

	v.videoDecoder = d2video.CreateBinkDecoder(videoBytes)
	return nil
}
