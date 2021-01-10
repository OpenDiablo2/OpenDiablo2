package filesystem

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset"
)

// OnAddSource is a shim method to allow loading of filesystem sources
func OnAddSource(path string) (asset.Source, error) {
	return &Source{
		Root: path,
	}, nil
}
