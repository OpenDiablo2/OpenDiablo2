package d2interface

// ArchiveDataStream is an archive data stream
type ArchiveDataStream interface {
	Read(p []byte) (n int, err error)
	Seek(offset int64, whence int) (int64, error)
	Close() error
}
