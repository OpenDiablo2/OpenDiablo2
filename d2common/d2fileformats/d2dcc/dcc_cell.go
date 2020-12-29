package d2dcc

// DCCCell represents a single cell in a DCC file.
type DCCCell struct {
	Width       int
	Height      int
	XOffset     int
	YOffset     int
	LastWidth   int
	LastHeight  int
	LastXOffset int
	LastYOffset int
}
