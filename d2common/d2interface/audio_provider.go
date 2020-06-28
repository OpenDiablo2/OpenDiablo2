package d2interface

type AudioProvider interface {
	PlayBGM(song string)
	LoadSoundEffect(sfx string) (SoundEffect, error)
	SetVolumes(bgmVolume, sfxVolume float64)
}
