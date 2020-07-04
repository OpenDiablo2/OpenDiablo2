package d2interface

// Configuration saves, loads, and returns the OpenDiablo2
// configuration. This is either loaded from disk, or generated
// when one is not found.
type Configuration interface {
	Load() error
	Save() error
	// Get() Configuration
	MpqLoadOrder() []string
	Language() string
	MpqPath() string
	TicksPerSecond() int
	FpsCap() int
	SfxVolume() float64
	BgmVolume() float64
	FullScreen() bool
	RunInBackground() bool
	VsyncEnabled() bool
	Backend() string
}
