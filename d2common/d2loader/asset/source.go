package asset

import (
	"fmt"
	"io"
)

// Source is an abstraction for something that can load and list assets
type Source interface {
	fmt.Stringer
	Open(name string) (io.ReadSeeker, error)
	Path() string
	Exists(subPath string) bool
}
