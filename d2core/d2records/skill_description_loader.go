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
			d.String("skilldesc"),
			d.Number("SkillPage"),
			d.Number("SkillRow"),
			d.Number("SkillColumn"),
			d.Number("ListRow"),
			d.String("ListPool"),
			d.Number("IconCel"),
			d.String("str name"),
			d.String("str short"),
			d.String("str long"),
			d.String("str alt"),
			d.String("str mana"),
			d.String("descdam"),
			parser.Parse(d.String("ddam calc1")),
			parser.Parse(d.String("ddam calc2")),
			d.String("p1dmelem"),
			parser.Parse(d.String("p1dmmin")),
			parser.Parse(d.String("p1dmmax")),
			d.String("p2dmelem"),
			parser.Parse(d.String("p2dmmin")),
			parser.Parse(d.String("p2dmmax")),
			d.String("p3dmelem"),
			parser.Parse(d.String("p3dmmin")),
			parser.Parse(d.String("p3dmmax")),
			d.String("descatt"),
			d.String("descmissile1"),
			d.String("descmissile2"),
			d.String("descmissile3"),
			d.String("descline1"),
			d.String("desctexta1"),
			d.String("desctextb1"),
			parser.Parse(d.String("desccalca1")),
			parser.Parse(d.String("desccalcb1")),
			d.String("descline2"),
			d.String("desctexta2"),
			d.String("desctextb2"),
			parser.Parse(d.String("desccalca2")),
			parser.Parse(d.String("desccalcb2")),
			d.String("descline3"),
			d.String("desctexta3"),
			d.String("desctextb3"),
			parser.Parse(d.String("desccalca3")),
			parser.Parse(d.String("desccalcb3")),
			d.String("descline4"),
			d.String("desctexta4"),
			d.String("desctextb4"),
			parser.Parse(d.String("desccalca4")),
			parser.Parse(d.String("desccalcb4")),
			d.String("descline5"),
			d.String("desctexta5"),
			d.String("desctextb5"),
			parser.Parse(d.String("desccalca5")),
			parser.Parse(d.String("desccalcb5")),
			d.String("descline6"),
			d.String("desctexta6"),
			d.String("desctextb6"),
			parser.Parse(d.String("desccalca6")),
			parser.Parse(d.String("desccalcb6")),
			d.String("dsc2line1"),
			d.String("dsc2texta1"),
			d.String("dsc2textb1"),
			parser.Parse(d.String("dsc2calca1")),
			parser.Parse(d.String("dsc2calcb1")),
			d.String("dsc2line2"),
			d.String("dsc2texta2"),
			d.String("dsc2textb2"),
			parser.Parse(d.String("dsc2calca2")),
			parser.Parse(d.String("dsc2calcb2")),
			d.String("dsc2line3"),
			d.String("dsc2texta3"),
			d.String("dsc2textb3"),
			parser.Parse(d.String("dsc2calca3")),
			parser.Parse(d.String("dsc2calcb3")),
			d.String("dsc2line4"),
			d.String("dsc2texta4"),
			d.String("dsc2textb4"),
			parser.Parse(d.String("dsc2calca4")),
			parser.Parse(d.String("dsc2calcb4")),
			d.String("dsc3line1"),
			d.String("dsc3texta1"),
			d.String("dsc3textb1"),
			parser.Parse(d.String("dsc3calca1")),
			parser.Parse(d.String("dsc3calcb1")),
			d.String("dsc3line2"),
			d.String("dsc3texta2"),
			d.String("dsc3textb2"),
			parser.Parse(d.String("dsc3calca2")),
			parser.Parse(d.String("dsc3calcb2")),
			d.String("dsc3line3"),
			d.String("dsc3texta3"),
			d.String("dsc3textb3"),
			parser.Parse(d.String("dsc3calca3")),
			parser.Parse(d.String("dsc3calcb3")),
			d.String("dsc3line4"),
			d.String("dsc3texta4"),
			d.String("dsc3textb4"),
			parser.Parse(d.String("dsc3calca4")),
			parser.Parse(d.String("dsc3calcb4")),
			d.String("dsc3line5"),
			d.String("dsc3texta5"),
			d.String("dsc3textb5"),
			parser.Parse(d.String("dsc3calca5")),
			parser.Parse(d.String("dsc3calcb5")),
			d.String("dsc3line6"),
			d.String("dsc3texta6"),
			d.String("dsc3textb6"),
			parser.Parse(d.String("dsc3calca6")),
			parser.Parse(d.String("dsc3calcb6")),
			d.String("dsc3line7"),
			d.String("dsc3texta7"),
			d.String("dsc3textb7"),
			parser.Parse(d.String("dsc3calca7")),
			parser.Parse(d.String("dsc3calcb7")),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Skill.Descriptions = records

	r.Logger.Infof("Loaded %d Skill Description records", len(records))

	return nil
}
