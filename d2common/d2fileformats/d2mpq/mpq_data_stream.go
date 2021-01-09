package d2mpq

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

var _ d2interface.DataStream = &MpqDataStream{} // Static check to confirm struct conforms to interface

// MpqDataStream represents a stream for MPQ data.
type MpqDataStream struct {
	stream *Stream
}

// Read reads data from the data stream
func (m *MpqDataStream) Read(p []byte) (n int, err error) {
	totalRead := m.stream.Read(p, 0, uint32(len(p)))
	return int(totalRead), nil
}

// Seek sets the position of the data stream
func (m *MpqDataStream) Seek(offset int64, whence int) (int64, error) {
	m.stream.CurrentPosition = uint32(offset + int64(whence))
	return int64(m.stream.CurrentPosition), nil
}

// Close closes the data stream
func (m *MpqDataStream) Close() error {
	m.stream = nil
	return nil
}
