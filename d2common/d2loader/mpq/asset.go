package mpq

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"
)

const (
	bufLength = 32
)

// static check that Asset implements Asset
var _ asset.Asset = &Asset{}

// Asset represents a file record within an MPQ archive
type Asset struct {
	stream d2interface.DataStream
	data   []byte
	path   string
	source *Source
}

// Type returns the asset type
func (a *Asset) Type() types.AssetType {
	return types.Ext2AssetType(filepath.Ext(a.Path()))
}

// Source returns the source of this asset
func (a *Asset) Source() asset.Source {
	return a.source
}

// Path returns the sub-path (within the source) of this asset
func (a *Asset) Path() string {
	return a.path
}

// Read will read asset data into the given buffer
func (a *Asset) Read(buf []byte) (n int, err error) {
	totalRead, err := a.stream.Read(buf)
	if totalRead == 0 {
		return 0, io.EOF
	}

	return totalRead, nil
}

// Seek will seek the read position for the next read operation
func (a *Asset) Seek(offset int64, whence int) (n int64, err error) {
	return a.stream.Seek(offset, whence)
}

// Seek will seek the read position for the next read operation
func (a *Asset) Close() (err error) {
	_, err = a.Seek(0, 0)
	return err
}

// Data returns the raw file data as a slice of bytes
func (a *Asset) Data() ([]byte, error) {
	if a.stream == nil {
		return nil, fmt.Errorf("asset has no file: %s", a.Path())
	}

	if a.data != nil {
		return a.data, nil
	}

	_, seekErr := a.Seek(0, 0)
	if seekErr != nil {
		return nil, seekErr
	}

	buf := make([]byte, bufLength)
	data := make([]byte, 0)

	for {
		numBytesRead, readErr := a.Read(buf)

		data = append(data, buf[:numBytesRead]...)

		if readErr != nil || numBytesRead == 0 {
			break
		}
	}

	a.data = data

	return data, nil
}

// Path returns the path
func (a *Asset) String() string {
	return a.Path()
}
