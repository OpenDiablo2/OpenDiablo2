package d2interface

// AudioProvider is something that can play music, load audio files managed
// by the asset manager, and set the game engine's volume levels
type AudioProvider interface {
	PlayBGM(song string)
	LoadSoundEffect(sfx string) (SoundEffect, error)
	SetVolumes(bgmVolume, sfxVolume float64)
}
