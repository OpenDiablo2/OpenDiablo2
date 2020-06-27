package d2mpq

type MpqDataStream struct {
	stream *Stream
}

func (m *MpqDataStream) Read(p []byte) (n int, err error) {
	totalRead := m.stream.Read(p, 0, uint32(len(p)))
	return int(totalRead), nil
}

func (m *MpqDataStream) Seek(offset int64, whence int) (int64, error) {
	m.stream.CurrentPosition = uint32(offset + int64(whence))
	return int64(m.stream.CurrentPosition), nil
}

func (m *MpqDataStream) Close() error {
	m.stream = nil
	return nil
}
