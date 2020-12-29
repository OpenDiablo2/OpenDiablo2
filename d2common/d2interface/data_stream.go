package d2interface

// DataStream is a data stream
type DataStream interface {
	Read(p []byte) (n int, err error)
	Seek(offset int64, whence int) (int64, error)
	Close() error
}
