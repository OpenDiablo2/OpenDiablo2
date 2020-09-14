package filesystem

import (
	"fmt"
	"os"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"
)

const (
	bufLength = 32
)

// static check that Asset implements Asset
var _ asset.Asset = &Asset{}

// Asset represents an asset that is in the host filesystem
type Asset struct {
	assetType types.AssetType
	source    *Source
	data      []byte
	path      string
	file      *os.File
}

// Type returns the asset type
func (fsa *Asset) Type() types.AssetType {
	return fsa.assetType
}

// Source returns the asset source that this asset was loaded from
func (fsa *Asset) Source() asset.Source {
	return fsa.source
}

// Path returns the sub-path (within the asset source Root) for this asset
func (fsa *Asset) Path() string {
	return fsa.path
}

// Read reads bytes into the given byte buffer
func (fsa *Asset) Read(p []byte) (n int, err error) {
	return fsa.file.Read(p)
}

// Seek seeks within the file
func (fsa *Asset) Seek(offset int64, whence int) (int64, error) {
	return fsa.file.Seek(offset, whence)
}

// Close closes the file
func (fsa *Asset) Close() error {
	return fsa.file.Close()
}

// Data returns the raw file data as a slice of bytes
func (fsa *Asset) Data() ([]byte, error) {
	if fsa.file == nil {
		return nil, fmt.Errorf("asset has no file: %s", fsa.Path())
	}

	if fsa.data != nil {
		return fsa.data, nil
	}

	_, seekErr := fsa.file.Seek(0, 0)
	if seekErr != nil {
		return nil, seekErr
	}

	buf := make([]byte, bufLength)
	data := make([]byte, 0)

	for {
		numBytesRead, readErr := fsa.Read(buf)

		data = append(data, buf[:numBytesRead]...)

		if readErr != nil {
			break
		}
	}

	fsa.data = data

	return data, nil
}

// String returns the path
func (fsa *Asset) String() string {
	return fsa.Path()
}
