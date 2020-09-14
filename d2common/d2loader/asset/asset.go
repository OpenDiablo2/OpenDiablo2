package asset

import (
	"fmt"
	"io"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"
)

// Asset represents a game asset. It has a type, an asset source, a sub-path (within the
// asset source), and it can read data and seek within the data
type Asset interface {
	fmt.Stringer
	io.Reader
	io.Seeker
	io.Closer
	Type() types.AssetType
	Source() Source
	Path() string
	Data() ([]byte, error)
}
