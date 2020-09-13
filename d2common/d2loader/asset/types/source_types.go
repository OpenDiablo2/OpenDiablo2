package types

import (
	"path/filepath"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2mpq"
)

// SourceType represents the type of the asset source
type SourceType int

// Asset sources
const (
	AssetSourceUnknown SourceType = iota
	AssetSourceFileSystem
	AssetSourceMPQ
)

// Ext2SourceType returns the SourceType from the given file extension
func Ext2SourceType(ext string) SourceType {
	ext = strings.ToLower(ext)
	ext = strings.ReplaceAll(ext, ".", "")

	lookup := map[string]SourceType{
		"mpq": AssetSourceMPQ,
	}

	if knownType, found := lookup[ext]; found {
		return knownType
	}

	return AssetSourceUnknown
}

// CheckSourceType attempts to determine the source type of the source
func CheckSourceType(path string) SourceType {
	// on MacOS, the MPQ's from blizzard don't have file extensions
	// so we just attempt to init the file as an mpq
	if _, err := d2mpq.Load(path); err == nil {
		return AssetSourceMPQ
	}

	ext := filepath.Ext(path)

	return Ext2SourceType(ext)
}
