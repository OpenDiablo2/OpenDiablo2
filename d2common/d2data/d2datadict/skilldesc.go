package d2datadict

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"log"
)

type SkillDescriptionRecord struct {
	Name         string // skilldesc
	SkillPage    string // SkillPage
	SkillRow     string // SkillRow
	SkillColumn  string // SkillColumn
	ListRow      string // ListRow
	ListPool     string // ListPool
	IconCel      string // IconCel
	NameKey      string // str name
	ShortKey     string // str short
	LongKey      string // str long
	AltKey       string // str alt
	ManaKey      string // str mana
	Descdam      string // descdam
	DdamCalc1    string // ddam calc1
	DdamCalc2    string // ddam calc2
	P1dmelem     string // p1dmelem
	P1dmmin      string // p1dmmin
	P1dmmax      string // p1dmmax
	P2dmelem     string // p2dmelem
	P2dmmin      string // p2dmmin
	P2dmmax      string // p2dmmax
	P3dmelem     string // p3dmelem
	P3dmmin      string // p3dmmin
	P3dmmax      string // p3dmmax
	Descatt      string // descatt
	Descmissile1 string // descmissile1
	Descmissile2 string // descmissile2
	Descmissile3 string // descmissile3
	Descline1    string // descline1
	Desctexta1   string // desctexta1
	Desctextb1   string // desctextb1
	Desccalca1   string // desccalca1
	Desccalcb1   string // desccalcb1
	Descline2    string // descline2
	Desctexta2   string // desctexta2
	Desctextb2   string // desctextb2
	Desccalca2   string // desccalca2
	Desccalcb2   string // desccalcb2
	Descline3    string // descline3
	Desctexta3   string // desctexta3
	Desctextb3   string // desctextb3
	Desccalca3   string // desccalca3
	Desccalcb3   string // desccalcb3
	Descline4    string // descline4
	Desctexta4   string // desctexta4
	Desctextb4   string // desctextb4
	Desccalca4   string // desccalca4
	Desccalcb4   string // desccalcb4
	Descline5    string // descline5
	Desctexta5   string // desctexta5
	Desctextb5   string // desctextb5
	Desccalca5   string // desccalca5
	Desccalcb5   string // desccalcb5
	Descline6    string // descline6
	Desctexta6   string // desctexta6
	Desctextb6   string // desctextb6
	Desccalca6   string // desccalca6
	Desccalcb6   string // desccalcb6
	Dsc2line1    string // dsc2line1
	Dsc2texta1   string // dsc2texta1
	Dsc2textb1   string // dsc2textb1
	Dsc2calca1   string // dsc2calca1
	Dsc2calcb1   string // dsc2calcb1
	Dsc2line2    string // dsc2line2
	Dsc2texta2   string // dsc2texta2
	Dsc2textb2   string // dsc2textb2
	Dsc2calca2   string // dsc2calca2
	Dsc2calcb2   string // dsc2calcb2
	Dsc2line3    string // dsc2line3
	Dsc2texta3   string // dsc2texta3
	Dsc2textb3   string // dsc2textb3
	Dsc2calca3   string // dsc2calca3
	Dsc2calcb3   string // dsc2calcb3
	Dsc2line4    string // dsc2line4
	Dsc2texta4   string // dsc2texta4
	Dsc2textb4   string // dsc2textb4
	Dsc2calca4   string // dsc2calca4
	Dsc2calcb4   string // dsc2calcb4
	Dsc3line1    string // dsc3line1
	Dsc3texta1   string // dsc3texta1
	Dsc3textb1   string // dsc3textb1
	Dsc3calca1   string // dsc3calca1
	Dsc3calcb1   string // dsc3calcb1
	Dsc3line2    string // dsc3line2
	Dsc3texta2   string // dsc3texta2
	Dsc3textb2   string // dsc3textb2
	Dsc3calca2   string // dsc3calca2
	Dsc3calcb2   string // dsc3calcb2
	Dsc3line3    string // dsc3line3
	Dsc3texta3   string // dsc3texta3
	Dsc3textb3   string // dsc3textb3
	Dsc3calca3   string // dsc3calca3
	Dsc3calcb3   string // dsc3calcb3
	Dsc3line4    string // dsc3line4
	Dsc3texta4   string // dsc3texta4
	Dsc3textb4   string // dsc3textb4
	Dsc3calca4   string // dsc3calca4
	Dsc3calcb4   string // dsc3calcb4
	Dsc3line5    string // dsc3line5
	Dsc3texta5   string // dsc3texta5
	Dsc3textb5   string // dsc3textb5
	Dsc3calca5   string // dsc3calca5
	Dsc3calcb5   string // dsc3calcb5
	Dsc3line6    string // dsc3line6
	Dsc3texta6   string // dsc3texta6
	Dsc3textb6   string // dsc3textb6
	Dsc3calca6   string // dsc3calca6
	Dsc3calcb6   string // dsc3calcb6
	Dsc3line7    string // dsc3line7
	Dsc3texta7   string // dsc3texta7
	Dsc3textb7   string // dsc3textb7
	Dsc3calca7   string // dsc3calca7
	Dsc3calcb7   string // dsc3calcb7
}

// ItemStatCosts stores all of the ItemStatCostRecords
//nolint:gochecknoglobals // Currently global by design
var SkillDescriptions map[string]*SkillDescriptionRecord

// LoadItemStatCosts loads ItemStatCostRecord's from text
func LoadSkillDescriptions(file []byte) {
	SkillDescriptions = make(map[string]*SkillDescriptionRecord)

	d := d2common.LoadDataDictionary(file)
	for d.Next() {
		record := &SkillDescriptionRecord{
			d.String("skilldesc"),
			d.String("SkillPage"),
			d.String("SkillRow"),
			d.String("SkillColumn"),
			d.String("ListRow"),
			d.String("ListPool"),
			d.String("IconCel"),
			d.String("str name"),
			d.String("str short"),
			d.String("str long"),
			d.String("str alt"),
			d.String("str mana"),
			d.String("descdam"),
			d.String("ddam calc1"),
			d.String("ddam calc2"),
			d.String("p1dmelem"),
			d.String("p1dmmin"),
			d.String("p1dmmax"),
			d.String("p2dmelem"),
			d.String("p2dmmin"),
			d.String("p2dmmax"),
			d.String("p3dmelem"),
			d.String("p3dmmin"),
			d.String("p3dmmax"),
			d.String("descatt"),
			d.String("descmissile1"),
			d.String("descmissile2"),
			d.String("descmissile3"),
			d.String("descline1"),
			d.String("desctexta1"),
			d.String("desctextb1"),
			d.String("desccalca1"),
			d.String("desccalcb1"),
			d.String("descline2"),
			d.String("desctexta2"),
			d.String("desctextb2"),
			d.String("desccalca2"),
			d.String("desccalcb2"),
			d.String("descline3"),
			d.String("desctexta3"),
			d.String("desctextb3"),
			d.String("desccalca3"),
			d.String("desccalcb3"),
			d.String("descline4"),
			d.String("desctexta4"),
			d.String("desctextb4"),
			d.String("desccalca4"),
			d.String("desccalcb4"),
			d.String("descline5"),
			d.String("desctexta5"),
			d.String("desctextb5"),
			d.String("desccalca5"),
			d.String("desccalcb5"),
			d.String("descline6"),
			d.String("desctexta6"),
			d.String("desctextb6"),
			d.String("desccalca6"),
			d.String("desccalcb6"),
			d.String("dsc2line1"),
			d.String("dsc2texta1"),
			d.String("dsc2textb1"),
			d.String("dsc2calca1"),
			d.String("dsc2calcb1"),
			d.String("dsc2line2"),
			d.String("dsc2texta2"),
			d.String("dsc2textb2"),
			d.String("dsc2calca2"),
			d.String("dsc2calcb2"),
			d.String("dsc2line3"),
			d.String("dsc2texta3"),
			d.String("dsc2textb3"),
			d.String("dsc2calca3"),
			d.String("dsc2calcb3"),
			d.String("dsc2line4"),
			d.String("dsc2texta4"),
			d.String("dsc2textb4"),
			d.String("dsc2calca4"),
			d.String("dsc2calcb4"),
			d.String("dsc3line1"),
			d.String("dsc3texta1"),
			d.String("dsc3textb1"),
			d.String("dsc3calca1"),
			d.String("dsc3calcb1"),
			d.String("dsc3line2"),
			d.String("dsc3texta2"),
			d.String("dsc3textb2"),
			d.String("dsc3calca2"),
			d.String("dsc3calcb2"),
			d.String("dsc3line3"),
			d.String("dsc3texta3"),
			d.String("dsc3textb3"),
			d.String("dsc3calca3"),
			d.String("dsc3calcb3"),
			d.String("dsc3line4"),
			d.String("dsc3texta4"),
			d.String("dsc3textb4"),
			d.String("dsc3calca4"),
			d.String("dsc3calcb4"),
			d.String("dsc3line5"),
			d.String("dsc3texta5"),
			d.String("dsc3textb5"),
			d.String("dsc3calca5"),
			d.String("dsc3calcb5"),
			d.String("dsc3line6"),
			d.String("dsc3texta6"),
			d.String("dsc3textb6"),
			d.String("dsc3calca6"),
			d.String("dsc3calcb6"),
			d.String("dsc3line7"),
			d.String("dsc3texta7"),
			d.String("dsc3textb7"),
			d.String("dsc3calca7"),
			d.String("dsc3calcb7"),
		}

		SkillDescriptions[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	log.Printf("Loaded %d Skill Description records", len(SkillDescriptions))
}
