package d2interface

type ArchiveDataStream interface {
	Read(p []byte) (n int, err error)
	Seek(offset int64, whence int) (int64, error)
	Close() error
}
