// Package d2records provides a RecordManager implementation which is used to parse
// the various txt files from the d2 mpq archives. Each data dictionary (txt file) is
// parsed into slices or maps of structs. There is a struct type defined for each txt file.
//
// The RecordManager is a singleton that  loads all of the txt files and export them as
// data members. The RecordManager is meant to be used a a singleton member, exported by the
// AssetManager in  d2core/d2asset.
package d2records
