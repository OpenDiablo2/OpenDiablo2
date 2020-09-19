package d2records

// The information has been gathered from [https://d2mods.info/forum/kb/viewarticle?a=465]

// Overlays contains all of the OverlayRecords from Overlay.txt
type Overlays map[string]*OverlayRecord

// OverlayRecord encapsulates information found in Overlay.txt
type OverlayRecord struct {
	// Overlay name
	Name string

	// .dcc file found in Data/Globals/Overlays
	Filename string

	// XOffset, YOffset the x,y offset of the overlay
	XOffset, YOffset int

	// These values modify Y-axis placement
	Height1 int
	Height2 int
	Height3 int
	Height4 int

	// AnimRate animation speed control
	AnimRate int

	// Trans controls overlay blending mode, check out the link for more info
	// This should probably become an "enum" later on
	Trans int

	// Radius maximum for light
	Radius int

	// InitRadius Light radius increase per frame
	InitRadius int

	// Red, Green, Blue color for light
	Red, Green, Blue uint8

	// Version is 100 for expansion, 0 for vanilla
	Version bool

	// PreDraw controls overlay drawing precedence
	PreDraw bool

	// Unknown fields, commenting out for now
	// NumDirections int
	// LocalBlood    int
	// OneOfN int
	// Dir  bool
	// Open bool
	// Beta bool

	// Apparently unused
	// Character string
	// LoopWaitTime int
	// Frames int
}
