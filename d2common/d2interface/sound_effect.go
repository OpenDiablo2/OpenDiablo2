package d2interface

// SoundEffect is something that that the AudioProvider can Play or Stop
type SoundEffect interface {
	Play()
	Stop()
	IsPlaying() bool
	SetVolume(volume float64)
}
