package filesystem

import (
	"os"
	"path/filepath"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"
)

// static check that Source implements AssetSource
var _ asset.Source = &Source{}

// Source represents an asset source which is a normal directory on the host file system
type Source struct {
	Root string
}

// Type returns the type of this asset source
func (s *Source) Type() types.SourceType {
	return types.AssetSourceFileSystem
}

// Open opens a file with the given sub-path within the Root dir of the file system source
func (s *Source) Open(subPath string) (asset.Asset, error) {
	file, err := os.Open(s.fullPath(subPath))

	if err == nil {
		a := &Asset{
			assetType: types.Ext2AssetType(filepath.Ext(subPath)),
			source:    s,
			path:      subPath,
			file:      file,
		}

		return a, nil
	}

	return nil, err
}

func (s *Source) fullPath(subPath string) string {
	return filepath.Clean(filepath.Join(s.Root, subPath))
}

// Path returns the Root dir of this file system source
func (s *Source) Path() string {
	return s.Root
}

// String returns the path
func (s *Source) String() string {
	return s.Path()
}
