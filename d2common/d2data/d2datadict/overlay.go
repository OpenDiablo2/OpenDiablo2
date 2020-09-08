package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// The information has been gathered from [https://d2mods.info/forum/kb/viewarticle?a=465]

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

// Overlays contains all of the OverlayRecords from Overlay.txt
var Overlays map[string]*OverlayRecord // nolint:gochecknoglobals // Currently global by design

// LoadOverlays loads overlay records from Overlay.txt
func LoadOverlays(file []byte) {
	Overlays = make(map[string]*OverlayRecord)
	d := d2txt.LoadDataDictionary(file)

	for d.Next() {
		record := &OverlayRecord{
			Name:       d.String("Overlay"),
			Filename:   d.String("Filename"),
			Version:    d.Bool("Version"),
			PreDraw:    d.Bool("PreDraw"),
			XOffset:    d.Number("Xoffset"),
			YOffset:    d.Number("Yoffset"),
			Height1:    d.Number("Height1"),
			Height2:    d.Number("Height1"),
			Height3:    d.Number("Height1"),
			Height4:    d.Number("Height1"),
			AnimRate:   d.Number("AnimRate"),
			Trans:      d.Number("Trans"),
			InitRadius: d.Number("InitRadius"),
			Radius:     d.Number("Radius"),
			Red:        uint8(d.Number("Red")),
			Green:      uint8(d.Number("Green")),
			Blue:       uint8(d.Number("Blue")),
		}
		Overlays[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Overlay records", len(Overlays))
}
