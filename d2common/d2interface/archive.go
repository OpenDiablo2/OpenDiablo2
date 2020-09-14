package d2interface

// Archive is an abstract representation of a game archive file
// For the original Diablo II, archives are always MPQ's, but
// OpenDiablo2 can handle any kind of archive file  as long as it
// implements this interface
type Archive interface {
	Path() string
	Contains(string) bool
	Size() uint32
	Close()
	FileExists(fileName string) bool
	ReadFile(fileName string) ([]byte, error)
	ReadFileStream(fileName string) (DataStream, error)
	ReadTextFile(fileName string) (string, error)
	GetFileList() ([]string, error)
}
