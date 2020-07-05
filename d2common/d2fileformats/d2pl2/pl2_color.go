package d2pl2

// PL2Color represents an RGBA color
type PL2Color struct {
	R uint8
	G uint8
	B uint8
	// padded 0's inside of the pl2 files
	// we keep this here because we use restruct to read
	// the bytes from the files and place into the structs
	_ uint8
}
