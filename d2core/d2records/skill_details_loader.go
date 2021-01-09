package d2records

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2calculation/d2parser"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2fileformats/d2txt"
)

// LoadSkills loads skills.txt file contents into a skill record map
//nolint:funlen // Makes no sense to split
func skillDetailsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(map[int]*SkillRecord)

	parser := d2parser.New()

	for d.Next() {
		name := d.String("skill")
		parser.SetCurrentReference("skill", name)

		anim, err := animToEnum(d.String("anim"))
		if err != nil {
			return err
		}

		record := &SkillRecord{
			Skill:             d.String("skill"),
			ID:                d.Number("Id"),
			Charclass:         d.String("charclass"),
			Skilldesc:         d.String("skilldesc"),
			Srvstfunc:         d.Number("srvstfunc"),
			Srvdofunc:         d.Number("srvdofunc"),
			Prgstack:          d.Bool("prgstack"),
			Srvprgfunc1:       d.Number("srvprgfunc1"),
			Srvprgfunc2:       d.Number("srvprgfunc2"),
			Srvprgfunc3:       d.Number("srvprgfunc3"),
			Prgcalc1:          parser.Parse(d.String("prgcalc1")),
			Prgcalc2:          parser.Parse(d.String("prgcalc2")),
			Prgcalc3:          parser.Parse(d.String("prgcalc3")),
			Prgdam:            d.Number("prgdam"),
			Srvmissile:        d.String("srvmissile"),
			Decquant:          d.Bool("decquant"),
			Lob:               d.Bool("lob"),
			Srvmissilea:       d.String("srvmissilea"),
			Srvmissileb:       d.String("srvmissileb"),
			Srvmissilec:       d.String("srvmissilec"),
			Srvoverlay:        d.String("srvoverlay"),
			Aurafilter:        d.Number("aurafilter"),
			Aurastate:         d.String("aurastate"),
			Auratargetstate:   d.String("auratargetstate"),
			Auralencalc:       parser.Parse(d.String("auralencalc")),
			Aurarangecalc:     parser.Parse(d.String("aurarangecalc")),
			Aurastat1:         d.String("aurastat1"),
			Aurastatcalc1:     parser.Parse(d.String("aurastatcalc1")),
			Aurastat2:         d.String("aurastat2"),
			Aurastatcalc2:     parser.Parse(d.String("aurastatcalc2")),
			Aurastat3:         d.String("aurastat3"),
			Aurastatcalc3:     parser.Parse(d.String("aurastatcalc3")),
			Aurastat4:         d.String("aurastat4"),
			Aurastatcalc4:     parser.Parse(d.String("aurastatcalc4")),
			Aurastat5:         d.String("aurastat5"),
			Aurastatcalc5:     parser.Parse(d.String("aurastatcalc5")),
			Aurastat6:         d.String("aurastat6"),
			Aurastatcalc6:     parser.Parse(d.String("aurastatcalc6")),
			Auraevent1:        d.String("auraevent1"),
			Auraeventfunc1:    d.Number("auraeventfunc1"),
			Auraevent2:        d.String("auraevent2"),
			Auraeventfunc2:    d.Number("auraeventfunc2"),
			Auraevent3:        d.String("auraevent3"),
			Auraeventfunc3:    d.Number("auraeventfunc3"),
			Auratgtevent:      d.String("auratgtevent"),
			Auratgteventfunc:  d.String("auratgteventfunc"),
			Passivestate:      d.String("passivestate"),
			Passiveitype:      d.String("passiveitype"),
			Passivestat1:      d.String("passivestat1"),
			Passivecalc1:      parser.Parse(d.String("passivecalc1")),
			Passivestat2:      d.String("passivestat2"),
			Passivecalc2:      parser.Parse(d.String("passivecalc2")),
			Passivestat3:      d.String("passivestat3"),
			Passivecalc3:      parser.Parse(d.String("passivecalc3")),
			Passivestat4:      d.String("passivestat4"),
			Passivecalc4:      parser.Parse(d.String("passivecalc4")),
			Passivestat5:      d.String("passivestat5"),
			Passivecalc5:      parser.Parse(d.String("passivecalc5")),
			Passiveevent:      d.String("passiveevent"),
			Passiveeventfunc:  d.String("passiveeventfunc"),
			Summon:            d.String("summon"),
			Pettype:           d.String("pettype"),
			Petmax:            parser.Parse(d.String("petmax")),
			Summode:           d.String("summode"),
			Sumskill1:         d.String("sumskill1"),
			Sumsk1calc:        parser.Parse(d.String("sumsk1calc")),
			Sumskill2:         d.String("sumskill2"),
			Sumsk2calc:        parser.Parse(d.String("sumsk2calc")),
			Sumskill3:         d.String("sumskill3"),
			Sumsk3calc:        parser.Parse(d.String("sumsk3calc")),
			Sumskill4:         d.String("sumskill4"),
			Sumsk4calc:        parser.Parse(d.String("sumsk4calc")),
			Sumskill5:         d.String("sumskill5"),
			Sumsk5calc:        parser.Parse(d.String("sumsk5calc")),
			Sumumod:           d.Number("sumumod"),
			Sumoverlay:        d.String("sumoverlay"),
			Stsuccessonly:     d.Bool("stsuccessonly"),
			Stsound:           d.String("stsound"),
			Stsoundclass:      d.String("stsoundclass"),
			Stsounddelay:      d.Bool("stsounddelay"),
			Weaponsnd:         d.Bool("weaponsnd"),
			Dosound:           d.String("dosound"),
			DosoundA:          d.String("dosound a"),
			DosoundB:          d.String("dosound b"),
			Tgtoverlay:        d.String("tgtoverlay"),
			Tgtsound:          d.String("tgtsound"),
			Prgoverlay:        d.String("prgoverlay"),
			Prgsound:          d.String("prgsound"),
			Castoverlay:       d.String("castoverlay"),
			Cltoverlaya:       d.String("cltoverlaya"),
			Cltoverlayb:       d.String("cltoverlayb"),
			Cltstfunc:         d.Number("cltstfunc"),
			Cltdofunc:         d.Number("cltdofunc"),
			Cltprgfunc1:       d.Number("cltprgfunc1"),
			Cltprgfunc2:       d.Number("cltprgfunc2"),
			Cltprgfunc3:       d.Number("cltprgfunc3"),
			Cltmissile:        d.String("cltmissile"),
			Cltmissilea:       d.String("cltmissilea"),
			Cltmissileb:       d.String("cltmissileb"),
			Cltmissilec:       d.String("cltmissilec"),
			Cltmissiled:       d.String("cltmissiled"),
			Cltcalc1:          parser.Parse(d.String("cltcalc1")),
			Cltcalc2:          parser.Parse(d.String("cltcalc2")),
			Cltcalc3:          parser.Parse(d.String("cltcalc3")),
			Warp:              d.Bool("warp"),
			Immediate:         d.Bool("immediate"),
			Enhanceable:       d.Bool("enhanceable"),
			Attackrank:        d.Number("attackrank"),
			Noammo:            d.Bool("noammo"),
			Range:             d.String("range"),
			Weapsel:           d.Number("weapsel"),
			Itypea1:           d.String("itypea1"),
			Itypea2:           d.String("itypea2"),
			Itypea3:           d.String("itypea3"),
			Etypea1:           d.String("etypea1"),
			Etypea2:           d.String("etypea2"),
			Itypeb1:           d.String("itypeb1"),
			Itypeb2:           d.String("itypeb2"),
			Itypeb3:           d.String("itypeb3"),
			Etypeb1:           d.String("etypeb1"),
			Etypeb2:           d.String("etypeb2"),
			Anim:              anim,
			Seqtrans:          d.String("seqtrans"),
			Monanim:           d.String("monanim"),
			Seqnum:            d.Number("seqnum"),
			Seqinput:          d.Number("seqinput"),
			Durability:        d.Bool("durability"),
			UseAttackRate:     d.Bool("UseAttackRate"),
			LineOfSight:       d.Number("LineOfSight"),
			TargetableOnly:    d.Bool("TargetableOnly"),
			SearchEnemyXY:     d.Bool("SearchEnemyXY"),
			SearchEnemyNear:   d.Bool("SearchEnemyNear"),
			SearchOpenXY:      d.Bool("SearchOpenXY"),
			SelectProc:        d.Number("SelectProc"),
			TargetCorpse:      d.Bool("TargetCorpse"),
			TargetPet:         d.Bool("TargetPet"),
			TargetAlly:        d.Bool("TargetAlly"),
			TargetItem:        d.Bool("TargetItem"),
			AttackNoMana:      d.Bool("AttackNoMana"),
			TgtPlaceCheck:     d.Bool("TgtPlaceCheck"),
			ItemEffect:        d.Number("ItemEffect"),
			ItemCltEffect:     d.Number("ItemCltEffect"),
			ItemTgtDo:         d.Number("ItemTgtDo"),
			ItemTarget:        d.Number("ItemTarget"),
			ItemCheckStart:    d.Bool("ItemCheckStart"),
			ItemCltCheckStart: d.Bool("ItemCltCheckStart"),
			ItemCastSound:     d.String("ItemCastSound"),
			ItemCastOverlay:   d.String("ItemCastOverlay"),
			Skpoints:          parser.Parse(d.String("skpoints")),
			Reqlevel:          d.Number("reqlevel"),
			Maxlvl:            d.Number("maxlvl"),
			Reqstr:            d.Number("reqstr"),
			Reqdex:            d.Number("reqdex"),
			Reqint:            d.Number("reqint"),
			Reqvit:            d.Number("reqvit"),
			Reqskill1:         d.String("reqskill1"),
			Reqskill2:         d.String("reqskill2"),
			Reqskill3:         d.String("reqskill3"),
			Restrict:          d.Number("restrict"),
			State1:            d.String("State1"),
			State2:            d.String("State2"),
			State3:            d.String("State3"),
			Delay:             d.Number("delay"),
			Leftskill:         d.Bool("leftskill"),
			Repeat:            d.Bool("repeat"),
			Checkfunc:         d.Number("checkfunc"),
			Nocostinstate:     d.Bool("nocostinstate"),
			Usemanaondo:       d.Bool("usemanaondo"),
			Startmana:         d.Number("startmana"),
			Minmana:           d.Number("minmana"),
			Manashift:         d.Number("manashift"),
			Mana:              d.Number("mana"),
			Lvlmana:           d.Number("lvlmana"),
			Interrupt:         d.Bool("interrupt"),
			InTown:            d.Bool("InTown"),
			Aura:              d.Bool("aura"),
			Periodic:          d.Bool("periodic"),
			Perdelay:          parser.Parse(d.String("perdelay")),
			Finishing:         d.Bool("finishing"),
			Passive:           d.Bool("passive"),
			Progressive:       d.Bool("progressive"),
			General:           d.Bool("general"),
			Scroll:            d.Bool("scroll"),
			Calc1:             parser.Parse(d.String("calc1")),
			Calc2:             parser.Parse(d.String("calc2")),
			Calc3:             parser.Parse(d.String("calc3")),
			Calc4:             parser.Parse(d.String("calc4")),
			Param1:            d.Number("Param1"),
			Param2:            d.Number("Param2"),
			Param3:            d.Number("Param3"),
			Param4:            d.Number("Param4"),
			Param5:            d.Number("Param5"),
			Param6:            d.Number("Param6"),
			Param7:            d.Number("Param7"),
			Param8:            d.Number("Param8"),
			InGame:            d.Bool("InGame"),
			ToHit:             d.Number("ToHit"),
			LevToHit:          d.Number("LevToHit"),
			ToHitCalc:         parser.Parse(d.String("ToHitCalc")),
			ResultFlags:       d.Number("ResultFlags"),
			HitFlags:          d.Number("HitFlags"),
			HitClass:          d.Number("HitClass"),
			Kick:              d.Bool("Kick"),
			HitShift:          d.Number("HitShift"),
			SrcDam:            d.Number("SrcDam"),
			MinDam:            d.Number("MinDam"),
			MinLevDam1:        d.Number("MinLevDam1"),
			MinLevDam2:        d.Number("MinLevDam2"),
			MinLevDam3:        d.Number("MinLevDam3"),
			MinLevDam4:        d.Number("MinLevDam4"),
			MinLevDam5:        d.Number("MinLevDam5"),
			MaxDam:            d.Number("MaxDam"),
			MaxLevDam1:        d.Number("MaxLevDam1"),
			MaxLevDam2:        d.Number("MaxLevDam2"),
			MaxLevDam3:        d.Number("MaxLevDam3"),
			MaxLevDam4:        d.Number("MaxLevDam4"),
			MaxLevDam5:        d.Number("MaxLevDam5"),
			DmgSymPerCalc:     parser.Parse(d.String("DmgSymPerCalc")),
			EType:             d.String("EType"),
			EMin:              d.Number("EMin"),
			EMinLev1:          d.Number("EMinLev1"),
			EMinLev2:          d.Number("EMinLev2"),
			EMinLev3:          d.Number("EMinLev3"),
			EMinLev4:          d.Number("EMinLev4"),
			EMinLev5:          d.Number("EMinLev5"),
			EMax:              d.Number("EMax"),
			EMaxLev1:          d.Number("EMaxLev1"),
			EMaxLev2:          d.Number("EMaxLev2"),
			EMaxLev3:          d.Number("EMaxLev3"),
			EMaxLev4:          d.Number("EMaxLev4"),
			EMaxLev5:          d.Number("EMaxLev5"),
			EDmgSymPerCalc:    parser.Parse(d.String("EDmgSymPerCalc")),
			ELen:              d.Number("ELen"),
			ELevLen1:          d.Number("ELevLen1"),
			ELevLen2:          d.Number("ELevLen2"),
			ELevLen3:          d.Number("ELevLen3"),
			ELenSymPerCalc:    parser.Parse(d.String("ELenSymPerCalc")),
			Aitype:            d.Number("aitype"),
			Aibonus:           d.Number("aibonus"),
			CostMult:          d.Number("cost mult"),
			CostAdd:           d.Number("cost add"),
		}

		records[record.ID] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Skill.Details = records

	r.Logger.Infof("Loaded %d Skill records", len(records))

	return nil
}

func animToEnum(anim string) (d2enum.PlayerAnimationMode, error) {
	switch anim {
	case "SC":
		return d2enum.PlayerAnimationModeCast, nil

	case "TH":
		return d2enum.PlayerAnimationModeThrow, nil

	case "KK":
		return d2enum.PlayerAnimationModeKick, nil

	case "SQ":
		return d2enum.PlayerAnimationModeSequence, nil

	case "S1":
		return d2enum.PlayerAnimationModeSkill1, nil

	case "S2":
		return d2enum.PlayerAnimationModeSkill2, nil

	case "S3":
		return d2enum.PlayerAnimationModeSkill3, nil

	case "S4":
		return d2enum.PlayerAnimationModeSkill4, nil

	case "A1":
		return d2enum.PlayerAnimationModeAttack1, nil

	case "A2":
		return d2enum.PlayerAnimationModeAttack2, nil

	case "":
		return d2enum.PlayerAnimationModeNone, nil
	}

	// should not be reached
	return d2enum.PlayerAnimationModeNone, fmt.Errorf("unknown skill anim value [%s]", anim)
}
