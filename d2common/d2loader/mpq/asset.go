package mpq

import (
	"path/filepath"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"
)

// static check that Asset implements Asset
var _ asset.Asset = &Asset{}

// Asset represents a file record within an MPQ archive
type Asset struct {
	stream d2interface.ArchiveDataStream
	name   string
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
	return a.name
}

// Read will read asset data into the given buffer
func (a *Asset) Read(buf []byte) (n int, err error) {
	return a.stream.Read(buf)
}

// Seek will seek the read position for the next read operation
func (a *Asset) Seek(offset int64, whence int) (n int64, err error) {
	return a.stream.Seek(offset, whence)
}
