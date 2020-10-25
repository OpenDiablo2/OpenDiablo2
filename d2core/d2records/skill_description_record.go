package d2records

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation"

// SkillDescriptions stores all of the SkillDescriptionRecords
type SkillDescriptions map[string]*SkillDescriptionRecord

// SkillDescriptionRecord is a single row from skilldesc.txt and is used for
// generating text strings for skills.
type SkillDescriptionRecord struct {
	Name         string                    // skilldesc
	SkillPage    int                       // SkillPage
	SkillRow     int                       // SkillRow
	SkillColumn  int                       // SkillColumn
	ListRow      int                       // ListRow
	ListPool     string                    // ListPool
	IconCel      int                       // IconCel
	NameKey      string                    // str name
	ShortKey     string                    // str short
	LongKey      string                    // str long
	AltKey       string                    // str alt
	ManaKey      string                    // str mana
	Descdam      string                    // descdam
	DdamCalc1    d2calculation.Calculation // ddam calc1
	DdamCalc2    d2calculation.Calculation // ddam calc2
	P1dmelem     string                    // p1dmelem
	P1dmmin      d2calculation.Calculation // p1dmmin
	P1dmmax      d2calculation.Calculation // p1dmmax
	P2dmelem     string                    // p2dmelem
	P2dmmin      d2calculation.Calculation // p2dmmin
	P2dmmax      d2calculation.Calculation // p2dmmax
	P3dmelem     string                    // p3dmelem
	P3dmmin      d2calculation.Calculation // p3dmmin
	P3dmmax      d2calculation.Calculation // p3dmmax
	Descatt      string                    // descatt
	Descmissile1 string                    // descmissile1
	Descmissile2 string                    // descmissile2
	Descmissile3 string                    // descmissile3
	Descline1    string                    // descline1
	Desctexta1   string                    // desctexta1
	Desctextb1   string                    // desctextb1
	Desccalca1   d2calculation.Calculation // desccalca1
	Desccalcb1   d2calculation.Calculation // desccalcb1
	Descline2    string                    // descline2
	Desctexta2   string                    // desctexta2
	Desctextb2   string                    // desctextb2
	Desccalca2   d2calculation.Calculation // desccalca2
	Desccalcb2   d2calculation.Calculation // desccalcb2
	Descline3    string                    // descline3
	Desctexta3   string                    // desctexta3
	Desctextb3   string                    // desctextb3
	Desccalca3   d2calculation.Calculation // desccalca3
	Desccalcb3   d2calculation.Calculation // desccalcb3
	Descline4    string                    // descline4
	Desctexta4   string                    // desctexta4
	Desctextb4   string                    // desctextb4
	Desccalca4   d2calculation.Calculation // desccalca4
	Desccalcb4   d2calculation.Calculation // desccalcb4
	Descline5    string                    // descline5
	Desctexta5   string                    // desctexta5
	Desctextb5   string                    // desctextb5
	Desccalca5   d2calculation.Calculation // desccalca5
	Desccalcb5   d2calculation.Calculation // desccalcb5
	Descline6    string                    // descline6
	Desctexta6   string                    // desctexta6
	Desctextb6   string                    // desctextb6
	Desccalca6   d2calculation.Calculation // desccalca6
	Desccalcb6   d2calculation.Calculation // desccalcb6
	Dsc2line1    string                    // dsc2line1
	Dsc2texta1   string                    // dsc2texta1
	Dsc2textb1   string                    // dsc2textb1
	Dsc2calca1   d2calculation.Calculation // dsc2calca1
	Dsc2calcb1   d2calculation.Calculation // dsc2calcb1
	Dsc2line2    string                    // dsc2line2
	Dsc2texta2   string                    // dsc2texta2
	Dsc2textb2   string                    // dsc2textb2
	Dsc2calca2   d2calculation.Calculation // dsc2calca2
	Dsc2calcb2   d2calculation.Calculation // dsc2calcb2
	Dsc2line3    string                    // dsc2line3
	Dsc2texta3   string                    // dsc2texta3
	Dsc2textb3   string                    // dsc2textb3
	Dsc2calca3   d2calculation.Calculation // dsc2calca3
	Dsc2calcb3   d2calculation.Calculation // dsc2calcb3
	Dsc2line4    string                    // dsc2line4
	Dsc2texta4   string                    // dsc2texta4
	Dsc2textb4   string                    // dsc2textb4
	Dsc2calca4   d2calculation.Calculation // dsc2calca4
	Dsc2calcb4   d2calculation.Calculation // dsc2calcb4
	Dsc3line1    string                    // dsc3line1
	Dsc3texta1   string                    // dsc3texta1
	Dsc3textb1   string                    // dsc3textb1
	Dsc3calca1   d2calculation.Calculation // dsc3calca1
	Dsc3calcb1   d2calculation.Calculation // dsc3calcb1
	Dsc3line2    string                    // dsc3line2
	Dsc3texta2   string                    // dsc3texta2
	Dsc3textb2   string                    // dsc3textb2
	Dsc3calca2   d2calculation.Calculation // dsc3calca2
	Dsc3calcb2   d2calculation.Calculation // dsc3calcb2
	Dsc3line3    string                    // dsc3line3
	Dsc3texta3   string                    // dsc3texta3
	Dsc3textb3   string                    // dsc3textb3
	Dsc3calca3   d2calculation.Calculation // dsc3calca3
	Dsc3calcb3   d2calculation.Calculation // dsc3calcb3
	Dsc3line4    string                    // dsc3line4
	Dsc3texta4   string                    // dsc3texta4
	Dsc3textb4   string                    // dsc3textb4
	Dsc3calca4   d2calculation.Calculation // dsc3calca4
	Dsc3calcb4   d2calculation.Calculation // dsc3calcb4
	Dsc3line5    string                    // dsc3line5
	Dsc3texta5   string                    // dsc3texta5
	Dsc3textb5   string                    // dsc3textb5
	Dsc3calca5   d2calculation.Calculation // dsc3calca5
	Dsc3calcb5   d2calculation.Calculation // dsc3calcb5
	Dsc3line6    string                    // dsc3line6
	Dsc3texta6   string                    // dsc3texta6
	Dsc3textb6   string                    // dsc3textb6
	Dsc3calca6   d2calculation.Calculation // dsc3calca6
	Dsc3calcb6   d2calculation.Calculation // dsc3calcb6
	Dsc3line7    string                    // dsc3line7
	Dsc3texta7   string                    // dsc3texta7
	Dsc3textb7   string                    // dsc3textb7
	Dsc3calca7   d2calculation.Calculation // dsc3calca7
	Dsc3calcb7   d2calculation.Calculation // dsc3calcb7
}
