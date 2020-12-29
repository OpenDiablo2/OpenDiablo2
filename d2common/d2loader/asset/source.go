package asset

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset/types"
)

// Source is an abstraction for something that can load and list assets
type Source interface {
	fmt.Stringer
	Type() types.SourceType
	Open(name string) (Asset, error)
	Path() string
}
