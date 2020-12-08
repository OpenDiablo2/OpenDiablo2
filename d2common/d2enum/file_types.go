package d2enum

// FileType represents the type of an asset
type FileType int

// File types
const (
	FileTypeUnknown FileType = iota
	FileTypeDirectory
	FileTypeMPQ
	FileTypeJSON
	FileTypeStringTable
	FileTypeFontTable
	FileTypeDataDictionary
	FileTypePalette
	FileTypePaletteTransform
	FileTypeCOF
	FileTypeDC6
	FileTypeDCC
	FileTypeDS1
	FileTypeDT1
	FileTypeWAV
	FileTypeD2
	FileTypeLocale
)
