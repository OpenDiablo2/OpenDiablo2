package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

func overlaysLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(Overlays)

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

		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d Overlay records", len(records))

	r.Layout.Overlays = records

	return nil
}
