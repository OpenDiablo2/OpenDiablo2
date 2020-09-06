package types

import "strings"

// AssetType represents the type of an asset
type AssetType int

// Asset types
const (
	AssetTypeUnknown AssetType = iota
	AssetTypeJSON
	AssetTypeStringTable
	AssetTypeDataDictionary
	AssetTypePalette
	AssetTypePaletteTransform
	AssetTypeCOF
	AssetTypeDC6
	AssetTypeDCC
	AssetTypeDS1
	AssetTypeDT1
	AssetTypeWAV
	AssetTypeD2
)

// Ext2AssetType determines the AssetType with the given file extension
func Ext2AssetType(ext string) AssetType {
	ext = strings.ToLower(ext)
	ext = strings.ReplaceAll(ext, ".", "")

	lookup := map[string]AssetType{
		"json": AssetTypeJSON,
		"tbl":  AssetTypeStringTable,
		"txt":  AssetTypeDataDictionary,
		"dat":  AssetTypePalette,
		"pl2":  AssetTypePaletteTransform,
		"cof":  AssetTypeCOF,
		"dc6":  AssetTypeDC6,
		"dcc":  AssetTypeDCC,
		"ds1":  AssetTypeDS1,
		"dt1":  AssetTypeDT1,
		"wav":  AssetTypeWAV,
		"d2":   AssetTypeD2,
	}

	if knownType, found := lookup[ext]; found {
		return knownType
	}

	return AssetTypeUnknown
}
