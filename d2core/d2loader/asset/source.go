package asset

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2loader/asset/types"
)

// Source is an abstraction for something that can load and list assets
type Source interface {
	Type() types.SourceType
	Open(name string) (Asset, error)
	String() string
}
