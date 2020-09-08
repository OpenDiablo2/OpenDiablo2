package types

import "strings"

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
