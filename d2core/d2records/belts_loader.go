package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// nolint:funlen // cant reduce
func beltsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(Belts)

	for d.Next() {
		record := &BeltRecord{
			Name:      d.String("name"),
			NumBoxes:  d.Number("numboxes"),
			BoxWidth:  d.Number("boxwidth"),
			BoxHeight: d.Number("boxheight"),

			Box1Left:   d.Number("box1left"),
			Box1Right:  d.Number("box1right"),
			Box1Top:    d.Number("box1top"),
			Box1Bottom: d.Number("box1bottom"),

			Box2Left:   d.Number("box2left"),
			Box2Right:  d.Number("box2right"),
			Box2Top:    d.Number("box2top"),
			Box2Bottom: d.Number("box2bottom"),

			Box3Left:   d.Number("box3left"),
			Box3Right:  d.Number("box3right"),
			Box3Top:    d.Number("box3top"),
			Box3Bottom: d.Number("box3bottom"),

			Box4Left:   d.Number("box4left"),
			Box4Right:  d.Number("box4right"),
			Box4Top:    d.Number("box4top"),
			Box4Bottom: d.Number("box4bottom"),

			Box5Left:   d.Number("box5left"),
			Box5Right:  d.Number("box5right"),
			Box5Top:    d.Number("box5top"),
			Box5Bottom: d.Number("box5bottom"),

			Box6Left:   d.Number("box6left"),
			Box6Right:  d.Number("box6right"),
			Box6Top:    d.Number("box6top"),
			Box6Bottom: d.Number("box6bottom"),

			Box7Left:   d.Number("box7left"),
			Box7Right:  d.Number("box7right"),
			Box7Top:    d.Number("box7top"),
			Box7Bottom: d.Number("box7bottom"),

			Box8Left:   d.Number("box8left"),
			Box8Right:  d.Number("box8right"),
			Box8Top:    d.Number("box8top"),
			Box8Bottom: d.Number("box8bottom"),

			Box9Left:   d.Number("box9left"),
			Box9Right:  d.Number("box9right"),
			Box9Top:    d.Number("box9top"),
			Box9Bottom: d.Number("box9bottom"),

			Box10Left:   d.Number("box10left"),
			Box10Right:  d.Number("box10right"),
			Box10Top:    d.Number("box10top"),
			Box10Bottom: d.Number("box10bottom"),

			Box11Left:   d.Number("box11left"),
			Box11Right:  d.Number("box11right"),
			Box11Top:    d.Number("box11top"),
			Box11Bottom: d.Number("box11bottom"),

			Box12Left:   d.Number("box12left"),
			Box12Right:  d.Number("box12right"),
			Box12Top:    d.Number("box12top"),
			Box12Bottom: d.Number("box12bottom"),

			Box13Left:   d.Number("box13left"),
			Box13Right:  d.Number("box13right"),
			Box13Top:    d.Number("box13top"),
			Box13Bottom: d.Number("box13bottom"),

			Box14Left:   d.Number("box14left"),
			Box14Right:  d.Number("box14right"),
			Box14Top:    d.Number("box14top"),
			Box14Bottom: d.Number("box14bottom"),

			Box15Left:   d.Number("box15left"),
			Box15Right:  d.Number("box15right"),
			Box15Top:    d.Number("box15top"),
			Box15Bottom: d.Number("box15bottom"),

			Box16Left:   d.Number("box16left"),
			Box16Right:  d.Number("box16right"),
			Box16Top:    d.Number("box16top"),
			Box16Bottom: d.Number("box16bottom"),
		}
		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Logger.Infof("Loaded %d belts", len(records))

	r.Item.Belts = records

	return nil
}
