package d2records

// SoundDetails is a map of the SoundEntries
type SoundDetails map[string]*SoundDetailsRecord

// SoundDetailsRecord represents a sound entry
type SoundDetailsRecord struct {
	Handle    string
	FileName  string
	Index     int
	Volume    int
	GroupSize int
	FadeIn    int
	FadeOut   int
	Duration  int
	Compound  int
	Reverb    int
	Falloff   int
	Priority  int
	Block1    int
	Block2    int
	Block3    int
	Loop      bool
	DeferInst bool
	StopInst  bool
	Cache     bool
	AsyncOnly bool
	Stream    bool
	Stereo    bool
	Tracking  bool
	Solo      bool
	MusicVol  bool
}
