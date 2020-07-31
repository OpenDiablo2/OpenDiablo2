package d2datadict

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
)

// OverlayRecord encapsulates information found in Overlay.txt
// The information has been gathered from [https://d2mods.info/forum/kb/viewarticle?a=465]
type OverlayRecord struct {
	// Overlay name
	Overlay string
	// .dcc file found in Data/Globals/Overlays
	Filename string
	// Apparently unused, a similar field in the .dcc file is used instead
	Frames int
	// Unused
	Character string
	// Controls overlay drawing precedence
	PreDraw bool
	// Unknown
	OneOfN int

	// Unknown
	Dir  bool
	Open bool
	Beta bool

	XOffset int
	YOffset int

	// These values modify Y-axis placement
	Height1 int
	Height2 int
	Height3 int
	Height4 int

	// Animation speed control
	AnimRate int
	// Unused
	LoopWaitTime int
	// Controls overlay blending mode, check out the link for more info
	// This should probably become an "enum" later on
	Trans int
	// Maximum light radius
	Radius int
	// Light radius increase per frame
	InitRadius int

	Red   uint8
	Green uint8
	Blue  uint8

	// Unknown
	NumDirections int
	LocalBlood    int
}

var Overlays map[string]*OverlayRecord

func LoadOverlays(file []byte) {
	Overlays = make(map[string]*OverlayRecord)
	d := d2common.LoadDataDictionary(file)

	for d.Next() {
		record := &OverlayRecord{
			Overlay:      d.String("Overlay"),
			Filename:     d.String("Filename"),
			Frames:       d.Number("Frames"),
			Character:    d.String("Character"),
			PreDraw:      d.Bool("PreDraw"),
			OneOfN:       d.Number("1ofN"),
			Dir:          d.Bool("Dir"),
			Open:         d.Bool("Open"),
			Beta:         d.Bool("Beta"),
			XOffset:      d.Number("Xoffset"),
			YOffset:      d.Number("Yoffset"),
			Height1:      d.Number("Height1"),
			Height2:      d.Number("Height1"),
			Height3:      d.Number("Height1"),
			Height4:      d.Number("Height1"),
			AnimRate:     d.Number("AnimRate"),
			LoopWaitTime: d.Number("LoopWaitTime"),
			Trans:        d.Number("Trans"),
			InitRadius:   d.Number("InitRadius"),
			Radius:       d.Number("Radius"),
			Red:          uint8(d.Number("Red")),
			Green:        uint8(d.Number("Green")),
			Blue:         uint8(d.Number("Blue")),
			LocalBlood:   d.Number("LocalBlood"),
		}
		Overlays[record.Overlay] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Overlay records", len(Overlays))

}
