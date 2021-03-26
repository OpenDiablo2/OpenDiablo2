package mpq

import (
	"io"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2loader/asset"
)

// static check that Source implements AssetSource
var _ asset.Source = &Source{}

// NewSource creates a new MPQ Source
func NewSource(sourcePath string) (asset.Source, error) {
	loaded, err := d2mpq.FromFile(sourcePath)
	if err != nil {
		return nil, err
	}

	return &Source{loaded}, nil
}

// Source is an implementation of an asset source for MPQ archives
type Source struct {
	MPQ d2interface.Archive
}

// Open attempts to open a file within the MPQ archive
func (v *Source) Open(name string) (a io.ReadSeeker, err error) {
	name = cleanName(name)
	return v.MPQ.ReadFileStream(name)
}

// Exists returns true if the file exists
func (v *Source) Exists(subPath string) bool {
	subPath = cleanName(subPath)
	return v.MPQ.Contains(subPath)
}

// Path returns the path of the MPQ on the host filesystem
func (v *Source) Path() string {
	return v.MPQ.Path()
}

// String returns the path
func (v *Source) String() string {
	return v.Path()
}

func cleanName(name string) string {
	name = strings.ReplaceAll(name, "/", "\\")

	if string(name[0]) == "\\" {
		name = name[1:]
	}

	return name
}
