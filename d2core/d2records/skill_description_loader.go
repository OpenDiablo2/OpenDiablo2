package d2records

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation/d2parser"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// Loadrecords loads skill description records from skilldesc.txt
//nolint:funlen // doesn't make sense to split
func skillDescriptionLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(map[string]*SkillDescriptionRecord)

	parser := d2parser.New()
	parser.SetCurrentReference("skill", "TODO: connect skill with description!")

	for d.Next() {
		record := &SkillDescriptionRecord{
			Name:         d.String("skilldesc"),
			SkillPage:    d.Number("SkillPage"),
			SkillRow:     d.Number("SkillRow"),
			SkillColumn:  d.Number("SkillColumn"),
			ListRow:      d.Number("ListRow"),
			ListPool:     d.String("ListPool"),
			IconCel:      d.Number("IconCel"),
			NameKey:      d.String("str name"),
			ShortKey:     d.String("str short"),
			LongKey:      d.String("str long"),
			AltKey:       d.String("str alt"),
			ManaKey:      d.String("str mana"),
			Descdam:      d.String("descdam"),
			DdamCalc1:    parser.Parse(d.String("ddam calc1")),
			DdamCalc2:    parser.Parse(d.String("ddam calc2")),
			P1dmelem:     d.String("p1dmelem"),
			P1dmmin:      parser.Parse(d.String("p1dmmin")),
			P1dmmax:      parser.Parse(d.String("p1dmmax")),
			P2dmelem:     d.String("p2dmelem"),
			P2dmmin:      parser.Parse(d.String("p2dmmin")),
			P2dmmax:      parser.Parse(d.String("p2dmmax")),
			P3dmelem:     d.String("p3dmelem"),
			P3dmmin:      parser.Parse(d.String("p3dmmin")),
			P3dmmax:      parser.Parse(d.String("p3dmmax")),
			Descatt:      d.String("descatt"),
			Descmissile1: d.String("descmissile1"),
			Descmissile2: d.String("descmissile2"),
			Descmissile3: d.String("descmissile3"),
			Descline1:    d.String("descline1"),
			Desctexta1:   d.String("desctexta1"),
			Desctextb1:   d.String("desctextb1"),
			Desccalca1:   parser.Parse(d.String("desccalca1")),
			Desccalcb1:   parser.Parse(d.String("desccalcb1")),
			Descline2:    d.String("descline2"),
			Desctexta2:   d.String("desctexta2"),
			Desctextb2:   d.String("desctextb2"),
			Desccalca2:   parser.Parse(d.String("desccalca2")),
			Desccalcb2:   parser.Parse(d.String("desccalcb2")),
			Descline3:    d.String("descline3"),
			Desctexta3:   d.String("desctexta3"),
			Desctextb3:   d.String("desctextb3"),
			Desccalca3:   parser.Parse(d.String("desccalca3")),
			Desccalcb3:   parser.Parse(d.String("desccalcb3")),
			Descline4:    d.String("descline4"),
			Desctexta4:   d.String("desctexta4"),
			Desctextb4:   d.String("desctextb4"),
			Desccalca4:   parser.Parse(d.String("desccalca4")),
			Desccalcb4:   parser.Parse(d.String("desccalcb4")),
			Descline5:    d.String("descline5"),
			Desctexta5:   d.String("desctexta5"),
			Desctextb5:   d.String("desctextb5"),
			Desccalca5:   parser.Parse(d.String("desccalca5")),
			Desccalcb5:   parser.Parse(d.String("desccalcb5")),
			Descline6:    d.String("descline6"),
			Desctexta6:   d.String("desctexta6"),
			Desctextb6:   d.String("desctextb6"),
			Desccalca6:   parser.Parse(d.String("desccalca6")),
			Desccalcb6:   parser.Parse(d.String("desccalcb6")),
			Dsc2line1:    d.String("dsc2line1"),
			Dsc2texta1:   d.String("dsc2texta1"),
			Dsc2textb1:   d.String("dsc2textb1"),
			Dsc2calca1:   parser.Parse(d.String("dsc2calca1")),
			Dsc2calcb1:   parser.Parse(d.String("dsc2calcb1")),
			Dsc2line2:    d.String("dsc2line2"),
			Dsc2texta2:   d.String("dsc2texta2"),
			Dsc2textb2:   d.String("dsc2textb2"),
			Dsc2calca2:   parser.Parse(d.String("dsc2calca2")),
			Dsc2calcb2:   parser.Parse(d.String("dsc2calcb2")),
			Dsc2line3:    d.String("dsc2line3"),
			Dsc2texta3:   d.String("dsc2texta3"),
			Dsc2textb3:   d.String("dsc2textb3"),
			Dsc2calca3:   parser.Parse(d.String("dsc2calca3")),
			Dsc2calcb3:   parser.Parse(d.String("dsc2calcb3")),
			Dsc2line4:    d.String("dsc2line4"),
			Dsc2texta4:   d.String("dsc2texta4"),
			Dsc2textb4:   d.String("dsc2textb4"),
			Dsc2calca4:   parser.Parse(d.String("dsc2calca4")),
			Dsc2calcb4:   parser.Parse(d.String("dsc2calcb4")),
			Dsc3line1:    d.String("dsc3line1"),
			Dsc3texta1:   d.String("dsc3texta1"),
			Dsc3textb1:   d.String("dsc3textb1"),
			Dsc3calca1:   parser.Parse(d.String("dsc3calca1")),
			Dsc3calcb1:   parser.Parse(d.String("dsc3calcb1")),
			Dsc3line2:    d.String("dsc3line2"),
			Dsc3texta2:   d.String("dsc3texta2"),
			Dsc3textb2:   d.String("dsc3textb2"),
			Dsc3calca2:   parser.Parse(d.String("dsc3calca2")),
			Dsc3calcb2:   parser.Parse(d.String("dsc3calcb2")),
			Dsc3line3:    d.String("dsc3line3"),
			Dsc3texta3:   d.String("dsc3texta3"),
			Dsc3textb3:   d.String("dsc3textb3"),
			Dsc3calca3:   parser.Parse(d.String("dsc3calca3")),
			Dsc3calcb3:   parser.Parse(d.String("dsc3calcb3")),
			Dsc3line4:    d.String("dsc3line4"),
			Dsc3texta4:   d.String("dsc3texta4"),
			Dsc3textb4:   d.String("dsc3textb4"),
			Dsc3calca4:   parser.Parse(d.String("dsc3calca4")),
			Dsc3calcb4:   parser.Parse(d.String("dsc3calcb4")),
			Dsc3line5:    d.String("dsc3line5"),
			Dsc3texta5:   d.String("dsc3texta5"),
			Dsc3textb5:   d.String("dsc3textb5"),
			Dsc3calca5:   parser.Parse(d.String("dsc3calca5")),
			Dsc3calcb5:   parser.Parse(d.String("dsc3calcb5")),
			Dsc3line6:    d.String("dsc3line6"),
			Dsc3texta6:   d.String("dsc3texta6"),
			Dsc3textb6:   d.String("dsc3textb6"),
			Dsc3calca6:   parser.Parse(d.String("dsc3calca6")),
			Dsc3calcb6:   parser.Parse(d.String("dsc3calcb6")),
			Dsc3line7:    d.String("dsc3line7"),
			Dsc3texta7:   d.String("dsc3texta7"),
			Dsc3textb7:   d.String("dsc3textb7"),
			Dsc3calca7:   parser.Parse(d.String("dsc3calca7")),
			Dsc3calcb7:   parser.Parse(d.String("dsc3calcb7")),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Skill.Descriptions = records

	r.Debugf("Loaded %d SkillDescription records", len(records))

	return nil
}
