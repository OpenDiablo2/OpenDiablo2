package filesystem

import (
	"os"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"
)

// static check that Asset implements Asset
var _ asset.Asset = &Asset{}

// Asset represents an asset that is in the host filesystem
type Asset struct {
	assetType types.AssetType
	source    *Source
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

// String returns the path
func (fsa *Asset) String() string {
	return fsa.Path()
}
