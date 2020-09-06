package asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2loader/asset/types"
	"io"
)

// Asset represents a game asset. It has a type, an asset source, a sub-path (within the
// asset source), and it can read data and seek within the data
type Asset interface {
	io.ReadSeeker
	Type() types.AssetType
	Source() Source
	Path() string
}
