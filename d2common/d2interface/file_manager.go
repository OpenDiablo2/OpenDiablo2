package d2interface

// FileManager manages file access to the archives being managed
// by the ArchiveManager
type FileManager interface {
	Cacher
	LoadFileStream(filePath string) (DataStream, error)
	LoadFile(filePath string) ([]byte, error)
	FileExists(filePath string) (bool, error)
}
