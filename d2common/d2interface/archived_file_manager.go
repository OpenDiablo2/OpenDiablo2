package d2interface

// ArchivedFileManager manages file access to the archives being managed
// by the ArchiveManager
type ArchivedFileManager interface {
	Cacher
	LoadFileStream(filePath string) (ArchiveDataStream, error)
	LoadFile(filePath string) ([]byte, error)
	FileExists(filePath string) (bool, error)
}
