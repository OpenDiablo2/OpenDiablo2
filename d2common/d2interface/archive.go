package d2interface

type Archive interface {
	Path() string
	Contains(string) bool
	Size() uint32
	Close()
	FileExists(fileName string) bool
	ReadFile(fileName string) ([]byte, error)
	ReadFileStream(fileName string) (ArchiveDataStream, error)
	ReadTextFile(fileName string) (string, error)
	GetFileList() ([]string, error)
}
